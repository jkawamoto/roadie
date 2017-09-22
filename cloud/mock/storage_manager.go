//
// cloud/mock/storage_manager.go
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

package mock

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/jkawamoto/roadie/cloud"
)

// ErrServiceFailure is an error used in tests.
var ErrServiceFailure = fmt.Errorf("this service is out of order")

// StorageManager is a memory based mock storage manager.
type StorageManager struct {
	// If true, all method returns an error.
	Failure bool
	// Represent a key-value storage.
	storage map[string]string
	// Lock.
	mutex *sync.Mutex
}

// NewStorageManager creates a new mock storage manager.
func NewStorageManager() *StorageManager {
	return &StorageManager{
		storage: make(map[string]string),
		mutex:   new(sync.Mutex),
	}
}

// Upload reads bytes from a given stream and stores it in the map storage
// with a given location as the key.
func (s *StorageManager) Upload(ctx context.Context, loc *url.URL, in io.Reader) (err error) {

	if s.Failure {
		return ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	body, err := ioutil.ReadAll(in)
	if err != nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.storage[loc.String()] = string(body)
	return

}

// Download writes bytes from the map storage with a given key loc to a given writer.
func (s *StorageManager) Download(ctx context.Context, loc *url.URL, out io.Writer) error {

	if s.Failure {
		return ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	body, ok := s.storage[loc.String()]
	if !ok {
		return fmt.Errorf("%v is not found", loc)
	}
	_, err := out.Write([]byte(body))
	return err

}

// GetFileInfo returns information of a file pointed by a given URL.
func (s *StorageManager) GetFileInfo(ctx context.Context, loc *url.URL) (*cloud.FileInfo, error) {

	if s.Failure {
		return nil, ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	v, ok := s.storage[loc.String()]
	if !ok {
		return nil, os.ErrNotExist
	}

	info := cloud.FileInfo{
		Name:        path.Base(loc.Path),
		URL:         loc,
		TimeCreated: time.Now(),
		Size:        int64(len(v)),
	}
	return &info, nil

}

// List lists up files in this storage and passes each information to a given handler.
func (s *StorageManager) List(ctx context.Context, loc *url.URL, handler cloud.FileInfoHandler) (err error) {

	if s.Failure {
		return ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	query := loc.String()
	for filename, body := range s.storage {

		if strings.HasPrefix(filename, query) {

			var u *url.URL
			u, err = url.Parse(filename)
			if err != nil {
				return
			}
			err = handler(&cloud.FileInfo{
				Name:        path.Base(filename),
				URL:         u,
				TimeCreated: time.Now(),
				Size:        int64(len(body)),
			})
			if err != nil {
				return
			}

			if s.Failure {
				return ErrServiceFailure
			}

		}
	}

	return nil
}

// Delete deletes a file pointed by a given URL.
func (s *StorageManager) Delete(ctx context.Context, loc *url.URL) error {

	if s.Failure {
		return ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.storage[loc.String()]; !ok {
		return os.ErrNotExist
	}

	delete(s.storage, loc.String())
	return nil
}
