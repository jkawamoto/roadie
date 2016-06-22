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

// CmdLog shows logs of a given instance.
func CmdLog(c *cli.Context) error {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected at most 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	name := c.Args()[0]

	ch := make(chan *RoadieLogEntry)
	chErr := make(chan error)
	go getLogEntries(conf.Gcp.Project, name, ch, chErr)

loop:
	for {
		select {
		case entry := <-ch:
			if entry == nil {
				break loop
			}
			fmt.Printf("%v: %s\n", entry.Timestamp.Format("2006/01/02 15:04:05"), entry.Payload.Log)
		case err := <-chErr:
			fmt.Println(err.Error())
			break loop
		}
	}

	return nil
}

func getLogEntries(project, name string, ch chan *RoadieLogEntry, chErr chan error) {

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
			ProjectIds: []string{project},
			Filter: fmt.Sprintf(
				// Instead of logName, which is specified TAG env in roadie-gce,
				// use instance name to distinguish instances. This update makes all logs
				// will have same log name, docker, so that such log can be stored into
				// GCS easily.
				//
				// "resource.type = \"gce_instance\" AND logName = \"projects/%s/logs/%s\"", project, name),
				"resource.type = \"gce_instance\" AND jsonPayload.instance_name = \"%s\"", name),
			PageToken: pageToken,
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
