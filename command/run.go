package command

import (
	"bytes"
	"log"
	"text/template"

	"github.com/jkawamoto/roadie-cli/util"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// CmdRun specifies the behavior of `run` command.
func CmdRun(c *cli.Context) error {

	if c.NArg() == 0 {
		return cli.NewExitError("No configuration file is given", 1)
	}

	yamlFile := c.Args()[0]

	conf := GetConfig(c)
	if conf.Gcp.Project == "" {
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
		if conf.Gcp.Bucket == "" {
			return cli.NewExitError("Bucket name is required when you use --local", 2)
		}
		if err := s.setLocalSource(path, conf.Gcp.Project, conf.Gcp.Bucket); err != nil {
			return cli.NewExitError(err.Error(), 2)
		}
	}

	// Check result section.
	if s.body.Result == "" {
		if conf.Gcp.Bucket == "" {
			return cli.NewExitError("Bucket name is required or you need to add result section to "+yamlFile, 2)
		}
		s.setResult(conf.Gcp.Bucket)
	}

	// debug:
	log.Printf("Script to be run is\n%s", s.String())

	// Prepare startup script.
	startup, err := util.Asset("assets/startup.sh")
	if err != nil {
		log.Fatal("Startup script was not found.")
		return cli.NewExitError(err.Error(), 1)
	}

	buf := &bytes.Buffer{}
	data := map[string]string{
		"Name":   s.instanceName,
		"Script": s.String(),
		// "Options": "--no-shutdown",
		"Options": " ",
	}
	temp, err := template.New("startup").Parse(string(startup))
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	if err := temp.ExecuteTemplate(buf, "startup", data); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	// Create an instance.
	builder, err := util.NewInstanceBuilder(conf.Gcp.Project)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	log.Printf("Creating an instance named %s.", chalk.Bold.TextStyle(s.instanceName))
	builder.CreateInstance(s.instanceName, []*util.MetadataItem{
		&util.MetadataItem{
			Key:   "startup-script",
			Value: buf.String(),
		},
	})

	return nil
}
