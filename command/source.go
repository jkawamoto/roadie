package command

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uitable"
	"github.com/jkawamoto/roadie-cli/util"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// CmdSourceList prints source files.
func CmdSourceList(c *cli.Context) error {

	if c.NArg() != 0 {
		fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	storage, _ := util.NewStorage(conf.Gcp.Project, conf.Gcp.Bucket)

	ch := make(chan *util.FileInfo)
	errCh := make(chan error)

	go storage.List(".roadie/source", ch, errCh)

	quiet := c.Bool("quiet")
	table := uitable.New()
	if !quiet {
		table.AddRow("FILE NAME", "SIZE", "TIME CREATED")
	}

loop:
	for {
		select {
		case item := <-ch:
			if item == nil {
				break loop
			}

			if quiet {
				table.AddRow(item.Name)
			} else {
				table.AddRow(item.Name, fmt.Sprintf("%dKB", item.Size/1024), item.TimeCreated.Format(PrintTimeFormat))
			}

		case err := <-errCh:
			return cli.NewExitError(err.Error(), 2)
		}
	}

	fmt.Println(table.String())
	return nil

}

// CmdSourceDelete deletes given source files.
func CmdSourceDelete(c *cli.Context) error {

	if c.NArg() == 0 {
		fmt.Printf(chalk.Red.Color("expected at least 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	storage, err := util.NewStorage(conf.Gcp.Project, conf.Gcp.Bucket)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	for _, name := range c.Args() {

		s.Prefix = fmt.Sprintf("Deleting %s...", name)
		s.FinalMSG = fmt.Sprintf("\nDeleted %s.    \n", name)
		s.Start()

		if err := storage.Delete(".roadie/source/" + name); err != nil {
			s.FinalMSG = fmt.Sprintf(chalk.Red.Color("\nCannot delete %s (%s)\n"), name, err.Error())
		}
		s.Stop()

	}

	if err != nil {
		return cli.NewExitError("Cannot delete all files.", 1)
	}
	return nil

}

// CmdSourceGet downloads one source tarball.
func CmdSourceGet(c *cli.Context) error {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	conf := GetConfig(c)
	storage, err := util.NewStorage(conf.Gcp.Project, conf.Gcp.Bucket)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	name := c.Args()[0]
	output := c.String("o")
	if output == "" {
		output = name
	} else {
		stat, err2 := os.Stat(output)
		if err2 != nil {
			output = name
		} else if stat.IsDir() {
			output = filepath.Join(output, name)
		}
	}

	f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	defer f.Close()

	buf := bufio.NewWriter(f)
	defer buf.Flush()

	if err := storage.Download(".roadie/source/"+name, buf); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}
