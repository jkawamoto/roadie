//
// command/util/archive_test.go
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

package util

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestArchive(t *testing.T) {

	var err error
	root := ".."

	cases := []struct {
		excludes []string
	}{
		{nil},
		{[]string{"*_test.go"}},
	}

	for _, c := range cases {

		expectedFiles := make(map[string]struct{})
		err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			var matched bool
			for _, e := range c.excludes {
				matched, err = filepath.Match(e, info.Name())
				if err != nil {
					return err
				} else if matched {
					return nil
				}
			}

			path, err = filepath.Rel(root, path)
			if err != nil {
				return err
			}
			// Archived files' names shoule be slashed.
			expectedFiles[filepath.ToSlash(path)] = struct{}{}
			return nil

		})
		if err != nil {
			t.Fatalf("Walking a directory returns an error: %v", err)
		}

		temp, err := ioutil.TempDir("", "TestArchive")
		if err != nil {
			t.Fatalf("TempDir returns an error: %v", err)
		}
		defer os.RemoveAll(temp)

		target := path.Join(temp, "test.tar.gz")
		t.Logf("Creating an archive file: %s", target)
		err = Archive(root, target, c.excludes)
		if err != nil {
			t.Fatalf("Archive returns an error: %v", err)
		}

		f, err := os.Open(target)
		if err != nil {
			t.Fatalf("Open(%v) returns an error: %v", target, err)
		}
		defer f.Close()

		zipReader, err := gzip.NewReader(f)
		if err != nil {
			t.Fatalf("gzip.NewReader returns an error: %v", err)
		}
		defer zipReader.Close()

		tarReader := tar.NewReader(zipReader)
		var header *tar.Header
		var expected, received []byte
		for {

			header, err = tarReader.Next()
			if err == io.EOF {
				break
			} else if err != nil {
				t.Fatalf("reading tarball header returns an error: %v", err)
			}

			expected, err = ioutil.ReadFile(filepath.Join(root, header.Name))
			if err != nil {
				t.Fatalf("cannot read %v: %v", header.Name, err)
			}
			received, err = ioutil.ReadAll(tarReader)
			if err != nil {
				t.Fatalf("reading a file from a tarball returns an error: %v", err)
			}
			if string(received) != string(expected) {
				t.Errorf("received %v, want %v", string(received), string(expected))
			}
			delete(expectedFiles, header.Name)

		}

		if len(expectedFiles) != 0 {
			t.Errorf("archived file misses %v", expectedFiles)
		}

	}

}

// TestArchiveFile tests archiving only one file by Archive.
func TestArchiveFile(t *testing.T) {

	target := "archive.go"

	temp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("TempDir returns an error: %v", err)
	}
	defer os.RemoveAll(temp)

	filename := filepath.Join(temp, target+".tar.gz")
	err = Archive(target, filename, nil)
	if err != nil {
		t.Fatalf("Archive returns an error: %v", err)
	}

	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Open(%v) returns an error: %v", filename, err)
	}
	defer f.Close()

	zipReader, err := gzip.NewReader(f)
	if err != nil {
		t.Fatalf("gzip.NewReader returns an error: %v", err)
	}
	defer zipReader.Close()

	tarReader := tar.NewReader(zipReader)
	var header *tar.Header
	var expected, received []byte
	for {

		header, err = tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatalf("reading tarball header returns an error: %v", err)
		}

		expected, err = ioutil.ReadFile(header.Name)
		if err != nil {
			t.Fatalf("cannot read %v: %v", header.Name, err)
		}
		received, err = ioutil.ReadAll(tarReader)
		if err != nil {
			t.Fatalf("reading a file from a tarball returns an error: %v", err)
		}
		if string(received) != string(expected) {
			t.Errorf("received %v, want %v", string(received), string(expected))
		}

	}

}

func TestArchiveNotExistingDirectory(t *testing.T) {

	temp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("TempDir returns an error: %v", err)
	}
	defer os.RemoveAll(temp)

	err = Archive("../_util", filepath.Join(temp, "test.tar.gz"), nil)
	if err == nil {
		t.Error("Archiving a not existing directory returns no errors")
	}

}
