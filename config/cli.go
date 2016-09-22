//
// config/clid.go
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

package config

import (
	"fmt"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// FromCliContext returns a config object from a context of cli.
func FromCliContext(c *cli.Context) (conf *Config) {

	conf, _ = c.App.Metadata["config"].(*Config)

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

	return

}
