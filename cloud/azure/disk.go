//
// cloud/azure/disk.go
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

// This source file is associated with Azure's Disk API of which Swagger's
// clients are stored in `disk` directory.

package azure

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/jkawamoto/roadie/cloud/azure/disk/client"
	"github.com/jkawamoto/roadie/cloud/azure/disk/client/disks"
	"github.com/jkawamoto/roadie/cloud/azure/disk/models"
)

const (
	// DiskAPIVersion defines API version of disk service.
	DiskAPIVersion = "2016-04-30-preview"
)

// DiskService provides an interface for Azure's disk service.
type DiskService struct {
	client    *client.DiskResourceProviderClient
	Config    *AzureConfig
	Logger    *log.Logger
	SleepTime time.Duration
}

// DiskSet is a map of which key represents a disk name and value is the
// associated disk information.
type DiskSet map[string]*models.Disk

// NewDiskService creates a new disk service interface assosiated with
// a given config; to authorize a authentication token
// is required.
func NewDiskService(ctx context.Context, cfg *AzureConfig, logger *log.Logger) (service *DiskService, err error) {

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
	switch transport := cli.Transport.(type) {
	case *httptransport.Runtime:
		transport.DefaultAuthentication = httptransport.BearerToken(cfg.Token.AccessToken)
		cli.Disks.SetTransport(transport)
	}

	service = &DiskService{
		client:    cli,
		Config:    cfg,
		Logger:    logger,
		SleepTime: DefaultSleepTime,
	}
	return

}

// CreateDiskFromImage creates a disk from an image.
func (s *DiskService) CreateDiskFromImage(ctx context.Context, name, imageID string, diskSize int32) (disk *models.Disk, err error) {

	s.Logger.Println("Creating disk", name, "from", imageID)
	created, creating, err := s.client.Disks.DisksCreateOrUpdate(
		disks.NewDisksCreateOrUpdateParamsWithContext(ctx).
			WithAPIVersion(DiskAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithDiskName(name).WithDisk(&models.Disk{
			Resource: models.Resource{
				Location: &s.Config.Location,
			},
			Properties: &models.DiskProperties{
				AccountType: models.DiskPropertiesAccountTypePremiumLRS,
				CreationData: &models.CreationData{
					CreateOption: toPtr(models.CreationDataCreateOptionFromImage),
					ImageReference: &models.ImageDiskReference{
						ID: &imageID,
					},
				},
				DiskSizeGB: diskSize,
				OsType:     models.DiskPropertiesOsTypeLinux,
			},
		}))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case created != nil:
		s.Logger.Println("Created disk", name)
		return created.Payload, nil

	case creating != nil:
		var info DiskSet
		s.Logger.Println("Waiting for creating disk", name)
		for {
			if info, err = s.Disks(ctx); err != nil {
				break
			} else if d, existing := info[name]; existing {
				s.Logger.Println("Created disk", name)
				return d, nil
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

	s.Logger.Println("Cannot create disk", name, ":", err.Error())
	return

}

// Disks retrieves disks.
func (s *DiskService) Disks(ctx context.Context) (info DiskSet, err error) {

	s.Logger.Println("Retrieving disks")
	res, err := s.client.Disks.DisksListByResourceGroup(
		disks.NewDisksListByResourceGroupParamsWithContext(ctx).
			WithAPIVersion(DiskAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve disks")
		return
	}

	info = make(DiskSet)
	for _, v := range res.Payload.Value {
		info[v.Name] = v
	}
	s.Logger.Println("Retrieved disks")
	return

}

// DeleteDisk deletes a given named disk.
func (s *DiskService) DeleteDisk(ctx context.Context, name string) (err error) {

	s.Logger.Println("Deleting disk", name)
	deleted, deleting, nocontent, err := s.client.Disks.DisksDelete(
		disks.NewDisksDeleteParamsWithContext(ctx).
			WithAPIVersion(DiskAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithDiskName(name))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case deleted != nil:
		s.Logger.Println("Deleted disk", name)
		return

	case deleting != nil:
		var diskset map[string]*models.Disk
		for {
			s.Logger.Println("Waiting for deleting disk", name)
			if diskset, err = s.Disks(ctx); err != nil {
				break
			} else if _, exist := diskset[name]; !exist {
				s.Logger.Panicln("Deleted disk", name)
				return
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-time.After(s.SleepTime):
			}
		}

	case nocontent != nil:
		s.Logger.Println("Deleting disk doesn't exist")
		return

	default:
		err = fmt.Errorf("Unexpected case has occurred")

	}

	s.Logger.Println("Cannot delete disk", name, ":", err.Error())
	return

}

// GetDiskInfo retrieves information of a given named disk.
func (s *DiskService) GetDiskInfo(ctx context.Context, name string) (info *models.Disk, err error) {

	s.Logger.Println("Retrieving information of disk", name)
	res, err := s.client.Disks.DisksGet(
		disks.NewDisksGetParamsWithContext(ctx).
			WithAPIVersion(DiskAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithDiskName(name))
	if err != nil {
		return
	}

	s.Logger.Println("Retrieved the disk information")
	info = res.Payload
	return

}
