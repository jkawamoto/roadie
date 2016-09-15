//
// command/util/gcs.go
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
	"net/http"
	"net/url"

	"github.com/jkawamoto/roadie/config"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/storage/v1"
)

const gcsScope = storage.DevstorageFullControlScope

// Storage object.
type Storage struct {
	BucketName string
	project    string
	ctx        context.Context
	client     *http.Client
	service    *storage.Service
}

// NewStorage creates a new storage accessor to a bucket name under the given contest.
// The context must have a config.
// If the given bucket does not exsits, it will be created.
func NewStorage(ctx context.Context) (s *Storage, err error) {

	cfg, ok := config.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Context dosen't have a config: %s", ctx)
	}

	// Create a client.
	client, err := google.DefaultClient(ctx, gcsScope)
	if err != nil {
		return
	}
	// Create a servicer.
	service, err := storage.New(client)
	if err != nil {
		return
	}

	return &Storage{
		BucketName: cfg.Gcp.Bucket,
		project:    cfg.Gcp.Project,
		ctx:        ctx,
		client:     client,
		service:    service,
	}, nil

}

// CreateIfNotExists creates the bucket if not exists.
func (s *Storage) CreateIfNotExists() error {

	var err error
	if _, exist := s.service.Buckets.Get(s.BucketName).Do(); exist != nil {
		_, err = s.service.Buckets.Insert(s.project, &storage.Bucket{Name: s.BucketName}).Do()
	}
	return err

}

// Upload a file to a location.
func (s *Storage) Upload(in io.Reader, location *url.URL) error {

	object := &storage.Object{Name: location.Path[1:]}
	if _, err := s.service.Objects.Insert(s.BucketName, object).Media(in).Do(); err != nil {
		return err
	}
	return nil

}

// Download downloads a file and write it to a given writer.
func (s *Storage) Download(filename string, out io.Writer) (err error) {

	res, err := s.service.Objects.Get(s.BucketName, filename).Download()
	if err != nil {
		return
	}
	defer res.Body.Close()

	reader := bufio.NewReader(res.Body)
	_, err = reader.WriteTo(out)
	return

}

// Status returns a file status of an object.
func (s *Storage) Status(filename string) (*FileInfo, error) {

	res, err := s.service.Objects.Get(s.BucketName, filename).Do()
	if err != nil {
		return nil, err
	}
	return NewFileInfo(res), nil

}

// List searches items, i.e. files and folders, matching a given prefix.
// Found items will be passed to a given handler item by item.
// If the handler returns a non nil value, listing up will be canceled.
// In that case, this function will also return the given value.
func (s *Storage) List(prefix string, handler func(*FileInfo) error) error {

	var token string
	for {

		res, err := s.service.Objects.List(s.BucketName).Prefix(prefix).PageToken(token).Do()
		if err != nil {
			return err
		}

		for _, item := range res.Items {

			select {
			case <-s.ctx.Done():
				return s.ctx.Err()
			default:
				if err := handler(NewFileInfo(item)); err != nil {
					return err
				}
			}

		}

		if res.NextPageToken == "" {
			return nil
		}
		token = res.NextPageToken

	}

}

// Delete deletes a given file.
func (s *Storage) Delete(name string) error {

	return s.service.Objects.Delete(s.BucketName, name).Do()

}
