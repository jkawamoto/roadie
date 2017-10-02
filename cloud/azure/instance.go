//
// cloud/azure/instance.go
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

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/jkawamoto/roadie/cloud"
	client "github.com/jkawamoto/roadie/cloud/azure/compute/client"
	"github.com/jkawamoto/roadie/cloud/azure/compute/client/virtual_machine_images"
	"github.com/jkawamoto/roadie/cloud/azure/compute/client/virtual_machine_sizes"
	"github.com/jkawamoto/roadie/cloud/azure/compute/client/virtual_machines"
	"github.com/jkawamoto/roadie/cloud/azure/compute/models"
	"github.com/jkawamoto/roadie/script"
)

const (
	// ComputeAPIVersion defines API version of compute service.
	ComputeAPIVersion = "2016-04-30-preview"
)

// ComputeService provides an interface for Azure's compute service.
type ComputeService struct {
	client    *client.ComputeManagementClient
	Config    *AzureConfig
	Logger    *log.Logger
	SleepTime time.Duration
}

// Entry is an entry of lists.
type Entry struct {
	ID   string
	Name string
}

// NewComputeService creates a new compute service interface assosiated with
// a given configuration.
func NewComputeService(ctx context.Context, cfg *AzureConfig, logger *log.Logger) (s *ComputeService, err error) {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	// Create a resource group if not exist.
	err = CreateResourceGroupIfNotExist(ctx, cfg, logger)
	if err != nil {
		return
	}

	// Create a management client.
	cli := client.NewHTTPClient(strfmt.NewFormats())
	s = &ComputeService{
		client:    cli,
		Config:    cfg,
		Logger:    logger,
		SleepTime: DefaultSleepTime,
	}
	return

}

// CreateInstance creates an instance which has a given name.
func (s *ComputeService) CreateInstance(ctx context.Context, name string, script *script.Script, disksize int64) (err error) {

	if script.Name == "" {
		script.Name = name
	}

	startup, err := StartupScript(s.Config, script)
	if err != nil {
		return
	}

	// Create dependent service interfaces.
	networkService, err := NewNetworkService(ctx, s.Config, s.Logger)
	if err != nil {
		return
	}
	diskService, err := NewDiskService(ctx, s.Config, s.Logger)
	if err != nil {
		return
	}

	// Create a network service.
	nif, err := networkService.CreateNetworkInterface(ctx, s.networkInterfaceName(name))
	if err != nil {
		return
	}

	osDiskName := s.osDiskName(name)
	param := &models.VirtualMachine{
		Resource: models.Resource{
			Location: &s.Config.Location,
		},
		Properties: &models.VirtualMachineProperties{
			HardwareProfile: &models.HardwareProfile{
				VMSize: s.Config.MachineType,
			},
			StorageProfile: &models.StorageProfile{
				ImageReference: &models.ImageReference{
					Publisher: s.Config.OS.PublisherName,
					Offer:     s.Config.OS.Offer,
					Sku:       s.Config.OS.Skus,
					Version:   s.Config.OS.Version,
				},
				OsDisk: &models.OSDisk{
					Name:         osDiskName,
					Caching:      models.CachingReadOnly,
					CreateOption: models.CreateOptionFromImage,
					DiskSizeGB:   int32(disksize),
				},
				DataDisks: []*models.DataDisk{},
			},
			OsProfile: &models.OSProfile{
				ComputerName:  name,
				AdminUsername: "roadie",
				AdminPassword: "pass2roadie-A",
				CustomData:    startup,
				LinuxConfiguration: &models.LinuxConfiguration{
					DisablePasswordAuthentication: false,
				},
			},
			NetworkProfile: &models.NetworkProfile{
				NetworkInterfaces: []*models.NetworkInterfaceReference{
					&models.NetworkInterfaceReference{
						SubResource: models.SubResource{
							ID: nif.ID,
						},
						Properties: &models.NetworkInterfaceReferenceProperties{
							Primary: true,
						},
					},
				},
			},
			DiagnosticsProfile: &models.DiagnosticsProfile{
				BootDiagnostics: &models.BootDiagnostics{
					Enabled: false,
				},
			},
		},
	}

	s.Logger.Println("Creating virtuan machine", name)
	creating, created, err := s.client.VirtualMachines.VirtualMachinesCreateOrUpdate(
		virtual_machines.NewVirtualMachinesCreateOrUpdateParamsWithContext(ctx).
			WithAPIVersion(ComputeAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithVMName(name).
			WithParameters(param), httptransport.BearerToken(s.Config.Token.AccessToken))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case created != nil || creating != nil:
		var info *models.VirtualMachine
		s.Logger.Println("Waiting for creating virtual machine", name)
		for {
			info, err = s.GetInstanceInfo(ctx, name)
			if err != nil {
				break
			} else if info.Properties.ProvisioningState == "Succeeded" {
				break
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-time.After(s.SleepTime):
			}
		}

	default:
		err = fmt.Errorf("Unexpected case has occurred")

	}

	// Waiting creation of the OS disk
	var diskSet DiskSet
	for {
		if diskSet, err = diskService.Disks(ctx); err != nil {
			break
		} else if info, exist := diskSet[osDiskName]; exist && info.Properties.ProvisioningState == "Succeeded" {
			break
		}

		select {
		case <-ctx.Done():
			err = ctx.Err()
			break
		case <-time.After(s.SleepTime):
		}
	}

	if err != nil {
		s.Logger.Println("Cannot create virtual machine", name, ":", err.Error())
		s.Logger.Println(toJSON(param))
		networkService.DeleteNetworkInterface(ctx, nif.Name)
	}
	return

}

