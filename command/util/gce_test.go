//
// command/util/gce_test.go
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
	"os"
	"testing"
)

// func TestCreateInstance(t *testing.T) {
//
// 	b, err := NewInstanceBuilder("jkawamoto-ppls")
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
//
// 	if err := b.CreateInstance("test-instance"); err != nil {
// 		t.Error(err.Error())
// 	}
//
// 	if err := b.StopInstance("test-instance"); err != nil {
// 		t.Error(err.Error())
// 	}
//
// }

func TestAvailableZones(t *testing.T) {

	id := os.Getenv("PROJECT_ID")
	if id == "" {
		t.Log("Skip this test because no project id is given.")
		return
	}

	b, err := NewInstanceBuilder(id)
	if err != nil {
		t.Error(err.Error())
	}

	zones, err := b.AvailableZones()
	if err != nil {
		t.Error(err.Error())
	}

	for _, v := range zones {
		t.Logf("Available zone: %s", v)
	}

}

func TestAvailableMachineTypes(t *testing.T) {

	id := os.Getenv("PROJECT_ID")
	if id == "" {
		t.Log("Skip this test because no project id is given.")
		return
	}

	b, err := NewInstanceBuilder(id)
	if err != nil {
		t.Error(err.Error())
	}

	types, err := b.AvailableMachineTypes()
	if err != nil {
		t.Error(err.Error())
	}

	for _, v := range types {
		t.Logf("Available machine type: %s", v)
	}

}

func TestNormalizedZone(t *testing.T) {

	id := os.Getenv("PROJECT_ID")
	if id == "" {
		t.Log("Skip this test because no project id is given.")
		return
	}

	b, err := NewInstanceBuilder(id)
	if err != nil {
		t.Error(err.Error())
	}

	b.Zone = "us-central1-c"
	if b.normalizedZone() != fmt.Sprintf("projects/%s/zones/us-central1-c", id) {
		t.Errorf("Zone is not correct: %s", b.Zone)
	}

}

func TestNormalizedMachineType(t *testing.T) {

	id := os.Getenv("PROJECT_ID")
	if id == "" {
		t.Log("Skip this test because no project id is given.")
		return
	}

	b, err := NewInstanceBuilder(id)
	if err != nil {
		t.Error(err.Error())
	}

	b.MachineType = "n1-standard-2"
	if b.normalizedMachineType() != fmt.Sprintf("projects/%s/zones/us-central1-b/machineTypes/n1-standard-2", id) {
		t.Errorf("Zone is not correct: %s", b.Zone)
	}

}
