//
// command/util/fileinfo.go
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
	"strings"
	"time"

	storage "google.golang.org/api/storage/v1"
)

// FileInfo defines file information structure.
type FileInfo struct {
	Name        string
	Path        string
	TimeCreated time.Time
	Size        uint64
}

// NewFileInfo creates a file info from an object.
func NewFileInfo(f *storage.Object) *FileInfo {

	splitedName := strings.Split(f.Name, "/")
	t, _ := time.Parse("2006-01-02T15:04:05", strings.Split(f.TimeCreated, ".")[0])

	return &FileInfo{
		Name:        splitedName[len(splitedName)-1],
		Path:        f.Name,
		TimeCreated: t.In(time.Local),
		Size:        f.Size,
	}
}
