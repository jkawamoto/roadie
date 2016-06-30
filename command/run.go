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
	"strings"
	"text/template"
	"time"

	"github.com/briandowns/spinner"
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

	yamlFile := c.Args()[0]

	conf := GetConfig(c)
	if conf.Gcp.Project == "" {
		return cli.NewExitError("Project must be given", 2)
	}
	if conf.Gcp.Bucket == "" {
		fmt.Printf(chalk.Red.Color("Bucket name is not given. Use %s\n."), conf.Gcp.Project)
		conf.Gcp.Bucket = conf.Gcp.Project
	}

	script, err := loadScript(yamlFile, c.StringSlice("e"))
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	if v := c.String("name"); v != "" {
		script.instanceName = v
	}

	// Prepare source section.
	if v := c.String("git"); v != "" {
		script.setGitSource(v)
	} else if v := c.String("url"); v != "" {
		script.setURLSource(v)
	} else if path := c.String("local"); path != "" {
		if err := script.setLocalSource(path, conf.Gcp.Project, conf.Gcp.Bucket); err != nil {
			return cli.NewExitError(err.Error(), 2)
		}
	} else if script.body.Source == "" {
		return cli.NewExitError("No source section and source flages are given.", 2)
	}

	// Check result section.
	if script.body.Result == "" || c.Bool("overwrite-result-section") {
		script.setResult(conf.Gcp.Bucket)
	} else {
		fmt.Printf(
			chalk.Red.Color("Since result section is given in %s, all outputs will be stored in %s.\n"), yamlFile, script.body.Result)
		fmt.Println(
			chalk.Red.Color("Those buckets might not be retrieved from this program and manually downloading results is required."))
		fmt.Println(
			chalk.Red.Color("To manage outputs by this program, delete result section or set --overwrite-result-section flag."))
	}

	// debug:
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
