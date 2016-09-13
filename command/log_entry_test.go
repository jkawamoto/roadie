package command

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"
	logging "google.golang.org/api/logging/v2beta1"
)

func TestGetLogEntries(t *testing.T) {

	// Test giving project name and filter are passed to requestDo.
	project := "test-project"
	filter := "test-filter"
	getLogEntries(context.Background(), project, filter,
		func(_ *LogEntry) error {
			return nil
		},
		func(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {

			// Checking project id.
			t.Log("ProjectIds is", req.ProjectIds)
			exist := false
			for _, id := range req.ProjectIds {
				if id == project {
					exist = true
				}
			}
			if !exist {
				t.Error("ProjectIds doesn't have the giving project id")
			}

			// Checking filter.
			t.Log("Filter is", req.Filter)
			if req.Filter != filter {
				t.Error("Filter doesn't match the giving filter")
			}

			return &logging.ListLogEntriesResponse{}, nil
		})

	// Test giving entries are passed to handler.
	samplePayload := struct{ Key, Value string }{Key: "key", Value: "value"}
	getLogEntries(context.Background(), project, filter,
		func(entry *LogEntry) error {

			t.Log("Timestamp is", entry.Timestamp)
			if entry.Timestamp.Year() != 2006 {
				t.Error("Timestamp is not correct")
			}

			if entry.Payload != samplePayload {
				t.Error("Payload doesn't match a passing payload")
			}

			return nil
		},
		func(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {

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

		})

	// Test giving token will be used and handler will be received entries gived from
	// another page.
	var invoked int
	getLogEntries(context.Background(), project, filter,
		func(entry *LogEntry) error {

			invoked++
			t.Log("Timestamp is", entry.Timestamp)
			if entry.Timestamp.Year() != 2006 {
				t.Error("Timestamp is not correct")
			}

			if entry.Payload != samplePayload {
				t.Error("Payload doesn't match a passing payload")
			}
			return nil
		},
		func() func(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {
			var counter int
			token := "next-token"
			return func(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {

				if counter > 1 {
					t.Error("requestDo is called too many times.")
				}

				if (counter == 0 && req.PageToken != "") || (counter != 0 && req.PageToken != token) {
					t.Error("Wrong page token is set")
				}

				if counter != 0 {
					token = ""
				}
				counter++

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
					},
					NextPageToken: token,
				}, nil

			}
		}())
	// Checking how many times handker called.
	if invoked != 4 {
		t.Error("NextPageToken doesn't work")
	}

}

func TestStopGetLogEntries(t *testing.T) {

	// Test getLogEntries will be canceled when handler returns non nil values.
	var invoked int
	project := "test-project"
	filter := "test-filter"
	getLogEntries(context.Background(), project, filter,
		func(entry *LogEntry) error {
			invoked++
			return fmt.Errorf("Test error.")
		},
		func(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {

			return &logging.ListLogEntriesResponse{
				Entries: []*logging.LogEntry{
					&logging.LogEntry{
						JsonPayload: "samplePayload",
						Timestamp:   "2006-01-02T15:04:05Z",
					},
					&logging.LogEntry{
						JsonPayload: "samplePayload",
						Timestamp:   "2006-01-02T15:04:05.12345Z",
					},
				}}, nil

		})

	if invoked != 1 {
		t.Error("handler reutns some error but getLogEntries didn't stop")
	}

}

func TestCancelGetLogEntries(t *testing.T) {

	// Test getLogEntries will be canceled via context.
	project := "test-project"
	filter := "test-filter"
	ctx, cancel := context.WithCancel(context.Background())

	var invoked int
	getLogEntries(ctx, project, filter,
		func(entry *LogEntry) error {
			invoked++
			cancel()
			return nil
		},
		func(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {

			return &logging.ListLogEntriesResponse{
				Entries: []*logging.LogEntry{
					&logging.LogEntry{
						JsonPayload: "samplePayload",
						Timestamp:   "2006-01-02T15:04:05Z",
					},
					&logging.LogEntry{
						JsonPayload: "samplePayload",
						Timestamp:   "2006-01-02T15:04:05.12345Z",
					},
				}}, nil

		})

	if invoked != 1 {
		t.Error("context was canceled but getLogEntries didn't stop")
	}

}
