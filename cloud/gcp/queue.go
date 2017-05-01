//
// cloud/gcp/queue.go
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

package gcp

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"

	"cloud.google.com/go/datastore"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
	compute "google.golang.org/api/compute/v1"
	"google.golang.org/api/iterator"
)

const (
	// QueueKind defines kind of entries stored in cloud datastore.
	QueueKind = "roadie-queue"
	// QueueManagerVersion defines the version of queue manager to be used.
	QueueManagerVersion = "0.2.0-dev"
)

// QueueName is a structure to obtaine QueueName attribute from entities
// in cloud datastore.
type QueueName struct {
	// Queue name.
	QueueName string
}

// QueueService implements cloud.QueueManager based on Google Cloud
// Datastore.
type QueueService struct {
	Config *Config
	Logger *log.Logger
}

// NewQueueService creates an interace for a queue service based on Google
// Cloud Datastore.
func NewQueueService(ctx context.Context, cfg *Config, logger *log.Logger) (*QueueService, error) {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	return &QueueService{
		Config: cfg,
		Logger: logger,
	}, nil

}

// Enqueue add a given script to a given named queue.
func (s *QueueService) Enqueue(ctx context.Context, queue string, task *script.Script) (err error) {

	s.Logger.Println("Enqueuing a task to queue", queue)
	id := time.Now().Unix()
	key := datastore.IDKey(QueueKind, id, nil)

	client, err := datastore.NewClient(ctx, s.Config.Project)
	if err != nil {
		s.Logger.Println("Cannot create a client for Google Cloud Datastore:", err.Error())
		return
	}
	defer client.Close()

	// Update URLs of which scheme is `roadie://` to `gs://`.
	ReplaceURLScheme(s.Config, task)
	s.Logger.Println("Script of the enqueuing task is\n", task.String())

	// Enqueue the task.
	_, err = client.RunInTransaction(ctx, func(tx *datastore.Transaction) (err error) {
		_, err = tx.Put(key, &Task{
			Name:      task.InstanceName,
			QueueName: queue,
			Script:    task,
			Status:    TaskStatusWaiting,
		})
		return
	})
	if err != nil {
		s.Logger.Println("Cannot add the task to the queue:", err.Error())
		return
	}

	// If there are no workers, create one worker.
	exist, err := s.workerExists(ctx, queue)
	if err != nil {
		s.Logger.Println("Cannot retrieve running worker instances:", err.Error())
		return
	} else if !exist {
		err = s.CreateWorkers(ctx, queue, 1, func(name string) error {
			s.Logger.Printf("New instance %v has started\n", name)
			return nil
		})
	}

	if err != nil {
		s.Logger.Println("Cannot enqueue the task to queue", queue, ":", err.Error())
	} else {
		s.Logger.Println("Enqueued the task to queue", queue)
	}
	return

}

// Fetch retrieves one task from a queue and returns it; status of the returned
// task is updated to running.
// If there is no task, return nil with nil error.
func (s *QueueService) Fetch(ctx context.Context, queue string) (task *Task, err error) {

	s.Logger.Println("Retrieving a task in queue", queue)
	query := datastore.NewQuery(QueueKind).Filter("QueueName=", queue).Filter("Status=", TaskStatusWaiting).Limit(1)

	client, err := datastore.NewClient(ctx, s.Config.Project)
	if err != nil {
		return
	}
	defer client.Close()

	task = new(Task)
	_, err = client.RunInTransaction(ctx, func(tx *datastore.Transaction) (err error) {

		iter := client.Run(ctx, query)
		key, err := iter.Next(task)
		if err != nil {
			return
		}

		task.Status = TaskStatusRunning
		_, err = tx.Put(key, task)
		return

	})

	if err == iterator.Done {
		s.Logger.Println("No tasks are found")
		return nil, nil
	}
	return

}

