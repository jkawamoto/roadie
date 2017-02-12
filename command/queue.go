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
// along with Roadie.  If not, see <http://www.gnu.org/licenses/>.
//

package command

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	pb "gopkg.in/cheggaaa/pb.v1"

	"golang.org/x/sync/errgroup"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/command/util"
	"github.com/jkawamoto/roadie/config"
	"github.com/jkawamoto/roadie/resource"
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

	fmt.Fprintln(os.Stderr, "Creating instances")
	bar := pb.New(instances)
	bar.Output = os.Stderr
	bar.Prefix("Instance")
	bar.Start()
	defer bar.Finish()

	wg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < instances; i++ {

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

		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			func(name, startup string) {
				wg.Go(func() error {
					defer bar.Increment()
					return createInstance(ctx, name, startup, size, os.Stderr)
				})
			}(name, startup)
		}

	}

	return wg.Wait()

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

	fmt.Fprintln(os.Stderr, "Creating instances")
	bar := pb.New(instances)
	bar.Output = os.Stderr
	bar.Prefix("Instance")
	bar.Start()
	defer bar.Finish()

	wg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < instances; i++ {

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

		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			func(name, startup string) {
				wg.Go(func() error {
					defer bar.Increment()
					return createInstance(ctx, name, startup, size, os.Stderr)
				})
			}(name, startup)
		}

	}

	return wg.Wait()

}

// updatePending updates pending attribute of tasks in a given queue.
func updatePending(ctx context.Context, queue string, pending bool) (err error) {

	store := cloud.NewDatastore(ctx)
	return store.UpdateTasks(queue, func(task *resource.Task) (*resource.Task, error) {
		task.Pending = pending
		return task, nil
	})

}
