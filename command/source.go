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
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/jkawamoto/roadie/script"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// CmdSourcePut archives a given folder and uploads it as a given named file.
// This command takes two arguments:
// - filepath: path for a file or a directory
// - name: uploading archive name
// Argument name can be omitted.
// If name is omitted, the name of file or directory where filepath points will
// be used as the name argument.
// Note that source files are archived by tar-gz method and actual uploaded file
// will have `.tar.gz` suffix.
//
// The source put command also takes --exclude flag.
// If the flag is given, any files mathing the excluding patters will be omitted
// from the source archive file.
func CmdSourcePut(c *cli.Context) (err error) {

	n := c.NArg()
	if n < 1 || n > 2 {
		fmt.Printf("expected 1 or 2 arguments. (%d given)\n", n)
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	err = cmdSourcePut(m, c.Args().First(), c.Args().Get(1), c.StringSlice("exclude"))
	if err != nil {
		err = cli.NewExitError(err, 2)
	}
	return

}

// cmdSourcePut uploads a directory `path` after making archive file named `name`.
// If `excludes` are given, any files match such exclude patters are omitted from
// the archive file.
func cmdSourcePut(m *Metadata, path, name string, excludes []string) (err error) {

	loc, err := uploadSourceFiles(m, path, name, excludes)
	if err != nil {
		return
	}
	fmt.Fprintln(m.Stdout, "Source files are uploaded to", chalk.Bold.TextStyle(loc.String()))
	return

}

// uploadSourceFiles uploads a file or a directory specified by a given path and
// stores them with a given name.
// If the name argument is not given, stored file will have the same name as
// the base name of the given path.
// If the given path represents a directory, files in the directory will be
// compressed and tarballed. In this case, the stored file will have a suffix
// `.tar.gz`.
func uploadSourceFiles(m *Metadata, fpath, name string, excludes []string) (location *url.URL, err error) {

	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		return
	}
	defer os.RemoveAll(tmp)

	if name == "" {
		var abs string
		if abs, err = filepath.Abs(fpath); err != nil {
			return
		}
		name = filepath.Base(abs)
	}
	name = fmt.Sprintf("%v.tar.gz", strings.TrimSuffix(name, ".tar.gz"))

	// File path to be uploaded.
	uploadingFile := filepath.Join(tmp, name)

	m.Spinner.Prefix = fmt.Sprint("Creating archived file", uploadingFile)
	m.Spinner.FinalMSG = fmt.Sprint("Finished creating archived file", uploadingFile)
	m.Spinner.Start()

	if err = util.Archive(fpath, uploadingFile, excludes); err != nil {
		m.Spinner.Stop()
		return
	}
	m.Spinner.Stop()

	service, err := m.StorageManager()
	if err != nil {
		return
	}
	storage := cloud.NewStorage(service, nil)

	loc, err := createURL(script.SourcePrefix, name)
	if err != nil {
		return
	}
	err = storage.UploadFile(m.Context, loc, uploadingFile)
	return loc, err

}
