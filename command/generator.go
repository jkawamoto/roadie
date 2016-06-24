package command

import (
	"fmt"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

func GenerateListAction(prefix string) func(*cli.Context) error {

	return func(c *cli.Context) error {

		if c.NArg() != 0 {
			fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		conf := GetConfig(c)
		return PrintFileList(conf.Gcp.Project, conf.Gcp.Bucket, prefix, c.Bool("quiet"))

	}

}

func GenerateGetAction(prefix string) func(*cli.Context) error {

	return func(c *cli.Context) error {

		if c.NArg() != 1 {
			fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		conf := GetConfig(c)
		return DownloadFromGCS(conf.Gcp.Project, conf.Gcp.Bucket, prefix, c.Args()[0], c.String("o"))

	}

}

func GenerateDeleteAction(prefix string) func(*cli.Context) error {

	return func(c *cli.Context) error {

		if c.NArg() == 0 {
			fmt.Printf(chalk.Red.Color("expected at least 1 argument. (%d given)\n"), c.NArg())
			return cli.ShowSubcommandHelp(c)
		}

		conf := GetConfig(c)
		return DeleteFromGCS(conf.Gcp.Project, conf.Gcp.Bucket, prefix, c.Args())

	}

}
