package util

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

const scope = compute.ComputeScope

type InstanceBuilder struct {
	Project     string
	Zone        string
	MachineType string
	service     *compute.Service
}

// NewInstanceManager creates a new instance manager.
func NewInstanceBuilder(project string) (*InstanceBuilder, error) {

	// Create a client.
	client, err := google.DefaultClient(context.Background(), scope)
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

func (b *InstanceBuilder) CreateInstance(name string) (error, error) {

	// POST https://www.googleapis.com/compute/v1/projects/jkawamoto-ppls/zones/us-central1-b/instances
	// {
	//   "name": "instance-1",
	//   "zone": "projects/jkawamoto-ppls/zones/us-central1-b",
	//   "machineType": "projects/jkawamoto-ppls/zones/us-central1-b/machineTypes/n1-standard-2",
	//   "metadata": {
	//     "items": []
	//   },
	//   "tags": {
	//     "items": []
	//   },
	//   "disks": [
	//     {
	//       "type": "PERSISTENT",
	//       "boot": true,
	//       "mode": "READ_WRITE",
	//       "autoDelete": true,
	//       "deviceName": "instance-1",
	//       "initializeParams": {
	//         "sourceImage": "https://www.googleapis.com/compute/v1/projects/coreos-cloud/global/images/coreos-stable-1010-5-0-v20160527",
	//         "diskType": "projects/jkawamoto-ppls/zones/us-central1-b/diskTypes/pd-standard",
	//         "diskSizeGb": "9"
	//       }
	//     }
	//   ],
	//   "canIpForward": false,
	//   "networkInterfaces": [
	//     {
	//       "network": "projects/jkawamoto-ppls/global/networks/default",
	//       "accessConfigs": [
	//         {
	//           "name": "External NAT",
	//           "type": "ONE_TO_ONE_NAT"
	//         }
	//       ]
	//     }
	//   ],
	//   "description": "",
	//   "scheduling": {
	//     "preemptible": false,
	//     "onHostMaintenance": "MIGRATE",
	//     "automaticRestart": true
	//   },
	//   "serviceAccounts": [
	//     {
	//       "email": "default",
	//       "scopes": [
	//         "https://www.googleapis.com/auth/cloud-platform"
	//       ]
	//     }
	//   ]
	// }

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
				Network: "projects/jkawamoto-ppls/global/networks/default",
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
	}

	// gcloud compute --project "{project}" instances create "{instance_name}" \
	//     --zone us-central1-b --image coreos --metadata exec="{exe}" \
	//     --metadata-from-file startup-script={startup} \
	//     --machine-type "n1-standard-2" \
	//

	b.service.Instances.Insert(b.Project, b.Zone, &bluepring)
	// if _, err := service.Buckets.Get(bucket).Do(); err != nil {
	//
	// 	if res, err := service.Buckets.Insert(project, &storage.Bucket{Name: bucket}).Do(); err == nil {
	// 		log.Printf("Bucket %s was created at %s", res.Name, res.SelfLink)
	// 	} else {
	// 		return nil, err
	// 	}
	//
	// }
	//
	// return &Storage{
	// 	BucketName: bucket,
	// 	client:     client,
	// 	service:    service,
	// }, nil
	return nil, nil

}

// normalizedZone returns the normalized zone string of Zone property.
func (b *InstanceBuilder) normalizedZone() string {
	return "projects/" + b.Project + "/zones/" + b.Zone
}

// normalizedMachineType returns the normalized instance type of MachineType property.
func (b *InstanceBuilder) normalizedMachineType() string {
	return b.normalizedZone() + "/machineTypes/" + b.MachineType
}
