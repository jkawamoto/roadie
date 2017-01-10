//
// command/cloud/storage.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
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

package cloud

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/cheggaaa/pb"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/jkawamoto/roadie/config"
	"github.com/urfave/cli"
)

// Storage provides APIs to access a cloud storage.
type Storage struct {
	// Context this Storage associates with.
	ctx context.Context
	// Servicer.
	service storageServicer
	// Writer logs to be printed.
	Log io.Writer
}

// FileInfoHandler is a handler to recieve a file info.
type FileInfoHandler func(*FileInfo) error

// storageServicer defines APIs a storage service provider must have.
type storageServicer interface {
	CreateIfNotExists() error
	Upload(in io.Reader, location *url.URL) error
	Download(filename string, out io.Writer) error
	Status(filename string) (*FileInfo, error)
	List(prefix string, handler FileInfoHandler) error
	Delete(name string) error
}

// NewStorage creates a cloud storage accessor with a given context.
// The context must have a Config.
func NewStorage(ctx context.Context) *Storage {

	return &Storage{
		ctx:     ctx,
		service: NewCloudStorageService(ctx),
		Log:     os.Stderr,
	}

}

// PrepareBucket makes a bucket if it doesn't exist under a given context.
// The given context must have a config.
func (s *Storage) PrepareBucket() error {
	return s.service.CreateIfNotExists()
}

// UploadFile uploads a file to a bucket associated with a project under a given
// context. Uploaded file will have a given name. This function returns a URL
// for the uploaded file with error object.
func (s *Storage) UploadFile(prefix, name, input string) (string, error) {

	var err error
	cfg, err := config.FromContext(s.ctx)
	if err != nil {
		return "", err
	}

	if err = s.service.CreateIfNotExists(); err != nil {
		return "", err
	}

	if name == "" {
		name = filepath.Base(input)
	}
	location := util.CreateURL(cfg.Bucket, prefix, name)

	info, err := os.Stat(input)
	if err != nil {
		return "", err
	}

	file, err := os.Open(input)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fmt.Fprintln(s.Log, "Uploading...")
	bar := pb.New64(int64(info.Size())).SetUnits(pb.U_BYTES).Prefix(name)
	bar.Output = s.Log
	bar.AlwaysUpdate = true
	bar.Start()
	defer bar.Finish()

	if err := s.service.Upload(bar.NewProxyReader(file), location); err != nil {
		return "", cli.NewExitError(err.Error(), 2)
	}
	return location.String(), nil

}

// ListupFiles lists up files in a bucket associated with a project and which
// have a prefix under a given context. Information of found files will be passed to a handler.
// If the handler returns non nil value, the listing up will be canceled.
// In this case, this function also returns the given error value.
func (s *Storage) ListupFiles(prefix string, handler FileInfoHandler) (err error) {

	return s.service.List(prefix, handler)

}

// DownloadFiles downloads files in a bucket associated with a project,
// which has a prefix and satisfies a query under a given context.
// Downloaded files will be put in a given directory.
func (s *Storage) DownloadFiles(prefix, dir string, queries []string) (err error) {

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

	fmt.Fprintln(s.Log, "Downloading...")
	pool, _ := pb.StartPool()
	defer pool.Stop()

	var wg sync.WaitGroup
	defer wg.Wait()

	return s.ListupFiles(prefix, func(info *FileInfo) error {

		select {
		case <-s.ctx.Done():
			return s.ctx.Err()

		default:
			// If Name is empty, it might be a folder or a special file.
			if info.Name == "" {
				return nil
			}

			if match(queries, info.Name) {

				bar := pb.New64(int64(info.Size)).SetUnits(pb.U_BYTES).Prefix(info.Name)
				bar.Output = s.Log
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

					if err := s.service.Download(info.Path, buf); err != nil {
						bar.FinishPrint(fmt.Sprintf(chalk.Red.Color("Cannot download %s (%s)"), info.Name, err.Error()))
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
func (s *Storage) DeleteFiles(prefix string, queries []string) error {

	fmt.Fprintln(s.Log, "Deleting...")
	pool, _ := pb.StartPool()
	defer pool.Stop()

	var wg sync.WaitGroup
	defer wg.Wait()

	return s.ListupFiles(prefix, func(info *FileInfo) error {

		select {
		case <-s.ctx.Done():
			return s.ctx.Err()

		default:

			// If Name is empty, it might be a folder or a special file.
			if info.Name == "" {
				return nil
			}

			if match(queries, info.Name) {

				bar := pb.New64(int64(info.Size)).SetUnits(pb.U_BYTES).Prefix(info.Name)
				bar.Output = s.Log
				pool.Add(bar)

				wg.Add(1)
				go func(info *FileInfo) {
					defer wg.Done()
					defer bar.Add64(int64(info.Size))

					if err := s.service.Delete(info.Path); err != nil {
						fmt.Fprintf(s.Log, chalk.Red.Color("Cannot delete %s (%s)\n"), info.Path, err.Error())
					}
				}(info)

			}

			return nil
		}

	})
}

// PrintFileBody prints file bodies which has a prefix and satisfies query under a context.
// If header is ture, additional messages well be printed.
func (s *Storage) PrintFileBody(prefix, query string, output io.Writer, header bool) error {

	return s.ListupFiles(prefix, func(info *FileInfo) error {

		select {
		case <-s.ctx.Done():
			return s.ctx.Err()

		default:
			if info.Name != "" && strings.HasPrefix(info.Name, query) {
				if header {
					fmt.Fprintf(output, "*** %s ***\n", info.Name)
				}
				return s.service.Download(info.Path, output)
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
