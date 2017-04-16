//
// cloud/cloud_test.go
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

package cloud_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/cloud/mock"
)

// Test uploading a file.
func TestUploadFile(t *testing.T) {

	container := "test"
	filename := "cloud_test.go" // This file.

	ctx := context.Background()
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	_, err := s.UploadFile(ctx, container, filename, filename)
	if err != nil {
		t.Error(err.Error())
	}

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}

	if elem, ok := service.Storage[filename]; !ok {
		t.Error("Cannot find the uploaded file")
	} else if elem != string(body) {
		t.Error("Uploaded file don't match the expected file")
	}

}

// Test uploading a file without specifying name.
func TestUploadFileWithoutName(t *testing.T) {

	container := "test"
	filename := "cloud_test.go" // This file.

	ctx := context.Background()
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	// Omit to give a file name to be used to store the file.
	_, err := s.UploadFile(ctx, container, "", filename)
	if err != nil {
		t.Error(err.Error())
	}

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}

	if elem, ok := service.Storage[filename]; !ok {
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
	s := cloud.NewStorage(mock.NewStorageManager(), ioutil.Discard)

	_, err := s.UploadFile(ctx, prefix, "", filename)
	if err == nil {
		t.Error("Uploading non-existing file but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

func TestUploadWithServiceFailuer(t *testing.T) {

	prefix := "test"
	filename := "cloud_test.go" // This file.

	ctx := context.Background()
	service := mock.NewStorageManager()
	service.Failure = true

	s := cloud.NewStorage(service, ioutil.Discard)

	_, err := s.UploadFile(ctx, prefix, "", filename)
	if err == nil {
		t.Error("Uploading to a out-of-service servicer but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

// Test uploading a file with a canceled context.
func TestCancelUpload(t *testing.T) {

	prefix := "test"
	filename := "cloud_test.go" // This file.

	ctx, cancel := context.WithCancel(context.Background())
	s := cloud.NewStorage(mock.NewStorageManager(), ioutil.Discard)

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
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.Storage["test1/a"] = "a"
	service.Storage["test1/b"] = "b"
	service.Storage["test2/c"] = "c"
	service.Storage["test2/d"] = "d"

	err := s.ListupFiles(ctx, "", prefix, func(info *cloud.FileInfo) error {
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
	service := mock.NewStorageManager()
	service.Failure = true

	service.Storage["test1/a"] = "a"
	service.Storage["test1/b"] = "b"
	service.Storage["test2/c"] = "c"
	service.Storage["test2/d"] = "d"

	s := cloud.NewStorage(service, ioutil.Discard)

	err := s.ListupFiles(ctx, "", "test1", func(info *cloud.FileInfo) error {
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
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	temp, err := ioutil.TempDir("", "TestDownloadFiles")
	if err != nil {
		t.Fatal("Cannot create a temporal directory:", err.Error())
	}
	defer os.RemoveAll(temp)

	prefix := "test1"
	service.Storage["test1/aaa.log"] = "a"
	service.Storage["test1/bab.log"] = "b"
	service.Storage["test1/caa.txt"] = "c"
	service.Storage["test2/ddd.log"] = "d"

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
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	temp := filepath.Join(os.TempDir(), "TestDownloadFilesToNotExistingDir")
	defer os.RemoveAll(temp)

	prefix := "test1"
	service.Storage["test1/aaa.log"] = "a"
	service.Storage["test1/bab.log"] = "b"
	service.Storage["test1/caa.txt"] = "c"
	service.Storage["test2/ddd.log"] = "d"

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
	service := mock.NewStorageManager()
	service.Failure = true

	s := cloud.NewStorage(service, ioutil.Discard)

	temp, err := ioutil.TempDir("", "TestDownloadFiles")
	if err != nil {
		t.Fatal("Cannot create a temporal directory:", err.Error())
	}
	defer os.RemoveAll(temp)

	prefix := "test1"
	service.Storage["test1/aaa.log"] = "a"
	service.Storage["test1/bab.log"] = "b"
	service.Storage["test1/caa.txt"] = "c"
	service.Storage["test2/ddd.log"] = "d"

	if err = s.DownloadFiles(ctx, "", prefix, temp, []string{"*.log"}); err == nil {
		t.Error("Download files from a out-of-service servicer but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

// Test cancel works in downloading files.
func TestCancelDownloadFiles(t *testing.T) {

	var err error
	ctx, cancel := context.WithCancel(context.Background())
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	temp, err := ioutil.TempDir("", "TestDownloadFiles")
	if err != nil {
		t.Fatal("Cannot create a temporal directory:", err.Error())
	}
	defer os.RemoveAll(temp)

	prefix := "test1"
	service.Storage["test1/aaa.log"] = "a"
	service.Storage["test1/bab.log"] = "b"
	service.Storage["test1/caa.txt"] = "c"
	service.Storage["test2/ddd.log"] = "d"

	cancel()
	if err = s.DownloadFiles(ctx, "", prefix, temp, []string{"*.log"}); err == nil {
		t.Error("Download files is canceled but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}

// Test deleting a file.
func TestDeleteFiles(t *testing.T) {

	ctx := context.Background()
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.Storage["test1/aaa.log"] = "a"
	service.Storage["test1/bab.log"] = "b"
	service.Storage["test1/caa.txt"] = "c"
	service.Storage["test2/ddd.log"] = "d"

	if err := s.DeleteFiles(ctx, "", prefix, []string{"*.log"}); err != nil {
		t.Error("DownloadFiiles returns an error:", err.Error())
	}

	if len(service.Storage) != 2 {
		t.Error("The number of eleted files are wrong")
	}
	for filename := range service.Storage {
		if filename != "test1/caa.txt" && filename != "test2/ddd.log" {
			t.Error("DeleteFiles has deleted wrong files")
		}
	}

}

func TestDeleteFilesWithServicerFailure(t *testing.T) {

	var err error
	ctx := context.Background()
	service := mock.NewStorageManager()
	service.Failure = true

	s := cloud.NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.Storage["test1/aaa.log"] = "a"
	service.Storage["test1/bab.log"] = "b"
	service.Storage["test1/caa.txt"] = "c"
	service.Storage["test2/ddd.log"] = "d"

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
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.Storage["test1/aaa.log"] = "a"
	service.Storage["test1/bab.log"] = "b"
	service.Storage["test1/caa.txt"] = "c"
	service.Storage["test2/ddd.log"] = "d"

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
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.Storage["test1/aaa.log"] = "a"
	service.Storage["test1/bab.log"] = "b"
	service.Storage["test1/caa.txt"] = "c"
	service.Storage["test2/ddd.log"] = "d"

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
	service := mock.NewStorageManager()
	service.Failure = true

	s := cloud.NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.Storage["test1/aaa.log"] = "a"
	service.Storage["test1/bab.log"] = "b"
	service.Storage["test1/caa.txt"] = "c"
	service.Storage["test2/ddd.log"] = "d"

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
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	prefix := "test1"
	service.Storage["test1/aaa.log"] = "a"
	service.Storage["test1/bab.log"] = "b"
	service.Storage["test1/caa.txt"] = "c"
	service.Storage["test2/ddd.log"] = "d"

	cancel()

	output := &bytes.Buffer{}
	err = s.PrintFileBody(ctx, "", prefix, "aaa.log", output, false)
	if err == nil {
		t.Error("Print file body has been canceled but no error occurred")
	}
	t.Log("Received error message:", err.Error())

}
