//
// cloud/gcp/helper_test.go
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

package gcp

import (
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/jkawamoto/roadie/script"

	yaml "gopkg.in/yaml.v2"
)

func GetConfig() *Config {

	data, err := ioutil.ReadFile("test_config.yml")
	if err != nil {
		return nil
	}

	cfg := new(Config)
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil
	}

	return cfg

}

// TestReplaceURLScheme checks that function replaces URLs which start with "roadie://".
// to "gs://<bucketname>/.roadie/".
func TestReplaceURLScheme(t *testing.T) {

	project := "sample-project"
	bucket := "test-bucket"
	region := "us-central1-c"
	machine := "n1-standard-2"
	cfg := &Config{
		Project:     project,
		Bucket:      bucket,
		Zone:        region,
		MachineType: machine,
	}

	task := script.Script{
		Source: "roadie://source/some-sourcefile",
		Data: []string{
			"roadie://data/some-datafile",
		},
		Result: "roadie://result/result-file",
	}
	ReplaceURLScheme(cfg, &task)

	// Check results.
	if task.Source != "gs://test-bucket/.roadie/source/some-sourcefile" {
		t.Error("source section is not correct:", task.Source)
	}
	if task.Data[0] != "gs://test-bucket/.roadie/data/some-datafile" {
		t.Error("data section is not correct:", task.Data)
	}
	if task.Result != "gs://test-bucket/.roadie/result/result-file" {
		t.Error("result section is not correct:", task.Result)
	}

}

func TestCreateURL(t *testing.T) {

	project := "sample-project"
	bucket := "test-bucket"
	region := "us-central1-c"
	machine := "n1-standard-2"
	cfg := &Config{
		Project:     project,
		Bucket:      bucket,
		Zone:        region,
		MachineType: machine,
	}

	u, err := url.Parse(CreateURL(cfg, "/path/to/file"))
	if err != nil {
		t.Error(err.Error())
	}
	if u.Scheme != "gs" {
		t.Errorf("Scheme is not correct: %s", u.Scheme)
	}
	if u.Host != bucket {
		t.Errorf("Host name is not correct: %s", u.Host)
	}
	if u.Path != "/.roadie/path/to/file" {
		t.Errorf("Path is not correct: %s", u.Path)
	}

}
