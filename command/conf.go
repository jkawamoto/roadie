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
