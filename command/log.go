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
	"time"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/config"
	"github.com/mitchellh/mapstructure"
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

// CmdLog shows logs of a given instance.
func CmdLog(c *cli.Context) error {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	name := c.Args()[0]
	timestamp := !c.Bool("no-timestamp")
	follow := c.Bool("follow")
	if err := cmdLog(conf, name, timestamp, follow); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil
}

func cmdLog(conf *config.Config, instanceName string, timestamp, follow bool) (err error) {

	ch := make(chan *LogEntry)
	chErr := make(chan error)

	baseFilter := fmt.Sprintf(
		// Instead of logName, which is specified TAG env in roadie-gce,
		// use instance name to distinguish instances. This update makes all logs
		// will have same log name, docker, so that such log can be stored into
		// GCS easily.
		//
		// "resource.type = \"gce_instance\" AND logName = \"projects/%s/logs/%s\"", project, name),
		"resource.type = \"gce_instance\" AND jsonPayload.instance_name = \"%s\"", instanceName)

	filter := baseFilter
	var lastTimestamp *time.Time
	for {

		go GetLogEntries(conf.Gcp.Project, filter, ch, chErr)

		err = func() error {

			// Listening channels and print logs.
			for {
				select {
				case entry := <-ch:

					if entry == nil {
						return nil
					}

					if payload, err2 := getRoadiePayload(entry); err2 == nil {

						var msg string
						if timestamp {
							msg = fmt.Sprintf("%v: %s", entry.Timestamp.Format(PrintTimeFormat), payload.Log)
						} else {
							msg = fmt.Sprintf("%s", payload.Log)
						}

						if payload.Stream == "stdout" {
							fmt.Println(msg)
						} else {
							fmt.Fprintln(os.Stderr, msg)
						}

						lastTimestamp = &entry.Timestamp

					} else {
						log.Println(chalk.Red.Color(err2.Error()))
					}

				case err := <-chErr:
					return err
				}
			}
		}()

		if err != nil || !follow {
			break
		}

		time.Sleep(30 * time.Second)

		// TODO: check status each loop
		utc := lastTimestamp.In(time.UTC)
		filter = fmt.Sprintf("%s AND timestamp > \"%s\"", baseFilter, utc.Format(LogTimeFormat))

	}
	return

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
