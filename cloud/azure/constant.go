//
// cloud/azure/constant.go
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

package azure

import "time"

const (
	// DefaultOSPublisherName is the default publisher name of the default OS image.
	DefaultOSPublisherName = "Canonical"
	// DefaultOSOffer is the default offer of the default OS image.
	DefaultOSOffer = "UbuntuServer"
	// DefaultOSSkus is the default skus of the default OS image.
	DefaultOSSkus = "17.04"
	// DefaultOSVersion is the default version of the default version.
	DefaultOSVersion = "latest"
	// DefaultLocation is the default location.
	DefaultLocation = "westus"

	// DefaultSleepTime is the default sleeping time to wait creating or deleting
	// objects.
	DefaultSleepTime = 30 * time.Second

	// ComputeServiceDefaultMachineType defines the default machine type.
	ComputeServiceDefaultMachineType = "Standard_A2"
	// ComputeServiceCustomScriptExtension defines the name of custom script
	// extention.
	ComputeServiceCustomScriptExtension = "CustomScriptForLinux"

	// BinContainer is the name of the container where binary programs will be stored.
	BinContainer = "bin"
	// SourceContainer is the name of the container where source files will be stored.
	SourceContainer = "source"
	// DataContainer is the name of the container where data files will be stored.
	DataContainer = "data"
	// StartupContainer is the name of the container where startup files will be stored.
	StartupContainer = "startup"
	// ResultContainer is the name of the container where result files will be stored.
	ResultContainer = "result"
	// LogContainer is the name of the container where logs will be stored.
	LogContainer = "log"

	// QueuePrefix is a prefix a job which is working for a queue must has.
	QueuePrefix = "queue-"
	// TaskPrefix is a prefix a task which is working for a queue must has.
	TaskPrefix = "task-"
)
