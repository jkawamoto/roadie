//
// command/util/cloudstorage.go
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

// newService creates a new service object.
func (s *CloudStorageService) newService() (service *storage.Service, err error) {

	var client *http.Client
	// Create a client.
	client, err = google.DefaultClient(s.ctx, gcsScope)
	if err != nil {
		return
	}
	// Create a servicer.
	return storage.New(client)

}

// CreateIfNotExists creates the bucket if not exists.
func (s *CloudStorageService) CreateIfNotExists() (err error) {

	service, err := s.newService()
	if err != nil {
		return
	}

	cfg, ok := config.FromContext(s.ctx)
	if !ok {
		return fmt.Errorf("Context dosen't have a config: %s", s.ctx)
	}

	if _, exist := service.Buckets.Get(cfg.Bucket).Do(); exist != nil {
		_, err = service.Buckets.Insert(cfg.Project, &storage.Bucket{Name: cfg.Bucket}).Do()
	}
	return

}

// Upload a file to a location.
func (s *CloudStorageService) Upload(in io.Reader, location *url.URL) (err error) {

	service, err := s.newService()
	if err != nil {
		return
	}

	cfg, ok := config.FromContext(s.ctx)
	if !ok {
		return fmt.Errorf("Context dosen't have a config: %s", s.ctx)
	}

	object := &storage.Object{Name: location.Path[1:]}
	_, err = service.Objects.Insert(cfg.Bucket, object).Media(in).Do()
	return

}

// Download downloads a file and write it to a given writer.
func (s *CloudStorageService) Download(filename string, out io.Writer) (err error) {

	service, err := s.newService()
	if err != nil {
		return
	}

	cfg, ok := config.FromContext(s.ctx)
	if !ok {
		return fmt.Errorf("Context dosen't have a config: %s", s.ctx)
	}

	res, err := service.Objects.Get(cfg.Bucket, filename).Download()
	if err != nil {
		return
	}
	defer res.Body.Close()

	reader := bufio.NewReader(res.Body)
	_, err = reader.WriteTo(out)
	return

}

// Status returns a file status of an object.
func (s *CloudStorageService) Status(filename string) (info *FileInfo, err error) {

	service, err := s.newService()
	if err != nil {
		return
	}

	cfg, ok := config.FromContext(s.ctx)
	if !ok {
		return nil, fmt.Errorf("Context dosen't have a config: %s", s.ctx)
	}

	res, err := service.Objects.Get(cfg.Bucket, filename).Do()
	if err != nil {
		return
	}
	info = NewFileInfo(res)
	return

}

// List searches items, i.e. files and folders, matching a given prefix.
// Found items will be passed to a given handler item by item.
// If the handler returns a non nil value, listing up will be canceled.
// In that case, this function will also return the given value.
func (s *CloudStorageService) List(prefix string, handler FileInfoHandler) error {

	service, err := s.newService()
	if err != nil {
		return err
	}

	cfg, ok := config.FromContext(s.ctx)
	if !ok {
		return fmt.Errorf("Context dosen't have a config: %s", s.ctx)
	}

	var token string
	for {

		res, err := service.Objects.List(cfg.Bucket).Prefix(prefix).PageToken(token).Do()
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
func (s *CloudStorageService) Delete(name string) (err error) {

	service, err := s.newService()
	if err != nil {
		return
	}

	cfg, ok := config.FromContext(s.ctx)
	if !ok {
		return fmt.Errorf("Context dosen't have a config: %s", s.ctx)
	}

	return service.Objects.Delete(cfg.Bucket, name).Do()

}
