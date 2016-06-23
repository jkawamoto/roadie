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

// PrintFileList prints a list of files having a given prefix.
func PrintFileList(project, bucket, prefix string, quiet bool) (err error) {

	storage, _ := util.NewStorage(project, bucket)

	ch := make(chan *util.FileInfo)
	errCh := make(chan error)

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "Loading information..."
	s.FinalMSG = "\n                        \r"
	s.Start()

	go storage.List(prefix, ch, errCh)

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

		case err = <-errCh:
			break loop
		}
	}

	s.Stop()
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	fmt.Println(table.String())
	return nil

}

// UploadToGCS uploads a file to GCS.
func UploadToGCS(project, bucket, prefix, name, filepath string) error {

	// TODO: sould return the URL storing the new object.
	storage, err := util.NewStorage(project, bucket)
	if err != nil {
		return err
	}

	if name == "" {
		name = util.Basename(filepath)
	}
	location := util.CreateURL(bucket, prefix, name)

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Uploading to %s...", location)
	s.FinalMSG = fmt.Sprintf("\n                   \rUploaded to %s.   \n", chalk.Bold.TextStyle(location.String()))
	s.Start()
	defer s.Stop()

	if err := storage.Upload(filepath, location); err != nil {
		s.FinalMSG = fmt.Sprintf(chalk.Red.Color("Cannot upload file %s. (%s)"), filepath, err.Error())
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// DownloadFromGCS downloads a file from GCS
func DownloadFromGCS(project, bucket, prefix, name, output string) error {

	storage, err := util.NewStorage(project, bucket)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	if output == "" {
		output = name
	} else {
		stat, err2 := os.Stat(output)
		if err2 == nil && stat.IsDir() {
			output = filepath.Join(output, name)
		}
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "Downloading..."
	s.FinalMSG = fmt.Sprintf("\n              \rDownloaded to %s.\n", chalk.Bold.TextStyle(output))
	s.Start()
	defer s.Stop()

	f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	defer f.Close()

	buf := bufio.NewWriter(f)
	defer buf.Flush()

	if err := storage.Download(filepath.Join(prefix, name), buf); err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// DeleteFromGCS deletes a file from GCS.
func DeleteFromGCS(project, bucket, prefix string, names []string) error {

	storage, err := util.NewStorage(project, bucket)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	for _, name := range names {

		s.Prefix = fmt.Sprintf("Deleting %s...", name)
		s.FinalMSG = fmt.Sprintf("\nDeleted %s.    \n", chalk.Bold.TextStyle(name))
		s.Start()

		if err := storage.Delete(filepath.Join(prefix, name)); err != nil {
			s.FinalMSG = fmt.Sprintf(chalk.Red.Color("\nCannot delete %s (%s)\n"), name, err.Error())
		}
		s.Stop()

	}

	if err != nil {
		return cli.NewExitError("Cannot delete all files.", 1)
	}
	return nil

}
