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
	"os"
	"strings"
	"time"

	pb "gopkg.in/cheggaaa/pb.v1"

	"github.com/gosuri/uitable"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// QueueName is a structure to obtaine QueueName attribute from entities
// in cloud datastore.
type QueueName struct {
	// Queue name.
	QueueName string
}

// optQueueAdd defines arguments for cmdQueueAdd.
type optQueueAdd struct {
	// Metadata to run a command.
	*Metadata
	// SourceOpt specifies options for source secrion of a script.
	SourceOpt
	// TaskName to be created.
	TaskName string
	// QueueName the task to be added to.
	QueueName string
	// ScriptFile to be run.
	ScriptFile string
	// ScriptArgs to fill place holders in the script.
	ScriptArgs []string
	// OverWriteResultSection if it is set True.
	OverWriteResultSection bool
}

// CmdQueueAdd adds a given script to a given named queue.
func CmdQueueAdd(c *cli.Context) (err error) {

	if c.NArg() != 2 {
		fmt.Printf(chalk.Red.Color("expected 2 arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	return cmdQueueAdd(&optQueueAdd{
		Metadata: getMetadata(c),
		SourceOpt: SourceOpt{
			Git:     c.String("git"),
			URL:     c.String("url"),
			Local:   c.String("local"),
			Exclude: c.StringSlice("exclude"),
			Source:  c.String("source"),
		},
		TaskName:               c.String("name"),
		QueueName:              c.Args().First(),
		ScriptFile:             c.Args().Get(1),
		ScriptArgs:             c.StringSlice("e"),
		OverWriteResultSection: c.Bool("overwrite-result-section"),
	})
}

func cmdQueueAdd(opt *optQueueAdd) (err error) {

	s, err := script.NewScript(opt.ScriptFile, opt.ScriptArgs)
	if err != nil {
		return
	}

	// Update instance name.
	// If an instance name is not given, use the default name.
	if opt.TaskName != "" {
		s.InstanceName = strings.ToLower(opt.TaskName)
	}

	// Check a specified bucket exists and create it if not.
	service, err := opt.StorageManager()
	if err != nil {
		return err
	}
	storage := cloud.NewStorage(service, nil)

	// Update source section.
	err = UpdateSourceSection(opt.Metadata, s, &opt.SourceOpt, storage, os.Stdout)
	if err != nil {
		return
	}

	// Update result section
	UpdateResultSection(s, opt.OverWriteResultSection, os.Stdout)

	queueManager, err := opt.QueueManager()
	if err != nil {
		return
	}

	opt.Spinner.Prefix = fmt.Sprintf("Enqueuing task %s to queue %s...", chalk.Bold.TextStyle(s.InstanceName), chalk.Bold.TextStyle(opt.QueueName))
	opt.Spinner.FinalMSG = fmt.Sprintf("\n%s\rInstance created.\n", strings.Repeat(" ", len(opt.Spinner.Prefix)+2))
	opt.Spinner.Start()
	defer opt.Spinner.Stop()

	err = queueManager.Enqueue(opt.Context, opt.QueueName, s)
	if err != nil {
		opt.Spinner.FinalMSG = fmt.Sprintf(chalk.Red.Color("\n%s\rCannot create instance.\n"), strings.Repeat(" ", len(opt.Spinner.Prefix)+2))
	}
	return

}

// CmdQueueStatus lists up existing queue information if no arguments given;
// otherwise lists up existing tasks' information in a given queue.
// Each information should have queue name, the number of items in the queue,
// the number of instances working to the queue.
func CmdQueueStatus(c *cli.Context) error {

	m := getMetadata(c)
	switch c.NArg() {
	case 0:
		return cmdQueueStatus(m)
	case 1:
		return cmdTaskStatus(m, c.Args().First())
	default:
		fmt.Printf(chalk.Red.Color("expected at most one argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

}

// cmdQueueStatus prints status of queues.
func cmdQueueStatus(m *Metadata) (err error) {

	m.Spinner.Prefix = "Loading information"
	m.Spinner.Start()
	defer m.Spinner.Stop()

	queue, err := m.QueueManager()
	if err != nil {
		return
	}

	table := uitable.New()
	table.AddRow("QUEUE NAME")
	err = queue.Queues(m.Context, func(name string) error {
		table.AddRow(name)
		return nil
	})
	if err != nil {
		return
	}

	m.Spinner.Stop()
	fmt.Println(table.String())
	return

}

// cmdTaskStatus prints status of tasks in a given queue.
func cmdTaskStatus(m *Metadata, queue string) (err error) {

	m.Spinner.Prefix = "Loading information"
	m.Spinner.Start()
	defer m.Spinner.Stop()

	manager, err := m.QueueManager()
	if err != nil {
		return
	}

	table := uitable.New()
	table.AddRow("TASK NAME", "STATUS")
	err = manager.Tasks(m.Context, queue, func(name, status string) error {
		table.AddRow(name, status)
		return nil
	})
	if err != nil {
		return
	}

	m.Spinner.Stop()
	fmt.Println(table.String())
	return

}

// CmdQueueLog prints all log from a queue if only queue name is given;
// otherwise prints log of a specific task/
func CmdQueueLog(c *cli.Context) error {

	m := getMetadata(c)
	switch c.NArg() {
	case 1:
		return cmdQueueLog(m, c.Args().First(), !c.Bool("no-timestamp"))
	case 2:
		return cmdTaskLog(m, c.Args().First(), c.Args().Get(1), !c.Bool("no-timestamp"))
	default:
		fmt.Printf(chalk.Red.Color("expected one or two arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

}

// cmdQueueLog prints log from a given queue.
func cmdQueueLog(m *Metadata, queue string, timestamp bool) (err error) {
	log, err := m.LogManager()
	if err != nil {
		return
	}
	return log.GetQueueLog(m.Context, queue, func(t time.Time, line string, stderr bool) (err error) {
		var msg string
		if timestamp {
			msg = fmt.Sprintf("%v %s", t.Format(PrintTimeFormat), line)
		} else {
			msg = line
		}
		if stderr {
			fmt.Fprintln(os.Stderr, msg)
		} else {
			fmt.Println(msg)
		}
		return
	})
}

// cmdTaskLog prints log from a task.
func cmdTaskLog(m *Metadata, queue, task string, timestamp bool) (err error) {
	log, err := m.LogManager()
	if err != nil {
		return
	}
	return log.GetTaskLog(m.Context, queue, task, func(t time.Time, line string, stderr bool) (err error) {
		var msg string
		if timestamp {
			msg = fmt.Sprintf("%v %s", t.Format(PrintTimeFormat), line)
		} else {
			msg = line
		}
		if stderr {
			fmt.Fprintln(os.Stderr, msg)
		} else {
			fmt.Println(msg)
		}
		return
	})
}

// CmdQueueInstanceList lists up instances working with a given queue.
func CmdQueueInstanceList(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m := getMetadata(c)
	queueManager, err := m.QueueManager()
	if err != nil {
		return
	}
	queue := c.Args().First()

	return queueManager.Workers(m.Context, queue, func(name string) (err error) {
		_, err = fmt.Println(name)
		return err
	})

}

// CmdQueueInstanceAdd creates instances working for a given queue.
func CmdQueueInstanceAdd(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf(chalk.Red.Color("expected 1 argument. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m := getMetadata(c)
	queue := c.Args().First()
	instances := c.Int("instances")
	queueManager, err := m.QueueManager()
	if err != nil {
		return
	}

	fmt.Fprintln(os.Stderr, "Creating instances")
	bar := pb.New(instances)
	bar.Output = os.Stderr
	bar.Prefix("Instance")
	bar.Start()
	defer bar.Finish()

	return queueManager.CreateWorkers(m.Context, queue, instances, func(name string) error {
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

	m := getMetadata(c)
	queue := c.Args().First()
	queueManager, err := m.QueueManager()
	if err != nil {
		return
	}

	return queueManager.Stop(m.Context, queue)

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
	m := getMetadata(c)
	queueManager, err := m.QueueManager()
	if err != nil {
		return
	}

	return queueManager.Restart(m.Context, queue)

}

// CmdQueueDelete deletes a task in a queue or whole queue.
func CmdQueueDelete(c *cli.Context) error {

	m := getMetadata(c)
	switch c.NArg() {
	case 1:
		return cmdQueueDelete(m, c.Args().First())
	case 2:
		return cmdTaskDelete(m, c.Args().First(), c.Args().Get(1))
	default:
		fmt.Printf(chalk.Red.Color("expected one or two arguments. (%d given)\n"), c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

}

// cmdQueueDelete deletes a given queue.
func cmdQueueDelete(m *Metadata, queue string) (err error) {

	manager, err := m.QueueManager()
	if err != nil {
		return
	}
	return manager.DeleteQueue(m.Context, queue)

}

// cmdTaskDelete deletes a task in a queue.
func cmdTaskDelete(m *Metadata, queue, task string) (err error) {

	manager, err := m.QueueManager()
	if err != nil {
		return
	}
	return manager.DeleteTask(m.Context, queue, task)

}