// Tasks retrieves tasks in a given names queue.
func (s *QueueService) Tasks(ctx context.Context, queue string, handler cloud.QueueManagerTaskHandler) (err error) {

	s.Logger.Println("Retrieving tasks in queue", queue)
	query := datastore.NewQuery(QueueKind).Filter("QueueName=", queue)

	client, err := datastore.NewClient(ctx, s.Config.Project)
	if err != nil {
		return
	}
	defer client.Close()

	res := client.Run(ctx, query)
	var task script.Script
	for {

		select {
		case <-ctx.Done():
			s.Logger.Println("Retrieving tasks is canceled")
			return ctx.Err()

		default:
		}

		_, err = res.Next(&task)
		if err == iterator.Done {
			s.Logger.Println("Retrieved tasks in queue", queue)
			return nil
		} else if err != nil {
			break
		}

		err = handler(task.InstanceName, "pending")
		if err != nil {
			break
		}

	}

	s.Logger.Println("Stopped retrieving tasks in queue", queue, ":", err.Error())
	return

}

// Queues retrieves existing queue names.
func (s *QueueService) Queues(ctx context.Context, handler cloud.QueueStatusHandler) (err error) {

	s.Logger.Println("Retrieving queue names")
	query := datastore.NewQuery(QueueKind).Project("QueueName").Distinct()

	client, err := datastore.NewClient(ctx, s.Config.Project)
	if err != nil {
		return
	}
	defer client.Close()

	res := client.Run(ctx, query)
	var name QueueName
	for {

		select {
		case <-ctx.Done():
			s.Logger.Println("Retrieving queue names is canceled")
			return ctx.Err()

		default:
		}

		_, err = res.Next(&name)
		if err == iterator.Done {
			s.Logger.Println("Retrieved queue names")
			return nil
		} else if err != nil {
			break
		}

		err = handler(name.QueueName)
		if err != nil {
			break
		}

	}

	s.Logger.Println("Stopped retrieving queue names:", err.Error())
	return

}

// UpdateTask updates tasks in a given named queue with a given modifier.
func (s *QueueService) UpdateTask(ctx context.Context, queue string, modifier func(*Task) *Task) (err error) {

	s.Logger.Println("Updating tasks' status in queue", queue)
	query := datastore.NewQuery(QueueKind).Filter("QueueName=", queue)

	client, err := datastore.NewClient(ctx, s.Config.Project)
	if err != nil {
		return
	}
	defer client.Close()

	_, err = client.RunInTransaction(ctx, func(tx *datastore.Transaction) (err error) {

		var task Task
		var key *datastore.Key
		res := client.Run(ctx, query)
		for {

			select {
			case <-ctx.Done():
				return ctx.Err()

			default:
			}

			key, err = res.Next(&task)
			if err == iterator.Done {
				return nil
			} else if err != nil {
				return err
			}

			_, err = tx.Put(key, modifier(&task))
			if err != nil {
				return err
			}

		}

	})

	if err != nil {
		s.Logger.Println("Failed to update tasks' status in queue", queue, ":", err.Error())
	} else {
		s.Logger.Println("Updated tasks' status in queue", queue)
	}
	return

}

// Stop executing tasks in a queue which has a given name.
func (s *QueueService) Stop(ctx context.Context, queue string) error {

	return s.UpdateTask(ctx, queue, func(task *Task) *Task {
		if task.Status == TaskStatusWaiting {
			task.Status = TaskStatusPending
		}
		return task
	})

}

// Restart executing tasks in a queue which has a given name.
func (s *QueueService) Restart(ctx context.Context, queue string) (err error) {

	s.Logger.Println("Restarting queue", queue)
	err = s.UpdateTask(ctx, queue, func(task *Task) *Task {
		if task.Status == TaskStatusPending {
			task.Status = TaskStatusWaiting
		}
		return task
	})
	if err != nil {
		return
	}

	if exist, err := s.workerExists(ctx, queue); err != nil {
		return err
	} else if !exist {
		err = s.CreateWorkers(ctx, queue, 1, func(name string) error {
			s.Logger.Printf("New instance %v has started\n", name)
			return nil
		})
		if err != nil {
			return err
		}
	}

	s.Logger.Println("Finished restarting queue", queue)
	return

}

