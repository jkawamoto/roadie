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

	"golang.org/x/net/context"

	"github.com/briandowns/spinner"
	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/jkawamoto/roadie/config"
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

	conf := GetConfig(c)
	if err := cmdSourcePut(conf, c.Args()[0], c.Args()[1], c.StringSlice("exclude")); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// cmdSourcePut uploads a directory `root` after making archive file named `name`.
// If `excludes` are given, any files match such exclude patters are omitted from
// the archive file.
func cmdSourcePut(conf *config.Config, root, name string, excludes []string) (err error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = config.NewContext(ctx, conf)

	filename := fmt.Sprintf("%s.tar.gz", filepath.Base(name))
	uploadingPath := filepath.Join(os.TempDir(), filename)

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Creating %s...", filename)
	s.FinalMSG = fmt.Sprintf("\n%s\rDone.\n", strings.Repeat(" ", len(s.Prefix)+2))

	s.Start()
	if err = util.Archive(root, uploadingPath, append(excludes, uploadingPath)); err != nil {
		s.FinalMSG = fmt.Sprintf(chalk.Red.Color("\n%s\rCannot create %s.\n"), strings.Repeat(" ", len(s.Prefix)+2), filename)
		s.Stop()
		return
	}
	s.Stop()
	defer os.Remove(uploadingPath)

	storage := util.NewStorage(ctx)
	url, err := storage.UploadFile(SourcePrefix, filename, uploadingPath)
	if err != nil {
		return
	}
	fmt.Printf("Source files are uploaded to %s\n", chalk.Bold.TextStyle(url))
	return nil

}
