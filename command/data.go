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
// along with Roadie.  If not, see <http://www.gnu.org/licenses/>.
//

package command

import (
	"fmt"
	"path/filepath"
	"runtime"

	"golang.org/x/sync/errgroup"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
	"github.com/urfave/cli"
)

// optDataPut defines options for data put command.
type optDataPut struct {
	*Metadata
	Filename   string
	StoredName string
}

// run defines process of data put command.
func (o *optDataPut) run() (err error) {

	if o.Filename == "" {
		return fmt.Errorf("Given filename is empty")
	}

	filenames, err := filepath.Glob(o.Filename)
	if err != nil {
		return
	}

	service, err := o.Metadata.StorageManager()
	if err != nil {
		return
	}
	storage := cloud.NewStorage(service, nil)

	wg, ctx := errgroup.WithContext(o.Context)
	semaphore := make(chan struct{}, runtime.NumCPU()-1)
	for _, target := range filenames {

		select {
		case <-ctx.Done():
			return ctx.Err()

		case semaphore <- struct{}{}:
			func(target string) {
				wg.Go(func() (err error) {
					defer func() { <-semaphore }()

					var output string
					if o.StoredName != "" && len(filenames) == 1 {
						output = o.StoredName
					} else {
						output = filepath.Base(target)
					}

					var location string
					location, err = storage.UploadFile(ctx, script.DataPrefix, output, target)
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

// CmdDataPut uploads a given file.
func CmdDataPut(c *cli.Context) error {

	n := c.NArg()
	if n < 1 || n > 2 {
		fmt.Printf(chalk.Red.Color("expected 1 or 2 arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return err
	}
	storedName := ""
	if n == 2 {
		storedName = c.Args()[1]
	}
	opt := &optDataPut{
		Metadata:   m,
		Filename:   c.Args().First(),
		StoredName: storedName,
	}

	if err := opt.run(); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}
