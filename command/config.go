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

// import (
// 	"fmt"
// 	"strings"
//
// 	"github.com/gosuri/uitable"
// 	"github.com/jkawamoto/roadie/cloud/gce"
// 	"github.com/jkawamoto/roadie/command/util"
// 	"github.com/jkawamoto/roadie/config"
// 	"github.com/ttacon/chalk"
// 	"github.com/urfave/cli"
// )
//
// // CmdConfigProject shows or sets project ID to config file.
// func CmdConfigProject(c *cli.Context) error {
// 	if c.Bool("help") {
// 		return cli.ShowSubcommandHelp(c)
// 	}
// 	if c.NArg() != 0 {
// 		fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
// 		return cli.ShowSubcommandHelp(c)
// 	}
// 	return CmdConfigProjectShow(c)
// }
//
// // CmdConfigProjectSet sets a given name to the current project ID.
// func CmdConfigProjectSet(c *cli.Context) error {
//
// 	if c.NArg() != 1 {
// 		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
// 		return cli.ShowSubcommandHelp(c)
// 	}
//
// 	conf := config.FromCliContext(c)
// 	var name string
// 	name = c.Args()[0]
// 	if strings.Contains(name, " ") {
// 		fmt.Println(chalk.Red.Color("The given project ID has spaces. They are replaced to '_'."))
// 		name = strings.Replace(name, " ", "_", -1)
// 	}
// 	if conf.GcpConfig.Project == "" {
// 		fmt.Printf("Set project ID:\n  %s\n", chalk.Green.Color(name))
// 	} else {
// 		fmt.Printf("Update project ID:\n  %s -> %s\n", conf.GcpConfig.Project, chalk.Green.Color(name))
// 	}
// 	conf.GcpConfig.Project = name
//
// 	if err := conf.Save(); err != nil {
// 		return cli.NewExitError(err.Error(), 3)
// 	}
// 	return nil
// }
//
// // CmdConfigProjectShow prints current project ID.
// func CmdConfigProjectShow(c *cli.Context) error {
// 	conf := config.FromCliContext(c)
// 	if conf.GcpConfig.Project != "" {
// 		fmt.Println(conf.GcpConfig.Project)
// 	} else {
// 		fmt.Println(chalk.Red.Color("Not set"))
// 	}
// 	return nil
// }
//
// // CmdConfigType shows current configuration of machine type,
// // or show help message when either -h or --help flag is set.
// func CmdConfigType(c *cli.Context) error {
// 	if c.Bool("help") {
// 		return cli.ShowSubcommandHelp(c)
// 	}
// 	if c.NArg() != 0 {
// 		fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
// 		return cli.ShowSubcommandHelp(c)
// 	}
// 	return CmdConfigTypeShow(c)
// }
//
// // CmdConfigTypeSet sets a new machine type.
// func CmdConfigTypeSet(c *cli.Context) error {
// 	if c.NArg() != 1 {
// 		fmt.Printf(
// 			chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
// 		return cli.ShowSubcommandHelp(c)
// 	}
//
// 	conf := config.FromCliContext(c)
// 	v := c.Args()[0]
// 	if conf.GcpConfig.MachineType == "" {
// 		fmt.Printf("Set machine type:\n  %s\n", chalk.Green.Color(v))
// 	} else {
// 		fmt.Printf("Update machine type:\n  %s -> %s\n", conf.GcpConfig.MachineType, chalk.Green.Color(v))
// 	}
//
// 	s := gce.NewComputeService(&conf.GcpConfig, nil)
// 	list, err := s.AvailableMachineTypes(util.GetContext(c))
// 	if err == nil {
// 		available := false
// 		for _, item := range list {
// 			if v == item.Name {
// 				available = true
// 			}
// 		}
// 		if !available {
// 			fmt.Printf(chalk.Red.Color("Updated but the given machine type '%s' is not available.\n"), v)
// 		}
// 	} else {
// 		fmt.Printf(chalk.Red.Color("Since project ID is not given, cannot check the given machine type '%s' is available.\n"), v)
// 	}
//
// 	conf.GcpConfig.MachineType = v
// 	if err = conf.Save(); err != nil {
// 		return cli.NewExitError(err.Error(), 3)
// 	}
//
// 	return nil
// }
//
// // CmdConfigTypeList lists up available machine types for the current project.
// func CmdConfigTypeList(c *cli.Context) error {
//
// 	conf := config.FromCliContext(c)
// 	if conf.GcpConfig.Project == "" {
// 		return cli.NewExitError("project ID is required to receive available machine types.", 2)
// 	}
//
// 	s := gce.NewComputeService(&conf.GcpConfig, nil)
// 	list, err := s.AvailableMachineTypes(util.GetContext(c))
// 	if err != nil {
// 		return cli.NewExitError(err.Error(), 1)
// 	}
//
// 	fmt.Println("Available machine types:")
// 	table := uitable.New()
// 	table.AddRow("MACHINE TYPE", "DESCRIPTION")
// 	for _, v := range list {
// 		if v.Name == conf.GcpConfig.MachineType {
// 			table.AddRow(chalk.Green.Color(v.Name)+"*", chalk.Green.Color(v.Description))
// 		} else {
// 			table.AddRow(chalk.ResetColor.Color(v.Name), v.Description)
// 		}
// 	}
// 	fmt.Println(table.String())
// 	return nil
//
// }
//
// // CmdConfigTypeShow shows current configuration of machine type.
// func CmdConfigTypeShow(c *cli.Context) error {
// 	conf := config.FromCliContext(c)
// 	if conf.GcpConfig.MachineType != "" {
// 		fmt.Println(conf.GcpConfig.MachineType)
// 	} else {
// 		fmt.Println(chalk.Red.Color("Not set") + " - 'n1-standard-1' will be used by default.")
// 	}
// 	return nil
// }
//
// // CmdConfigZone shows current configuration of zone,
// // or show help message when either -h or --help flag is set.
// func CmdConfigZone(c *cli.Context) error {
// 	if c.Bool("help") {
// 		return cli.ShowSubcommandHelp(c)
// 	}
// 	if c.NArg() != 0 {
// 		fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
// 		return cli.ShowSubcommandHelp(c)
// 	}
// 	return CmdConfigZoneShow(c)
// }
//
// // CmdConfigZoneSet sets a zone.
// func CmdConfigZoneSet(c *cli.Context) error {
//
// 	if c.NArg() != 1 {
// 		fmt.Printf(
// 			chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
// 		return cli.ShowSubcommandHelp(c)
// 	}
//
// 	conf := config.FromCliContext(c)
// 	v := c.Args()[0]
// 	if conf.GcpConfig.Zone == "" {
// 		fmt.Printf("Set zone:\n  %s\n", chalk.Green.Color(v))
// 	} else {
// 		fmt.Printf("Update zone:\n  %s -> %s\n", conf.GcpConfig.Zone, chalk.Green.Color(v))
// 	}
//
// 	s := gce.NewComputeService(&conf.GcpConfig, nil)
// 	list, err := s.AvailableRegions(util.GetContext(c))
// 	if err == nil {
// 		available := false
// 		for _, item := range list {
// 			if v == item.Name {
// 				available = true
// 			}
// 		}
// 		if !available {
// 			fmt.Printf(chalk.Red.Color("Updated but the given zone '%s' is not available.\n"), v)
// 		}
// 	} else {
// 		fmt.Printf(chalk.Red.Color("Since project ID is not given, cannot check the given zone '%s' is available.\n"), v)
// 	}
//
// 	conf.GcpConfig.Zone = v
// 	if err = conf.Save(); err != nil {
// 		return cli.NewExitError(err.Error(), 2)
// 	}
// 	return nil
//
// }
//
// // CmdConfigZoneList lists up available zones for the current project.
// func CmdConfigZoneList(c *cli.Context) error {
//
// 	conf := config.FromCliContext(c)
// 	if conf.GcpConfig.Project == "" {
// 		return cli.NewExitError("project ID is required to receive available zones.", 2)
// 	}
//
// 	s := gce.NewComputeService(&conf.GcpConfig, nil)
// 	list, err := s.AvailableRegions(util.GetContext(c))
// 	if err != nil {
// 		return cli.NewExitError(err.Error(), 1)
// 	}
//
// 	fmt.Println("Available zones:")
// 	table := uitable.New()
// 	table.AddRow(chalk.ResetColor.Color("ZONE"), "STATUS")
// 	for _, v := range list {
// 		if v.Name == conf.GcpConfig.Zone {
// 			table.AddRow(chalk.Green.Color(v.Name)+"*", v.Status)
// 		} else {
// 			table.AddRow(chalk.ResetColor.Color(v.Name), v.Status)
// 		}
// 	}
// 	fmt.Println(table.String())
// 	return nil
//
// }
//
// // CmdConfigZoneShow shows current configuration of zone.
// func CmdConfigZoneShow(c *cli.Context) error {
// 	conf := config.FromCliContext(c)
// 	if conf.GcpConfig.Zone != "" {
// 		fmt.Println(conf.GcpConfig.Zone)
// 	} else {
// 		fmt.Println(chalk.Red.Color("Not set") + " - 'us-central1-b' will be used by default.")
// 	}
// 	return nil
// }
//
// // CmdConfigBucket shows current configuration of bucket name,
// // or show help message when either -h or --help flag is set.
// func CmdConfigBucket(c *cli.Context) error {
// 	if c.Bool("help") {
// 		return cli.ShowSubcommandHelp(c)
// 	}
// 	if c.NArg() != 0 {
// 		fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
// 		return cli.ShowSubcommandHelp(c)
// 	}
// 	return CmdConfigBucketShow(c)
// }
//
// // CmdConfigBucketSet sets bucket name.
// func CmdConfigBucketSet(c *cli.Context) error {
// 	if c.NArg() != 1 {
// 		fmt.Printf(chalk.Red.Color("expected at most 1 argument. (%d given)\n"), c.NArg())
// 		return cli.ShowSubcommandHelp(c)
// 	}
//
// 	conf := config.FromCliContext(c)
// 	name := c.Args()[0]
// 	if conf.GcpConfig.Bucket == "" {
// 		fmt.Printf("Set bucket name:\n  %s\n", chalk.Green.Color(name))
// 	} else {
// 		fmt.Printf("Update bucket name:\n  %s -> %s\n", conf.GcpConfig.Bucket, chalk.Green.Color(name))
// 	}
//
// 	conf.GcpConfig.Bucket = name
// 	if err := conf.Save(); err != nil {
// 		return cli.NewExitError(err.Error(), 3)
// 	}
// 	return nil
//
// }
//
// // func CmdConfigBucketList(c *cli.Context) error {
// // 	return nil
// // }
//
// // CmdConfigBucketShow shows current bucket name.
// func CmdConfigBucketShow(c *cli.Context) error {
// 	conf := config.FromCliContext(c)
// 	if conf.GcpConfig.Bucket != "" {
// 		fmt.Println(conf.GcpConfig.Bucket)
// 	} else {
// 		fmt.Println(chalk.Red.Color("Not set"))
// 	}
// 	return nil
// }
//
// // func CmdConfigBucketShow(c *cli.Context) error {
// // 	return nil
// // }
