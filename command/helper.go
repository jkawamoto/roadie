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
	"io"
	"net/url"
	"path"
	"strings"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// createURL returns a URL of which scheme is roadie:// and represents a path
// in a container.
func createURL(container, p string) (loc *url.URL, err error) {

	loc, err = url.Parse(script.RoadieSchemePrefix + path.Join(container, p))
	if p == "" || strings.HasSuffix(p, "/") {
		loc.Path = loc.Path + "/"
	}
	return

}

// GenerateListAction generates an action which prints list of files in a given
// container. If url is true, show urls, too.
func GenerateListAction(container string) func(*cli.Context) error {

	return func(c *cli.Context) error {

		if c.NArg() != 0 {
			fmt.Printf("expected no arguments. (%d given)\n", c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		m, err := getMetadata(c)
		if err != nil {
			return cli.NewExitError(err, 3)
		}
		err = PrintFileList(m, container, "", c.Bool("url"), c.Bool("quiet"))
		if err != nil {
			return cli.NewExitError(err, 2)
		}
		return nil

	}

}

// GenerateGetAction generates an action which downloads files from a given
// container.
func GenerateGetAction(container string) func(*cli.Context) error {

	return func(c *cli.Context) (err error) {

		if c.NArg() == 0 {
			fmt.Printf("expected at least 1 argument. (%d given)\n", c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		m, err := getMetadata(c)
		if err != nil {
			return cli.NewExitError(err, 3)
		}

		err = cmdGet(m, container, c.Args(), c.String("o"))
		if err != nil {
			return cli.NewExitError(err, 2)
		}
		return

	}

}

// cmdGet implements a general get command.
func cmdGet(m *Metadata, container string, queries []string, dir string) (err error) {

	service, err := m.StorageManager()
	if err != nil {
		return
	}
	storage := cloud.NewStorage(service, m.Spinner.Writer)

	loc, err := createURL(container, "")
	if err != nil {
		return
	}
	return storage.DownloadFiles(m.Context, loc, dir, queries)

}

// GenerateDeleteAction generates an action which deletes files from a given
// container.
func GenerateDeleteAction(container string) func(*cli.Context) error {

	return func(c *cli.Context) (err error) {

		if c.NArg() == 0 {
			fmt.Printf("expected at least 1 argument. (%d given)\n", c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		m, err := getMetadata(c)
		if err != nil {
			return cli.NewExitError(err, 3)
		}

		err = cmdDelete(m, container, c.Args())
		if err != nil {
			return cli.NewExitError(err, 2)
		}
		return

	}

}

// cmdDelete implements a general delete command.
func cmdDelete(m *Metadata, container string, queries []string) (err error) {

	service, err := m.StorageManager()
	if err != nil {
		return
	}
	storage := cloud.NewStorage(service, m.Spinner.Writer)

	loc, err := createURL(container, "")
	if err != nil {
		return
	}

	m.Spinner.Prefix = "Delete files..."
	m.Spinner.Start()
	defer m.Spinner.Stop()
	return storage.DeleteFiles(m.Context, loc, queries)

}

// SourceOpt defines options to update source section
type SourceOpt struct {
	// Git specifies a URL for a git repository to be used as source code.
	Git string
	// URL specifies a URL of an archive or executable file.
	URL string
	// Local specifies a local path which contains source code. All files except
	// matched the exclude pattern are archived and uploaded to a cloud storage.
	Local string
	// Exclude specifies a glob patters. Matched paths are excluded from the
	// source code archive. This option works with the local option.
	Exclude []string
	// Source specifies a file name in uploaded as source code already.
	Source string
}

// UpdateSourceSection updates source secrion of the given script according to
// the given option.
func UpdateSourceSection(m *Metadata, s *script.Script, opt *SourceOpt, storage *cloud.Storage) (err error) {

	// Check source section.
	switch {
	case opt.Git != "":
		if s.Source != "" {
			fmt.Fprintf(
				m.Stdout,
				chalk.Red.Color("The source section of the script will be overwritten to '%s' since a Git repository is given.\n"),
				opt.Git)
		}
		if err = setGitSource(s, opt.Git); err != nil {
			return
		}

	case opt.URL != "":
		if s.Source != "" {
			fmt.Fprintf(
				m.Stdout,
				chalk.Red.Color("The source section the script will be overwritten to '%s' since a repository URL is given.\n"),
				opt.URL)
		}
		s.Source = opt.URL

	case opt.Local != "":
		if err = setLocalSource(m, s, opt.Local, opt.Exclude); err != nil {
			return
		}

	case opt.Source != "":
		err = setUploadedSource(s, opt.Source)
		if err != nil {
			return
		}

	case s.Source == "":
		fmt.Fprintln(m.Stdout, chalk.Red.Color("No source section and source flags are given."))
	}

	return

}

// setGitSource sets a Git repository `repo` to source section in a given `script`.
// If overwriting source section, it prints warning, too.
func setGitSource(s *script.Script, repo string) (err error) {

	if strings.HasPrefix(repo, "git@") {
		sp := strings.SplitN(repo[len("git@"):], ":", 2)
		if len(sp) != 2 {
			return fmt.Errorf("Given git repository URL is invalid: %s", repo)
		}
		s.Source = fmt.Sprintf("https://%s/%s", sp[0], sp[1])
	} else {
		u, err := url.Parse(repo)
		if err != nil {
			return err
		}
		if !u.IsAbs() {
			u.Scheme = "https"
		}
		if !strings.HasSuffix(u.Path, ".git") {
			u.Path += ".git"
		}
		s.Source = u.String()

	}
	return

}

// setLocalSource sets a URL of a cloud storage to source section in a given `script` under a given context.
// It uploads source codes specified by `path` to GCS and set the URL pointing
// the uploaded files to the source section. If filename patters are given
// by `excludes`, files matching such patters are excluded to upload.
// To upload files to GCS, `conf` is used.
func setLocalSource(m *Metadata, s *script.Script, path string, excludes []string) (err error) {

	location, err := uploadSourceFiles(m, path, s.Name, excludes)
	if err != nil {
		return
	}
	s.Source = location.String()
	return

}

// setUploadedSource sets a URL to a `file` in source directory to a given `script`.
// Source codes will be downloaded from the URL. If overwriting the source
// section, it prints warning, too.
func setUploadedSource(s *script.Script, file string) (err error) {

	if !strings.HasSuffix(file, ".tar.gz") {
		file += ".tar.gz"
	}

	url, err := createURL(script.SourcePrefix, file)
	if err != nil {
		return
	}
	if s.Source != "" {
		fmt.Printf("Source section will be overwritten to '%s' since a filename is given.\n", url)
	}
	s.Source = url.String()
	return

}

// UpdateResultSection updates result section of the given script file.
func UpdateResultSection(s *script.Script, overwrite bool, warning io.Writer) (err error) {

	if s.Result == "" || overwrite {
		var loc *url.URL
		loc, err = createURL(script.ResultPrefix, s.Name)
		if err != nil {
			return
		}
		s.Result = loc.String()
	} else {
		fmt.Fprintf(
			warning,
			`Since result section is given, all outputs will be stored in %s.\n
Those buckets might not be retrieved from this program and manually downloading results is required.
To manage outputs by this program, delete result section or set --overwrite-result-section flag.`, s.Result)
	}
	return

}
