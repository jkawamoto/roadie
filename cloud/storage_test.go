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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package cloud

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/config"
)

type mockStorageServicer struct {
	t        *testing.T
	bucket   string
	prefix   string
	filename string
	location string
	testList bool
}

func newMockStorageServicer(t *testing.T, cfg *config.Config) *mockStorageServicer {

	prefix := "sample/prefix"
	filename := "storage_test.go"

	if cfg == nil {
		cfg = &config.Config{}
	}

	return &mockStorageServicer{
		t:        t,
		bucket:   cfg.Bucket,
		prefix:   prefix,
		filename: filename,
		location: fmt.Sprintf("gs://%s/%s/%s", cfg.Bucket, prefix, filename),
	}

}

func (s *mockStorageServicer) CreateIfNotExists() error {
	return nil
}

// Upload is a mock function to check uploaded file is correct or not.
func (s *mockStorageServicer) Upload(in io.Reader, location *url.URL) error {

	obtain, err := ioutil.ReadAll(in)
	if err != nil {
		s.t.Error(err.Error())
		return err
	}

	expected, err := ioutil.ReadFile(s.filename)
	if err != nil {
		s.t.Error(err.Error())
		return err
	}

	if string(obtain) != string(expected) {
		s.t.Error("Received file body isn't correct:", string(obtain))
	}

	if location.String() != s.location {
		s.t.Error("Given location isn't correct:", location)
	}

	return nil
}

func (s *mockStorageServicer) Download(filename string, out io.Writer) error {

	if filename != s.filename {
		return fmt.Errorf("File %s is not found", filename)
	}

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	out.Write(body)
	return nil

}

func (s *mockStorageServicer) Status(filename string) (*FileInfo, error) {
	return nil, nil
}

// List is a mock function of List.
func (s *mockStorageServicer) List(prefix string, handler FileInfoHandler) error {

	// Actual file to be downloaded.
	info, err := os.Stat(s.filename)
	if err != nil {
		s.t.Fatal("Testing file doesn't exist:", s.filename, err.Error())
	}
	if err := handler(&FileInfo{
		Name: s.filename,
		Path: s.filename,
		Size: info.Size(),
	}); err != nil {
		return err
	}

	// Dummy file.
	if err := handler(&FileInfo{
		Name: fmt.Sprintf("%s-%d", prefix, 1),
		Path: s.filename,
		Size: info.Size(),
	}); err != nil {
		return err
	}

	if s.testList {
		// Dummy file for raising an error.
		if err := handler(&FileInfo{
			Name: fmt.Sprintf("%s-end", prefix),
		}); err == nil {
			s.t.Error("Handler doesn't stop listing up.")
		}
	}

	return nil
}

func (s *mockStorageServicer) Delete(name string) error {

	if name != s.filename {
		s.t.Error("Delete a wrong file:", name)
	}
	return nil
}

// TestPrepareBucket.
func TestPrepareBucket(t *testing.T) {

	s := &Storage{
		service: newMockStorageServicer(t, nil),
		Log:     os.Stderr,
	}
	if err := s.PrepareBucket(); err != nil {
		t.Error("PrepareBucket returns an error:", err.Error())
	}

}

// Test uploads a file.
func TestUploadFile(t *testing.T) {

	cfg := &config.Config{
		Gcp: config.Gcp{
			Project: "sample-project",
			Bucket:  "sample-bucket",
		},
	}

	service := newMockStorageServicer(t, cfg)
	s := &Storage{
		ctx:     config.NewContext(context.Background(), cfg),
		service: service,
		Log:     os.Stderr,
	}
	s.UploadFile(service.prefix, service.filename, service.filename)

}

// Test lists up files.
func TestListupFiles(t *testing.T) {

	cfg := &config.Config{
		Gcp: config.Gcp{
			Project: "sample-project",
			Bucket:  "sample-bucket",
		},
	}

	service := newMockStorageServicer(t, cfg)
	service.testList = true
	s := &Storage{
		ctx:     config.NewContext(context.Background(), cfg),
		service: service,
		Log:     os.Stderr,
	}

	stop := fmt.Errorf("Listing up should stop")
	// err := s.ListupFiles(service.prefix, func(info *FileInfo) error {
	s.ListupFiles(service.prefix, func(info *FileInfo) error {

		if strings.HasSuffix(info.Name, "end") {
			return stop
		}
		return nil

	})
	// if err != stop {
	// 	t.Error("Returned value isn't correct:", err)
	// }

}

