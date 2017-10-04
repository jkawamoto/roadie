//
// cloud/azure/storage_remote_test.go
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

package azure

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
)

func TestStorageServiceWithRemoteConnection(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	ctx := context.Background()
	account := fmt.Sprintf("test-storage%v", time.Now().Unix())
	cfg.AccountName = account
	s, err := NewStorageService(ctx, cfg, nil)
	if err != nil {
		t.Fatalf("cannot create a storage service: %v", err)
	}
	defer func() {
		err = s.deleteAccount(ctx)
		if err != nil {
			t.Errorf("cannot delete storage account: %v", err)
		}
	}()

	fp, err := os.Open("Makefile")
	defer fp.Close()

	loc, err := url.Parse(script.RoadieSchemePrefix + fmt.Sprintf("f%d", time.Now().Unix()))
	if err != nil {
		t.Fatalf("cannot parse a URL: %v", err)
	}
	t.Log("uploading...", loc)

	err = s.Upload(ctx, loc, fp)
	if err != nil {
		t.Fatalf("cannot upload a file: %v", err)
	}

	info, err := s.GetFileInfo(ctx, loc)
	if err != nil {
		t.Fatalf("cannot get file information: %v", err)
	}
	t.Log(info)

	query := *loc
	query.Path = ""
	err = s.List(ctx, &query, func(info *cloud.FileInfo) error {
		t.Log(info)
		return nil
	})
	if err != nil {
		t.Fatalf("List returns an error: %v", err)
	}

	err = s.Delete(ctx, loc)
	if err != nil {
		t.Fatalf("Delete returns an error: %v", err)
	}

}
