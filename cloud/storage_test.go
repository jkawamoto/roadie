//
// cloud/storage_test.go
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
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// ErrServiceFailure is an error used in tests.
var ErrServiceFailure = fmt.Errorf("this service is out of order")

type MockStorageService struct {
	// Represent a key-value storage.
	storage map[string]string
	// If true, all method returns an error.
	failure bool
}

func NewMockStorageService() *MockStorageService {
	return &MockStorageService{
		storage: make(map[string]string),
	}
}

// Upload is a mock function to check uploaded file is correct or not.
func (s *MockStorageService) Upload(ctx context.Context, container, filename string, in io.Reader) (uri string, err error) {

	if s.failure {
		err = ErrServiceFailure
		return
	}

	body, err := ioutil.ReadAll(in)
	if err != nil {
		return
	}
	s.storage[filename] = string(body)
	return

}

func (s *MockStorageService) Download(ctx context.Context, container, filename string, out io.Writer) error {

	if s.failure {
		return ErrServiceFailure
	}

	body, ok := s.storage[filename]
	if !ok {
		return fmt.Errorf("File %s is not found", filename)
	}
	_, err := out.Write([]byte(body))
	return err

}

func (s *MockStorageService) GetFileInfo(ctx context.Context, container, filename string) (*FileInfo, error) {
	return nil, nil
}

// List is a mock function of List.
func (s *MockStorageService) List(ctx context.Context, container, prefix string, handler FileInfoHandler) error {

	// Represent a directory.
	err := handler(&FileInfo{
		Name:        "",
		Path:        prefix,
		TimeCreated: time.Now(),
		Size:        0,
	})
	if err != nil {
		return err
	}

	for filename, body := range s.storage {

		if strings.HasPrefix(filename, prefix) {

			err := handler(&FileInfo{
				Name:        filepath.Base(filename),
				Path:        filename,
				TimeCreated: time.Now(),
				Size:        int64(len(body)),
			})
			if err != nil {
				return err
			}

			if s.failure {
				return ErrServiceFailure
			}

		}
	}

	return nil
}

func (s *MockStorageService) Delete(ctx context.Context, container, filename string) error {

	if s.failure {
		return ErrServiceFailure
	}

	if _, ok := s.storage[filename]; !ok {
		return fmt.Errorf("Given file is not found: %v", filename)
	}

	delete(s.storage, filename)
	return nil
}

// Test uploading a file.
func TestUploadFile(t *testing.T) {

	container := "test"
	filename := "storage_test.go" // This file.

	ctx := context.Background()
	service := NewMockStorageService()
	s := NewStorage(service, ioutil.Discard)

	_, err := s.UploadFile(ctx, container, filename, filename)
	if err != nil {
		t.Error(err.Error())
	}

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}

	if elem, ok := service.storage[filename]; !ok {
		t.Error("Cannot find the uploaded file")
	} else if elem != string(body) {
		t.Error("Uploaded file don't match the expected file")
	}

}

// Test uploading a file without specifying name.
func TestUploadFileWithoutName(t *testing.T) {

	container := "test"
	filename := "storage_test.go" // This file.

	ctx := context.Background()
	service := NewMockStorageService()
	s := NewStorage(service, ioutil.Discard)

	// Omit to give a file name to be used to store the file.
	_, err := s.UploadFile(ctx, container, "", filename)
	if err != nil {
		t.Error(err.Error())
	}

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}

	if elem, ok := service.storage[filename]; !ok {
		t.Error("Cannot find the uploaded file")
	} else if elem != string(body) {
		t.Error("Uploaded file don't match the expected file")
	}

}

