package util

import "testing"

func TestAvailableZones(t *testing.T) {

	b, err := NewInstanceBuilder("jkawamoto-ppls")
	if err != nil {
		t.Error(err.Error())
	}

	zones, err := b.AvailableZones()
	if err != nil {
		t.Error(err.Error())
	}

	for _, v := range zones {
		t.Logf("Available zone: %s", v)
	}

}

func TestAvailableMachineTypes(t *testing.T) {

	b, err := NewInstanceBuilder("jkawamoto-ppls")
	if err != nil {
		t.Error(err.Error())
	}

	types, err := b.AvailableMachineTypes()
	if err != nil {
		t.Error(err.Error())
	}

	for _, v := range types {
		t.Logf("Available machine type: %s", v)
	}

}

func TestNormalizedZone(t *testing.T) {

	b, err := NewInstanceBuilder("jkawamoto-ppls")
	if err != nil {
		t.Error(err.Error())
	}

	b.Zone = "us-central1-c"
	if b.normalizedZone() != "projects/jkawamoto-ppls/zones/us-central1-c" {
		t.Errorf("Zone is not correct: %s", b.Zone)
	}

}

func TestNormalizedMachineType(t *testing.T) {

	b, err := NewInstanceBuilder("jkawamoto-ppls")
	if err != nil {
		t.Error(err.Error())
	}

	b.MachineType = "n1-standard-2"
	if b.normalizedMachineType() != "projects/jkawamoto-ppls/zones/us-central1-b/machineTypes/n1-standard-2" {
		t.Errorf("Zone is not correct: %s", b.Zone)
	}

}
