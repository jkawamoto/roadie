//
// cloud/gce/ignition.go
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

package gce

import (
	"bytes"
	"encoding/json"
	"regexp"
	"text/template"

	"github.com/jkawamoto/roadie/assets"
)

const (
	// DefaultIgnitionVersion defines the version of ignition configuration.
	DefaultIgnitionVersion = "2.0.0"
)

// SystemdUnit is configuration for a systemd unit.
type SystemdUnit struct {
	Name     string `json:"name,omitempty"`
	Enable   bool   `json:"enable,omitempty"`
	Contents string `json:"contents,omitempty"`
}

// IgnitionVersion is configuration of ignition version.
type IgnitionVersion struct {
	Version string `json:"version,omitempty"`
}

// SystemdConfig defines a set of units to be set up.
type SystemdConfig struct {
	Units []SystemdUnit `json:"units,omitempty"`
}

// IgnitionConfig defines a simple structure for ignition config file which has
// only systemd's unit configurations.
//
// {
//   "ignition": { "version": "2.0.0" },
//   "systemd": {
//     "units": [{
//       "name": "example.service",
//       "enable": true,
//       "contents": "[Service]\nType=oneshot\nExecStart=/usr/bin/echo Hello World\n\n[Install]\nWantedBy=multi-user.target"
//     }]
//   }
// }
type IgnitionConfig struct {
	Ignition IgnitionVersion `json:"ignition,omitempty"`
	Systemd  SystemdConfig   `json:"systemd,omitempty"`
}

// NewIgnitionConfig creates a new ignition configuration.
func NewIgnitionConfig() *IgnitionConfig {

	return &IgnitionConfig{
		Ignition: IgnitionVersion{
			Version: DefaultIgnitionVersion,
		},
	}

}

// String marshals this config to a JSON string.
func (c *IgnitionConfig) String() string {

	// data, err := json.Marshal(c)
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		// This error can be solved during developping phase.
		panic(err)
	}

	return string(data)

}

// Append a new unit and returns updated this configuration.
func (c *IgnitionConfig) Append(unit SystemdUnit) *IgnitionConfig {

	c.Systemd.Units = append(c.Systemd.Units, unit)
	return c

}

// fluentdUnitOpt defines options for processing a template of fluentd.service.
type fluentdUnitOpt struct {
	Name string
}

// FluentdUnit creates a new SystemdUnit for fluentd.
func FluentdUnit(name string) (unit SystemdUnit, err error) {

	data, err := assets.Asset("assets/fluentd.service")
	if err != nil {
		return
	}

	buf := &bytes.Buffer{}
	temp, err := template.New("fluentd").Parse(string(data))
	if err != nil {
		return
	}
	err = temp.ExecuteTemplate(buf, "fluentd", &fluentdUnitOpt{
		Name: name,
	})
	if err != nil {
		return
	}

	unit.Name = "fluentd.service"
	unit.Enable = true
	unit.Contents = removeComments(buf.String())
	return

}

// roadieUnitOpt defines options for processing a template of roadie.service.
type roadieUnitOpt struct {
	Name    string
	Image   string
	Options string
}

// RoadieUnit creates a new SystemdUnit for roadie-gcp.
func RoadieUnit(name, image, options string) (unit SystemdUnit, err error) {

	data, err := assets.Asset("assets/roadie.service")
	if err != nil {
		return
	}

	buf := &bytes.Buffer{}
	temp, err := template.New("roadie").Parse(string(data))
	if err != nil {
		return
	}
	err = temp.ExecuteTemplate(buf, "roadie", &roadieUnitOpt{
		Name:    name,
		Image:   image,
		Options: options,
	})
	if err != nil {
		return
	}

	unit.Name = "roadie.service"
	unit.Enable = true
	unit.Contents = removeComments(buf.String())
	return

}

// LogcastUnit creates a new unit for forwarding log to Fluentd.
func LogcastUnit() (unit SystemdUnit, err error) {

	data, err := assets.Asset("assets/logcast.service")
	if err != nil {
		return
	}

	unit.Name = "logcast.service"
	unit.Enable = true
	unit.Contents = removeComments(string(data))
	return

}

func removeComments(str string) string {
	r := regexp.MustCompile("#.*\n")
	return r.ReplaceAllString(str, "")
}
