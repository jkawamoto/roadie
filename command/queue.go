//
// command/queue.go
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

package command

import (
	"fmt"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/jkawamoto/roadie/command/cloud"
	"github.com/jkawamoto/roadie/command/resource"
	"github.com/jkawamoto/roadie/config"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
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

// CmdQueueList lists up existing queue information.
// Each information should have queue name, the number of items in the queue,
// the number of instances working to the queue.
func CmdQueueList(c *cli.Context) (err error) {

	if c.NArg() != 0 {
		fmt.Printf(chalk.Red.Color("expected no arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	cfg := config.FromCliContext(c)
	ctx, cancel := context.WithCancel(config.NewContext(context.Background(), cfg))
	defer cancel()

	store := cloud.NewDatastore(ctx)
	err = store.QueueNames(func(name string) error {
		fmt.Println(name)
		return nil
	})
	return

}

// CmdQueueShow prints information about a given queue.
// It prints how many items in the queue and how many instance working for the queue.
func CmdQueueShow(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	cfg := config.FromCliContext(c)
	ctx, cancel := context.WithCancel(config.NewContext(context.Background(), cfg))
	defer cancel()

	name := c.Args().First()
	store := cloud.NewDatastore(ctx)
	err = store.FindTasks(name, func(item *resource.Task) error {
		fmt.Println(item.InstanceName)
		return nil
	})

	return
}

// CmdQueueInstanceList lists up instances working with a given queue.
func CmdQueueInstanceList(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	ctx := config.NewContext(context.Background(), config.FromCliContext(c))
	instances, err := runningInstances(ctx)
	if err != nil {
		return
	}

	queue := c.Args().First()
	for name := range instances {

		if strings.HasPrefix(name, queue) {
			fmt.Println(name)
		}

	}

	return nil
}

// CmdQueueInstanceAdd creates instances working for a given queue.
func CmdQueueInstanceAdd(c *cli.Context) error {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	cfg := config.FromCliContext(c)
	ctx, cancel := context.WithCancel(config.NewContext(context.Background(), cfg))
	defer cancel()

	queue := c.Args().First()
	startup, err := resource.WorkerStartup(&resource.WorkerStartupOpt{
		ProjectID: cfg.Project,
		Name:      queue,
		Version:   QueueManagerVersion,
	})
	if err != nil {
		return err
	}

	instances := c.Int("instances")
	size := c.Int64("disk-size")
	for i := 0; i < instances; i++ {

		fmt.Fprintf(os.Stderr, "Creating an instance (%d/%d)\n", i+1, instances)
		name := fmt.Sprintf("%s-%d", queue, time.Now().Unix())
		err = createInstance(ctx, name, startup, size, os.Stderr)
		if err != nil {
			return err
		}

	}

	return nil
}

// CmdQueueStop stops executing a queue. In order to stop a queue,
// It updates pending property of all tasks to true.
func CmdQueueStop(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	ctx, cancel := context.WithCancel(config.NewContext(context.Background(), config.FromCliContext(c)))
	defer cancel()

	return updatePending(ctx, c.Args().First(), true)

}

// CmdQueueRestart restarts executing a queue. In order to restart a queue,
// It updates pending property of all tasks to false.
// Then create instances working for the queue.
func CmdQueueRestart(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	queue := c.Args().First()
	cfg := config.FromCliContext(c)
	ctx, cancel := context.WithCancel(config.NewContext(context.Background(), cfg))
	defer cancel()

	err = updatePending(ctx, queue, false)
	if err != nil {
		return
	}

	startup, err := resource.WorkerStartup(&resource.WorkerStartupOpt{
		ProjectID: cfg.Project,
		Name:      queue,
		Version:   QueueManagerVersion,
	})
	if err != nil {
		return err
	}

	instances := c.Int("instances")
	size := c.Int64("disk-size")
	for i := 0; i < instances; i++ {

		fmt.Fprintf(os.Stderr, "Creating an instance (%d/%d)\n", i+1, instances)
		name := fmt.Sprintf("%s-%d", queue, time.Now().Unix())
		err = createInstance(ctx, name, startup, size, os.Stderr)
		if err != nil {
			return err
		}

	}

	return

}

// updatePending updates pending attribute of tasks in a given queue.
func updatePending(ctx context.Context, queue string, pending bool) (err error) {

	cfg, err := config.FromContext(ctx)
	if err != nil {
		return
	}

	client, err := datastore.NewClient(ctx, cfg.Project)
	if err != nil {
		return
	}
	defer client.Close()

	query := datastore.NewQuery(QueueKind).Filter("QueueName=", queue)
	_, err = client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {

		res := client.Run(ctx, query)
		for {

			select {
			case <-ctx.Done():
				return ctx.Err()

			default:
				var task resource.Task
				key, err := res.Next(&task)
				if err == datastore.Done {
					return nil
				} else if err != nil {
					return err
				}

				task.Pending = pending
				tx.Put(key, &task)

			}

		}

	})

	return
}
