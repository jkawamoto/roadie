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
