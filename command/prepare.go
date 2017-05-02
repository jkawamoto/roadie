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
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/cloud/azure"
	"github.com/jkawamoto/roadie/cloud/gcp"
	"github.com/jkawamoto/roadie/config"
	colorable "github.com/mattn/go-colorable"
	"github.com/urfave/cli"
)

const (
	// metadataKey is a key for metadata.
	metadataKey = "metadata"
)

var (
	// ErrNoMetadata is an error raised then no metadata are attached to contexts.
	ErrNoMetadata = fmt.Errorf("No metadata are attached")
	// ErrNoConfiguration is an error raised when no configurations are given.
	ErrNoConfiguration = fmt.Errorf("No configurations are given")
	// ErrServiceConfiguration is an error raised when the given service
	// configuration is not correct.
	ErrServiceConfiguration = fmt.Errorf("Cloud service configuration is not correct")
)

// Metadata is a set of data used for any commands.
type Metadata struct {
	// Context for running a command.
	Context context.Context
	// Config for running a command.
	Config *config.Config
	// Stdin in an io.Readre to read user's input.
	Stdin io.Reader
	// Stdout is an io.Writer to output messages to the standard output.
	Stdout io.Writer
	// Stderr is an io.Writer to output messages to the standard error.
	Stderr io.Writer
	// Logger to output logs.
	Logger *log.Logger
	// Spinner for decorating output standard message; not logging information.
	// If verbose mode is set, the spinner will be disabled.
	Spinner *spinner.Spinner
	// Provider of a cloud service.
	provider cloud.Provider
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

// prepareProvider
func (m *Metadata) prepareProvider(forceAuth bool) (err error) {

	m.Logger.Println("Checking authentication information")
	switch {
	case m.provider != nil:
		// This case will be used to run tests.

	case m.Config.GcpConfig.Project != "":
		m.provider, err = gcp.NewProvider(m.Context, &m.Config.GcpConfig, m.Logger, forceAuth)
		if err != nil {
			return
		}
		m.Config.Save()

	case m.Config.AzureConfig.SubscriptionID != "":
		if m.Config.AzureConfig.TenantID == "" {
			return fmt.Errorf("Azure's tenant ID is not given. Check config file %v", m.Config.FileName)
		}
		m.provider, err = azure.NewProvider(m.Context, &m.Config.AzureConfig, m.Logger, forceAuth)
		if err != nil {
			return
		}
		m.Config.Save()

	default:
		// return ErrServiceConfiguration
	}
	return

}

// getMetadata gets metadata from a cli context.
func getMetadata(c *cli.Context) (meta *Metadata, err error) {

	meta, ok := c.App.Metadata[metadataKey].(*Metadata)
	if !ok {
		err = ErrNoMetadata
	} else if meta.Config == nil {
		err = ErrNoConfiguration
	} else if meta.provider == nil {
		err = ErrServiceConfiguration
	}
	return

}

// PrepareCommand prepares executing any command; it loads the configuratio file,
// checkes global flags.
func PrepareCommand(c *cli.Context) (err error) {
	meta := new(Metadata)

	meta.Stdin = os.Stdin
	if c.GlobalBool("no-color") {
		meta.Stdout = colorable.NewNonColorable(os.Stdout)
		meta.Stderr = colorable.NewNonColorable(os.Stderr)
	} else {
		meta.Stdout = colorable.NewColorableStdout()
		meta.Stderr = colorable.NewColorableStderr()
	}

	// Get a context from main function.
	meta.Context, _ = c.App.Metadata["context"].(context.Context)

	// Prepare a logger and decorator.
	meta.Spinner = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	var logger *log.Logger
	if c.GlobalBool("verbose") {
		logger = log.New(meta.Stderr, "", log.LstdFlags)
		// If verbose mode, spinner is disabled, since it may conflict logging information.
		meta.Spinner.Writer = ioutil.Discard
	} else {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
		meta.Spinner.Writer = meta.Stdout
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
		err = meta.prepareProvider(c.GlobalBool("auth"))
		if err != nil {
			return
		}
	}

	c.App.Metadata[metadataKey] = meta
	return nil

}
