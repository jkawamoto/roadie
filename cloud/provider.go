//
// cloud/provider.go
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

import "context"

// Provider is an interface of cloud service provider.
type Provider interface {
	// InstanceManager returns an instance manager interface.
	InstanceManager(context.Context) (InstanceManager, error)
	// QueueManager returns a queue manager interface.
	QueueManager(context.Context) (QueueManager, error)
	// StorageManager returns a storage manager interface.
	StorageManager(context.Context) (StorageManager, error)
	// LogManager returns a log manager interface.
	LogManager(context.Context) (LogManager, error)
}
