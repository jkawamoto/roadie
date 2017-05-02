//
// cloud/gcp/config_test.go
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
	"fmt"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestUnmarshalYAML(t *testing.T) {

	var cfg Config
	var res Config

	cfg = Config{
		Project:     "sample-project",
		Bucket:      "sample-bucket",
		Zone:        "sample-zone",
		MachineType: "sample-machine-type",
	}
	data, err := yaml.Marshal(&cfg)
	if err != nil {
		t.Error(err.Error())
	}
	yaml.Unmarshal(data, &res)

	if res.Project != cfg.Project {
		t.Error("Project doesn't match:", res.Project)
	}
	if res.Bucket != cfg.Bucket {
		t.Error("Bucket doesn't match:", res.Bucket)
	}
	if res.Zone != cfg.Zone {
		t.Error("Zone doesn't match:", res.Zone)
	}
	if res.MachineType != cfg.MachineType {
		t.Error("MachineType doesn't match:", res.MachineType)
	}

	cfg = Config{
		Project: "sample-project",
		Bucket:  "sample-bucket",
	}
	data, err = yaml.Marshal(&cfg)
	if err != nil {
		t.Error(err.Error())
	}
	yaml.Unmarshal(data, &res)

	if res.Project != cfg.Project {
		t.Error("Project doesn't match:", res.Project)
	}
	if res.Bucket != cfg.Bucket {
		t.Error("Bucket doesn't match:", res.Bucket)
	}
	if res.Zone != DefaultZone {
		t.Error("Zone doesn't match:", res.Zone)
	}
	if res.MachineType != DefaultMachineType {
		t.Error("MachineType doesn't match:", res.MachineType)
	}

}

func TestNormalizedZone(t *testing.T) {

	project := "sample-project"
	region := "us-central1-c"
	machine := "n1-standard-2"
	cfg := Config{
		Project:     project,
		Zone:        region,
		MachineType: machine,
	}

	res := cfg.normalizedZone()
	if res != fmt.Sprintf("projects/%s/zones/%s", project, region) {
		t.Error("Normalized zone isn's correct:", res)
	}

}

func TestNormalizedMachineType(t *testing.T) {

	project := "sample-project"
	region := "us-central1-c"
	machine := "n1-standard-2"
	cfg := Config{
		Project:     project,
		Zone:        region,
		MachineType: machine,
	}

	res := cfg.normalizedMachineType()
	if res != fmt.Sprintf("projects/%s/zones/%s/machineTypes/%s", project, region, machine) {
		t.Error("Normalized machine type isn't correct:", res)
	}

}

func TestDiskType(t *testing.T) {

	project := "sample-project"
	region := "us-central1-c"
	machine := "n1-standard-2"
	cfg := Config{
		Project:     project,
		Zone:        region,
		MachineType: machine,
	}

	res := cfg.diskType()
	if res != fmt.Sprintf("projects/%s/zones/%s/diskTypes/pd-standard", project, region) {
		t.Error("Disk type isn't correct:", res)
	}

}

func TestNetwork(t *testing.T) {

	project := "sample-project"
	region := "us-central1-c"
	machine := "n1-standard-2"
	cfg := Config{
		Project:     project,
		Zone:        region,
		MachineType: machine,
	}

	res := cfg.network()
	if res != fmt.Sprintf("projects/%s/global/networks/default", project) {
		t.Error("Network isn't correct:", res)
	}

}
