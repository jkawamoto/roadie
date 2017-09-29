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
	"strings"
	"time"

	pb "gopkg.in/cheggaaa/pb.v1"

	"github.com/gosuri/uitable"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

const (
	// QueueStatusHeaderName is the header name for queue names.
	QueueStatusHeaderName = "QUEUE NAME"
	// QueueStatusHeaderRunning is the header name for the numbers of running tasks.
	QueueStatusHeaderRunning = "RUNNING"
	// QueueStatusHeaderWaiting is the header name for the numbers of waiting tasks.
	QueueStatusHeaderWaiting = "WAITING"
	// QueueStatusHeaderPending is the header name for the numbers of pending tasks.
	QueueStatusHeaderPending = "PENDING"
	// QueueStatusHeaderWorker is the header name for the number of worker instances.
	QueueStatusHeaderWorker = "WORKER"
	// TaskStatusHeaderName is the header name for task names.
	TaskStatusHeaderName = "TASK NAME"
	// TaskStatusHeaderStatus is the header name for task statuses.
	TaskStatusHeaderStatus = "STATUS"
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
		fmt.Printf("expected 2 arguments. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	err = cmdQueueAdd(&optQueueAdd{
		Metadata: m,
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
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

func cmdQueueAdd(opt *optQueueAdd) (err error) {

	s, err := script.NewScriptTemplate(opt.ScriptFile, opt.ScriptArgs)
	if err != nil {
		return
	}

	// Update instance name.
	// If an instance name is not given, use the default name.
	if opt.TaskName != "" {
		s.Name = strings.ToLower(opt.TaskName)
	}

	// Check a specified bucket exists and create it if not.
	service, err := opt.StorageManager()
	if err != nil {
		return
	}
	storage := cloud.NewStorage(service, opt.Stdout)

	// Update source section.
	err = UpdateSourceSection(opt.Metadata, s, &opt.SourceOpt, storage)
	if err != nil {
		return
	}

	// Update result section
	UpdateResultSection(s, opt.OverWriteResultSection, opt.Stdout)

	queueManager, err := opt.QueueManager()
	if err != nil {
		return
	}

	opt.Spinner.Prefix = fmt.Sprintf("Enqueuing task %s to queue %s...", chalk.Bold.TextStyle(s.Name), chalk.Bold.TextStyle(opt.QueueName))
	opt.Spinner.FinalMSG = fmt.Sprintf("Task %s has been added to queue %s", s.Name, opt.QueueName)
	opt.Spinner.Start()
	defer opt.Spinner.Stop()

	err = queueManager.Enqueue(opt.Context, opt.QueueName, s)
	if err != nil {
		opt.Spinner.FinalMSG = fmt.Sprint(chalk.Bold.TextStyle("Cannot add the task:"), s.Name, ":", err.Error())
	}
	return

}

// CmdQueueStatus lists up existing queue information if no arguments given;
// otherwise lists up existing tasks' information in a given queue.
// Each information should have queue name, the number of items in the queue,
// the number of instances working to the queue.
func CmdQueueStatus(c *cli.Context) error {

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	switch c.NArg() {
	case 0:
		return cmdQueueStatus(m)
	case 1:
		return cmdTaskStatus(m, c.Args().First())
	default:
		fmt.Printf("expected at most one argument. (%d given)\n", c.NArg())
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
	table.AddRow(
		QueueStatusHeaderName, QueueStatusHeaderRunning,
		QueueStatusHeaderWaiting, QueueStatusHeaderPending, QueueStatusHeaderWorker)
	err = queue.Queues(m.Context, func(name string, status cloud.QueueStatus) error {
		table.AddRow(name, status.Running, status.Waiting, status.Pending, status.Worker)
		return nil
	})
	if err != nil {
		return
	}

	m.Spinner.Stop()
	fmt.Fprintln(m.Stdout, table.String())
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
	table.AddRow(TaskStatusHeaderName, TaskStatusHeaderStatus)
	err = manager.Tasks(m.Context, queue, func(name, status string) error {
		table.AddRow(name, status)
		return nil
	})
	if err != nil {
		return
	}

	m.Spinner.Stop()
	fmt.Fprintln(m.Stdout, table.String())
	return

}

// CmdQueueLog prints all log from a queue if only queue name is given;
// otherwise prints log of a specific task.
func CmdQueueLog(c *cli.Context) error {

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	switch c.NArg() {
	case 1:
		return cmdQueueLog(m, c.Args().First(), !c.Bool("no-timestamp"))
	case 2:
		return cmdTaskLog(m, c.Args().First(), c.Args().Get(1), !c.Bool("no-timestamp"))
	default:
		fmt.Printf("expected one or two arguments. (%d given)\n", c.NArg())
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
			fmt.Fprintln(m.Stderr, msg)
		} else {
			fmt.Fprintln(m.Stdout, msg)
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
			fmt.Fprintln(m.Stderr, msg)
		} else {
			fmt.Fprintln(m.Stdout, msg)
		}
		return
	})
}

// CmdQueueInstanceList implements `queue instance list` command.
func CmdQueueInstanceList(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf("expected 1 argument. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}
	err = cmdQueueInstanceList(m, c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdQueueInstanceList lists up instances working with a given queue.
func cmdQueueInstanceList(m *Metadata, queue string) (err error) {

	queueManager, err := m.QueueManager()
	if err != nil {
		return
	}
	return queueManager.Workers(m.Context, queue, func(name string) (err error) {
		_, err = fmt.Fprintln(m.Stdout, name)
		return
	})

}

// CmdQueueInstanceAdd implements `queue instance add`.
func CmdQueueInstanceAdd(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf("expected 1 argument. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}
	err = cmdQueueInstanceAdd(m, c.Args().First(), c.Int("instances"))
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdQueueInstanceAdd creates instances working for a given queue.
func cmdQueueInstanceAdd(m *Metadata, queue string, n int) (err error) {

	queueManager, err := m.QueueManager()
	if err != nil {
		return
	}

	fmt.Fprintln(m.Stdout, "Creating instances")
	bar := pb.New(n)
	bar.Output = m.Stdout
	bar.Prefix("Instance")
	bar.Start()
	defer bar.Finish()

	return queueManager.CreateWorkers(m.Context, queue, n, func(name string) error {
		bar.Increment()
		return nil
	})

}

// CmdQueueStop implements `queue stop` command.
func CmdQueueStop(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf("expected 1 argument. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	err = cmdQueueStop(m, c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdQueueStop stops a given queue.
func cmdQueueStop(m *Metadata, queue string) (err error) {

	queueManager, err := m.QueueManager()
	if err != nil {
		return
	}
	return queueManager.Stop(m.Context, queue)

}

// CmdQueueRestart implements `queue restart` command.
func CmdQueueRestart(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		fmt.Printf("expected 1 argument. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}
	err = cmdQueueRestart(m, c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// CmdQueueRestart restarts executing a queue.
func cmdQueueRestart(m *Metadata, queue string) (err error) {

	queueManager, err := m.QueueManager()
	if err != nil {
		return
	}
	return queueManager.Restart(m.Context, queue)

}

// CmdQueueDelete deletes a task in a queue or whole queue.
func CmdQueueDelete(c *cli.Context) (err error) {

	m, err := getMetadata(c)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	switch c.NArg() {
	case 1:
		err = cmdQueueDelete(m, c.Args().First())
	case 2:
		err = cmdTaskDelete(m, c.Args().First(), c.Args().Get(1))
	default:
		fmt.Printf("expected one or two arguments. (%d given)\n", c.NArg())
		return cli.ShowSubcommandHelp(c)
	}
	if err != nil {
		return cli.NewExitError(err, 2)
	}
	return

}

// cmdQueueDelete deletes a given queue.
func cmdQueueDelete(m *Metadata, queue string) (err error) {

	manager, err := m.QueueManager()
	if err != nil {
		return
	}

	m.Spinner.Prefix = fmt.Sprint("Deleting queue", queue)
	m.Spinner.Start()
	defer m.Spinner.Stop()

	return manager.DeleteQueue(m.Context, queue)

}

// cmdTaskDelete deletes a task in a queue.
func cmdTaskDelete(m *Metadata, queue, task string) (err error) {

	manager, err := m.QueueManager()
	if err != nil {
		return
	}

	m.Spinner.Prefix = fmt.Sprint("Deleting task", task, "in queue", queue)
	m.Spinner.Start()
	defer m.Spinner.Stop()

	return manager.DeleteTask(m.Context, queue, task)

}
