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

	"github.com/jkawamoto/pb" // Use `public_pool_add` branch.
	"github.com/jkawamoto/roadie/util"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// ListupFilesWorker is goroutine of a woker called from listupFiles.
type ListupFilesWorker func(storage *util.Storage, file <-chan *util.FileInfo, done chan<- struct{})

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
func ListupFiles(project, bucket, prefix string, worker ListupFilesWorker) (err error) {

	storage, err := util.NewStorage(project, bucket)
	if err != nil {
		return
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

	return

}

// DownloadFiles downloads files in a bucket associated with a project,
// which has a prefix and satisfies a query. Downloaded files will be put in
// a given directory.
func DownloadFiles(project, bucket, prefix, dir string, queries []string) (err error) {

	var info os.FileInfo
	if info, err = os.Stat(dir); err != nil {
		// Given dir does not exist.
		if err = os.MkdirAll(dir, 0777); err != nil {
			return
		}
	} else {
		if !info.IsDir() {
			return fmt.Errorf("Cannot create the directory tree: %s", dir)
		}
	}

	return ListupFiles(
		project, bucket, prefix,
		func(storage *util.Storage, file <-chan *util.FileInfo, done chan<- struct{}) {

			var wg sync.WaitGroup
			fmt.Println("Downloading...")

			pool, _ := pb.StartPool()
			defer pool.Stop()

			for {

				info := <-file
				if info == nil {
					break
				} else if info.Name == "" {
					continue
				}

				if match(queries, info.Name) {

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

				}

			}

			wg.Wait()
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

				if match(queries, info.Name) {

					wg.Add(1)
					go func(info *util.FileInfo) {

						defer wg.Done()
						if err := storage.Delete(info.Path); err != nil {
							fmt.Printf(chalk.Red.Color("Cannot delete %s (%s)\n"), info.Path, err.Error())
						}

					}(info)

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

// match returns true if there are at least one pattern matching to name.
func match(patterns []string, name string) bool {
	for _, pat := range patterns {
		if res, _ := filepath.Match(pat, name); res {
			return true
		}
	}
	return false
}
