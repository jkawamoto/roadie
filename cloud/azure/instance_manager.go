//
// cloud/azure/instance_manager.go
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

package azure

import (
	"context"
	"log"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
)

// InstanceManager implements cloud.InstanceManager interface to run a script
// on Azure.
type InstanceManager struct {
	service *BatchService
	Config  *AzureConfig
	Logger  *log.Logger
}

// NewInstanceManager creates a new instance manager.
func NewInstanceManager(ctx context.Context, cfg *AzureConfig, logger *log.Logger) (m *InstanceManager, err error) {

	service, err := NewBatchService(ctx, cfg, logger)
	if err != nil {
		return
	}

	m = &InstanceManager{
		service: service,
		Config:  cfg,
		Logger:  logger,
	}
	return

}

// CreateInstance creates an instance which has a given name.
func (m *InstanceManager) CreateInstance(ctx context.Context, task *script.Script) (err error) {

	err = m.service.CreateJob(ctx, task.Name)
	if err != nil {
		return
	}

	err = m.service.CreateTask(ctx, task.Name, task)
	if err != nil {
		m.service.DeleteJob(ctx, task.Name)
	}
	return

}

// DeleteInstance deletes the given named instance.
func (m *InstanceManager) DeleteInstance(ctx context.Context, name string) error {

	return m.service.DeleteJob(ctx, name)

}

// Instances returns a list of running instances
func (m *InstanceManager) Instances(ctx context.Context, handler cloud.InstanceHandler) (err error) {

	jobs, err := m.service.Jobs(ctx)
	if err != nil {
		return
	}
	for name := range jobs {
		err = handler(name, "Running")
		if err != nil {
			return
		}
	}
	return

}

// AvailableRegions returns a list of available regions.
func (m *InstanceManager) AvailableRegions(ctx context.Context) (regions []cloud.Region, err error) {

	m.Logger.Println("Retrieving available regions")
	regions, err = Locations(ctx, &m.Config.Token, m.Config.SubscriptionID)
	if err != nil {
		m.Logger.Println("Cannot retrieve available regions")
	} else {
		m.Logger.Println("Retrieved available regions")
	}
	return

}

// AvailableMachineTypes returns a list of available machine types.
func (m *InstanceManager) AvailableMachineTypes(ctx context.Context) (types []cloud.MachineType, err error) {
	return m.service.AvailableMachineTypes(ctx)
}
