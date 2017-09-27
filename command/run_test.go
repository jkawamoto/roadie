//
// command/run_test.go
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

package command

import (
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"

	yaml "gopkg.in/yaml.v2"

	"github.com/jkawamoto/roadie/cloud/mock"
	"github.com/jkawamoto/roadie/script"
)

// TestCmdRun tests cmdRun function. cmdRun processes the following steps:
// 1. preparing a cloud storage,
// 2. loading a given script file,
// 3. uploading source files if necessary, and updating the script file,
// 4. creating an instance.
func TestCmdRun(t *testing.T) {

	var err error
	tmp, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("TempFile returns an error: %v", err)
	}
	defer os.Remove(tmp.Name())
	testScript := script.Script{
		Source: "roadie://source/somefile.tar.gz",
		Run: []string{
			"cmd {{args}}",
		},
	}
	data, err := yaml.Marshal(&testScript)
	if err != nil {
		t.Fatalf("Marchalling test script returns an error: %v", err)
	}
	tmp.Write(data)
	tmp.Close()

	p := mock.NewProvider()
	m := testMetadata(nil, p)

	t.Run("no instance name given", func(t *testing.T) {
		opt := runOpt{
			Metadata:   m,
			SourceOpt:  SourceOpt{},
			ScriptFile: tmp.Name(),
			ScriptArgs: []string{
				"args=a",
			},
		}
		err = cmdRun(&opt)
		if err != nil {
			t.Fatalf("cmdRun returns an error: %v", err)
		}

		var name string
		for key, status := range p.MockInstanceManager.Status {
			if strings.HasPrefix(key, path.Base(tmp.Name())) && status == mock.StatusRunning {
				name = key
			}
		}
		if name == "" {
			t.Error("instance is not running")
		} else {

			var expect *url.URL
			if uploadedScript := p.MockInstanceManager.Script[name]; uploadedScript == nil {
				t.Error("uploaded script is not found")
			} else if expect, err = createURL("result", name); err != nil {
				t.Fatalf("createURL returns an error: %v", err)
			} else if uploadedScript.Result != expect.String() {
				t.Errorf("result section is %q, want %q", uploadedScript.Result, expect)
			}

		}

	})

	t.Run("an instance name given", func(t *testing.T) {

		opt := runOpt{
			Metadata:   m,
			SourceOpt:  SourceOpt{},
			ScriptFile: tmp.Name(),
			ScriptArgs: []string{
				"args=a",
			},
			InstanceName: "test-instance",
		}
		err = cmdRun(&opt)
		if err != nil {
			t.Fatalf("cmdRun returns an error: %v", err)
		}

		if p.MockInstanceManager.Status[opt.InstanceName] != mock.StatusRunning {
			t.Errorf("instance %q is not running", opt.InstanceName)
		} else {

			if uploadedScript := p.MockInstanceManager.Script[opt.InstanceName]; uploadedScript == nil {
				t.Error("uploaded script is not found")
			} else if len(uploadedScript.Run) != 1 {
				t.Errorf("run section is not correct: %v", uploadedScript.Run)
			} else if uploadedScript.Run[0] != "cmd a" {
				t.Errorf("run section is %q, want %q", uploadedScript.Run[0], "cmd a")
			}

		}

	})

}
