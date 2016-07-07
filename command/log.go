//
// command/log.go
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
	"os"
	"strings"

	"github.com/jkawamoto/roadie/config"
	"github.com/mitchellh/mapstructure"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// RoadiePayload defines the payload structure of insance logs.
type RoadiePayload struct {
	Username     string
	Stream       string
	Log          string
	ContainerID  string `mapstructure:"container_id"`
	InstanceName string `mapstructure:"instance_name"`
}

// PrintTimeFormat defines time format to be used to print logs.
const PrintTimeFormat = "2006/01/02 15:04:05"

// CmdLog shows logs of a given instance.
func CmdLog(c *cli.Context) error {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	name := c.Args()[0]
	timestamp := !c.Bool("no-timestamp")
	if err := cmdLog(conf, name, timestamp); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil
}

func cmdLog(conf *config.Config, instanceName string, timestamp bool) error {

	ch := make(chan *LogEntry)
	chErr := make(chan error)

	filter := fmt.Sprintf(
		// Instead of logName, which is specified TAG env in roadie-gce,
		// use instance name to distinguish instances. This update makes all logs
		// will have same log name, docker, so that such log can be stored into
		// GCS easily.
		//
		// "resource.type = \"gce_instance\" AND logName = \"projects/%s/logs/%s\"", project, name),
		"resource.type = \"gce_instance\" AND jsonPayload.instance_name = \"%s\"", instanceName)

	go GetLogEntries(conf.Gcp.Project, filter, ch, chErr)

	return func() error {

		// Listening channels and print logs.
		for {
			select {
			case entry := <-ch:

				if entry == nil {
					return nil
				}

				if payload, err := getRoadiePayload(entry); err == nil {

					var msg string
					if timestamp {
						msg = fmt.Sprintf("%v: %s\n", entry.Timestamp.Format(PrintTimeFormat), payload.Log)
					} else {
						msg = fmt.Sprintf("%s\n", payload.Log)
					}

					if payload.Stream == "stdout" {
						fmt.Println(msg)
					} else {
						fmt.Fprintln(os.Stderr, msg)
					}

				} else {
					log.Println(chalk.Red.Color(err.Error()))
				}

			case err := <-chErr:
				return err
			}
		}
	}()

}

// getRoadiePayload converts LogEntry's payload to a RoadiePayload.
func getRoadiePayload(entry *LogEntry) (*RoadiePayload, error) {

	var res RoadiePayload
	if err := mapstructure.Decode(entry.Payload, &res); err != nil {
		return nil, err
	}
	res.Log = strings.TrimRight(res.Log, "\n")

	return &res, nil
}
