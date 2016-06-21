package command

import (
	"log"

	"github.com/jkawamoto/roadie-cli/util"
	"github.com/urfave/cli"
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
	if v := c.String("name"); v != "" {
		s.instanceName = v
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
	if s.body.Result == "" {
		if conf.GCP.Bucket == "" {
			return cli.NewExitError("Bucket name is required or you need to add result section to "+yamlFile, 2)
		}
		s.setResult(conf.GCP.Bucket)
	}

	// debug:
	log.Printf("Script to be run is\n%s", s.String())

	// Run
	startup, err := util.Asset("assets/startup.sh")
	if err != nil {
		log.Fatal("Startup script was not found.")
		return cli.NewExitError(err.Error(), 1)
	}

	builder, err := util.NewInstanceBuilder(conf.GCP.Project)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	log.Printf("Creating an instance named %s.", s.instanceName)
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
