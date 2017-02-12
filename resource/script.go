//
// resource/script.go
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

package resource

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jkawamoto/roadie/command/util"
	"gopkg.in/yaml.v2"
)

// ScriptBody defines the roadie script body.
type ScriptBody struct {
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
}

// Script defines a data structure of script file.
type Script struct {
	// ScriptBody.
	ScriptBody
	// Filename of the script.
	Filename string
	// InstanceName to run the script.
	InstanceName string
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
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost"
	} else if strings.Contains(hostname, ".") {
		hostname = strings.Split(hostname, ".")[0]
	}

	res = &Script{
		Filename: filename,
		InstanceName: strings.ToLower(fmt.Sprintf(
			"%s-%s-%s", hostname, util.Basename(filename), time.Now().Format("20060102150405"))),
	}

	// Unmarshal YAML file.
	if err = yaml.Unmarshal(buf.Bytes(), &res.ScriptBody); err != nil {
		return
	}
	return
}

// String converts this script to a string.
func (s *Script) String() string {
	res, _ := yaml.Marshal(s.ScriptBody)
	return string(res)
}
