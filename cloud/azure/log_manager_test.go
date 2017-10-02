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
		storage: &StorageService{
			blobClient: cli.GetBlobService(),
			Logger:     logger,
		},
		Logger: logger,
	}

	ctx := context.Background()
	name := "test-name"
	var data []string
	for i := 0; i != 10; i++ {
		data = append(data, fmt.Sprintf("%v %v", time.Now().UTC().Format("2006/01/02 15:04:05"), i))
	}

	var loc *url.URL
	for _, format := range []string{"%v-init.log", "%v.log"} {
		loc, err = url.Parse(script.RoadieSchemePrefix + path.Join(LogContainer, fmt.Sprintf(format, name)))
		if err != nil {
			t.Fatalf("cannot parse a URL: %v", err)
		}
		err = m.storage.Upload(ctx, loc, strings.NewReader(strings.Join(data, "\n")+"\n"))
		if err != nil {
			t.Fatalf("cannot upload a file to %v: %v", loc, err)
		}
	}

	t.Run("retrieve logs", func(t *testing.T) {
		var c int
		now := time.Now()
		err = m.Get(ctx, name, time.Now(), func(ti time.Time, line string, stderr bool) error {
			if !now.After(ti) || now.Sub(ti) > 5*time.Minute {
				t.Errorf("time a log entry issued is %q, now %q", ti, now)
			}
			if line != fmt.Sprint(c%10) {
				t.Errorf("log entry is %q, want %v", line, c%10)
			}
			c++
			return nil
		})
		if err != nil {
			t.Errorf("Get returns an error: %v", err)
		}
		if c != len(data)*2 {
			t.Errorf("%v log entries, want %v", c, len(data)*2)
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
