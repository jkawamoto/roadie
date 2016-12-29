//
// command/log_test.go
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
	"bufio"
	"bytes"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/logging"
	"golang.org/x/net/context"

	"github.com/jkawamoto/roadie/command/log"
	"github.com/jkawamoto/roadie/config"
)

// resource a structure used in ActivityPayload.Resource.
type PayloadResource struct {
	Zone string
	Type string
	ID   string
	Name string
}

// TestCmdLog tests cmdLog outputs correct log entries.
func TestCmdLog(t *testing.T) {

	instance := "test-instance"
	samplePayload := log.RoadiePayload{
		Username:     "not used",
		Stream:       "stdout",
		Log:          "sample log output",
		ContainerID:  "not used",
		InstanceName: instance,
	}

	var requester log.EntryRequesterFunc
	// Make a mock requester which doesn't requests but returns pre-defined log entries.
	requester = func(project, filter string, handler log.EntryHandler) error {

		if strings.Contains(filter, "GCE_OPERATION_DONE") {
			// If the request is for operation logs,
			// returns a dummy log of starting an instance.

			return handler(&logging.Entry{
				Timestamp: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
				Payload: &log.ActivityPayload{
					EventType: "GCE_OPERATION_DONE",
					Resource: PayloadResource{
						Name: instance,
					},
					EventSubtype: log.EventSubtypeInsert,
				},
			})

		}

		// Otherwise, requests must be for instance logs.
		if !strings.Contains(filter, instance) {
			t.Error("Filter doesn't have the given instance name")
		}

		handler(&logging.Entry{
			Timestamp: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			Payload:   &samplePayload,
		})
		handler(&logging.Entry{
			Timestamp: time.Date(2006, 1, 2, 15, 4, 5, 12345, time.UTC),
			Payload:   &samplePayload,
		})
		return nil

	}

	// Start tests.
	var output *bytes.Buffer
	var scanner *bufio.Scanner
	cfg := &config.Config{}
	ctx := config.NewContext(context.Background(), cfg)

	// Without timestamp.
	output = &bytes.Buffer{}
	cmdLog(&logOpt{
		Context:      ctx,
		InstanceName: instance,
		Output:       output,
		Requester:    requester,
	})
	scanner = bufio.NewScanner(output)
	for scanner.Scan() {
		res := strings.TrimSpace(scanner.Text())
		if res != samplePayload.Log {
			t.Error("Log message is not correct:", res)
		}
	}

	// With timestamp.
	output = &bytes.Buffer{}
	cmdLog(&logOpt{
		Context:      ctx,
		InstanceName: instance,
		Timestamp:    true,
		Output:       output,
		Requester:    requester,
	})
	scanner = bufio.NewScanner(output)
	for scanner.Scan() {
		res := scanner.Text()
		if !strings.Contains(res, samplePayload.Log) {
			t.Error("Log message is not correct:", res)
		}
		if !strings.Contains(res, "2006/01/02") {
			t.Error("Timestampe in log message is not correct:", res)
		}
	}

}

// TestCmdLogWithReusedInstanceName tests cmdLog when a user reuses some instance name.
// In this case, cmdLog should output log entries only coming from the newest instance.
func TestCmdLogWithReusedInstanceName(t *testing.T) {

	instance := "test-instance"
	newPayload := log.RoadiePayload{
		Username:     "not used",
		Stream:       "stdout",
		Log:          "this log entries must be appeared",
		ContainerID:  "not used",
		InstanceName: instance,
	}

	var requester log.EntryRequesterFunc
	// Make a mock requester which doesn't requests but returns pre-defined log entries.
	requester = func(project, filter string, handler log.EntryHandler) error {

		if strings.Contains(filter, "GCE_OPERATION_DONE") {
			// If the request is for operation logs,
			// returns a dummy log of starting an instance.

			// Start an old instance.
			handler(&logging.Entry{
				Timestamp: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
				Payload: &log.ActivityPayload{
					EventType: "GCE_OPERATION_DONE",
					Resource: PayloadResource{
						Name: instance,
					},
					EventSubtype: log.EventSubtypeInsert,
				},
			})

			// Stop the old instance.
			handler(&logging.Entry{
				Timestamp: time.Date(2006, 2, 2, 15, 4, 5, 0, time.UTC),
				Payload: &log.ActivityPayload{
					EventType: "GCE_OPERATION_DONE",
					Resource: PayloadResource{
						Name: instance,
					},
					EventSubtype: log.EventSubtypeDelete,
				},
			})

			// Start a new instance.
			handler(&logging.Entry{
				Timestamp: time.Date(2006, 3, 2, 15, 4, 5, 0, time.UTC),
				Payload: &log.ActivityPayload{
					EventType: "GCE_OPERATION_DONE",
					Resource: PayloadResource{
						Name: instance,
					},
					EventSubtype: log.EventSubtypeInsert,
				},
			})

			return nil

		}

		// Otherwise, requests must be for instance logs.
		if !strings.Contains(filter, instance) {
			t.Error("Filter doesn't have the given instance name:", filter)
		}
		if !strings.Contains(filter, "timestamp > \"2006-03-02T15:04:05Z\"") {
			t.Error("Filter doesn't request newer log entries after the newest instance created:", filter)
		}

		handler(&logging.Entry{
			Timestamp: time.Date(2006, 4, 2, 15, 4, 5, 0, time.UTC),
			Payload:   &newPayload,
		})
		handler(&logging.Entry{
			Timestamp: time.Date(2006, 4, 2, 15, 4, 5, 12345, time.UTC),
			Payload:   &newPayload,
		})

		return nil

	}

	// Send a request.
	cfg := &config.Config{}
	ctx := config.NewContext(context.Background(), cfg)
	cmdLog(&logOpt{
		Context:      ctx,
		InstanceName: instance,
		Requester:    requester,
	})

}
