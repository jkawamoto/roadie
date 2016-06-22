package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

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
		fmt.Printf(chalk.Red.Color("expected at most 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	name := c.Args()[0]

	ch := make(chan *LogEntry)
	chErr := make(chan error)
	go GetLogEntries(
		conf.Gcp.Project,
		fmt.Sprintf(
			// Instead of logName, which is specified TAG env in roadie-gce,
			// use instance name to distinguish instances. This update makes all logs
			// will have same log name, docker, so that such log can be stored into
			// GCS easily.
			//
			// "resource.type = \"gce_instance\" AND logName = \"projects/%s/logs/%s\"", project, name),
			"resource.type = \"gce_instance\" AND jsonPayload.instance_name = \"%s\"", name),
		ch, chErr)

loop:
	for {
		select {
		case entry := <-ch:

			if entry == nil {
				break loop
			}

			if payload, err := getRoadiePayload(entry); err == nil {
				fmt.Printf("%v: %s\n", entry.Timestamp.Format("2006/01/02 15:04:05"), payload.Log)
			} else {
				log.Println(chalk.Red.Color(err.Error()))
			}

		case err := <-chErr:
			fmt.Println(err.Error())
			break loop
		}
	}

	return nil
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
