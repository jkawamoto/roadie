//
// cloud/gce/config.go
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

package gce

import (
	"path/filepath"
)

// TODO: Rename to "Config".
// GcpConfig defines information to access Google Cloud Platform.
type GcpConfig struct {
	// Project name.
	Project string `yaml:"project"`
	// Bucket name
	Bucket string `yaml:"bucket"`
	// Zone where instances will run.
	Zone string `yaml:"zone"`
	// Default machine type of new instances.
	MachineType string `yaml:"machine_type"`
	// Instance disk size.
	DiskSize int64 `yaml:"disk_size,omitempty"`
}

// UnmarshalYAML helps to unmarshal GcpConfig objects.
func (cfg *GcpConfig) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {

	type AuxGcpConfig struct {
		Project     string `yaml:"project"`
		Bucket      string `yaml:"bucket"`
		Zone        string `yaml:"zone"`
		MachineType string `yaml:"machine_type"`
		DiskSize    int64  `yaml:"disk_size,omitempty"`
	}

	aux := AuxGcpConfig{}
	err = unmarshal(&aux)
	if err != nil {
		return
	}

	cfg.Project = aux.Project
	cfg.Bucket = aux.Bucket

	if aux.Zone != "" {
		cfg.Zone = aux.Zone
	} else {
		cfg.Zone = DefaultZone
	}

	if aux.MachineType != "" {
		cfg.MachineType = aux.MachineType
	} else {
		cfg.MachineType = DefaultMachineType
	}

	if aux.DiskSize != 0 {
		cfg.DiskSize = aux.DiskSize
	} else {
		cfg.DiskSize = DefaultDiskSize
	}

	return

}

// normalizedZone returns the normalized zone string of Zone property.
func (cfg *GcpConfig) normalizedZone() string {
	return filepath.Join("projects", cfg.Project, "zones", cfg.Zone)
}

// normalizedMachineType returns the normalized instance type of MachineType property.
func (cfg *GcpConfig) normalizedMachineType() string {
	return filepath.Join(cfg.normalizedZone(), "machineTypes", cfg.MachineType)
}

// diskType returns default disk type.
func (cfg *GcpConfig) diskType() string {
	return filepath.Join(cfg.normalizedZone(), "/diskTypes/pd-standard")
}

// network returns default network name
func (cfg *GcpConfig) network() string {
	return filepath.Join("projects", cfg.Project, "/global/networks/default")
}
