// +build remote
//
// cloud/azure/disk_test.go
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

func TestDiskService(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	ctx := context.Background()
	location := "westus2"
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	service, err := NewDiskService(ctx, cfg.Token, cfg.SubscriptionID, location, logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	publisherName := "Canonical"
	offer := "UbuntuServer"
	skus := "16.10"
	version := "16.10.201703070"
	compute, err := NewComputeService(ctx, cfg.Token, cfg.SubscriptionID, location, os.Stdout)
	if err != nil {
		t.Fatal(err.Error())
	}

	imageID, err := compute.ImageID(ctx, publisherName, offer, skus, version)
	if err != nil {
		t.Fatal(err.Error())
	}

	name := fmt.Sprintf("disk%v", time.Now().Unix())
	_, err = service.CreateDiskFromImage(ctx, name, imageID, 10)
	if err != nil {
		t.Error(err.Error())
	}

	disks, err := service.Disks(ctx)
	if err != nil {
		t.Error(err.Error())
	}
	if _, exist := disks[name]; !exist {
		t.Error("Cannot find created disk")
	}

	err = service.DeleteDisk(ctx, name)
	if err != nil {
		t.Error(err.Error())
	}

	disks, err = service.Disks(ctx)
	if err != nil {
		t.Error(err.Error())
	}
	if _, exist := disks[name]; exist {
		t.Error("Deleted disk is found")
	}

}
