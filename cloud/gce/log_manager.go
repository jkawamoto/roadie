//
// cloud/gce/log_manager.go
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
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/jkawamoto/roadie/cloud"

	"cloud.google.com/go/logging/apiv2"
	"google.golang.org/api/iterator"
	loggingpb "google.golang.org/genproto/googleapis/logging/v2"
)

// LogManager implements cloud.LogManager interface.
// It requests logs to google cloud logging service.
type LogManager struct {
	// Config is a reference for a configuration of GCP.
	Config    *GcpConfig
	Logger    *log.Logger
	SleepTime time.Duration
}

// EntryHandler is a function type to handler Entries.
type EntryHandler func(*loggingpb.LogEntry) error

// RoadiePayloadHandler is a function type to handle RoadiePayloads.
type RoadiePayloadHandler func(time.Time, string) error

// ActivityPayloadHandler is a function type to handle ActivityPayloads.
type ActivityPayloadHandler func(time.Time, *ActivityPayload) error

// NewLogManager creates a new log manager.
func NewLogManager(cfg *GcpConfig, logger *log.Logger) (m *LogManager) {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	return &LogManager{
		Config:    cfg,
		Logger:    logger,
		SleepTime: 30 * time.Second,
	}

}

// Get requests log entries of the given named instance.
func (s *LogManager) Get(ctx context.Context, instanceName string, from time.Time, handler cloud.LogHandler) (err error) {

	// Determine when the newest instance starts.
	err = s.OperationLogEntries(ctx, from, func(timestamp time.Time, payload *ActivityPayload) (err error) {
		if payload.Resource.Name == instanceName {
			if payload.EventSubtype == LogEventSubtypeInsert {
				from = timestamp
			}
		}
		return
	})
	if err != nil {
		return
	}

	// Request log entries.
	return s.InstanceLogEntries(ctx, instanceName, from, func(timestamp time.Time, payload string) (err error) {
		return handler(timestamp, payload, false)
	})



}

// Entries get log entries matching with a given filter from given project logs.
// Found log entries will be passed a given handler one by one.
// If the handler returns non-nil value as an error, this function will end.
func (s *LogManager) Entries(ctx context.Context, filter string, handler EntryHandler) (err error) {

	s.Logger.Println("Retrieving log:", filter)
	client, err := logging.NewClient(ctx)
	if err != nil {
		return
	}
	defer client.Close()

	iter := client.ListLogEntries(ctx, &loggingpb.ListLogEntriesRequest{
		ResourceNames: []string{
			fmt.Sprintf("projects/%v", s.Config.Project),
		},
		Filter: filter,
	})
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		e, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return err
		}
		if err := handler(e); err != nil {
			return err
		}
	}

	s.Logger.Println("Finished retrieving log")
	return
}

// InstanceLogEntries requests log entries of a given instance.
// Obtained log entries will be passed a given handler entry by entry.
// If the handler returns non nil value, obtaining log entries is canceled immediately.
func (s *LogManager) InstanceLogEntries(ctx context.Context, instanceName string, from time.Time, handler RoadiePayloadHandler) error {

	// Instead of logName, which is specified TAG env in roadie-gce,
	// use instance name to distinguish instances. This update makes all logs
	// will have same log name, docker, so that such log can be stored into
	// GCS easily.
	filter := fmt.Sprintf(
		`resource.type = "gce_instance" AND jsonPayload.instance_name = "%s" AND timestamp > "%s"`,
		instanceName, from.In(time.UTC).Format(LogTimeFormat))

	return s.Entries(ctx, filter, func(entry *loggingpb.LogEntry) error {
		payload := entry.GetJsonPayload()
		if value, ok := payload.GetFields()["MESSAGE"]; !ok {
			return nil
		} else if msg := value.GetStringValue(); msg == "" {
			return nil
		} else {
			ts := entry.GetTimestamp()
			return handler(time.Unix(ts.Seconds, int64(ts.Nanos)).In(time.Local), msg)
		}
	})

}

// OperationLogEntries requests log entries about google cloud platform operations.
// Obtained log entries will be passed a given handler entry by entry.
// If the handler returns non nil value, obtaining log entries is canceled immediately.
func (s *LogManager) OperationLogEntries(ctx context.Context, from time.Time, handler ActivityPayloadHandler) error {

	filter := fmt.Sprintf(
		`jsonPayload.event_type = "GCE_OPERATION_DONE" AND timestamp > "%s"`,
		from.In(time.UTC).Format(LogTimeFormat))
	return s.Entries(ctx, filter, func(entry *loggingpb.LogEntry) error {
		payload, err := NewActivityPayload(entry.GetJsonPayload())
		if err != nil {
			return err
		}
		ts := entry.GetTimestamp()
		return handler(time.Unix(ts.Seconds, int64(ts.Nanos)).In(time.Local), payload)
	})

}
