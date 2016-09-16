//
// command/util/cloudinstance.go
//
// Copyright (c) 2016 Junpei Kawamoto
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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package util

import (
	"fmt"
	"strings"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/config"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

const gceScope = compute.CloudPlatformScope

// MetadataItem has Key and Value properties.
type MetadataItem struct {
	Key   string
	Value string
}

// MachineType defines a structure of machine type infoemation.
type MachineType struct {
	Name        string
	Description string
}

// Zone defines a structure of zone information.
type Zone struct {
	Name   string
	Status string
}

// newComputeService creates a new service under a given context.
func newComputeService(ctx context.Context) (*compute.Service, error) {

	// Create a client.
	client, err := google.DefaultClient(context.Background(), gceScope)
	if err != nil {
		return nil, err
	}

	// Create a servicer.
	return compute.New(client)

}

// AvailableZones returns a slice of zone names.
func AvailableZones(ctx context.Context) (zones []Zone, err error) {

	cfg, ok := config.FromContext(ctx)
	if !ok {
		err = fmt.Errorf("Given context doesn't have a config: %s", ctx)
		return
	}

	service, err := newComputeService(ctx)
	if err != nil {
		return
	}

	res, err := service.Zones.List(cfg.Gcp.Project).Do()
	if err != nil {
		return nil, err
	}

	zones = make([]Zone, len(res.Items))
	for i, v := range res.Items {
		zones[i] = Zone{
			Name:   v.Name,
			Status: v.Status,
		}
	}
	return

}

// AvailableMachineTypes returns a slice of machie type names.
func AvailableMachineTypes(ctx context.Context) (types []MachineType, err error) {

	cfg, ok := config.FromContext(ctx)
	if !ok {
		err = fmt.Errorf("Given context doesn't have a config: %s", ctx)
		return
	}

	service, err := newComputeService(ctx)
	if err != nil {
		return
	}

	res, err := service.MachineTypes.List(cfg.Gcp.Project, cfg.Gcp.Zone).Do()
	if err != nil {
		return
	}

	types = make([]MachineType, len(res.Items))
	for i, v := range res.Items {
		types[i] = MachineType{Name: v.Name, Description: v.Description}
	}
	return

}

// CreateInstance creates a new instance based on the bilder's configuration.
func CreateInstance(ctx context.Context, name string, metadata []*MetadataItem, disksize int64) (err error) {

	cfg, ok := config.FromContext(ctx)
	if !ok {
		err = fmt.Errorf("Given context doesn't have a config: %s", ctx)
		return
	}

	service, err := newComputeService(ctx)
	if err != nil {
		return
	}

	matadataItems := make([]*compute.MetadataItems, len(metadata))
	for i, v := range metadata {
		matadataItems[i] = &compute.MetadataItems{
			Key:   v.Key,
			Value: &v.Value,
		}
	}

	bluepring := compute.Instance{
		Name:        strings.ToLower(name),
		Zone:        normalizedZone(cfg.Gcp.Project, cfg.Gcp.Zone),
		MachineType: normalizedMachineType(cfg.Gcp.Project, cfg.Gcp.Zone, cfg.Gcp.MachineType),
		Disks: []*compute.AttachedDisk{
			&compute.AttachedDisk{
				Type:       "PERSISTENT",
				Boot:       true,
				Mode:       "READ_WRITE",
				AutoDelete: true,
				InitializeParams: &compute.AttachedDiskInitializeParams{
					SourceImage: "https://www.googleapis.com/compute/v1/projects/coreos-cloud/global/images/coreos-stable-1010-5-0-v20160527",
					DiskType:    normalizedZone(cfg.Gcp.Project, cfg.Gcp.Zone) + "/diskTypes/pd-standard",
					DiskSizeGb:  disksize,
				},
			},
		},
		CanIpForward: false,
		NetworkInterfaces: []*compute.NetworkInterface{
			&compute.NetworkInterface{
				Network: "projects/" + cfg.Gcp.Project + "/global/networks/default",
				AccessConfigs: []*compute.AccessConfig{
					&compute.AccessConfig{
						Name: "External NAT",
						Type: "ONE_TO_ONE_NAT",
					},
				},
			},
		},
		Scheduling: &compute.Scheduling{
			Preemptible:       false,
			OnHostMaintenance: "MIGRATE",
			AutomaticRestart:  true,
		},
		ServiceAccounts: []*compute.ServiceAccount{
			&compute.ServiceAccount{
				Email: "default",
				Scopes: []string{
					"https://www.googleapis.com/auth/cloud-platform",
				},
			},
		},
		Metadata: &compute.Metadata{
			Items: matadataItems,
		},
	}

	res, err := service.Instances.Insert(cfg.Gcp.Project, cfg.Gcp.Zone, &bluepring).Do()
	if err == nil {
		if res.StatusMessage != "" {
			fmt.Println(res.StatusMessage)
		}
		for _, v := range res.Warnings {
			fmt.Println(chalk.Red.Color(v.Message))
		}
	}
	return

}

// DeleteInstance deletes a given named instance.
func DeleteInstance(ctx context.Context, name string) (err error) {

	cfg, ok := config.FromContext(ctx)
	if !ok {
		err = fmt.Errorf("Given context doesn't have a config: %s", ctx)
		return
	}

	service, err := newComputeService(ctx)
	if err != nil {
		return
	}

	res, err := service.Instances.Stop(cfg.Gcp.Project, cfg.Gcp.Zone, name).Do()
	if err == nil {
		if res.StatusMessage != "" {
			fmt.Println(res.StatusMessage)
		}
		for _, v := range res.Warnings {
			fmt.Println(chalk.Red.Color(v.Message))
		}
	}
	return
}

// normalizedZone returns the normalized zone string of Zone property.
func normalizedZone(project, zone string) string {
	return "projects/" + project + "/zones/" + zone
}

// normalizedMachineType returns the normalized instance type of MachineType property.
func normalizedMachineType(project, zone, mtype string) string {
	return normalizedZone(project, zone) + "/machineTypes/" + mtype
}
