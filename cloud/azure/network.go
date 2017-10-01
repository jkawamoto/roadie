//
// cloud/azure/network.go
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

// This source file is associated with Azure's Network API of which Swagger's
// clients are stored in `network` directory.

package azure

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	client "github.com/jkawamoto/roadie/cloud/azure/network/client"
	"github.com/jkawamoto/roadie/cloud/azure/network/client/network_interfaces"
	"github.com/jkawamoto/roadie/cloud/azure/network/client/public_ip_addresses"
	"github.com/jkawamoto/roadie/cloud/azure/network/client/virtual_networks"
	"github.com/jkawamoto/roadie/cloud/azure/network/models"
)

const (
	// NetworkAPIVersion defines API version of network service.
	NetworkAPIVersion = "2016-03-30"

	// DefaultAddressPrefix defines a default address prefix.
	DefaultAddressPrefix = "10.0.0.0/16"
)

// NetworkService provides an interface for Azure's network service.
type NetworkService struct {
	client        *client.NetworkManagementClient
	Config        *AzureConfig
	AddressPrefix string
	Logger        *log.Logger
	SleepTime     time.Duration
}

// NewNetworkService creates a new network service interface assosiated with
// a given subscription id and location; to authorize a authentication token
// is required.
func NewNetworkService(ctx context.Context, cfg *AzureConfig, logger *log.Logger) (service *NetworkService, err error) {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags|log.Lshortfile)
	}

	// Create a resource group if not exist.
	err = CreateResourceGroupIfNotExist(ctx, cfg, logger)
	if err != nil {
		return
	}

	// Create a management client.
	cli := client.NewHTTPClient(strfmt.NewFormats())
	return &NetworkService{
		client:        cli,
		Config:        cfg,
		AddressPrefix: DefaultAddressPrefix,
		Logger:        logger,
		SleepTime:     DefaultSleepTime,
	}, nil

}

// CreatePublicIPAddress created a new public IP address which has a given name.
// The IP address will be added to the resource group given when this network
// service was created.
func (s *NetworkService) CreatePublicIPAddress(ctx context.Context, name string) (*models.PublicIPAddress, error) {

	s.Logger.Println("Creating a public IP address", name)
	created, creating, err := s.client.PublicIPAddresses.PublicIPAddressesCreateOrUpdate(
		public_ip_addresses.NewPublicIPAddressesCreateOrUpdateParamsWithContext(ctx).
			WithAPIVersion(NetworkAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithPublicIPAddressName(name).
			WithParameters(&models.PublicIPAddress{
				Resource: models.Resource{
					Location: s.Config.Location,
				},
				Properties: &models.PublicIPAddressPropertiesFormat{
					PublicIPAddressVersion:   models.PublicIPAddressPropertiesFormatPublicIPAddressVersionIPV4,
					PublicIPAllocationMethod: models.PublicIPAddressPropertiesFormatPublicIPAllocationMethodDynamic,
					IDLETimeoutInMinutes:     4,
				},
			}), httptransport.BearerToken(s.Config.Token.AccessToken))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case created != nil || creating != nil:
		var addresses map[string]*models.PublicIPAddress
		for {
			s.Logger.Println("Waiting for creating new public IP address", name)
			if addresses, err = s.PublicIPAddresses(ctx); err != nil {
				break
			} else if address, existing := addresses[name]; existing {
				s.Logger.Println("Created new public IP address", name)
				return address, nil
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-wait(s.SleepTime):
			}
		}

	default:
		err = fmt.Errorf("Unexpected case has occurred")

	}

	s.Logger.Println("Cannot create new public IP address", name, ":", err.Error())
	return nil, err

}