// Test uploading a non-existing file.
func TestUploadInvalidFile(t *testing.T) {

	prefix := "test"
	filename := "_storage_test.go" // Non-existing file.

	ctx := context.Background()
	s := NewStorage(NewMockStorageService(), ioutil.Discard)

	_, err := s.UploadFile(ctx, prefix, "", filename)
	if err == nil {
		t.Error("Uploading non-existing file but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

func TestUploadWithServiceFailuer(t *testing.T) {

	prefix := "test"
	filename := "storage_test.go" // This file.

	ctx := context.Background()
	service := NewMockStorageService()
	service.failure = true

	s := NewStorage(service, ioutil.Discard)

	_, err := s.UploadFile(ctx, prefix, "", filename)
	if err == nil {
		t.Error("Uploading to a out-of-service servicer but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

// Test uploading a file with a canceled context.
func TestCancelUpload(t *testing.T) {

	prefix := "test"
	filename := "storage_test.go" // This file.

	ctx, cancel := context.WithCancel(context.Background())
	s := NewStorage(NewMockStorageService(), ioutil.Discard)

	cancel()

	_, err := s.UploadFile(ctx, prefix, "", filename)
	if err == nil {
		t.Error("Uploading with a canceled context but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

// Test lists up files.
func TestListupFiles(t *testing.T) {

	ctx := context.Background()
	service := NewMockStorageService()
	s := NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.storage["test1/a"] = "a"
	service.storage["test1/b"] = "b"
	service.storage["test2/c"] = "c"
	service.storage["test2/d"] = "d"

	err := s.ListupFiles(ctx, "", prefix, func(info *FileInfo) error {
		if info.Name == "" {
			return nil
		}
		if !strings.HasPrefix(info.Path, prefix) {
			t.Error("List up a wrong file:", info.Path)
		}
		return nil
	})
	if err != nil {
		t.Error(err.Error())
	}

}

func TestListupFilesWithServiceFailuer(t *testing.T) {

	ctx := context.Background()
	service := NewMockStorageService()
	service.failure = true

	service.storage["test1/a"] = "a"
	service.storage["test1/b"] = "b"
	service.storage["test2/c"] = "c"
	service.storage["test2/d"] = "d"

	s := NewStorage(service, ioutil.Discard)

	err := s.ListupFiles(ctx, "", "test1", func(info *FileInfo) error {
		return nil
	})
	if err == nil {
		t.Error("List up files in a out-of-service servicer but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

// Test downloads a file.
func TestDownloadFiles(t *testing.T) {

	var err error
	ctx := context.Background()
	service := NewMockStorageService()
	s := NewStorage(service, ioutil.Discard)

	temp, err := ioutil.TempDir("", "TestDownloadFiles")
	if err != nil {
		t.Fatal("Cannot create a temporal directory:", err.Error())
	}
	defer os.RemoveAll(temp)

	prefix := "test1"
	service.storage["test1/aaa.log"] = "a"
	service.storage["test1/bab.log"] = "b"
	service.storage["test1/caa.txt"] = "c"
	service.storage["test2/ddd.log"] = "d"

	if err = s.DownloadFiles(ctx, "", prefix, temp, []string{"*.log"}); err != nil {
		t.Fatal("DownloadFiiles returns an error:", err.Error())
	}

	matches, err := filepath.Glob(filepath.Join(temp, "*"))
	if err != nil {
		t.Error(err.Error())
	}

	if len(matches) != 2 {
		t.Error("The number of downloaded files is not correct.")
	}

	var data []byte
	for _, f := range matches {
		if !strings.HasSuffix(f, ".log") {
			t.Error("Downloaded file is not matched to the query:", f)
		}
		data, err = ioutil.ReadFile(f)
		if err != nil {
			t.Error(err.Error())
		}
		body := string(data)
		if body != "a" && body != "b" {
			t.Error("Downloaded file body is not correct:", body)
		}
	}

}

func TestDownloadFilesToNotExistingDir(t *testing.T) {

	var err error
	ctx := context.Background()
	service := NewMockStorageService()
	s := NewStorage(service, ioutil.Discard)

	temp := filepath.Join(os.TempDir(), "TestDownloadFilesToNotExistingDir")
	defer os.RemoveAll(temp)

	prefix := "test1"
	service.storage["test1/aaa.log"] = "a"
	service.storage["test1/bab.log"] = "b"
	service.storage["test1/caa.txt"] = "c"
	service.storage["test2/ddd.log"] = "d"

	if err = s.DownloadFiles(ctx, "", prefix, temp, []string{"*.log"}); err != nil {
		t.Fatal("DownloadFiiles returns an error:", err.Error())
	}

	matches, err := filepath.Glob(filepath.Join(temp, "*"))
	if err != nil {
		t.Error(err.Error())
	}

	if len(matches) != 2 {
		t.Error("The number of downloaded files is not correct.")
	}

	var data []byte
	for _, f := range matches {
		if !strings.HasSuffix(f, ".log") {
			t.Error("Downloaded file is not matched to the query:", f)
		}
		data, err = ioutil.ReadFile(f)
		if err != nil {
			t.Error(err.Error())
		}
		body := string(data)
		if body != "a" && body != "b" {
			t.Error("Downloaded file body is not correct:", body)
		}
	}

}

func TestDownloadFilesWithServicerFailure(t *testing.T) {

	var err error
	ctx := context.Background()
	service := NewMockStorageService()
	service.failure = true

	s := NewStorage(service, ioutil.Discard)

	temp, err := ioutil.TempDir("", "TestDownloadFiles")
	if err != nil {
		t.Fatal("Cannot create a temporal directory:", err.Error())
	}
	defer os.RemoveAll(temp)

	prefix := "test1"
	service.storage["test1/aaa.log"] = "a"
	service.storage["test1/bab.log"] = "b"
	service.storage["test1/caa.txt"] = "c"
	service.storage["test2/ddd.log"] = "d"

	if err = s.DownloadFiles(ctx, "", prefix, temp, []string{"*.log"}); err == nil {
		t.Error("Download files from a out-of-service servicer but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

// Test cancel works in downloading files.
func TestCancelDownloadFiles(t *testing.T) {

	var err error
	ctx, cancel := context.WithCancel(context.Background())
	service := NewMockStorageService()
	s := NewStorage(service, ioutil.Discard)

	temp, err := ioutil.TempDir("", "TestDownloadFiles")
	if err != nil {
		t.Fatal("Cannot create a temporal directory:", err.Error())
	}
	defer os.RemoveAll(temp)

	prefix := "test1"
	service.storage["test1/aaa.log"] = "a"
	service.storage["test1/bab.log"] = "b"
	service.storage["test1/caa.txt"] = "c"
	service.storage["test2/ddd.log"] = "d"

	cancel()
	if err = s.DownloadFiles(ctx, "", prefix, temp, []string{"*.log"}); err == nil {
		t.Error("Download files is canceled but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

// Test deleting a file.
func TestDeleteFiles(t *testing.T) {

	ctx := context.Background()
	service := NewMockStorageService()
	s := NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.storage["test1/aaa.log"] = "a"
	service.storage["test1/bab.log"] = "b"
	service.storage["test1/caa.txt"] = "c"
	service.storage["test2/ddd.log"] = "d"

	if err := s.DeleteFiles(ctx, "", prefix, []string{"*.log"}); err != nil {
		t.Error("DownloadFiiles returns an error:", err.Error())
	}

	if len(service.storage) != 2 {
		t.Error("The number of eleted files are wrong")
	}
	for filename := range service.storage {
		if filename != "test1/caa.txt" && filename != "test2/ddd.log" {
			t.Error("DeleteFiles has deleted wrong files")
		}
	}

}

func TestDeleteFilesWithServicerFailure(t *testing.T) {

	var err error
	ctx := context.Background()
	service := NewMockStorageService()
	service.failure = true

	s := NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.storage["test1/aaa.log"] = "a"
	service.storage["test1/bab.log"] = "b"
	service.storage["test1/caa.txt"] = "c"
	service.storage["test2/ddd.log"] = "d"

	err = s.DeleteFiles(ctx, "", prefix, []string{"*.log"})
	if err == nil {
		t.Error("Delete files from a out-of-service servicer but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

// Test deleting a file.
func TestCancelDeleteFiles(t *testing.T) {

	var err error
	ctx, cancel := context.WithCancel(context.Background())
	service := NewMockStorageService()
	s := NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.storage["test1/aaa.log"] = "a"
	service.storage["test1/bab.log"] = "b"
	service.storage["test1/caa.txt"] = "c"
	service.storage["test2/ddd.log"] = "d"

	cancel()

	err = s.DeleteFiles(ctx, "", prefix, []string{"*.log"})
	if err == nil {
		t.Error("Delete files has been canceled but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

func TestPrintFileBody(t *testing.T) {

	var err error
	ctx := context.Background()
	service := NewMockStorageService()
	s := NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.storage["test1/aaa.log"] = "a"
	service.storage["test1/bab.log"] = "b"
	service.storage["test1/caa.txt"] = "c"
	service.storage["test2/ddd.log"] = "d"

	output := &bytes.Buffer{}
	if err = s.PrintFileBody(ctx, "", prefix, "aaa.log", output, false); err != nil {
		t.Error("PrintFileBody returns an error:", err.Error())
	}
	if output.String() != "a" {
		t.Error("Printed file body isn't correct:", output.String())
	}

	output.Reset()
	if err = s.PrintFileBody(ctx, "", prefix, "aaa.log", output, true); err != nil {
		t.Error("PrintFileBody returns an error:", err.Error())
	}
	if !strings.HasSuffix(output.String(), "a") {
		t.Error("Printed file body isn't correct:", output.String())
	}

}

func TestPrintFileBodyWithServicerFailure(t *testing.T) {

	var err error
	ctx := context.Background()
	service := NewMockStorageService()
	service.failure = true

	s := NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.storage["test1/aaa.log"] = "a"
	service.storage["test1/bab.log"] = "b"
	service.storage["test1/caa.txt"] = "c"
	service.storage["test2/ddd.log"] = "d"

	output := &bytes.Buffer{}
	err = s.PrintFileBody(ctx, "", prefix, "aaa.log", output, false)
	if err == nil {
		t.Error("Request sent to a out-of-service servicer but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

func TestCancelPrintFileBody(t *testing.T) {

	var err error
	ctx, cancel := context.WithCancel(context.Background())
	service := NewMockStorageService()
	s := NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.storage["test1/aaa.log"] = "a"
	service.storage["test1/bab.log"] = "b"
	service.storage["test1/caa.txt"] = "c"
	service.storage["test2/ddd.log"] = "d"

	cancel()

	output := &bytes.Buffer{}
	err = s.PrintFileBody(ctx, "", prefix, "aaa.log", output, false)
	if err == nil {
		t.Error("Print file body has been canceled but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}
