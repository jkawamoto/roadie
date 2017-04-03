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
	// Client of the GCE's storage service.
	client *storage.Client
	// Config is a reference for a configuration of GCP.
	Config *GcpConfig
}

// NewStorageService creates a new storage accessor to a bucket specified in a
// given configuration.
func NewStorageService(ctx context.Context, cfg *GcpConfig) (s *StorageService, err error) {

	cli, err := storage.NewClient(ctx)
	if err != nil {
		return
	}
	s = &StorageService{
		client: cli,
		Config: cfg,
	}

	// Check the given project has the given bucket; if not, create a new bucket.
	var attrs *storage.BucketAttrs
	iter := cli.Buckets(ctx, cfg.Project)
	for {
		attrs, err = iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return
		}
		if attrs.Name == cfg.Bucket {
			return
		}
	}

	err = cli.Bucket(cfg.Bucket).Create(ctx, cfg.Project, nil)
	return

}

// Upload a file to a location.
func (s *StorageService) Upload(ctx context.Context, filename string, in io.Reader) (uri string, err error) {

	obj := s.client.Bucket(s.Config.Bucket).Object(filename)
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
		return "", fmt.Errorf("Faild to upload object %v", filename)
	}

	u := url.URL{
		Scheme: "gs",
		Host:   s.Config.Bucket,
		Path:   filepath.Join(StoragePrefix, filename),
	}
	uri = u.String()
	return

}

// Download downloads a file and write it to a given writer.
func (s *StorageService) Download(ctx context.Context, filename string, out io.Writer) (err error) {

	obj := s.client.Bucket(s.Config.Bucket).Object(filepath.Join(StoragePrefix, filename))
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
	return

}

// GetFileInfo returns a file status of an object.
func (s *StorageService) GetFileInfo(ctx context.Context, filename string) (info *cloud.FileInfo, err error) {

	attrs, err := s.client.Bucket(s.Config.Bucket).Object(filepath.Join(StoragePrefix, filename)).Attrs(ctx)
	if err != nil {
		return
	}

	info = &cloud.FileInfo{
		Name:        filepath.Base(attrs.Name),
		Path:        attrs.Name,
		TimeCreated: attrs.Created,
		Size:        attrs.Size,
	}
	return

}

// List searches items, i.e. files and folders, matching a given prefix.
// Found items will be passed to a given handler item by item.
// If the handler returns a non nil value, listing up will be canceled.
// In that case, this function will also return the given value.
func (s *StorageService) List(ctx context.Context, prefix string, handler cloud.FileInfoHandler) (err error) {

	iter := s.client.Bucket(s.Config.Bucket).Objects(ctx, &storage.Query{
		Prefix: filepath.Join(StoragePrefix, prefix),
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
		err = handler(&cloud.FileInfo{
			Name:        base,
			Path:        attrs.Name,
			TimeCreated: attrs.Created,
			Size:        attrs.Size,
		})
		if err != nil {
			return err
		}
	}

	return

}

// Delete deletes a given file.
func (s *StorageService) Delete(ctx context.Context, filename string) (err error) {
	return s.client.
		Bucket(s.Config.Bucket).
		Object(filepath.Join(StoragePrefix, filename)).
		Delete(ctx)
}

// Close the connection to a GCE's storage server.
func (s *StorageService) Close() error {
	return s.client.Close()
}
