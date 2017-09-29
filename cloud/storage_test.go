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

package cloud_test

// TODO: Add tests for xz compressed files.

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/cloud/mock"
)

// uploadDummyFiles uploads a given list of dummy files to a given storage service.
func uploadDummyFiles(ctx context.Context, m cloud.StorageManager, prefix string, files []string) (err error) {

	var loc *url.URL
	for _, f := range files {
		loc, err = url.Parse(prefix + f)
		if err != nil {
			return
		}
		err = m.Upload(ctx, loc, strings.NewReader(f))
		if err != nil {
			return
		}
	}
	return

}

// Test uploading a file.
func TestUploadFile(t *testing.T) {

	var err error
	filename := "storage.go"
	loc, err := url.Parse("roadie://cloud_test/" + filename)
	if err != nil {
		t.Fatalf("cannot parse a URL: %v", err)
	}

	ctx := context.Background()
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	err = s.UploadFile(ctx, loc, filename)
	if err != nil {
		t.Fatalf("cannot upload: %v", err)
	}

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("cannot read the target file: %v", err)
	}

	out := bytes.Buffer{}
	err = service.Download(ctx, loc, &out)
	if err != nil {
		t.Fatalf("Download returns an error: %v", err)
	}

	if out.String() != string(body) {
		t.Errorf("Uploaded file = %v, want %v", out.String(), string(body))
	}

}

// Test uploading a non-existing file.
func TestUploadInvalidFile(t *testing.T) {

	var err error
	filename := "_storage.go" // Non-existing file.
	loc, err := url.Parse("roadie://cloud_test/" + filename)
	if err != nil {
		t.Fatalf("cannot parse a URL: %v", err)
	}

	ctx := context.Background()
	s := cloud.NewStorage(mock.NewStorageManager(), ioutil.Discard)

	err = s.UploadFile(ctx, loc, filename)
	if _, ok := err.(*os.PathError); !ok {
		t.Errorf("uploading non-existing file but no PathError: %v", err)
	}

}

func TestUploadWithServiceFailuer(t *testing.T) {

	var err error
	filename := "storage.go"
	loc, err := url.Parse("roadie://cloud_test/" + filename)
	if err != nil {
		t.Fatalf("cannot parse a URL: %v", err)
	}

	ctx := context.Background()
	service := mock.NewStorageManager()
	service.Failure = true
	s := cloud.NewStorage(service, ioutil.Discard)

	err = s.UploadFile(ctx, loc, filename)
	if err == nil {
		t.Error("uploading to a out-of-service servicer but no error occurred")
	}

}

