package config

import (
	"io/ioutil"
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
func LoadConfig(filename string) (*Config, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := toml.Unmarshal(buf, &config); err != nil {
		return nil, err
	}

	return &config, nil

}
