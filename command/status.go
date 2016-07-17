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
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uitable"
	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/jkawamoto/roadie/config"
	"github.com/mitchellh/mapstructure"
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

	if c.Bool("help") {
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	if err := cmdStatus(conf, c.Bool("all")); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// cmdStatus shows instance information. To obtain such information, `conf` is
// required. If all is true, print all instances otherwise information of
// instances of which results are deleted already will be omitted.
func cmdStatus(conf *config.Config, all bool) error {

	instances := make(map[string]struct{})
	if !all {

		err := ListupFiles(conf.Gcp.Project, conf.Gcp.Bucket, ResultPrefix,
			func(storage *util.Storage, file <-chan *util.FileInfo, done chan<- struct{}) {
				defer func() {
					done <- struct{}{}
				}()

				for {
					info := <-file
					if info == nil {
						return
					}

					rel, _ := filepath.Rel(ResultPrefix, info.Path)
					rel = filepath.Dir(rel)
					instances[rel] = struct{}{}

				}
			})

		if err != nil {
			return err
		}

	}

	ch := make(chan *LogEntry)
	chErr := make(chan error)

	go GetLogEntries(conf.Gcp.Project,
		"jsonPayload.event_type = \"GCE_OPERATION_DONE\"", ch, chErr)

	runnings := make(map[string]bool)

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "Loading information..."
	s.FinalMSG = fmt.Sprintf("\n%s\r", strings.Repeat(" ", len(s.Prefix)+2))
	s.Start()

	func() {

		for {
			select {
			case entry := <-ch:

				if entry == nil {
					return
				}

				// If all flag is not set show only following instances;
				//  - running instances,
				//  - ended but results are note deleted instances.
				if payload, err := getActivityPayload(entry); err == nil {

					switch payload.EventSubtype {
					case eventSubtypeInsert:
						runnings[payload.Resource.Name] = true

					case eventSubtypeDelete:
						runnings[payload.Resource.Name] = false
						if _, exist := instances[payload.Resource.Name]; !all && !exist {
							delete(runnings, payload.Resource.Name)
						}
					}

				} else {
					log.Println(chalk.Red.Color(err.Error()))
				}

			case err := <-chErr:
				fmt.Println(err.Error())
				return
			}
		}

	}()
	s.Stop()

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
	if err := cmdStatusKill(conf, c.Args()[0]); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// cmdStatusKill kills a given instance named `instanceName`.
func cmdStatusKill(conf *config.Config, instanceName string) (err error) {

	b, err := util.NewInstanceBuilder(conf.Gcp.Project)
	if err != nil {
		return
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Killing instance %s...", instanceName)
	s.FinalMSG = fmt.Sprintf("\n%s\rKilled Instance %s.\n", strings.Repeat(" ", len(s.Prefix)+2), instanceName)
	s.Start()
	defer s.Stop()

	if err = b.DeleteInstance(instanceName); err != nil {
		s.FinalMSG = fmt.Sprintf(
			chalk.Red.Color("\n%s\rCannot kill instance %s (%s)\n"),
			strings.Repeat(" ", len(s.Prefix)+2), instanceName, err.Error())
	}
	return

}
