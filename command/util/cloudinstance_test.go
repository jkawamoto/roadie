//
// command/util/cloudinstance_test.go
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

package util

import (
	"fmt"
	"testing"
)

func TestNormalizedZone(t *testing.T) {

	project := "sample-project"
	zone := "us-central1-c"
	if res := normalizedZone(project, zone); res != fmt.Sprintf("projects/%s/zones/%s", project, zone) {
		t.Error("Normalized zone isn's correct:", res)
	}

}

func TestNormalizedMachineType(t *testing.T) {

	project := "sample-project"
	zone := "us-central1-c"
	machine := "n1-standard-2"
	if res := normalizedMachineType(project, zone, machine); res != fmt.Sprintf("projects/%s/zones/%s/machineTypes/%s", project, zone, machine) {

		t.Error("Normalized machine type isn't correct:", res)
	}

}
