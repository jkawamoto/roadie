//
// cloud/mock/storage_manager_test.go
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
	"bytes"
	"context"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/cloud"
)

func TestUpload(t *testing.T) {

	var err error
	m := NewStorageManager()
	ctx := context.Background()

	data := "1234567890"
	loc, err := url.Parse("roadie://mock/test.txt")
	if err != nil {
		t.Fatalf("Cannot parse a URL: %v", err)
	}

	err = m.Upload(ctx, loc, strings.NewReader(data))
	if err != nil {
		t.Fatalf("Upload returns an error: %v", err)
	}

	res, ok := m.storage[loc.String()]
	if !ok {
		t.Errorf("%v doesn't exist", loc)
	} else if res != data {
		t.Errorf("%v is uploaded, want %v", res, data)
	}

	// Test when Failure is true.
	m.Failure = true
	err = m.Upload(ctx, loc, strings.NewReader(data))
	if err == nil {
		t.Error("Failure is true but Upload doesn't return any error")
	}

}

func TestDownload(t *testing.T) {

	var err error
	m := NewStorageManager()
	ctx := context.Background()

	data := "1234567890"
	loc, err := url.Parse("roadie://mock/test.txt")
	if err != nil {
		t.Fatalf("Cannot parse a URL: %v", err)
	}
	err = m.Upload(ctx, loc, strings.NewReader(data))
	if err != nil {
		t.Fatalf("Cannot upload to %v: %v", loc, err)
	}

	var out bytes.Buffer
	err = m.Download(ctx, loc, &out)
	if err != nil {
		t.Fatalf("Download returns an error: %v", err)
	}
	if out.String() != data {
		t.Errorf("%v is downloaded, want %v", out.String(), data)
	}

	// Test with not existing URL.
	dummy, err := url.Parse("roadie://mock/test2.txt")
	if err != nil {
		t.Fatalf("Cannot parse a url: %v", err)
	}
	err = m.Download(ctx, dummy, &out)
	if err == os.ErrNotExist {
		t.Errorf("Download doesn't return ErrNotExist: %v", err)
	}

	// Test when Failure is true.
	m.Failure = true
	err = m.Download(ctx, loc, &out)
	if err == nil {
		t.Error("Failure is true but Download doesn't return any error")
	}

}

func TestGetFileInfo(t *testing.T) {

	var err error
	m := NewStorageManager()
	ctx := context.Background()

	data := "1234567890"
	loc, err := url.Parse("roadie://mock/test.txt")
	if err != nil {
		t.Fatalf("Cannot parse a URL: %v", err)
	}
	err = m.Upload(ctx, loc, strings.NewReader(data))
	if err != nil {
		t.Fatalf("Cannot upload to %v: %v", loc, err)
	}

	info, err := m.GetFileInfo(ctx, loc)
	if err != nil {
		t.Fatalf("GetFileInfo returns an error: %v", err)
	}
	if info.Name != path.Base(loc.Path) {
		t.Errorf("Name = %v, want %v", info.Name, path.Base(loc.Path))
	}
	if info.Path != loc.Path {
		t.Errorf("Path = %v, want %v", info.Path, loc.Path)
	}
	if info.Size != int64(len(data)) {
		t.Errorf("Size = %v, want %v", info.Size, len(data))
	}

	// Test with not existing URL.
	dummy, err := url.Parse("roadie://mock/test2.txt")
	if err != nil {
		t.Fatalf("Cannot parse a url: %v", err)
	}
	_, err = m.GetFileInfo(ctx, dummy)
	if err != os.ErrNotExist {
		t.Errorf("GetFileInfo doesn't return ErrNotExist: %v", err)
	}

	// Test when Failure is true.
	m.Failure = true
	_, err = m.GetFileInfo(ctx, loc)
	if err == nil {
		t.Error("Failure is true but GetFileInfo doesn't return any error")
	}

}

func TestList(t *testing.T) {

	var err error
	m := NewStorageManager()
	ctx := context.Background()

	files := []string{
		"roadie://mock/test.txt",
		"roadie://mock/test2.txt",
		"roadie://mock/dir/test.txt",
		"roadie://mock/dir/test2.txt",
	}
	for _, v := range files {
		var loc *url.URL
		loc, err = url.Parse(v)
		if err != nil {
			t.Fatalf("Cannot parse a URL: %v", err)
		}
		err = m.Upload(ctx, loc, strings.NewReader(v))
		if err != nil {
			t.Fatalf("Cannot upload to %v: %v", loc, err)
		}
	}

	cases := []struct {
		query    string
		expected []string
	}{
		{"roadie://mock/test", files[:2]},
		{"roadie://mock/", append(files, "roadie://mock/", "roadie://mock/dir/")},
		{"roadie://mock/dir", append(files[2:], "roadie://mock/dir/")},
		{"roadie://mock/dummy", []string{}},
	}

	for _, c := range cases {

		var query *url.URL
		query, err = url.Parse(c.query)
		if err != nil {
			t.Fatalf("Cannot parse a URL: %v", err)
		}
		res := make(map[string]struct{})
		m.List(ctx, query, func(info *cloud.FileInfo) error {
			res[info.Path] = struct{}{}
			return nil
		})
		if len(res) != len(c.expected) {
			t.Errorf("the number of found files = %v, want %v", len(res), len(c.expected))
		}
		for _, v := range c.expected {
			if _, exist := res[v]; !exist {
				t.Errorf("%v is not listed", v)
			}
		}

	}

	// Test that a handler returns an error.
	query, err := url.Parse(cases[0].query)
	if err != nil {
		t.Fatalf("Cannot parse a URL: %v", err)
	}
	expectedError := fmt.Errorf("test error")
	err = m.List(ctx, query, func(info *cloud.FileInfo) error {
		return expectedError
	})
	if err != expectedError {
		t.Errorf("handler returns an error but List doesn't return the error: %v", err)
	}

	// Test when Failure is true.
	m.Failure = true
	err = m.List(ctx, query, func(info *cloud.FileInfo) error {
		return nil
	})
	if err == nil {
		t.Error("Failure is true but GetFileInfo doesn't return any error")
	}

}

func TestDelete(t *testing.T) {

	var err error
	m := NewStorageManager()
	ctx := context.Background()

	data := "1234567890"
	loc, err := url.Parse("roadie://mock/test.txt")
	if err != nil {
		t.Fatalf("Cannot parse a URL: %v", err)
	}

	err = m.Upload(ctx, loc, strings.NewReader(data))
	if err != nil {
		t.Fatalf("Upload returns an error: %v", err)
	}

	err = m.Delete(ctx, loc)
	if err != nil {
		t.Fatalf("Delete returns an error: %v", err)
	}
	if _, exist := m.storage[loc.String()]; exist {
		t.Fatal("Deleted file still exists")
	}

	// Delete not existing file.
	err = m.Delete(ctx, loc)
	if err != os.ErrNotExist {
		t.Errorf("Delete doesn't return os.ErrNotExist: %v", err)
	}

	m.Failure = true
	err = m.Delete(ctx, loc)
	if err == nil {
		t.Error("Failure is true but Delete doesn't return any error")
	}

}
