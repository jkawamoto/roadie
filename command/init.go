//
// command/init.go
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
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/deiwin/interact"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// CmdInit helps to create a configuration file.
func CmdInit(c *cli.Context) error {

	var err error
	actor := interact.NewActor(os.Stdin, os.Stdout)

	fmt.Printf(`%s.
This command will create ."roadie" file in current directory. Configurations
can be updated with "roadie config" command. See "roadie config --help",
for more detail. Type ctrl-c at anytime to quite.

`, chalk.Bold.TextStyle("Initialize Roadie"))

	// TODO: Windows support.

	// Check basic requirements.
	fmt.Println("Checking requirements...")
	if _, err = exec.LookPath("python"); err != nil {
		return cli.NewExitError(chalk.Red.Color("`python` is not found in PATH. It is required."), -1)
	} else if _, err = exec.LookPath("curl"); err != nil {
		return cli.NewExitError(chalk.Red.Color("`curl` is not found in PATH. It is required."), -1)
	}

	// Check gcloud command.
	if _, err = exec.LookPath("gcloud"); err != nil {
		var ans bool
		fmt.Println(chalk.Red.Color("`Google Cloud SDK` is not foune."))
		fmt.Println("If you have installed it already, make sure your `PATH` includes `gcloud` command and reloaded it.")
		ans, err = actor.Confirm(chalk.Yellow.Color("Install `Google Cloud SDK`?"), interact.ConfirmDefaultToYes)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		} else if !ans {
			return cli.NewExitError(chalk.Red.Color("Please install it by yourself. See https://cloud.google.com/sdk/"), -1)
		}

		fmt.Println("Installing `Google Cloud SDK`...")
		if err = installSDK(); err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		fmt.Println(chalk.Yellow.Color(`Please restart your shell and continue initialization by typing the following commands:

    $ exec -l $SHELL
    $ roadie init
    `))
		return nil

	}

	// Get gcloud configuration.
	fmt.Println("Loading configurations of `Google Cloud SDK`...")
	gcloud, err := getGcloudConf()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	} else if gcloud.Project == "" {
		var ans bool
		ans, err = actor.Confirm(chalk.Yellow.Color("`Google Cloud SDK` does not seem to be set up. Setup?"), interact.ConfirmDefaultToYes)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		} else if !ans {
			return cli.NewExitError(chalk.Red.Color("Please setup it by yourself. Run `gcloud init`."), -1)
		}

		fmt.Println("Setting up `Google Cloud SDK`...")
		if err = setupGcloud(); err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		gcloud, err = getGcloudConf()
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
	}

	conf := GetConfig(c)
	conf.Gcp.Project = gcloud.Project
	conf.Gcp.Zone = gcloud.Zone

	message := "Please enter project name"
	conf.Gcp.Project, err = actor.PromptOptional(message, conf.Gcp.Project, checkNotEmpty)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	message = "Please enter bucket name"
	conf.Gcp.Bucket, err = actor.PromptOptional(message, conf.Gcp.Project, checkNotEmpty)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	fmt.Println("")

	abs, _ := filepath.Abs(".roadie")
	fmt.Printf("About to write to %s:\n", abs)
	conf.Print()
	save, err := actor.Confirm(chalk.Yellow.Color("Is this ok?"), interact.ConfirmNoDefault)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	if save {
		fmt.Println("Saving configuarions...")
		if err = conf.Save(); err != nil {
			return cli.NewExitError(err.Error(), 2)
		}
	}
	return nil
}

func checkNotEmpty(value string) error {
	if value == "" {
		return fmt.Errorf("Input value is empty.")
	}
	return nil
}

func installSDK() (err error) {

	curl := exec.Command("curl", "https://sdk.cloud.google.com")
	runner := exec.Command("bash")

	curlOut, err := curl.StdoutPipe()
	if err != nil {
		return
	}
	runnerIn, err := runner.StdinPipe()
	if err != nil {
		return
	}
	go func() {
		io.Copy(runnerIn, curlOut)
		runnerIn.Close()
	}()

	runnerOut, err := runner.StdoutPipe()
	if err != nil {
		return
	}
	go io.Copy(os.Stdout, runnerOut)

	curlErr, err := curl.StderrPipe()
	if err != nil {
		return
	}
	runnerErr, err := runner.StderrPipe()
	if err != nil {
		return
	}
	go io.Copy(os.Stderr, io.MultiReader(curlErr, runnerErr))

	curl.Start()
	runner.Run()
	curl.Wait()

	return nil
}

func getGcloudConf() (res GcloudConfig, err error) {

	output, err := exec.Command("gcloud", "config", "list").Output()
	if err != nil {
		return
	}

	res = GcloudConfig{}
	for _, v := range strings.Split(string(output), "\n") {
		if strings.Contains(v, "=") {

			kv := strings.Split(v, " = ")
			switch kv[0] {
			case "zone":
				res.Zone = kv[1]
			case "account":
				res.Account = kv[1]
			case "project":
				res.Project = kv[1]
			}
		}
	}
	return

}

type GcloudConfig struct {
	Zone    string
	Account string
	Project string
}

func setupGcloud() (err error) {

	cmd := exec.Command("gcloud", "init")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	defer stdin.Close()
	go io.Copy(stdin, os.Stdin)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	defer stdout.Close()
	go io.Copy(os.Stdout, stdout)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return
	}
	defer stderr.Close()
	go io.Copy(os.Stderr, stderr)

	return cmd.Run()

}
