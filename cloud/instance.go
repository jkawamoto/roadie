//
// cloud/instance.go
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

package cloud

import (
	"context"

	"github.com/jkawamoto/roadie/script"
)

// InstanceHandler is a handler function to retrive instances' status.
type InstanceHandler func(name, status string) error

// InstanceManager is a service interface of an instance manager.
type InstanceManager interface {
	// CreateInstance creates an instance which has a given name.
	CreateInstance(ctx context.Context, script *script.Script) error
	// DeleteInstance deletes the given named instance.
	DeleteInstance(ctx context.Context, name string) error
	// Instances returns a list of running instances
	Instances(ctx context.Context, handler InstanceHandler) error
}

// MetadataItem has Key and Value properties.
type MetadataItem struct {
	Key   string
	Value string
}
