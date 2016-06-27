// +build windows

package command

import (
	"fmt"
	"os/exec"

	"github.com/deiwin/interact"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// checkGcloud checks there are gcloud command.
func checkGcloud(actor interact.Actor) error {

	if _, err := exec.LookPath("gcloud"); err != nil {
		fmt.Println(chalk.Red.Color("`Google Cloud SDK` is not found."))
		fmt.Println("Please visit https://cloud.google.com/sdk/ and install Google Cloud SDK.")
		fmt.Println("If you have installed it already, make sure your `PATH` includes `gcloud` command and reloaded it.")
		return cli.NewExitError("", 0)
	}

	return nil

}
