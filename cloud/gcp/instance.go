//
// cloud/gcp/instance.go
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

package gcp

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"

	"google.golang.org/api/compute/v1"
)

const (
	// SourceImage defines the ID of the source image to be used for instance.
	SourceImage = "projects/coreos-cloud/global/images/coreos-stable-1298-7-0-v20170401"
)

var (
	// RoadieSchemeURLOffset defines an offset value to remove scheme name from
	// URLs.
	RoadieSchemeURLOffset = len(script.RoadieSchemePrefix)
)

// ComputeService implements cloud.InstanceManager based on Google Cloud
// Platform.
type ComputeService struct {
	Config    *Config
	Logger    *log.Logger
	SleepTime time.Duration
}

// NewComputeService creates a new compute service client.
func NewComputeService(cfg *Config, logger *log.Logger) *ComputeService {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}
	return &ComputeService{
		Config:    cfg,
		Logger:    logger,
		SleepTime: 10 * time.Second,
	}
}

// newService creates a new service under a given context.
func (s *ComputeService) newService(ctx context.Context) (*compute.Service, error) {

	// Create a client.
	cfg := NewAuthorizationConfig(0)
	client := cfg.Client(ctx, s.Config.Token)

	// Create a servicer.
	return compute.New(client)

}

// AvailableRegions returns a slice of region information.
func (s *ComputeService) AvailableRegions(ctx context.Context) (regions []cloud.Region, err error) {

	s.Logger.Println("Retrieving available regions")
	service, err := s.newService(ctx)
	if err != nil {
		return
	}

	res, err := service.Zones.List(s.Config.Project).Do()
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

	s.Logger.Println("Finished retrieving available regions")
	return

}

// AvailableMachineTypes returns a slice of machie type names.
func (s *ComputeService) AvailableMachineTypes(ctx context.Context) (types []cloud.MachineType, err error) {

	s.Logger.Println("Retrieving available machine types")
	service, err := s.newService(ctx)
	if err != nil {
		return
	}

	res, err := service.MachineTypes.List(s.Config.Project, s.Config.Zone).Do()
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

	s.Logger.Println("Finished retrieving available machine types")
	return

}

// CreateInstance creates a new instance based on the builder's configuration.
func (s *ComputeService) CreateInstance(ctx context.Context, task *script.Script) (err error) {

	if task.Image == "" {
		task.Image = DefaultBaseImage
	}

	// Create an ignition config.
	fluentd, err := FluentdUnit(task.Name)
	if err != nil {
		return
	}
	options := ""
	if s.Config.NoShutdown {
		options += "--no-shutdown"
	}
	roadie, err := RoadieUnit(task.Name, task.Image, options)
	if err != nil {
		return
	}
	logcast, err := LogcastUnit("roadie.service")
	if err != nil {
		return
	}
	ignition := NewIgnitionConfig().Append(fluentd).Append(roadie).Append(logcast).String()
	s.Logger.Println("Ignition configuration is", ignition)

	// Update URLs of which scheme is `roadie://` to `gs://`.
	ReplaceURLScheme(s.Config, task)
	s.Logger.Printf("Updated script file is \n%v\n", task.String())

	scriptStr := task.String()
	err = s.createInstance(ctx, task.Name, []*compute.MetadataItems{
		&compute.MetadataItems{
			Key:   "script",
			Value: &scriptStr,
		},
		&compute.MetadataItems{
			Key:   "user-data",
			Value: &ignition,
		},
	})
	if err != nil {
		return
	}

	s.Logger.Println("Finished creating instance", task.Name)
	return

}

