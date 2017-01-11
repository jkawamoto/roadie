//
// cloud/fileinfo.go
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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package cloud

import "time"

// FileInfo defines file information structure.
type FileInfo struct {
	// Name of the file, which means the base name.
	Name string
	// Path of the file.
	Path string
	// TimeCreated is the time when the file was created.
	TimeCreated time.Time
	// Size of the file.
	Size int64
}
