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
	"fmt"
	"io"
	"io/ioutil"
	"os"

	pb "gopkg.in/cheggaaa/pb.v1"

	"github.com/jkawamoto/roadie/cloud/gce"
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

	var log io.Writer
	if c.Bool("verbose") {
		log = os.Stderr
	} else {
		log = ioutil.Discard
	}

	ctx := util.GetContext(c)
	cfg, err := config.FromContext(ctx)
	if err != nil {
		return
	}
	queueManager, err := gce.NewQueueService(ctx, cfg.Project, cfg.Zone, cfg.MachineType, log)
	if err != nil {
		return
	}
	defer queueManager.Close()

	return queueManager.Queues(ctx, func(name string) error {
		_, err := fmt.Println(name)
		return err
	})

}

// CmdQueueShow prints information about a given queue.
// It prints how many items in the queue and how many instance working for the queue.
func CmdQueueShow(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	name := c.Args().First()

	var log io.Writer
	if c.Bool("verbose") {
		log = os.Stderr
	} else {
		log = ioutil.Discard
	}

	ctx := util.GetContext(c)
	cfg, err := config.FromContext(ctx)
	if err != nil {
		return
	}
	queueManager, err := gce.NewQueueService(ctx, cfg.Project, cfg.Zone, cfg.MachineType, log)
	if err != nil {
		return
	}
	defer queueManager.Close()

	return queueManager.Tasks(ctx, name, func(item *resource.ScriptBody) error {
		// TODO: Print task information here.
		return nil
	})

}

// CmdQueueInstanceList lists up instances working with a given queue.
func CmdQueueInstanceList(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}
	queue := c.Args().First()

	var log io.Writer
	if c.Bool("verbose") {
		log = os.Stderr
	} else {
		log = ioutil.Discard
	}

	ctx := util.GetContext(c)
	cfg, err := config.FromContext(ctx)
	if err != nil {
		return
	}
	queueManager, err := gce.NewQueueService(ctx, cfg.Project, cfg.Zone, cfg.MachineType, log)
	if err != nil {
		return
	}
	defer queueManager.Close()

	return queueManager.Workers(ctx, queue, func(name string) error {
		_, err := fmt.Println(name)
		return err
	})

}

// CmdQueueInstanceAdd creates instances working for a given queue.
func CmdQueueInstanceAdd(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	queue := c.Args().First()
	instances := c.Int("instances")
	diskSize := c.Int64("disk-size")

	var log io.Writer
	if c.Bool("verbose") {
		log = os.Stderr
	} else {
		log = ioutil.Discard
	}

	ctx := util.GetContext(c)
	cfg, err := config.FromContext(ctx)
	if err != nil {
		return
	}
	queueManager, err := gce.NewQueueService(ctx, cfg.Project, cfg.Zone, cfg.MachineType, log)
	if err != nil {
		return
	}
	defer queueManager.Close()

	fmt.Fprintln(os.Stderr, "Creating instances")
	bar := pb.New(instances)
	bar.Output = os.Stderr
	bar.Prefix("Instance")
	bar.Start()
	defer bar.Finish()

	return queueManager.CreateWorkers(ctx, queue, diskSize, instances, func(name string) error {
		bar.Increment()
		return nil
	})

}

// CmdQueueStop stops executing a queue. In order to stop a queue,
// It updates pending property of all tasks to true.
func CmdQueueStop(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}
	queue := c.Args().First()

	var log io.Writer
	if c.Bool("verbose") {
		log = os.Stderr
	} else {
		log = ioutil.Discard
	}

	ctx := util.GetContext(c)
	cfg, err := config.FromContext(ctx)
	if err != nil {
		return
	}
	queueManager, err := gce.NewQueueService(ctx, cfg.Project, cfg.Zone, cfg.MachineType, log)
	if err != nil {
		return
	}
	defer queueManager.Close()

	return queueManager.Stop(ctx, queue)

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

	var log io.Writer
	if c.Bool("verbose") {
		log = os.Stderr
	} else {
		log = ioutil.Discard
	}

	ctx := util.GetContext(c)
	cfg, err := config.FromContext(ctx)
	if err != nil {
		return
	}
	queueManager, err := gce.NewQueueService(ctx, cfg.Project, cfg.Zone, cfg.MachineType, log)
	if err != nil {
		return
	}
	defer queueManager.Close()

	err = queueManager.Restart(ctx, queue)
	if err != nil {
		return
	}

	instances := c.Int("instances")
	diskSize := c.Int64("disk-size")

	fmt.Fprintln(os.Stderr, "Creating instances")
	bar := pb.New(instances)
	bar.Output = os.Stderr
	bar.Prefix("Instance")
	bar.Start()
	defer bar.Finish()

	return queueManager.CreateWorkers(ctx, queue, diskSize, instances, func(name string) error {
		bar.Increment()
		return nil
	})

}
