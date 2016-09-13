package command

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	logging "google.golang.org/api/logging/v2beta1"

	"github.com/jkawamoto/roadie/config"
)

func TestCmdLog(t *testing.T) {

	instance := "test-instance"
	samplePayload := RoadiePayload{
		Username:     "not used",
		Stream:       "stdout",
		Log:          "sample log output",
		ContainerID:  "not used",
		InstanceName: instance,
	}

	var requester LogEntryRequesterFunc
	requester = func(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {

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
	cmdLog(&config.Config{}, &logOpt{
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
	cmdLog(&config.Config{}, &logOpt{
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
