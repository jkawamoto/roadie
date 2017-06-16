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
	"fmt"
	"io"
	"os"
	"time"

	"github.com/urfave/cli"
)

// optLog defines arguments of log command.
type optLog struct {
	*Metadata
	// InstanceName of which logs are shown.
	InstanceName string
	// If true, timestamp is also printed.
	Timestamp bool
	// If true, keep waiting new logs.
	Follow bool
	// SleepTime defines sleep time.
	SleepTime time.Duration
	// From defines the time retriving log entries from.
	From time.Time
}

// CmdLog shows logs of a given instance.
func CmdLog(c *cli.Context) error {

	// Checking the number of arguments.
	if c.NArg() != 1 {
		fmt.Printf("expected 1 argument. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return err
	}

	// Run the log command.
	if err := cmdLog(&optLog{
		Metadata:     m,
		InstanceName: c.Args()[0],
		Timestamp:    !c.Bool("no-timestamp"),
		Follow:       c.Bool("follow"),
		SleepTime:    DefaultSleepTime,
	}); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil
}

// cmdLog retrieves and prints log entries according to the given options.
func cmdLog(opt *optLog) (err error) {

	log, err := opt.LogManager()
	if err != nil {
		return
	}

	for {

		err = log.Get(opt.Context, opt.InstanceName, opt.From, func(timestamp time.Time, line string, stderr bool) error {

			var msg string
			if opt.Timestamp {
				msg = fmt.Sprintf("%v %s", timestamp.Format(PrintTimeFormat), line)
			} else {
				msg = fmt.Sprintf("%s", line)
			}

			if !stderr {
				fmt.Fprintln(os.Stdout, msg)
			} else {
				fmt.Fprintln(os.Stderr, msg)
			}

			opt.From = timestamp
			return nil

		})

		if err != nil || !opt.Follow {
			break
		}

		select {
		case <-opt.Context.Done():
			return opt.Context.Err()
		case <-time.After(opt.SleepTime):
		}

	}

	if err == io.EOF {
		err = nil
	}
	return

}
