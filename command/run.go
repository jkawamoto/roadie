//
// command/run.go
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
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jkawamoto/roadie/config"
	"github.com/jkawamoto/roadie/util"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// runOpt manages all arguments and flags defined in run command.
type runOpt struct {

	// Git repository URL which will be cloned as source codes.
	Git string

	// URL where source codes are stored.
	URL string

	// Directory or file path which will be used as source codes.
	Local string

	// File patterns which will be excluded from source codes in a given
	// directory. This flag will only work when `local` flag is given with
	// a directory.
	Exclude []string

	// Path for the script file to be run.
	ScriptFile string

	// Arguments for the script.
	ScriptArgs []string

	// Instance name. If not set, named by script file name and current time.
	InstanceName string

	// Specify disk size of new instance.
	DiskSize int64

	// If true, result section will be overwritten so that roadie can manage
	// result data. Otherwise, users require to manage them by theirself.
	OverWriteResultSection bool

	// If true, created instance will not shutdown automatically. So, users
	// require to do it by theirself. This flag can be useful for debugging.
	NoShutdown bool

	// If true, do not create any instances but show startup script.
	// This flag is for debugging.
	Dry bool
}

// CmdRun specifies the behavior of `run` command.
func CmdRun(c *cli.Context) error {

	if c.NArg() == 0 {
		fmt.Println(chalk.Red.Color("Script file is not given."))
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	opt := runOpt{
		Git:                    c.String("git"),
		URL:                    c.String("url"),
		Local:                  c.String("local"),
		Exclude:                c.StringSlice("exclude"),
		ScriptFile:             c.Args()[0],
		ScriptArgs:             c.StringSlice("e"),
		InstanceName:           c.String("name"),
		DiskSize:               c.Int64("disk-size"),
		OverWriteResultSection: c.Bool("overwrite-result-section"),
		NoShutdown:             c.Bool("no-shutdown"),
		Dry:                    c.Bool("dry"),
	}
	if err := cmdRun(conf, &opt); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// cmdRun implements the main logic of run command.
func cmdRun(conf *config.Config, opt *runOpt) (err error) {

	if conf.Gcp.Project == "" {
		return fmt.Errorf("Project name must be given.")
	}
	if conf.Gcp.Bucket == "" {
		fmt.Printf(chalk.Red.Color("Bucket name is not given. Use %s\n."), conf.Gcp.Project)
		conf.Gcp.Bucket = conf.Gcp.Project
	}

	script, err := NewScript(opt.ScriptFile, opt.ScriptArgs)
	if err != nil {
		return
	}

	// Update instance name.
	if opt.InstanceName != "" {
		script.InstanceName = strings.ToLower(opt.InstanceName)
	}

	// Check source section.
	switch {
	case opt.Git != "":
		setGitSource(script, opt.Git)

	case opt.URL != "":
		setURLSource(script, opt.URL)

	case opt.Local != "":
		if err = setLocalSource(conf, script, opt.Local, opt.Exclude); err != nil {
			return
		}

	case script.Body.Source == "":
		fmt.Println(chalk.Red.Color("No source section and source flages are given."))
	}

	// Check result section.
	if script.Body.Result == "" || opt.OverWriteResultSection {
		location := util.CreateURL(conf.Gcp.Bucket, ResultPrefix, script.InstanceName)
		script.Body.Result = location.String()
	} else {
		fmt.Printf(
			chalk.Red.Color("Since result section is given in %s, all outputs will be stored in %s.\n"), script.Filename, script.Body.Result)
		fmt.Println(
			chalk.Red.Color("Those buckets might not be retrieved from this program and manually downloading results is required."))
		fmt.Println(
			chalk.Red.Color("To manage outputs by this program, delete result section or set --overwrite-result-section flag."))
	}

	// Debugging info.
	fmt.Printf("Script to be run:\n%s\n", script.String())

	// Prepare startup script.
	startup, err := util.Asset("assets/startup.sh")
	if err != nil {
		fmt.Println(chalk.Red.Color("Startup script was not found."))
		return
	}

	options := " "
	if opt.NoShutdown {
		options = "--no-shutdown"
	}

	buf := &bytes.Buffer{}
	data := map[string]string{
		"Name":    script.InstanceName,
		"Script":  script.String(),
		"Options": options,
	}
	temp, err := template.New("startup").Parse(string(startup))
	if err != nil {
		return
	}
	if err = temp.ExecuteTemplate(buf, "startup", data); err != nil {
		return
	}

	// Create an instance.
	builder, err := util.NewInstanceBuilder(conf.Gcp.Project)
	if err != nil {
		return
	}

	if opt.Dry {
		fmt.Printf("Startup script:\n%s\n", buf.String())
	} else {

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Prefix = fmt.Sprintf("Creating an instance named %s...", chalk.Bold.TextStyle(script.InstanceName))
		s.FinalMSG = fmt.Sprintf("\n%s\rInstance created.\n", strings.Repeat(" ", len(s.Prefix)+2))
		s.Start()
		defer s.Stop()

		err = builder.CreateInstance(script.InstanceName, []*util.MetadataItem{
			&util.MetadataItem{
				Key:   "startup-script",
				Value: buf.String(),
			},
		}, opt.DiskSize)

		if err != nil {
			s.FinalMSG = fmt.Sprintf(chalk.Red.Color("\n%s\rCannot create instance.\n"), strings.Repeat(" ", len(s.Prefix)+2))
			return
		}

	}
	return nil

}

// setGitSource sets a Git repository `repo` to source section in a given `script`.
// If overwriting source section, it prints warning, too.
func setGitSource(script *Script, repo string) {
	if script.Body.Source != "" {
		fmt.Printf(
			chalk.Red.Color("The source section of %s will be overwritten to '%s' since a Git repository is given.\n"),
			script.Filename, repo)
	}
	script.Body.Source = repo
}

// setURLSource sets a `url` to source section in a given `script`.
// Source codes will be downloaded from the URL.
// The file pointed by the URL must be either executable, zipped, or tarballed
// file. If overwriting source section, it prints warning, too.
func setURLSource(script *Script, url string) {
	if script.Body.Source != "" {
		fmt.Printf(
			chalk.Red.Color("The source section of %s will be overwritten to '%s' since a repository URL is given.\n"),
			script.Filename, url)
	}
	script.Body.Source = url
}

// setLocalSource sets a GCS URL to source section in a given `script`.
// It uploads source codes specified by `path` to GCS and set the URL pointing
// the uploaded files to the source section. If filename patters are given
// by `excludes`, files matching such patters are excluded to upload.
// To upload files to GCS, `conf` is used.
func setLocalSource(conf *config.Config, script *Script, path string, excludes []string) (err error) {

	info, err := os.Stat(path)
	if err != nil {
		return
	}

	var filename string      // File name on GCS.
	var uploadingPath string // File path to be uploaded.

	if info.IsDir() { // Directory will be archived.

		filename = script.InstanceName + ".tar.gz"
		uploadingPath = filepath.Join(os.TempDir(), filename)

		spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		spin.Prefix = fmt.Sprintf("Creating an archived file %s...", uploadingPath)
		spin.FinalMSG = fmt.Sprintf("\n%s\rCreating the archived file %s.    \n", strings.Repeat(" ", len(spin.Prefix)+2), uploadingPath)
		spin.Start()

		if err = util.Archive(path, uploadingPath, excludes); err != nil {
			spin.Stop()
			return
		}

		spin.Stop()

	} else { // One source file just will be uploaded.

		uploadingPath = path
		filename = util.Basename(path)

	}

	location, err := UploadToGCS(conf.Gcp.Project, conf.Gcp.Bucket, SourcePrefix, filename, uploadingPath)
	if err != nil {
		return
	}
	script.Body.Source = location
	return nil

}
