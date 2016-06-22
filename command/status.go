package command

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
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

func CmdStatus(c *cli.Context) error {
	// Write your code here
	return nil
}

func Logging() {

	ch := make(chan *RoadieLogEntry)
	chErr := make(chan error)

	go getLogEntries(ch, chErr)

	fmt.Println("Current Logs:")
loop:
	for {
		select {
		case entry := <-ch:
			if entry == nil {
				break loop
			}
			fmt.Printf("  %v: %s\n", entry.Timestamp.Format("2006/01/02 15:04:05"), entry.Payload.Log)
		case err := <-chErr:
			fmt.Println(err.Error())
			break loop
		}
	}

}

func getLogEntries(ch chan *RoadieLogEntry, chErr chan error) {

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
			ProjectIds: []string{"jkawamoto-ppls"},
			Filter:     "resource.type = \"gce_instance\" AND logName = \"projects/jkawamoto-ppls/logs/docker\"",
			PageToken:  pageToken,
		}).Do()
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
