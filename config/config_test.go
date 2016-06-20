package config

import "testing"

func TestLoadConfig(t *testing.T) {

	conf, err := LoadConfig("../.roadie")
	if err != nil {
		t.Error(err.Error())
	}

	if conf.GCP.Project != "jkawamoto-ppls" {
		t.Errorf("GCP Project is not correct: %s", conf.GCP.Project)
	}
	if conf.GCP.MachineType != "n1-standard-2" {
		t.Errorf("GCP MachieType is not correct: %s", conf.GCP.MachineType)
	}
	if conf.GCP.Zone != "us-central1-b" {
		t.Errorf("GCP Zone is not correct: %s", conf.GCP.Zone)
	}
	if conf.GCP.Bucket != "jkawamoto-ppls" {
		t.Errorf("GCP Bucket is not correct: %s", conf.GCP.Bucket)
	}

}
