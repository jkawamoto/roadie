package util

import "testing"

func TestBasename(t *testing.T) {

	if res := Basename("somefile.go"); res != "somefile" {
		t.Errorf("%s does not match 'somefile'", res)
	}

	if res := Basename("noext"); res != "noext" {
		t.Errorf("%s does not match 'noext'", res)
	}

	if res := Basename("/path/to/somefile.go"); res != "somefile" {
		t.Errorf("%s does not match 'somefile'", res)
	}

}

func TestCreateURL(t *testing.T) {

	u := CreateURL("bucket_name", "source", "/path/to/file")
	if u.Scheme != "gs" {
		t.Errorf("Scheme is not correct: %s", u.Scheme)
	}
	if u.Host != "bucket_name" {
		t.Errorf("Host name is not correct: %s", u.Host)
	}
	if u.Path != "/.roadie/source/path/to/file" {
		t.Errorf("Path is not correct: %s", u.Path)
	}

}
