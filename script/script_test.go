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
- roadie://somebucket/somedata
run:
- abc def
result: roadie://somebucket/result
upload:
- xyz
`

	// complexScript is a script which has a place holder.
	complexScript = `
apt:
- python-numpy
source: https://github.com/jkawamoto/roadie
data:
- roadie://somebucket/somedata
run:
- abc {{args}}
result: roadie://somebucket/result
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
		t.Fatalf("cannot open file %q: %v", filename, err)
	}
	_, err = fp.WriteString(simpleScript)
	if err != nil {
		t.Fatalf("cannot write a sample script: %v", err)
	}
	if err = fp.Close(); err != nil {
		t.Fatalf("cannot close the file: %v", err)
	}
	t.Logf("script file created in %s", filename)
	defer os.Remove(filename)

	// Loading test.
	script, err := NewScriptTemplate(filename, nil)
	if err != nil {
		t.Fatalf("cannot read script file %q: %v", filename, err)
	}

	// Tests
	if script.Name != strings.ToLower(script.Name) {
		t.Errorf("Instance name %s has upper cases.", script.Name)
	}

	if strings.Contains(script.Name, ".") {
		t.Errorf("Instance name %s has dots.", script.Name)
	}

	if len(script.APT) != 1 || script.APT[0] != "python-numpy" {
		t.Errorf("apt section is not correct: %s", script.APT)
	}

	if script.Source != "https://github.com/jkawamoto/roadie" {
		t.Errorf("source section is not correct: %s", script.Source)
	}

	if len(script.Data) != 1 || script.Data[0] != "roadie://somebucket/somedata" {
		t.Errorf("data section is not correct: %s", script.Data)
	}

	if len(script.Run) != 1 || script.Run[0] != "abc def" {
		t.Errorf("run section is not correct: %s", script.Run)
	}

	if script.Result != "roadie://somebucket/result" {
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
		t.Fatalf("cannot create %q: %v", filename, err)
	}
	_, err = fp.WriteString(complexScript)
	if err != nil {
		t.Fatalf("cannot write to file %q: %v", filename, err)
	}
	if err = fp.Close(); err != nil {
		t.Fatalf("cannot close the script file: %v", err)
	}
	t.Logf("script file is created in %s", filename)
	defer os.Remove(filename)

	cases := []struct {
		params   string
		expected string
	}{
		{"args=xyz", "abc xyz"},
		{"args=param=x", "abc param=x"},
	}

	var script *Script
	for _, c := range cases {

		// Loading test.
		script, err = NewScriptTemplate(filename, []string{c.params})
		if err != nil {
			t.Fatalf("cannot read script %q: %v", filename, err)
		}

		// Tests
		if len(script.Run) != 1 || script.Run[0] != c.expected {
			t.Errorf("run section is not correct: %q, want %q", script.Run, c.expected)
		}

	}

	// Loading without parameters.
	_, err = NewScriptTemplate(filename, nil)
	if err == nil {
		t.Error("Placeholders are not given but script is created.")
	}

}
