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

	"github.com/jkawamoto/roadie/cloud"
)

// Provider is a mock provider for tests.
type Provider struct {
	MockInstanceManager *InstanceManager
	MockQueueManager    *QueueManager
	MockStorageManager  *StorageManager
	MockLogManager      *LogManager
	MockResourceManager *ResourceManager
}

// NewProvider creates a new mock provider.
func NewProvider() *Provider {
	return &Provider{
		MockInstanceManager: NewInstanceManager(),
		MockQueueManager:    NewQueueManager(),
		MockStorageManager:  NewStorageManager(),
		MockLogManager:      NewLogManager(),
		MockResourceManager: NewResourceManager(),
	}
}

// InstanceManager returns an instance manager interface.
func (m *Provider) InstanceManager(ctx context.Context) (cloud.InstanceManager, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return m.MockInstanceManager, nil
	}
}

// QueueManager returns a queue manager interface.
func (m *Provider) QueueManager(ctx context.Context) (cloud.QueueManager, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return m.MockQueueManager, nil
	}
}

// StorageManager returns a storage manager interface.
func (m *Provider) StorageManager(ctx context.Context) (cloud.StorageManager, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return m.MockStorageManager, nil
	}
}

// LogManager returns a log manager interface.
func (m *Provider) LogManager(ctx context.Context) (cloud.LogManager, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return m.MockLogManager, nil
	}
}

// ResourceManager returns a mock resrouce manager.
func (m *Provider) ResourceManager(ctx context.Context) (cloud.ResourceManager, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return m.MockResourceManager, nil
	}
}
