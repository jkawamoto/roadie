//
// cloud/gce/task.go
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

import "github.com/jkawamoto/roadie/script"

// Task defines a data structure of enqueued script file.
type Task struct {
	// Name of this task.
	Name string `yaml:"name,omitempty"`
	// The script body.
	Script *script.Script `yaml:"script,omitempty"`
	// Queue name.
	QueueName string `yaml:"queue-name"`
	// If true, NextQueuedScript will skip this script.
	// In order to stop a queue, this flag will be used.
	Pending bool `yaml:"pending"`
}
