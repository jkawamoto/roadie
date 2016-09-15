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
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/storage/v1"
)

const gcsScope = storage.DevstorageFullControlScope

// Storage object.
type Storage struct {
	BucketName string
	project    string
	client     *http.Client
	service    *storage.Service
}

// FileInfo defines file information structure.
type FileInfo struct {
	Name        string
	Path        string
	TimeCreated time.Time
	Size        uint64
}

// NewStorage creates a new storage accessor to a given bucket name
// under the given contest. If the given bucket does not exsits, it will be created.
func NewStorage(ctx context.Context, project, bucket string) (s *Storage, err error) {

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

	s = &Storage{
		BucketName: bucket,
		project:    project,
		client:     client,
		service:    service,
	}

	// Check the given bucket exists.
	err = s.createIfNotExists()
	return

}

// createIfNotExists creates the bucket if not exists.
func (s *Storage) createIfNotExists() error {

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

// List is a goroutine to list up files in a bucket.
func (s *Storage) List(prefix string, resCh chan<- *FileInfo, errCh chan<- error) {

	token := ""
	for {

		res, err := s.service.Objects.List(s.BucketName).Prefix(prefix).PageToken(token).Do()
		if err != nil {
			errCh <- err
			return
		}

		for _, item := range res.Items {
			resCh <- NewFileInfo(item)
		}

		token = res.NextPageToken
		if token == "" {
			resCh <- nil
			return
		}

	}
}

// Delete deletes a given file.
func (s *Storage) Delete(name string) error {

	return s.service.Objects.Delete(s.BucketName, name).Do()

}

// NewFileInfo creates a file info from an object.
func NewFileInfo(f *storage.Object) *FileInfo {

	splitedName := strings.Split(f.Name, "/")
	t, _ := time.Parse("2006-01-02T15:04:05", strings.Split(f.TimeCreated, ".")[0])

	return &FileInfo{
		Name:        splitedName[len(splitedName)-1],
		Path:        f.Name,
		TimeCreated: t.In(time.Local),
		Size:        f.Size,
	}
}
