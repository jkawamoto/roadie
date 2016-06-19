package command

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type config map[interface{}]interface{}

const (
	source = "source"
	git    = "git"
	url    = "url"
	local  = "local"
	result = "result"
)

// CmdRun specifies the behavior of `run` command.
func CmdRun(c *cli.Context) error {

	if c.NArg() == 0 {
		return cli.NewExitError("No configuration file is given", 1)
	}

	conf := make(config)
	err := loadScript(c.Args()[0], c.StringSlice("e"), &conf)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	// Prepare source section.
	if v := c.String(git); v != "" {

		if _, ok := conf[source]; ok {
			log.Printf(
				chalk.Red.Color("%s has source section but a Git repository is given. The source section will be overwritten to '%s'."),
				c.Args()[0],
				v,
			)
		}
		conf[source] = v

	} else if v := c.String(url); v != "" {

		if _, ok := conf[source]; ok {
			log.Printf(
				chalk.Red.Color("%s has source section but a repository URL is given. The source section will be overwritten to '%s'."),
				c.Args()[0],
				v,
			)

		}
		conf[source] = v

	} else if path := c.String(local); path != "" {

		if _, ok := conf[source]; ok {
			log.Printf(
				chalk.Red.Color("%s has source section but a path for source codes is given. The source section will be overwritten."),
				c.Args()[0],
			)
		}

		if info, notExists := os.Stat(path); notExists != nil {
			// TODO: File doesn't exist.
		} else if info.IsDir() {
			// TODO: Make a tar ball and upload it to a bucket.
		} else {
			// TODO: If v is a already archived file, just upload it.
			if !strings.HasSuffix(path, ".zip") && !strings.HasSuffix(path, ".tar") && !strings.HasSuffix(path, ".tar.gz") {
				// TODO: Warning.
			}

		}

		// TODO: path must be updated to a valid gs url.
		conf[source] = path

	} else {
		// TODO: if no source flag given, what shoud it do?

	}

	// Check result section.
	// checkResultSection

	// debug:
	res, err := yaml.Marshal(conf)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	fmt.Println(string(res))

	return nil
}

// Load a given script file and apply arguments.
func loadScript(filename string, args []string, out *config) error {

	// Define function map to replace place holders.
	funcs := template.FuncMap{}
	for _, v := range args {
		sp := strings.Split(v, "=")
		if len(sp) >= 2 {
			funcs[sp[0]] = func() string {
				return sp[1]
			}
		}
	}

	// Load YAML config file.
	conf, err := template.New(filename).Funcs(funcs).ParseFiles(filename)
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			return err
		default:
			return fmt.Errorf("Cannot apply variables to the place holders in %s", filename)
		}
	}

	// Replace place holders with given args.
	buf := &bytes.Buffer{}
	if err := conf.Execute(buf, nil); err != nil {
		return err
	}

	// Unmarshal YAML file.
	return yaml.Unmarshal(buf.Bytes(), out)

}

// checkResultSection validates config has result section.
func checkResultSection(c *cli.Context, conf *config) error {

	// if _, ok := conf[result]; !ok {
	//
	// 	// if c.Bool("quiet") {
	// 	// 	return cli.NewExitError("Configuration doesn't have result section.", 3)
	// 	// } else {
	// 	//
	// 	// }
	// 	return cli.NewExitError("Configuration doesn't have result section.", 3)
	//
	// }

	return nil

}

// Return a unique file name to be used for uploading file.
func getUniqueFilename() string {

	if dir, err := os.Getwd(); err != nil {
		return "hoge"

	} else {

		return dir

	}

}
