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
	"io/ioutil"

	"github.com/jkawamoto/roadie/cloud/azure/auth"
	yaml "gopkg.in/yaml.v2"
)

// OSInformation defines OS information of creating instances.
type OSInformation struct {
	PublisherName string `yaml:"publisher_name,omitempty"`
	Offer         string `yaml:"offer,omitempty"`
	Skus          string `yaml:"skus,omitempty"`
	Version       string `yaml:"version,omitempty"`
}

// AzureConfig defines configuration to access Azure's API.
type AzureConfig struct {
	TenantID          string `yaml:"tenant_id,omitempty"`
	ClientID          string `yaml:"client_id,omitempty"`
	SubscriptionID    string `yaml:"subscription_id,omitempty"`
	Location          string `yaml:"location,omitempty"`
	ResourceGroupName string `yaml:"resource_group_name,omitempty"`
	MachineType       string `yaml:"machine_type,omitempty"`
	StorageAccount    string `yaml:"storage_account,omitempty"`
	OS                OSInformation
	Token             auth.Token
}

// NewAzureConfig creates a new AzureConfig with default values.
func NewAzureConfig() *AzureConfig {

	return &AzureConfig{
		ResourceGroupName: ComputeServiceResourceGroupName,
		MachineType:       ComputeServiceDefaultMachineType,
		StorageAccount:    DefaultStorageAccount,
		OS: OSInformation{
			PublisherName: DefaultOSPublisherName,
			Offer:         DefaultOSOffer,
			Skus:          DefaultOSSkus,
			Version:       DefaultOSVersion,
		},
	}

}

// NewAzureConfigFromFile creates a new AzureConfig from a file.
func NewAzureConfigFromFile(filename string) (cfg *AzureConfig, err error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	cfg = new(AzureConfig)
	err = yaml.Unmarshal(data, cfg)

	if cfg.ResourceGroupName == "" {
		cfg.ResourceGroupName = ComputeServiceResourceGroupName
	}
	if cfg.MachineType == "" {
		cfg.MachineType = ComputeServiceDefaultMachineType
	}
	if cfg.StorageAccount == "" {
		cfg.StorageAccount = DefaultStorageAccount
	}

	return

}

// WriteFilr writes this configuration to a file.
func (cfg *AzureConfig) WriteFile(filename string) (err error) {

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return
	}
	return ioutil.WriteFile(filename, data, 0644)

}

// String returns a string representing this configuration.
func (cfg *AzureConfig) String() (str string, err error) {

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return
	}
	str = string(data)
	return

}

// UnmarshalYAML unmarshals this configuration from a YAML document.
func (cfg *AzureConfig) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {

	var aux struct {
		TenantID          string `yaml:"tenant_id,omitempty"`
		ClientID          string `yaml:"client_id,omitempty"`
		SubscriptionID    string `yaml:"subscription_id,omitempty"`
		Location          string `yaml:"location,omitempty"`
		ResourceGroupName string `yaml:"resource_group_name,omitempty"`
		MachineType       string `yaml:"machine_type,omitempty"`
		StorageAccount    string `yaml:"storage_account,omitempty"`
		OS                OSInformation
		Token             auth.Token
	}
	err = unmarshal(&aux)
	if err != nil {
		return
	}
	*cfg = aux

	if cfg.ResourceGroupName == "" {
		cfg.ResourceGroupName = ComputeServiceResourceGroupName
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