// PublicIPAddresses retrieves a set of public IP addresses registered to the
// resource group given when this network service created.
func (s *NetworkService) PublicIPAddresses(ctx context.Context) (addresses map[string]*models.PublicIPAddress, err error) {

	s.Logger.Println("Retrieving existing public IP addresses")
	res, err := s.client.PublicIPAddresses.PublicIPAddressesList(
		public_ip_addresses.NewPublicIPAddressesListParamsWithContext(ctx).
			WithAPIVersion(NetworkAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName), httptransport.BearerToken(s.Config.Token.AccessToken))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve existing public IP addresses:", err.Error())
		return
	}

	addresses = make(map[string]*models.PublicIPAddress)
	for _, v := range res.Payload.Value {
		addresses[v.Name] = v
	}

	s.Logger.Println("Retrieved existing public IP addresses")
	return

}

// DeletePublicIPAddress deletes an IP address which has a given name.
func (s *NetworkService) DeletePublicIPAddress(ctx context.Context, name string) (err error) {

	s.Logger.Println("Deleting public IP address", name)
	deleted, deleting, nocontent, err := s.client.PublicIPAddresses.PublicIPAddressesDelete(
		public_ip_addresses.NewPublicIPAddressesDeleteParamsWithContext(ctx).
			WithAPIVersion(NetworkAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithPublicIPAddressName(name), httptransport.BearerToken(s.Config.Token.AccessToken))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case deleted != nil || deleting != nil:
		var addresses map[string]*models.PublicIPAddress
		for {
			s.Logger.Println("Waiting for deleting public IP address", name)
			if addresses, err = s.PublicIPAddresses(ctx); err != nil {
				break
			} else if _, existing := addresses[name]; !existing {
				s.Logger.Println("Deleted public IP address", name)
				return
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-wait(s.SleepTime):
			}
		}

	case nocontent != nil:
		s.Logger.Println("Deleting public IP address doesn't exist")
		return

	default:
		err = fmt.Errorf("Unexpected case has occurred")

	}

	s.Logger.Println("Cannot delete public IP address", name, ":", err.Error())
	return

}

// CreateVirtualNetwork creates a new virtual network which has a given name.
// Created virtual network will belong to the subscription and resource group
// specified whtn the network service was created.
func (s *NetworkService) CreateVirtualNetwork(ctx context.Context, name string) (vnet *models.VirtualNetwork, err error) {

	s.Logger.Println("Creating virtual network", name)
	created, creating, err := s.client.VirtualNetworks.VirtualNetworksCreateOrUpdate(
		virtual_networks.NewVirtualNetworksCreateOrUpdateParamsWithContext(ctx).
			WithAPIVersion(NetworkAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithVirtualNetworkName(name).
			WithParameters(&models.VirtualNetwork{
				Resource: models.Resource{
					Location: s.Config.Location,
				},
				Properties: &models.VirtualNetworkPropertiesFormat{
					AddressSpace: &models.AddressSpace{
						AddressPrefixes: []string{
							s.AddressPrefix,
						},
					},
					Subnets: []*models.Subnet{
						&models.Subnet{
							Name: s.subnetName(name),
							Properties: &models.SubnetPropertiesFormat{
								AddressPrefix: s.AddressPrefix,
							},
						},
					},
				},
			}), httptransport.BearerToken(s.Config.Token.AccessToken))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case created != nil || creating != nil:
		var networks map[string]*models.VirtualNetwork
		for {
			s.Logger.Println("Waiting for creating virtual network", name)
			if networks, err = s.VirtualNetworks(ctx); err != nil {
				break
			} else if vnet, existing := networks[name]; existing {
				s.Logger.Println("Created virtual network", name)
				return vnet, nil
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-wait(s.SleepTime):
			}
		}

	default:
		err = fmt.Errorf("Unexpected case has occurred")

	}

	s.Logger.Println("Cannot create virtual network", name, ":", err.Error())
	return

}

