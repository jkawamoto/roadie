//
// util/path.go
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
	"net/url"
	"path/filepath"
)

// Basename computes the basename of a given filename.
func Basename(filename string) string {

	ext := filepath.Ext(filename)
	bodySize := len(filename) - len(ext)

	return filepath.Base(filename[:bodySize])

}

// CreateURL creates a valid URL for uploaing object.
func CreateURL(bucket, group, name string) *url.URL {

	return &url.URL{
		Scheme: "gs",
		Host:   bucket,
		Path:   filepath.ToSlash(filepath.Join("/", group, name)),
	}

}
