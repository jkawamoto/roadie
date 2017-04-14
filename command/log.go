//
// command/log.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
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
// along with Roadie.  If not, see <http://www.gnu.org/licenses/>.
//

package command

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/cloud/gce"
	"github.com/jkawamoto/roadie/command/util"
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
		Context:      util.GetContext(c),
		InstanceName: c.Args()[0],
		Timestamp:    !c.Bool("no-timestamp"),
		Follow:       c.Bool("follow"),
		Output:       os.Stdout,
		Config:       config.FromCliContext(c),
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
	// Config.
	Config *config.Config
}

func cmdLog(opt *logOpt) (err error) {

	// Validate option.
	if opt.Output == nil {
		opt.Output = ioutil.Discard
	}

	var from time.Time
	log := gce.NewLogManager(&opt.Config.GcpConfig)

	for {

		err = log.Get(opt.Context, opt.InstanceName, from, func(timestamp time.Time, line string, stderr bool) error {

			var msg string
			if opt.Timestamp {
				msg = fmt.Sprintf("%v: %s", timestamp.Format(PrintTimeFormat), line)
			} else {
				msg = fmt.Sprintf("%s", line)
			}

			if !stderr {
				fmt.Fprintln(opt.Output, msg)
			} else {
				fmt.Fprintln(os.Stderr, msg)
			}

			from = timestamp
			return nil

		})
		if err != nil {
			return
		}

		if !opt.Follow {
			break
		}
		time.Sleep(30 * time.Second)

	}
	return

}
