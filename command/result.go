//
// command/result.go
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

	"github.com/deiwin/interact"
	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/urfave/cli"
)

const (
	// ResultPrefix defines a prefix to store result files.
	ResultPrefix = ".roadie/result"
	// StdoutFilePrefix defines a prefix for stdout result files.
	StdoutFilePrefix = "stdout"
)

// CmdResult defines the default behavior which is showing help with --help or
// -h. Otherwise, do as list command.
func CmdResult(c *cli.Context) error {

	if c.Bool("help") {
		return cli.ShowSubcommandHelp(c)
	}
	return CmdResultList(c)

}

// CmdResultList shows a list of instance names or result files belonging to an instance.
func CmdResultList(c *cli.Context) error {

	conf := GetConfig(c)
	var err error
	switch c.NArg() {
	case 0:
		err = PrintDirList(conf.Gcp.Project, conf.Gcp.Bucket, ResultPrefix, c.Bool("url"), c.Bool("quiet"))
	case 1:
		instance := c.Args().First()
		err = PrintFileList(conf.Gcp.Project, conf.Gcp.Bucket, filepath.Join(ResultPrefix, instance), c.Bool("url"), c.Bool("quiet"))
	default:
		fmt.Printf(chalk.Red.Color("expected at most 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// CmdResultShow shows results of stdout for a given instance names or result files belonging to an instance.
func CmdResultShow(c *cli.Context) error {

	conf := GetConfig(c)
	var err error
	switch c.NArg() {
	case 1:
		instance := c.Args().First()
		err = util.PrintFileBody(conf.Gcp.Project, conf.Gcp.Bucket, filepath.Join(ResultPrefix, instance), StdoutFilePrefix, false)

	case 2:
		instance := c.Args().First()
		filePrefix := StdoutFilePrefix + c.Args().Get(1)
		err = util.PrintFileBody(conf.Gcp.Project, conf.Gcp.Bucket, filepath.Join(ResultPrefix, instance), filePrefix, true)

	default:
		fmt.Printf(chalk.Red.Color("expected 1 or 2 arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// CmdResultGet downloads results for a given instance and file names are matched to queries.
func CmdResultGet(c *cli.Context) error {

	if c.NArg() < 2 {
		fmt.Printf(chalk.Red.Color("expected at least 2 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	instance := c.Args().First()

	err := util.DownloadFiles(
		conf.Gcp.Project, conf.Gcp.Bucket, filepath.Join(ResultPrefix, instance),
		c.String("o"), c.Args().Tail())

	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// CmdResultDelete deletes results for a given instance and file names are matched to queries.
func CmdResultDelete(c *cli.Context) error {

	if c.NArg() == 0 {
		fmt.Printf(chalk.Red.Color("expected at least 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	instance := c.Args().First()
	var patterns []string
	if c.NArg() == 1 {

		actor := interact.NewActor(os.Stdin, os.Stdout)
		deleteAll, err := actor.Confirm(
			chalk.Red.Color("File names are not given. Do you want to delete all files?"),
			interact.ConfirmDefaultToNo)

		if err != nil {
			return cli.NewExitError(err.Error(), -1)
		} else if deleteAll {
			patterns = []string{"*"}
		} else {
			return nil
		}

	} else {
		patterns = c.Args().Tail()
	}

	err := util.DeleteFiles(
		conf.Gcp.Project, conf.Gcp.Bucket, filepath.Join(ResultPrefix, instance), patterns)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}
