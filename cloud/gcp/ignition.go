//
// cloud/gcp/ignition.go
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

package gcp

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

var (
	// RegexpCommentLine defines a regular expression for a comment line.
	RegexpCommentLine = regexp.MustCompile("#.*\n")
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

// loadUnitTemplate loads a given named template and applies opt.
func loadUnitTemplate(name string, opt interface{}) (str string, err error) {

	data, err := assets.Asset(name)
	if err != nil {
		return
	}

	buf := &bytes.Buffer{}
	temp, err := template.New("").Parse(string(data))
	if err != nil {
		return
	}
	err = temp.ExecuteTemplate(buf, "", opt)
	if err != nil {
		return
	}

	str = RegexpCommentLine.ReplaceAllString(buf.String(), "")
	return

}

// fluentdUnitOpt defines options for processing a template of fluentd.service.
type fluentdUnitOpt struct {
	Name string
}

// FluentdUnit creates a new SystemdUnit for fluentd.
func FluentdUnit(name string) (unit SystemdUnit, err error) {

	unit.Name = "fluentd.service"
	unit.Enable = true
	unit.Contents, err = loadUnitTemplate("assets/fluentd.service", &fluentdUnitOpt{
		Name: name,
	})
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

	unit.Name = "roadie.service"
	unit.Enable = true
	unit.Contents, err = loadUnitTemplate("assets/roadie.service", &roadieUnitOpt{
		Name:    name,
		Image:   image,
		Options: options,
	})
	return

}

// logcaseUnitOpt defines options for processing a template of logcast.service.
type logcaseUnitOpt struct {
	Service string
}

// LogcastUnit creates a new unit for forwarding log to Fluentd.
func LogcastUnit(service string) (unit SystemdUnit, err error) {

	unit.Name = "logcast.service"
	unit.Enable = true
	unit.Contents, err = loadUnitTemplate("assets/logcast.service", &logcaseUnitOpt{
		Service: service,
	})
	return

}

// queueManagerUnitOpt defines options for queue manager service unit.
type queueManagerUnitOpt struct {
	Project   string
	Version   string
	QueueName string
}

// QueueManagerUnit creates a new unit for Roadie queue manager.
func QueueManagerUnit(project, version, queueName string) (unit SystemdUnit, err error) {

	unit.Name = "queue.service"
	unit.Enable = true
	unit.Contents, err = loadUnitTemplate("assets/queue.service", &queueManagerUnitOpt{
		Project:   project,
		Version:   version,
		QueueName: queueName,
	})
	return

}
