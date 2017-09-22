//
// command/result.go
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
	"strings"

	"github.com/deiwin/interact"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

const (
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

// CmdResultList shows a list of instance names or result files belonging to an
// instance.
// This command takes one optional argument, instance name.
// If an instance name is given, this command shows result files associated with
// the given instance.
// If any instance names are not given, this command shows names of instances of
// which results are stored in a cloud storage.
func CmdResultList(c *cli.Context) (err error) {

	n := c.NArg()
	if n > 1 {
		fmt.Printf("expected at most 1 argument. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}
	err = cmdResultList(m, c.Args().First(), c.Bool("url"), c.Bool("quiet"))
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdResultList implements result list command.
// If showURL is true, the URLs of each file will be shown.
// If quiet is true, only file names will be shown.
func cmdResultList(m *Metadata, instance string, showURL, quiet bool) error {

	if instance == "" {
		return PrintDirList(m, script.ResultPrefix, "", showURL, quiet)
	}
	if !strings.HasSuffix(instance, "/") {
		instance += "/"
	}
	return PrintFileList(m, script.ResultPrefix, instance, showURL, quiet)

}

// CmdResultShow shows results of a given named instance.
// This command takes two arguments, instance name and file name; and prints
// the body of the given named file belonging to the given named instance.
// File name argument can be omitted and then all files which have prefix
// `stdout` will be printed.
func CmdResultShow(c *cli.Context) (err error) {

	n := c.NArg()
	if n < 1 || n > 2 {
		fmt.Printf("expected 1 or 2 arguments. (%d given)\n", n)
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	err = cmdResultShow(m, c.Args().First(), c.Args().Get(1))
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return nil

}

// cmdResultShow implements result show command.
func cmdResultShow(m *Metadata, instance, prefix string) error {

	service, err := m.StorageManager()
	if err != nil {
		return err
	}
	storage := cloud.NewStorage(service, m.Stdout)

	if !strings.HasSuffix(instance, "/") {
		instance += "/"
	}
	loc, err := createURL(script.ResultPrefix, instance)
	if err != nil {
		return err
	}

	var header bool
	if prefix == "" {
		prefix = StdoutFilePrefix
		header = true
	}
	return storage.PrintFileBody(m.Context, loc, prefix, m.Stdout, header)

}

// CmdResultGet downloads results files which belong to a given named instance
// and matches one of given queries.
// If no queries are given, the wild card * will be used instead.
func CmdResultGet(c *cli.Context) error {

	if c.NArg() == 0 {
		fmt.Println("expected at least 1 argument. (0 given)")
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	err = cmdResultGet(m, c.Args().First(), c.Args().Tail(), c.String("o"))
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return nil

}

// cmdResultGet implements result get command.
func cmdResultGet(m *Metadata, instance string, queries []string, dir string) (err error) {

	service, err := m.StorageManager()
	if err != nil {
		return
	}

	storage := cloud.NewStorage(service, m.Stdout)
	if len(queries) == 0 {
		queries = []string{"*"}
	}

	if !strings.HasSuffix(instance, "/") {
		instance += "/"
	}
	loc, err := createURL(script.ResultPrefix, instance)
	if err != nil {
		return
	}
	return storage.DownloadFiles(m.Context, loc, dir, queries)

}

// CmdResultDelete deletes result files which belong to a given instance and
// matches given queries.
func CmdResultDelete(c *cli.Context) (err error) {

	if c.NArg() == 0 {
		fmt.Println("expected at least 1 argument. (0 given)")
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	if c.NArg() == 1 {

		var deleteAll bool
		actor := interact.NewActor(os.Stdin, m.Stdout)
		deleteAll, err = actor.Confirm(
			chalk.Red.Color("File names are not given. Do you want to delete all files?"),
			interact.ConfirmDefaultToNo)

		if err != nil {
			return cli.NewExitError(err, -1)

		} else if deleteAll {

			var logManager cloud.LogManager
			logManager, err = m.LogManager()
			if err != nil {
				return
			}
			err = logManager.Delete(m.Context, c.Args().First())
			if err != nil {
				return cli.NewExitError(err, 2)
			}

		} else {
			return nil
		}

	}

	err = cmdResultDelete(m, c.Args().First(), c.Args().Tail())
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdResultDelete implements resutl delete command.
func cmdResultDelete(m *Metadata, instance string, queries []string) (err error) {

	service, err := m.StorageManager()
	if err != nil {
		return
	}

	storage := cloud.NewStorage(service, m.Stdout)
	if len(queries) == 0 {
		queries = []string{"*"}
	}

	if !strings.HasSuffix(instance, "/") {
		instance += "/"
	}
	loc, err := createURL(script.ResultPrefix, instance)
	if err != nil {
		return
	}
	return storage.DeleteFiles(m.Context, loc, queries)

}