// Test downloads a file.
func TestDownloadFiles(t *testing.T) {

	cfg := &config.Config{
		Gcp: config.Gcp{
			Project: "sample-project",
			Bucket:  "sample-bucket",
		},
	}

	service := newMockStorageServicer(t, cfg)
	s := &Storage{
		ctx:     config.NewContext(context.Background(), cfg),
		service: service,
		Log:     os.Stderr,
	}

	var err error
	temp, err := ioutil.TempDir("", "TestDownloadFiles")
	if err != nil {
		t.Fatal("Cannot create a temporal directory:", err.Error())
	}
	defer os.RemoveAll(temp)

	if err = s.DownloadFiles(service.prefix, temp, []string{service.filename}); err != nil {
		t.Error("DownloadFiiles returns an error:", err.Error())
	}

	res, err := ioutil.ReadFile(filepath.Join(temp, service.filename))
	if err != nil {
		t.Error(err.Error())
	}
	expected, err := ioutil.ReadFile(service.filename)
	if err != nil {
		t.Error(err.Error())
	}

	if string(res) != string(expected) {
		t.Error("Download a wrong file or the file is broken")
	}

}

// Test cancel works in downloading files.
func TestCancelDownloadFiles(t *testing.T) {

	cfg := &config.Config{
		Gcp: config.Gcp{
			Project: "sample-project",
			Bucket:  "sample-bucket",
		},
	}

	ctx, cancel := context.WithCancel(config.NewContext(context.Background(), cfg))
	service := newMockStorageServicer(t, cfg)
	s := &Storage{
		ctx:     ctx,
		service: service,
		Log:     os.Stderr,
	}

	var err error
	temp, err := ioutil.TempDir("", "TestDownloadFiles")
	if err != nil {
		t.Fatal("Cannot create a temporal directory:", err.Error())
	}
	defer os.RemoveAll(temp)

	// Cancel it.
	cancel()

	// Start download.
	if err = s.DownloadFiles(service.prefix, temp, []string{service.filename}); err == nil {

		t.Error("Download was canceled but doesn't return error")
		res, err := ioutil.ReadFile(filepath.Join(temp, service.filename))
		if err != nil {
			t.Log(err.Error())
		} else {
			t.Log("The downloaded file:", res)
		}

	}

}

// Test deleting a file.
func TestDeleteFiles(t *testing.T) {

	cfg := &config.Config{
		Gcp: config.Gcp{
			Project: "sample-project",
			Bucket:  "sample-bucket",
		},
	}

	service := newMockStorageServicer(t, cfg)
	s := &Storage{
		ctx:     config.NewContext(context.Background(), cfg),
		service: service,
		Log:     os.Stderr,
	}

	// Start download.
	if err := s.DeleteFiles(service.prefix, []string{service.filename}); err != nil {
		t.Error("DownloadFiiles returns an error:", err.Error())
	}

}

func TestPrintFileBody(t *testing.T) {

	var err error
	cfg := &config.Config{
		Gcp: config.Gcp{
			Project: "sample-project",
			Bucket:  "sample-bucket",
		},
	}

	service := newMockStorageServicer(t, cfg)
	s := &Storage{
		ctx:     config.NewContext(context.Background(), cfg),
		service: service,
		Log:     os.Stderr,
	}

	expected, err := ioutil.ReadFile(service.filename)
	if err != nil {
		t.Fatal("Cannot read a sample file:", err.Error())
	}

	output := &bytes.Buffer{}
	if err = s.PrintFileBody(service.prefix, service.filename, output, false); err != nil {
		t.Error("PrintFileBody returns an error:", err.Error())
	}
	if output.String() != string(expected) {
		t.Error("Printed file body isn't correct:", output.String())
	}

	output.Reset()
	if err = s.PrintFileBody(service.prefix, service.filename, output, true); err != nil {
		t.Error("PrintFileBody returns an error:", err.Error())
	}
	if !strings.HasSuffix(output.String(), string(expected)) {
		t.Error("Printed file body isn't correct:", output.String())
	}

}
