//
// command/util/archive_test.go
//
// Copyright (c) 2016 Junpei Kawamoto
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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package util

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

const archiveFile = "test.tar.gz"

func TestArchive(t *testing.T) {

	temp := os.TempDir()
	target := path.Join(temp, archiveFile)
	t.Logf("Creating an archive file: %s", target)

	if err := Archive("..", target, []string{"*.tar"}); err != nil {
		t.Error(err.Error())
	}

	root, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
	}

	os.Chdir(temp)
	defer func() {
		os.Chdir(root)
	}()
	exec.Command("tar", "-zxvf", archiveFile)

	if err := filepath.Walk(path.Join(root, ".."), checkExistence(temp)); err != nil {
		t.Error(err.Error())
	}

}

func checkExistence(target string) filepath.WalkFunc {

	return func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			if strings.HasSuffix(path, ".git") {
				return filepath.SkipDir
			}
			return nil
		} else if strings.HasSuffix(path, ".tar") {
			return nil
		}

		_, check := os.Stat(path)
		return check

	}

}
