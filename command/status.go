//
// command/status.go
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
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uitable"
	"github.com/jkawamoto/roadie-cli/util"
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

// getActivityPayload converts LogEntry's payload to a ActivityPayload.
func getActivityPayload(entry *LogEntry) (*ActivityPayload, error) {
	var res ActivityPayload
	if err := mapstructure.Decode(entry.Payload, &res); err != nil {
		return nil, err
	}
	return &res, nil
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

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "Loading information..."
	s.FinalMSG = fmt.Sprintf("\n%s\r", strings.Repeat(" ", len(s.Prefix)+2))
	s.Start()

loop:
	for {
		select {
		case entry := <-ch:

			if entry == nil {
				s.Stop()
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
			s.Stop()
			break loop
		}
	}

	table := uitable.New()
	table.AddRow("INSTANCE NAME", "STATUS")
	for name, status := range runnings {
		if status {
			table.AddRow(name, "running")
		} else {
			table.AddRow(name, "end")
		}
	}
	fmt.Println(table)

	return nil
}

// CmdStatusKill kills an instance.
func CmdStatusKill(c *cli.Context) error {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	b, err := util.NewInstanceBuilder(conf.Gcp.Project)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	name := c.Args()[0]
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Killing instance %s...", name)
	s.FinalMSG = fmt.Sprintf("\n%s\rKilled Instance %s.    \n", strings.Repeat(" ", len(s.Prefix)+2), name)
	s.Start()

	if err = b.StopInstance(name); err != nil {
		s.FinalMSG = fmt.Sprintf(
			chalk.Red.Color("\n%s\rCannot kill instance %s (%s)\n"),
			strings.Repeat(" ", len(s.Prefix)+2), name, err.Error())
	}
	s.Stop()

	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}
