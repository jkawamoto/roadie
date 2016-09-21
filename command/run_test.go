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

	"golang.org/x/net/context"

	"github.com/jkawamoto/roadie/command/cloud"
	"github.com/jkawamoto/roadie/command/resource"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/jkawamoto/roadie/config"
)

func TestCmdRun(t *testing.T) {
	// t.Error("Test is not implemented.")

	// With instance name/without instance name

	// with correct script / no script

	// Get flag.

	// URL flag.

	// Local flag.

	// Source flag.

	// OverWriteResultSection

	// shutdown option

}

// TestSetGitSource checks setGitSource sets correct repository URL.
func TestSetGitSource(t *testing.T) {

	script := resource.Script{}
	setGitSource(&script, "https://github.com/jkawamoto/roadie.git")

	if script.Body.Source != "https://github.com/jkawamoto/roadie.git" {
		t.Errorf("source section is not correct: %s", script.Body.Source)
	}

}

// TestSetURLSource checks setURLSource sets correct url.
func TestSetURLSource(t *testing.T) {

	script := resource.Script{}
	setURLSource(&script, "https://github.com/jkawamoto/roadie")

	if script.Body.Source != "https://github.com/jkawamoto/roadie" {
		t.Errorf("source section is not correct: %s", script.Body.Source)
	}

}

// TestSetLocalSource checks setLocalSource sets correct url with any directories,
// and file paths. This test doesn't check excludes parameters since those parameters
// are tested in tests for util.Archive.
func TestSetLocalSource(t *testing.T) {

	conf := &config.Config{
		Gcp: config.Gcp{
			Bucket: "somebucket",
		},
	}
	ctx := config.NewContext(context.Background(), conf)
	storage := &cloud.Storage{}

	var script resource.Script
	var err error

	// Test with directories.
	for _, target := range []string{".", "../command", ".."} {

		script = resource.Script{
			InstanceName: "test",
		}

		t.Logf("Trying target %s", target)
		if err = setLocalSource(ctx, storage, &script, target, nil, true); err != nil {
			t.Error(err.Error())
		}
		if script.Body.Source != util.CreateURL("somebucket", SourcePrefix, "test.tar.gz").String() {
			t.Errorf("source section is not correct: %s", script.Body.Source)
		}

	}

	// Test with a file.
	script = resource.Script{
		InstanceName: "test",
	}
	if err = setLocalSource(ctx, storage, &script, "run.go", nil, true); err != nil {
		t.Error(err.Error())
	}
	if script.Body.Source != util.CreateURL("somebucket", SourcePrefix, "run.go").String() {
		t.Errorf("source section is not correct: %s", script.Body.Source)
	}

	// Test with unexisting file.
	if err = setLocalSource(ctx, storage, &script, "abcd.efg", nil, true); err == nil {
		t.Error("Give an unexisting path but no error occurs.")
	}
	t.Logf("Give an unexisting path to setLocalSource and got an error: %s", err.Error())

}

// TestSetSource checks setSource sets correct url from a given filename.
func TestSetSource(t *testing.T) {

	conf := &config.Config{
		Gcp: config.Gcp{
			Bucket: "somebucket",
		},
	}
	script := &resource.Script{}

	setSource(conf, script, "abc.zip")

	if script.Body.Source != util.CreateURL("somebucket", SourcePrefix, "abc.zip").String() {
		t.Errorf("source section is not correct: %s", script.Body.Source)
	}

}

// TestReplaceURLScheme checks that function replaces URLs which start with "roadie://".
// to "gs://<bucketname>/.roadie/".
func TestReplaceURLScheme(t *testing.T) {

	type ScriptBody struct {
		APT    []string `yaml:"apt,omitempty"`
		Source string   `yaml:"source,omitempty"`
		Data   []string `yaml:"data,omitempty"`
		Run    []string `yaml:"run,omitempty"`
		Result string   `yaml:"result,omitempty"`
		Upload []string `yaml:"upload,omitempty"`
	}

	conf := config.Config{
		Gcp: config.Gcp{
			Bucket: "test-bucket",
		},
	}

	script := resource.Script{
		Body: resource.ScriptBody{
			Source: "roadie://some-sourcefile",
			Data: []string{
				"roadie://some-datafile",
			},
			Result: "roadie://result-file",
		},
	}

	// Run.
	if err := replaceURLScheme(&conf, &script); err != nil {
		t.Fatal("replaceURLScheme returns an error:", err.Error())
	}

	// Check results.
	if script.Body.Source != "gs://test-bucket/.roadie/source/some-sourcefile" {
		t.Error("source section is not correct:", script.Body.Source)
	}
	if script.Body.Data[0] != "gs://test-bucket/.roadie/data/some-datafile" {
		t.Error("data section is not correct:", script.Body.Data)
	}
	if script.Body.Result != "gs://test-bucket/.roadie/result/result-file" {
		t.Error("result section is not correct:", script.Body.Result)
	}

}
