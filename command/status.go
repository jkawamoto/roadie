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

	"github.com/gosuri/uitable"
	"github.com/jkawamoto/roadie/chalk"
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

	m.Spinner.Prefix = "Loading information..."
	m.Spinner.Start()
	defer m.Spinner.Stop()

	compute, err := m.InstanceManager()
	if err != nil {
		return
	}

	table := uitable.New()
	table.AddRow("INSTANCE NAME", "STATUS")
	err = compute.Instances(m.Context, func(name, status string) (err error) {
		table.AddRow(name, status)
		return
	})
	if err != nil {
		return
	}

	m.Spinner.Stop()
	fmt.Println(table.String())
	return

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

	m.Spinner.Prefix = fmt.Sprintf("Killing instance %s...", instanceName)
	m.Spinner.FinalMSG = fmt.Sprintln("Killed Instance", instanceName)
	m.Spinner.Start()
	defer m.Spinner.Stop()

	compute, err := m.InstanceManager()
	if err != nil {
		return
	}

	if err = compute.DeleteInstance(m.Context, instanceName); err != nil {
		m.Spinner.FinalMSG = fmt.Sprintf(
			chalk.Red.Color("Cannot kill instance %s (%s)\n"), instanceName, err.Error())
	}
	return

}
