//
// cloud/queue_manager.go
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

// QueueStatus holds the numbers of waiting, pending, and running tasks in a queue.
// It also has the number of workers working for the queue, too.
type QueueStatus struct {
	// The number of waiting tasks.
	Waiting int
	// The number of pending tasks.
	Pending int
	// The number of running tasks.
	Running int
	// The number of worker instances.
	Worker int
}

// QueueManagerNameHandler is a type of handler function to retrieve names.
type QueueManagerNameHandler func(string) error

// QueueStatusHandler is a type of handler function to retrieve queues' stauts.
type QueueStatusHandler func(name string, status QueueStatus) error

// QueueManagerTaskHandler is a type of handler function to retrieve tasks.
type QueueManagerTaskHandler func(name string, status string) error

// QueueManager is a service interface of a queuing task manager.
type QueueManager interface {
	// Enqueue a new task to a given named queue.
	Enqueue(ctx context.Context, queue string, task *script.Script) error
	// Tasks retrieves tasks in a given names queue.
	Tasks(ctx context.Context, queue string, handler QueueManagerTaskHandler) error
	// Queues retrieves existing queue names.
	Queues(ctx context.Context, handler QueueStatusHandler) error
	// Stop executing tasks in a given named queue.
	Stop(ctx context.Context, queue string) error
	// Restart executing tasks in a given names queue.
	Restart(ctx context.Context, queue string) error
	// CreateWorkers creates worker instances working for a given named queue.
	CreateWorkers(ctx context.Context, queue string, n int, handler QueueManagerNameHandler) error
	// Workers retrieves worker instance names for a given queue.
	Workers(ctx context.Context, queue string, handler QueueManagerNameHandler) error
	// DeleteQueue deletes a given named queue.
	DeleteQueue(ctx context.Context, queue string) error
	// DeleteTask deletes a given named task in a given named queue.
	DeleteTask(ctx context.Context, queue, task string) error
}
