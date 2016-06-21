package config

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	"github.com/naoina/toml"
)

// Config defines a structure of config file.
type Config struct {
	Gcp struct {
		Project     string
		MachineType string
		Zone        string
		Bucket      string
	}
}

// LoadConfig loads config from a given file.
func LoadConfig(filename string) *Config {

	if f, err := os.Open(filename); err == nil {
		defer f.Close()

		if buf, err := ioutil.ReadAll(f); err == nil {

			var config Config
			if err := toml.Unmarshal(buf, &config); err == nil {
				return &config
			}

			log.Println(err.Error())

		} else {
			log.Println(err.Error())
		}

	} else {
		log.Println(err.Error())
	}

	log.Println("Cannot read default configuration: .roadie")
	return &Config{}

}

// Save config stores configurations to a given file.
func (c *Config) Save(filename string) (err error) {

	writeFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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
