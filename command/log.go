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
	"io"
	"io/ioutil"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/config"
	"github.com/urfave/cli"
)

// CmdLog shows logs of a given instance.
func CmdLog(c *cli.Context) error {

	// Checking the number of arguments.
	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	// Run the log command.
	if err := cmdLog(GetConfig(c), &logOpt{
		InstanceName: c.Args()[0],
		Timestamp:    !c.Bool("no-timestamp"),
		Follow:       c.Bool("follow"),
		Output:       os.Stdout,
	}); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil
}

// TODO: logOpt should has config.Config.
// logOpt manages arguments for log command.
type logOpt struct {
	// InstanceName of which logs are shown.
	InstanceName string
	// If true, timestamp is also printed.
	Timestamp bool
	// If true, keep waiting new logs.
	Follow bool
	// io.Writer to be outputted logs.
	Output io.Writer
	// Context, default is context.Background.
	Context context.Context
	// Used to obtain log entries.
	Requester LogEntryRequester
}

func cmdLog(conf *config.Config, opt *logOpt) (err error) {

	// Validate option.
	if opt.Output == nil {
		opt.Output = ioutil.Discard
	}
	if opt.Context == nil {
		opt.Context = context.Background()
	}
	if opt.Requester == nil {
		opt.Requester, err = NewCloudLoggingService(opt.Context)
		if err != nil {
			return err
		}
	}

	// Determine when the newest instance starts.
	var start time.Time
	if err = GetOperationLogEntries(opt.Context, conf.Gcp.Project, opt.Requester, func(timestamp time.Time, payload *ActivityPayload) (err error) {
		if payload.Resource.Name == opt.InstanceName {
			if payload.EventSubtype == EventSubtypeInsert {
				start = timestamp
			}
		}
		return
	}); err != nil {
		return
	}

	for {

		err = GetInstanceLogEntries(
			opt.Context, conf.Gcp.Project, opt.InstanceName, start, opt.Requester, func(timestamp time.Time, payload *RoadiePayload) (err error) {

				var msg string
				if opt.Timestamp {
					msg = fmt.Sprintf("%v: %s", timestamp.Format(PrintTimeFormat), payload.Log)
				} else {
					msg = fmt.Sprintf("%s", payload.Log)
				}

				if payload.Stream == "stdout" {
					fmt.Fprintln(opt.Output, msg)
				} else {
					fmt.Fprintln(os.Stderr, msg)
				}

				start = timestamp
				return

			})
		if err != nil {
			break
		}

		if !opt.Follow {
			break
		}
		time.Sleep(30 * time.Second)

	}
	return

}
