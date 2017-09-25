//
// command/status_test.go
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

package command

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/cloud/mock"
)

func TestCmdStatus(t *testing.T) {

	var err error
	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)

	statuses := map[string]string{
		"instance1":  mock.StatusRunning,
		"instance2":  mock.StatusRunning,
		"instance3":  mock.StatusRunning,
		"instance11": mock.StatusTerminated,
		"instance12": mock.StatusTerminated,
	}
	for name, status := range statuses {
		p.MockInstanceManager.Status[name] = status
	}

	err = cmdStatus(m, false)
	if err != nil {
		t.Fatalf("cmdStatus returns an error: %v", err)
	}

	lines := strings.Split(output.String(), "\n")
	if !strings.HasPrefix(lines[0], "INSTANCE NAME") {
		t.Errorf("1st line is not header: %v", lines[0])
	}
	if len(lines) != len(statuses)+1 {
		t.Errorf("%v lines outputted, want %v lines ", len(lines), len(statuses)+1)
	}

	for _, line := range lines[1:] {
		kv := strings.Split(line, "\t")
		if len(kv) < 2 {
			t.Errorf("line has missing items: %v", line)
			continue
		}

		name := strings.TrimSpace(kv[0])
		status := strings.TrimSpace(kv[1])
		if statuses[name] != status {
			t.Errorf("status of %v is %v, want %v", name, status, statuses[name])
		}
	}

}

func TestCmdStatusKill(t *testing.T) {

	var err error
	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)

	statuses := map[string]string{
		"instance1":  mock.StatusRunning,
		"instance2":  mock.StatusRunning,
		"instance3":  mock.StatusRunning,
		"instance11": mock.StatusTerminated,
		"instance12": mock.StatusTerminated,
	}
	for name, status := range statuses {
		p.MockInstanceManager.Status[name] = status
	}

	t.Run("kill an instance", func(t *testing.T) {
		err = cmdStatusKill(m, "instance1")
		if err != nil {
			t.Fatalf("cmdStatusKill of instance1 returns an error: %v", err)
		}
		if p.MockInstanceManager.Status["instance1"] != mock.StatusTerminated {
			t.Errorf("killed instance's status %q", p.MockInstanceManager.Status["instance1"])
		}
	})

	t.Run("kill a terminated instance", func(t *testing.T) {
		err = cmdStatusKill(m, "instance11")
		if err == nil {
			t.Error("killed a terminated instance but no errors are returned")
		}
	})

	t.Run("kill an unknown instance", func(t *testing.T) {
		err = cmdStatusKill(m, "instance42")
		if err == nil {
			t.Error("killed not existing instance but no errors are returned")
		}
	})

}
