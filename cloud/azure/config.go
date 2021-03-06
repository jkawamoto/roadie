//
// cloud/azure/config.go
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
	"hash/fnv"
	"io/ioutil"

	"github.com/Azure/go-autorest/autorest/adal"
	yaml "gopkg.in/yaml.v2"
)

// TODO: Support async to not wait finishing each operation.

// OSInformation defines OS information of creating instances.
type OSInformation struct {
	PublisherName string `yaml:"publisher_name,omitempty"`
	Offer         string `yaml:"offer,omitempty"`
	Skus          string `yaml:"skus,omitempty"`
	Version       string `yaml:"version,omitempty"`
}

// Config defines configuration to access Azure's API.
type Config struct {
	TenantID       string `yaml:"tenant_id,omitempty"`
	SubscriptionID string `yaml:"subscription_id,omitempty"`
	Location       string `yaml:"location,omitempty"`
	ProjectID      string `yaml:"project_id,omitempty"`
	AccountName    string `yaml:"-"`
	MachineType    string `yaml:"machine_type,omitempty"`
	OS             OSInformation
	Token          adal.Token
}

// NewConfig creates a new Config with default values.
func NewConfig() *Config {

	return &Config{
		Location:    DefaultLocation,
		MachineType: ComputeServiceDefaultMachineType,
		OS: OSInformation{
			PublisherName: DefaultOSPublisherName,
			Offer:         DefaultOSOffer,
			Skus:          DefaultOSSkus,
			Version:       DefaultOSVersion,
		},
	}

}

// NewConfigFromFile creates a new Config from a file.
func NewConfigFromFile(filename string) (cfg *Config, err error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	cfg = new(Config)
	err = yaml.Unmarshal(data, cfg)
	if cfg.Location == "" {
		cfg.Location = DefaultLocation
	}
	if cfg.MachineType == "" {
		cfg.MachineType = ComputeServiceDefaultMachineType
	}
	if cfg.OS.PublisherName == "" {
		cfg.OS.PublisherName = DefaultOSPublisherName
	}
	if cfg.OS.Offer == "" {
		cfg.OS.Offer = DefaultOSOffer
	}
	if cfg.OS.Skus == "" {
		cfg.OS.Skus = DefaultOSSkus
	}
	if cfg.OS.Version == "" {
		cfg.OS.Version = DefaultOSVersion
	}

	cfg.updateAccountName()
	return

}

// updateAccountName updates AccountName combining ProjectID and Location.
// When those values are modified, this function must be called.
func (cfg *Config) updateAccountName() {
	hash := fnv.New32()
	hash.Write([]byte(cfg.Location))
	cfg.AccountName = fmt.Sprintf("%v%x", cfg.ProjectID, hash.Sum32())
}

// Valid returns true if this config has values in required fields; otherwise
// return false.
// The required fields are
// - TenantID
// - SubscriptionID
// - ProjectID
func (cfg *Config) Valid() bool {

	if cfg.TenantID == "" || cfg.SubscriptionID == "" || cfg.ProjectID == "" {
		return false
	}
	return true

}

// WriteFile writes this configuration to a file.
func (cfg *Config) WriteFile(filename string) (err error) {

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return
	}
	return ioutil.WriteFile(filename, data, 0644)

}

// String returns a string representing this configuration.
func (cfg *Config) String() (str string, err error) {

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return
	}
	str = string(data)
	return

}

// UnmarshalYAML unmarshals this configuration from a YAML document.
func (cfg *Config) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {

	var aux struct {
		TenantID       string `yaml:"tenant_id,omitempty"`
		SubscriptionID string `yaml:"subscription_id,omitempty"`
		Location       string `yaml:"location,omitempty"`
		ProjectID      string `yaml:"project_id,omitempty"`
		AccountName    string `yaml:"-"`
		MachineType    string `yaml:"machine_type,omitempty"`
		OS             OSInformation
		Token          adal.Token
	}
	err = unmarshal(&aux)
	if err != nil {
		return
	}
	*cfg = aux

	if cfg.Location == "" {
		cfg.Location = DefaultLocation
	}
	if cfg.MachineType == "" {
		cfg.MachineType = ComputeServiceDefaultMachineType
	}
	cfg.updateAccountName()

	return

}

// UnmarshalYAML unmarshals configuration form a YAML document.
func (info *OSInformation) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {

	var aux struct {
		PublisherName string `yaml:"publisher_name,omitempty"`
		Offer         string `yaml:"offer,omitempty"`
		Skus          string `yaml:"skus,omitempty"`
		Version       string `yaml:"version,omitempty"`
	}
	err = unmarshal(&aux)
	if err != nil {
		return
	}
	*info = aux

	if info.PublisherName == "" {
		info.PublisherName = DefaultOSPublisherName
	}
	if info.Offer == "" {
		info.Offer = DefaultOSOffer
	}
	if info.Skus == "" {
		info.Skus = DefaultOSSkus
	}
	if info.Version == "" {
		info.Version = DefaultOSVersion
	}

	return

}
