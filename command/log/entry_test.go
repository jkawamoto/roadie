//
// command/log/entry_test.go
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

package log

import (
	"context"
	"testing"

	"github.com/jkawamoto/roadie/config"

	"cloud.google.com/go/logging"
)

// Test for GetLogEntries method.
func TestGetLogEntries(t *testing.T) {

	// Test giving project name and filter are passed to requestDo.
	cfg := &config.Config{
		Gcp: config.Gcp{
			Project: "test-project",
		},
	}
	ctx := config.NewContext(context.Background(), cfg)
	filter := "test-filter"

	// Test giving entries are passed to handler.
	sampleEntry := logging.Entry{}
	GetEntriesFunc(ctx, filter, func(project, recvfilter string, handler EntryHandler) error {

		// Checking project id.
		t.Log("ProjectIds is", project)
		if project != cfg.Project {
			t.Error("ProjectIds doesn't have the giving project id")
		}

		// Checking filter.
		t.Log("Filter is", recvfilter)
		if recvfilter != filter {
			t.Error("Filter doesn't match the giving filter")
		}

		// Pass an entry to the handler.
		return handler(&sampleEntry)

	},
		// Handler checkes the given entry matches the passed one.
		func(entry *logging.Entry) error {
			if &sampleEntry != entry {
				t.Error("Entry doesn't match a passing one")
			}
			return nil
		})

}
