//
// command/config_test.go
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

package command

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/cloud/mock"
)

var (
	testMachineTypes = []cloud.MachineType{
		{Name: "type A", Description: "sample machine type"},
		{Name: "type B", Description: "another machine type"},
	}
	testRegions = []cloud.Region{
		{Name: "region A", Status: "up"},
		{Name: "region B", Status: "down"},
	}
)

func TestCmdConfigProjectSet(t *testing.T) {

	var err error
	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)

	resource, err := p.ResourceManager(m.Context)
	if err != nil {
		t.Fatalf("ResourceManager returns an error: %v", err)
	}

	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("cannot create a temp file: %v", err)
	}
	f.Close()
	m.Config.FileName = f.Name()
	defer os.Remove(f.Name())

	cases := []struct {
		input  string
		expect string
		output string
	}{
		{"sample1", "sample1", "sample1"},
		{"sample2", "sample2", "sample1 -> sample2"},
		{"sample-3", "sample-3", "sample2 -> sample-3"},
		{"sample 4", "sample_4", "sample-3 -> sample_4"},
	}

	for _, c := range cases {

		err = cmdConfigProjectSet(m, c.input)
		if err != nil {
			t.Fatalf("cmdConfigProjectSet returns an error: %v", err)
		}
		if id := resource.GetProjectID(); id != c.expect {
			t.Errorf("updated project name is %q, want %q", id, c.expect)
		}
		lines := strings.Split(strings.TrimRight(output.String(), "\n"), "\n")
		if len(lines) < 2 {
			t.Errorf("output message is too short: %v", lines)
		}
		if output := strings.TrimSpace(lines[len(lines)-1]); output != c.output {
			t.Errorf("output message is %q, want %q", output, c.output)
		}
		output.Reset()

	}

	// With a wrong file name
	m.Config.FileName = ""
	err = cmdConfigProjectSet(m, "sample 1")
	if err == nil {
		t.Error("a wrong config file path is used but no errors are returned")
	}

}

func TestCmdConfigProjectShow(t *testing.T) {

	var err error
	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)

	resource, err := m.ResourceManager()
	if err != nil {
		t.Fatalf("ResourceManager returns an error: %v", err)
	}

	cases := []struct {
		set    string
		expect string
	}{
		{"", MsgNotSet},
		{"test-id", "test-id"},
	}

	for _, c := range cases {

		resource.SetProjectID(c.set)
		err = cmdConfigProjectShow(m)
		if err != nil {
			t.Fatalf("cmdConfigProjectShow returns an error: %v", err)
		}
		if res := strings.TrimSpace(output.String()); res != c.expect {
			t.Errorf("printed %q, want %q", res, c.expect)
		}
		output.Reset()

	}

}

func TestConfigMachineTypeSet(t *testing.T) {

	var err error
	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)
	p.MockResourceManager.AvailableMachineTypes = testMachineTypes

	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("cannot create a temp file: %v", err)
	}
	f.Close()
	m.Config.FileName = f.Name()
	defer os.Remove(f.Name())

	cases := []struct {
		input  string
		expect string
	}{
		{"type A", "type A"},
		{"type B", "type A -> type B"},
	}

	for _, c := range cases {
		err = cmdConfigMachineTypeSet(m, c.input)
		if err != nil {
			t.Fatalf("cmdConfigMachineTypeSet with %q returns an error: %v", c.input, err)
		}
		lines := strings.Split(strings.TrimRight(output.String(), "\n"), "\n")
		if len(lines) < 2 {
			t.Errorf("output message is too short: %v", lines)
		}
		if res := strings.TrimSpace(lines[len(lines)-1]); res != c.expect {
			t.Errorf("output message is %q, want %q", res, c.expect)
		}
		output.Reset()
	}

	// With a wrong machine type.
	err = cmdConfigMachineTypeSet(m, "type C")
	if err == nil {
		t.Fatal("set a wrong machine type but no errors are returned")
	}

	// With an out-of-service resource manager.
	p.MockResourceManager.Failure = true
	err = cmdConfigMachineTypeSet(m, "type A")
	if err == nil {
		t.Error("failure is true but no errors are returned")
	}
	p.MockResourceManager.Failure = false

	// With a wrong file name
	m.Config.FileName = ""
	err = cmdConfigMachineTypeSet(m, "type A")
	if err == nil {
		t.Error("a wrong config file path is used but no errors are returned")
	}

}

func TestCmdConfigMachineTypeList(t *testing.T) {

	var err error
	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)
	p.MockResourceManager.AvailableMachineTypes = testMachineTypes
	p.MockResourceManager.SetMachineType("type A")

	err = cmdConfigMachineTypeList(m)
	if err != nil {
		t.Fatalf("cmdConfigMachineTypeList returns an error: %v", err)
	}

	lines := strings.Split(output.String(), "\n")
	if !strings.Contains(lines[0], "MACHINE TYPE") || !strings.Contains(lines[0], "DESCRIPTION") {
		t.Errorf("a wrong table header: %v", lines[0])
	}
	if len(lines) < 2 {
		t.Fatalf("output message is too short: %v", lines)
	}
	for _, row := range lines[1:] {
		if row == "" {
			continue
		}
		kv := strings.Split(row, "\t")
		if len(kv) < 2 {
			t.Fatalf("table row doesn't have enough information: %v", kv)
		}
		if name := strings.TrimSpace(kv[0]); name != "type A*" && name != "type B" {
			t.Errorf("machine type name is incorrect: %q", name)
		}
		if desc := strings.TrimSpace(kv[1]); desc != "sample machine type" && desc != "another machine type" {
			t.Errorf("machine description is incorrect: %q", desc)
		}
	}

	// With an out-of-service resource manager.
	p.MockResourceManager.Failure = true
	err = cmdConfigMachineTypeList(m)
	if err == nil {
		t.Error("resource manager is out-of-service but no errors are returned")
	}
	p.MockResourceManager.Failure = false

	// With a canceled context.
	var cancel context.CancelFunc
	m.Context, cancel = context.WithCancel(context.Background())
	cancel()
	if cmdConfigMachineTypeList(m) == nil {
		t.Error("context is canceled but no errors are returned")
	}

}

