//
// command/source_test.go
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
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"testing"
)

// Since cmdSourcePut is a tiny wrapper of uploadSourceFiles.
// We only test uploadSourceFiles, here.

// TestUploadSourceFiles tests uploadSourceFiles uploads a directory or a file
// to the source section of a cloud storage.
func TestUploadSourceFiles(t *testing.T) {

	var err error
	m := testMetadata(nil)
	s, err := m.StorageManager()
	if err != nil {
		t.Fatalf("cannot get a storage manager: %v", err)
	}

	cases := []struct {
		path     string
		rename   string
		expected string
		files    map[string]string
	}{
		// Test uploading a directory without renaming.
		{"util", "", "roadie://source/util.tar.gz", map[string]string{
			"archive_test.go": "util/archive_test.go",
			"archive.go":      "util/archive.go",
		}},
		// Test uploading a directory with a file name.
		{"util", "dir", "roadie://source/dir.tar.gz", map[string]string{
			"archive_test.go": "util/archive_test.go",
			"archive.go":      "util/archive.go",
		}},
		// Test uploading a file without renaming.
		{"source.go", "", "roadie://source/source.go.tar.gz", map[string]string{
			"source.go": "source.go",
		}},
		// Test uploading a file with renaming.
		{"helper_test.go", "another_test", "roadie://source/another_test.tar.gz", map[string]string{
			"helper_test.go": "helper_test.go",
		}},
	}

	var loc *url.URL
	for _, c := range cases {

		loc, err = uploadSourceFiles(m, c.path, c.rename, nil)
		if err != nil {
			t.Errorf("error from uploadFiles of %q and %q: %v", c.path, c.rename, err)
		}
		if loc.String() != c.expected {
			t.Errorf("uploadFiles of %q and %q returns %v, want %v", c.path, c.rename, loc, c.expected)
		}

		reader, writer := io.Pipe()
		ch := make(chan error)
		go func(out io.WriteCloser) {
			err := s.Download(m.Context, loc, out)
			out.Close()
			ch <- err
		}(writer)

		var gzReader *gzip.Reader
		gzReader, err = gzip.NewReader(reader)
		if err != nil {
			t.Fatalf("cannot create a gzip.Reader: %v", err)
		}
		defer gzReader.Close()

		var header *tar.Header
		tarReader := tar.NewReader(gzReader)
		for {

			header, err = tarReader.Next()
			if err == io.EOF {
				break
			} else if err != nil {
				t.Fatalf("reading a tarball returns an error: %v", err)
			}

			if path, exist := c.files[header.Name]; !exist {

				t.Errorf("archive file contains wrong file %v", header.Name)

			} else {

				var original []byte
				original, err = ioutil.ReadFile(path)
				if err != nil {
					t.Fatalf("ReadFile(%v) returns an error: %v", path, err)
				}

				var data []byte
				data, err = ioutil.ReadAll(tarReader)
				if err != nil {
					t.Fatalf("ReadAll returns an error: %v", err)
				}

				if string(original) != string(data) {
					t.Errorf("downloaded file %v, want %v", string(data), string(original))
				}

			}

		}

		err = <-ch
		if err != nil {
			t.Fatal("Download returns an error: %v", err)
		}
		close(ch)

	}

	// Test uploading nonexisting files.
	_, err = uploadSourceFiles(m, "_source.go", "", nil)
	if _, ok := err.(*os.PathError); !ok {
		t.Errorf("no os.PathError returned for an invalid file: %v", err)
	}

}
