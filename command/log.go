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
	"github.com/jkawamoto/roadie/command/log"
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
	if err := cmdLog(&logOpt{
		Context:      config.NewContext(context.Background(), config.FromCliContext(c)),
		InstanceName: c.Args()[0],
		Timestamp:    !c.Bool("no-timestamp"),
		Follow:       c.Bool("follow"),
		Output:       os.Stdout,
	}); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil
}

// logOpt manages arguments for log command.
type logOpt struct {
	// Context, a config must be attached to.
	Context context.Context
	// InstanceName of which logs are shown.
	InstanceName string
	// If true, timestamp is also printed.
	Timestamp bool
	// If true, keep waiting new logs.
	Follow bool
	// io.Writer to be outputted logs.
	Output io.Writer
	// Used to obtain log entries.
	Requester log.EntryRequester
}

func cmdLog(opt *logOpt) (err error) {

	// Validate option.
	if opt.Output == nil {
		opt.Output = ioutil.Discard
	}
	if opt.Requester == nil {
		opt.Requester = log.NewCloudLoggingService(opt.Context)
	}

	// Determine when the newest instance starts.
	var start time.Time
	if err = log.GetOperationLogEntries(opt.Context, opt.Requester, func(timestamp time.Time, payload *log.ActivityPayload) (err error) {
		if payload.Resource.Name == opt.InstanceName {
			if payload.EventSubtype == log.EventSubtypeInsert {
				start = timestamp
			}
		}
		return
	}); err != nil {
		return
	}

	for {

		err = log.GetInstanceLogEntries(
			opt.Context, opt.InstanceName, start, opt.Requester, func(timestamp time.Time, payload *log.RoadiePayload) (err error) {

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
