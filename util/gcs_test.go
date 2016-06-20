package util

import (
	"net/url"
	"testing"
)

func TestNew(t *testing.T) {

	_, err := New("jkawamoto-ppls", "jkawamoto-ppls")
	if err != nil {
		t.Error(err.Error())
	}

}

func TestUpload(t *testing.T) {

	s, err := New("jkawamoto-ppls", "jkawamoto-ppls")
	if err != nil {
		t.Error(err.Error())
	}

	location, err := url.Parse("gs://jkawamoto-ppls/.roadie/gcs_test.go")
	if err != nil {
		t.Error(err.Error())
	}

	if err := s.Upload("./gcs_test.go", location); err != nil {
		t.Error(err.Error())
	}

}
