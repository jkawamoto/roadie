//
// cloud/resource_manager.go
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

package cloud

import "context"

// ResourceManager defines an interface for configuration to use a cloud service.
type ResourceManager interface {
	// GetProjectID returns an ID of the current project.
	GetProjectID() string
	// SetProjectID sets an ID to the current project.
	SetProjectID(string)
	// GetMachineType returns a machine type the current project uses by default.
	GetMachineType() string
	// SetMachineType sets a machine type as the default one.
	SetMachineType(string)
	// MachineTypes returns a set of available machine types.
	MachineTypes(context.Context) ([]MachineType, error)
	// GetRegion returns a region name the current project working on.
	GetRegion() string
	// SetRegion sets a region to the current project.
	SetRegion(string)
	// Regions returns a set of available regions.
	Regions(context.Context) ([]Region, error)
}
