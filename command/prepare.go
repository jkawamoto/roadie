//
// command/prepare.go
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

package command

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/cloud/gcp"
	"github.com/jkawamoto/roadie/config"
	"github.com/urfave/cli"
)

const (
	// metadataKey is a key for metadata.
	metadataKey = "metadata"
)

// Metadata is a set of data used for any commands.
type Metadata struct {
	// Context for running a command.
	Context context.Context
	// Config for running a command.
	Config *config.Config
	// Provider of a cloud service.
	provider cloud.Provider
	// Logger to output logs.
	Logger *log.Logger
	// Spinner for decorating output standard message; not logging information.
	// If verbose mode is set, the spinner will be disabled.
	Spinner *spinner.Spinner
}

// InstanceManager returns an instance manager interface.
func (m *Metadata) InstanceManager() (cloud.InstanceManager, error) {
	return m.provider.InstanceManager(m.Context)
}

// QueueManager returns a queue manager interface.
func (m *Metadata) QueueManager() (cloud.QueueManager, error) {
	return m.provider.QueueManager(m.Context)
}

// StorageManager returns a storage manager interface.
func (m *Metadata) StorageManager() (cloud.StorageManager, error) {
	return m.provider.StorageManager(m.Context)
}

// LogManager returns a log manager interface.
func (m *Metadata) LogManager() (cloud.LogManager, error) {
	return m.provider.LogManager(m.Context)
}

// ResourceManager returns a resource manager interface.
func (m *Metadata) ResourceManager() (cloud.ResourceManager, error) {
	return m.provider.ResourceManager(m.Context)
}

// getMetadata gets metadata from a cli context.
func getMetadata(c *cli.Context) (meta *Metadata, err error) {

	meta, ok := c.App.Metadata[metadataKey].(*Metadata)
	if !ok {
		err = fmt.Errorf("No metadata is attached")
	} else if meta.Config == nil {
		err = fmt.Errorf("No configuration is given")
	} else if meta.provider == nil {
		err = fmt.Errorf("Cloud service configuration is not correct")
	}
	return

}

// PrepareCommand prepares executing any command; it loads the configuratio file,
// checkes global flags.
func PrepareCommand(c *cli.Context) (err error) {
	meta := new(Metadata)

	// Get a context from main function.
	meta.Context, _ = c.App.Metadata["context"].(context.Context)

	// Prepare a logger and decorator.
	meta.Spinner = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	var logger *log.Logger
	if c.GlobalBool("verbose") {
		logger = log.New(os.Stderr, "", log.LstdFlags)
		// If verbose mode, spinner is disabled, since it may conflict logging information.
		meta.Spinner.Writer = ioutil.Discard
	} else {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)

	}
	meta.Logger = logger

	// Load the configuration file.
	var cfg *config.Config
	if conf := c.GlobalString("config"); conf != "" {
		cfg = &config.Config{
			FileName: conf,
		}
		err = cfg.Load()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot read the given config file:", err.Error())
			cfg = nil
		}

	} else {
		cfg, err = config.NewConfig()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot read any config files:", err.Error())
			cfg = nil
		}

	}
	meta.Config = cfg

	if meta.Config != nil {
		// Prepare a service provider.
		logger.Println("Checking authentication information")
		var provider cloud.Provider
		switch {
		case cfg.GcpConfig.Project != "":
			provider, err = gcp.NewProvider(meta.Context, &cfg.GcpConfig, logger, c.GlobalBool("auth"))
			if err != nil {
				return
			}
			cfg.Save()

		default:
			// return fmt.Errorf("Configuration isn't correct")
		}
		meta.provider = provider

	}

	c.App.Metadata[metadataKey] = meta
	return nil

}
