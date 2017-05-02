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

// CmdConfigProject shows or sets project ID to config file.
func CmdConfigProject(c *cli.Context) error {
	if c.Bool("help") {
		return cli.ShowSubcommandHelp(c)
	}
	if c.NArg() != 0 {
		fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}
	return CmdConfigProjectShow(c)
}

// CmdConfigProjectSet sets a given name to the current project ID.
func CmdConfigProjectSet(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return
	}
	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	name := c.Args().First()
	if strings.Contains(name, " ") {
		fmt.Println(chalk.Red.Color("The given project ID has spaces. They are replaced to '_'."))
		name = strings.Replace(name, " ", "_", -1)
	}
	if id := resource.GetProjectID(); id == "" {
		fmt.Printf("Set project ID:\n  %s\n", chalk.Green.Color(name))
	} else {
		fmt.Printf("Update project ID:\n  %s -> %s\n", id, chalk.Green.Color(name))
	}
	resource.SetProjectID(name)

	err = m.Config.Save()
	if err != nil {
		return cli.NewExitError(err.Error(), 3)
	}
	return
}

// CmdConfigProjectShow prints current project ID.
func CmdConfigProjectShow(c *cli.Context) (err error) {

	m, err := getMetadata(c)
	if err != nil {
		return
	}
	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	if id := resource.GetProjectID(); id != "" {
		fmt.Println(id)
	} else {
		fmt.Println(chalk.Red.Color("Not set"))
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
		fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}
	return CmdConfigMachineTypeShow(c)
}

// CmdConfigMachineTypeSet sets a new machine type.
func CmdConfigMachineTypeSet(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(
			chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return
	}
	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	newType := c.Args().First()

	types, err := resource.MachineTypes(m.Context)
	if err != nil {
		fmt.Println("Cannot get available machine types:", err.Error())
	} else {

		var exist bool
		for _, v := range types {
			if v.Name == newType {
				exist = true
				break
			}
		}
		if !exist {
			return fmt.Errorf("Given machine type %v is not available", newType)
		}

	}

	if t := resource.GetMachineType(); t == "" {
		fmt.Println("Set machine type:", chalk.Green.Color(newType))
	} else {
		fmt.Println("Update machine type:", t, "->", chalk.Green.Color(newType))
	}

	resource.SetMachineType(newType)
	err = m.Config.Save()
	if err != nil {
		return cli.NewExitError(err.Error(), 3)
	}

	return nil
}

// CmdConfigMachineTypeList lists up available machine types for the current project.
func CmdConfigMachineTypeList(c *cli.Context) (err error) {

	m, err := getMetadata(c)
	if err != nil {
		return
	}
	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	types, err := resource.MachineTypes(m.Context)
	if err != nil {
		return
	}

	fmt.Println("Available machine types:")
	table := uitable.New()
	table.AddRow("MACHINE TYPE", "DESCRIPTION")
	for _, v := range types {
		if v.Name == resource.GetMachineType() {
			table.AddRow(chalk.Green.Color(v.Name)+"*", chalk.Green.Color(v.Description))
		} else {
			table.AddRow(chalk.ResetColor.Color(v.Name), v.Description)
		}
	}
	fmt.Println(table.String())
	return

}

// CmdConfigMachineTypeShow shows current configuration of machine type.
func CmdConfigMachineTypeShow(c *cli.Context) (err error) {

	m, err := getMetadata(c)
	if err != nil {
		return
	}
	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	if t := resource.GetMachineType(); t != "" {
		fmt.Println(t)
	} else {
		fmt.Println(chalk.Red.Color("Not set"))
	}
	return nil
}

// CmdConfigRegion shows current configuration of zone,
// or show help message when either -h or --help flag is set.
func CmdConfigRegion(c *cli.Context) error {
	if c.Bool("help") {
		return cli.ShowSubcommandHelp(c)
	}
	if c.NArg() != 0 {
		fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}
	return CmdConfigRegionShow(c)
}

// CmdConfigRegionSet sets a zone.
func CmdConfigRegionSet(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(
			chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return
	}
	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	newRegion := c.Args().First()
	regions, err := resource.Regions(m.Context)
	if err != nil {
		fmt.Println("Cannot obtain available regions:", err.Error())
	} else {

		var exist bool
		for _, v := range regions {
			if v.Name == newRegion {
				exist = true
				break
			}
		}
		if !exist {
			return fmt.Errorf("Given region %v is not available", newRegion)
		}

	}

	if old := resource.GetRegion(); old == "" {
		fmt.Println("Set region:", chalk.Green.Color(newRegion))
	} else {
		fmt.Println("Update region:", old, "->", chalk.Green.Color(newRegion))
	}
	resource.SetRegion(newRegion)

	err = m.Config.Save()
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return

}

// CmdConfigRegionList lists up available zones for the current project.
func CmdConfigRegionList(c *cli.Context) (err error) {

	m, err := getMetadata(c)
	if err != nil {
		return
	}
	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	regions, err := resource.Regions(m.Context)
	if err != nil {
		return
	}

	fmt.Println("Available zones:")
	table := uitable.New()
	table.AddRow(chalk.ResetColor.Color("ZONE"), "STATUS")
	for _, v := range regions {
		if v.Name == resource.GetRegion() {
			table.AddRow(chalk.Green.Color(v.Name)+"*", v.Status)
		} else {
			table.AddRow(chalk.ResetColor.Color(v.Name), v.Status)
		}
	}
	fmt.Println(table.String())
	return

}

// CmdConfigRegionShow shows current configuration of zone.
func CmdConfigRegionShow(c *cli.Context) (err error) {

	m, err := getMetadata(c)
	if err != nil {
		return
	}
	resource, err := m.ResourceManager()
	if err != nil {
		return
	}

	if region := resource.GetRegion(); region != "" {
		fmt.Println(region)
	} else {
		fmt.Println(chalk.Red.Color("Not set"))
	}
	return

}
