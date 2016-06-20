package util

import (
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

const scope = compute.ComputeScope

type InstanceBuilder struct {
	Project string
	Zone    string
	service *compute.Service
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
		Project: project,
		service: service,
	}
	res.SetZone("us-central1-b")
	return res, nil

}

// SetZone sets a given zone to the builder object.
func (b *InstanceBuilder) SetZone(zone string) {

	if strings.HasPrefix(zone, "projects") {
		b.Zone = zone
	} else {
		b.Zone = "projects/" + b.Project + "/zones/" + zone
	}

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

func NewInstance(project, zone string) (error, error) {

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
		Name:        "one-1",
		Zone:        "projects/jkawamoto-ppls/zones/us-central1-b",
		MachineType: "projects/jkawamoto-ppls/zones/us-central1-b/machineTypes/n1-standard-2",
		Disks: []*compute.AttachedDisk{
			&compute.AttachedDisk{
				Type:       "PERSISTENT",
				Boot:       true,
				Mode:       "READ_WRITE",
				AutoDelete: true,
				InitializeParams: &compute.AttachedDiskInitializeParams{
					SourceImage: "https://www.googleapis.com/compute/v1/projects/coreos-cloud/global/images/coreos-stable-1010-5-0-v20160527",
					DiskType:    "projects/jkawamoto-ppls/zones/us-central1-b/diskTypes/pd-standard",
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

	service.Instances.Insert(project, zone, &bluepring)
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