// DeleteInstance deletes a given named instance.
func (s *ComputeService) DeleteInstance(ctx context.Context, name string) (err error) {

	s.Logger.Println("Deleting instance", name)
	service, err := s.newService(ctx)
	if err != nil {
		return
	}

	res, err := service.Instances.Delete(s.Config.Project, s.Config.Zone, name).Do()
	if err == nil {
		s.Logger.Println("Finished deleting instance")
		if res.StatusMessage != "" {
			s.Logger.Println("*", res.StatusMessage)
		}
		for _, v := range res.Warnings {
			s.Logger.Println("*", v.Message)
		}
	}
	return
}

// Instances returns a list of running instances
func (s *ComputeService) Instances(ctx context.Context, handler cloud.InstanceHandler) (err error) {

	s.Logger.Println("Retrieving running instances")
	instances := make(map[string]struct{})
	log := NewLogManager(s.Config, s.Logger)
	err = log.OperationLogEntries(ctx, time.Time{}, func(_ time.Time, payload *ActivityPayload) error {

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		switch payload.EventSubtype {
		case LogEventSubtypeInsert:
			instances[payload.Resource.Name] = struct{}{}

		case LogEventSubtypeDelete:
			delete(instances, payload.Resource.Name)
		}
		return nil

	})
	if err != nil {
		return
	}

	var instanceNames []string
	for name := range instances {
		instanceNames = append(instanceNames, name)
	}
	sort.Strings(instanceNames)
	for _, name := range instanceNames {
		err = handler(name, StatusRunning)
		if err != nil {
			return
		}
	}

	s.Logger.Println("Retrieving terminated instances")
	storage, err := NewStorageService(ctx, s.Config, s.Logger)
	if err != nil {
		return err
	}

	var prev string
	err = storage.List(ctx, script.ResultPrefix, "", func(info *cloud.FileInfo) (err error) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if info.Name == "" {
			return
		}

		dir := path.Base(path.Dir(info.Path))
		if prev != dir {
			if _, exist := instances[dir]; !exist {
				handler(dir, StatusTerminated)
			}
			prev = dir
		}
		return
	})
	if err != nil {
		return
	}

	s.Logger.Println("Finished retrieving instances")
	return
}

// CreateInstance creates a new instance based on the builder's configuration.
func (s *ComputeService) createInstance(ctx context.Context, instanceName string, matadataItems []*compute.MetadataItems) (err error) {

	s.Logger.Println("Creating instance", instanceName)
	service, err := s.newService(ctx)
	if err != nil {
		return
	}

	blueprint := compute.Instance{
		Name:        strings.ToLower(instanceName),
		Zone:        s.Config.normalizedZone(),
		MachineType: s.Config.normalizedMachineType(),
		Disks: []*compute.AttachedDisk{
			&compute.AttachedDisk{
				Type:       "PERSISTENT",
				Boot:       true,
				Mode:       "READ_WRITE",
				AutoDelete: true,
				InitializeParams: &compute.AttachedDiskInitializeParams{
					SourceImage: SourceImage,
					DiskType:    s.Config.diskType(),
					DiskSizeGb:  s.Config.DiskSize,
				},
			},
		},
		CanIpForward: false,
		NetworkInterfaces: []*compute.NetworkInterface{
			&compute.NetworkInterface{
				Network: s.Config.network(),
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

	res, err := service.Instances.Insert(s.Config.Project, s.Config.Zone, &blueprint).Do()
	if err != nil {
		return
	}
	if res.StatusMessage != "" {
		s.Logger.Println(res.StatusMessage)
	}
	for _, v := range res.Warnings {
		s.Logger.Println("*", v.Message)
	}

	log := NewLogManager(s.Config, s.Logger)
	from := time.Now().In(time.UTC)
	// Token for specifying a target instance has been started.
	instanceStarted := fmt.Errorf("Target instance has been started.")
	for {

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(s.SleepTime):
		}

		err := log.OperationLogEntries(ctx, from, func(t time.Time, payload *ActivityPayload) (err error) {
			if payload.EventSubtype == LogEventSubtypeInsert && payload.Resource.Name == blueprint.Name {
				return instanceStarted
			}
			from = t
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
