package command

import (
	"fmt"

	"github.com/jkawamoto/roadie-cli/config"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// GetConfig returns a config object from a context.
func GetConfig(c *cli.Context) *config.Config {

	conf, _ := c.App.Metadata["config"].(*config.Config)

	if v := c.GlobalString("project"); v != "" {
		fmt.Printf("Overwrite project configuration: %s -> %s\n", conf.Gcp.Project, chalk.Green.Color(v))
		conf.Gcp.Project = v
	}
	if v := c.GlobalString("type"); v != "" {
		fmt.Printf("Overwrite machine type configuration: %s -> %s\n", conf.Gcp.MachineType, chalk.Green.Color(v))
		conf.Gcp.MachineType = v
	}
	if v := c.GlobalString("zone"); v != "" {
		fmt.Printf("Overwrite zone configuration: %s -> %s\n", conf.Gcp.Zone, chalk.Green.Color(v))
		conf.Gcp.Zone = v
	}
	if v := c.GlobalString("bucket"); v != "" {
		fmt.Printf("Overwrite bucket configuration: %s -> %s\n", conf.Gcp.Bucket, chalk.Green.Color(v))
		conf.Gcp.Bucket = v
	}

	return conf

}

// GenerateListAction generates an action which prints list of files satisfies a given prefix.
// If url is true, show urls, too.
func GenerateListAction(prefix string) func(*cli.Context) error {

	return func(c *cli.Context) error {

		if c.NArg() != 0 {
			fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		conf := GetConfig(c)
		return PrintFileList(conf.Gcp.Project, conf.Gcp.Bucket, prefix, c.Bool("url"), c.Bool("quiet"))

	}

}

// GenerateGetAction generates an action which downloads files from a given prefix.
func GenerateGetAction(prefix string) func(*cli.Context) error {

	return func(c *cli.Context) error {

		if c.NArg() == 0 {
			fmt.Printf(chalk.Red.Color("expected at least 1 argument. (%d given)\n"), c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		conf := GetConfig(c)
		return DownloadFiles(conf.Gcp.Project, conf.Gcp.Bucket, prefix, c.String("o"), c.Args())

	}

}

// GenerateDeleteAction generates an action which deletes files from a given prefix.
func GenerateDeleteAction(prefix string) func(*cli.Context) error {

	return func(c *cli.Context) error {

		if c.NArg() == 0 {
			fmt.Printf(chalk.Red.Color("expected at least 1 argument. (%d given)\n"), c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		conf := GetConfig(c)
		return DeleteFiles(conf.Gcp.Project, conf.Gcp.Bucket, prefix, c.Args())

	}

}
