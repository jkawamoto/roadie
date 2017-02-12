//
// command/log/entry.go
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

package log

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/logging"

	"github.com/jkawamoto/roadie/config"
)

// EntryHandler is a function type to handler Entries.
type EntryHandler func(*logging.Entry) error

// GetEntriesFunc is a helper function to call GetEntries with EntryRequesterFunc.
func GetEntriesFunc(ctx context.Context, filter string, requester EntryRequesterFunc, handler EntryHandler) (err error) {
	return GetEntries(ctx, filter, requester, handler)
}

// GetEntries requests log entries of a project via a given requester under a given context.
// Obtained log entries are filtered by a given filter query and will be passed
// a given handler entry by entry. If the handler returns non nil value,
// obtaining log entries is canceled immediately.
func GetEntries(ctx context.Context, filter string, requester EntryRequester, handler EntryHandler) (err error) {

	cfg, err := config.FromContext(ctx)
	if err != nil {
		return
	}
	return requester.Entries(cfg.Project, filter, handler)

}

// RoadiePayloadHandler is a function type to handle RoadiePayloads.
type RoadiePayloadHandler func(time.Time, *RoadiePayload) error

// GetInstanceLogEntriesFunc is a helper function to call GetInstanceLogEntries with LogEntryRequesterFunc.
func GetInstanceLogEntriesFunc(
	ctx context.Context, instanceName string, start time.Time, requester EntryRequesterFunc, handler RoadiePayloadHandler) (err error) {

	return GetInstanceLogEntries(ctx, instanceName, start, requester, handler)
}

// GetInstanceLogEntries requests log entries of a given instance.
// Obtained log entries will be passed a given handler entry by entry.
// If the handler returns non nil value, obtaining log entries is canceled immediately.
func GetInstanceLogEntries(
	ctx context.Context, instanceName string, start time.Time, requester EntryRequester, handler RoadiePayloadHandler) (err error) {

	// Instead of logName, which is specified TAG env in roadie-gce,
	// use instance name to distinguish instances. This update makes all logs
	// will have same log name, docker, so that such log can be stored into
	// GCS easily.
	filter := fmt.Sprintf(
		"resource.type = \"gce_instance\" AND jsonPayload.instance_name = \"%s\" AND timestamp > \"%s\"",
		instanceName, start.In(time.UTC).Format(TimeFormat))

	return GetEntries(ctx, filter, requester, func(entry *logging.Entry) error {
		payload, err := NewRoadiePayload(entry.Payload)
		if err != nil {
			return err
		}
		return handler(entry.Timestamp.In(time.Local), payload)
	})
}

// ActivityPayloadHandler is a function type to handle ActivityPayloads.
type ActivityPayloadHandler func(time.Time, *ActivityPayload) error

// GetOperationLogEntriesFunc is a helper function to call GetOperationLogEntries with LogEntryRequesterFunc.
func GetOperationLogEntriesFunc(ctx context.Context, requester EntryRequesterFunc, handler ActivityPayloadHandler) (err error) {

	return GetOperationLogEntries(ctx, requester, handler)
}

// GetOperationLogEntries requests log entries about google cloud platform operations.
// Obtained log entries will be passed a given handler entry by entry.
// If the handler returns non nil value, obtaining log entries is canceled immediately.
func GetOperationLogEntries(ctx context.Context, requester EntryRequester, handler ActivityPayloadHandler) (err error) {

	return GetEntries(
		ctx, "jsonPayload.event_type = \"GCE_OPERATION_DONE\"", requester,
		func(entry *logging.Entry) error {
			payload, err := NewActivityPayload(entry.Payload)
			if err != nil {
				return err
			}
			return handler(entry.Timestamp.In(time.Local), payload)
		})
}
