//
// cloud/mock/resource.go
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

	"github.com/jkawamoto/roadie/cloud"
)

// ResourceManager is a mock service implementing cloud.ResourceManager interface.
type ResourceManager struct {
	// Project ID.
	projectID string
	// Machine type.
	machineType string
	// Region
	region string
	// Available machine types
	AvailableMachineTypes []cloud.MachineType
	// Available regions
	AvailableRegions []cloud.Region
	// Failure if set true, MachineTypes and Regions will return ErrServiceFailure.
	Failure bool
}

// NewResourceManager returns a mock resource manager.
func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}

// GetProjectID returns an ID of the current project.
func (m *ResourceManager) GetProjectID() string {
	return m.projectID
}

// SetProjectID sets an ID to the current project.
func (m *ResourceManager) SetProjectID(v string) {
	m.projectID = v
}

// GetMachineType returns a machine type the current project uses by default.
func (m *ResourceManager) GetMachineType() string {
	return m.machineType
}

// SetMachineType sets a machine type as the default one.
func (m *ResourceManager) SetMachineType(v string) {
	m.machineType = v
}

// MachineTypes returns a set of available machine types.
func (m *ResourceManager) MachineTypes(ctx context.Context) ([]cloud.MachineType, error) {

	if m.Failure {
		return nil, ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return m.AvailableMachineTypes, nil
	}

}

// GetRegion returns a region name the current project working on.
func (m *ResourceManager) GetRegion() string {
	return m.region
}

// SetRegion sets a region to the current project.
func (m *ResourceManager) SetRegion(v string) {
	m.region = v
}

// Regions returns a set of available regions.
func (m *ResourceManager) Regions(ctx context.Context) ([]cloud.Region, error) {

	if m.Failure {
		return nil, ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return m.AvailableRegions, nil
	}

}