// VirtualNetworks retrieve a set of available virtual networks belonging to
// the subscription and resource groupt specified when this network service was
// created.
func (s *NetworkService) VirtualNetworks(ctx context.Context) (networks map[string]*models.VirtualNetwork, err error) {

	s.Logger.Println("Retrieving virtual networks")
	res, err := s.client.VirtualNetworks.VirtualNetworksList(
		virtual_networks.NewVirtualNetworksListParamsWithContext(ctx).
			WithAPIVersion(NetworkAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName), httptransport.BearerToken(s.Config.Token.AccessToken))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrive virtual networks:", err.Error())
		return
	}

	networks = make(map[string]*models.VirtualNetwork)
	for _, v := range res.Payload.Value {
		networks[v.Name] = v
	}
	s.Logger.Println("Retrieved virtual networks")
	return

}

// DeleteVirtualNetwork deletes a given named virtual network.
func (s *NetworkService) DeleteVirtualNetwork(ctx context.Context, name string) (err error) {

	s.Logger.Println("Deleting virtual network", name)
	deleted, deleting, nocontent, err := s.client.VirtualNetworks.VirtualNetworksDelete(
		virtual_networks.NewVirtualNetworksDeleteParamsWithContext(ctx).
			WithAPIVersion(NetworkAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithVirtualNetworkName(name), httptransport.BearerToken(s.Config.Token.AccessToken))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case deleted != nil || deleting != nil:
		var names map[string]*models.VirtualNetwork
		for {
			s.Logger.Println("Waiting for deleting virtual network", name)
			if names, err = s.VirtualNetworks(ctx); err != nil {
				break
			} else if _, existing := names[name]; !existing {
				s.Logger.Println("Deleted virtual network", name)
				return
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-wait(s.SleepTime):
			}
		}

	case nocontent != nil:
		s.Logger.Println("Deleting virtual network doesn't exist")
		return

	default:
		err = fmt.Errorf("Unexpected case has occurred")

	}

	s.Logger.Println("Cannot delete virtual network", name, ":", err.Error())
	return

}

// CreateNetworkInterface creates a new network interface which has a given
// name; the created network interface will belong to the subscription and
// resource groupt specified when this network service was creates.
// This function also creates a virtual network and a public IP address; and
// attaches it to the creating network interface. When you delete the created
// network interface, the both virtuanl network and public IP address will
// be deleted.
func (s *NetworkService) CreateNetworkInterface(ctx context.Context, name string) (nif *models.NetworkInterface, err error) {

	vnet, err := s.CreateVirtualNetwork(ctx, s.vnetName(name))
	if err != nil {
		return
	} else if len(vnet.Properties.Subnets) == 0 {
		s.DeleteVirtualNetwork(ctx, vnet.Name)
		return nil, fmt.Errorf("Subnet is not found")
	}
	subnet := vnet.Properties.Subnets[0]

	ipAddress, err := s.CreatePublicIPAddress(ctx, s.ipName(name))
	if err != nil {
		s.DeleteVirtualNetwork(ctx, vnet.Name)
	}

	s.Logger.Println("Creating network interface", name)
	created, creating, err := s.client.NetworkInterfaces.NetworkInterfacesCreateOrUpdate(
		network_interfaces.NewNetworkInterfacesCreateOrUpdateParamsWithContext(ctx).
			WithAPIVersion(NetworkAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithNetworkInterfaceName(name).
			WithParameters(&models.NetworkInterface{
				Resource: models.Resource{
					Location: s.Config.Location,
				},
				Properties: &models.NetworkInterfacePropertiesFormat{
					IPConfigurations: []*models.NetworkInterfaceIPConfiguration{
						&models.NetworkInterfaceIPConfiguration{
							Name: s.ipConfigName(name),
							Properties: &models.NetworkInterfaceIPConfigurationPropertiesFormat{
								Subnet:                    subnet,
								PublicIPAddress:           ipAddress,
								PrivateIPAllocationMethod: models.IPConfigurationPropertiesFormatPrivateIPAllocationMethodDynamic,
							},
						},
					},
				},
			}), httptransport.BearerToken(s.Config.Token.AccessToken))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case created != nil || creating != nil:
		var interfaces map[string]*models.NetworkInterface
		for {
			s.Logger.Println("Waiting for creating network interface", name)
			if interfaces, err = s.NetworkInterfaces(ctx); err != nil {
				break
			} else if nif, existing := interfaces[name]; existing {
				s.Logger.Println("Created network interface", name)
				return nif, nil
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-wait(s.SleepTime):
			}
		}

	default:
		err = fmt.Errorf("Unexpected case has occurred")
	}

	s.Logger.Println("Cannot create network interface", name, ":", err.Error())
	s.DeletePublicIPAddress(ctx, ipAddress.Name)
	s.DeleteVirtualNetwork(ctx, vnet.Name)
	return nil, err

}

