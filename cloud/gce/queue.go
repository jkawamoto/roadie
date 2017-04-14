//
// cloud/gce/queue.go
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
	"google.golang.org/api/iterator"
)

const (
	// QueueKind defines kind of entries stored in cloud datastore.
	QueueKind = "roadie-queue"
	// QueueManagerVersion defines the version of queue manager to be used.
	QueueManagerVersion = "0.1.3"
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
	Config *GcpConfig
	Logger *log.Logger
}

// NewQueueService creates an interace for a queue service based on Google
// Cloud Datastore.
func NewQueueService(ctx context.Context, cfg *GcpConfig, logger *log.Logger) (*QueueService, error) {

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
		return
	}
	defer client.Close()

	_, err = client.RunInTransaction(ctx, func(tx *datastore.Transaction) (err error) {
		_, err = tx.Put(key, &Task{
			Name:      task.InstanceName,
			Image:     task.Image,
			QueueName: queue,
			Script:    task,
			Pending:   false,
		})
		return err
	})

	if err != nil {
		s.Logger.Println("Cannot enqueue the task to queue", queue, ":", err.Error())
	} else {
		s.Logger.Println("Enqueued the task to queue", queue)
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

		err = handler(&task)
		if err != nil {
			break
		}

	}

	s.Logger.Println("Stopped retrieving tasks in queue", queue, ":", err.Error())
	return

}

// Queues retrieves existing queue names.
func (s *QueueService) Queues(ctx context.Context, handler cloud.QueueManagerNameHandler) (err error) {

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
		task.Pending = true
		return task
	})

}

// Restart executing tasks in a queue which has a given name.
func (s *QueueService) Restart(ctx context.Context, queue string) error {

	// TODO: Start one instance.
	return s.UpdateTask(ctx, queue, func(task *Task) *Task {
		task.Pending = false
		return task
	})

}

// CreateWorkers creates worker instances working for a given named queue.
func (s *QueueService) CreateWorkers(ctx context.Context, queue string, diskSize int64, n int, handler cloud.QueueManagerNameHandler) error {

	s.Logger.Println("Creating worker instances for queue", queue)
	compute := NewComputeService(s.Config, s.Logger)
	eg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < n; i++ {

		eg.Go(func() error {

			var err error
			name := fmt.Sprintf("%s-%d%d", queue, time.Now().Unix(), i)
			startup, err := WorkerStartup(&WorkerStartupOpt{
				ProjectID:    s.Config.Project,
				InstanceName: name,
				Name:         queue,
				Version:      QueueManagerVersion,
			})
			if err != nil {
				return err
			}

			err = compute.createInstance(ctx, name, startup, diskSize)
			if err == nil {
				err = handler(name)
			}
			return err

		})

	}

	err := eg.Wait()
	if err != nil {
		s.Logger.Println("Failed to create worker instances:", err.Error())
	} else {
		s.Logger.Println("Created worker instances for queue", queue)
	}
	return err

}

// Workers retrieves worker instance names for a given queue.
func (s *QueueService) Workers(ctx context.Context, queue string, handler cloud.QueueManagerNameHandler) error {

	s.Logger.Println("Retrieving workers in queue", queue)
	compute := NewComputeService(s.Config, s.Logger)
	instances, err := compute.Instances(ctx)
	if err != nil {
		return err
	}

	for name := range instances {
		if strings.HasPrefix(name, queue) {
			err = handler(name)
			if err != nil {
				return err
			}
		}
	}

	s.Logger.Println("Finishes retrieving workers in queue", queue)
	return nil

}
