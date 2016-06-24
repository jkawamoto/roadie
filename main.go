package main

import (
	"os"

	"github.com/jkawamoto/roadie-cli/config"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = Author
	app.Email = Email
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Metadata = map[string]interface{}{
		"config": config.LoadConfig("./.roadie"),
	}

	app.Run(os.Args)
}
