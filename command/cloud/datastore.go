//
// command/cloud/datastore.go
//
// Copyright (c) 2016 Junpei Kawamoto
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

package cloud

import (
	"cloud.google.com/go/datastore"
	"github.com/jkawamoto/roadie/command/resource"
	"github.com/jkawamoto/roadie/config"
	"golang.org/x/net/context"
)

// QueueKind defines kind of entries stored in cloud datastore.
const QueueKind = "roadie-queue"

// QueueName is a structure to obtaine QueueName attribute from entities
// in cloud datastore.
type QueueName struct {
	// Queue name.
	QueueName string
}

type Datastore struct {
	ctx context.Context
}

// NewDatastore creates a new datastore interface,
// It requires a context which has a config object.
func NewDatastore(ctx context.Context) *Datastore {
	return &Datastore{
		ctx: ctx,
	}
}

func (d *Datastore) Insert(id int64, task *resource.Task) (err error) {

	cfg, err := config.FromContext(d.ctx)
	if err != nil {
		return
	}

	client, err := datastore.NewClient(d.ctx, cfg.Project)
	if err != nil {
		return
	}
	defer client.Close()

	key := datastore.NewKey(d.ctx, "roadie-queue", "", id, nil)
	trans, err := client.NewTransaction(d.ctx)
	if err != nil {
		return
	}

	_, err = trans.Put(key, task)
	if err != nil {
		trans.Rollback()
	} else {
		trans.Commit()
	}
	return

}

// QueueNames lists up queue names. Founded names are passed to a given handler function.
// If the handler returns non-nil error, listing up will be stopped.
func (d *Datastore) QueueNames(handler func(string) error) (err error) {

	cfg, err := config.FromContext(d.ctx)
	if err != nil {
		return
	}

	client, err := datastore.NewClient(d.ctx, cfg.Project)
	if err != nil {
		return err
	}
	defer client.Close()

	query := datastore.NewQuery(QueueKind).Project("QueueName").Distinct()
	res := client.Run(d.ctx, query)
	for {

		select {
		case <-d.ctx.Done():
			return d.ctx.Err()

		default:
			var name QueueName
			_, err = res.Next(&name)
			if err == datastore.Done {
				return nil
			} else if err != nil {
				return err
			}

			err = handler(name.QueueName)
			if err != nil {
				return err
			}

		}

	}

}

// FindTasks lists up tasks in a given named queue. Founded tasks will be passed
// to a given handler function. If the hunder function returns non-nil error,
// the listing up will be stopped.
func (d *Datastore) FindTasks(name string, handler func(*resource.Task) error) (err error) {

	cfg, err := config.FromContext(d.ctx)
	if err != nil {
		return
	}

	client, err := datastore.NewClient(d.ctx, cfg.Project)
	if err != nil {
		return
	}
	defer client.Close()

	query := datastore.NewQuery(QueueKind).Filter("QueueName=", name)
	res := client.Run(d.ctx, query)
	for {

		select {
		case <-d.ctx.Done():
			return d.ctx.Err()

		default:
			var item resource.Task
			_, err = res.Next(&item)
			if err == datastore.Done {
				return nil
			} else if err != nil {
				return
			}

			err = handler(&item)
			if err != nil {
				return
			}

		}

	}

}

// UpdateTasks updates tasks in a given named queue. Each task will be passed to
// a given handler. The handler should return modified tasks. If the handler
// returns non-nil error, the update will be stopped.
func (d *Datastore) UpdateTasks(name string, handler func(*resource.Task) (*resource.Task, error)) (err error) {

	cfg, err := config.FromContext(d.ctx)
	if err != nil {
		return
	}

	client, err := datastore.NewClient(d.ctx, cfg.Project)
	if err != nil {
		return
	}
	defer client.Close()

	_, err = client.RunInTransaction(d.ctx, func(tx *datastore.Transaction) error {

		query := datastore.NewQuery(QueueKind).Filter("QueueName=", name)
		res := client.Run(d.ctx, query)
		for {

			select {
			case <-d.ctx.Done():
				return d.ctx.Err()

			default:
				var task resource.Task
				key, err := res.Next(&task)
				if err == datastore.Done {
					return nil
				} else if err != nil {
					return err
				}

				newTask, err := handler(&task)
				if err != nil {
					return err
				}
				tx.Put(key, newTask)

			}

		}

	})
	return

}
