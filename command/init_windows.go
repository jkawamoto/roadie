// +build windows
//
// command/init_windows.go
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
	"os/exec"

	"github.com/deiwin/interact"
	"github.com/urfave/cli"
)

// checkGcloud checks there are gcloud command.
func checkGcloud(actor interact.Actor) error {

	if _, err := exec.LookPath("gcloud"); err != nil {
		fmt.Println("`Google Cloud SDK` is not found.")
		fmt.Println("Please visit https://cloud.google.com/sdk/ and install Google Cloud SDK.")
		fmt.Println("If you have installed it already, make sure your `PATH` includes `gcloud` command and reloaded it.")
		return cli.NewExitError("", 0)
	}

	return nil

}