// NetworkInterfaces retrieves network interfaces belonging to the subscription
// and resource group specified when this network service was created.
func (s *NetworkService) NetworkInterfaces(ctx context.Context) (interfaces map[string]*models.NetworkInterface, err error) {

	s.Logger.Println("Retriving network interfaces")
	res, err := s.client.NetworkInterfaces.NetworkInterfacesList(
		network_interfaces.NewNetworkInterfacesListParamsWithContext(ctx).
			WithAPIVersion(NetworkAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName), httptransport.BearerToken(s.Config.Token.AccessToken))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve network interfaces:", err.Error())
		return
	}

	interfaces = make(map[string]*models.NetworkInterface)
	for _, v := range res.Payload.Value {
		interfaces[v.Name] = v
	}
	s.Logger.Println("Retrieved network interfaces")
	return

}

// DeleteNetworkInterface deletes a given named network interface; it also
// deletes associated virtual network and public IP address.
func (s *NetworkService) DeleteNetworkInterface(ctx context.Context, name string) (err error) {

	s.Logger.Println("Deleting network interface", name)
	deleted, deleting, nocontent, err := s.client.NetworkInterfaces.NetworkInterfacesDelete(
		network_interfaces.NewNetworkInterfacesDeleteParamsWithContext(ctx).
			WithAPIVersion(NetworkAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithNetworkInterfaceName(name), httptransport.BearerToken(s.Config.Token.AccessToken))

	switch {
	case err != nil:
		err = NewAPIError(err)
		s.Logger.Println("Cannot delete network interface", name, ":", err.Error())

	case deleted != nil || deleting != nil:
		var interfaces map[string]*models.NetworkInterface
		for {
			s.Logger.Println("Waiting for deleting network interface", name)
			if interfaces, err = s.NetworkInterfaces(ctx); err != nil {
				s.Logger.Println("Cannot delete network interface", name, ":", err.Error())
				break
			} else if _, existing := interfaces[name]; !existing {
				s.Logger.Println("Deleted network interface", name)
				break
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-wait(s.SleepTime):
			}
		}

	case nocontent != nil:
		s.Logger.Println("Deleting network interface doesn't exist")

	default:
		err = fmt.Errorf("Unexpected case has occurred")
		s.Logger.Println("Cannot delete network interface", name, ":", err.Error())

	}

	e1 := s.DeletePublicIPAddress(ctx, s.ipName(name))
	e2 := s.DeleteVirtualNetwork(ctx, s.vnetName(name))
	if e1 != nil {
		err = e1
	} else if e2 != nil {
		err = e2
	}
	return

}

// vnetName creates a virtual network name associated with a given network
// interface name.
func (s *NetworkService) vnetName(name string) string {
	return fmt.Sprintf("%s-vnet", name)
}

// subnetName creates a subnet name associated with a given network interface
// name.
func (s *NetworkService) subnetName(name string) string {
	return fmt.Sprintf("%s-subnet", name)
}

// ipName creates an IP address name associated with a given network interface
// name.
func (s *NetworkService) ipName(name string) string {
	return fmt.Sprintf("%s-ip", name)
}

// ipConfigName create a configuration name of an IP address associated with a
// given network interface name.
func (s *NetworkService) ipConfigName(name string) string {
	return fmt.Sprintf("%s-ip", name)
}