// DeleteInstance deletes the given named instance.
func (s *ComputeService) DeleteInstance(ctx context.Context, name string) (err error) {

	s.Logger.Println("Deleting virtual machine", name)
	deleted, deleting, nocontent, err := s.client.VirtualMachines.VirtualMachinesDelete(
		virtual_machines.NewVirtualMachinesDeleteParamsWithContext(ctx).
			WithAPIVersion(ComputeAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithVMName(name), httptransport.BearerToken(s.Config.Token.AccessToken))

	switch {
	case err != nil:
		err = NewAPIError(err)
		s.Logger.Println("Cannot delete virtual machine", name)

	case deleted != nil || deleting != nil:
		var instances map[string]struct{}
		for {
			s.Logger.Println("Waiting for deleting virtual machine", name)
			if instances, err = s.Instances(ctx); err != nil {
				s.Logger.Println("Cannot delete virtual machine", name)
				break
			} else if _, ok := instances[name]; !ok {
				s.Logger.Println("Deleted virtual machine", name)
				break
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-time.After(s.SleepTime):
			}
		}

	case nocontent != nil:
		s.Logger.Println("Deleting virtual machine doesn't exist")

	default:
		err = fmt.Errorf("Unexpected case has occurred")
		s.Logger.Println(err.Error())

	}

	networkService, e1 := NewNetworkService(ctx, s.Config, s.Logger)
	if e1 != nil {
		s.Logger.Println("Cannote delete a network interface associated with virtual machine", name, ":", err.Error())
	} else {
		e1 = networkService.DeleteNetworkInterface(ctx, s.networkInterfaceName(name))
	}

	diskService, e2 := NewDiskService(ctx, s.Config, s.Logger)
	if e2 != nil {
		s.Logger.Println("Cannot delete a disk associated with virtual machine", name, ":", err.Error())
	} else {
		e2 = diskService.DeleteDisk(ctx, s.osDiskName(name))
	}

	if err == nil {
		if e1 != nil {
			err = e1
		} else {
			err = e2
		}
	}
	return

}

// Instances returns a list of running instances
func (s *ComputeService) Instances(ctx context.Context) (instances map[string]struct{}, err error) {

	s.Logger.Println("Retrieving instances")
	res, err := s.client.VirtualMachines.VirtualMachinesList(
		virtual_machines.NewVirtualMachinesListParamsWithContext(ctx).
			WithAPIVersion(ComputeAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName), httptransport.BearerToken(s.Config.Token.AccessToken))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve instances")
		return
	}

	instances = make(map[string]struct{})
	for _, v := range res.Payload.Value {
		fmt.Println(v.Properties.ProvisioningState)
		instances[v.Name] = struct{}{}
	}
	s.Logger.Println("Retrieved instances")
	return

}

