//
// command/log/roadie.go
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

package log

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/ptypes/struct"
)

// RoadiePayload defines the payload structure of instance logs.
type RoadiePayload struct {
	Username     string
	Stream       string
	Log          string
	ContainerID  string `structpb:"container_id"`
	InstanceName string `structpb:"instance_name"`
}

// NewRoadiePayload converts LogEntry's payload to a RoadiePayload.
func NewRoadiePayload(payload interface{}) (res *RoadiePayload, err error) {

	switch s := payload.(type) {
	case *RoadiePayload:
		res = s
	case *structpb.Struct:
		res = &RoadiePayload{}
		ConvertStructPB(s, res)
	default:
		return nil, fmt.Errorf("Given payload is not an instance of *structpb.Struct: %v", payload)
	}

	res.Log = strings.TrimRight(res.Log, "\n")
	return
}