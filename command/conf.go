package command

import (
	"github.com/jkawamoto/roadie-cli/config"
	"github.com/urfave/cli"
)

// GetConfig returns a config object from a context.
func GetConfig(c *cli.Context) *config.Config {

	conf, _ := c.App.Metadata["config"].(*config.Config)
	if v := c.String("project"); v != "" {
		conf.Gcp.Project = v
	}
	if v := c.String("bucket"); v != "" {
		conf.Gcp.Bucket = v
	}

	return conf

}