// GetInstanceInfo retrieves information of a given named instance.
func (s *ComputeService) GetInstanceInfo(ctx context.Context, name string) (info *models.VirtualMachine, err error) {

	s.Logger.Println("Retrieving information of instance", name)
	res, err := s.client.VirtualMachines.VirtualMachinesGet(
		virtual_machines.NewVirtualMachinesGetParamsWithContext(ctx).
			WithAPIVersion(ComputeAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithVMName(name), httptransport.BearerToken(s.Config.Token.AccessToken))
	if err != nil {
		return
	}

	s.Logger.Println("Retrieved the instance information")
	info = res.Payload
	return

}

// AvailableRegions returns a list of available regions.
func (s *ComputeService) AvailableRegions(ctx context.Context) (regions []cloud.Region, err error) {
	s.Logger.Println("Retrieving available regions")
	regions, err = Locations(ctx, &s.Config.Token, s.Config.SubscriptionID)
	if err != nil {
		s.Logger.Println("Cannot retrieve available regions")
	} else {
		s.Logger.Println("Retrieved available regions")
	}
	return
}

// AvailableMachineTypes returns a list of available machine types.
func (s *ComputeService) AvailableMachineTypes(ctx context.Context) (types []cloud.MachineType, err error) {

	s.Logger.Println("Retrieving available machine types")
	res, err := s.client.VirtualMachineSizes.VirtualMachineSizesList(
		virtual_machine_sizes.NewVirtualMachineSizesListParamsWithContext(ctx).
			WithAPIVersion(ComputeAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithLocation(s.Config.Location), httptransport.BearerToken(s.Config.Token.AccessToken))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve available machine types:", err.Error())
		return
	}

	types = make([]cloud.MachineType, len(res.Payload.Value))
	for i, v := range res.Payload.Value {
		types[i] = cloud.MachineType{
			Name:        v.Name,
			Description: fmt.Sprintf("%v Cores, %v MB RAM", v.NumberOfCores, v.MemoryInMB),
		}
	}
	s.Logger.Println("Retrieved available machine types")
	return

}

// ImagePublishers retrieves a set of image publishers.
func (s *ComputeService) ImagePublishers(ctx context.Context) (publishers []Entry, err error) {

	s.Logger.Println("Retrieving image publishers")
	res, err := s.client.VirtualMachineImages.VirtualMachineImagesListPublishers(
		virtual_machine_images.NewVirtualMachineImagesListPublishersParamsWithContext(ctx).
			WithAPIVersion(ComputeAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithLocation(s.Config.Location), httptransport.BearerToken(s.Config.Token.AccessToken))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve image publishers:", err.Error())
		return
	}

	publishers = make([]Entry, len(res.Payload))
	for i, v := range res.Payload {
		publishers[i] = Entry{
			ID:   v.ID,
			Name: *v.Name,
		}
	}
	s.Logger.Println("Retrieved image publishers")
	return

}

// ImageOffers retrieves a set of offers provided by a given publisher.
func (s *ComputeService) ImageOffers(ctx context.Context, publisher string) (offers []Entry, err error) {

	s.Logger.Println("Retrieving image offers of ", publisher)
	res, err := s.client.VirtualMachineImages.VirtualMachineImagesListOffers(
		virtual_machine_images.NewVirtualMachineImagesListOffersParamsWithContext(ctx).
			WithAPIVersion(ComputeAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithLocation(s.Config.Location).
			WithPublisherName(publisher), httptransport.BearerToken(s.Config.Token.AccessToken))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve image offers:", err.Error())
		return
	}

	offers = make([]Entry, len(res.Payload))
	for i, v := range res.Payload {
		offers[i] = Entry{
			ID:   v.ID,
			Name: *v.Name,
		}
	}
	s.Logger.Println("Retrieved image offers")
	return

}

// ImageSkus retrieves a set of skus provded by a given publisher and offer.
func (s *ComputeService) ImageSkus(ctx context.Context, publisherName, offer string) (skus []Entry, err error) {

	s.Logger.Println("Retrieving image skus of ", publisherName, ":", offer)
	res, err := s.client.VirtualMachineImages.VirtualMachineImagesListSkus(
		virtual_machine_images.NewVirtualMachineImagesListSkusParamsWithContext(ctx).
			WithAPIVersion(ComputeAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithLocation(s.Config.Location).
			WithPublisherName(publisherName).
			WithOffer(offer), httptransport.BearerToken(s.Config.Token.AccessToken))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve image skus:", err.Error())
		return
	}

	skus = make([]Entry, len(res.Payload))
	for i, v := range res.Payload {
		skus[i] = Entry{
			ID:   v.ID,
			Name: *v.Name,
		}
	}
	s.Logger.Println("Retrieved image skus")
	return

}

// ImageVersions retrieves a set of versions provided by a given publisher,
// offer, and skus.
func (s *ComputeService) ImageVersions(ctx context.Context, publisherName, offer, skus string) (versions []Entry, err error) {

	s.Logger.Println("Retrieving image versions of ", publisherName, ":", offer, ":", skus)
	res, err := s.client.VirtualMachineImages.VirtualMachineImagesList(
		virtual_machine_images.NewVirtualMachineImagesListParamsWithContext(ctx).
			WithAPIVersion(ComputeAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithLocation(s.Config.Location).
			WithPublisherName(publisherName).
			WithOffer(offer).
			WithSkus(skus), httptransport.BearerToken(s.Config.Token.AccessToken))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve image versions:", err.Error())
		return
	}

	versions = make([]Entry, len(res.Payload))
	for i, v := range res.Payload {
		versions[i] = Entry{
			ID:   v.ID,
			Name: *v.Name,
		}
	}
	s.Logger.Println("REtrieved image versions")
	return

}

// ImageID retrieves an image ID.
func (s *ComputeService) ImageID(ctx context.Context, publisherName, offer, skus, version string) (id string, err error) {

	s.Logger.Println("Retrieving an image ID")
	res, err := s.client.VirtualMachineImages.VirtualMachineImagesGet(
		virtual_machine_images.NewVirtualMachineImagesGetParamsWithContext(ctx).
			WithAPIVersion(ComputeAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithLocation(s.Config.Location).
			WithPublisherName(publisherName).
			WithOffer(offer).
			WithSkus(skus).
			WithVersion(version), httptransport.BearerToken(s.Config.Token.AccessToken))

	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve an image ID:", err.Error())
		return
	}

	s.Logger.Println("Retreived the image ID")
	return res.Payload.ID, nil

}

// networkInterfaceName creates a network interface name associated with a given
// virtual machine name.
func (s *ComputeService) networkInterfaceName(name string) string {
	return fmt.Sprintf("%s-network", name)
}

// osDiskName creates an OS disk name associated with a given virtual machine
// name.
func (s *ComputeService) osDiskName(name string) string {
	return fmt.Sprintf("%s-os", name)
}
