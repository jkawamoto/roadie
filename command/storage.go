package command

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uitable"
	"github.com/jkawamoto/pb"
	"github.com/jkawamoto/roadie-cli/util"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// AddRecorder is a callback to add file information to a table.
type AddRecorder func(table *uitable.Table, info *util.FileInfo, quiet bool)

// ListupFilesWorker is goroutine of a woker called from listupFiles.
type ListupFilesWorker func(storage *util.Storage, file <-chan *util.FileInfo, done chan<- struct{})

// PrintFileList prints a list of files having a given prefix.
func PrintFileList(project, bucket, prefix string, quiet bool) (err error) {
	return printList(
		project, bucket, prefix, quiet, []string{"FILE NAME", "SIZE", "TIME CREATED"},
		func(table *uitable.Table, info *util.FileInfo, quiet bool) {

			if info.Name != "" {
				if quiet {
					table.AddRow(info.Name)
				} else {
					table.AddRow(info.Name, fmt.Sprintf("%dKB", info.Size/1024), info.TimeCreated.Format(PrintTimeFormat))
				}
			}

		})
}

// PrintDirList prints a list of directoris in a given prefix.
func PrintDirList(project, bucket, prefix string, quiet bool) (err error) {

	return printList(
		project, bucket, prefix, quiet, []string{"INSTANCE NAME", "TIME CREATED"},
		func(table *uitable.Table, info *util.FileInfo, quiet bool) {

			dir, _ := filepath.Split(info.Path)
			name := strings.Replace(dir[len(ResultPrefix):], "/", "", -1)
			if info.Name != "" {
				if quiet {
					table.AddRow(name)
				} else {
					table.AddRow(name, info.TimeCreated.Format(PrintTimeFormat))
				}
			}

		})
}

func printList(project, bucket, prefix string, quiet bool, headers []string, addRecorder AddRecorder) (err error) {

	storage, err := util.NewStorage(project, bucket)
	if err != nil {
		return
	}

	ch := make(chan *util.FileInfo)
	errCh := make(chan error)

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "Loading information..."
	s.FinalMSG = "\n                        \r"
	s.Start()

	go storage.List(prefix, ch, errCh)

	table := uitable.New()
	if !quiet {
		rawHeaders := make([]interface{}, len(headers))
		for i, v := range headers {
			rawHeaders[i] = v
		}
		table.AddRow(rawHeaders...)
	}

loop:
	for {
		select {
		case item := <-ch:
			if item == nil {
				break loop
			}
			addRecorder(table, item, quiet)
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
func UploadToGCS(project, bucket, prefix, name, input string) (string, error) {

	storage, err := util.NewStorage(project, bucket)
	if err != nil {
		return "", err
	}

	if name == "" {
		name = filepath.Base(input)
	}
	location := util.CreateURL(bucket, prefix, name)

	info, err := os.Stat(input)
	if err != nil {
		return "", err
	}

	file, err := os.Open(input)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fmt.Println("Uploading...")
	bar := pb.New64(int64(info.Size())).SetUnits(pb.U_BYTES).Prefix(name)
	bar.Start()
	defer bar.Finish()

	if err := storage.Upload(bar.NewProxyReader(file), location); err != nil {
		return "", cli.NewExitError(err.Error(), 2)
	}
	return location.String(), nil

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

	filename := filepath.Join(prefix, name)
	info, err := storage.Status(filename)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	defer f.Close()

	fmt.Println("Downloading...")
	bar := pb.New64(int64(info.Size)).SetUnits(pb.U_BYTES).Prefix(name)
	bar.Start()
	defer bar.Finish()

	writer := io.MultiWriter(f, bar)
	buf := bufio.NewWriter(writer)
	defer buf.Flush()

	if err := storage.Download(filename, buf); err != nil {
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

	// TODO: Support glob and use ListupFiles.

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
