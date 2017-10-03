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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/Azure/go-autorest/autorest/adal"
)

func TestNewConfig(t *testing.T) {

	cfg := NewConfig()
	if cfg.OS.PublisherName != DefaultOSPublisherName {
		t.Error("Default publisher name is not correct:", cfg.OS.PublisherName)
	}
	if cfg.OS.Offer != DefaultOSOffer {
		t.Error("Default offer is not correct:", cfg.OS.Offer)
	}
	if cfg.OS.Skus != DefaultOSSkus {
		t.Error("Default skus is not correct:", cfg.OS.Skus)
	}
	if cfg.OS.Version != DefaultOSVersion {
		t.Error("Default version is not correct:", cfg.OS.Version)
	}

}

func TestNewConfigFromFile(t *testing.T) {

	var err error

	cfg := NewConfig()
	cfg.ResourceGroupName = ""
	cfg.OS.PublisherName = ""
	cfg.OS.Offer = ""
	cfg.OS.Skus = ""
	cfg.OS.Version = ""

	filename := filepath.Join(os.TempDir(), fmt.Sprintf("azure_config%v.yml", time.Now().Unix()))
	err = cfg.WriteFile(filename)
	if err != nil {
		t.Error(err.Error())
	}

	res, err := NewConfigFromFile(filename)
	if err != nil {
		t.Error(err.Error())
	}
	if res.ResourceGroupName != ComputeServiceResourceGroupName {
		t.Error("Default resource group name is not correct:", res.ResourceGroupName)
	}
	if res.OS.PublisherName != DefaultOSPublisherName {
		t.Error("Default publisher name is not correct:", res.OS.PublisherName)
	}
	if res.OS.Offer != DefaultOSOffer {
		t.Error("Default offer is not correct:", res.OS.Offer)
	}
	if res.OS.Skus != DefaultOSSkus {
		t.Error("Default skus is not correct:", res.OS.Skus)
	}
	if res.OS.Version != DefaultOSVersion {
		t.Error("Default version is not correct:", res.OS.Version)
	}

}

func TestConfigWriteFile(t *testing.T) {

	var err error

	cfg := NewConfig()
	cfg.SubscriptionID = "subscription"
	cfg.ResourceGroupName = "resource"
	cfg.Location = "location"
	cfg.Token = adal.Token{
		AccessToken: "token",
	}

	filename := filepath.Join(os.TempDir(), fmt.Sprintf("azure_config%v.yml", time.Now().Unix()))
	err = cfg.WriteFile(filename)
	if err != nil {
		t.Error(err.Error())
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(string(data))

}
