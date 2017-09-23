//
// cloud/gcp/storage.go
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

package gcp

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"path"
	"strings"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"

	"cloud.google.com/go/storage"
)

// getObjectName converts a URL to an object name for Google Cloud Storage.
// Given URL must be in the following format;
// - roadie://category/path
// and the converted name will be in the following format;
// - .roadie/category/path
func getObjectName(loc *url.URL) (name string) {

	name = path.Join(StoragePrefix, loc.Hostname(), loc.Path)
	if strings.HasSuffix(loc.Path, "/") && !strings.HasSuffix(name, "/") {
		name += "/"
	}
	return

}

// StorageService implements cloud.StorageServicer interface for accessing GCP's
// cloud storage.
type StorageService struct {
	// Config is a reference for a configuration of GCP.
	Config *Config
	// Logger
	Logger *log.Logger
}

// NewStorageService creates a new storage accessor to a bucket specified in a
// given configuration.
func NewStorageService(ctx context.Context, cfg *Config, logger *log.Logger) (s *StorageService, err error) {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	var cli *storage.Client
	if cfg.Token == nil || cfg.Token.AccessToken == "" {
		// If any token is not given, use a normal client.
		cli, err = storage.NewClient(ctx)
	} else {
		c := NewAuthorizationConfig(0)
		cli, err = storage.NewClient(ctx, option.WithTokenSource(c.TokenSource(ctx, cfg.Token)))
	}
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

// newClient creates a new storage client.
func (s *StorageService) newClient(ctx context.Context) (*storage.Client, error) {

	cfg := NewAuthorizationConfig(0)
	return storage.NewClient(ctx, option.WithTokenSource(cfg.TokenSource(ctx, s.Config.Token)))

}

// Upload a file to a location.
func (s *StorageService) Upload(ctx context.Context, loc *url.URL, in io.Reader) (err error) {

	s.Logger.Println("Uploading a file to", loc)
	client, err := s.newClient(ctx)
	if err != nil {
		return
	}
	defer client.Close()

	name := getObjectName(loc)
	obj := client.Bucket(s.Config.Bucket).Object(name)

	// After the writer will be closed, actual uploading will be finished.
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
		return fmt.Errorf("Faild to upload object %v", name)
	}

	s.Logger.Println("Finished uploading a file to", loc)
	return

}

// Download downloads a file and write it to a given writer.
func (s *StorageService) Download(ctx context.Context, loc *url.URL, out io.Writer) (err error) {

	s.Logger.Println("Downloading a file from", loc)
	client, err := s.newClient(ctx)
	if err != nil {
		return
	}
	defer client.Close()

	name := getObjectName(loc)
	obj := client.Bucket(s.Config.Bucket).Object(name)
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
		return fmt.Errorf("Faild to download object %v", name)
	}

	s.Logger.Println("Finished downloading the file from", loc)
	return

}

// GetFileInfo returns a file status of an object.
func (s *StorageService) GetFileInfo(ctx context.Context, loc *url.URL) (info *cloud.FileInfo, err error) {

	s.Logger.Println("Retrieving information of a file in", loc)
	client, err := s.newClient(ctx)
	if err != nil {
		return
	}
	defer client.Close()

	name := getObjectName(loc)
	attrs, err := client.Bucket(s.Config.Bucket).Object(name).Attrs(ctx)
	if err != nil {
		return
	}

	info = &cloud.FileInfo{
		Name:        path.Base(attrs.Name),
		URL:         loc,
		TimeCreated: attrs.Created,
		Size:        attrs.Size,
	}

	s.Logger.Println("Finished retrieving the information of the file in", loc)
	return

}

// List searches items, i.e. files and folders, matching a given prefix.
// Found items will be passed to a given handler item by item.
// If the handler returns a non nil value, listing up will be canceled.
// In that case, this function will also return the given value.
func (s *StorageService) List(ctx context.Context, loc *url.URL, handler cloud.FileInfoHandler) (err error) {

	s.Logger.Printf(`Retrieving the list of files matching to prefix "%v"`, loc)
	client, err := s.newClient(ctx)
	if err != nil {
		return
	}
	defer client.Close()

	iter := client.Bucket(s.Config.Bucket).Objects(ctx, &storage.Query{
		Prefix:   getObjectName(loc),
		Versions: false,
	})

	var attrs *storage.ObjectAttrs
	var u *url.URL
	for {
		attrs, err = iter.Next()
		if err == iterator.Done {
			err = nil
			break
		} else if err != nil {
			return
		}

		u, err = url.Parse(script.RoadieSchemePrefix + strings.TrimPrefix(attrs.Name, StoragePrefix+"/"))
		if err != nil {
			return
		}
		err = handler(&cloud.FileInfo{
			Name:        path.Base(attrs.Name),
			URL:         u,
			TimeCreated: attrs.Created,
			Size:        attrs.Size,
		})
		if err != nil {
			return
		}
	}

	s.Logger.Println("Finished retrieving the file list")
	return

}

// Delete deletes a given file.
func (s *StorageService) Delete(ctx context.Context, loc *url.URL) (err error) {

	s.Logger.Println("Deleting a file in", loc)
	client, err := s.newClient(ctx)
	if err != nil {
		return
	}
	defer client.Close()

	name := getObjectName(loc)
	err = client.Bucket(s.Config.Bucket).Object(name).Delete(ctx)
	if err != nil {
		return
	}

	s.Logger.Println("Finished deleting the file in", loc)
	return

}
