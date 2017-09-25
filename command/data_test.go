//
// command/data_test.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
//
// This file is part of Roadie.
//
// Roadie is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Roadie is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Roadie.  If not, see <http://www.gnu.org/licenses/>.
//

package command

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
)

// locationURLs parses a multi-lines text and retuens a map from file names to
// stored locations of them.
// The multi-lines test should be formatted in
// > file1.txt -> roadie://data/file1.txt
// > file2.dat -> roadie://data/file2.data
// The last line may containes new line.
func locationURLs(s string) (res map[string]string) {

	res = make(map[string]string)
	for _, line := range strings.Split(s, "\n") {
		pair := strings.Split(line, "->")
		if len(pair) == 2 {
			file := strings.TrimSpace(pair[0])
			loc := strings.TrimSpace(pair[1])
			res[file] = loc
		}
	}
	return

}

// TestLocationURLs tasts locationURLs.
func TestLocationURLs(t *testing.T) {

	cases := []struct {
		input  string
		expect map[string]string
	}{
		{"file1.txt -> roadie://data/file1.txt", map[string]string{"file1.txt": "roadie://data/file1.txt"}},
		{"file1.txt -> roadie://data/file1.txt\n", map[string]string{"file1.txt": "roadie://data/file1.txt"}},
		{"\nfile1.txt -> roadie://data/file1.txt", map[string]string{"file1.txt": "roadie://data/file1.txt"}},
		{"\nfile1.txt -> roadie://data/file1.txt\n", map[string]string{"file1.txt": "roadie://data/file1.txt"}},
		{"abc\nfile1.txt -> roadie://data/file1.txt", map[string]string{"file1.txt": "roadie://data/file1.txt"}},
		{"abc\nfile1.txt -> roadie://data/file1.txt\n", map[string]string{"file1.txt": "roadie://data/file1.txt"}},
		{
			"file1.txt -> roadie://data/file1.txt\nfile2.txt -> roadie://data/file2.txt",
			map[string]string{
				"file1.txt": "roadie://data/file1.txt",
				"file2.txt": "roadie://data/file2.txt",
			},
		},
		{
			"file1.txt -> roadie://data/file1.txt\nfile2.txt -> roadie://data/file2.txt\n",
			map[string]string{
				"file1.txt": "roadie://data/file1.txt",
				"file2.txt": "roadie://data/file2.txt",
			},
		},
	}

	for _, c := range cases {
		res := locationURLs(c.input)
		if len(res) != len(c.expect) {
			t.Errorf("the number of parsed urls is %v, want %v", len(res), len(c.expect))
		}
		for file, loc := range c.expect {
			if res[file] != loc {
				t.Errorf("URL of %v is %v, want %v", file, res[file], loc)
			}
		}
	}

}

// TestCmdDataPut tests optDataPut.run.
func TestCmdDataPut(t *testing.T) {

	var err error
	var output bytes.Buffer
	opt := optDataPut{
		Metadata: testMetadata(&output, nil),
	}
	s, err := opt.StorageManager()
	if err != nil {
		t.Fatalf("cannot get a storage manager: %v", err)
	}

	cases := []struct {
		filename   string
		storedName string
		expected   map[string]string
	}{
		// Put a file without renaming.
		{"data.go", "", map[string]string{"data.go": "roadie://data/data.go"}},
		// Put a file with renaming it.
		{"data.go", "source.go", map[string]string{"data.go": "roadie://data/source.go"}},
		// Put files matching a glob pattern.
		{"util/*", "", map[string]string{
			"util/archive.go":      "roadie://data/archive.go",
			"util/archive_test.go": "roadie://data/archive_test.go",
		}},
		// Put files matching a glob pattern into a directory.
		{"util/*", "util", map[string]string{
			"util/archive.go":      "roadie://data/util/archive.go",
			"util/archive_test.go": "roadie://data/util/archive_test.go",
		}},
		// Put a directory.
		// {"util", "", map[string]string{
		// 	"archive.go":      "roadie://data/archive.go",
		// 	"archive_test.go": "roadie://data/archive_test.go",
		// }},
	}

	for _, c := range cases {

		t.Run(fmt.Sprintf("File=%q,Name=%q", c.filename, c.storedName), func(t *testing.T) {
			defer output.Reset()

			opt.Filename = c.filename
			opt.StoredName = c.storedName
			if err = opt.run(); err != nil {
				t.Fatalf("run returns an error: %v", err)
			}

			var matches []string
			locations := locationURLs(output.String())
			matches, err = filepath.Glob(opt.Filename)
			if err != nil {
				t.Fatalf("Glob returns an error: %v", err)
			}
			for _, f := range matches {

				if locations[f] != c.expected[f] {
					t.Errorf("uploaded location of %v is %q, want %v", f, locations[f], c.expected[f])
				}

				var loc *url.URL
				loc, err = url.Parse(locations[f])
				if err != nil {
					t.Fatalf("cannot parse a URL: %v", err)
				}

				var data bytes.Buffer
				err = s.Download(opt.Context, loc, &data)
				if err != nil {
					t.Fatalf("Download returns an error: %v", err)
				}

				var original []byte
				original, err = ioutil.ReadFile(f)
				if err != nil {
					t.Fatalf("ReadFile(%v) returns an error: %v", f, err)
				}
				if data.String() != string(original) {
					t.Errorf("uploaded file is broken %v, want %v", data.String(), string(original))
				}

			}
		})

	}

	t.Run("not existing file", func(t *testing.T) {
		opt.Filename = "_data.go"
		opt.StoredName = ""
		if err = opt.run(); err == nil {
			t.Error("putting not existing files doesn't return any errors")
		}
	})

	t.Run("invalid glob pattern", func(t *testing.T) {
		opt.Filename = "[b-a"
		opt.StoredName = ""
		if err = opt.run(); err == nil {
			t.Error("Give a wrong pattern but no errors occur.")
		}
	})

	t.Run("empty pattern", func(t *testing.T) {
		opt.Filename = ""
		opt.StoredName = ""
		if err = opt.run(); err == nil {
			t.Error("Give empty pattern but no errors occur.")
		}
	})

}
