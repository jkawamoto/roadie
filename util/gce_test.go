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

func TestSetZone(t *testing.T) {

	b, err := NewInstanceBuilder("jkawamoto-ppls")
	if err != nil {
		t.Error(err.Error())
	}

	for _, v := range []string{"us-central1-c", "projects/jkawamoto-ppls/zones/us-central1-c"} {
		b.SetZone(v)
		if b.Zone != "projects/jkawamoto-ppls/zones/us-central1-c" {
			t.Errorf("Zone is not correct: %s", b.Zone)
		}
	}

}
