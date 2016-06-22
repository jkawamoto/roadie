package util

import (
	"net/url"
	"path/filepath"
)

// Basename computes the basename of a given filename.
func Basename(filename string) string {

	ext := filepath.Ext(filename)
	bodySize := len(filename) - len(ext)

	return filepath.Base(filename[:bodySize])

}

// CreateURL creates a valid URL for uploaing object.
func CreateURL(bucket, group, name string) *url.URL {

	return &url.URL{
		Scheme: "gs",
		Host:   bucket,
		Path:   filepath.ToSlash(filepath.Join("/.roadie", group, name)),
	}

}
