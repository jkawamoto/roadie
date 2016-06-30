//
// command/source.go
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
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jkawamoto/roadie/util"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// SourcePrefix defines a prefix to store source files.
const SourcePrefix = ".roadie/source"

// CmdSourcePut archives a given folder and uploads it as a given named file.
func CmdSourcePut(c *cli.Context) error {

	if c.NArg() != 2 {
		fmt.Printf(chalk.Red.Color("expected 2 arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	root := c.Args()[0]
	name := fmt.Sprintf("%s.tar.gz", filepath.Base(c.Args()[1]))
	archPath := filepath.Join(os.TempDir(), name)

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Creating %s...", name)
	s.FinalMSG = fmt.Sprintf("\n%s\rDone.\n", strings.Repeat(" ", len(s.Prefix)+2))

	s.Start()
	if err := util.Archive(root, archPath, append(c.StringSlice("exclude"), archPath)); err != nil {
		s.FinalMSG = fmt.Sprintf(chalk.Red.Color("\n%s\rCannot create %s.\n"), strings.Repeat(" ", len(s.Prefix)+2), name)
		s.Stop()
		return cli.NewExitError(err.Error(), 2)
	}
	s.Stop()

	conf := GetConfig(c)
	url, err := UploadToGCS(conf.Gcp.Project, conf.Gcp.Bucket, SourcePrefix, name, archPath)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	fmt.Printf("Source files are uploaded to %s\n", chalk.Bold.TextStyle(url))
	return nil
}
