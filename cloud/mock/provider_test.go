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

func TestStorageManager(t *testing.T) {

	p := NewProvider()
	m, err := p.StorageManager(context.Background())
	if err != nil {
		t.Fatalf("StorageManager returns an error: %v", err)
	}

	if _, ok := m.(*StorageManager); !ok {
		t.Errorf("StorageManager doesn't return a mock manager: %T", m)
	}

}
