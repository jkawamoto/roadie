package main

import (
	"fmt"
	"os"

	"github.com/jkawamoto/roadie-cli/command"
	"github.com/urfave/cli"
)

// GlobalFlags manages golabal flags.
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "project, p",
		Usage: "overwrite project name configuration.",
	},
	cli.StringFlag{
		Name:  "type, t",
		Usage: "overwrite machi type configuration.",
	},
	cli.StringFlag{
		Name:  "zone, z",
		Usage: "overwrite zone configuration.",
	},
	cli.StringFlag{
		Name:  "bucket, b",
		Usage: "overwrite bucket name configuration.",
	},
}

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
				Name:  "name",
				Usage: "Instance name.",
			},
			cli.StringSliceFlag{
				Name:  "e",
				Usage: "key=value to be set in place holders of the script.",
			},
			cli.BoolFlag{
				Name:  "overwrite-result-section",
				Usage: "if set this flag, result section in a given script will be overwritten to default value.",
			},
			cli.IntFlag{
				Name:  "disk-size",
				Usage: "set disk size in GB. (Minimum: 9)",
			},
		},
	},
	{
		Name:        "status",
		Usage:       "show instance status.",
		Description: "Show status of instances. Stopped insances will be deleted from the output after certain time.",
		ArgsUsage:   " ",
		Action:      command.CmdStatus,
		// TODO: Add kill command to delete instance by hand.
	},
	{
		Name:  "log",
		Usage: "show logs.",
		Description: "Show logs for a given instance name. Logs consists of messages from the framework and messages written to stderr. " +
			"To see messages written stdout from script, use 'result' command. This command required project name is required. " +
			"To set project name, use 'config project set' command. To find instance names, use 'status' command.",
		ArgsUsage: "<instance name>",
		Action:    command.CmdLog,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "no-timestamp",
				Usage: "Not print timestamps.",
			},
		},
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
				Name:      "project",
				Usage:     "show and update project name of Google Cloud Platform.",
				ArgsUsage: "[<project name>]",
				Action:    command.CmdConfigProject,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "help, h",
						Usage: "show help",
					},
				},
				Subcommands: cli.Commands{
					cli.Command{
						Name:        "set",
						Usage:       "set project name",
						Description: "Set a new name to the current project. Project name should start with alphabet and not have spaces.",
						ArgsUsage:   "<project name>",
						Action:      command.CmdConfigProjectSet,
					},
					cli.Command{
						Name:        "show",
						Usage:       "show the current project name.",
						Description: "Show the current project name.",
						ArgsUsage:   " ",
						Action:      command.CmdConfigProjectShow,
					},
				},
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
				Name:   "zone",
				Usage:  "show and update zone used to run scripts.",
				Action: command.CmdConfigZone,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "help, h",
						Usage: "show help",
					},
				},
				Subcommands: cli.Commands{
					cli.Command{
						Name:        "set",
						Usage:       "set zone where scripts run.",
						Description: "Set zone. Available zones are shown in 'list' command.",
						ArgsUsage:   "<zone>",
						Action:      command.CmdConfigZoneSet,
					},
					cli.Command{
						Name:  "list",
						Usage: "show available zones.",
						Description: "Show a list of zones for the current project. " +
							"To receive available zones, project name must be set. See 'roadie config project'. " +
							"This command takes no arguments.",
						ArgsUsage: " ",
						Action:    command.CmdConfigZoneList,
					},
					cli.Command{
						Name:  "show",
						Usage: "show current zone.",
						Description: "Show current zone. If it is not set, show default zone. " +
							"This command takes no arguments.",
						ArgsUsage: " ",
						Action:    command.CmdConfigZoneShow,
					},
				},
			},
			cli.Command{
				Name:   "bucket",
				Usage:  "show and update bucket name.",
				Action: command.CmdConfigBucket,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "help, h",
						Usage: "show help",
					},
				},
				Subcommands: cli.Commands{
					cli.Command{
						Name:        "set",
						Usage:       "set bucket used to store source codes and results.",
						Description: "Set bucket. If the bucket does not exist, it will be created, automatically.",
						// Description: "Set bucket. The given bucket must exist. Use create command to prepare new bucket." +
						// 	"list command shows buckets names associated with the current project.",
						ArgsUsage: "<bucket name>",
						Action:    command.CmdConfigBucketSet,
					},
					// cli.Command{
					// 	Name:  "list",
					// 	Usage: "show available buckets.",
					// 	Description: "Show a list of buckets the current project can access. " +
					// 		"To receive the bucket names, project name must be set. " +
					// 		"This command takes no arguments.",
					// 	ArgsUsage: " ",
					// 	Action:    command.CmdConfigBucketList,
					// },
					cli.Command{
						Name:        "show",
						Usage:       "show current bucket name.",
						Description: "Show current bucket name. This command takes no arguments.",
						ArgsUsage:   " ",
						Action:      command.CmdConfigBucketShow,
					},
					// cli.Command{
					// 	Name:  "create",
					// 	Usage: "create a new bucket.",
					// 	Description: "Create a new bucket with a given name. " +
					// 		"To create a new bucket, project name must be set.",
					// 	ArgsUsage: "<bucket name>",
					// 	Action:    command.CmdConfigBucketShow,
					// },
				},
			},
		},
	},
}

// CommandNotFound shows error message and exit when a given command is not found.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
