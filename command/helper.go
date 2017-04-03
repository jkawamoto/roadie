//
// command/helper.go
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
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/cloud/gce"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/jkawamoto/roadie/config"
	"github.com/urfave/cli"
)

// PrintTimeFormat defines time format to be used to print logs.
const PrintTimeFormat = "2006/01/02 15:04:05"

// GenerateListAction generates an action which prints list of files satisfies a given prefix.
// If url is true, show urls, too.
func GenerateListAction(prefix string) func(*cli.Context) error {

	return func(c *cli.Context) error {

		if c.NArg() != 0 {
			fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		err := PrintFileList(util.GetContext(c), prefix, c.Bool("url"), c.Bool("quiet"))
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

		ctx := util.GetContext(c)
		cfg, err := config.FromContext(ctx)
		if err != nil {
			return err
		}

		service, err := gce.NewStorageService(ctx, &cfg.GcpConfig)
		if err != nil {
			return err
		}
		defer service.Close()

		storage := cloud.NewStorage(service, nil)
		if err := storage.DownloadFiles(ctx, prefix, c.String("o"), c.Args()); err != nil {
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

		ctx := util.GetContext(c)
		cfg, err := config.FromContext(ctx)
		if err != nil {
			return err
		}

		service, err := gce.NewStorageService(ctx, &cfg.GcpConfig)
		if err != nil {
			return err
		}
		defer service.Close()

		storage := cloud.NewStorage(service, nil)
		if err := storage.DeleteFiles(ctx, prefix, c.Args()); err != nil {
			return cli.NewExitError(err.Error(), 2)
		}

		return nil

	}

}
