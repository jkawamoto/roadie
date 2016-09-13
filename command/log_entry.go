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
	"github.com/mitchellh/mapstructure"
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

// LogEntryRequester is an inteface used in GetLogEntries.
// This interface requests supplying Do method which process a request of
// obtaining log entries.
type LogEntryRequester interface {
	Do(*logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error)
}

// CloudLoggingService implements LogEntryRequester interface.
// It requests logs to google cloud logging service.
type CloudLoggingService struct {
	service *logging.Service
}

// NewCloudLoggingService creates a new CloudLoggingService with a given context.
func NewCloudLoggingService(ctx context.Context) (res *CloudLoggingService, err error) {

	client, err := google.DefaultClient(ctx, logging.CloudPlatformReadOnlyScope)
	if err != nil {
		return
	}

	service, err := logging.New(client)
	if err != nil {
		return
	}

	return &CloudLoggingService{service: service}, nil

}

// Do requests a given request with the specified context.
func (s *CloudLoggingService) Do(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {

	return s.service.Entries.List(req).Do()

}

// LogEntryRequesterFunc will be used to implement LogEntryRequester interface
// on functions.
type LogEntryRequesterFunc func(*logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error)

// Do implements LogEntryRequester interface.
func (f LogEntryRequesterFunc) Do(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {
	return f(req)
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

// RoadiePayload defines the payload structure of insance logs.
type RoadiePayload struct {
	Username     string
	Stream       string
	Log          string
	ContainerID  string `mapstructure:"container_id"`
	InstanceName string `mapstructure:"instance_name"`
}

// NewRoadiePayload converts LogEntry's payload to a RoadiePayload.
func NewRoadiePayload(entry *LogEntry) (*RoadiePayload, error) {

	var res RoadiePayload
	if err := mapstructure.Decode(entry.Payload, &res); err != nil {
		return nil, err
	}
	res.Log = strings.TrimRight(res.Log, "\n")

	return &res, nil
}

// ActivityPayload defines the payload structure of activity log.
type ActivityPayload struct {
	EventTimestampUs string `mapstructure:"event_timestamp_us"`
	EventType        string `mapstructure:"vent_type"`
	TraceID          string `mapstructure:"trace_id"`
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
	EventSubtype string `mapstructure:"event_subtype"`
	Operation    struct {
		Zone string
		Type string
		ID   string
		Name string
	}
}

// NewActivityPayload converts LogEntry's payload to a ActivityPayload.
func NewActivityPayload(entry *LogEntry) (*ActivityPayload, error) {
	var res ActivityPayload
	if err := mapstructure.Decode(entry.Payload, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

const (
	// EventSubtypeInsert means this event is creating an instance.
	EventSubtypeInsert = "compute.instances.insert"
	// EventSubtypeDelete means this event is deleting an instance.
	EventSubtypeDelete = "compute.instances.delete"
)
