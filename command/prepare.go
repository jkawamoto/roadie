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

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/cloud/gce"
	"github.com/jkawamoto/roadie/config"
	"github.com/urfave/cli"
)

// Metadata is a set of data used for any commands.
type Metadata struct {
	Context  context.Context
	Config   *config.Config
	provider cloud.Provider
	Logger   *log.Logger
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

// getMetadata gets metadata from a cli context.
func getMetadata(c *cli.Context) *Metadata {

	rawCtx, _ := c.App.Metadata["context"]
	ctx, _ := rawCtx.(context.Context)

	rawCfg, _ := c.App.Metadata["config"]
	cfg, _ := rawCfg.(*config.Config)

	rawProvier, _ := c.App.Metadata["provider"]
	provider, _ := rawProvier.(cloud.Provider)

	rawLogger, _ := c.App.Metadata["logger"]
	logger, _ := rawLogger.(*log.Logger)

	return &Metadata{
		Context:  ctx,
		Config:   cfg,
		provider: provider,
		Logger:   logger,
	}

}

// PrepareCommand prepares executing any command; it loads the configuratio file,
// checkes global flags.
func PrepareCommand(c *cli.Context) (err error) {

	// Load the configuration file.
	var cfg *config.Config
	if conf := c.GlobalString("config"); conf != "" {

		cfg = &config.Config{
			FileName: conf,
		}
		err = cfg.Load()
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("Cannot read the given config file: %v", err.Error()), 1)
		}

	} else {

		cfg, err = config.NewConfig()
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("Cannot read the given config file: %v", err.Error()), 1)
		}

	}
	c.App.Metadata["config"] = cfg

	// Prepare a logger.
	var logger *log.Logger
	if c.Bool("verbose") {
		logger = log.New(os.Stderr, "", log.LstdFlags)
	} else {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}
	c.App.Metadata["logger"] = logger

	// Prepare a service provider.
	switch {
	case cfg.GcpConfig.Project != "":
		c.App.Metadata["provider"] = gce.NewProvider(&cfg.GcpConfig, logger)
	default:
		// TODO: Return an error.
		return fmt.Errorf("")
	}

	return

}
