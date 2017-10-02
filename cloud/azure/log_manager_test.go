//
// cloud/azure/log_manager_test.go
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
	"io/ioutil"
	"log"
	"net/url"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/jkawamoto/roadie/script"
)

func TestLogManagerGet(t *testing.T) {

	var err error
	server := newMockStorageServer()
	defer server.Close()

	cli, err := server.GetClient()
	if err != nil {
		t.Fatalf("cannot get a client: %v", err)
	}
	logger := log.New(ioutil.Discard, "", log.LstdFlags)

	m := LogManager{
		service: &StorageService{
			blobClient: cli.GetBlobService(),
			Logger:     logger,
		},
		Logger: logger,
	}

	ctx := context.Background()
	name := "test-name"
	data := "a\nb\nc\nd\n"

	var loc *url.URL
	for _, format := range []string{"%v-init.log", "%v.log"} {
		loc, err = url.Parse(script.RoadieSchemePrefix + path.Join(LogContainer, fmt.Sprintf(format, name)))
		if err != nil {
			t.Fatalf("cannot parse a URL: %v", err)
		}
		err = m.service.Upload(ctx, loc, strings.NewReader(data))
		if err != nil {
			t.Fatalf("cannot upload a file to %v: %v", loc, err)
		}
	}

	t.Run("retrieve logs", func(t *testing.T) {
		var res []string
		err = m.Get(ctx, name, time.Now(), func(t time.Time, line string, stderr bool) error {
			res = append(res, line)
			return nil
		})
		if err != nil {
			t.Errorf("Get returns an error: %v", err)
		}
		if strings.Join(res, "\n") != strings.TrimRight(data+data, "\n") {
			t.Errorf("received %v, want %v", res, data)
		}
	})

	t.Run("handler returns an error", func(t *testing.T) {
		expected := fmt.Errorf("some error")
		err = m.Get(ctx, name, time.Now(), func(t time.Time, line string, stderr bool) error {
			return expected
		})
		if err != expected {
			t.Errorf("Get returns %q, want %v", err, expected)
		}
	})

	t.Run("cancel retrieving logs", func(t *testing.T) {
		ctx2, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		err = m.Get(ctx2, name, time.Now(), func(t time.Time, line string, stderr bool) error {
			select {
			case <-time.After(4 * time.Second):
				return nil
			case <-ctx2.Done():
				return ctx.Err()
			}
		})
		if err == nil {
			t.Error("Get is canceled but no errors are returned")
		}
	})

}
