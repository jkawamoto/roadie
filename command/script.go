//
// command/script.go
//
// Copyright (c) 2016 Junpei Kawamoto
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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package command

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jkawamoto/roadie/util"
	"gopkg.in/yaml.v2"
)

type Script struct {
	filename     string
	instanceName string
	body         struct {
		APT    []string `yaml:"apt,omitempty"`
		Source string   `yaml:"source,omitempty"`
		Data   []string `yaml:"data,omitempty"`
		Run    []string `yaml:"run,omitempty"`
		Result string   `yaml:"result,omitempty"`
		Upload []string `yaml:"upload,omitempty"`
	}
}

// Load a given script file and apply arguments.
func loadScript(filename string, args []string) (*Script, error) {

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
			return nil, err
		default:
			return nil, fmt.Errorf("Cannot apply variables to the place holders in %s", filename)
		}
	}

	// Replace place holders with given args.
	buf := &bytes.Buffer{}
	if err := conf.Execute(buf, nil); err != nil {
		return nil, err
	}

	// Construct a script object.
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost"
	} else if strings.Contains(hostname, ".") {
		hostname = strings.Split(hostname, ".")[0]
	}

	res := Script{
		filename: filename,
		instanceName: strings.ToLower(fmt.Sprintf(
			"%s-%s-%s", hostname, util.Basename(filename), time.Now().Format("20060102150405"))),
	}

	// Unmarshal YAML file.
	if err := yaml.Unmarshal(buf.Bytes(), &res.body); err != nil {
		return nil, err
	}

	return &res, nil

}

// Set result section with a given bucket name.
func (s *Script) setResult(bucket string) {

	location := util.CreateURL(bucket, ResultPrefix, s.instanceName)
	s.body.Result = location.String()

}

// Convert to string.
func (s *Script) String() string {
	res, _ := yaml.Marshal(s.body)
	return string(res)
}
