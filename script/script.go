//
// script/script.go
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

package script

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	// RoadieSchemePrefix is the prefix of roadie scheme URLs.
	RoadieSchemePrefix = "roadie://"
	// SourcePrefix defines a prefix to store source files.
	SourcePrefix = "source"
	// DataPrefix defines a prefix to store data files.
	DataPrefix = "data"
	// ResultPrefix defines a prefix to store result files.
	ResultPrefix = "result"
)

// Script defines roadie's script format.
type Script struct {
	// List of apt packages to be installed.
	APT []string `yaml:"apt,omitempty"`
	// URL to the source code.
	Source string `yaml:"source,omitempty"`
	// List of URLs to be downloaded as data files.
	Data []string `yaml:"data,omitempty"`
	// List of commands to be run.
	Run []string `yaml:"run,omitempty"`
	// URL where the computational results will be stored.
	Result string `yaml:"result,omitempty"`
	// List of glob pattern, files matches of one of them are uploaded as resuts.
	Upload []string `yaml:"upload,omitempty"`
	// List of option flags.
	Options []string `yaml:"options,omitempty"`

	// ** Following attributes are given by roadie. **
	// InstanceName to run the script.
	InstanceName string `yaml:"instance_name,omitempty"`
	// Image is a docker image name used to run this script.
	Image string
}

// NewScript loads a given script file and apply arguments.
func NewScript(filename string, args []string) (res *Script, err error) {

	// Define function map to replace place holders.
	funcs := template.FuncMap{}
	for _, v := range args {
		sp := strings.Split(v, "=")
		if len(sp) >= 2 {
			funcs[sp[0]] = func() string {
				return sp[1]
			}
		}
	}

	// Load YAML config file.
	conf, err := template.New(filepath.Base(filename)).Funcs(funcs).ParseFiles(filename)
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			return
		default:
			return nil, fmt.Errorf("Cannot apply variables to the place holders in %s", filename)
		}
	}

	// Replace place holders with given args.
	buf := &bytes.Buffer{}
	if err = conf.Execute(buf, nil); err != nil {
		return
	}

	// Construct a script object.
	res = &Script{}

	// Unmarshal YAML file.
	if err = yaml.Unmarshal(buf.Bytes(), res); err != nil {
		return
	}

	res.InstanceName = strings.ToLower(
		strings.Replace(fmt.Sprintf(
			"%s-%s", basename(filename), time.Now().Format("20060102150405")),
			".", "-", -1))
	return

}

// String converts this script to a string.
func (s *Script) String() string {
	res, _ := yaml.Marshal(s)
	return string(res)
}