// Test uploading a file with a canceled context.
func TestCancelUpload(t *testing.T) {

	var err error
	filename := "storage.go"
	loc, err := url.Parse("roadie://cloud_test/" + filename)
	if err != nil {
		t.Fatalf("cannot parse a URL: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	s := cloud.NewStorage(mock.NewStorageManager(), ioutil.Discard)
	cancel()

	err = s.UploadFile(ctx, loc, filename)
	if err == nil {
		t.Error("uploading with a canceled context but no error occurred")
	}

}

// Test lists up files.
func TestListupFiles(t *testing.T) {

	var err error
	ctx := context.Background()
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	prefix := "roadie://cloud_test/"
	files := []string{"test1/a", "test1/b", "test2/c", "test2/d"}
	err = uploadDummyFiles(ctx, service, prefix, files)
	if err != nil {
		t.Fatalf("cannot upload dummy files: %v", err)
	}

	query, err := url.Parse(prefix + "test1")
	if err != nil {
		t.Fatalf("cannot parse a URL: %v", err)
	}
	err = s.ListupFiles(ctx, query, func(info *cloud.FileInfo) error {
		if info.Name == "" {
			return nil
		}
		if !strings.HasPrefix(info.URL.String(), prefix) {
			t.Errorf("a wrong file is listed up: %v", info.URL)
		}
		return nil
	})
	if err != nil {
		t.Errorf("ListupFiles returns an error: %v", err)
	}

	// Test with Failure is true.
	service.Failure = true
	err = s.ListupFiles(ctx, query, func(info *cloud.FileInfo) error {
		return nil
	})
	if err == nil {
		t.Error("List up files in a out-of-service servicer but no error occurred")
	}

}

// Test downloads a file.
func TestDownloadFiles(t *testing.T) {

	var err error
	ctx := context.Background()
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	prefix := "roadie://cloud_test/"
	files := []string{"test1/aaa.log", "test1/bab.log", "test1/caa.txt", "test2/ddd.log"}
	err = uploadDummyFiles(ctx, service, prefix, files)
	if err != nil {
		t.Fatalf("cannot upload dummy files: %v", err)
	}

	loc, err := url.Parse(prefix + "test1")
	if err != nil {
		t.Fatalf("cannot parse a URL: %v", err)
	}

	temp1, err := ioutil.TempDir("", "TestDownloadFiles")
	if err != nil {
		t.Fatalf("cannot create a temporal directory: %v", err)
	}
	temp2 := filepath.Join(os.TempDir(), "TestDownloadFilesToNotExistingDir")

	for _, dir := range []string{temp1, temp2} {
		defer os.RemoveAll(dir)

		err = s.DownloadFiles(ctx, loc, dir, []string{"*.log"})
		if err != nil {
			t.Fatalf("DownloadFiiles returns an error: %v", err)
		}

		var matches []string
		matches, err = filepath.Glob(filepath.Join(dir, "*"))
		if err != nil {
			t.Fatalf("Glob returns an error: %v", err)
		}

		if len(matches) != 2 {
			t.Errorf("the number of downloaded files is %v, want 2", len(matches))
		}

		var data []byte
		for _, f := range matches {
			if !strings.HasSuffix(f, ".log") {
				t.Errorf("downloaded file doesn't match the query: %v", f)
			}
			data, err = ioutil.ReadFile(f)
			if err != nil {
				t.Fatalf("cannot read a downloaded file: %v", err)
			}
			body := string(data)
			if body != files[0] && body != files[1] {
				t.Errorf("downloaded file body is %v, want %v or %v", body, files[0], files[1])
			}
		}

	}

	// Failure is true.
	service.Failure = true
	err = s.DownloadFiles(ctx, loc, temp1, []string{"*.log"})
	if err == nil {
		t.Error("Download files from a out-of-service servicer but no error occurred")
	}
	service.Failure = false

	// With a canceled context.
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = s.DownloadFiles(ctx, loc, temp1, []string{"*.log"})
	if err == nil {
		t.Error("Download files is canceled but no error occurred")
	}

}

// Test deleting a file.
func TestDeleteFiles(t *testing.T) {

	var err error
	ctx := context.Background()
	service := mock.NewStorageManager()
	s := cloud.NewStorage(service, ioutil.Discard)

	prefix := "roadie://cloud_test/"
	files := []string{"test1/aaa.log", "test1/bab.log", "test1/caa.txt", "test2/ddd.log"}
	err = uploadDummyFiles(ctx, service, prefix, files)
	if err != nil {
		t.Fatalf("cannot upload dummy files: %v", err)
	}

	loc, err := url.Parse(prefix + "test1")
	if err != nil {
		t.Fatalf("cannot parse a URL: %v", err)
	}

	err = s.DeleteFiles(ctx, loc, []string{"*.log"})
	if err != nil {
		t.Fatalf("DownloadFiiles returns an error: %v", err)
	}

	err = s.ListupFiles(ctx, loc, func(info *cloud.FileInfo) error {
		if info.Name == "" {
			return nil
		}
		if info.URL.String() != loc.String()+"/caa.txt" {
			t.Errorf("a wrong file is still in the storage %v", info.URL)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("ListupFiles returns an error: %v", err)
	}

	// With out-of-service server.
	service.Failure = true
	err = s.DeleteFiles(ctx, loc, []string{"*.log"})
	if err == nil {
		t.Error("Delete files from a out-of-service servicer but no error occurred")
	}
	service.Failure = false

	// With a canceled context.
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	err = s.DeleteFiles(ctx, loc, []string{"*.log"})
	if err == nil {
		t.Error("Delete files has been canceled but no error occurred")
	}

}

func TestPrintFileBody(t *testing.T) {

	var err error
	ctx := context.Background()
	m := mock.NewStorageManager()
	s := cloud.NewStorage(m, ioutil.Discard)

	prefix := "roadie://cloud_test/"
	files := []string{"test1/aaa.log", "test1/bab.log", "test1/caa.txt", "test2/ddd.log"}
	err = uploadDummyFiles(ctx, m, prefix, files)
	if err != nil {
		t.Fatalf("cannot upload dummy files: %v", err)
	}

	loc, err := url.Parse(prefix + "test1")
	if err != nil {
		t.Fatalf("cannot parse a URL: %v", err)
	}

	var output bytes.Buffer
	err = s.PrintFileBody(ctx, loc, "aaa.log", &output, false)
	if err != nil {
		t.Fatalf("PrintFileBody returns an error: %v", err)
	}
	if output.String() != "test1/aaa.log" {
		t.Errorf("Printed file body is %v, want test1/aaa.log", output)
	}

	output.Reset()
	err = s.PrintFileBody(ctx, loc, "aaa.log", &output, true)
	if err != nil {
		t.Fatalf("PrintFileBody returns an error: %v", err)
	}
	if !strings.HasSuffix(output.String(), "aaa.log") {
		t.Errorf("Printed file body isn't correct: %v", output)
	}

	// With out-of-service server.
	m.Failure = true
	err = s.PrintFileBody(ctx, loc, "aaa.log", &output, false)
	if err == nil {
		t.Error("Request sent to a out-of-service servicer but no error occurred")
	}
	m.Failure = false

	// With a canceled context.
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	err = s.PrintFileBody(ctx, loc, "aaa.log", &output, false)
	if err == nil {
		t.Error("Print file body has been canceled but no error occurred")
	}

}
