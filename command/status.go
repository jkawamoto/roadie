//
// command/status.go
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
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uitable"
	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/command/log"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/jkawamoto/roadie/config"
	"github.com/urfave/cli"
)

// CmdStatus shows status of instances.
func CmdStatus(c *cli.Context) error {

	if c.Bool("help") {
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	if err := cmdStatus(conf, c.Bool("all")); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// cmdStatus shows instance information. To obtain such information, `conf` is
// required. If all is true, print all instances otherwise information of
// instances of which results are deleted already will be omitted.
func cmdStatus(conf *config.Config, all bool) error {

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "Loading information..."
	s.FinalMSG = fmt.Sprintf("\n%s\r", strings.Repeat(" ", len(s.Prefix)+2))
	s.Start()

	instances := make(map[string]struct{})
	if !all {

		err := util.ListupFiles(conf.Gcp.Project, conf.Gcp.Bucket, ResultPrefix,
			func(storage *util.Storage, file <-chan *util.FileInfo, done chan<- struct{}) {
				defer func() {
					done <- struct{}{}
				}()

				for {
					info := <-file
					if info == nil {
						return
					}

					rel, _ := filepath.Rel(ResultPrefix, info.Path)
					rel = filepath.Dir(rel)
					instances[rel] = struct{}{}

				}
			})

		if err != nil {
			return err
		}

	}

	runnings := make(map[string]bool)
	ctx := context.Background()
	requester, _ := log.NewCloudLoggingService(ctx)

	err := log.GetOperationLogEntries(ctx, conf.Gcp.Project, requester, func(_ time.Time, payload *log.ActivityPayload) (err error) {

		// If all flag is not set, show only following instances;
		//  - running instances,
		//  - ended but results are note deleted instances.
		switch payload.EventSubtype {
		case log.EventSubtypeInsert:
			runnings[payload.Resource.Name] = true

		case log.EventSubtypeDelete:
			runnings[payload.Resource.Name] = false
			if _, exist := instances[payload.Resource.Name]; !all && !exist {
				delete(runnings, payload.Resource.Name)
			}
		}

		return

	})
	s.Stop()
	if err != nil {
		return err
	}

	table := uitable.New()
	table.AddRow("INSTANCE NAME", "STATUS")
	for name, status := range runnings {
		if status {
			table.AddRow(name, "running")
		} else {
			table.AddRow(name, "end")
		}
	}
	fmt.Println(table)

	return nil
}

// CmdStatusKill kills an instance.
func CmdStatusKill(c *cli.Context) error {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	if err := cmdStatusKill(conf, c.Args()[0]); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// cmdStatusKill kills a given instance named `instanceName`.
func cmdStatusKill(conf *config.Config, instanceName string) (err error) {

	b, err := util.NewInstanceBuilder(conf.Gcp.Project)
	if err != nil {
		return
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Killing instance %s...", instanceName)
	s.FinalMSG = fmt.Sprintf("\n%s\rKilled Instance %s.\n", strings.Repeat(" ", len(s.Prefix)+2), instanceName)
	s.Start()
	defer s.Stop()

	if err = b.DeleteInstance(instanceName); err != nil {
		s.FinalMSG = fmt.Sprintf(
			chalk.Red.Color("\n%s\rCannot kill instance %s (%s)\n"),
			strings.Repeat(" ", len(s.Prefix)+2), instanceName, err.Error())
	}
	return

}
