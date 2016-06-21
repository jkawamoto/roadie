package main

import (
	"fmt"
	"os"

	"github.com/jkawamoto/roadie-cli/command"
	"github.com/urfave/cli"
)

// GlobalFlags manages golabal flags.
var GlobalFlags = []cli.Flag{}

// Commands manage sub commands.
var Commands = []cli.Command{
	{
		Name:   "run",
		Usage:  "Run a script on Google Cloud Platform.",
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
		Name:  "config",
		Usage: "Show and upate configuration.",
		Description: "Show and update configurations. Every configurations are stored to '.roadie' in the current working directory. " +
			"You can also update configurations without this command by editing that file.",
		Action: cli.ShowSubcommandHelp,
		Flags:  []cli.Flag{},
		Subcommands: cli.Commands{
			cli.Command{
				Name:        "project",
				Usage:       "show and update project name of Google Cloud Platform.",
				Description: "Set the given name as the project name when <project name> is given. Otherwise show the current project name.",
				ArgsUsage:   "[<project name>]",
				Action:      command.CmdConfigProject,
			},
			cli.Command{
				Name:   "type",
				Usage:  "show and update machine type used to run scripts.",
				Action: command.CmdConfigType,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "help, h",
						Usage: "show help",
					},
				},
				Subcommands: cli.Commands{
					cli.Command{
						Name:        "set",
						Usage:       "set machine type.",
						Description: "Set a new machine type. Available machine types are shown in 'list' command.",
						ArgsUsage:   "<machine type>",
						Action:      command.CmdConfigTypeSet,
					},
					cli.Command{
						Name:  "list",
						Usage: "show available machine types.",
						Description: "Show a list of available machine types for the current project. " +
							"To receive available machine types, project name must be set. See 'roadie config project'. " +
							"This command takes no arguments.",
						ArgsUsage: " ",
						Action:    command.CmdConfigTypeList,
					},
					cli.Command{
						Name:  "show",
						Usage: "show current machine type.",
						Description: "Show current machine type. If it is not set, show default machine type. " +
							"This command takes no arguments.",
						ArgsUsage: " ",
						Action:    command.CmdConfigTypeShow,
					},
				},
			},
			cli.Command{
				Name: "zone",
			},
			cli.Command{
				Name: "bucket",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
