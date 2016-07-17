//
// config/config.go
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

package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jkawamoto/roadie/chalk"
	"github.com/naoina/toml"
)

// ConfigureFile defines configuration file name.
const ConfigureFile = ".roadie"

// DotGit defines a git repository name.
const DotGit = ".git"

// Config defines a structure of config file.
type Config struct {
	Filename string `toml:"-"`
	Gcp      struct {
		Project     string
		MachineType string
		Zone        string
		Bucket      string
	}
}

// NewConfig creates a config object. If there is a configure file,
// it also loads the file, too.
func NewConfig() (c *Config, err error) {

	c = &Config{
		Filename: lookup(),
	}

	if _, exist := os.Stat(c.Filename); exist == nil {

		f, err := os.Open(c.Filename)
		if err != nil {
			return nil, fmt.Errorf(
				chalk.Red.Color("Cannot open configuration file %s. (%s)"),
				c.Filename, err.Error())
		}
		defer f.Close()

		var buf []byte
		if buf, err = ioutil.ReadAll(f); err == nil {
			err = toml.Unmarshal(buf, c)
		}
		if err != nil {
			return nil, fmt.Errorf(
				chalk.Red.Color("Configuration file %s is broken. Fix or delete it, first. (%s)"),
				c.Filename, err.Error())
		}

	}

	return

}

// Save config stores configurations to a given file.
func (c *Config) Save() (err error) {

	writeFile, err := os.OpenFile(c.Filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer writeFile.Close()

	writer := bufio.NewWriter(writeFile)
	defer writer.Flush()

	data, err := toml.Marshal(*c)
	if err != nil {
		return
	}

	_, err = writer.Write(data)
	return

}

// Print shows current configurations as a TOML style.
func (c *Config) Print() (err error) {

	data, err := toml.Marshal(*c)
	if err != nil {
		return
	}
	fmt.Println(string(data))
	return

}

// lookup checks suitable configuration file name.
// If there is some configuration file in a path from current directory
// to root, use the found file. If there is a git repository in the same path,
// use a configuration file set in the same directory of the repository root.
// Otherwise, use a configuration file in the current directory.
func lookup() (res string) {

	res = ConfigureFile
	path, err := filepath.Abs(".")
	if err != nil {
		return
	}

	var repo string
	for ; path != "/"; path = filepath.Join(path, "..") {

		cand := filepath.Join(path, ConfigureFile)
		if _, exist := os.Stat(cand); exist == nil {
			return cand
		}

		if repo == "" {
			if _, exist := os.Stat(filepath.Join(path, DotGit)); exist == nil {
				repo = path
			}
		}

	}

	if repo != "" {
		res = filepath.Join(repo, ConfigureFile)
	}
	return
}
