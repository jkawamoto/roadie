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
	DefaultOSSkus = "16.10"
	// DefaultOSVersion is the default version of the default version.
	DefaultOSVersion = "latest"
	// DefaultStorageAccount is the default storage account name.
	DefaultStorageAccount = "roadie"
	// DefaultBatchAccount is the default batch account name.
	DefaultBatchAccount = "roadie"
	// DefaultSleepTime is the default sleeping time to wait creating or deleting
	// objects.
	DefaultSleepTime = 30 * time.Second

	// ComputeServiceResourceGroupName defines the default resource name.
	ComputeServiceResourceGroupName = "roadie"
	// ComputeServiceDefaultMachineType defines the default machine type.
	ComputeServiceDefaultMachineType = "Standard_A2"
	// ComputeServiceCustomScriptExtension defines the name of custom script
	// extention.
	ComputeServiceCustomScriptExtension = "CustomScriptForLinux"

	// ProvisioningSucceeded defines succeeded provisioning status
	ProvisioningSucceeded = "Succeeded"
)
