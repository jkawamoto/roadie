//
// resource/task.go
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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package resource

// Task defines a data structure of enqueued script file.
type Task struct {
	// InstanceName to be created.
	InstanceName string `yaml:"instance-name,omitempty"`
	// Image name to be used to create the instance.
	Image string `yaml:"image,omitempty"`
	// The script body.
	Body ScriptBody `yaml:"body,omitempty"`
	// Queue name.
	QueueName string `yaml:"queue-name"`
	// If true, NextQueuedScript will skip this script.
	// In order to stop a queue, this flag will be used.
	Pending bool `yaml:"pending"`
}