// CreateWorkers creates worker instances working for a given named queue.
func (s *QueueService) CreateWorkers(ctx context.Context, queue string, n int, handler cloud.QueueManagerNameHandler) (err error) {

	s.Logger.Println("Creating worker instances for queue", queue)
	cService := NewComputeService(s.Config, s.Logger)

	// Create an ignition config.
	fluentd, err := FluentdUnit(queueLogKey(queue))
	if err != nil {
		return
	}
	qManager, err := QueueManagerUnit(s.Config.Project, QueueManagerVersion, queue)
	if err != nil {
		return
	}
	logcast, err := LogcastUnit("queue.service")
	if err != nil {
		return
	}
	ignition := NewIgnitionConfig().Append(fluentd).Append(qManager).Append(logcast).String()
	s.Logger.Println("Ignition configuration is", ignition)

	eg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < n; i++ {

		name := fmt.Sprintf("%s-%d%d", queue, time.Now().Unix(), i)
		eg.Go(func() (err error) {

			err = cService.createInstance(ctx, name, []*compute.MetadataItems{
				&compute.MetadataItems{
					Key:   "user-data",
					Value: &ignition,
				},
			})
			if err != nil {
				return
			}
			return handler(name)

		})

	}

	err = eg.Wait()
	if err != nil {
		s.Logger.Println("Failed to create worker instances:", err.Error())
	} else {
		s.Logger.Println("Created worker instances for queue", queue)
	}
	return

}

// Workers retrieves worker instance names for a given queue.
func (s *QueueService) Workers(ctx context.Context, queue string, handler cloud.QueueManagerNameHandler) (err error) {

	s.Logger.Println("Retrieving workers in queue", queue)
	cService := NewComputeService(s.Config, s.Logger)
	prefix := fmt.Sprintf("%v-", queue)
	err = cService.Instances(ctx, func(name, status string) error {
		if strings.HasPrefix(name, prefix) && status == StatusRunning {
			s.Logger.Println("Worker", name, "is working for queue", queue)
			return handler(name)
		}
		return nil
	})
	if err != nil {
		return
	}

	s.Logger.Println("Finishes retrieving workers in queue", queue)
	return

}

// DeleteQueue deletes a given named queue. This function deletes all tasks
// in a given queue and deletes all workers for that queue.
func (s *QueueService) DeleteQueue(ctx context.Context, queue string) (err error) {

	s.Logger.Println("Deleting queue", queue)
	query := datastore.NewQuery(QueueKind).Filter("QueueName=", queue)
	err = s.deleteTask(ctx, query)
	if err != nil {
		return
	}

	cService := NewComputeService(s.Config, s.Logger)
	err = s.Workers(ctx, queue, func(name string) (err error) {
		return cService.DeleteInstance(ctx, name)
	})
	if err != nil {
		return
	}

	s.Logger.Println("Finished deleting queue", queue)
	return
}

// DeleteTask deletes a given named task in a given named queue.
func (s *QueueService) DeleteTask(ctx context.Context, queue, task string) (err error) {

	s.Logger.Println("Deleting task", task)
	query := datastore.NewQuery(QueueKind).Filter("QueueName=", queue).Filter("Name=", task)
	err = s.deleteTask(ctx, query)
	if err != nil {
		s.Logger.Println("Failed to delete task", task, ":", err.Error())
		return
	}

	s.Logger.Println("Finished deleting task", task)
	return
}

// workerExists returns true if there is at lease one worker is working for the
// given named queue.
func (s *QueueService) workerExists(ctx context.Context, queue string) (exist bool, err error) {

	done := fmt.Errorf("Iteration done")
	err = s.Workers(ctx, queue, func(name string) error {
		exist = true
		return done
	})
	if err == done {
		err = nil
	}
	return

}

func (s *QueueService) deleteTask(ctx context.Context, query *datastore.Query) (err error) {

	client, err := datastore.NewClient(ctx, s.Config.Project)
	if err != nil {
		return
	}
	defer client.Close()

	_, err = client.RunInTransaction(ctx, func(tx *datastore.Transaction) (err error) {

		var key *datastore.Key
		iter := client.Run(ctx, query)
		for {

			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			key, err = iter.Next(nil)
			if err == iterator.Done {
				return nil
			} else if err != nil {
				return err
			}

			err = tx.Delete(key)
			if err != nil {
				return
			}

		}

	})
	return

}

// queueLogKey returns the log key associated with a given queue.
func queueLogKey(queue string) string {
	return fmt.Sprintf("queue-%v", queue)
}
