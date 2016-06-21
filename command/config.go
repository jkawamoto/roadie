package command

import (
	"fmt"
	"strings"

	"github.com/jkawamoto/roadie-cli/util"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// CmdConfigProject shows or sets project name to config file.
func CmdConfigProject(c *cli.Context) error {

	conf := GetConfig(c)

	switch c.NArg() {
	case 0:
		if conf.Gcp.Project != "" {
			fmt.Println(conf.Gcp.Project)
		} else {
			fmt.Println(chalk.Red.Color("Not set"))
		}
	case 1:
		var name string
		name = c.Args()[0]
		if strings.Contains(name, " ") {
			fmt.Println(chalk.Red.Color("The given project name has spaces. They are replaced to '_'."))
			name = strings.Replace(name, " ", "_", -1)
		}
		if conf.Gcp.Project == "" {
			fmt.Printf("Set project name:\n  %s\n", chalk.Green.Color(name))
		} else {
			fmt.Printf("Update project name:\n  %s -> %s\n", conf.Gcp.Project, chalk.Green.Color(name))
		}
		conf.Gcp.Project = name
	default:
		fmt.Printf(
			chalk.Red.Color("'config project' expected at most 1 argument. (%d given)\n"), c.NArg())
		// fmt.Println("  roadie config project: shows project name of Goole Cloud Platform.")
		// fmt.Println("  roadie config project <new name>: sets the given name to the project name.")
		cli.ShowSubcommandHelp(c)
	}

	return nil

}

// CmdConfigType shows current configuration of machine type,
// or show help message when either -h or --help flag is set.
func CmdConfigType(c *cli.Context) error {
	if c.Bool("help") {
		return cli.ShowSubcommandHelp(c)
	}
	return CmdConfigTypeShow(c)
}

// CmdConfigTypeSet sets a new machine type.
func CmdConfigTypeSet(c *cli.Context) error {
	// TODO: With -i, --interactive option, set such value interactively.
	if c.NArg() != 1 {
		fmt.Printf(
			chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	v := c.Args()[0]
	if conf.Gcp.MachineType == "" {
		fmt.Printf("Set machine type:\n  %s\n", chalk.Green.Color(v))
	} else {
		fmt.Printf("Update machine type:\n  %s -> %s\n", conf.Gcp.MachineType, chalk.Green.Color(v))
	}

	list, err := getAvailableTypeList(conf.Gcp.Project)
	if err == nil {
		available := false
		for _, item := range list {
			if v == item {
				available = true
			}
		}
		if !available {
			fmt.Printf(chalk.Red.Color("Updated but the given machine type '%s' is not available.\n"), v)
		}
	} else {
		fmt.Printf(chalk.Red.Color("Since project name is not given, cannot check the given machine type '%s' is available.\n"), v)
	}

	conf.Gcp.MachineType = v
	return nil
}

// CmdConfigTypeList lists up available machine types for the current project.
func CmdConfigTypeList(c *cli.Context) error {

	conf := GetConfig(c)
	if conf.Gcp.Project == "" {
		return cli.NewExitError("Project name is required to receive available machine types.", 2)
	}

	list, err := getAvailableTypeList(conf.Gcp.Project)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Println("Available machine types:")
	for _, v := range list {
		if v == conf.Gcp.MachineType {
			fmt.Println("* " + chalk.Green.Color(v))
		} else {
			fmt.Println("  " + v)
		}
	}
	return nil

}

// CmdConfigTypeShow shows current configuration of machine type.
func CmdConfigTypeShow(c *cli.Context) error {
	conf := GetConfig(c)
	if conf.Gcp.MachineType != "" {
		fmt.Println(conf.Gcp.MachineType)
	} else {
		fmt.Println(chalk.Red.Color("Not set") + " - 'n1-standard-1' will be used by default.")
	}
	return nil
}

// getAvailableTypeList retunrs a list of machine types for a given project.
func getAvailableTypeList(project string) (res []string, err error) {

	var b *util.InstanceBuilder
	res = nil

	b, err = util.NewInstanceBuilder(project)
	if err != nil {
		return
	}
	res, err = b.AvailableMachineTypes()
	return

}
