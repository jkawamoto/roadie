// +build remote
//
// cloud/azure/network_test.go
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
	"log"
	"os"
	"testing"
	"time"
)

func TestVirtualNetworks(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	ctx := context.Background()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	service, err := NewNetworkService(ctx, cfg.Token, cfg.SubscriptionID, "westus2", logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	name := fmt.Sprintf("test%v-vnet", time.Now().Unix())
	_, err = service.CreateVirtualNetwork(ctx, name)
	if err != nil {
		t.Fatal(err.Error())
	}

	names, err := service.VirtualNetworks(ctx)
	if err != nil {
		t.Error(err.Error())
	}
	if _, existing := names[name]; !existing {
		t.Error("Created virtual network is not found")
	}

	err = service.DeleteVirtualNetwork(ctx, name)
	if err != nil {
		t.Fatal(err.Error())
	}

	names, err = service.VirtualNetworks(ctx)
	if _, existing := names[name]; existing {
		t.Error("Deleted virtual network is found")
	}

}

func TestPublicIPAddress(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	ctx := context.Background()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	service, err := NewNetworkService(ctx, cfg.Token, cfg.SubscriptionID, "westus2", logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	name := fmt.Sprintf("test%v-ip", time.Now().Unix())
	_, err = service.CreatePublicIPAddress(ctx, name)
	if err != nil {
		t.Fatal(err.Error())
	}

	names, err := service.PublicIPAddresses(ctx)
	if err != nil {
		t.Error(err.Error())
	}
	if _, existing := names[name]; !existing {
		t.Error("Created public IP address is not found")
	}

	err = service.DeletePublicIPAddress(ctx, name)
	if err != nil {
		t.Fatal(err.Error())
	}

	names, err = service.PublicIPAddresses(ctx)
	if _, existing := names[name]; existing {
		t.Error("Deleted public IP address is found")
	}

}

func TestNetworkInterface(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	ctx := context.Background()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	service, err := NewNetworkService(ctx, cfg.Token, cfg.SubscriptionID, "westus2", logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	name := fmt.Sprintf("test%v-network", time.Now().Unix())
	nif, err := service.CreateNetworkInterface(ctx, name)
	if err != nil {
		t.Fatal(err.Error())
	}
	if nif.Name != name {
		t.Error("Name of the created network interface doesn't match")
	}

	interfaces, err := service.NetworkInterfaces(ctx)
	if err != nil {
		t.Error(err.Error())
	}
	if _, existing := interfaces[name]; !existing {
		t.Error("Created network interface is not found")
	}

	err = service.DeleteNetworkInterface(ctx, name)
	if err != nil {
		t.Error(err.Error())
	}

	interfaces, err = service.NetworkInterfaces(ctx)
	if err != nil {
		t.Error(err.Error())
	}
	if _, existing := interfaces[name]; existing {
		t.Error("Deleted network interface exists")
	}

}
