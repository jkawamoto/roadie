//
// cloud/mock/resource_test.go
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

package mock

import (
	"context"
	"testing"

	"github.com/jkawamoto/roadie/cloud"
)

func TestGetProjectID(t *testing.T) {

	expected := "test id"
	m := NewResourceManager()
	m.projectID = expected

	if res := m.GetProjectID(); res != expected {
		t.Errorf("GetProjectID returns %q, want %v", res, expected)
	}

}

func TestSetProjectID(t *testing.T) {

	expected := "test id"
	m := NewResourceManager()
	m.projectID = "another id"

	m.SetProjectID(expected)
	if m.projectID != expected {
		t.Errorf("SetProjectID sets %q, want %v", m.projectID, expected)
	}

}

func TestGetMatchineType(t *testing.T) {

	expected := "test type"
	m := NewResourceManager()
	m.machineType = expected

	if res := m.GetMachineType(); res != expected {
		t.Errorf("GetMachineType returns %q, want %v", res, expected)
	}

}

func TestSetMachineType(t *testing.T) {

	expected := "test type"
	m := NewResourceManager()
	m.machineType = "another type"

	m.SetMachineType(expected)
	if m.machineType != expected {
		t.Errorf("SetMachineType sets %q, want %v", m.machineType, expected)
	}

}

func TestMachineTypes(t *testing.T) {

	testTypes := map[string]string{
		"type A": "some machine type",
		"type B": "another machine type",
	}

	m := NewResourceManager()
	for k, v := range testTypes {
		m.AvailableMachineTypes = append(m.AvailableMachineTypes, cloud.MachineType{
			Name:        k,
			Description: v,
		})
	}

	ctx := context.Background()
	res, err := m.MachineTypes(ctx)
	if err != nil {
		t.Fatalf("MachineTypes returns an error: %v", err)
	}
	if len(res) != len(testTypes) {
		t.Errorf("the number of machine types is %v, want %v", len(res), len(testTypes))
	}
	for _, v := range res {
		if desc := testTypes[v.Name]; desc != v.Description {
			t.Errorf("descrption of %q is %q, want %q", v.Name, v.Description, desc)
		}
	}

	// With failure option.
	m.Failure = true
	_, err = m.MachineTypes(ctx)
	if err == nil {
		t.Error("failure is true but no errors are returned")
	}
	m.Failure = false

	// With a canceled context.
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	_, err = m.MachineTypes(ctx)
	if err == nil {
		t.Error("context is canceled but no errors are returned")
	}

}

func TestGetRegion(t *testing.T) {

	test := "test region"
	m := NewResourceManager()
	m.region = test

	if res := m.GetRegion(); res != test {
		t.Errorf("GetRegion returns %q, want %q", res, test)
	}

}

func TestSetRegion(t *testing.T) {

	test := "test region"
	m := NewResourceManager()
	m.region = "another region"

	m.SetRegion(test)
	if m.region != test {
		t.Errorf("SetRegion sets %q, want %q", m.region, test)
	}

}

func TestRegions(t *testing.T) {

	testRegions := map[string]string{
		"region A": "up",
		"region B": "up",
		"region C": "down",
	}

	m := NewResourceManager()
	for k, v := range testRegions {
		m.AvailableRegions = append(m.AvailableRegions, cloud.Region{
			Name:   k,
			Status: v,
		})
	}

	ctx := context.Background()
	res, err := m.Regions(ctx)
	if err != nil {
		t.Fatalf("Regions returns an error: %v", err)
	}
	if len(res) != len(testRegions) {
		t.Errorf("the number of regions is %v, want %v", len(res), len(testRegions))
	}
	for _, v := range res {
		if status := testRegions[v.Name]; status != v.Status {
			t.Errorf("status of %q is %q, want %q", v.Name, v.Status, status)
		}
	}

	// With failure option.
	m.Failure = true
	_, err = m.Regions(ctx)
	if err == nil {
		t.Error("failure is true but no errors are returned")
	}
	m.Failure = false

	// With a canceled context.
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	_, err = m.Regions(ctx)
	if err == nil {
		t.Error("context is canceled but no errors are returned")
	}

}
