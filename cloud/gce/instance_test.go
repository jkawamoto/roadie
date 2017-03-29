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
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNormalizedZone(t *testing.T) {

	project := "sample-project"
	region := "us-central1-c"
	machine := "n1-standard-2"
	service := NewComputeService(project, region, machine, ioutil.Discard)

	res := service.normalizedZone()
	if res != fmt.Sprintf("projects/%s/zones/%s", project, region) {
		t.Error("Normalized zone isn's correct:", res)
	}

}

func TestNormalizedMachineType(t *testing.T) {

	project := "sample-project"
	region := "us-central1-c"
	machine := "n1-standard-2"
	service := NewComputeService(project, region, machine, ioutil.Discard)

	res := service.normalizedMachineType()
	if res != fmt.Sprintf("projects/%s/zones/%s/machineTypes/%s", project, region, machine) {
		t.Error("Normalized machine type isn't correct:", res)
	}

}
