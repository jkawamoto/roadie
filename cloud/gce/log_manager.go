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
	"time"

	"github.com/jkawamoto/roadie/cloud"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/logging/logadmin"
	"google.golang.org/api/iterator"
)

// LogManager implements cloud.LogManager interface.
// It requests logs to google cloud logging service.
type LogManager struct {
	// Config is a reference for a configuration of GCP.
	Config    *GcpConfig
	SleepTime time.Duration
}

// EntryHandler is a function type to handler Entries.
type EntryHandler func(*logging.Entry) error

// RoadiePayloadHandler is a function type to handle RoadiePayloads.
type RoadiePayloadHandler func(time.Time, *RoadiePayload) error

// ActivityPayloadHandler is a function type to handle ActivityPayloads.
type ActivityPayloadHandler func(time.Time, *ActivityPayload) error

// NewLogManager creates a new log manager.
func NewLogManager(cfg *GcpConfig) (m *LogManager) {

	return &LogManager{
		Config:    cfg,
		SleepTime: 30 * time.Second,
	}

}

// Get requests log entries of the given named instance.
func (s *LogManager) Get(ctx context.Context, instanceName string, from time.Time, handler cloud.LogHandler) (err error) {

	var stopIteration = fmt.Errorf("Stop Iteration")

	// Determine when the newest instance starts.
	err = s.OperationLogEntries(ctx, func(timestamp time.Time, payload *ActivityPayload) (err error) {
		if payload.Resource.Name == instanceName {
			if payload.EventSubtype == LogEventSubtypeInsert {
				from = timestamp
				return stopIteration
			}
		}
		return
	})
	if err != nil && err != stopIteration {
		return
	}

	// Request log entries.
	return s.InstanceLogEntries(ctx, instanceName, from, func(timestamp time.Time, payload *RoadiePayload) (err error) {

		return handler(timestamp, payload.Log, payload.Stream != "stdout")

	})

}

// Entries get log entries matching with a given filter from given project logs.
// Found log entries will be passed a given handler one by one.
// If the handler returns non-nil value as an error, this function will end.
func (s *LogManager) Entries(ctx context.Context, filter string, handler EntryHandler) (err error) {

	client, err := logadmin.NewClient(ctx, s.Config.Project)
	if err != nil {
		return
	}
	defer client.Close()

	iter := client.Entries(ctx, logadmin.Filter(filter))
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		e, err := iter.Next()
		if err == iterator.Done {
			return nil
		} else if err != nil {
			return err
		}
		if err := handler(e); err != nil {
			return err
		}
	}
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

	return s.Entries(ctx, filter, func(entry *logging.Entry) error {
		payload, err := NewRoadiePayload(entry.Payload)
		if err != nil {
			return err
		}
		return handler(entry.Timestamp.In(time.Local), payload)
	})

}

// OperationLogEntries requests log entries about google cloud platform operations.
// Obtained log entries will be passed a given handler entry by entry.
// If the handler returns non nil value, obtaining log entries is canceled immediately.
func (s *LogManager) OperationLogEntries(ctx context.Context, handler ActivityPayloadHandler) error {

	filter := `jsonPayload.event_type = "GCE_OPERATION_DONE"`
	return s.Entries(ctx, filter, func(entry *logging.Entry) error {
		payload, err := NewActivityPayload(entry.Payload)
		if err != nil {
			return err
		}
		return handler(entry.Timestamp.In(time.Local), payload)
	})

}
