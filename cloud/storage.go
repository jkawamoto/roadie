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
	"github.com/urfave/cli"
)

// Storage provides APIs to access a cloud storage.
type Storage struct {
	// Servicer.
	service StorageServicer
	// Writer logs to be printed.
	Log io.Writer
}

// NewStorage creates a cloud storage accessor with a given context.
func NewStorage(servicer StorageServicer, log io.Writer) (s *Storage) {

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
func (s *Storage) UploadFile(ctx context.Context, prefix, name, input string) (uri string, err error) {

	if name == "" {
		name = filepath.Base(input)
	}
	location := filepath.Join(prefix, name)

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

	if uri, err = s.service.Upload(ctx, location, bar.NewProxyReader(file)); err != nil {
		return "", cli.NewExitError(err.Error(), 2)
	}
	return

}

// ListupFiles lists up files in a bucket associated with a project and which
// have a prefix under a given context. Information of found files will be passed to a handler.
// If the handler returns non nil value, the listing up will be canceled.
// In this case, this function also returns the given error value.
func (s *Storage) ListupFiles(ctx context.Context, prefix string, handler FileInfoHandler) (err error) {

	return s.service.List(ctx, prefix, handler)

}

// DownloadFiles downloads files in a bucket associated with a project,
// which has a prefix and satisfies a query under a given context.
// Downloaded files will be put in a given directory.
func (s *Storage) DownloadFiles(ctx context.Context, prefix, dir string, queries []string) (err error) {

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
	err = s.ListupFiles(ctx, prefix, func(info *FileInfo) error {

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

				goerr = s.service.Download(ctx, info.Path, io.MultiWriter(writer, bar))
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
func (s *Storage) DeleteFiles(ctx context.Context, prefix string, queries []string) error {

	fmt.Fprintln(s.Log, "Deleting...")
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
	err = s.ListupFiles(ctx, prefix, func(info *FileInfo) error {

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
				defer bar.Add64(int64(info.Size))
				goerr := s.service.Delete(ctx, info.Path)
				if goerr != nil {
					fmt.Fprintf(s.Log, chalk.Red.Color("Cannot delete %s (%s)\n"), info.Path, goerr.Error())
				}
				return goerr
			})

		}
		return nil

	})

	if err != nil {
		return err
	}
	return eg.Wait()

}

// PrintFileBody prints file bodies which has a prefix and satisfies query under a context.
// If header is ture, additional messages well be printed.
func (s *Storage) PrintFileBody(ctx context.Context, prefix, query string, output io.Writer, header bool) error {

	return s.ListupFiles(ctx, prefix, func(info *FileInfo) error {

		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			if info.Name != "" && strings.HasPrefix(info.Name, query) {
				if header {
					fmt.Fprintf(output, "*** %s ***\n", info.Name)
				}
				return s.service.Download(ctx, info.Path, output)
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
