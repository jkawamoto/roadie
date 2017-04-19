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
	"os"
	"path/filepath"
	"strings"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/jkawamoto/roadie/script"
	"github.com/urfave/cli"
)

// CmdSourcePut archives a given folder and uploads it as a given named file.
func CmdSourcePut(c *cli.Context) error {

	if c.NArg() != 2 {
		fmt.Printf(chalk.Red.Color("expected 2 arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m := getMetadata(c)
	if err := cmdSourcePut(m, c.Args()[0], c.Args()[1], c.StringSlice("exclude")); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// cmdSourcePut uploads a directory `root` after making archive file named `name`.
// If `excludes` are given, any files match such exclude patters are omitted from
// the archive file.
func cmdSourcePut(m *Metadata, root, name string, excludes []string) (err error) {

	filename := fmt.Sprintf("%s.tar.gz", filepath.Base(name))
	uploadingPath := filepath.Join(os.TempDir(), filename)

	m.Spinner.Prefix = fmt.Sprintf("Creating %s...", filename)
	m.Spinner.FinalMSG = fmt.Sprintf("\n%s\rDone.\n", strings.Repeat(" ", len(m.Spinner.Prefix)+2))

	m.Spinner.Start()
	if err = util.Archive(root, uploadingPath, append(excludes, uploadingPath)); err != nil {
		m.Spinner.FinalMSG = fmt.Sprintf(chalk.Red.Color("\n%s\rCannot create %s.\n"), strings.Repeat(" ", len(m.Spinner.Prefix)+2), filename)
		m.Spinner.Stop()
		return
	}
	m.Spinner.Stop()
	defer os.Remove(uploadingPath)

	service, err := m.StorageManager()
	if err != nil {
		return err
	}
	storage := cloud.NewStorage(service, nil)

	url, err := storage.UploadFile(m.Context, script.SourcePrefix, filename, uploadingPath)
	if err != nil {
		return
	}
	fmt.Printf("Source files are uploaded to %s\n", chalk.Bold.TextStyle(url))
	return nil

}
