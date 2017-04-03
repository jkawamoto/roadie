//
// script/script_test.go
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

package script

import (
	"os"
	"path"
	"strings"
	"testing"
)

const (
	// simpleScript is a script which doesn't have any place holders.
	simpleScript = `
apt:
- python-numpy
source: https://github.com/jkawamoto/roadie
data:
- gs://somebucket/somedata
run:
- abc def
result: gs://somebucket/result
upload:
- xyz
`

	// complexScript is a script which has a place holder.
	complexScript = `
apt:
- python-numpy
source: https://github.com/jkawamoto/roadie
data:
- gs://somebucket/somedata
run:
- abc {{args}}
result: gs://somebucket/result
upload:
- xyz
`
)

// TestLoadScript tests loading a script which doesn't have place holders.
func TestLoadScript(t *testing.T) {

	var err error

	// Prepare testing.
	filename := path.Join(os.TempDir(), "test.yaml")
	fp, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Error(err.Error())
	}
	_, err = fp.WriteString(simpleScript)
	if err != nil {
		t.Error(err.Error())
	}
	if err = fp.Close(); err != nil {
		t.Error(err.Error())
	}
	t.Logf("Create a test script file in %s", filename)
	defer os.Remove(filename)

	// Loading test.
	script, err := NewScript(filename, nil)
	if err != nil {
		t.Error(err.Error())
	}

	// Tests
	if script.InstanceName != strings.ToLower(script.InstanceName) {
		t.Errorf("Instance name %s has upper cases.", script.InstanceName)
	}

	if strings.Contains(script.InstanceName, ".") {
		t.Errorf("Instance name %s has dots.", script.InstanceName)
	}

	if len(script.APT) != 1 || script.APT[0] != "python-numpy" {
		t.Errorf("apt section is not correct: %s", script.APT)
	}

	if script.Source != "https://github.com/jkawamoto/roadie" {
		t.Errorf("source section is not correct: %s", script.Source)
	}

	if len(script.Data) != 1 || script.Data[0] != "gs://somebucket/somedata" {
		t.Errorf("data section is not correct: %s", script.Data)
	}

	if len(script.Run) != 1 || script.Run[0] != "abc def" {
		t.Errorf("run section is not correct: %s", script.Run)
	}

	if script.Result != "gs://somebucket/result" {
		t.Errorf("result section is not correct: %s", script.Result)
	}

	if len(script.Upload) != 1 || script.Upload[0] != "xyz" {
		t.Errorf("upload section is not correct: %s", script.Upload)
	}

}

// TestLoadScriptWithPlaceholders tests loading a script which has place holders.
func TestLoadScriptWithPlaceholders(t *testing.T) {

	var err error

	// Prepare testing.
	filename := path.Join(os.TempDir(), "test.yaml")
	fp, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Error(err.Error())
	}
	_, err = fp.WriteString(complexScript)
	if err != nil {
		t.Error(err.Error())
	}
	if err = fp.Close(); err != nil {
		t.Error(err.Error())
	}
	t.Logf("Create a test script file in %s", filename)
	defer os.Remove(filename)

	// Loading test.
	script, err := NewScript(filename, []string{"args=xyz"})
	if err != nil {
		t.Error(err.Error())
	}

	// Tests
	if len(script.Run) != 1 || script.Run[0] != "abc xyz" {
		t.Errorf("run section is not correct: %s", script.Run)
	}

	// Loading without parameters test.
	_, err = NewScript(filename, nil)
	if err == nil {
		t.Error("Placeholders are not given but script is created.")
	}
	t.Logf("Load script w/o arguments gets an error: %s", err.Error())

}
