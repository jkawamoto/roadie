//
// command/prepare.go
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
	"context"
	"fmt"

	"github.com/jkawamoto/roadie/config"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// PrepareCommand prepares executing any command; it loads the configuratio file,
// checkes global flags.
func PrepareCommand(c *cli.Context) (err error) {

	// Load the configuration file.
	conf, err := config.NewConfig()
	if err != nil {
		return
	}

	// Update configuration.
	if conf.Project == "" && c.Command.Name != "init" {
		fmt.Println(chalk.Yellow.Color("Project ID is not given. It is recommended to run `roadie init`."))
	}
	if v := c.GlobalString("project"); v != "" {
		fmt.Printf("Overwrite project configuration: %s -> %s\n", conf.Project, chalk.Green.Color(v))
		conf.Project = v
	}
	if v := c.GlobalString("type"); v != "" {
		fmt.Printf("Overwrite machine type configuration: %s -> %s\n", conf.MachineType, chalk.Green.Color(v))
		conf.MachineType = v
	}
	if v := c.GlobalString("zone"); v != "" {
		fmt.Printf("Overwrite zone configuration: %s -> %s\n", conf.Zone, chalk.Green.Color(v))
		conf.Zone = v
	}
	if v := c.GlobalString("bucket"); v != "" {
		fmt.Printf("Overwrite bucket configuration: %s -> %s\n", conf.Bucket, chalk.Green.Color(v))
		conf.Bucket = v
	}

	// Append the config to the context.
	ctx, ok := c.App.Metadata["context"].(context.Context)
	if !ok {
		return fmt.Errorf("Context is broken: %v", c.App.Metadata)
	}
	c.App.Metadata["context"] = config.NewContext(ctx, conf)

	return

}
