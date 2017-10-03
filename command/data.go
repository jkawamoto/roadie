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
	"net/url"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
	"github.com/urfave/cli"
)

// optDataPut defines options for data put command.
type optDataPut struct {
	*Metadata
	// Filename is a path for an actual file or a glob pattern.
	Filename string
	// StoredName is a path where the file to be stored.
	// If Filename is a glob pattern, matched files will be uploaded in a
	// directory named by this argument.
	StoredName string
}

// run defines process of data put command.
func (o *optDataPut) run() (err error) {

	filenames, err := filepath.Glob(o.Filename)
	if err != nil {
		return
	} else if len(filenames) == 0 {
		return fmt.Errorf("no files matching %q", o.Filename)
	}

	service, err := o.StorageManager()
	if err != nil {
		return
	}
	storage := cloud.NewStorage(service, o.Spinner.Writer)

	wg, ctx := errgroup.WithContext(o.Context)
	semaphore := make(chan struct{}, runtime.NumCPU()-1)
	var outputLock sync.Mutex
	var outputs []string
	for _, target := range filenames {

		select {
		case <-ctx.Done():
			return ctx.Err()

		case semaphore <- struct{}{}:
			func(target string) {
				wg.Go(func() (err error) {
					defer func() { <-semaphore }()

					var output string
					if o.StoredName == "" {
						output = filepath.Base(target)
					} else if len(filenames) == 1 {
						output = o.StoredName
					} else {
						output = path.Join(o.StoredName, filepath.Base(target))
					}

					var loc *url.URL
					loc, err = createURL(script.DataPrefix, output)
					if err != nil {
						return
					}
					err = storage.UploadFile(ctx, loc, target)
					if err != nil {
						return
					}

					outputLock.Lock()
					defer outputLock.Unlock()
					outputs = append(outputs, fmt.Sprintf("%v -> %v", target, loc))
					return

				})
			}(target)

		}
	}

	err = wg.Wait()
	if err != nil {
		return
	}
	for _, v := range outputs {
		fmt.Fprintln(o.Stdout, v)
	}
	return

}

// CmdDataPut uploads a given file.
func CmdDataPut(c *cli.Context) (err error) {

	// Checn the number of arguments.
	n := c.NArg()
	if n < 1 || n > 2 {
		fmt.Printf("expected 1 or 2 arguments. (%d given)\n", n)
		return cli.ShowSubcommandHelp(c)
	}

	// Get metadata.
	m, err := getMetadata(c)
	if err != nil {
		return
	}

	// Set up options.
	opt := &optDataPut{
		Metadata:   m,
		Filename:   c.Args().First(),
		StoredName: filepath.ToSlash(c.Args().Get(1)),
	}

	// Execute the command.
	if err := opt.run(); err != nil {
		err = cli.NewExitError(err.Error(), 2)
	}
	return

}
