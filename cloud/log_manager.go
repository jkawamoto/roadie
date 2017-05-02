//
// cloud/log_manager.go
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
	"time"
)

// LogHandler defines a hanler function for log entries.
type LogHandler func(timestamp time.Time, line string, stderr bool) error

// LogManager defines a service interface for obtaining log entries.
type LogManager interface {
	// Get instance log.
	Get(ctx context.Context, instanceName string, from time.Time, handler LogHandler) error
	// Delete instance log.
	Delete(ctx context.Context, instanceName string) error
	GetQueueLog(ctx context.Context, queue string, handler LogHandler) error
	GetTaskLog(ctx context.Context, queue, task string, handler LogHandler) error
}
