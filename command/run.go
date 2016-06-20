package command

import (
	"fmt"

	"github.com/jkawamoto/roadie-cli/util"
	"github.com/urfave/cli"
)

const (
	source = "source"
	result = "result"
)

// CmdRun specifies the behavior of `run` command.
func CmdRun(c *cli.Context) error {

	if c.NArg() == 0 {
		return cli.NewExitError("No configuration file is given", 1)
	}

	yamlFile := c.Args()[0]

	conf := GetConfig(c)
	if conf.GCP.Project == "" {
		return cli.NewExitError("Project must be given", 2)
	}

	s, err := loadScript(yamlFile, c.StringSlice("e"))
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	// Prepare source section.
	if v := c.String("git"); v != "" {
		s.setGitSource(v)
	} else if v := c.String("url"); v != "" {
		s.setURLSource(v)
	} else if path := c.String("local"); path != "" {

		if conf.GCP.Bucket == "" {
			return cli.NewExitError("Bucket name is required when you use --local", 2)
		}
		if err := s.setLocalSource(path, conf.GCP.Project, conf.GCP.Bucket); err != nil {
			return cli.NewExitError(err.Error(), 2)
		}
	}

	// Check result section.
	if _, ok := s.body["result"]; !ok {
		if conf.GCP.Bucket == "" {
			return cli.NewExitError("Bucket name is required or you need to add result section to "+yamlFile, 2)
		}
		s.setResult(conf.GCP.Bucket)
	}

	// debug:
	fmt.Println(s.String())

	// Run
	startup, err := util.Asset("assets/startup.sh")
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	builder, err := util.NewInstanceBuilder(conf.GCP.Project)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	builder.CreateInstance(s.instanceName, []*util.MetadataItem{
		&util.MetadataItem{
			Key:   "startup-script",
			Value: string(startup),
		},
		&util.MetadataItem{
			Key:   "script",
			Value: s.String(),
		},
	})

	return nil
}
