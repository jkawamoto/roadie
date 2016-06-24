package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jkawamoto/roadie-cli/util"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

const (
	// ResultPrefix defines a prefix to store result files.
	ResultPrefix = ".roadie/result"
	// StdoutFilePrefix defines a prefix for stdout result files.
	StdoutFilePrefix = "stdout"
)

// CmdResultList shows a list of instance names or result files belonging to an instance.
func CmdResultList(c *cli.Context) error {

	conf := GetConfig(c)
	switch c.NArg() {
	case 0:
		return PrintDirList(conf.Gcp.Project, conf.Gcp.Bucket, ResultPrefix, c.Bool("url"), c.Bool("quiet"))
	case 1:
		instance := c.Args()[0]
		return PrintFileList(conf.Gcp.Project, conf.Gcp.Bucket, filepath.Join(ResultPrefix, instance), c.Bool("url"), c.Bool("quiet"))
	default:
		fmt.Printf(chalk.Red.Color("expected at most 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

}

// CmdResultShow shows results of stdout for a given instance names or result files belonging to an instance.
func CmdResultShow(c *cli.Context) error {

	conf := GetConfig(c)
	switch c.NArg() {
	case 1:
		instance := c.Args()[0]
		return printFileBody(conf.Gcp.Project, conf.Gcp.Bucket, filepath.Join(ResultPrefix, instance), StdoutFilePrefix, false)

	case 2:
		instance := c.Args()[0]
		filePrefix := StdoutFilePrefix + c.Args()[1]
		return printFileBody(conf.Gcp.Project, conf.Gcp.Bucket, filepath.Join(ResultPrefix, instance), filePrefix, true)

	default:
		fmt.Printf(chalk.Red.Color("expected 1 or 2 arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

}

// CmdResultGet downloads results for a given instance and file names are matched to queries.
func CmdResultGet(c *cli.Context) error {

	if c.NArg() < 2 {
		fmt.Printf(chalk.Red.Color("expected at least 2 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	instance := c.Args().First()
	return DownloadFiles(
		conf.Gcp.Project, conf.Gcp.Bucket, filepath.Join(ResultPrefix, instance),
		c.String("o"), c.Args().Tail())

}

// CmdResultDelete deletes results for a given instance and file names are matched to queries.
func CmdResultDelete(c *cli.Context) error {

	if c.NArg() < 2 {
		fmt.Printf(chalk.Red.Color("expected at least 2 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	instance := c.Args().First()
	return DeleteFiles(
		conf.Gcp.Project, conf.Gcp.Bucket, filepath.Join(ResultPrefix, instance), c.Args().Tail())

}

// printFileBody prints file bodies in a bucket associated with a project,
// which has a prefix ans satisfies query. If quiet is ture, additional messages
// well be suppressed.
func printFileBody(project, bucket, prefix, query string, quiet bool) error {

	return ListupFiles(
		project, bucket, prefix,
		func(storage *util.Storage, file <-chan *util.FileInfo, done chan<- struct{}) {

			for {
				info := <-file
				if info == nil {
					done <- struct{}{}
					return
				}

				if info.Name != "" && strings.HasPrefix(info.Name, query) {
					if !quiet {
						fmt.Printf(chalk.Bold.TextStyle("*** %s ***\n"), info.Name)
					}
					if err := storage.Download(info.Path, os.Stdout); err != nil {
						fmt.Printf(chalk.Red.Color("Cannot download %s (%s)."), info.Name, err.Error())
					}
				}

			}

		})

}
