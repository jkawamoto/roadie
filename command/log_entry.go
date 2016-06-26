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
