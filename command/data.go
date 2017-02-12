//
// command/data.go
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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package command

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	"golang.org/x/sync/errgroup"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/command/util"
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

	filename := c.Args()[0]
	storedName := ""
	if n == 2 {
		storedName = c.Args()[1]
	}

	if err := cmdDataPut(util.GetContext(c), filename, storedName); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

func cmdDataPut(ctx context.Context, filename, storedName string) (err error) {

	filenames, err := filepath.Glob(filename)
	if err != nil {
		return
	}

	wg, ctx := errgroup.WithContext(ctx)
	semaphore := make(chan struct{}, runtime.NumCPU()-1)

	storage := cloud.NewStorage(ctx)
	for _, target := range filenames {

		select {
		case <-ctx.Done():
			return ctx.Err()

		case semaphore <- struct{}{}:
			func(target string) {
				wg.Go(func() (err error) {
					defer func() { <-semaphore }()

					var output string
					if storedName != "" && len(filenames) == 1 {
						output = storedName
					} else {
						output = filepath.Base(target)
					}

					var location string
					location, err = storage.UploadFile(DataPrefix, output, target)
					if err != nil {
						return
					}
					fmt.Printf("File uploaded to %s.\n", chalk.Bold.TextStyle(location))
					return

				})

			}(target)
		}
	}

	return wg.Wait()

}
