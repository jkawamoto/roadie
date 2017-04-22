//
// cloud/gce/storage.go
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

package gce

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"path/filepath"
	"strings"

	"google.golang.org/api/iterator"

	"github.com/jkawamoto/roadie/cloud"

	"cloud.google.com/go/storage"
)

// StorageService implements cloud.StorageServicer interface for accessing GCE's
// cloud storage.
type StorageService struct {
	// Config is a reference for a configuration of GCP.
	Config *GcpConfig
	// Logger
	Logger *log.Logger
}

// NewStorageService creates a new storage accessor to a bucket specified in a
// given configuration.
func NewStorageService(ctx context.Context, cfg *GcpConfig, logger *log.Logger) (s *StorageService, err error) {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	cli, err := storage.NewClient(ctx)
	if err != nil {
		return
	}
	defer cli.Close()

	// Check the given project has the given bucket; if not, create a new bucket.
	logger.Println("Checking bucket", cfg.Bucket, "exists")

	bucket := cli.Bucket(cfg.Bucket)
	_, err = bucket.Attrs(ctx)
	if err != nil {
		logger.Println("Creating bucket", cfg.Bucket)
		err = bucket.Create(ctx, cfg.Project, nil)
		if err != nil {
			return
		}
	}

	s = &StorageService{
		Config: cfg,
		Logger: logger,
	}
	return

}

// Upload a file to a location.
func (s *StorageService) Upload(ctx context.Context, container, filename string, in io.Reader) (uri string, err error) {

	s.Logger.Println("Uploading file", filename)
	client, err := storage.NewClient(ctx)
	if err != nil {
		return
	}
	defer client.Close()

	path := filepath.Join(StoragePrefix, container, filename)
	obj := client.Bucket(s.Config.Bucket).Object(path)
	writer := obj.NewWriter(ctx)
	size, err := io.Copy(writer, in)
	writer.Close()
	if err != nil {
		return
	}

	info, err := obj.Attrs(ctx)
	if err != nil {
		return
	} else if info.Size != size {
		obj.Delete(ctx)
		return "", fmt.Errorf("Faild to upload object %v", path)
	}

	u := url.URL{
		Scheme: "gs",
		Host:   s.Config.Bucket,
		Path:   filepath.Join(StoragePrefix, path),
	}
	uri = u.String()
	s.Logger.Println("Finished uploading the file to", uri)
	return

}

// Download downloads a file and write it to a given writer.
func (s *StorageService) Download(ctx context.Context, container, filename string, out io.Writer) (err error) {

	s.Logger.Println("Downloading file", filename)
	client, err := storage.NewClient(ctx)
	if err != nil {
		return
	}
	defer client.Close()

	obj := client.Bucket(s.Config.Bucket).Object(filepath.Join(StoragePrefix, container, filename))
	info, err := obj.Attrs(ctx)
	if err != nil {
		return
	}

	reader, err := obj.NewReader(ctx)
	if err != nil {
		return
	}
	defer reader.Close()

	size, err := io.Copy(out, reader)
	if size != info.Size {
		return fmt.Errorf("Faild to download object %v", filename)
	}

	s.Logger.Println("Finished downloading the file")
	return

}

// GetFileInfo returns a file status of an object.
func (s *StorageService) GetFileInfo(ctx context.Context, container, filename string) (info *cloud.FileInfo, err error) {

	s.Logger.Println("Retrieving information of file", filename)
	client, err := storage.NewClient(ctx)
	if err != nil {
		return
	}
	defer client.Close()

	attrs, err := client.Bucket(s.Config.Bucket).Object(filepath.Join(StoragePrefix, container, filename)).Attrs(ctx)
	if err != nil {
		return
	}

	info = &cloud.FileInfo{
		Name:        filepath.Base(attrs.Name),
		Path:        attrs.Name,
		TimeCreated: attrs.Created,
		Size:        attrs.Size,
	}

	s.Logger.Println("Finished retrieving the file information")
	return

}

// List searches items, i.e. files and folders, matching a given prefix.
// Found items will be passed to a given handler item by item.
// If the handler returns a non nil value, listing up will be canceled.
// In that case, this function will also return the given value.
func (s *StorageService) List(ctx context.Context, container, prefix string, handler cloud.FileInfoHandler) (err error) {

	s.Logger.Printf(`Retrieving the list of files matching to prefix "%v"`, prefix)
	client, err := storage.NewClient(ctx)
	if err != nil {
		return
	}
	defer client.Close()

	iter := client.Bucket(s.Config.Bucket).Objects(ctx, &storage.Query{
		Prefix:   filepath.Join(StoragePrefix, container, prefix),
		Versions: false,
	})
	for {
		attrs, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return err
		}

		var base string
		if strings.HasSuffix(attrs.Name, "/") {
			base = ""
		} else {
			base = filepath.Base(attrs.Name)
		}
		dir, _ := filepath.Rel(filepath.Join(StoragePrefix, container), attrs.Name)
		err = handler(&cloud.FileInfo{
			Name:        base,
			Path:        dir,
			TimeCreated: attrs.Created,
			Size:        attrs.Size,
		})
		if err != nil {
			return err
		}
	}

	s.Logger.Println("Finished retrieving the file list")
	return

}

// Delete deletes a given file.
func (s *StorageService) Delete(ctx context.Context, container, filename string) (err error) {

	s.Logger.Println("Deleting file", filename)
	client, err := storage.NewClient(ctx)
	if err != nil {
		return
	}
	defer client.Close()

	err = client.Bucket(s.Config.Bucket).
		Object(filepath.Join(StoragePrefix, container, filename)).
		Delete(ctx)
	if err != nil {
		return
	}

	s.Logger.Println("Finished deleting file", filename)
	return

}
