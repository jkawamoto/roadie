//
// cloud/gce/log_payload.go
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

package gce

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/ptypes/struct"
	"github.com/jkawamoto/structpbconv"
)

// ActivityPayload defines the payload structure of activity log.
type ActivityPayload struct {
	EventTimestampUs string `structpb:"event_timestamp_us"`
	EventType        string `structpb:"vent_type"`
	TraceID          string `structpb:"trace_id"`
	Actor            struct {
		User string
	}
	Resource struct {
		Zone string
		Type string
		ID   string
		Name string
	}
	Version      string
	EventSubtype string `structpb:"event_subtype"`
	Operation    struct {
		Zone string
		Type string
		ID   string
		Name string
	}
}

// NewActivityPayload converts LogEntry's payload to a ActivityPayload.
func NewActivityPayload(payload interface{}) (res *ActivityPayload, err error) {

	switch s := payload.(type) {
	case *ActivityPayload:
		res = s
	case *structpb.Struct:
		res = &ActivityPayload{}
		structpbconv.Convert(s, res)
	default:
		return nil, fmt.Errorf("Given payload is not an instance of *structpb.Struct: %v", payload)
	}

	return
}

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
		structpbconv.Convert(s, res)
	default:
		return nil, fmt.Errorf("Given payload is not an instance of *structpb.Struct: %v", payload)
	}

	res.Log = strings.TrimRight(res.Log, "\n")
	return
}
