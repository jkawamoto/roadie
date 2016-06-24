package command

import (
	"bytes"
	"fmt"
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
	if conf.Gcp.Bucket == "" {
		fmt.Printf(chalk.Red.Color("Bucket name is not given. Use %s\n."), conf.Gcp.Project)
		conf.Gcp.Bucket = conf.Gcp.Project
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
		if err := s.setLocalSource(path, conf.Gcp.Project, conf.Gcp.Bucket); err != nil {
			return cli.NewExitError(err.Error(), 2)
		}
	} else if s.body.Source == "" {
		return cli.NewExitError("No source section and source flages are given.", 2)
	}

	// Check result section.
	if s.body.Result == "" || c.Bool("overwrite-result-section") {
		s.setResult(conf.Gcp.Bucket)
	} else {
		fmt.Printf(
			chalk.Red.Color("Since result section is given in %s, all outputs will be stored in %s.\n"), yamlFile, s.body.Result)
		fmt.Println(
			chalk.Red.Color("Those buckets might not be retrieved from this program and manually downloading results is required."))
		fmt.Println(
			chalk.Red.Color("To manage outputs by this program, delete result section or set --overwrite-result-section flag."))
	}

	// debug:
	log.Printf("Script to be run:\n%s\n", s.String())

	// Prepare startup script.
	startup, err := util.Asset("assets/startup.sh")
	if err != nil {
		log.Fatal("Startup script was not found.")
		return cli.NewExitError(err.Error(), 1)
	}

	options := " "
	if c.Bool("no-shoutdown") {
		options = "--no-shutdown"
	}

	buf := &bytes.Buffer{}
	data := map[string]string{
		"Name":    s.instanceName,
		"Script":  s.String(),
		"Options": options,
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

	disksize := c.Int64("disk-size")
	if disksize < 9 {
		disksize = 9
	}

	if c.Bool("dry") {
		log.Printf("Startup script:\n%s\n", buf.String())
	} else {
		log.Printf("Creating an instance named %s.", chalk.Bold.TextStyle(s.instanceName))
		builder.CreateInstance(s.instanceName, []*util.MetadataItem{
			&util.MetadataItem{
				Key:   "startup-script",
				Value: buf.String(),
			},
		}, disksize)
	}

	return nil
}
