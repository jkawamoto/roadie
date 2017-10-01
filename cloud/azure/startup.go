//
// cloud/azure/startup.go
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
	"bytes"
	"encoding/base64"
	"text/template"

	yaml "gopkg.in/yaml.v2"

	"github.com/jkawamoto/roadie/assets"
	"github.com/jkawamoto/roadie/script"
)

const (
	// StartupTemplate is the asset name of the startup template.
	StartupTemplate = "assets/azure_startup.sh"
)

// startupScriptOpt defines options for creating a startup script.
type startupScriptOpt struct {
	// Config is a string representing a config in YAML.
	Config string
	// Script is a string representing a script in YAML.
	Script string
}

// StartupScript creates a base64 encoded string representing a starup script.
func StartupScript(cfg *AzureConfig, task *script.Script) (res string, err error) {

	data, err := assets.Asset(StartupTemplate)
	if err != nil {
		return
	}

	base, err := template.New("").Parse(string(data))
	if err != nil {
		return
	}

	cfgData, err := yaml.Marshal(cfg)
	if err != nil {
		return
	}
	taskData, err := yaml.Marshal(task)
	if err != nil {
		return
	}

	buf := bytes.Buffer{}
	err = base.ExecuteTemplate(&buf, "", &startupScriptOpt{
		Config: string(cfgData),
		Script: string(taskData),
	})
	if err != nil {
		return
	}

	res = base64.StdEncoding.EncodeToString(buf.Bytes())
	return

}
