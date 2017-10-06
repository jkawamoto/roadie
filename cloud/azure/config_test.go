//
// cloud/azure/config_test.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
//
// This file is part of Roadie.
//
// Roadie is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Roadie is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Roadie.  If not, see <http://www.gnu.org/licenses/>.
//

package azure

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewConfig(t *testing.T) {

	cfg := NewConfig()
	if cfg.Location != DefaultLocation {
		t.Errorf("location is %q, want %v", cfg.Location, DefaultLocation)
	}
	if cfg.OS.PublisherName != DefaultOSPublisherName {
		t.Errorf("publisher name is %q, want %v", cfg.OS.PublisherName, DefaultOSPublisherName)
	}
	if cfg.OS.Offer != DefaultOSOffer {
		t.Errorf("offer is %q, want %v", cfg.OS.Offer, DefaultOSOffer)
	}
	if cfg.OS.Skus != DefaultOSSkus {
		t.Errorf("skus is %q, want %v", cfg.OS.Skus, DefaultOSSkus)
	}
	if cfg.OS.Version != DefaultOSVersion {
		t.Errorf("version is %q, want %v", cfg.OS.Version, DefaultOSVersion)
	}

}

func TestReadWriteConfigFile(t *testing.T) {

	var err error
	cfg := Config{
		ProjectID: "test-project",
	}
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("cannot create a temporary directory: %v", err)
	}
	defer os.RemoveAll(dir)

	filename := filepath.Join(dir, "azure_config.yml")
	err = cfg.WriteFile(filename)
	if err != nil {
		t.Fatalf("cannot save a config file to %v: %v", filename, err)
	}

	res, err := NewConfigFromFile(filename)
	if err != nil {
		t.Fatalf("cannote read the configu file %v: %v", filename, err)
	}
	if res.Location != DefaultLocation {
		t.Errorf("location is %q, want %v", res.Location, DefaultLocation)
	}
	if res.OS.PublisherName != DefaultOSPublisherName {
		t.Errorf("publisher name is %q, want %v", res.OS.PublisherName, DefaultOSPublisherName)
	}
	if res.OS.Offer != DefaultOSOffer {
		t.Errorf("offer is %q, want %v", res.OS.Offer, DefaultOSOffer)
	}
	if res.OS.Skus != DefaultOSSkus {
		t.Errorf("skus is %q, want %v", res.OS.Skus, DefaultOSSkus)
	}
	if res.OS.Version != DefaultOSVersion {
		t.Errorf("version is %q, want %v", res.OS.Version, DefaultOSVersion)
	}
	if !strings.HasPrefix(res.AccountName, res.ProjectID) {
		t.Errorf("account name dosen't have the project ID %q as the prefix: %v", res.ProjectID, res.AccountName)
	}

}
