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
	"io/ioutil"
	"log"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/logging"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

var (
	// Token for specifying a target instance has been started.
	instanceStarted = fmt.Errorf("Target instance has been started.")
)

// ComputeService implements cloud.InstanceManager based on Google Cloud
// Platform.
type ComputeService struct {
	Config    *GcpConfig
	Logger    *log.Logger
	SleepTime time.Duration
}

// NewComputeService creates a new compute service client.
func NewComputeService(cfg *GcpConfig, logger *log.Logger) *ComputeService {

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
	client, err := google.DefaultClient(ctx, gceScope)
	if err != nil {
		return nil, err
	}

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
func (s *ComputeService) CreateInstance(ctx context.Context, name string, task *script.Script, disksize int64) (err error) {

	s.Logger.Println("Creating instance", name)

	// Update URLs of which scheme is `roadie://` to `gs://`.
	s.replaceURLScheme(task)

	// Create a startup script.
	startup, err := s.createStartupScript(name, task)
	if err != nil {
		return
	}
	err = s.createInstance(ctx, name, startup, disksize)
	if err != nil {
		return
	}

	s.Logger.Println("Finished creating instance", name)
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
func (s *ComputeService) Instances(ctx context.Context) (instances map[string]struct{}, err error) {

	s.Logger.Println("Retrieving running instances")
	instances = make(map[string]struct{})
	log := NewLogManager(s.Config)
	err = log.OperationLogEntries(ctx, func(_ time.Time, payload *ActivityPayload) error {

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

	s.Logger.Println("Finished retrieving running instances")
	return
}

// CreateInstance creates a new instance based on the builder's configuration.
func (s *ComputeService) createInstance(ctx context.Context, name string, startup string, disksize int64) (err error) {

	service, err := s.newService(ctx)
	if err != nil {
		return
	}

	matadataItems := []*compute.MetadataItems{
		&compute.MetadataItems{
			Key:   "startup-script",
			Value: &startup,
		},
	}

	blueprint := compute.Instance{
		Name:        strings.ToLower(name),
		Zone:        s.Config.normalizedZone(),
		MachineType: s.Config.normalizedMachineType(),
		Disks: []*compute.AttachedDisk{
			&compute.AttachedDisk{
				Type:       "PERSISTENT",
				Boot:       true,
				Mode:       "READ_WRITE",
				AutoDelete: true,
				InitializeParams: &compute.AttachedDiskInitializeParams{
					SourceImage: "https://www.googleapis.com/compute/v1/projects/coreos-cloud/global/images/coreos-stable-1010-5-0-v20160527",
					DiskType:    s.Config.diskType(),
					DiskSizeGb:  disksize,
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

	log := NewLogManager(s.Config)
	filter := fmt.Sprintf(
		`jsonPayload.event_type = "GCE_OPERATION_DONE" AND timestamp > "%s"`,
		time.Now().In(time.UTC).Format(LogTimeFormat))

	for {

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-wait(10 * time.Second):
		}

		err := log.Entries(ctx, filter, func(entry *logging.Entry) (err error) {
			payload, err := NewActivityPayload(entry.Payload)
			if err != nil {
				return
			}
			if payload.EventSubtype == LogEventSubtypeInsert && payload.Resource.Name == blueprint.Name {
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

// replaceURLScheme replaced URLs which start with "roadie://".
// Those URLs are modified to "gs://<bucketname>/.roadie/".
func (s *ComputeService) replaceURLScheme(task *script.Script) {

	offset := len(script.RoadieSchemePrefix)

	// Replace source section.
	if strings.HasPrefix(task.Source, script.RoadieSchemePrefix) {
		task.Source = s.createURL(script.SourcePrefix, task.Source[offset:])
	}

	// Replace data section.
	for i, url := range task.Data {
		if strings.HasPrefix(url, script.RoadieSchemePrefix) {
			task.Data[i] = s.createURL(script.DataPrefix, url[offset:])
		}
	}

	// Replace result section.
	if strings.HasPrefix(task.Result, script.RoadieSchemePrefix) {
		task.Result = s.createURL(script.ResultPrefix, task.Result[offset:])
	}

}

// createURL creates a valid URL for uploaing object.
func (s *ComputeService) createURL(group, name string) string {

	u := url.URL{
		Scheme: "gs",
		Host:   s.Config.Bucket,
		Path:   filepath.Join("/", StoragePrefix, group, name),
	}
	return u.String()

}

// createStartupScript creates a start up script with a given name and task.
func (s *ComputeService) createStartupScript(name string, task *script.Script) (startup string, err error) {

	retry := 10
	options := ""
	for _, v := range task.Options {
		switch {
		case strings.HasPrefix(v, "retry:"):
			retry, err = strconv.Atoi(v[len("retry:"):])
			if err != nil {
				retry = 10
			}

		default:
			options += "--" + v

		}
	}

	startup, err = Startup(&StartupOpt{
		Name:    name,
		Script:  task.String(),
		Options: options,
		Image:   task.Image,
		Retry:   retry,
	})
	return

}

// Wait a given duration.
func wait(d time.Duration) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.Sleep(d)
		close(ch)
	}()
	return ch
}
