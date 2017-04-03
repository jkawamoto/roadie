//
// cloud/gce/instance_test.go
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
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/jkawamoto/roadie/script"
)

func TestNewComputeService(t *testing.T) {

	project := "sample-project"
	region := "us-central1-c"
	machine := "n1-standard-2"
	cfg := &GcpConfig{
		Project:     project,
		Zone:        region,
		MachineType: machine,
	}

	s := NewComputeService(cfg, nil)

	if s.Config.Project != project {
		t.Error("Project name doesn't match:", s.Config.Project)
	}
	if s.Config.Zone != region {
		t.Error("Zone name doesn't match:", s.Config.Zone)
	}
	if s.Config.MachineType != machine {
		t.Error("Machine type doesn't match:", s.Config.MachineType)
	}

	if s.Log == nil {
		t.Error("Logger is nil")
	}

}

func TestCreateStartupScript(t *testing.T) {

	project := "sample-project"
	bucket := "test-bucket"
	region := "us-central1-c"
	machine := "n1-standard-2"
	cfg := &GcpConfig{
		Project:     project,
		Bucket:      bucket,
		Zone:        region,
		MachineType: machine,
	}
	service := NewComputeService(cfg, nil)

	var err error
	var task script.Script
	var startup string
	name := "test-name"

	task = script.Script{
		Options: []string{},
	}
	startup, err = service.createStartupScript(name, &task)
	if err != nil {
		t.Error(err.Error())
	}
	if !strings.Contains(startup, "$(seq 10)") {
		t.Error("Default retry is not set:", startup)
	}

	task = script.Script{
		Options: []string{
			"no-shutdown",
			"retry:15",
		},
	}
	startup, err = service.createStartupScript(name, &task)
	if err != nil {
		t.Error(err.Error())
	}
	if !strings.Contains(startup, "$(seq 15)") {
		t.Error("Retry is not correct:", startup)
	}
	if !strings.Contains(startup, "--no-shutdown") {
		t.Error("Options are not set:", startup)
	}

	task = script.Script{
		Options: []string{
			"retry:?",
		},
	}
	startup, err = service.createStartupScript(name, &task)
	if err != nil {
		t.Error(err.Error())
	}
	if !strings.Contains(startup, "$(seq 10)") {
		t.Error("Retry is not correct:", startup)
	}

}

// TestReplaceURLScheme checks that function replaces URLs which start with "roadie://".
// to "gs://<bucketname>/.roadie/".
func TestReplaceURLScheme(t *testing.T) {

	project := "sample-project"
	bucket := "test-bucket"
	region := "us-central1-c"
	machine := "n1-standard-2"
	cfg := &GcpConfig{
		Project:     project,
		Bucket:      bucket,
		Zone:        region,
		MachineType: machine,
	}
	service := NewComputeService(cfg, nil)

	task := script.Script{
		Source: "roadie://some-sourcefile",
		Data: []string{
			"roadie://some-datafile",
		},
		Result: "roadie://result-file",
	}
	service.replaceURLScheme(&task)

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
	cfg := &GcpConfig{
		Project:     project,
		Bucket:      bucket,
		Zone:        region,
		MachineType: machine,
	}
	service := NewComputeService(cfg, nil)

	u, err := url.Parse(service.createURL("source", "/path/to/file"))
	if err != nil {
		t.Error(err.Error())
	}
	if u.Scheme != "gs" {
		t.Errorf("Scheme is not correct: %s", u.Scheme)
	}
	if u.Host != bucket {
		t.Errorf("Host name is not correct: %s", u.Host)
	}
	if u.Path != "/.roadie/source/path/to/file" {
		t.Errorf("Path is not correct: %s", u.Path)
	}

}

func TestWait(t *testing.T) {

	select {
	case <-wait(1 * time.Minute):
		t.Fatal("Returned waiting 1min function first.")
	case <-wait(1 * time.Second):
		t.Log("Waiting 1 second returns first.")
	}

}
