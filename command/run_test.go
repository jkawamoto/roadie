//
// command/run_test.go
//
// Copyright (c) 2016 Junpei Kawamoto
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

package command

import (
	"testing"

	"github.com/jkawamoto/roadie/config"
	"github.com/jkawamoto/roadie/util"
)

func TestCmdRun(t *testing.T) {
	// t.Error("Test is not implemented.")

	// With instance name/without instance name

	// with corerct script / no script

	// Get flag.

	// URL flag.

	// Local flag.

	// Source flag.

	// OverWriteResultSection

	// shutdown option

}

// TestSetGitSource checks setGitSource sets correct repository url.
func TestSetGitSource(t *testing.T) {

	script := Script{}
	setGitSource(&script, "https://github.com/jkawamoto/roadie.git")

	if script.Body.Source != "https://github.com/jkawamoto/roadie.git" {
		t.Errorf("source section is not correct: %s", script.Body.Source)
	}

}

// TestSetURLSource checks setURLSource sets correct url.
func TestSetURLSource(t *testing.T) {

	script := Script{}
	setURLSource(&script, "https://github.com/jkawamoto/roadie")

	if script.Body.Source != "https://github.com/jkawamoto/roadie" {
		t.Errorf("source section is not correct: %s", script.Body.Source)
	}

}

// TestSetLocalSource checks setLocalSource sets correct url with any directories,
// and file paths. This test doesn't check excludes parameters since those parameters
// are tested in tests for util.Archive.
func TestSetLocalSource(t *testing.T) {

	conf := config.Config{}
	conf.Gcp.Bucket = "somebucket"

	var script Script
	var err error

	// Test with directories.
	for _, target := range []string{".", "../command", ".."} {

		script = Script{
			InstanceName: "test",
		}

		t.Logf("Trying target %s", target)
		if err = setLocalSource(&conf, &script, target, nil, true); err != nil {
			t.Error(err.Error())
		}
		if script.Body.Source != util.CreateURL("somebucket", SourcePrefix, "test.tar.gz").String() {
			t.Errorf("source section is not correct: %s", script.Body.Source)
		}

	}

	// Test with a file.
	script = Script{
		InstanceName: "test",
	}
	if err = setLocalSource(&conf, &script, "run.go", nil, true); err != nil {
		t.Error(err.Error())
	}
	if script.Body.Source != util.CreateURL("somebucket", SourcePrefix, "run.go").String() {
		t.Errorf("source section is not correct: %s", script.Body.Source)
	}

	// Test with unexisting file.
	if err = setLocalSource(&conf, &script, "abcd.efg", nil, true); err == nil {
		t.Error("Give an unexisting path but no error occurs.")
	}
	t.Logf("Give an unexisting path to setLocalSource and got an error: %s", err.Error())

}

// TestSetSource checks setSource sets correct url from a given filename.
func TestSetSource(t *testing.T) {

	conf := config.Config{}
	conf.Gcp.Bucket = "somebucket"

	script := Script{}

	setSource(&conf, &script, "abc.zip")

	if script.Body.Source != util.CreateURL("somebucket", SourcePrefix, "abc.zip").String() {
		t.Errorf("source section is not correct: %s", script.Body.Source)
	}

}
