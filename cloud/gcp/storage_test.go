package gcp

import (
	"net/url"
	"testing"
)

// TestGetObjectName tests getObjectName converts URLs to correct names for
// Google Cloud Storage.
// Given URL must be in the following format;
// - roadie://category/path
// and the converted name must be in the following format;
// - .roadie/category/path
func TestGetObjectName(t *testing.T) {

	cases := []struct {
		input    string
		expected string
	}{
		{"roadie://cloud_test/test_dir/test_file", StoragePrefix + "/cloud_test/test_dir/test_file"},
		{"roadie://cloud_test/test_dir/", StoragePrefix + "/cloud_test/test_dir/"},
	}

	for _, c := range cases {

		loc, err := url.Parse(c.input)
		if err != nil {
			t.Fatalf("cannot parse URL %q: %v", c.input, err)
		}
		res := getObjectName(loc)
		if res != c.expected {
			t.Errorf("getObjectPath returns %v, want %v", res, c.expected)
		}

	}

}
