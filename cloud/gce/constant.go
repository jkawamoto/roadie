//
// cloud/gce/doc.go
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

import compute "google.golang.org/api/compute/v1"

const (
	// GCP's scope.
	gceScope = compute.CloudPlatformScope
	// StoragePrefix is a prefix used to store related data into the cloud
	// storage.
	StoragePrefix = ".roadie"
	// DefaultZone defines the default zone.
	DefaultZone = "us-central1-b"
	// DefaultMachineType defines the default machine type.
	DefaultMachineType = "n1-standard-1"

	// LogTimeFormat defines time format of Google Logging.
	LogTimeFormat = "2006-01-02T15:04:05Z"
	// LogEventSubtypeInsert means this event is creating an instance.
	LogEventSubtypeInsert = "compute.instances.insert"
	// LogEventSubtypeDelete means this event is deleting an instance.
	LogEventSubtypeDelete = "compute.instances.delete"
)
