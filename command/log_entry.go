package command

import (
	"log"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/ttacon/chalk"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/logging/v2beta1"
)

// RoadieLogEntry is a structure of one log entry.
type RoadieLogEntry struct {
	Timestamp time.Time
	Payload   struct {
		Username     string
		Stream       string
		Log          string
		ContainerID  string `mapstructure:"container_id"`
		InstanceName string `mapstructure:"instance_name"`
	}
}

// GetLogEntries downloads log entries as a goroutine.
func GetLogEntries(project, filter string, ch chan *RoadieLogEntry, chErr chan error) {

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

				timestamp, err := time.Parse("2006-01-02T15:04:05.000Z", v.Timestamp)
				if err != nil {
					log.Println(chalk.Red.Color(err.Error()))
					continue
				}

				// Converting general interface to specific structure.
				entry := &RoadieLogEntry{Timestamp: timestamp}
				if err := mapstructure.Decode(v.JsonPayload, &entry.Payload); err != nil {
					log.Println(chalk.Red.Color(err.Error()))
					continue
				}

				entry.Payload.Log = strings.TrimRight(entry.Payload.Log, "\n")
				ch <- entry

			}
		}

		pageToken = res.NextPageToken
		if pageToken == "" {
			break
		}

	}

	ch <- nil

}
