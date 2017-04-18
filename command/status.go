//
// command/status.go
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
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uitable"
	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
	"github.com/urfave/cli"
)

// CmdStatus shows status of instances.
func CmdStatus(c *cli.Context) error {

	if c.Bool("help") {
		return cli.ShowSubcommandHelp(c)
	}

	m := getMetadata(c)
	if err := cmdStatus(m, c.Bool("all")); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// cmdStatus shows instance information. If all is true, print all instances otherwise information of
// instances of which results are deleted already will be omitted.
func cmdStatus(m *Metadata, all bool) (err error) {

	var runningInstances []string
	var terminatedInstances []string

	err = func() (err error) {
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Prefix = "Loading information..."
		s.FinalMSG = fmt.Sprintf("\n%s\r", strings.Repeat(" ", len(s.Prefix)+2))
		s.Start()
		defer s.Stop()

		compute, err := m.InstanceManager()
		if err != nil {
			return
		}
		instances, err := compute.Instances(m.Context)
		if err != nil {
			return err
		}
		for name := range instances {
			runningInstances = append(runningInstances, name)
		}
		sort.Strings(runningInstances)

		service, err := m.StorageManager()
		if err != nil {
			return err
		}
		storage := cloud.NewStorage(service, nil)

		var prev string
		err = storage.ListupFiles(m.Context, script.ResultPrefix, "", func(info *cloud.FileInfo) error {
			select {
			case <-m.Context.Done():
				return m.Context.Err()
			default:
			}

			rel := filepath.Dir(info.Path)
			if prev != rel {
				terminatedInstances = append(terminatedInstances, rel)
				prev = rel
			}
			return nil
		})
		if err != nil {
			return err
		}
		sort.Strings(terminatedInstances)
		return
	}()
	if err != nil {
		return err
	}

	table := uitable.New()
	table.AddRow("INSTANCE NAME", "STATUS")
	for _, name := range runningInstances {
		table.AddRow(name, "running")
	}
	for _, name := range terminatedInstances {
		table.AddRow(name, "terminated")
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

	m := getMetadata(c)
	if err := cmdStatusKill(m, c.Args().First()); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// cmdStatusKill kills a given instance named `instanceName`.
func cmdStatusKill(m *Metadata, instanceName string) (err error) {

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Killing instance %s...", instanceName)
	s.FinalMSG = fmt.Sprintf("\n%s\rKilled Instance %s.\n", strings.Repeat(" ", len(s.Prefix)+2), instanceName)
	s.Start()
	defer s.Stop()

	compute, err := m.InstanceManager()
	if err != nil {
		return
	}

	if err = compute.DeleteInstance(m.Context, instanceName); err != nil {
		s.FinalMSG = fmt.Sprintf(
			chalk.Red.Color("\n%s\rCannot kill instance %s (%s)\n"),
			strings.Repeat(" ", len(s.Prefix)+2), instanceName, err.Error())
	}
	return

}
