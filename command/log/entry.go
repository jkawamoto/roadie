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
	"github.com/jkawamoto/roadie/config"
	"golang.org/x/net/context"
	"google.golang.org/api/logging/v2beta1"
)

// Entry defines a generic structure of one log entry.
type Entry struct {
	Timestamp time.Time
	Payload   interface{}
}

// GetEntriesFunc is a helper function to call GetEntries with EntryRequesterFunc.
func GetEntriesFunc(ctx context.Context, filter string, requester EntryRequesterFunc, handler func(*Entry) error) (err error) {
	return GetEntries(ctx, filter, requester, handler)
}

// GetEntries requests log entries of a project via a given requester under a given context.
// Obtained log entries are filtered by a given filter query and will be passed
// a given handler entry by entry. If the handler returns non nil value,
// obtaining log entries is canceled immediately.
func GetEntries(ctx context.Context, filter string, requester EntryRequester, handler func(*Entry) error) (err error) {

	cfg, ok := config.FromContext(ctx)
	if !ok {
		return fmt.Errorf("The given context doesn't have any config: %s", ctx)
	}

	// pageToken will be used when logs are divided into several pages.
	pageToken := ""
	for {

		res, err := requester.Do(&logging.ListLogEntriesRequest{
			ProjectIds: []string{cfg.Project},
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

			timestamp, err := time.Parse(TimeFormat, v.Timestamp)
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
				if err = handler(&Entry{
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
	ctx context.Context, project, instanceName string, start time.Time, requester EntryRequesterFunc, handler func(time.Time, *RoadiePayload) error) (err error) {

	return GetInstanceLogEntries(ctx, instanceName, start, requester, handler)
}

// GetInstanceLogEntries requests log entries of a given instance.
// Obtained log entries will be passed a given handler entry by entry.
// If the handler returns non nil value, obtaining log entries is canceled immediately.
func GetInstanceLogEntries(
	ctx context.Context, instanceName string, start time.Time, requester EntryRequester, handler func(time.Time, *RoadiePayload) error) (err error) {

	// Instead of logName, which is specified TAG env in roadie-gce,
	// use instance name to distinguish instances. This update makes all logs
	// will have same log name, docker, so that such log can be stored into
	// GCS easily.
	filter := fmt.Sprintf(
		"resource.type = \"gce_instance\" AND jsonPayload.instance_name = \"%s\" AND timestamp > \"%s\"",
		instanceName, start.In(time.UTC).Format(TimeFormat))

	return GetEntries(ctx, filter, requester, func(entry *Entry) error {
		payload, err := NewRoadiePayload(entry)
		if err != nil {
			return err
		}
		return handler(entry.Timestamp, payload)
	})
}

// GetOperationLogEntriesFunc is a helper function to call GetOperationLogEntries with LogEntryRequesterFunc.
func GetOperationLogEntriesFunc(ctx context.Context,
	requester EntryRequesterFunc, handler func(time.Time, *ActivityPayload) error) (err error) {

	return GetOperationLogEntries(ctx, requester, handler)
}

// GetOperationLogEntries requests log entries about google cloud platform operations.
// Obtained log entries will be passed a given handler entry by entry.
// If the handler returns non nil value, obtaining log entries is canceled immediately.
func GetOperationLogEntries(ctx context.Context,
	requester EntryRequester, handler func(time.Time, *ActivityPayload) error) (err error) {

	return GetEntries(
		ctx, "jsonPayload.event_type = \"GCE_OPERATION_DONE\"", requester,
		func(entry *Entry) error {
			payload, err := NewActivityPayload(entry)
			if err != nil {
				return err
			}
			return handler(entry.Timestamp, payload)
		})
}
