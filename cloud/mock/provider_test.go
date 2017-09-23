//
// cloud/mock/provider_test.go
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
	"testing"
)

func TestInstanceManager(t *testing.T) {

	ctx := context.Background()
	p := NewProvider()
	m, err := p.InstanceManager(ctx)
	if err != nil {
		t.Fatalf("InstanceManager returns an error: %v", err)
	}

	if _, ok := m.(*InstanceManager); !ok {
		t.Errorf("InstanceManager doesn't return a mock manager: %T", m)
	}

	// With a canceled context.
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	_, err = p.InstanceManager(ctx)
	if err == nil {
		t.Error("context is canceled but no errors are returned")
	}

}

func TestStorageManager(t *testing.T) {

	ctx := context.Background()
	p := NewProvider()
	m, err := p.StorageManager(ctx)
	if err != nil {
		t.Fatalf("StorageManager returns an error: %v", err)
	}

	if _, ok := m.(*StorageManager); !ok {
		t.Errorf("StorageManager doesn't return a mock manager: %T", m)
	}

	// With a canceled context.
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	_, err = p.StorageManager(ctx)
	if err == nil {
		t.Error("context is canceled but no errors are returned")
	}

}

func TestResouceManager(t *testing.T) {

	ctx := context.Background()
	p := NewProvider()
	m, err := p.ResourceManager(ctx)
	if err != nil {
		t.Fatalf("ResourceManager returns an error: %v", err)
	}

	if _, ok := m.(*ResourceManager); !ok {
		t.Errorf("ResourceManager doesn't return a mock manager: %T", m)
	}

	// With a canceled context.
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	_, err = p.ResourceManager(ctx)
	if err == nil {
		t.Error("context is canceled but no errors are returned")
	}

}
