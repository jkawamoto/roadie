//
// cloud/storage.go
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
// along with Roadie.  If not, see <http://www.gnu.org/licenses/>.
//

package cloud

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"

	pb "gopkg.in/cheggaaa/pb.v1"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/ulikunitz/xz"
	"github.com/urfave/cli"
)

// Storage provides APIs to access a cloud storage.
type Storage struct {
	// Servicer.
	service StorageManager
	// TODO: Delete or update this.
	// Writer logs to be printed.
	Log io.Writer
}

// NewStorage creates a cloud storage accessor with a given context.
func NewStorage(servicer StorageManager, log io.Writer) (s *Storage) {

	if log == nil {
		log = os.Stderr
	}

	s = &Storage{
		service: servicer,
		Log:     log,
	}
	return

}

// UploadFile uploads a file to a bucket associated with a project under a given
// context. Uploaded file will have a given name. This function returns a URL
// for the uploaded file with error object.
func (s *Storage) UploadFile(ctx context.Context, container, name, input string) (uri string, err error) {

	if name == "" {
		name = filepath.Base(input)
	}

	file, err := os.Open(input)
	if err != nil {
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	fmt.Fprintln(s.Log, "Uploading...")
	bar := pb.New64(int64(info.Size())).SetUnits(pb.U_BYTES).Prefix(name)
	bar.Output = s.Log
	bar.AlwaysUpdate = true
	bar.Start()
	defer bar.Finish()

	if uri, err = s.service.Upload(ctx, container, name, bar.NewProxyReader(file)); err != nil {
		return "", cli.NewExitError(err.Error(), 2)
	}
	return

}

// ListupFiles lists up files in a bucket associated with a project and which
// have a prefix under a given context. Information of found files will be passed to a handler.
// If the handler returns non nil value, the listing up will be canceled.
// In this case, this function also returns the given error value.
func (s *Storage) ListupFiles(ctx context.Context, container, prefix string, handler FileInfoHandler) (err error) {

	return s.service.List(ctx, container, prefix, handler)

}

// DownloadFiles downloads files in a bucket associated with a project,
// which has a prefix and satisfies a query under a given context.
// Downloaded files will be put in a given directory.
func (s *Storage) DownloadFiles(ctx context.Context, container, prefix, dir string, queries []string) (err error) {

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

	fmt.Fprintln(s.Log, "")
	pool, err := pb.StartPool()
	if err != nil {
		log.Println("cannot create a progress bar:", err.Error())
	} else {
		pool.Output = s.Log
		defer pool.Stop()
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)
	err = s.ListupFiles(ctx, container, prefix, func(info *FileInfo) error {

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// If Name is empty, it might be a folder or a special file.
		if info.Name == "" {
			return nil
		}

		if match(queries, info.Name) {

			bar := pb.New64(int64(info.Size)).SetUnits(pb.U_BYTES).Prefix(info.Name)
			if pool != nil {
				pool.Add(bar)
			}

			eg.Go(func() error {
				var goerr error
				filename := filepath.Join(dir, info.Name)
				f, goerr := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
				if goerr != nil {
					bar.FinishPrint(fmt.Sprintf(chalk.Red.Color("Cannot create file %s (%s)"), filename, goerr.Error()))
					return goerr
				}
				defer f.Close()

				writer := bufio.NewWriter(f)
				defer writer.Flush()

				goerr = s.service.Download(ctx, container, info.Path, io.MultiWriter(writer, bar))
				if goerr != nil {
					bar.FinishPrint(fmt.Sprintf(chalk.Red.Color("Cannot download %s (%s)"), info.Name, goerr.Error()))
				} else {
					bar.Finish()
				}
				return goerr
			})

		}
		return nil
	})

	if err != nil {
		return
	}
	return eg.Wait()

}

// DeleteFiles deletes files in a bucket associated with a project,
// which has a prefix and satisfies a query. This request will be done under a
// given context.
func (s *Storage) DeleteFiles(ctx context.Context, container, prefix string, queries []string) (err error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)
	err = s.ListupFiles(ctx, container, prefix, func(info *FileInfo) error {

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// If Name is empty, it might be a folder or a special file.
		if info.Name == "" {
			return nil
		}

		if match(queries, info.Name) {

			eg.Go(func() (err error) {
				err = s.service.Delete(ctx, container, info.Path)
				if err != nil {
					fmt.Fprintf(s.Log, chalk.Red.Color("Cannot delete %s (%s)\n"), info.Path, err.Error())
				} else {
					fmt.Fprintln(s.Log, info.Path)
				}
				return
			})

		}
		return nil

	})

	if err != nil {
		return
	}
	return eg.Wait()

}

// PrintFileBody prints file bodies which has a prefix and satisfies query under a context.
// If header is ture, additional messages well be printed.
func (s *Storage) PrintFileBody(ctx context.Context, container, prefix, query string, output io.Writer, header bool) error {

	return s.ListupFiles(ctx, container, prefix, func(info *FileInfo) error {

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if info.Name != "" && strings.HasPrefix(info.Name, query) {

			if header {
				fmt.Fprintf(output, "*** %s ***\n", info.Name)
			}

			if strings.HasSuffix(info.Name, ".xz") {
				pipeReader, pipeWriter := io.Pipe()
				defer pipeWriter.Close()

				xzReader, err := xz.NewReader(pipeReader)
				if err != nil {
					return err
				}

				go io.Copy(output, xzReader)
				output = pipeWriter
			}

			return s.service.Download(ctx, container, info.Path, output)
		}

		return nil

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
