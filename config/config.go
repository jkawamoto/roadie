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

	"github.com/jkawamoto/roadie/chalk"
	"github.com/naoina/toml"
)

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

// LoadConfig loads config from a given file.
func LoadConfig(filename string) *Config {

	var err error
	var f *os.File

	if f, err = os.Open(filename); err == nil {
		defer f.Close()

		var buf []byte
		if buf, err = ioutil.ReadAll(f); err == nil {

			var config Config
			if err = toml.Unmarshal(buf, &config); err == nil {
				config.Filename = filename
				return &config
			}
		}
	}

	fmt.Printf(chalk.Red.Color("Cannot read configuration file %s: %s\n"), filename, err.Error())
	return &Config{
		Filename: filename,
	}

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
