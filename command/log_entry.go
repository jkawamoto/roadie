//
// command/log_entry.go
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

package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/jkawamoto/roadie/chalk"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/logging/v2beta1"
)

// LogTimeFormat defines time format of Google Logging.
const LogTimeFormat = "2006-01-02T15:04:05Z"

// LogEntry defines a generic structure of one log entry.
type LogEntry struct {
	Timestamp time.Time
	Payload   interface{}
}

// GetLogEntries downloads log entries as a goroutine.
func GetLogEntries(ctx context.Context, project, filter string, handler func(*LogEntry) error) (err error) {

	// context may be canceled in this function.
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client, err := google.DefaultClient(ctx, logging.CloudPlatformReadOnlyScope)
	if err != nil {
		return
	}

	service, err := logging.New(client)
	if err != nil {
		return
	}

	return getLogEntries(ctx, project, filter, handler, func(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {

		return service.Entries.List(req).Do()

	})

}

// getLogEntries requests log entries of a given project via requestDo function.
// Obtained log entries are filtered by a given filter query and will be passed
// a given handler entry by entry. If the handler returns non nil value,
// obtaining log entries is canceled immediately.
func getLogEntries(ctx context.Context, project, filter string,
	handler func(*LogEntry) error, requestDo func(*logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error)) (err error) {

	// pageToken will be used when logs are divided into several pages.
	pageToken := ""
	for {

		res, err := requestDo(&logging.ListLogEntriesRequest{
			ProjectIds: []string{project},
			Filter:     filter,
			PageToken:  pageToken,
		})
		if err != nil {
			return err
		}

		for _, v := range res.Entries {
			// TODO: Entries which don't have JsonPayload may containe system messages.
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
