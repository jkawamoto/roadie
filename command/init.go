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
	"strings"

	"github.com/jkawamoto/roadie/cloud/azure"
	"github.com/jkawamoto/roadie/cloud/gcp"
	"github.com/jkawamoto/roadie/config"

	"github.com/deiwin/interact"
	"github.com/urfave/cli"
)

// CmdInit helps to create a configuration file.
func CmdInit(c *cli.Context) (err error) {

	actor := interact.NewActor(os.Stdin, os.Stdout)

	// Initialization steps:
	// 1. Choose cloud service provider (GCP only).
	// 2. Ask nessesarry information (project id for gcp).
	// 3. Authentication.
	// 4. Store configurations.

	m, err := getMetadata(c)
	if err != nil {
		fmt.Println(`This command will create file "roadie.yml" in current directory.
Configurations can be updated with "roadie config" command.
See "roadie config --help", for more detail.
`)
		m.Config = new(config.Config)
		m.Config.FileName = "roadie.yml"

	} else {
		filename := m.Config.FileName
		m.Config = new(config.Config)
		m.Config.FileName = filename

	}

	fmt.Println(m.Decorator.Green("Initialize Roadie"))
	servicer, err := actor.PromptAndRetry("Choose a cloud service provider from [azure, gcp]", checkOption("azure", "gcp"))
	if err != nil {
		return
	}

	switch servicer {
	case "gcp":
		var project string
		project, err = actor.PromptAndRetry("Enter your project ID", checkNotEmpty)
		if err != nil {
			return cli.NewExitError(err.Error(), 10)
		}
		m.Config.GcpConfig.Project = project

		fmt.Println("")
		fmt.Println(m.Decorator.Green("Cheking authorization..."))
		_, err = gcp.NewProvider(m.Context, &m.Config.GcpConfig, m.Logger, true)
		if err != nil {
			return
		}

	case "azure":
		m.Config.AzureConfig.TenantID, err = actor.PromptAndRetry("Enter your tenant ID", checkNotEmpty)
		if err != nil {
			return
		}
		m.Config.AzureConfig.SubscriptionID, err = actor.PromptAndRetry("Enter your subscription ID", checkNotEmpty)
		if err != nil {
			return
		}
		m.Config.AzureConfig.ResourceGroupName, err = actor.PromptAndRetry(
			"Enter your project ID (you can choose any name but it must be unique in the world)", checkNotEmpty)
		if err != nil {
			return
		}

		fmt.Println("")
		fmt.Println(m.Decorator.Green("Cheking authorization..."))
		_, err = azure.NewProvider(m.Context, &m.Config.AzureConfig, m.Logger, true)
		if err != nil {
			return
		}

	}

	fmt.Println(m.Decorator.Green("Saving configuarions..."))
	return m.Config.Save()

}

// checkNotEmpty is a predicate used in actor to check input message is empty.
func checkNotEmpty(value string) error {
	if value == "" {
		return fmt.Errorf("Input value is empty.")
	}
	return nil
}

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
