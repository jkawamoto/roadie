//
// command/init.go
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

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/cloud/gcp"
	"github.com/jkawamoto/roadie/config"

	"github.com/deiwin/interact"
	"github.com/urfave/cli"
)

// GcloudConfig defines information recived from `gcloud config list`.
type GcloudConfig struct {
	Zone    string
	Account string
	Project string
}

// CmdInit helps to create a configuration file.
func CmdInit(c *cli.Context) (err error) {

	fmt.Println(chalk.Green.Color("Initialize Roadie"))
	actor := interact.NewActor(os.Stdin, os.Stdout)

	// Initialization steps:
	// 1. Choose cloud service provider (GCP only).
	// 2. Ask nessesarry information (project id for gcp).
	// 3. Authentication.
	// 4. Store configurations.

	m, err := getMetadata(c)
	if err != nil {
		fmt.Println(`This command will create "roadie.yml" file in current directory.
Configurations can be updated with "roadie config" command.
See "roadie config --help", for more detail.
`)
		m.Config = new(config.Config)
		m.Config.FileName = "roadie.yml"
	}

	var project string
	message := "Please enter your project ID"
	project, err = actor.PromptAndRetry(message, checkNotEmpty)
	if err != nil {
		return cli.NewExitError(err.Error(), 10)
	}
	m.Config.GcpConfig.Project = project

	fmt.Println("")
	fmt.Println(chalk.Green.Color("Cheking authorization..."))
	provider, err := gcp.NewProvider(m.Context, &m.Config.GcpConfig, m.Logger, true)
	if err != nil {
		return
	}
	_ = provider

	fmt.Println(chalk.Green.Color("Saving configuarions..."))
	return m.Config.Save()

}

// checkNotEmpty is a predicate used in actor to check input message is empty.
func checkNotEmpty(value string) error {
	if value == "" {
		return fmt.Errorf("Input value is empty.")
	}
	return nil
}
