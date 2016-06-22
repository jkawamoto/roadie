package command

import (
	"fmt"
	"log"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uitable"
	"github.com/mitchellh/mapstructure"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

const (
	eventSubtypeInsert = "compute.instances.insert"
	eventSubtypeDelete = "compute.instances.delete"
)

// ActivityPayload defines the payload structure of activity log.
type ActivityPayload struct {
	EventTimestampUs string `mapstructure:"event_timestamp_us"`
	EventType        string `mapstructure:"vent_type"`
	TraceID          string `mapstructure:"trace_id"`
	Actor            struct {
		User string
	}
	Resource struct {
		Zone string
		Type string
		ID   string
		Name string
	}
	Version      string
	EventSubtype string `mapstructure:"event_subtype"`
	Operation    struct {
		Zone string
		Type string
		ID   string
		Name string
	}
}

// CmdStatus shows status of instances.
func CmdStatus(c *cli.Context) error {

	conf := GetConfig(c)
	ch := make(chan *LogEntry)
	chErr := make(chan error)

	// TODO: filter by jsonPayload.actor.user.
	// jsonPayload.actor.user is an email address of instance owner.
	// To omit other users instance, filter logs by such data.
	go GetLogEntries(conf.Gcp.Project,
		"jsonPayload.event_type = \"GCE_OPERATION_DONE\"", ch, chErr)

	runnings := make(map[string]bool)

	s := spinner.New(spinner.CharSets[23], 100*time.Millisecond)
	s.FinalMSG = "\n"
	s.Start()

loop:
	for {
		select {
		case entry := <-ch:

			if entry == nil {
				break loop
			}

			if payload, err := getActivityPayload(entry); err == nil {

				switch payload.EventSubtype {
				case eventSubtypeInsert:
					runnings[payload.Resource.Name] = true
				case eventSubtypeDelete:
					runnings[payload.Resource.Name] = false
				}

			} else {
				log.Println(chalk.Red.Color(err.Error()))
			}

		case err := <-chErr:
			fmt.Println(err.Error())
			break loop
		}
	}

	s.Stop()

	table := uitable.New()
	table.AddRow("INSTANCE NAME", "STATUS")
	for name, status := range runnings {
		if status {
			table.AddRow(name, "running")
		} else {
			table.AddRow(name, "stop")
		}
	}
	fmt.Println(table)

	return nil
}

// getActivityPayload converts LogEntry's payload to a ActivityPayload.
func getActivityPayload(entry *LogEntry) (*ActivityPayload, error) {
	var res ActivityPayload
	if err := mapstructure.Decode(entry.Payload, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
