//
// cloud/mock/instance.go
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

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
)

const (
	// StatusRunning means an instance is still running.
	StatusRunning = "running"
	// StatusTerminated means an instance has been terminated.
	StatusTerminated = "terminated"
)

// InstanceManager is a mock instance manager.
type InstanceManager struct {
	// Failure is set ture, all methods return ErrServiceFailure.
	Failure bool
	// Status is a map to maintain instance's status; each key represents an
	// instance name and the associated value represents its status.
	Status map[string]string
	// Script is a map to maintain scripts: each key represents an instance name
	// and the associated value represents a pointer of the script the instance
	// runs on.
	Script map[string]*script.Script
}

// NewInstanceManager creates a new mock instance manager.
func NewInstanceManager() *InstanceManager {
	return &InstanceManager{
		Status: make(map[string]string),
		Script: make(map[string]*script.Script),
	}
}

// CreateInstance creates an instance which has a given name.
func (m *InstanceManager) CreateInstance(ctx context.Context, s *script.Script) (err error) {

	if m.Failure {
		return ErrServiceFailure
	}
	m.Status[s.Name] = StatusRunning
	m.Script[s.Name] = s
	return

}

// DeleteInstance deletes the given named instance.
func (m *InstanceManager) DeleteInstance(ctx context.Context, name string) (err error) {

	if m.Failure {
		return ErrServiceFailure
	}
	if _, exist := m.Status[name]; !exist {
		return fmt.Errorf("instance %q doesn't exist", name)
	}

	m.Status[name] = StatusTerminated
	delete(m.Script, name)
	return

}

// Instances returns a list of running instances
func (m *InstanceManager) Instances(ctx context.Context, handler cloud.InstanceHandler) (err error) {

	if m.Failure {
		return ErrServiceFailure
	}
	for name, status := range m.Status {
		err = handler(name, status)
		if err != nil {
			return
		}
	}
	return

}
