//
// command/source.go
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

	"github.com/jkawamoto/roadie/chalk"
	"github.com/urfave/cli"
)

// CmdSourcePut archives a given folder and uploads it as a given named file.
func CmdSourcePut(c *cli.Context) error {

	m := getMetadata(c)
	switch c.NArg() {
	case 1:
		return cmdSourcePut(m, c.Args().First(), "", c.StringSlice("exclude"))
	case 2:
		return cmdSourcePut(m, c.Args().First(), c.Args().Get(1), c.StringSlice("exclude"))
	default:
		fmt.Printf(chalk.Red.Color("expected 2 arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

}

// cmdSourcePut uploads a directory `path` after making archive file named `name`.
// If `excludes` are given, any files match such exclude patters are omitted from
// the archive file.
func cmdSourcePut(m *Metadata, path, name string, excludes []string) (err error) {

	location, err := uploadFiles(m, path, name, excludes)
	if err != nil {
		return
	}
	fmt.Println("Source files are uploaded to", chalk.Bold.TextStyle(location))
	return

}
