package main

import (
	"fmt"
	"os"

	"github.com/jkawamoto/roadie-cli/command"
	"github.com/ttacon/chalk"
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
		Name:  "init",
		Usage: "initialize roadie.",
		Description: "Check requirements. Install and set up `Google Cloud SDK` if necessary. " +
			"Create configuration file `.roadie`.",
		ArgsUsage: " ",
		Action:    command.CmdInit,
	},
	{
		Name:  "run",
		Usage: "run a script on Google Cloud Platform.",
		Description: "Create an instance and run a given script on it. " +
			"`git`, `url`, `local` flags help to deploy source files to the instance. " +
			"Although source section in script file is used to specify where source files are, " +
			"those flags overwrite such configuration, " +
			"and `local` flag uploads local files so that the instance can access it. " +
			"With the `local` flag, you don't need to make zip files and upload them to somewhere. " +
			"Script file might have some variables, i.e. parameters. " +
			"`e` option replaces placeholders by given key-value pairs. " +
			"A placeholder named `name` looks like {{name}} in script. " +
			"Option `-e name=abcdefg` replaces {{name}} as abcdefg.",
		ArgsUsage: "<script file>",
		Action:    command.CmdRun,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "git",
				Usage: "git repository `URL`. Souce files will be cloned from there.",
			},
			cli.StringFlag{
				Name:  "url",
				Usage: "source files will be downloaded from `URL`.",
			},
			cli.StringFlag{
				Name:  "local",
				Usage: "upload source files from given `PATH` and use it the new instance.",
			},
			cli.StringFlag{
				Name:  "name",
				Usage: "new instance uses the given `NAME`.",
			},
			cli.StringSliceFlag{
				Name:  "e",
				Usage: "`VALUE` must be key=value form which will be set in place holders of the script. This flag can be set multiply.",
			},
			cli.BoolFlag{
				Name:  "no-shutdown",
				Usage: "not shoutdown instance automatically. To stop instance use 'status kill' command.",
			},
			cli.BoolFlag{
				Name:  "overwrite-result-section",
				Usage: "if set, result section in a given script will be overwritten to default value.",
			},
			cli.Int64Flag{
				Name:  "disk-size",
				Usage: "set disk size in GB.",
				Value: 9,
			},
			cli.BoolFlag{
				Name:  "dry",
				Usage: "not create any actual instances but pring the startup script to be run instead.",
			},
		},
	},
	{
		Name:        "status",
		Usage:       "show instance status.",
		Description: "Show status of instances. Stopped insances will be deleted from the output after certain time.",
		ArgsUsage:   " ",
		Action:      command.CmdStatus,
		Subcommands: cli.Commands{
			{
				Name:        "kill",
				Usage:       "kill an instance.",
				Description: "kill a given instance. Any outputs except messages written to stderr will not be stored.",
				ArgsUsage:   "<instance name>",
				Action:      command.CmdStatusKill,
			},
		},
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
		Name:        "result",
		Usage:       "list up and get results.",
		Description: "list up, show, and download computation results.",
		Subcommands: cli.Commands{
			{
				Name:  "list",
				Usage: "list up result files for a given instance.",
				Description: "List up instance names or result file names. " +
					"If instance name is given, show result file names belonging to the instance. " +
					"Otherwise show instance names which have result files.",
				ArgsUsage: "[<instance name>]",
				Action:    command.CmdResultList,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "quiet, q",
						Usage: "only display file names",
					},
					cli.BoolFlag{
						Name:  "url",
						Usage: "show url of each file.",
					},
				},
			},
			{
				Name:  "show",
				Usage: "show massages written in stdout.",
				Description: "print messages written in stdout." +
					"If an index is given, only messages associated with the index will be printed. " +
					"Otherwise, all messages will be printed. " +
					"The index is 0-origin and associated with the steps in running scripts. " +
					"For example, suppose your script has 4 steps in run section, " +
					"there will be messages indexed 0 to 3.",
				ArgsUsage: "<instance name> [<index>]",
				Action:    command.CmdResultShow,
			},
			{
				Name:  "get",
				Usage: "get result files.",
				Description: "download result files from a given instance and matching given filenames. " +
					"Filenames accept wildcard characters. " +
					"Downloaded file will be stored in the current working directory. " +
					"If '-o' option is given, downloaded file will be stored in that directory.\n\n" +
					chalk.Bold.TextStyle("Note that") + " your shell may expand wildcards in unexpected way. " +
					"To avoid this problem, quote each filename.",
				ArgsUsage: "<instance name> <filename>...",
				Action:    command.CmdResultGet,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "o",
						Usage: "output directory. Files will be stored in `DIRECTORY`. If not exists, it will be made.",
						Value: ".",
					},
				},
			},
			{
				Name:  "delete",
				Usage: "delete result files.",
				Description: "delete result files from a given instance and match given filenames. " +
					"Filenames accept wildcard characters. ",
				ArgsUsage: "<instance name> <filename>...",
				Action:    command.CmdResultDelete,
			},
		},
	},
	{
		Name:  "config",
		Usage: "show and upate configuration.",
		Description: "Show and update configurations. Every configurations are stored to '.roadie' in the current working directory. " +
			"You can also update configurations without this command by editing that file.",
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
					{
						Name:        "set",
						Usage:       "set project name",
						Description: "Set a new name to the current project. Project name should start with alphabet and not have spaces.",
						ArgsUsage:   "<project name>",
						Action:      command.CmdConfigProjectSet,
					},
					{
						Name:        "show",
						Usage:       "show the current project name.",
						Description: "Show the current project name.",
						ArgsUsage:   " ",
						Action:      command.CmdConfigProjectShow,
					},
				},
			},
			{
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
					{
						Name:        "set",
						Usage:       "set machine type.",
						Description: "Set a new machine type. Available machine types are shown in 'list' command.",
						ArgsUsage:   "<machine type>",
						Action:      command.CmdConfigTypeSet,
					},
					{
						Name:  "list",
						Usage: "show available machine types.",
						Description: "Show a list of available machine types for the current project. " +
							"To receive available machine types, project name must be set. See 'roadie config project'. " +
							"This command takes no arguments.",
						ArgsUsage: " ",
						Action:    command.CmdConfigTypeList,
					},
					{
						Name:  "show",
						Usage: "show current machine type.",
						Description: "Show current machine type. If it is not set, show default machine type. " +
							"This command takes no arguments.",
						ArgsUsage: " ",
						Action:    command.CmdConfigTypeShow,
					},
				},
			},
			{
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
					{
						Name:        "set",
						Usage:       "set zone where scripts run.",
						Description: "Set zone. Available zones are shown in 'list' command.",
						ArgsUsage:   "<zone>",
						Action:      command.CmdConfigZoneSet,
					},
					{
						Name:  "list",
						Usage: "show available zones.",
						Description: "Show a list of zones for the current project. " +
							"To receive available zones, project name must be set. See 'roadie config project'. " +
							"This command takes no arguments.",
						ArgsUsage: " ",
						Action:    command.CmdConfigZoneList,
					},
					{
						Name:  "show",
						Usage: "show current zone.",
						Description: "Show current zone. If it is not set, show default zone. " +
							"This command takes no arguments.",
						ArgsUsage: " ",
						Action:    command.CmdConfigZoneShow,
					},
				},
			},
			{
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
					{
						Name:        "set",
						Usage:       "set bucket used to store source codes and results.",
						Description: "Set bucket. If the bucket does not exist, it will be created, automatically.",
						// Description: "Set bucket. The given bucket must exist. Use create command to prepare new bucket." +
						// 	"list command shows buckets names associated with the current project.",
						ArgsUsage: "<bucket name>",
						Action:    command.CmdConfigBucketSet,
					},
					// {
					// 	Name:  "list",
					// 	Usage: "show available buckets.",
					// 	Description: "Show a list of buckets the current project can access. " +
					// 		"To receive the bucket names, project name must be set. " +
					// 		"This command takes no arguments.",
					// 	ArgsUsage: " ",
					// 	Action:    command.CmdConfigBucketList,
					// },
					{
						Name:        "show",
						Usage:       "show current bucket name.",
						Description: "Show current bucket name. This command takes no arguments.",
						ArgsUsage:   " ",
						Action:      command.CmdConfigBucketShow,
					},
					// {
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
	{
		Name:  "source",
		Usage: "manage source files uploaded by this command.",
		Description: "If running scripts with --local flag, source files are uploaded to Google Cloud Storage. " +
			"This commange lists up those scripts and delete them if necessary.",
		Subcommands: cli.Commands{
			{
				Name:  "list",
				Usage: "list up source files.",
				Description: "List up source files in Google Cloud Storage. Those files can be reused for other scripts. " +
					"To reuse them, use URL like 'gs://<bucket name>/.roadie/source/<filename>'. " +
					"Otherwise, those files are not used automatically. To reduce storage size, use delete command.",
				ArgsUsage: " ",
				Action:    command.GenerateListAction(command.SourcePrefix),
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "quiet, q",
						Usage: "only display file names",
					},
					cli.BoolFlag{
						Name:  "url",
						Usage: "show url of each file.",
					},
				},
			},
			{
				Name:  "delete",
				Usage: "delete source files.",
				Description: "delete source files which match given filenames. " +
					"Filenames accept wildcard characters. ",
				ArgsUsage: "<filename>...",
				Action:    command.GenerateDeleteAction(command.SourcePrefix),
			},
			{
				Name:  "get",
				Usage: "get source files.",
				Description: "download source files which match given filenames. " +
					"Filenames accept wildcard characters. " +
					"Downloaded file will be stored in the current working directory. " +
					"If '-o' option is given, downloaded file will be stored in that directory.\n\n" +
					chalk.Bold.TextStyle("Note that") + " your shell may expand wildcards in unexpected way. " +
					"To avoid this problem, quote each filename.",
				ArgsUsage: "<filename>...",
				Action:    command.GenerateGetAction(command.SourcePrefix),
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "o",
						Usage: "output directory. Files will be stored in `DIRECTORY`. If not exists, it will be made.",
						Value: ".",
					},
				},
			},
		},
	},
	{
		Name:  "data",
		Usage: "manage data files.",
		Description: "Manage data files. Data files can be loaded from instance using their url, " +
			"such url is based on 'gs://<bucket name>/.roadie/data/<filename>'. '" +
			"Use data section in your script to loda data files in your instance.",
		Subcommands: cli.Commands{
			{
				Name:        "list",
				Usage:       "show lists of data.",
				Description: "List up data files. This command does not take any arguments.",
				ArgsUsage:   " ",
				Action:      command.GenerateListAction(command.DataPrefix),
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "quiet, q",
						Usage: "only display file names",
					},
					cli.BoolFlag{
						Name:  "url",
						Usage: "show url of each file.",
					},
				},
			},
			{
				Name:  "put",
				Usage: "put a data file.",
				Description: "Upload a data file. " +
					"If stored name is given, uploaded file will be renamed and stored as the given name. " +
					"Otherwise, basename of original file will be used.",
				ArgsUsage: "<file path> [<stored name>]",
				Action:    command.CmdDataPut,
			},
			{
				Name:  "delete",
				Usage: "delete data files.",
				Description: "delete data files which match given filenames. " +
					"Filenames accept wildcard characters. ",
				ArgsUsage: "<filename>...",
				Action:    command.GenerateDeleteAction(command.DataPrefix),
			},
			{
				Name:  "get",
				Usage: "get data files.",
				Description: "download data files which match given filenames. " +
					"Filenames accept wildcard characters. " +
					"Downloaded file will be stored in the current working directory. " +
					"If '-o' option is given, downloaded file will be stored in that directory.\n\n" +
					chalk.Bold.TextStyle("Note that") + " your shell may expand wildcards in unexpected way. " +
					"To avoid this problem, quote each filename.",
				ArgsUsage: "<filename>...",
				Action:    command.GenerateGetAction(command.DataPrefix),
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "o",
						Usage: "output directory. Files will be stored in `DIRECTORY`. If not exists, it will be made.",
						Value: ".",
					},
				},
			},
		},
	},
}

// CommandNotFound shows error message and exit when a given command is not found.
func CommandNotFound(c *cli.Context, command string) {

	fmt.Fprintf(os.Stderr, chalk.Red.Color("'%s' is not a %s command..\n"), command, c.App.Name)
	cli.ShowAppHelp(c)
	os.Exit(2)

}
