package command

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/command/log"

	logging "google.golang.org/api/logging/v2beta1"
)

// resource a structure used in ActivityPayload.Resource.
type resource struct {
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

	var requester log.LogEntryRequesterFunc
	// Make a mock requester which doesn't requests but returns pre-defined log entries.
	requester = func(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {

		if strings.Contains(req.Filter, "GCE_OPERATION_DONE") {
			// If the request is for operation logs,
			// returns a dummy log of starting an instance.
			return &logging.ListLogEntriesResponse{
				Entries: []*logging.LogEntry{
					&logging.LogEntry{
						JsonPayload: log.ActivityPayload{
							EventType: "GCE_OPERATION_DONE",
							Resource: resource{
								Name: instance,
							},
							EventSubtype: log.EventSubtypeInsert,
						},
						Timestamp: "2006-01-02T15:04:05Z",
					},
				}}, nil

		}

		// Otherwise, requests must be for instance logs.
		if !strings.Contains(req.Filter, instance) {
			t.Error("Filter doesn't have the given instance name")
		}

		return &logging.ListLogEntriesResponse{
			Entries: []*logging.LogEntry{
				&logging.LogEntry{
					JsonPayload: samplePayload,
					Timestamp:   "2006-01-02T15:04:05Z",
				},
				&logging.LogEntry{
					JsonPayload: samplePayload,
					Timestamp:   "2006-01-02T15:04:05.12345Z",
				},
			}}, nil
	}

	// Start tests.
	var output *bytes.Buffer
	var scanner *bufio.Scanner

	// Without timestamp.
	output = &bytes.Buffer{}
	cmdLog(&logOpt{
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

	var requester log.LogEntryRequesterFunc
	// Make a mock requester which doesn't requests but returns pre-defined log entries.
	requester = func(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {

		if strings.Contains(req.Filter, "GCE_OPERATION_DONE") {
			// If the request is for operation logs,
			// returns a dummy log of starting an instance.
			return &logging.ListLogEntriesResponse{
				Entries: []*logging.LogEntry{
					// Start an old instance.
					&logging.LogEntry{
						JsonPayload: log.ActivityPayload{
							EventType: "GCE_OPERATION_DONE",
							Resource: resource{
								Name: instance,
							},
							EventSubtype: log.EventSubtypeInsert,
						},
						Timestamp: "2006-01-02T15:04:05Z",
					},
					// Stop the old instance.
					&logging.LogEntry{
						JsonPayload: log.ActivityPayload{
							EventType: "GCE_OPERATION_DONE",
							Resource: resource{
								Name: instance,
							},
							EventSubtype: log.EventSubtypeDelete,
						},
						Timestamp: "2006-02-02T15:04:05Z",
					},
					// Start a new instance.
					&logging.LogEntry{
						JsonPayload: log.ActivityPayload{
							EventType: "GCE_OPERATION_DONE",
							Resource: resource{
								Name: instance,
							},
							EventSubtype: log.EventSubtypeInsert,
						},
						Timestamp: "2006-03-02T15:04:05Z",
					},
				}}, nil

		}

		// Otherwise, requests must be for instance logs.
		if !strings.Contains(req.Filter, instance) {
			t.Error("Filter doesn't have the given instance name:", req.Filter)
		}
		if !strings.Contains(req.Filter, "timestamp > \"2006-03-02T15:04:05Z\"") {
			t.Error("Filter doesn't request newer log entries after the newest instance created:", req.Filter)
		}

		return &logging.ListLogEntriesResponse{
			Entries: []*logging.LogEntry{
				&logging.LogEntry{
					JsonPayload: newPayload,
					Timestamp:   "2006-04-02T15:04:05Z",
				},
				&logging.LogEntry{
					JsonPayload: newPayload,
					Timestamp:   "2006-04-02T15:04:05.12345Z",
				},
			}}, nil
	}

	// Send a request.
	cmdLog(&logOpt{
		InstanceName: instance,
		Requester:    requester,
	})

}
