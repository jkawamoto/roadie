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
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Archive makes a tar.gz file consists of file tree
func Archive(root string, filename string, excludes []string) error {

	writeFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer writeFile.Close()

	writer := bufio.NewWriter(writeFile)
	defer writer.Flush()

	zipWriter, err := gzip.NewWriterLevel(writer, gzip.BestCompression)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer zipWriter.Close()

	tarWriter := tar.NewWriter(zipWriter)
	defer tarWriter.Close()

	return filepath.Walk(root, tarballing(tarWriter, excludes))

}

// Create a call back to add founded files to a given archive.
func tarballing(writer *tar.Writer, excludes []string) filepath.WalkFunc {

	return func(path string, info os.FileInfo, err error) error {

		// Directory won't include the archive.
		if info.IsDir() {
			if strings.HasSuffix(path, ".git") {
				return filepath.SkipDir
			}
			return nil
		}

		// Check the found path matches exclude rules.
		if m, e := excludePath(path, excludes); e != nil {
			return e
		} else if m {
			return nil
		}

		// For Windows: Replace path delimiters.
		path = filepath.ToSlash(path)

		// Write a file header.
		header, err := tar.FileInfoHeader(info, path)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		if strings.HasPrefix(path, "../") {
			header.Name = path[3:]
		} else {
			header.Name = path
		}
		writer.WriteHeader(header)

		// Prepare to write a file body.
		fp, err := os.Open(path)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		defer fp.Close()

		// Write the body.
		reader := bufio.NewReader(fp)
		if _, err := reader.WriteTo(writer); err != nil {
			fmt.Println(err.Error())
			return err
		}

		return nil

	}

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
