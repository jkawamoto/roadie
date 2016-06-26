package command

import (
	"fmt"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// DataPrefix defines a prefix to store data files.
const DataPrefix = ".roadie/data"

// CmdDataPut uploads a given file.
func CmdDataPut(c *cli.Context) error {

	n := c.NArg()
	if n < 1 || n > 2 {
		fmt.Printf(chalk.Red.Color("expected 1 or 2 arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	var name string
	if n == 1 {
		name = ""
	} else {
		name = c.Args()[1]
	}
	location, err := UploadToGCS(conf.Gcp.Project, conf.Gcp.Bucket, DataPrefix, name, c.Args()[0])
	if err != nil {
		return err
	}

	fmt.Printf("File uploaded to %s.\n", chalk.Bold.TextStyle(location))
	return nil
}
