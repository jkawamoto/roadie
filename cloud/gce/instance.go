//
// cloud/gce/instance.go
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

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/logging"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/command/log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

const (
	// GCP's scope.
	gceScope = compute.CloudPlatformScope
)

var (
	// Token for specifying a target instance has been started.
	instanceStarted = fmt.Errorf("Target instance has been started.")
)

type ComputeService struct {
	Project     string
	Region      string
	MachineType string
	Log         io.Writer
}

func NewComputeService(project, region, machine string, log io.Writer) *ComputeService {

	if log == nil {
		log = os.Stderr
	}

	return &ComputeService{
		Project:     project,
		Region:      region,
		MachineType: machine,
		Log:         log,
	}
}

// newService creates a new service under a given context.
func (s *ComputeService) newService(ctx context.Context) (*compute.Service, error) {

	// Create a client.
	client, err := google.DefaultClient(ctx, gceScope)
	if err != nil {
		return nil, err
	}

	// Create a servicer.
	return compute.New(client)

}

// AvailableRegions returns a slice of region information.
func (s *ComputeService) AvailableRegions(ctx context.Context) (regions []cloud.Region, err error) {

	service, err := s.newService(ctx)
	if err != nil {
		return
	}

	res, err := service.Zones.List(s.Project).Do()
	if err != nil {
		return
	}

	regions = make([]cloud.Region, len(res.Items))
	for i, v := range res.Items {
		regions[i] = cloud.Region{
			Name:   v.Name,
			Status: v.Status,
		}
	}
	return

}

// AvailableMachineTypes returns a slice of machie type names.
func (s *ComputeService) AvailableMachineTypes(ctx context.Context) (types []cloud.MachineType, err error) {

	service, err := s.newService(ctx)
	if err != nil {
		return
	}

	res, err := service.MachineTypes.List(s.Project, s.Region).Do()
	if err != nil {
		return
	}

	types = make([]cloud.MachineType, len(res.Items))
	for i, v := range res.Items {
		types[i] = cloud.MachineType{
			Name:        v.Name,
			Description: v.Description,
		}
	}
	return

}

// CreateInstance creates a new instance based on the bilder's configuration.
func (s *ComputeService) CreateInstance(ctx context.Context, name string, metadata []*cloud.MetadataItem, disksize int64) (err error) {

	service, err := s.newService(ctx)
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

	blueprint := compute.Instance{
		Name:        strings.ToLower(name),
		Zone:        s.normalizedZone(),
		MachineType: s.normalizedMachineType(),
		Disks: []*compute.AttachedDisk{
			&compute.AttachedDisk{
				Type:       "PERSISTENT",
				Boot:       true,
				Mode:       "READ_WRITE",
				AutoDelete: true,
				InitializeParams: &compute.AttachedDiskInitializeParams{
					SourceImage: "https://www.googleapis.com/compute/v1/projects/coreos-cloud/global/images/coreos-stable-1010-5-0-v20160527",
					DiskType:    s.normalizedZone() + "/diskTypes/pd-standard",
					DiskSizeGb:  disksize,
				},
			},
		},
		CanIpForward: false,
		NetworkInterfaces: []*compute.NetworkInterface{
			&compute.NetworkInterface{
				Network: "projects/" + s.Project + "/global/networks/default",
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

	res, err := service.Instances.Insert(s.Project, s.Region, &blueprint).Do()
	if err != nil {
		return
	}
	if res.StatusMessage != "" {
		fmt.Fprintln(s.Log, res.StatusMessage)
	}
	for _, v := range res.Warnings {
		fmt.Fprintln(s.Log, chalk.Red.Color(v.Message))
	}

	logService := log.NewCloudLoggingService(ctx)
	filter := fmt.Sprintf(
		"jsonPayload.event_type = \"GCE_OPERATION_DONE\" AND timestamp > \"%s\"",
		time.Now().In(time.UTC).Format(log.TimeFormat))

	for {

		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-wait(10 * time.Second):
			err := log.GetEntries(ctx, filter, logService, func(entry *logging.Entry) (err error) {
				payload, err := log.NewActivityPayload(entry.Payload)
				if err != nil {
					return
				}
				if payload.EventSubtype == log.EventSubtypeInsert && payload.Resource.Name == blueprint.Name {
					return instanceStarted
				}
				return
			})

			switch err {
			case instanceStarted:
				return nil
			case nil:
				continue
			default:
				return err
			}
		}

	}

}

// DeleteInstance deletes a given named instance.
func (s *ComputeService) DeleteInstance(ctx context.Context, name string) (err error) {

	service, err := s.newService(ctx)
	if err != nil {
		return
	}

	res, err := service.Instances.Delete(s.Project, s.Region, name).Do()
	if err == nil {
		if res.StatusMessage != "" {
			fmt.Fprintln(s.Log, res.StatusMessage)
		}
		for _, v := range res.Warnings {
			fmt.Fprintln(s.Log, chalk.Red.Color(v.Message))
		}
	}
	return
}

// Instances returns a list of running instances
func (s *ComputeService) Instances(ctx context.Context) (instances map[string]struct{}, err error) {

	instances = make(map[string]struct{})
	requester := log.NewCloudLoggingService(ctx)

	err = log.GetOperationLogEntries(ctx, requester, func(_ time.Time, payload *log.ActivityPayload) error {

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		switch payload.EventSubtype {
		case log.EventSubtypeInsert:
			instances[payload.Resource.Name] = struct{}{}

		case log.EventSubtypeDelete:
			delete(instances, payload.Resource.Name)
		}
		return nil

	})

	return
}

// normalizedZone returns the normalized zone string of Zone property.
func (s *ComputeService) normalizedZone() string {
	return "projects/" + s.Project + "/zones/" + s.Region
}

// normalizedMachineType returns the normalized instance type of MachineType property.
func (s *ComputeService) normalizedMachineType() string {
	return s.normalizedZone() + "/machineTypes/" + s.MachineType
}

// Wait a given duration.
func wait(d time.Duration) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.Sleep(d)
		ch <- struct{}{}
	}()
	return ch
}
