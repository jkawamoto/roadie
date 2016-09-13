//
// command/log/entry.go
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
	"time"

	"github.com/jkawamoto/roadie/chalk"
	"golang.org/x/net/context"
	"google.golang.org/api/logging/v2beta1"
)

// LogEntry defines a generic structure of one log entry.
type LogEntry struct {
	Timestamp time.Time
	Payload   interface{}
}

// GetLogEntriesFunc is a helper function to call GetLogEntries with LogEntryRequesterFunc.
func GetLogEntriesFunc(ctx context.Context, project, filter string, requester LogEntryRequesterFunc, handler func(*LogEntry) error) (err error) {
	return GetLogEntries(ctx, project, filter, requester, handler)
}

// GetLogEntries requests log entries of a given project via a given requester.
// Obtained log entries are filtered by a given filter query and will be passed
// a given handler entry by entry. If the handler returns non nil value,
// obtaining log entries is canceled immediately.
func GetLogEntries(ctx context.Context, project, filter string, requester LogEntryRequester, handler func(*LogEntry) error) (err error) {

	// pageToken will be used when logs are divided into several pages.
	pageToken := ""
	for {

		res, err := requester.Do(&logging.ListLogEntriesRequest{
			ProjectIds: []string{project},
			Filter:     filter,
			PageToken:  pageToken,
		})
		if err != nil {
			return err
		}

		for _, v := range res.Entries {
			// TODO: Entries which don't have JsonPayload may contain system messages.
			if v.JsonPayload == nil {
				continue
			}

			// Time format of log entries aren't generalized. Thus reformat it here.
			if strings.Contains(v.Timestamp, ".") {
				v.Timestamp = strings.Split(v.Timestamp, ".")[0] + "Z"
			}

			timestamp, err := time.Parse(LogTimeFormat, v.Timestamp)
			// TODO: Should be replaced to outputting to stderr via some logging method.
			if err != nil {
				fmt.Println(chalk.Red.Color(err.Error()))
				continue
			}
			timestamp = timestamp.In(time.Local)

			select {
			case <-ctx.Done():
				// If canceled, return with a given error.
				return ctx.Err()

			default:
				// Not canceled yet, pass an obtained entry to the handler.
				if err = handler(&LogEntry{
					Timestamp: timestamp,
					Payload:   v.JsonPayload,
				}); err != nil {
					return err
				}
			}

		}

		pageToken = res.NextPageToken
		if pageToken == "" {
			break
		}

	}

	return

}

// GetInstanceLogEntriesFunc is a helper function to call GetInstanceLogEntries with LogEntryRequesterFunc.
func GetInstanceLogEntriesFunc(
	ctx context.Context, project, instanceName string, start time.Time, requester LogEntryRequesterFunc, handler func(time.Time, *RoadiePayload) error) (err error) {

	return GetInstanceLogEntries(ctx, project, instanceName, start, requester, handler)
}

// GetInstanceLogEntries requests log entries of a given instance.
// Obtained log entries will be passed a given handler entry by entry.
// If the handler returns non nil value, obtaining log entries is canceled immediately.
func GetInstanceLogEntries(
	ctx context.Context, project, instanceName string, start time.Time, requester LogEntryRequester, handler func(time.Time, *RoadiePayload) error) (err error) {

	// Instead of logName, which is specified TAG env in roadie-gce,
	// use instance name to distinguish instances. This update makes all logs
	// will have same log name, docker, so that such log can be stored into
	// GCS easily.
	filter := fmt.Sprintf(
		"resource.type = \"gce_instance\" AND jsonPayload.instance_name = \"%s\" AND timestamp > \"%s\"",
		instanceName, start.In(time.UTC).Format(LogTimeFormat))

	return GetLogEntries(ctx, project, filter, requester, func(entry *LogEntry) error {
		payload, err := NewRoadiePayload(entry)
		if err != nil {
			return err
		}
		return handler(entry.Timestamp, payload)
	})
}

// GetOperationLogEntriesFunc is a helper function to call GetOperationLogEntries with LogEntryRequesterFunc.
func GetOperationLogEntriesFunc(ctx context.Context,
	project string, requester LogEntryRequesterFunc, handler func(time.Time, *ActivityPayload) error) (err error) {

	return GetOperationLogEntries(ctx, project, requester, handler)
}

// GetOperationLogEntries requests log entries about google cloud platform operations.
// Obtained log entries will be passed a given handler entry by entry.
// If the handler returns non nil value, obtaining log entries is canceled immediately.
func GetOperationLogEntries(ctx context.Context,
	project string, requester LogEntryRequester, handler func(time.Time, *ActivityPayload) error) (err error) {

	return GetLogEntries(
		ctx, project, "jsonPayload.event_type = \"GCE_OPERATION_DONE\"", requester,
		func(entry *LogEntry) error {
			payload, err := NewActivityPayload(entry)
			if err != nil {
				return err
			}
			return handler(entry.Timestamp, payload)
		})
}
