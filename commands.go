package main

import (
	"fmt"
	"os"

	"github.com/jkawamoto/roadie-cli/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "quiet, q",
		Usage: "If set, no ask to user.",
	},
}

var Commands = []cli.Command{
	{
		Name:   "run",
		Usage:  "",
		Action: command.CmdRun,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "git",
				Usage: "Git repository.",
			},
			cli.StringFlag{
				Name:  "url",
				Usage: "URL of the source code.",
			},
			cli.StringFlag{
				Name:  "local",
				Usage: "Local path to be run.",
			},
			cli.StringFlag{
				Name:  "project",
				Usage: "Project name.",
			},
			cli.StringFlag{
				Name:  "bucket",
				Usage: "Bucket name.",
			},
			cli.StringFlag{
				Name:  "name",
				Usage: "Instance name.",
			},
			cli.StringSliceFlag{
				Name:  "e",
				Usage: "key=value to be set in place holders of the script.",
			},
			// TODO: Add disk size option
		},
	},
	{
		Name:   "status",
		Usage:  "",
		Action: command.CmdStatus,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "result",
		Usage:  "",
		Action: command.CmdResult,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "config",
		Usage:  "",
		Action: command.CmdConfig,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