func TestCmdConfigMachineTypeShow(t *testing.T) {

	var err error
	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)

	resource, err := m.ResourceManager()
	if err != nil {
		t.Fatalf("ResourceManager returns an error: %v", err)
	}

	cases := []struct {
		set    string
		expect string
	}{
		{"", MsgNotSet},
		{"test-type", "test-type"},
	}

	for _, c := range cases {

		resource.SetMachineType(c.set)
		err = cmdConfigMachineTypeShow(m)
		if err != nil {
			t.Fatalf("cmdConfigMachineTypeShow returns an error: %v", err)
		}
		if res := strings.TrimSpace(output.String()); res != c.expect {
			t.Errorf("printed %q, want %q", res, MsgNotSet)
		}
		output.Reset()

	}

}

func TestConfigRegionSet(t *testing.T) {

	var err error
	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)
	p.MockResourceManager.AvailableRegions = testRegions

	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("cannot create a temp file: %v", err)
	}
	f.Close()
	m.Config.FileName = f.Name()
	defer os.Remove(f.Name())

	cases := []struct {
		input  string
		expect string
	}{
		{"region A", "region A"},
		{"region B", "region A -> region B"},
	}

	for _, c := range cases {
		err = cmdConfigRegionSet(m, c.input)
		if err != nil {
			t.Fatalf("cmdConfigRegionSet with %q returns an error: %v", c.input, err)
		}
		lines := strings.Split(strings.TrimRight(output.String(), "\n"), "\n")
		if len(lines) < 2 {
			t.Errorf("output message is too short: %v", lines)
		}
		if res := strings.TrimSpace(lines[len(lines)-1]); res != c.expect {
			t.Errorf("output message is %q, want %q", res, c.expect)
		}
		output.Reset()
	}

	// With a wrong machine type.
	err = cmdConfigRegionSet(m, "region C")
	if err == nil {
		t.Fatal("set a wrong machine type but no errors are returned")
	}

	// With an out-of-service resource manager.
	p.MockResourceManager.Failure = true
	err = cmdConfigRegionSet(m, "region A")
	if err == nil {
		t.Error("failure is true but no errors are returned")
	}
	p.MockResourceManager.Failure = false

	// With a wrong file name
	m.Config.FileName = ""
	err = cmdConfigRegionSet(m, "region A")
	if err == nil {
		t.Error("a wrong config file path is used but no errors are returned")
	}

}

func TestCmdConfigRegionList(t *testing.T) {

	var err error
	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)
	p.MockResourceManager.AvailableRegions = testRegions
	p.MockResourceManager.SetRegion("region A")

	err = cmdConfigRegionList(m)
	if err != nil {
		t.Fatalf("cmdConfigRegionList returns an error: %v", err)
	}

	lines := strings.Split(output.String(), "\n")
	if !strings.Contains(lines[0], "REGION") || !strings.Contains(lines[0], "STATUS") {
		t.Errorf("a wrong table header: %v", lines[0])
	}
	if len(lines) < 2 {
		t.Fatalf("output message is too short: %v", lines)
	}
	for _, row := range lines[1:] {
		if row == "" {
			continue
		}
		kv := strings.Split(row, "\t")
		if len(kv) < 2 {
			t.Fatalf("table row doesn't have enough information: %v", kv)
		}
		if name := strings.TrimSpace(kv[0]); name != "region A*" && name != "region B" {
			t.Errorf("region name is incorrect: %q", name)
		}
		if desc := strings.TrimSpace(kv[1]); desc != "up" && desc != "down" {
			t.Errorf("region status is incorrect: %q", desc)
		}
	}

	// With an out-of-service resource manager.
	p.MockResourceManager.Failure = true
	err = cmdConfigRegionList(m)
	if err == nil {
		t.Error("resource manager is out-of-service but no errors are returned")
	}
	p.MockResourceManager.Failure = false

	// With a canceled context.
	var cancel context.CancelFunc
	m.Context, cancel = context.WithCancel(context.Background())
	cancel()
	if cmdConfigRegionList(m) == nil {
		t.Error("context is canceled but no errors are returned")
	}

}

func TestCmdConfigRegionShow(t *testing.T) {

	var err error
	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)

	resource, err := m.ResourceManager()
	if err != nil {
		t.Fatalf("ResourceManager returns an error: %v", err)
	}

	cases := []struct {
		set    string
		expect string
	}{
		{"", MsgNotSet},
		{"test-region", "test-region"},
	}

	for _, c := range cases {

		resource.SetRegion(c.set)
		err = cmdConfigRegionShow(m)
		if err != nil {
			t.Fatalf("cmdConfigRegionShow returns an error: %v", err)
		}
		if res := strings.TrimSpace(output.String()); res != c.expect {
			t.Errorf("printed %q, want %q", res, MsgNotSet)
		}
		output.Reset()

	}

}
