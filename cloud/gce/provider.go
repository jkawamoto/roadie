//
// cloud/gce/provider.go
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

package gce

import (
	"context"
	"log"

	"github.com/jkawamoto/roadie/cloud"
)

type Provider struct {
	Config *GcpConfig
	Logger *log.Logger
}

func NewProvider(cfg *GcpConfig, logger *log.Logger) *Provider {

	return &Provider{
		Config: cfg,
		Logger: logger,
	}

}

// InstanceManager returns an instance manager interface.
func (p *Provider) InstanceManager(ctx context.Context) (cloud.InstanceManager, error) {

	return NewComputeService(p.Config, p.Logger), nil

}

// QueueManager returns a queue manager interface.
func (p *Provider) QueueManager(ctx context.Context) (cloud.QueueManager, error) {

	return NewQueueService(ctx, p.Config, p.Logger)

}

// StorageManager returns a storage manager interface.
func (p *Provider) StorageManager(ctx context.Context) (cloud.StorageManager, error) {

	return NewStorageService(ctx, p.Config)

}

// LogManager returns a log manager interface.
func (p *Provider) LogManager(ctx context.Context) (cloud.LogManager, error) {

	return NewLogManager(p.Config), nil

}
