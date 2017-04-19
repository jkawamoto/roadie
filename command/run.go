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
	"os"
	"strings"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/jkawamoto/roadie/script"
	"github.com/urfave/cli"
)

// runOpt manages all arguments and flags defined in run command.
type runOpt struct {
	// Metadata to run a command.
	*Metadata
	// SourceOpt specifies options for source secrion of a script.
	util.SourceOpt
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
	// If true, created instance will not shutdown automatically. So, users
	// require to do it by their self. This flag can be useful for debugging.
	NoShutdown bool
	// If true, do not create any instances but show startup script.
	// This flag is for debugging.
	Dry bool
	// The number of times retry roadie-gcp container when GCP's error happens.
	Retry int64
}

// CmdRun specifies the behavior of `run` command.
func CmdRun(c *cli.Context) error {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m := getMetadata(c)
	opt := &runOpt{
		Metadata: m,
		SourceOpt: util.SourceOpt{
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
		NoShutdown:             c.Bool("no-shutdown"),
		Dry:                    c.Bool("dry"),
		Retry:                  c.Int64("retry") + 1,
	}
	if err := cmdRun(opt); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	if c.Bool("follow") {
		return cmdLog(&optLog{
			Metadata:     m,
			InstanceName: opt.InstanceName,
			Timestamp:    true,
			Follow:       true,
			SleepTime:    DefaultSleepTime,
		})
	}
	return nil

}

// cmdRun implements the main logic of run command.
func cmdRun(opt *runOpt) (err error) {

	s, err := script.NewScript(opt.ScriptFile, opt.ScriptArgs)
	if err != nil {
		return
	}

	// Update instance name.
	// If an instance name is not given, use the default name.
	if opt.InstanceName != "" {
		s.InstanceName = strings.ToLower(opt.InstanceName)
	}

	// Check a specified bucket exists and create it if not.
	service, err := opt.StorageManager()
	if err != nil {
		return err
	}
	storage := cloud.NewStorage(service, nil)

	// Update source section.
	err = util.UpdateSourceSection(opt.Context, s, &opt.SourceOpt, storage, os.Stdout)
	if err != nil {
		return
	}

	// Update result section
	util.UpdateResultSection(s, opt.OverWriteResultSection, os.Stdout)

	// Debugging info.
	opt.Logger.Printf("Script to be run:\n%s\n", s.String())

	// Prepare options.
	if opt.NoShutdown {
		s.Options = append(s.Options, "no-shutdown")
	}
	if opt.Retry <= 0 {
		opt.Retry = 10
	}
	s.Options = append(s.Options, fmt.Sprintf("retry:%d", opt.Retry))

	opt.Spinner.Prefix = fmt.Sprintf("Creating an instance named %s...", chalk.Bold.TextStyle(s.InstanceName))
	opt.Spinner.FinalMSG = fmt.Sprintf("\n%s\rInstance created.\n", strings.Repeat(" ", len(opt.Spinner.Prefix)+2))
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
