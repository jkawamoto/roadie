//
// command/file.go
//
// Copyright (c) 2016 Junpei Kawamoto
//
// This file is part of Roadie.
//
// Roadie is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Roadie is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package command

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uitable"
	"github.com/jkawamoto/pb" // Use `public_pool_add` branch.
	"github.com/jkawamoto/roadie/util"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// AddRecorder is a callback to add file information to a table.
type AddRecorder func(table *uitable.Table, info *util.FileInfo, quiet bool)

// ListupFilesWorker is goroutine of a woker called from listupFiles.
type ListupFilesWorker func(storage *util.Storage, file <-chan *util.FileInfo, done chan<- struct{})

// PrintFileList prints a list of files having a given prefix.
func PrintFileList(project, bucket, prefix string, url, quiet bool) (err error) {

	var headers []string
	if url {
		headers = []string{"FILE NAME", "SIZE", "TIME CREATED", "URL"}
	} else {
		headers = []string{"FILE NAME", "SIZE", "TIME CREATED"}
	}

	return printList(
		project, bucket, prefix, quiet, headers,
		func(table *uitable.Table, info *util.FileInfo, quiet bool) {

			if info.Name != "" {
				if quiet {
					table.AddRow(info.Name)
				} else if url {
					table.AddRow(info.Name, fmt.Sprintf(
						"%dKB", info.Size/1024), info.TimeCreated.Format(PrintTimeFormat),
						fmt.Sprintf("gs://%s/%s", bucket, info.Path))
				} else {
					table.AddRow(info.Name, fmt.Sprintf(
						"%dKB", info.Size/1024), info.TimeCreated.Format(PrintTimeFormat))
				}
			}

		})
}

// PrintDirList prints a list of directoris in a given prefix.
func PrintDirList(project, bucket, prefix string, url, quiet bool) (err error) {

	var headers []string
	if url {
		headers = []string{"INSTANCE NAME", "TIME CREATED", "URL"}
	} else {
		headers = []string{"INSTANCE NAME", "TIME CREATED"}
	}

	// Storing previous folder name.
	prev := ""

	return printList(
		project, bucket, prefix, quiet, headers,
		func(table *uitable.Table, info *util.FileInfo, quiet bool) {

			rel, _ := filepath.Rel(prefix, info.Path)
			rel = filepath.Dir(rel)

			if rel != "." && rel != prev {
				if quiet {
					table.AddRow(rel)
				} else if url {
					table.AddRow(
						rel, info.TimeCreated.Format(PrintTimeFormat),
						fmt.Sprintf("gs://%s/%s", bucket, rel))
				} else {
					table.AddRow(rel, info.TimeCreated.Format(PrintTimeFormat))
				}
				prev = rel
			}

		})
}

func printList(project, bucket, prefix string, quiet bool, headers []string, addRecorder AddRecorder) (err error) {

	// TODO: Refactoring this method using ListupFiles.
	storage, err := util.NewStorage(project, bucket)
	if err != nil {
		return
	}

	ch := make(chan *util.FileInfo)
	errCh := make(chan error)

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "Loading information..."
	s.FinalMSG = fmt.Sprintf("\n%s\r", strings.Repeat(" ", len(s.Prefix)+2))
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

// UploadToGCS uploads a file to a bucket associated with a project.
// Uploaded file will have a given name. This function returns a URL
// for the uploaded file with error object.
func UploadToGCS(project, bucket, prefix, name, input string) (string, error) {

	// TODO: Parallel uploading.
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
	bar.AlwaysUpdate = true
	bar.Start()
	defer bar.Finish()

	if err := storage.Upload(bar.NewProxyReader(file), location); err != nil {
		return "", cli.NewExitError(err.Error(), 2)
	}
	return location.String(), nil

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

	func() {
		for {
			select {
			case <-done:
				// ListupFilesWorker ends.
				return
			case err = <-errCh:
				// storage.List ends but ListupFilesWorker is still working.
				file <- nil
			}
		}
	}()

	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	return nil

}

// DownloadFiles downloads files in a bucket associated with a project,
// which has a prefix and satisfies a query. Downloaded files will be put in
// a given directory.
func DownloadFiles(project, bucket, prefix, dir string, queries []string) error {

	if info, err := os.Stat(dir); err != nil {
		// Given dir does not exist.
		if err2 := os.MkdirAll(dir, 0777); err2 != nil {
			return cli.NewExitError(err2.Error(), 2)
		}
	} else {
		if !info.IsDir() {
			return cli.NewExitError(fmt.Sprintf("Cannot create the directory tree: %s", dir), 2)
		}
	}

	return ListupFiles(
		project, bucket, prefix,
		func(storage *util.Storage, file <-chan *util.FileInfo, done chan<- struct{}) {

			var wg sync.WaitGroup
			fmt.Println("Downloading...")

			pool, _ := pb.StartPool()
			for {

				info := <-file
				if info == nil {
					break
				} else if info.Name == "" {
					continue
				}

				for _, q := range queries {

					if matched, _ := filepath.Match(q, info.Name); matched {

						bar := pb.New64(int64(info.Size)).SetUnits(pb.U_BYTES).Prefix(info.Name)
						pool.Add(bar)

						wg.Add(1)
						go func(info *util.FileInfo, bar *pb.ProgressBar) {

							defer wg.Done()

							filename := filepath.Join(dir, info.Name)
							f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
							if err != nil {
								bar.FinishPrint(fmt.Sprintf(chalk.Red.Color("Cannot create file %s (%s)"), filename, err.Error()))
								return
							}
							defer f.Close()

							buf := bufio.NewWriter(io.MultiWriter(f, bar))
							defer buf.Flush()

							if err := storage.Download(info.Path, buf); err != nil {
								bar.FinishPrint(fmt.Sprintf(chalk.Red.Color("Cannot doenload %s (%s)"), info.Name, err.Error()))
							} else {
								bar.Finish()
							}

						}(info, bar)

						break
					}

				}

			}

			wg.Wait()
			pool.Stop()
			done <- struct{}{}

		})

}

// DeleteFiles deletes files in a bucket associated with a project,
// which has a prefix and satisfies a query.
func DeleteFiles(project, bucket, prefix string, queries []string) error {

	return ListupFiles(
		project, bucket, prefix,
		func(storage *util.Storage, file <-chan *util.FileInfo, done chan<- struct{}) {

			var wg sync.WaitGroup
			fmt.Println("Deleting...")

			for {

				info := <-file
				if info == nil {
					break
				} else if info.Name == "" {
					continue
				}

				for _, q := range queries {

					if matched, _ := filepath.Match(q, info.Name); matched {

						wg.Add(1)
						go func(info *util.FileInfo) {
							defer wg.Done()
							if err := storage.Delete(info.Path); err != nil {
								fmt.Printf(chalk.Red.Color("Cannot delete %s (%s)\n"), info.Path, err.Error())
							}
						}(info)

						// Break from checking queries.
						break
					}

				}

			}

			wg.Wait()
			done <- struct{}{}

		})

}

// PrintFileBody prints file bodies in a bucket associated with a project,
// which has a prefix ans satisfies query. If quiet is ture, additional messages
// well be suppressed.
func PrintFileBody(project, bucket, prefix, query string, quiet bool) error {

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
