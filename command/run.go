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

// CmdRun specifies the behavior of `run` command.
func CmdRun(c *cli.Context) error {

	if c.NArg() == 0 {
		fmt.Println(chalk.Red.Color("Script file is not given."))
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	if conf.Gcp.Project == "" {
		return cli.NewExitError("Project must be given", 2)
	}
	if conf.Gcp.Bucket == "" {
		fmt.Printf(chalk.Red.Color("Bucket name is not given. Use %s\n."), conf.Gcp.Project)
		conf.Gcp.Bucket = conf.Gcp.Project
	}

	script, err := loadScript(c.Args()[0], c.StringSlice("e"))
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	// Check source section.
	if script.body.Source == "" {
		return cli.NewExitError("No source section and source flages are given.", 2)
	}
	return runScript(conf, script, c)
}

// CmdRunGit sets a git repository to the source and creates an instance.
func CmdRunGit(c *cli.Context) error {

	if c.NArg() != 2 {
		fmt.Printf(chalk.Red.Color("expected 2 arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	if conf.Gcp.Project == "" {
		return cli.NewExitError("Project must be given", 2)
	}
	if conf.Gcp.Bucket == "" {
		fmt.Printf(chalk.Red.Color("Bucket name is not given. Use %s\n."), conf.Gcp.Project)
		conf.Gcp.Bucket = conf.Gcp.Project
	}

	script, err := loadScript(c.Args()[1], c.StringSlice("e"))
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	// Prepare source section.
	repo := c.Args()[0]
	if script.body.Source != "" {
		fmt.Printf(
			chalk.Red.Color("The source section of %s will be overwritten to '%s' since a Git repository is given.\n"),
			script.filename, repo)
	}
	script.body.Source = repo

	return runScript(conf, script, c)
}

// CmdRunURL sets a url to the source and creates an instance.
func CmdRunURL(c *cli.Context) error {

	if c.NArg() != 2 {
		fmt.Printf(chalk.Red.Color("expected 2 arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	if conf.Gcp.Project == "" {
		return cli.NewExitError("Project must be given", 2)
	}
	if conf.Gcp.Bucket == "" {
		fmt.Printf(chalk.Red.Color("Bucket name is not given. Use %s\n."), conf.Gcp.Project)
		conf.Gcp.Bucket = conf.Gcp.Project
	}

	script, err := loadScript(c.Args()[1], c.StringSlice("e"))
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	// Prepare source section.
	url := c.Args()[0]
	if script.body.Source != "" {
		fmt.Printf(
			chalk.Red.Color("The source section of %s will be overwritten to '%s' since a repository URL is given.\n"),
			script.filename, url)
	}
	script.body.Source = url

	return runScript(conf, script, c)

}

// CmdRunLocal uploads a directory as the source and creates an instance.
func CmdRunLocal(c *cli.Context) error {

	if c.NArg() != 2 {
		fmt.Printf(chalk.Red.Color("expected 2 arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	if conf.Gcp.Project == "" {
		return cli.NewExitError("Project must be given", 2)
	}
	if conf.Gcp.Bucket == "" {
		fmt.Printf(chalk.Red.Color("Bucket name is not given. Use %s\n."), conf.Gcp.Project)
		conf.Gcp.Bucket = conf.Gcp.Project
	}

	script, err := loadScript(c.Args()[1], c.StringSlice("e"))
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	// Prepare source section.
	if script.body.Source != "" {
		fmt.Printf(
			chalk.Red.Color("The source section of %s is overwritten since a path for source codes is given.\n"),
			script.filename)
	}

	path := c.Args()[0]
	info, err := os.Stat(path)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	var name string
	var arcPath string
	if info.IsDir() {

		// Directory will be archived.
		filename := script.instanceName + ".tar.gz"
		arcPath = filepath.Join(os.TempDir(), filename)

		spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		spin.Prefix = fmt.Sprintf("Creating an archived file %s...", arcPath)
		spin.FinalMSG = fmt.Sprintf("\n%s\rCreating the archived file %s.    \n", strings.Repeat(" ", len(spin.Prefix)+2), arcPath)
		spin.Start()
		if err := util.Archive(path, arcPath, c.StringSlice("exclude")); err != nil {
			spin.Stop()
			return cli.NewExitError(err.Error(), 2)
		}
		name = filename
		spin.Stop()

	} else {

		// One source file just will be uploaded.
		arcPath = path
		name = util.Basename(path)

	}

	location, err := UploadToGCS(conf.Gcp.Project, conf.Gcp.Bucket, SourcePrefix, name, arcPath)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	script.body.Source = location

	return runScript(conf, script, c)
}

// runScript run a given script with config and context information.
func runScript(conf *config.Config, script *Script, c *cli.Context) error {

	// Update instance name.
	if v := c.String("name"); v != "" {
		script.instanceName = v
	}

	// Check result section.
	if script.body.Result == "" || c.Bool("overwrite-result-section") {
		script.setResult(conf.Gcp.Bucket)
	} else {
		fmt.Printf(
			chalk.Red.Color("Since result section is given in %s, all outputs will be stored in %s.\n"), script.filename, script.body.Result)
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
		return cli.NewExitError(err.Error(), 1)
	}

	options := " "
	if c.Bool("no-shoutdown") {
		options = "--no-shutdown"
	}

	buf := &bytes.Buffer{}
	data := map[string]string{
		"Name":    script.instanceName,
		"Script":  script.String(),
		"Options": options,
	}
	temp, err := template.New("startup").Parse(string(startup))
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	if err := temp.ExecuteTemplate(buf, "startup", data); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	// Create an instance.
	builder, err := util.NewInstanceBuilder(conf.Gcp.Project)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	if c.Bool("dry") {
		fmt.Printf("Startup script:\n%s\n", buf.String())
	} else {

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Prefix = fmt.Sprintf("Creating an instance named %s...", chalk.Bold.TextStyle(script.instanceName))
		s.FinalMSG = fmt.Sprintf("\n%s\rInstance created.\n", strings.Repeat(" ", len(s.Prefix)+2))
		s.Start()
		defer s.Stop()

		err := builder.CreateInstance(script.instanceName, []*util.MetadataItem{
			&util.MetadataItem{
				Key:   "startup-script",
				Value: buf.String(),
			},
		}, c.Int64("disk-size"))

		if err != nil {
			s.FinalMSG = fmt.Sprintf(chalk.Red.Color("\n%s\rCannot create instance.\n"), strings.Repeat(" ", len(s.Prefix)+2))
			return cli.NewExitError(err.Error(), 2)
		}

	}
	return nil

}
