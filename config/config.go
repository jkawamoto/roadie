package config

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	"github.com/naoina/toml"
	"github.com/ttacon/chalk"
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

	log.Printf(chalk.Red.Color("Cannot read configuration file %s: %s\n"), filename, err.Error())
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
