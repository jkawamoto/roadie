package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/naoina/toml"
)

// Config defines a structure of config file.
type Config struct {
	GCP struct {
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
