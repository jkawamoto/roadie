package cloud

import (
	"context"

	"github.com/jkawamoto/roadie/script"
)

// QueueManagerNameHandler is a type of handler function to retrieve names.
type QueueManagerNameHandler func(string) error

// QueueManagerTaskHandler is a type of handler function to retrieve tasks.
type QueueManagerTaskHandler func(*script.Script) error

// QueueManager is a service interface of a queuing task manager.
type QueueManager interface {
	// Enqueue a new task to a given named queue.
	Enqueue(ctx context.Context, queue string, task *script.Script) error
	// Tasks retrieves tasks in a given names queue.
	Tasks(ctx context.Context, queue string, handler QueueManagerTaskHandler) error
	// Queues retrieves existing queue names.
	Queues(ctx context.Context, handler QueueManagerNameHandler) error
	// Stop executing tasks in a given named queue.
	Stop(ctx context.Context, queue string) error
	// Restart executing tasks in a given names queue.
	Restart(ctx context.Context, queue string) error
	// CreateWorkers creates worker instances working for a given named queue.
	CreateWorkers(ctx context.Context, queue string, diskSize int64, n int, handler QueueManagerNameHandler) error
	// Workers retrieves worker instance names for a given queue.
	Workers(ctx context.Context, queue string, handler QueueManagerNameHandler) error
}
