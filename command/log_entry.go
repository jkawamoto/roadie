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

	"github.com/ttacon/chalk"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/logging/v2beta1"
)

// LogEntry defines a generic structure of one log entry.
type LogEntry struct {
	Timestamp time.Time
	Payload   interface{}
}

// GetLogEntries downloads log entries as a goroutine.
func GetLogEntries(project, filter string, ch chan<- *LogEntry, chErr chan<- error) {

	client, err := google.DefaultClient(context.Background(), logging.CloudPlatformReadOnlyScope)
	if err != nil {
		chErr <- err
		return
	}

	service, err := logging.New(client)
	if err != nil {
		chErr <- err
		return
	}

	pageToken := ""
	for {

		res, err := service.Entries.List(&logging.ListLogEntriesRequest{
			ProjectIds: []string{project}, Filter: filter, PageToken: pageToken}).Do()
		if err != nil {
			chErr <- err
			return
		}

		for _, v := range res.Entries {
			if v.JsonPayload != nil {

				if strings.Contains(v.Timestamp, ".") {
					v.Timestamp = strings.Split(v.Timestamp, ".")[0] + "Z"
				}

				timestamp, err := time.Parse("2006-01-02T15:04:05Z", v.Timestamp)
				if err != nil {
					fmt.Println(chalk.Red.Color(err.Error()))
					continue
				}
				timestamp = timestamp.In(time.Local)

				ch <- &LogEntry{
					Timestamp: timestamp,
					Payload:   v.JsonPayload,
				}

			}
		}

		pageToken = res.NextPageToken
		if pageToken == "" {
			break
		}

	}

	ch <- nil

}
