//
// config/config_test.go
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

package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// Test saving config.
func TestSave(t *testing.T) {

	var err error

	dir, err := ioutil.TempDir("", "config-test")
	if err != nil {
		t.Fatal("Cannot make a temporal directory:", err.Error())
	}
	defer os.RemoveAll(dir)

	cfg := Config{
		Filename: filepath.Join(dir, "config.toml"),
		Gcp: Gcp{
			Project:     "sample-project",
			Bucket:      "sample-bucket",
			Zone:        "sample-zone",
			MachineType: "sample-machine-type",
		},
	}

	// Test Save method.
	if err = cfg.Save(); err != nil {
		t.Error("Save returns an error:", err.Error())
	}

	raw, err := ioutil.ReadFile(cfg.Filename)
	if err != nil {
		t.Error("Cannot read a saved config file:", err.Error())
	}

	data := string(raw)
	t.Log("Saved config:\n", data)
	if !strings.Contains(data, cfg.Project) {
		t.Error("Project isn't saved")
	}
	if !strings.Contains(data, cfg.Bucket) {
		t.Error("Bucket isn't saved")
	}
	if !strings.Contains(data, cfg.Zone) {
		t.Error("Zone isn't saved")
	}
	if !strings.Contains(data, cfg.MachineType) {
		t.Error("MachineType isn't saved")
	}

}

// TestLookup tests lookup function.
func TestLookup(t *testing.T) {

	// Prepare temporary directory.
	temp := filepath.Join(os.TempDir(), "roadie-test", time.Now().Format("20060102150405"), "config")
	err := os.MkdirAll(temp, 0744)
	if err != nil {
		t.Error(err.Error())
		return
	}
	temp, err = filepath.EvalSymlinks(temp)
	if err != nil {
		t.Error(err.Error())
		return
	}

	// Move to the temporary directory.
	cd, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
		return
	}
	if err = os.Chdir(temp); err != nil {
		t.Error(err.Error())
		return
	}
	defer os.Chdir(cd)

	// Lookup from a directory w/o configuration file ans git repository.
	ans := filepath.Join(temp, ConfigureFile)
	test, err := filepath.Abs(lookup())
	if err != nil {
		t.Error(err.Error())
		return
	}
	if ans != test {
		t.Errorf("%s shoult be %s", test, ans)
	}

	// Lookup to a directory which has a configuration file.
	ans = filepath.Join(temp, "..", ConfigureFile)
	if err = ioutil.WriteFile(ans, []byte{}, 0644); err != nil {
		t.Error(err.Error())
		return
	}
	test, err = filepath.Abs(lookup())
	if err != nil {
		t.Error(err.Error())
		return
	}
	if ans != test {
		t.Errorf("%s shoult be %s", test, ans)
	}
	os.Remove(ans)

	// Lookup to a directory which has a git repository.
	os.Mkdir(filepath.Join(temp, "..", ".git"), 755)
	test, err = filepath.Abs(lookup())
	if err != nil {
		t.Error(err.Error())
		return
	}
	if ans != test {
		t.Errorf("%s shoult be %s", test, ans)
	}

}
