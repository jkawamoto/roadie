//
// cloud/gce/startup_test.go
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

// TestStartup tests startup function.
func TestStartup(t *testing.T) {

	opt := &StartupOpt{
		Name:    "test-name",
		Script:  "test-script\nline2",
		Options: "--options",
		Image:   "repos/image",
		Retry:   25,
	}

	res, err := Startup(opt)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(res)
	if !strings.Contains(res, fmt.Sprintf("INSTANCE=%s", opt.Name)) {
		t.Error("Instance name is wrong")
	}

	if !strings.Contains(res, fmt.Sprintf("docker rm -f %s", opt.Name)) {
		t.Error("Before retrying, previous container must be deleted.")
	}

	if !strings.Contains(res, opt.Script) {
		t.Error("Script data are wrong")
	}

	if !strings.Contains(res, fmt.Sprintf("%s < run.yml", opt.Options)) {
		t.Error("Options are wrong")
	}

	if !strings.Contains(res, opt.Image) {
		t.Error("Image information is wrong")
	}

	if !strings.Contains(res, fmt.Sprintf("seq %d", opt.Retry)) {
		t.Error("Retry number is wrong")
	}

}
