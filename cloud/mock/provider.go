//
// cloud/mock/provider.go
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

package mock

import (
	"context"
	"fmt"

	"github.com/jkawamoto/roadie/cloud"
)

// ErrServiceFailure is an error used in tests.
var ErrServiceFailure = fmt.Errorf("this service is out of order")

// Provider is a mock provider for tests.
type Provider struct {
	MockInstanceManager *InstanceManager
	MockQueueManager    cloud.QueueManager
	MockStorageManager  *StorageManager
	MockLogManager      cloud.LogManager
	MockResourceManager cloud.ResourceManager
}

// NewProvider creates a new mock provider.
func NewProvider() *Provider {
	return &Provider{
		MockInstanceManager: NewInstanceManager(),
		MockStorageManager:  NewStorageManager(),
	}
}

// InstanceManager returns an instance manager interface.
func (m *Provider) InstanceManager(context.Context) (cloud.InstanceManager, error) {
	return m.MockInstanceManager, nil
}

// QueueManager returns a queue manager interface.
func (m *Provider) QueueManager(context.Context) (cloud.QueueManager, error) {
	return m.MockQueueManager, nil
}

// StorageManager returns a storage manager interface.
func (m *Provider) StorageManager(context.Context) (cloud.StorageManager, error) {
	return m.MockStorageManager, nil
}

// LogManager returns a log manager interface.
func (m *Provider) LogManager(context.Context) (cloud.LogManager, error) {
	return m.MockLogManager, nil
}

// ResourceManager returns a mock resrouce manager.
func (m *Provider) ResourceManager(context.Context) (cloud.ResourceManager, error) {
	return m.MockResourceManager, nil
}
