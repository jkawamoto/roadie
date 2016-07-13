//
// command/helper.go
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

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/config"
	"github.com/urfave/cli"
)

// PrintTimeFormat defines time format to be used to print logs.
const PrintTimeFormat = "2006/01/02 15:04:05"

// GetConfig returns a config object from a context.
func GetConfig(c *cli.Context) *config.Config {

	conf, _ := c.App.Metadata["config"].(*config.Config)

	if v := c.GlobalString("project"); v != "" {
		fmt.Printf("Overwrite project configuration: %s -> %s\n", conf.Gcp.Project, chalk.Green.Color(v))
		conf.Gcp.Project = v
	}
	if v := c.GlobalString("type"); v != "" {
		fmt.Printf("Overwrite machine type configuration: %s -> %s\n", conf.Gcp.MachineType, chalk.Green.Color(v))
		conf.Gcp.MachineType = v
	}
	if v := c.GlobalString("zone"); v != "" {
		fmt.Printf("Overwrite zone configuration: %s -> %s\n", conf.Gcp.Zone, chalk.Green.Color(v))
		conf.Gcp.Zone = v
	}
	if v := c.GlobalString("bucket"); v != "" {
		fmt.Printf("Overwrite bucket configuration: %s -> %s\n", conf.Gcp.Bucket, chalk.Green.Color(v))
		conf.Gcp.Bucket = v
	}

	return conf

}

// GenerateListAction generates an action which prints list of files satisfies a given prefix.
// If url is true, show urls, too.
func GenerateListAction(prefix string) func(*cli.Context) error {

	return func(c *cli.Context) error {

		if c.NArg() != 0 {
			fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		conf := GetConfig(c)
		err := PrintFileList(conf.Gcp.Project, conf.Gcp.Bucket, prefix, c.Bool("url"), c.Bool("quiet"))
		if err != nil {
			return cli.NewExitError(err.Error(), 2)
		}
		return nil

	}

}

// GenerateGetAction generates an action which downloads files from a given prefix.
func GenerateGetAction(prefix string) func(*cli.Context) error {

	return func(c *cli.Context) error {

		if c.NArg() == 0 {
			fmt.Printf(chalk.Red.Color("expected at least 1 argument. (%d given)\n"), c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		conf := GetConfig(c)
		err := DownloadFiles(conf.Gcp.Project, conf.Gcp.Bucket, prefix, c.String("o"), c.Args())
		if err != nil {
			return cli.NewExitError(err.Error(), 2)
		}
		return nil

	}

}

// GenerateDeleteAction generates an action which deletes files from a given prefix.
func GenerateDeleteAction(prefix string) func(*cli.Context) error {

	return func(c *cli.Context) error {

		if c.NArg() == 0 {
			fmt.Printf(chalk.Red.Color("expected at least 1 argument. (%d given)\n"), c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		conf := GetConfig(c)
		if err := DeleteFiles(conf.Gcp.Project, conf.Gcp.Bucket, prefix, c.Args()); err != nil {
			return cli.NewExitError(err.Error(), 2)
		}
		return nil

	}

}