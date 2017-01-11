//
// cloud/cloudstorage.go
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
	"context"
	"fmt"
	"io"
	"net/url"
	"path/filepath"

	"google.golang.org/api/iterator"

	"github.com/jkawamoto/roadie/config"

	"cloud.google.com/go/storage"
)

// CloudStorageService object.
type CloudStorageService struct {
	// Context must have Config.
	ctx context.Context
}

// NewCloudStorageService creates a new storage accessor to a bucket name under the given contest.
// The context must have a config.
func NewCloudStorageService(ctx context.Context) *CloudStorageService {

	return &CloudStorageService{
		ctx: ctx,
	}

}

// CreateIfNotExists creates the bucket if not exists.
func (s *CloudStorageService) CreateIfNotExists() (err error) {

	cli, err := storage.NewClient(s.ctx)
	if err != nil {
		return
	}
	defer cli.Close()

	cfg, err := config.FromContext(s.ctx)
	if err != nil {
		return
	}

	iter := cli.Buckets(s.ctx, cfg.Project)
	for {
		e, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return err
		}
		if e.Name == cfg.Bucket {
			return nil
		}
	}

	return cli.Bucket(cfg.Bucket).Create(s.ctx, cfg.Project, nil)

}

// Upload a file to a location.
func (s *CloudStorageService) Upload(in io.Reader, location *url.URL) (err error) {

	cli, err := storage.NewClient(s.ctx)
	if err != nil {
		return
	}
	defer cli.Close()

	cfg, err := config.FromContext(s.ctx)
	if err != nil {
		return
	}

	obj := cli.Bucket(cfg.Bucket).Object(location.Path[1:])

	writer := obj.NewWriter(s.ctx)
	size, err := io.Copy(writer, in)
	writer.Close()
	if err != nil {
		return
	}

	info, err := obj.Attrs(s.ctx)
	if err != nil {
		return
	} else if info.Size != size {
		obj.Delete(s.ctx)
		return fmt.Errorf("Faild to upload object %v", location)
	}

	return

}

// Download downloads a file and write it to a given writer.
func (s *CloudStorageService) Download(filename string, out io.Writer) (err error) {

	cli, err := storage.NewClient(s.ctx)
	if err != nil {
		return
	}
	defer cli.Close()

	cfg, err := config.FromContext(s.ctx)
	if err != nil {
		return
	}

	obj := cli.Bucket(cfg.Bucket).Object(filename)

	info, err := obj.Attrs(s.ctx)
	if err != nil {
		return
	}

	reader, err := obj.NewReader(s.ctx)
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

// Status returns a file status of an object.
func (s *CloudStorageService) Status(filename string) (info *FileInfo, err error) {

	cli, err := storage.NewClient(s.ctx)
	if err != nil {
		return
	}
	defer cli.Close()

	cfg, err := config.FromContext(s.ctx)
	if err != nil {
		return
	}

	attrs, err := cli.Bucket(cfg.Bucket).Object(filename).Attrs(s.ctx)
	if err != nil {
		return
	}

	info = &FileInfo{
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
func (s *CloudStorageService) List(prefix string, handler FileInfoHandler) (err error) {

	cli, err := storage.NewClient(s.ctx)
	if err != nil {
		return
	}
	defer cli.Close()

	cfg, err := config.FromContext(s.ctx)
	if err != nil {
		return
	}

	iter := cli.Bucket(cfg.Bucket).Objects(s.ctx, &storage.Query{
		Prefix: prefix,
	})
	for {
		attrs, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return err
		}

		err = handler(&FileInfo{
			Name:        filepath.Base(attrs.Name),
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
func (s *CloudStorageService) Delete(name string) (err error) {

	cli, err := storage.NewClient(s.ctx)
	if err != nil {
		return
	}
	defer cli.Close()

	cfg, err := config.FromContext(s.ctx)
	if err != nil {
		return
	}

	return cli.Bucket(cfg.Bucket).Object(name).Delete(s.ctx)

}
