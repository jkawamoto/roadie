//
// cloud/gcp/helper.go
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

package gcp

import (
	"net/url"
	"path"
	"strings"

	"github.com/jkawamoto/roadie/script"
)

var (
	// truePtr is a pointer of a true value.
	truePtr = toBoolPtr(true)
	// falsePtr is a pointer of a false value.
	falsePtr = toBoolPtr(false)
)

// ReplaceURLScheme replaced URLs which start with "roadie://".
// Those URLs are modified to "gs://<bucketname>/.roadie/".
func ReplaceURLScheme(cfg *Config, task *script.Script) {

	// Replace source section.
	if strings.HasPrefix(task.Source, script.RoadieSchemePrefix) {
		task.Source = CreateURL(cfg, task.Source[RoadieSchemeURLOffset:])
	}

	// Replace data section.
	for i, url := range task.Data {
		if strings.HasPrefix(url, script.RoadieSchemePrefix) {
			task.Data[i] = CreateURL(cfg, url[RoadieSchemeURLOffset:])
		}
	}

	// Replace result section.
	if strings.HasPrefix(task.Result, script.RoadieSchemePrefix) {
		task.Result = CreateURL(cfg, task.Result[RoadieSchemeURLOffset:])
	}

}

// CreateURL creates a valid URL for uploaing object.
func CreateURL(cfg *Config, name string) string {

	u := url.URL{
		Scheme: "gs",
		Host:   cfg.Bucket,
		Path:   path.Join("/", StoragePrefix, name),
	}
	return u.String()

}

// toBoolPtr returns a pointer of the given boolean.
func toBoolPtr(b bool) *bool {
	return &b
}
