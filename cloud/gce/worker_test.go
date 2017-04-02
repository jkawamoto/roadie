//
// cloud/gce/woker_test.go
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

package gce

import (
	"fmt"
	"strings"
	"testing"
)

// Test WorkerStartup returns a correct startup script.
func TestWorkerStartup(t *testing.T) {

	opt := WorkerStartupOpt{
		ProjectID:    "sample-project",
		InstanceName: "sample-instance",
		Name:         "sample-queue-1",
		Version:      "1.0.0",
	}

	res, err := WorkerStartup(&opt)
	if err != nil {
		t.Error("WorkerStartup returns an error:", err.Error())
	}

	t.Log(res)
	if !strings.Contains(res, fmt.Sprintf("INSTANCE=%s", opt.InstanceName)) {
		t.Error("Generated script doesn't have an instane name:", res)
	}
	if !strings.Contains(res, fmt.Sprintf(
		"roadie-queue-manager_%s_linux_amd64", opt.Version)) {

		t.Error("Generated script doesn't have a correct file name:", res)
	}
	if !strings.Contains(res, fmt.Sprintf(
		"./roadie-queue-manager %s %s", opt.ProjectID, opt.Name)) {

		t.Error("Generated script doesn't have correct variables for queue manager.")
	}

}
