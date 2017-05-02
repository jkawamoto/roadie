//
// cloud/gcp/resource.go
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
	"context"
	"io/ioutil"
	"log"

	"github.com/jkawamoto/roadie/cloud"
)

// ResourceService is a service to get and set cloud configuration.
type ResourceService struct {
	Config *Config
	Logger *log.Logger
}

// NewResourceService creates a new resource service.
func NewResourceService(cfg *Config, logger *log.Logger) *ResourceService {
	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}
	return &ResourceService{
		Config: cfg,
		Logger: logger,
	}
}

// GetProjectID returns an ID of the current project.
func (m *ResourceService) GetProjectID() string {
	return m.Config.Project
}

// SetProjectID sets an ID to the current project.
func (m *ResourceService) SetProjectID(id string) {
	m.Config.Project = id
}

// GetMachineType returns a machine type the current project uses by default.
func (m *ResourceService) GetMachineType() string {
	return m.Config.MachineType
}

// SetMachineType sets a machine type as the default one.
func (m *ResourceService) SetMachineType(t string) {
	m.Config.MachineType = t
}

// MachineTypes returns a set of available machine types.
func (m *ResourceService) MachineTypes(ctx context.Context) ([]cloud.MachineType, error) {
	c := NewComputeService(m.Config, m.Logger)
	return c.AvailableMachineTypes(ctx)
}

// GetRegion returns a region name the current project working on.
func (m *ResourceService) GetRegion() string {
	return m.Config.Zone
}

// SetRegion sets a region to the current project.
func (m *ResourceService) SetRegion(region string) {
	m.Config.Zone = region
}

// Regions returns a set of available regions.
func (m *ResourceService) Regions(ctx context.Context) ([]cloud.Region, error) {
	c := NewComputeService(m.Config, m.Logger)
	return c.AvailableRegions(ctx)
}
