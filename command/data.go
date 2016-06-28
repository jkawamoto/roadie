//
// command/data.go
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

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// DataPrefix defines a prefix to store data files.
const DataPrefix = ".roadie/data"

// CmdDataPut uploads a given file.
func CmdDataPut(c *cli.Context) error {

	n := c.NArg()
	if n < 1 || n > 2 {
		fmt.Printf(chalk.Red.Color("expected 1 or 2 arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)

	filenames, err := filepath.Glob(c.Args()[0])
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	for _, target := range filenames {

		name := ""
		if n != 1 && len(filenames) == 1 {
			name = c.Args()[1]
		}

		location, err := UploadToGCS(conf.Gcp.Project, conf.Gcp.Bucket, DataPrefix, name, target)
		if err != nil {
			return err
		}

		fmt.Printf("File uploaded to %s.\n", chalk.Bold.TextStyle(location))
	}

	return nil
}
