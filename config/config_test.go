package config

import "testing"

func TestLoadConfig(t *testing.T) {

	conf := LoadConfig("../.roadie")

	if conf.Gcp.Project != "jkawamoto-ppls" {
		t.Errorf("GCP Project is not correct: %s", conf.Gcp.Project)
	}
	if conf.Gcp.MachineType != "n1-standard-2" {
		t.Errorf("GCP MachieType is not correct: %s", conf.Gcp.MachineType)
	}
	if conf.Gcp.Zone != "us-central1-b" {
		t.Errorf("GCP Zone is not correct: %s", conf.Gcp.Zone)
	}
	if conf.Gcp.Bucket != "jkawamoto-ppls" {
		t.Errorf("GCP Bucket is not correct: %s", conf.Gcp.Bucket)
	}

}
