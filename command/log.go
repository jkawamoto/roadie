package command

import (
	"fmt"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

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
			fmt.Printf("%v: %s\n", entry.Timestamp.Format("2006/01/02 15:04:05"), entry.Payload.Log)
		case err := <-chErr:
			fmt.Println(err.Error())
			break loop
		}
	}

	return nil
}
