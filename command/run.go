//
// command/run.go
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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// runOpt manages all arguments and flags defined in run command.
type runOpt struct {
	// Metadata to run a command.
	*Metadata
	// SourceOpt specifies options for source secrion of a script.
	SourceOpt
	// Path for the script file to be run.
	ScriptFile string
	// Arguments for the script.
	ScriptArgs []string
	// Instance name. If not set, named by script file name and current time.
	InstanceName string
	// Base docker image name.
	Image string
	// If true, result section will be overwritten so that roadie can manage
	// result data. Otherwise, users require to manage them by their self.
	OverWriteResultSection bool
}

// CmdRun specifies the behavior of `run` command.
func CmdRun(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf("expected 1 argument. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	opt := &runOpt{
		Metadata: m,
		SourceOpt: SourceOpt{
			Git:     c.String("git"),
			URL:     c.String("url"),
			Local:   c.String("local"),
			Exclude: c.StringSlice("exclude"),
			Source:  c.String("source"),
		},
		ScriptFile:   c.Args().First(),
		ScriptArgs:   c.StringSlice("e"),
		InstanceName: c.String("name"),
		Image:        c.String("image"),
		OverWriteResultSection: c.Bool("overwrite-result-section"),
	}

	currentTime := time.Now()
	err = cmdRun(opt)
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	if c.Bool("follow") {

		opt.Spinner.Prefix = "Waiting logging information..."
		opt.Spinner.FinalMSG = ""
		opt.Spinner.Start()
		defer opt.Spinner.Stop()

		select {
		case <-opt.Context.Done():
			return opt.Context.Err()
		case <-time.After(DefaultWaitTimeOfInstanceCreation):
			opt.Spinner.Stop()
		}

		return cmdLog(&optLog{
			Metadata:     m,
			InstanceName: opt.InstanceName,
			Timestamp:    true,
			Follow:       true,
			SleepTime:    DefaultSleepTime,
			After:        currentTime,
		})

	}
	return

}

// cmdRun implements the main logic of run command.
func cmdRun(opt *runOpt) (err error) {

	s, err := script.NewScriptWithArgs(opt.ScriptFile, opt.ScriptArgs)
	if err != nil {
		return
	}

	// Update instance name.
	// If an instance name is not given, use the default name.
	if opt.InstanceName != "" {
		s.Name = strings.ToLower(opt.InstanceName)
	}
	s.Image = opt.Image

	// Check a specified bucket exists and create it if not.
	service, err := opt.StorageManager()
	if err != nil {
		return
	}
	storage := cloud.NewStorage(service, opt.Spinner.Writer)

	// Update source section.
	err = UpdateSourceSection(opt.Metadata, s, &opt.SourceOpt, storage)
	if err != nil {
		return
	}

	// Update result section
	UpdateResultSection(s, opt.OverWriteResultSection, opt.Stdout)

	// Debugging info.
	opt.Logger.Printf("Script to be run:\n%s\n", s.String())

	opt.Spinner.Prefix = fmt.Sprintf("Creating instance %v...", chalk.Bold.TextStyle(s.Name))
	opt.Spinner.FinalMSG = fmt.Sprintf("Instance %v created.\n", chalk.Bold.TextStyle(s.Name))
	opt.Spinner.Start()
	defer opt.Spinner.Stop()

	instanceManager, err := opt.InstanceManager()
	if err != nil {
		return
	}

	err = instanceManager.CreateInstance(opt.Context, s)
	if err != nil {
		opt.Spinner.FinalMSG = fmt.Sprintf(chalk.Red.Color("\n%s\rCannot create instance.\n"), strings.Repeat(" ", len(opt.Spinner.Prefix)+2))
	}
	return

}
