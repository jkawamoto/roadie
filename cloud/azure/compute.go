//
// cloud/azure/compute.go
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

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/jkawamoto/roadie/cloud"
	client "github.com/jkawamoto/roadie/cloud/azure/compute/client"
	"github.com/jkawamoto/roadie/cloud/azure/compute/client/virtual_machine_images"
	"github.com/jkawamoto/roadie/cloud/azure/compute/client/virtual_machine_sizes"
)

const (
	// ComputeAPIVersion defines API version of compute service.
	ComputeAPIVersion = "2016-04-30-preview"
)

// ComputeService provides an interface for Azure's compute service.
type ComputeService struct {
	client *client.ComputeManagementClient
	Config *Config
	Logger *log.Logger
}

// Entry is an entry of lists.
type Entry struct {
	ID   string
	Name string
}

// NewComputeService creates a new compute service interface assosiated with
// a given configuration.
func NewComputeService(ctx context.Context, cfg *Config, logger *log.Logger) (s *ComputeService, err error) {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	// Create a management client.
	cli := client.NewHTTPClient(strfmt.NewFormats())
	s = &ComputeService{
		client: cli,
		Config: cfg,
		Logger: logger,
	}
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
