//
// commands.go
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

package main

import (
	"fmt"
	"os"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/command"
	"github.com/jkawamoto/roadie/script"
	"github.com/urfave/cli"
)

// GlobalFlags manages global flags.
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Usage: "specify a config file `NAME`.",
	},
	cli.BoolFlag{
		Name:  "verbose",
		Usage: "verbose outputs.",
	},
	// cli.BoolFlag{
	// 	Name: "no-color",
	// 	Usage: "disable colorized output."
	// },
	cli.BoolFlag{
		Name:  "auth",
		Usage: "force running an authentication process even if already logged in",
	},
}

// Commands manage sub commands.
var Commands = []cli.Command{
	{
		Name:  "init",
		Usage: "initialize roadie.",
		Description: "Check requirements. Install and set up `Google Cloud SDK` if necessary. " +
			"Create configuration file `.roadie`.",
		Category:  "Configuration",
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
		Category:  "Execution",
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
			cli.StringSliceFlag{
				Name:  "exclude",
				Usage: "`path` to be excluded to upload as the source files. This flag can be set multiply but only works with --local.",
			},
			cli.StringFlag{
				Name:  "source",
				Usage: "use `FILE` in source, shown in `roadie source list`, as source codes.",
			},
			cli.StringFlag{
				Name:  "name, n",
				Usage: "new instance uses the given `NAME`.",
			},
			cli.StringSliceFlag{
				Name:  "e",
				Usage: "`VALUE` must be key=value form which will be set in place holders of the script. This flag can be set multiply.",
			},
			cli.BoolFlag{
				Name:  "overwrite-result-section",
				Usage: "if set, result section in a given script will be overwritten to default value.",
			},
			cli.StringFlag{
				Name:  "image",
				Usage: "customize the base image which given program will run on.",
			},
			// cli.BoolFlag{
			// 	Name:  "follow, f",
			// 	Usage: "after creating instance, keep watching logs.",
			// },
			// cli.Int64Flag{
			// 	Name:  "retry",
			// 	Usage: "retry the program a given times when GCP's error happens.",
			// 	Value: 10,
			// },
		},
	},
	{
		Name:  "status",
		Usage: "show instance status.",
		Description: "Show status of instances. Stopped instances will be deleted from the output after certain time. " +
			"Without `--all` flag, this command shows status of instances of which results are not deleted.",
		Category:  "Execution",
		ArgsUsage: " ",
		Action:    command.CmdStatus,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "help, h",
				Usage: "show help",
			},
			cli.BoolFlag{
				Name:  "all",
				Usage: "show all instance status.",
			},
		},
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
			"To see messages written stdout from script, use 'result' command. This command required project ID is required. " +
			"To set project ID, use 'config project set' command. To find instance names, use 'status' command.",
		Category:  "Execution",
		ArgsUsage: "<instance name>",
		Action:    command.CmdLog,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "no-timestamp",
				Usage: "not print time stamps.",
			},
			cli.BoolFlag{
				Name:  "follow, f",
				Usage: "keep waiting new logs coming.",
			},
		},
	},
	{
		Name:        "result",
		Usage:       "list up and get results.",
		Description: "list up, show, and download computation results.",
		Category:    "Data handling",
		Action:      command.CmdResult,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "help, h",
				Usage: "show help",
			},
		},
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
				Description: "download result files from a given instance and matching given file names. " +
					"File names accept wild-card characters. " +
					"Downloaded file will be stored in the current working directory. " +
					"If '-o' option is given, downloaded file will be stored in that directory.\n\n" +
					chalk.Bold.TextStyle("Note that") + " your shell may expand wild-cards in unexpected way. " +
					"To avoid this problem, quote each file name.",
				ArgsUsage: "<instance name> <file name>...",
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
				Description: `delete result files from a given instance and match given file names.
File names accept wild card characters.　If file names are not given, delete all
files belonging to the instance.`,
				ArgsUsage: "<instance name> [<file name>...]",
				Action:    command.CmdResultDelete,
			},
		},
	},
	{
		Name:  "config",
		Usage: "show and update configuration.",
		Description: "Show and update configurations. Every configurations are stored to 'roadie.yml' in the current working directory. " +
			"You can also update configurations without this command by editing that file.",
		Category: "Configuration",
		Subcommands: cli.Commands{
			cli.Command{
				Name:      "project",
				Usage:     "show and update project ID.",
				ArgsUsage: "[<project ID>]",
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
						Usage:       "set project ID",
						Description: "Set a new name to the current project. project ID should start with alphabet and not have spaces.",
						ArgsUsage:   "<project ID>",
						Action:      command.CmdConfigProjectSet,
					},
					{
						Name:        "show",
						Usage:       "show the current project ID.",
						Description: "Show the current project ID.",
						ArgsUsage:   " ",
						Action:      command.CmdConfigProjectShow,
					},
				},
			},
			{
				Name:   "machine",
				Usage:  "show and update machine type used to run scripts.",
				Action: command.CmdConfigMachineType,
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
						Action:      command.CmdConfigMachineTypeSet,
					},
					{
						Name:  "list",
						Usage: "show available machine types.",
						Description: "Show a list of available machine types for the current project. " +
							"To receive available machine types, project ID must be set. See 'roadie config project'. " +
							"This command takes no arguments.",
						ArgsUsage: " ",
						Action:    command.CmdConfigMachineTypeList,
					},
					{
						Name:  "show",
						Usage: "show current machine type.",
						Description: "Show current machine type. If it is not set, show default machine type. " +
							"This command takes no arguments.",
						ArgsUsage: " ",
						Action:    command.CmdConfigMachineTypeShow,
					},
				},
			},
			{
				Name:   "region",
				Usage:  "show and update region information used to run scripts.",
				Action: command.CmdConfigRegion,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "help, h",
						Usage: "show help",
					},
				},
				Subcommands: cli.Commands{
					{
						Name:        "set",
						Usage:       "set a region where scripts run.",
						Description: "Set a region. Available regions are shown in 'list' command.",
						ArgsUsage:   "<zone>",
						Action:      command.CmdConfigRegionSet,
					},
					{
						Name:        "list",
						Usage:       "show available regions.",
						Description: "Show a list of regions for the current project. ",
						ArgsUsage:   " ",
						Action:      command.CmdConfigRegionList,
					},
					{
						Name:        "show",
						Usage:       "show current zone.",
						Description: "Show current region.",
						ArgsUsage:   " ",
						Action:      command.CmdConfigRegionShow,
					},
				},
			},
		},
	},
	{
		Name:  "source",
		Usage: "manage source files uploaded by this command.",
		Description: "If running scripts with --local flag, source files are uploaded to Google Cloud Storage. " +
			"This command lists up those scripts and delete them if necessary.",
		Category: "Data handling",
		Action:   command.GenerateListAction(script.SourcePrefix),
		Subcommands: cli.Commands{
			{
				Name:  "list",
				Usage: "list up source files.",
				Description: "List up source files in Google Cloud Storage. Those files can be reused for other scripts. " +
					"To reuse them, use URL like 'gs://<bucket name>/.roadie/source/<filename>'. " +
					"Otherwise, those files are not used automatically. To reduce storage size, use delete command.",
				ArgsUsage: " ",
				Action:    command.GenerateListAction(script.SourcePrefix),
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
				Description: "delete source files which match given file names. " +
					"File names accept wild card characters. ",
				ArgsUsage: "<file name>...",
				Action:    command.GenerateDeleteAction(script.SourcePrefix),
			},
			{
				Name:  "get",
				Usage: "get source files.",
				Description: "download source files which match given file names. " +
					"File names accept wild card characters. " +
					"Downloaded file will be stored in the current working directory. " +
					"If '-o' option is given, downloaded file will be stored in that directory.\n\n" +
					chalk.Bold.TextStyle("Note that") + " your shell may expand wild cards in unexpected way. " +
					"To avoid this problem, quote each file name.",
				ArgsUsage: "<file name>...",
				Action:    command.GenerateGetAction(script.SourcePrefix),
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "o",
						Usage: "output directory. Files will be stored in `DIRECTORY`. If not exists, it will be made.",
						Value: ".",
					},
				},
			},
			{
				Name:  "put",
				Usage: "put source files.",
				Description: "upload a given path as source code to be run. If the given \n" +
					"path points a directory, files in the directory are tarballed; in this case \n" +
					"uploaded file name is the directory name followd by `.tar.gz`. \n" +
					"If name option is given, uploaded file is renamed to the given name.",
				ArgsUsage: "<filepath> [<name>]",
				Action:    command.CmdSourcePut,
				Flags: []cli.Flag{
					cli.StringSliceFlag{
						Name:  "exclude, e",
						Usage: "specify excluding `PATH`. This flag can be set multiply.",
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
			"Use data section in your script to load data files in your instance.",
		Category: "Data handling",
		Action:   command.GenerateListAction(script.DataPrefix),
		Subcommands: cli.Commands{
			{
				Name:        "list",
				Usage:       "show lists of data.",
				Description: "List up data files. This command does not take any arguments.",
				ArgsUsage:   " ",
				Action:      command.GenerateListAction(script.DataPrefix),
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
					"Otherwise, basename of original file will be used. " +
					"File path accepts wild card characters, but if the given file path matches more than 2, " +
					"stored name will be ignored.",
				ArgsUsage: "<file path> [<stored name>]",
				Action:    command.CmdDataPut,
			},
			{
				Name:  "delete",
				Usage: "delete data files.",
				Description: "delete data files which match given file names. " +
					"File names accept wild card characters. ",
				ArgsUsage: "<file name>...",
				Action:    command.GenerateDeleteAction(script.DataPrefix),
			},
			{
				Name:  "get",
				Usage: "get data files.",
				Description: "download data files which match given file names. " +
					"File names accept wild card characters. " +
					"Downloaded file will be stored in the current working directory. " +
					"If '-o' option is given, downloaded file will be stored in that directory.\n\n" +
					chalk.Bold.TextStyle("Note that") + " your shell may expand wild cards in unexpected way. " +
					"To avoid this problem, quote each file name.",
				ArgsUsage: "<file name>...",
				Action:    command.GenerateGetAction(script.DataPrefix),
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
		Name:        "queue",
		Usage:       "manage queues and enqueued jobs.",
		Description: "",
		Category:    "Queue based execution",
		Action:      command.CmdQueueStatus,
		Subcommands: cli.Commands{
			{
				Name:  "add",
				Usage: "add a new task to a queue.",
				Description: "add a new task to a queue. If the specified queue does not\n" +
					"exist, a new queue and one worker instance are created for the task.",
				ArgsUsage: "<queue name> <script file>",
				Action:    command.CmdQueueAdd,
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
					cli.StringSliceFlag{
						Name:  "exclude",
						Usage: "`path` to be excluded to upload as the source files. This flag can be set multiply but only works with --local.",
					},
					cli.StringFlag{
						Name:  "source",
						Usage: "use `FILE` in source, shown in `roadie source list`, as source codes.",
					},
					cli.StringFlag{
						Name:  "name, n",
						Usage: "new instance uses the given `NAME`.",
					},
					cli.StringSliceFlag{
						Name:  "e",
						Usage: "`VALUE` must be key=value form which will be set in place holders of the script. This flag can be set multiply.",
					},
					cli.BoolFlag{
						Name:  "overwrite-result-section",
						Usage: "if set, result section in a given script will be overwritten to default value.",
					},
				},
			},
			{
				Name:        "status",
				Usage:       "show status of queues or tasks.",
				Description: "show status of queues if no quene names given; otherwise show status of tasks in the given queue.",
				ArgsUsage:   "[queue name]",
				Action:      command.CmdQueueStatus,
			},
			{
				Name:        "log",
				Usage:       "show log of queues or tasks.",
				Description: "show all log in a queue if only quene name is given; otherwise show log of a specific task.",
				ArgsUsage:   "<queue name> [task name]",
				Action:      command.CmdQueueLog,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "no-timestamp",
						Usage: "not print time stamps.",
					},
				},
			},
			{
				Name:        "instance",
				Usage:       "manage instances associated with a queue.",
				Description: "manage instances associated with a queue.",
				Subcommands: cli.Commands{
					{
						Name:        "list",
						Usage:       "list up instances working for a queue.",
						Description: "list up instances working for a queue.",
						ArgsUsage:   "<queue name>",
						Action:      command.CmdQueueInstanceList,
					},
					{
						Name:        "add",
						Usage:       "add a new instance for a queue.",
						Description: "add a new instance for a queue. Zone and instance type can be set by global flags",
						ArgsUsage:   "<queue name>",
						Action:      command.CmdQueueInstanceAdd,
						Flags: []cli.Flag{
							cli.IntFlag{
								Name:  "instances",
								Usage: "`number` of instance to be created.",
								Value: 1,
							},
						},
					},
				},
			},
			{
				Name:        "stop",
				Usage:       "stop a executing queue.",
				Description: "To reduce the number of instances working with a queue, this command helps.",
				ArgsUsage:   "<queue name>",
				Action:      command.CmdQueueStop,
			},
			{
				Name:        "restart",
				Usage:       "restart a stopping queue.",
				Description: "restart a stopping queue. By default, one instance will be created to handle the queue.",
				ArgsUsage:   "<queue name>",
				Action:      command.CmdQueueRestart,
			},
			{
				Name:        "delete",
				Usage:       "delete a queue or a task",
				Description: "if both a queue name and a task name are given, delete the tasks; otherwise delete the queue.",
				ArgsUsage:   "<queue name> [<task name>]",
				Action:      command.CmdQueueDelete,
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
