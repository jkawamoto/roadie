package command

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/jkawamoto/roadie-cli/util"
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

	yamlFile := c.Args()[0]

	conf := make(config)
	err := loadScript(yamlFile, c.StringSlice("e"), &conf)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	// Prepare source section.
	if v := c.String(git); v != "" {

		if _, ok := conf[source]; ok {
			log.Printf(
				chalk.Red.Color("%s has source section but a Git repository is given. The source section will be overwritten to '%s'."),
				yamlFile,
				v,
			)
		}
		conf[source] = v

	} else if v := c.String(url); v != "" {

		if _, ok := conf[source]; ok {
			log.Printf(
				chalk.Red.Color("%s has source section but a repository URL is given. The source section will be overwritten to '%s'."),
				yamlFile,
				v,
			)

		}
		conf[source] = v

	} else if path := c.String(local); path != "" {

		if _, ok := conf[source]; ok {
			log.Printf(
				chalk.Red.Color("%s has source section but a path for source codes is given. The source section will be overwritten."),
				yamlFile,
			)
		}

		if info, notExists := os.Stat(path); notExists != nil {
			// Target path does not exits.
			return cli.NewExitError(notExists.Error(), 2)

		} else if info.IsDir() {

			filename := basename(yamlFile) + time.Now().Format("20060102150405") + ".tar.gz"
			arcPath := filepath.Join(os.TempDir(), filename)
			log.Printf("Create an archived file %s", arcPath)
			util.Archive(path, arcPath, nil)

			url := "gs://" + c.String("bucket") + "/.roadie/source/" + filename
			log.Printf("Uploading to %s", url)

			// TODO: Upload the archive file to url.

			conf[source] = url

		} else {

			url := "gs://" + c.String("bucket") + "/.roadie/source/" + basename(path)
			log.Printf("Uploading to %s", url)

			// TODO: Upload the file to url.

			conf[source] = url

		}

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

// Get the basename of a given filename.
func basename(filename string) string {

	ext := filepath.Ext(filename)
	bodySize := len(filename) - len(ext)

	return filepath.Base(filename[:bodySize])

}
