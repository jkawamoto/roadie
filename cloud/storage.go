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
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"

	pb "gopkg.in/cheggaaa/pb.v1"

	"github.com/ulikunitz/xz"
)

// StorageManager defines methods which a storage service provider must provides.
// Each method takes a URL to point a file stored in this storage.
// The URL should be
// - roadie://category/path
// where category is one of
// - script.SourcePrefix
// - script.DataPrefix
// - script.ResultPrefix
type StorageManager interface {

	// Upload a given stream to a given URL.
	Upload(ctx context.Context, loc *url.URL, in io.Reader) error

	// Download a file pointed by a given URL and write it to a given stream.
	Download(ctx context.Context, loc *url.URL, out io.Writer) error

	// GetFileInfo retrieves information of a file pointed by a given URL.
	GetFileInfo(ctx context.Context, loc *url.URL) (*FileInfo, error)

	// List up files of which URLs start with a given URL.
	// It takes a handler; information of found files are sent to it.
	List(ctx context.Context, loc *url.URL, handler FileInfoHandler) error

	// Delete a file pointed by a given URL.
	Delete(ctx context.Context, loc *url.URL) error
}

// FileInfoHandler is a handler to receive a file info.
type FileInfoHandler func(*FileInfo) error

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
		log = ioutil.Discard
	}

	s = &Storage{
		service: servicer,
		Log:     log,
	}
	return

}

// UploadFile uploads a file where a given URL points.
func (s *Storage) UploadFile(ctx context.Context, loc *url.URL, input string) (err error) {

	file, err := os.Open(input)
	if err != nil {
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return
	}

	fmt.Fprintln(s.Log, "Uploading...")
	bar := pb.New64(int64(info.Size())).SetUnits(pb.U_BYTES).Prefix(path.Base(loc.Path))
	bar.Output = s.Log
	bar.AlwaysUpdate = true
	bar.Start()
	defer bar.Finish()

	return s.service.Upload(ctx, loc, bar.NewProxyReader(bufio.NewReader(file)))

}

// ListupFiles lists up files which location is matching to a given URL.
// Information of found files will be passed to a handler.
// If the handler returns non nil value, the listing up will be canceled.
// In this case, this function also returns the given error value.
func (s *Storage) ListupFiles(ctx context.Context, loc *url.URL, handler FileInfoHandler) (err error) {

	return s.service.List(ctx, loc, handler)

}

// DownloadFiles downloads files matching a given prefix and queries.
// Downloaded files will be put in a given directory.
func (s *Storage) DownloadFiles(ctx context.Context, prefix *url.URL, dir string, queries []string) (err error) {

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
					bar.FinishPrint(fmt.Sprintf("Cannot create file %s (%s)", filename, goerr.Error()))
					return goerr
				}
				defer f.Close()

				writer := bufio.NewWriter(f)
				defer writer.Flush()

				goerr = s.service.Download(ctx, info.URL, io.MultiWriter(writer, bar))
				if goerr != nil {
					bar.FinishPrint(fmt.Sprintf("Cannot download %s (%s)", info.Name, goerr.Error()))
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

// DeleteFiles deletes files matching a given URL prefix and queries.
func (s *Storage) DeleteFiles(ctx context.Context, prefix *url.URL, queries []string) (err error) {

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

			eg.Go(func() (err error) {

				err = s.service.Delete(ctx, info.URL)
				if err != nil {
					fmt.Fprintf(s.Log, "Cannot delete %v: %v\n", info.URL, err.Error())
				} else {
					fmt.Fprintln(s.Log, info.URL)
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

// PrintFileBody prints file bodies which has a prefix and satisfies query.
// If header is ture, additional messages well be printed.
func (s *Storage) PrintFileBody(ctx context.Context, prefix *url.URL, query string, output io.Writer, header bool) error {

	return s.ListupFiles(ctx, prefix, func(info *FileInfo) error {

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

			return s.service.Download(ctx, info.URL, output)
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
