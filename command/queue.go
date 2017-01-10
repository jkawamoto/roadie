//
// command/queue.go
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

package command

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jkawamoto/roadie/command/cloud"
	"github.com/jkawamoto/roadie/command/resource"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/jkawamoto/roadie/config"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
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

	store := cloud.NewDatastore(util.GetContext(c))
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

	name := c.Args().First()
	store := cloud.NewDatastore(util.GetContext(c))
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

	instances, err := runningInstances(util.GetContext(c))
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
	ctx := util.GetContext(c)

	queue := c.Args().First()
	instances := c.Int("instances")
	size := c.Int64("disk-size")

	var wg sync.WaitGroup
	for i := 0; i < instances; i++ {

		fmt.Fprintf(os.Stderr, "Creating an instance (%d/%d)\n", i+1, instances)
		name := fmt.Sprintf("%s-%d", queue, time.Now().Unix())

		startup, err := resource.WorkerStartup(&resource.WorkerStartupOpt{
			ProjectID:    cfg.Project,
			InstanceName: name,
			Name:         queue,
			Version:      QueueManagerVersion,
		})
		if err != nil {
			return err
		}

		wg.Add(1)
		go func(name, startup string) {
			defer wg.Done()
			createInstance(ctx, name, startup, size, os.Stderr)
		}(name, startup)

	}

	wg.Wait()
	return nil
}

// CmdQueueStop stops executing a queue. In order to stop a queue,
// It updates pending property of all tasks to true.
func CmdQueueStop(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}
	return updatePending(util.GetContext(c), c.Args().First(), true)

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
	ctx := util.GetContext(c)
	err = updatePending(ctx, queue, false)
	if err != nil {
		return
	}

	instances := c.Int("instances")
	size := c.Int64("disk-size")
	for i := 0; i < instances; i++ {

		fmt.Fprintf(os.Stderr, "Creating an instance (%d/%d)\n", i+1, instances)
		name := fmt.Sprintf("%s-%d", queue, time.Now().Unix())
		startup, err := resource.WorkerStartup(&resource.WorkerStartupOpt{
			ProjectID:    cfg.Project,
			InstanceName: name,
			Name:         queue,
			Version:      QueueManagerVersion,
		})
		if err != nil {
			return err
		}

		err = createInstance(ctx, name, startup, size, os.Stderr)
		if err != nil {
			return err
		}

	}

	return

}

// updatePending updates pending attribute of tasks in a given queue.
func updatePending(ctx context.Context, queue string, pending bool) (err error) {

	store := cloud.NewDatastore(ctx)
	if err != nil {
		return
	}
	err = store.UpdateTasks(queue, func(task *resource.Task) (*resource.Task, error) {
		task.Pending = pending
		return task, nil
	})
	return

}
