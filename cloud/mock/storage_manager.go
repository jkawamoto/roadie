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
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/jkawamoto/roadie/cloud"
)

// ErrServiceFailure is an error used in tests.
var ErrServiceFailure = fmt.Errorf("this service is out of order")

// StorageManager is a memory based mock storage manager.
type StorageManager struct {
	// Represent a key-value storage.
	Storage map[string]string
	// If true, all method returns an error.
	Failure bool
	// Lock.
	mutex *sync.Mutex
}

// NewStorageManager creates a new mock storage manager.
func NewStorageManager() *StorageManager {
	return &StorageManager{
		Storage: make(map[string]string),
		mutex:   new(sync.Mutex),
	}
}

// Upload is a mock function to check uploaded file is correct or not.
func (s *StorageManager) Upload(ctx context.Context, container, filename string, in io.Reader) (uri string, err error) {

	if s.Failure {
		err = ErrServiceFailure
		return
	}

	body, err := ioutil.ReadAll(in)
	if err != nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Storage[filename] = string(body)
	uri = fmt.Sprint("file://mock/", filepath.Join(container, filename))
	return

}

// Download writes a file in a container to a writer.
func (s *StorageManager) Download(ctx context.Context, container, filename string, out io.Writer) error {

	if s.Failure {
		return ErrServiceFailure
	}

	body, ok := s.Storage[filename]
	if !ok {
		return fmt.Errorf("File %s is not found", filename)
	}
	_, err := out.Write([]byte(body))
	return err

}

// GetFileInfo returns information of a file.
func (s *StorageManager) GetFileInfo(ctx context.Context, container, filename string) (*cloud.FileInfo, error) {
	return nil, nil
}

// List is a mock function of List.
func (s *StorageManager) List(ctx context.Context, container, prefix string, handler cloud.FileInfoHandler) error {

	// Represent a directory.
	err := handler(&cloud.FileInfo{
		Name:        "",
		Path:        prefix,
		TimeCreated: time.Now(),
		Size:        0,
	})
	if err != nil {
		return err
	}

	for filename, body := range s.Storage {

		if strings.HasPrefix(filename, prefix) {

			err := handler(&cloud.FileInfo{
				Name:        filepath.Base(filename),
				Path:        filename,
				TimeCreated: time.Now(),
				Size:        int64(len(body)),
			})
			if err != nil {
				return err
			}

			if s.Failure {
				return ErrServiceFailure
			}

		}
	}

	return nil
}

// Delete deletes a file.
func (s *StorageManager) Delete(ctx context.Context, container, filename string) error {

	if s.Failure {
		return ErrServiceFailure
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.Storage[filename]; !ok {
		return fmt.Errorf("Given file is not found: %v", filename)
	}

	delete(s.Storage, filename)
	return nil
}
