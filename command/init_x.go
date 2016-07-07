// +build !windows
//
// command/init_x.go
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
	"fmt"
	"os/exec"

	"github.com/deiwin/interact"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// checkGcloud checks there are gcloud command.
// If not, it can install that command.
func checkGcloud(actor interact.Actor) error {

	var err error
	if _, err = exec.LookPath("gcloud"); err != nil {
		var ans bool
		fmt.Println(chalk.Red.Color("`Google Cloud SDK` is not found."))
		fmt.Println("If you have installed it already, make sure your `PATH` includes `gcloud` command and reloaded it.")

		ans, err = actor.Confirm(chalk.Yellow.Color("Install `Google Cloud SDK`?"), interact.ConfirmDefaultToYes)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		} else if !ans {
			return cli.NewExitError(chalk.Red.Color("Please install it by yourself. See https://cloud.google.com/sdk/"), -1)
		}

		// Check basic requirements.
		fmt.Println("Checking requirements...")
		if _, err = exec.LookPath("python"); err != nil {
			return cli.NewExitError(chalk.Red.Color("`python` is not found in PATH. It is required to install Google Cloud SDK."), -1)
		} else if _, err = exec.LookPath("curl"); err != nil {
			return cli.NewExitError(chalk.Red.Color("`curl` is not found in PATH. It is required to install Google Cloud SDK."), -1)
		}

		fmt.Println("Installing `Google Cloud SDK`...")
		if err = installSDK(); err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		return cli.NewExitError(chalk.Yellow.Color(`Please restart your shell and continue initialization by typing the following commands:

  $ exec -l $SHELL
  $ roadie init
      `), 0)
	}

	return nil

}
