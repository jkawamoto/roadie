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

	"github.com/briandowns/spinner"
	"github.com/jkawamoto/roadie/util"
	"github.com/ttacon/chalk"
	"gopkg.in/yaml.v2"
)

type script struct {
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
func loadScript(filename string, args []string) (*script, error) {

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

	res := script{
		filename: filename,
		instanceName: fmt.Sprintf(
			"%s-%s-%s", hostname, util.Basename(filename), time.Now().Format("20060102150405")),
	}

	// Unmarshal YAML file.
	if err := yaml.Unmarshal(buf.Bytes(), &res.body); err != nil {
		return nil, err
	}

	return &res, nil

}

// Set a git repository to source section.
func (s *script) setGitSource(repo string) {
	if s.body.Source != "" {
		fmt.Printf(
			chalk.Red.Color("The source section of %s will be overwritten to '%s' since a Git repository is given.\n"),
			s.filename, repo)
	}
	s.body.Source = repo
}

// Set a URL to source section.
func (s *script) setURLSource(url string) {
	if s.body.Source != "" {
		fmt.Printf(
			chalk.Red.Color("The source section of %s will be overwritten to '%s' since a repository URL is given.\n"),
			s.filename, url)
	}
	s.body.Source = url
}

// Upload source files and set that location to source section.
func (s *script) setLocalSource(path, project, bucket string) error {
	if s.body.Source != "" {
		fmt.Printf(
			chalk.Red.Color("The source section of %s is overwritten since a path for source codes is given.\n"),
			s.filename)
	}

	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	var name string
	var arcPath string
	if info.IsDir() {

		filename := s.instanceName + ".tar.gz"
		arcPath = filepath.Join(os.TempDir(), filename)

		spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		spin.Prefix = fmt.Sprintf("Creating an archived file %s...", arcPath)
		spin.FinalMSG = fmt.Sprintf("\n%s\rCreating the archived file %s.    \n", strings.Repeat(" ", len(spin.Prefix)+2), arcPath)
		spin.Start()
		if err := util.Archive(path, arcPath, nil); err != nil {
			spin.Stop()
			return err
		}
		name = filename
		spin.Stop()

	} else {

		arcPath = path
		name = util.Basename(path)

	}

	location, err := UploadToGCS(project, bucket, SourcePrefix, name, arcPath)
	if err != nil {
		return err
	}
	s.body.Source = location
	return nil

}

// Set result section with a given bucket name.
func (s *script) setResult(bucket string) {

	location := util.CreateURL(bucket, ResultPrefix, s.instanceName)
	s.body.Result = location.String()

}

// Convert to string.
func (s *script) String() string {
	res, _ := yaml.Marshal(s.body)
	return string(res)
}
