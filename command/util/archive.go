//
// command/util/archive.go
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
	"bufio"
	"compress/gzip"
	"os"
	"path/filepath"
	"strings"
)

// Archive makes a tar.gz file consists of file tree
func Archive(root string, filename string, excludes []string) (err error) {

	writeFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer writeFile.Close()

	writer := bufio.NewWriter(writeFile)
	defer writer.Flush()

	zipWriter, err := gzip.NewWriterLevel(writer, gzip.BestCompression)
	if err != nil {
		return
	}
	defer zipWriter.Close()

	tarWriter := tar.NewWriter(zipWriter)
	defer tarWriter.Close()

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Directory won't include the archive.
		if info.IsDir() {
			if strings.HasSuffix(path, ".git") {
				return filepath.SkipDir
			}
			return nil
		}

		// Check the found path matches exclude rules.
		matched, err := excludePath(path, excludes)
		if err != nil {
			return err
		} else if matched {
			return nil
		}

		// Write a file header. (for Windows: path should be slashed)
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		if root != path {
			var rel string
			rel, err = filepath.Rel(root, path)
			if err != nil {
				return err
			}
			header.Name = filepath.ToSlash(rel)
		}
		tarWriter.WriteHeader(header)

		// Prepare to write a file body.
		fp, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fp.Close()

		// Write the body.
		reader := bufio.NewReader(fp)
		_, err = reader.WriteTo(tarWriter)
		return err

	})

}

// Check whether a given path is an exclude path.
func excludePath(path string, excludes []string) (bool, error) {

	for _, part := range strings.Split(filepath.ToSlash(path), "/") {

		for _, pattern := range excludes {

			if match, err := filepath.Match(pattern, part); err != nil {
				return false, err
			} else if match {
				return true, nil
			}
		}

	}

	return false, nil

}
