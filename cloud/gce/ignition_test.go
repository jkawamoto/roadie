//
// cloud/gce/ignition_test.go
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

func TestNewIgnitionConfig(t *testing.T) {

	cfg := NewIgnitionConfig()
	if cfg.Ignition.Version != DefaultIgnitionVersion {
		t.Error("Created ignition config has wrong version:", cfg.Ignition.Version)
	}

}

func TestFluentdUnit(t *testing.T) {

	name := "testname"
	unit, err := FluentdUnit(name)
	if err != nil {
		t.Error(err.Error())
	}

	if !strings.Contains(unit.Contents, fmt.Sprintf("INSTANCE=%v", name)) {
		t.Error("Created fluentd unit doesn't have a correct instane name")
	}
	if strings.Contains(unit.Contents, "# fluentd.service") {
		t.Error("Unit contents still has comments")
	}

}

func TestRoadieUnit(t *testing.T) {

	name := "testname"
	image := "testimage"
	options := "options"
	unit, err := RoadieUnit(name, image, options)
	if err != nil {
		t.Error(err.Error())
	}

	if !strings.Contains(unit.Contents, fmt.Sprintf("--name %v", name)) {
		t.Error("Created unit doesn't have a correct instance name")
	}
	if !strings.Contains(unit.Contents, image) {
		t.Error("Created unit doesn't have a correct image name")
	}
	if !strings.Contains(unit.Contents, options) {
		t.Error("Created unit doesn't have correct options")
	}
	if strings.Contains(unit.Contents, "# roadie.service") {
		t.Error("Unit contents still has comments")
	}

}

func TestLogcastUnit(t *testing.T) {

	unit, err := LogcastUnit()
	if err != nil {
		t.Error(err.Error())
	}

	if !strings.Contains(unit.Contents, "/usr/bin/ncat 127.0.0.1 24225") {
		t.Error("Created unit doesn't have a correct cmd")
	}
	if strings.Contains(unit.Contents, "# logcast.service") {
		t.Error("Unit contents still has comments")
	}

}
