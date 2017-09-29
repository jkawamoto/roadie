//
// command/config.go
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
	"strings"

	"github.com/gosuri/uitable"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

const (
	// MsgNotSet is the message to be printed when some config isn't set.
	MsgNotSet = "Not set"
)

// CmdConfigProject shows or sets project ID to config file.
func CmdConfigProject(c *cli.Context) error {
	if c.Bool("help") {
		return cli.ShowSubcommandHelp(c)
	}
	if c.NArg() != 0 {
		fmt.Printf("expected no arguments. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}
	return CmdConfigProjectShow(c)
}

// CmdConfigProjectSet implements `config project set` command.
func CmdConfigProjectSet(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf("expected 1 argument. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	err = cmdConfigProjectSet(m, c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 4)
	}
	return

}

// cmdConfigProjectSet sets a new project name to the configuration file.
func cmdConfigProjectSet(m *Metadata, name string) (err error) {

	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	if strings.Contains(name, " ") {
		fmt.Fprintln(m.Stdout, chalk.Red.Color("The given project ID has spaces. They are replaced to '_'."))
		name = strings.Replace(name, " ", "_", -1)
	}
	if id := resource.GetProjectID(); id == "" {
		fmt.Fprintf(m.Stdout, "Set project ID:\n  %s\n", chalk.Green.Color(name))
	} else {
		fmt.Fprintf(m.Stdout, "Update project ID:\n  %s -> %s\n", id, chalk.Green.Color(name))
	}
	resource.SetProjectID(name)

	err = m.Config.Save()
	if err != nil {
		err = fmt.Errorf("cannot save the configuration to %q: %v", m.Config.FileName, err)
	}
	return

}

// CmdConfigProjectShow implements `config project show` command.
func CmdConfigProjectShow(c *cli.Context) (err error) {

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	err = cmdConfigProjectShow(m)
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdConfigProjectShow prints the current project ID.
func cmdConfigProjectShow(m *Metadata) (err error) {

	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	if id := resource.GetProjectID(); id != "" {
		fmt.Fprintln(m.Stdout, id)
	} else {
		fmt.Fprintln(m.Stdout, chalk.Red.Color(MsgNotSet))
	}
	return

}

// CmdConfigMachineType shows current configuration of machine type,
// or show help message when either -h or --help flag is set.
func CmdConfigMachineType(c *cli.Context) error {
	if c.Bool("help") {
		return cli.ShowSubcommandHelp(c)
	}
	if c.NArg() != 0 {
		fmt.Printf("expected no arguments. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}
	return CmdConfigMachineTypeShow(c)
}

// CmdConfigMachineTypeSet implements `config machine set` command.
func CmdConfigMachineTypeSet(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf("expected 1 argument. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}
	err = cmdConfigMachineTypeSet(m, c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdConfigMachineTypeSet checks a given machine type is available and then
// sets it to the defaule machine type for the current project.
func cmdConfigMachineTypeSet(m *Metadata, machineType string) (err error) {

	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	types, err := resource.MachineTypes(m.Context)
	if err != nil {
		return
	}

	var exist bool
	for _, v := range types {
		if v.Name == machineType {
			exist = true
			break
		}
	}
	if !exist {
		return fmt.Errorf("Given machine type %q is not available", machineType)
	}

	old := resource.GetMachineType()
	resource.SetMachineType(machineType)
	err = m.Config.Save()
	if err != nil {
		return fmt.Errorf("cannot save the configuration to %q: %v", m.Config.FileName, err)
	}

	if old == "" {
		fmt.Fprintf(m.Stdout, "Set machine type:\n  %v\n", chalk.Green.Color(machineType))
	} else {
		fmt.Fprintf(m.Stdout, "Update machine type: \n  %v -> %v\n", old, chalk.Green.Color(machineType))
	}
	return

}

// CmdConfigMachineTypeList implements `config machine list` command.
func CmdConfigMachineTypeList(c *cli.Context) (err error) {

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}
	err = cmdConfigMachineTypeList(m)
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdConfigMAchineTypeList prints available machine types for the current project.
func cmdConfigMachineTypeList(m *Metadata) (err error) {

	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	types, err := resource.MachineTypes(m.Context)
	if err != nil {
		return
	}

	table := uitable.New()
	table.AddRow("MACHINE TYPE", "DESCRIPTION")
	for _, v := range types {
		if v.Name == resource.GetMachineType() {
			table.AddRow(chalk.Green.Color(v.Name)+"*", chalk.Green.Color(v.Description))
		} else {
			table.AddRow(chalk.White.Color(v.Name), v.Description)
		}
	}
	fmt.Fprintln(m.Stdout, table.String())
	return

}

// CmdConfigMachineTypeShow implements `config machine set` command.
func CmdConfigMachineTypeShow(c *cli.Context) (err error) {

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}
	err = cmdConfigMachineTypeShow(m)
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdConfigMachineTypeShow prints the machine type currently chosen.
func cmdConfigMachineTypeShow(m *Metadata) (err error) {

	resource, err := m.ResourceManager()
	if err != nil {
		return
	}
	if t := resource.GetMachineType(); t != "" {
		fmt.Fprintln(m.Stdout, t)
	} else {
		fmt.Fprintln(m.Stdout, chalk.Red.Color(MsgNotSet))
	}
	return

}

// CmdConfigRegion shows current configuration of zone,
// or show help message when either -h or --help flag is set.
func CmdConfigRegion(c *cli.Context) error {
	if c.Bool("help") {
		return cli.ShowSubcommandHelp(c)
	}
	if c.NArg() != 0 {
		fmt.Printf("expected no arguments. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}
	return CmdConfigRegionShow(c)
}

// CmdConfigRegionSet implements `config region set` command.
func CmdConfigRegionSet(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf("expected 1 argument. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}
	err = cmdConfigRegionSet(m, c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdConfigRegionSet sets a given region as the default region of the current
// project. It also checkes the given region is available or not.
func cmdConfigRegionSet(m *Metadata, region string) (err error) {

	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	regions, err := resource.Regions(m.Context)
	if err != nil {
		return
	}

	var exist bool
	for _, v := range regions {
		if v.Name == region {
			exist = true
			break
		}
	}
	if !exist {
		return fmt.Errorf("Given region %q is not available", region)
	}

	old := resource.GetRegion()
	resource.SetRegion(region)

	err = m.Config.Save()
	if err != nil {
		return fmt.Errorf("cannot save the configuration to %q: %v", m.Config.FileName, err)
	}

	if old == "" {
		fmt.Fprintf(m.Stdout, "Set region:\n  %v\n", chalk.Green.Color(region))
	} else {
		fmt.Fprintf(m.Stdout, "Update region:\n  %v -> %v", old, chalk.Green.Color(region))
	}
	return

}

// CmdConfigRegionList implements `config region list` command.
func CmdConfigRegionList(c *cli.Context) (err error) {

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}
	err = cmdConfigRegionList(m)
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdConfigRegionList prints available regions.
func cmdConfigRegionList(m *Metadata) (err error) {

	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	regions, err := resource.Regions(m.Context)
	if err != nil {
		return
	}

	table := uitable.New()
	table.AddRow("REGION", "STATUS")
	for _, v := range regions {
		if v.Name == resource.GetRegion() {
			table.AddRow(chalk.Green.Color(v.Name)+"*", v.Status)
		} else {
			table.AddRow(chalk.White.Color(v.Name), v.Status)
		}
	}
	fmt.Fprintln(m.Stdout, table.String())
	return

}

// CmdConfigRegionShow implements `config region show` command.
func CmdConfigRegionShow(c *cli.Context) (err error) {

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}
	err = cmdConfigRegionShow(m)
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdConfigRegionShow prints the current selected region.
func cmdConfigRegionShow(m *Metadata) (err error) {

	resource, err := m.ResourceManager()
	if err != nil {
		return
	}
	if region := resource.GetRegion(); region != "" {
		fmt.Fprintln(m.Stdout, region)
	} else {
		fmt.Fprintln(m.Stdout, chalk.Red.Color(MsgNotSet))
	}
	return

}
