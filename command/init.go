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
	"strings"

	"github.com/jkawamoto/roadie/config"
	"github.com/ttacon/chalk"

	"github.com/deiwin/interact"
	"github.com/urfave/cli"
)

// CmdInit helps to create a configuration file.
func CmdInit(c *cli.Context) (err error) {

	// If metadata don't have complete information, it's okay for now.
	m, _ := getMetadata(c)

	err = cmdInit(m)
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdInit initialize Roadie.
// The initialization steps:
// 1. Choose cloud service provider (GCP only).
// 2. Ask nessesarry information (project id for gcp).
// 3. Authentication.
// 4. Store configurations.
func cmdInit(m *Metadata) (err error) {

	if m.Config == nil {
		m.Config = new(config.Config)
	}
	if m.Config.FileName == "" {
		m.Config.FileName = ConfigFile
	}

	fmt.Fprintf(m.Stdout, `%v
This command will create file %q in current directory.
Configurations can be updated with "roadie config" command.
See "roadie config --help", for more detail.

`, chalk.Green.Color("Initialize Roadie"), m.Config.FileName)

	actor := interact.NewActor(m.Stdin, m.Stdout)

	provider, err := actor.PromptOptionalAndRetry("Choose a cloud provider from a) Microsoft Azure, g) Google Cloud Platform", "g", checkOption("a", "g"))
	if err != nil {
		return
	}
	switch provider {
	case "a":
		var tennant, subscription string
		tennant, err = actor.PromptAndRetry("Enter your tennant ID", checkNotEmpty)
		if err != nil {
			return
		}
		fmt.Fprintln(m.Stdout, "")
		m.Config.AzureConfig.TenantID = tennant

		subscription, err = actor.PromptAndRetry("Enter your subscription ID", checkNotEmpty)
		if err != nil {
			return
		}
		fmt.Fprintln(m.Stdout, "")
		m.Config.AzureConfig.SubscriptionID = subscription

	case "g":
		var project string
		project, err = actor.PromptAndRetry("Enter your project ID", checkNotEmpty)
		if err != nil {
			return
		}
		fmt.Fprintln(m.Stdout, "")
		m.Config.GcpConfig.Project = project

	}

	fmt.Fprintln(m.Stdout, chalk.Green.Color("Cheking authorization..."))
	err = m.prepareProvider(true)
	if err != nil {
		return
	}

	fmt.Fprintln(m.Stdout, chalk.Green.Color("Saving configuarions..."))
	return m.Config.Save()

}

// checkNotEmpty is a predicate used in actor to check input message is empty.
func checkNotEmpty(value string) error {
	if value == "" {
		return fmt.Errorf("Input value is empty.")
	}
	return nil
}

// checkOption returns InputCheck checks a given string matches at least one of
// given options.
func checkOption(options ...string) interact.InputCheck {
	return func(str string) error {
		for _, v := range options {
			if v == str {
				return nil
			}
		}
		return fmt.Errorf("Input must be one of [%v]", strings.Join(options, ", "))
	}
}
