//
// cloud/mock/instance_test.go
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
	"fmt"
	"testing"

	"github.com/jkawamoto/roadie/script"
)

func TestCreateInstance(t *testing.T) {

	m := NewInstanceManager()
	if len(m.Status) != 0 {
		t.Fatalf("initializing an instance manager is failed, has %v items", len(m.Status))
	}

	var err error
	ctx := context.Background()
	s := script.Script{
		Name: "test-instance",
	}

	err = m.CreateInstance(ctx, &s)
	if err != nil {
		t.Fatalf("CreateInstance returns an error: %v", err)
	}

	if len(m.Status) != 1 || m.Status[s.Name] != StatusRunning {
		t.Errorf("created instance's status is %q, want %v", m.Status[s.Name], StatusRunning)
	}
	if len(m.Script) != 1 || m.Script[s.Name] != &s {
		t.Errorf("script of the created instance is %+v, want %+v", m.Script[s.Name], s)
	}

	// Failure test.
	m.Failure = true
	err = m.CreateInstance(ctx, &s)
	if err == nil {
		t.Error("Failuer is true but no errors are returned")
	}

}

func TestDeleteInstance(t *testing.T) {

	m := NewInstanceManager()
	if len(m.Status) != 0 {
		t.Fatalf("initializing an instance manager is failed, has %v items", len(m.Status))
	}

	var err error
	ctx := context.Background()
	s := script.Script{
		Name: "test-instance",
	}
	err = m.CreateInstance(ctx, &s)
	if err != nil {
		t.Fatalf("CreateInstance returns an error: %v", err)
	}
	if len(m.Status) != 1 || m.Status[s.Name] != StatusRunning {
		t.Fatalf("created instance's status is %q, want %v", m.Status[s.Name], StatusRunning)
	}
	if len(m.Script) != 1 || m.Script[s.Name] != &s {
		t.Fatalf("script of the created instance is %+v, want %+v", m.Script[s.Name], s)
	}

	err = m.DeleteInstance(ctx, s.Name)
	if err != nil {
		t.Fatalf("DeleteInstance returns an error: %v", err)
	}
	if m.Status[s.Name] != StatusTerminated {
		t.Errorf("status of the deleted instance is %q, want %v", m.Status[s.Name], StatusTerminated)
	}
	if _, exist := m.Script[s.Name]; exist {
		t.Error("script of the deleted instance is still remining")
	}

	// Delete not existing instance.
	err = m.DeleteInstance(ctx, "another_instance")
	if err == nil {
		t.Error("deleting a not existing instance but no errors are returned")
	}

	// Failure test.
	m.Failure = true
	err = m.DeleteInstance(ctx, s.Name)
	if err == nil {
		t.Error("Failuer is true but no errors are returned")
	}

}

func TestInstances(t *testing.T) {

	m := NewInstanceManager()
	if len(m.Status) != 0 {
		t.Fatalf("initializing an instance manager is failed, has %v items", len(m.Status))
	}

	instances := map[string]string{
		"instance1":  StatusRunning,
		"instance2":  StatusRunning,
		"instance3":  StatusRunning,
		"instance11": StatusTerminated,
		"instance12": StatusTerminated,
	}
	for name, status := range instances {
		m.Status[name] = status
	}

	var err error
	var c int
	err = m.Instances(context.Background(), func(name, status string) error {
		if instances[name] != status {
			t.Errorf("status of %q is %v, want %v", name, status, instances[name])
		}
		c++
		return nil
	})
	if err != nil {
		t.Fatalf("Instances returns an error: %v", err)
	}
	if c != len(instances) {
		t.Errorf("the number of found instances %v, want %v", c, len(instances))
	}

	// Handler returns an error.
	expected := fmt.Errorf("expected error")
	err = m.Instances(context.Background(), func(name, status string) error {
		return expected
	})
	if err != expected {
		t.Errorf("handler returns an error but Instances returns a different error: %v", err)
	}

	// Failure test.
	m.Failure = true
	err = m.Instances(context.Background(), func(name, status string) error {
		return nil
	})
	if err == nil {
		t.Error("Failuer is true but no errors are returned")
	}

}
