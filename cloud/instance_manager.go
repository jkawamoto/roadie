//
// cloud/instance_manager.go
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

type InstanceManager interface {

	// CreateInstance creates an instance which has a given name.
	CreateInstance(ctx context.Context, name string, metadata []*MetadataItem, disksize int64) error

	// DeleteInstance deletes the given named instance.
	DeleteInstance(ctx context.Context, name string) error

	// Instances returns a list of running instances
	Instances(context.Context) (instances map[string]struct{}, err error)

	// AvailableRegions returns a list of available regions.
	AvailableRegions(context.Context) (zones []Region, err error)

	// AvailableMachineTypes returns a list of available machine types.
	AvailableMachineTypes(context.Context) (types []MachineType, err error)
}

// MetadataItem has Key and Value properties.
type MetadataItem struct {
	Key   string
	Value string
}

// MachineType defines a structure of machine type infoemation.
type MachineType struct {
	Name        string
	Description string
}

// Region defines a structure of region information.
type Region struct {
	Name   string
	Status string
}
