package command

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

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

// ListupFilesWorker is goroutine of a woker called from listupFiles.
type ListupFilesWorker func(storage *util.Storage, file <-chan *util.FileInfo, done chan<- struct{})

// CmdResultList shows a list of instance names or result files belonging to an instance.
func CmdResultList(c *cli.Context) error {

	conf := GetConfig(c)
	switch c.NArg() {
	case 0:
		return PrintDirList(conf.Gcp.Project, conf.Gcp.Bucket, ResultPrefix, c.Bool("quiet"))
	case 1:
		instance := c.Args()[0]
		return PrintFileList(conf.Gcp.Project, conf.Gcp.Bucket, filepath.Join(ResultPrefix, instance), c.Bool("quiet"))
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

// CmdResultGet downloads results for a given instance names or result files belonging to an instance.
func CmdResultGet(c *cli.Context) error {

	if c.NArg() < 2 {
		fmt.Printf(chalk.Red.Color("expected at least 2 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	instance := c.Args().First()
	for _, query := range c.Args().Tail() {

		downloadFiles(conf.Gcp.Project, conf.Gcp.Bucket, filepath.Join(ResultPrefix, instance), query)

	}

	return nil

}

func CmdResultGetAll(c *cli.Context) error {
	return nil
}

// ListupFiles lists up files in a bucket associated with a project and which
// have a prefix. Information of found files will be sent to worker function via channgel.
// The worker function will be started as a goroutine.
func ListupFiles(project, bucket, prefix string, worker ListupFilesWorker) error {

	storage, err := util.NewStorage(project, bucket)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	file := make(chan *util.FileInfo, 10)
	done := make(chan struct{})
	errCh := make(chan error)

	go storage.List(prefix, file, errCh)
	go worker(storage, file, done)

loop:
	for {
		select {
		case <-done:
			// printFileBodyWorker ends.
			break loop
		case err = <-errCh:
			// storage.List ends but printFileBodyWorker is still working.
			file <- nil
		}
	}

	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

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

// DownloadFiles downloads files in a bucket associated with a project,
// which has a prefix and satisfies a query.
func downloadFiles(project, bucket, prefix, query string) error {

	return ListupFiles(
		project, bucket, prefix,
		func(storage *util.Storage, file <-chan *util.FileInfo, done chan<- struct{}) {

			var wg sync.WaitGroup
			for {

				info := <-file
				if info == nil {
					break
				}

				fmt.Println(info.Path)

				if matched, _ := filepath.Match(query, info.Name); matched {
					fmt.Println(matched)

					wg.Add(1)
					go func() {

						f, err := os.OpenFile(info.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
						if err != nil {
						}
						defer f.Close()

						buf := bufio.NewWriter(f)
						defer buf.Flush()

						if err := storage.Download(info.Path, buf); err != nil {
						}
						wg.Done()

					}()

				}

			}

			wg.Wait()
			done <- struct{}{}

		})

}
