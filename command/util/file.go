//
// command/util/file.go
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

package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/context"

	"github.com/cheggaaa/pb"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/config"
	"github.com/urfave/cli"
)

// ListupFilesHandler is a handler of LlistupFiles.
type ListupFilesHandler func(storage *Storage, info *FileInfo) error

// PrepareBucket makes a bucket if it doesn't exist under a given context.
// The given context must have a config.
func PrepareBucket(ctx context.Context) error {

	// Check a specified bucket exists and create it if not.
	if storage, e := NewStorage(ctx); e != nil {
		return e
	} else if e := storage.CreateIfNotExists(); e != nil {
		return e
	}
	return nil

}

// UploadFiles uploads a file to a bucket associated with a project under a given
// context. Uploaded file will have a given name. This function returns a URL
// for the uploaded file with error object.
func UploadFiles(ctx context.Context, prefix, name, input string) (string, error) {

	cfg, ok := config.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("Config is not attached to the given Context: %s", ctx)
	}

	storage, err := NewStorage(ctx)
	if err != nil {
		return "", err
	}
	if err = storage.CreateIfNotExists(); err != nil {
		return "", err
	}

	if name == "" {
		name = filepath.Base(input)
	}
	location := CreateURL(cfg.Gcp.Bucket, prefix, name)

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
// have a prefix under a given context. Information of found files will be passed to a handler.
// If the handler returns non nil value, the listing up will be canceled.
// In this case, this function also returns the given error value.
func ListupFiles(ctx context.Context, prefix string, handler ListupFilesHandler) (err error) {

	storage, err := NewStorage(ctx)
	if err != nil {
		return
	}
	if err = storage.CreateIfNotExists(); err != nil {
		return
	}

	return storage.List(prefix, func(info *FileInfo) error {
		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			return handler(storage, info)
		}
	})

}

// DownloadFiles downloads files in a bucket associated with a project,
// which has a prefix and satisfies a query under a given context.
// Downloaded files will be put in a given directory.
func DownloadFiles(ctx context.Context, prefix, dir string, queries []string) (err error) {

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

	fmt.Println("Downloading...")
	pool, _ := pb.StartPool()
	defer pool.Stop()

	var wg sync.WaitGroup
	defer wg.Wait()

	return ListupFiles(ctx, prefix, func(storage *Storage, info *FileInfo) error {

		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			// If Name is empty, it might be a folder or a special file.
			if info.Name == "" {
				return nil
			}

			if match(queries, info.Name) {

				bar := pb.New64(int64(info.Size)).SetUnits(pb.U_BYTES).Prefix(info.Name)
				pool.Add(bar)

				wg.Add(1)
				go func(info *FileInfo, bar *pb.ProgressBar) {

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

			return nil
		}

	})
}

// DeleteFiles deletes files in a bucket associated with a project,
// which has a prefix and satisfies a query. This request will be done under a
// given context.
func DeleteFiles(ctx context.Context, prefix string, queries []string) error {

	fmt.Println("Deleting...")

	var wg sync.WaitGroup
	defer wg.Wait()

	return ListupFiles(ctx, prefix, func(storage *Storage, info *FileInfo) error {

		select {
		case <-ctx.Done():
			return ctx.Err()

		default:

			// If Name is empty, it might be a folder or a special file.
			if info.Name == "" {
				return nil
			}

			if match(queries, info.Name) {

				wg.Add(1)
				go func(info *FileInfo) {

					defer wg.Done()
					if err := storage.Delete(info.Path); err != nil {
						fmt.Printf(chalk.Red.Color("Cannot delete %s (%s)\n"), info.Path, err.Error())
					}

				}(info)

			}

			return nil
		}

	})
}

// PrintFileBody prints file bodies in a bucket associated with a project,
// which has a prefix and satisfies query under a context.
// If quiet is ture, additional messages well be suppressed.
func PrintFileBody(ctx context.Context, project, bucket, prefix, query string, quiet bool) error {

	return ListupFiles(ctx, prefix, func(storage *Storage, info *FileInfo) error {

		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			// If Name is empty, it might be a folder or a special file.
			if info == nil {
				return nil
			}

			if info.Name != "" && strings.HasPrefix(info.Name, query) {
				if !quiet {
					fmt.Printf(chalk.Bold.TextStyle("*** %s ***\n"), info.Name)
				}
				if err := storage.Download(info.Path, os.Stdout); err != nil {
					fmt.Printf(chalk.Red.Color("Cannot download %s (%s)."), info.Name, err.Error())
				}
			}

			return nil
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
