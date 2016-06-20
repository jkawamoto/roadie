package util

import (
	"log"

	"github.com/ttacon/chalk"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

const gceScope = compute.CloudPlatformScope

// InstanceBuilder maintains configurations to create new instances.
type InstanceBuilder struct {
	Project     string
	Zone        string
	MachineType string
	service     *compute.Service
}

// MetadataItem has Key and Value properties.
type MetadataItem struct {
	Key   string
	Value string
}

// NewInstanceBuilder creates a new instance builder associated with
// a given project.
func NewInstanceBuilder(project string) (*InstanceBuilder, error) {

	// Create a client.
	client, err := google.DefaultClient(context.Background(), gceScope)
	if err != nil {
		return nil, err
	}

	// Create a servicer.
	service, err := compute.New(client)
	if err != nil {
		return nil, err
	}

	res := &InstanceBuilder{
		Project:     project,
		Zone:        "us-central1-b",
		MachineType: "n1-standard-1",
		service:     service,
	}
	return res, nil

}

// AvailableZones returns a slice of zone names.
func (b *InstanceBuilder) AvailableZones() ([]string, error) {

	res, err := b.service.Zones.List(b.Project).Do()
	if err != nil {
		return nil, err
	}

	zones := make([]string, len(res.Items))
	for i, v := range res.Items {
		zones[i] = v.Name
	}

	return zones, nil

}

// AvailableMachineTypes returns a slice of machie type names.
func (b *InstanceBuilder) AvailableMachineTypes() ([]string, error) {

	res, err := b.service.MachineTypes.List(b.Project, "us-central1-b").Do()
	if err != nil {
		return nil, err
	}

	types := make([]string, len(res.Items))
	for i, v := range res.Items {
		types[i] = v.Name
	}

	return types, nil

}

// CreateInstance creates a new instance based on the bilder's configuration.
func (b *InstanceBuilder) CreateInstance(name string, metadata []*MetadataItem) (err error) {

	matadataItems := make([]*compute.MetadataItems, len(metadata))
	for i, v := range metadata {
		matadataItems[i] = &compute.MetadataItems{
			Key:   v.Key,
			Value: &v.Value,
		}
	}

	bluepring := compute.Instance{
		Name:        name,
		Zone:        b.normalizedZone(),
		MachineType: b.normalizedMachineType(),
		Disks: []*compute.AttachedDisk{
			&compute.AttachedDisk{
				Type:       "PERSISTENT",
				Boot:       true,
				Mode:       "READ_WRITE",
				AutoDelete: true,
				InitializeParams: &compute.AttachedDiskInitializeParams{
					SourceImage: "https://www.googleapis.com/compute/v1/projects/coreos-cloud/global/images/coreos-stable-1010-5-0-v20160527",
					DiskType:    b.normalizedZone() + "/diskTypes/pd-standard",
					DiskSizeGb:  9,
				},
			},
		},
		CanIpForward: false,
		NetworkInterfaces: []*compute.NetworkInterface{
			&compute.NetworkInterface{
				Network: "projects/" + b.Project + "/global/networks/default",
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

	res, err := b.service.Instances.Insert(b.Project, b.Zone, &bluepring).Do()
	if err == nil {
		if res.StatusMessage != "" {
			log.Println(res.StatusMessage)
		}
		for _, v := range res.Warnings {
			log.Println(chalk.Red.Color(v.Message))
		}
	}
	return

}

// StopInstance stops a given named instance.
func (b *InstanceBuilder) StopInstance(name string) (err error) {
	res, err := b.service.Instances.Stop(b.Project, b.Zone, name).Do()
	if err == nil {
		if res.StatusMessage != "" {
			log.Println(res.StatusMessage)
		}
		for _, v := range res.Warnings {
			log.Println(chalk.Red.Color(v.Message))
		}
	}
	return
}

// normalizedZone returns the normalized zone string of Zone property.
func (b *InstanceBuilder) normalizedZone() string {
	return "projects/" + b.Project + "/zones/" + b.Zone
}

// normalizedMachineType returns the normalized instance type of MachineType property.
func (b *InstanceBuilder) normalizedMachineType() string {
	return b.normalizedZone() + "/machineTypes/" + b.MachineType
}
