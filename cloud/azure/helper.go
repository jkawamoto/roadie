//
// cloud/azure/helper.go
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

package azure

import (
	"encoding/json"
	"fmt"
)

const (
	// ComputeServiceResourceGroupName defines the default resource name.
	ComputeServiceResourceGroupName = "roadie"
)

// toPtr returns a pointer of a copy of a given string.
func toPtr(s string) *string {
	return &s
}

// toJSON returns a JSON string representing a given data.
func toJSON(param interface{}) string {
	buf, err := json.MarshalIndent(param, "", "  ")
	if err != nil {
		return fmt.Sprintln(param)
	}
	return string(buf)
}
