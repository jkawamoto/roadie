//
// cloud/mock/queue_test.go
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

package mock

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
)

func TestParseIndexedName(t *testing.T) {

	base := "test-name"
	index := 10
	name := indexedName(base, index)

	rBase, rIdx, err := parseIndexedName(name)
	if err != nil {
		t.Errorf("parseIndexedName returns an error: %v", err)
	}
	if rBase != base {
		t.Errorf("base name %q, want %v", rBase, base)
	}
	if rIdx != index {
		t.Errorf("index %v, want %v", rIdx, index)
	}

}

func insertDummyTasks(m *QueueManager, base string, task script.Script) (err error) {
	for i := 0; i != 10; i++ {
		newQueue := indexedName(base, i)
		for j := 0; j != i+1; j++ {
			newTask := task
			newTask.Name = indexedName(newTask.Name, j)
			err = m.Enqueue(context.Background(), newQueue, &newTask)
			if err != nil {
				return
			}
		}
	}
	return
}

func TestEnqueue(t *testing.T) {

	ctx := context.Background()
	m := NewQueueManager()

	testQueue := "test-queue"
	task := script.Script{
		Name: "test-task",
	}

	for i := 0; i != 5; i++ {
		t.Run(fmt.Sprintf("enqueue(n=%d)", i), func(t *testing.T) {
			newTask := task
			newTask.Name = indexedName(task.Name, i)

			err := m.Enqueue(ctx, testQueue, &newTask)
			if err != nil {
				t.Fatalf("Enqueue returns an error: %v", err)
			}
			if len(m.Status[testQueue]) != i+1 {
				t.Fatal("no tasks exists")
			}
			if i == 0 {
				if m.Status[testQueue][i] != StatusRunning {
					t.Errorf("task %v is not %v: %v", i, StatusRunning, m.Status[testQueue][i])
				}
			} else {
				if m.Status[testQueue][i] != StatusWaiting {
					t.Errorf("task %v is not %v: %v", i, StatusWaiting, m.Status[testQueue][i])
				}
			}
			if len(m.Script[testQueue]) != i+1 || m.Script[testQueue][i].Name != newTask.Name {
				t.Errorf("enqueued taks is %+v, want %+v", m.Script[testQueue][i], newTask)
			}

		})
	}

	t.Run("out-of-service", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()

		if m.Enqueue(ctx, testQueue, &task) == nil {
			t.Error("out-of-service manager doesn't return any error")
		}

	})

	t.Run("canceled", func(t *testing.T) {
		ctx2, cancel := context.WithCancel(ctx)
		cancel()
		if m.Enqueue(ctx2, testQueue, &task) == nil {
			t.Error("context is canceled but no errors are returned")
		}
	})
}

func TestTasks(t *testing.T) {

	var err error
	ctx := context.Background()
	m := NewQueueManager()

	testQueue := "test-queue"
	task := script.Script{
		Name: "test-task",
	}

	err = insertDummyTasks(m, testQueue, task)
	if err != nil {
		t.Fatalf("insertDummyTests returns an error: %v", err)
	}

	index := 5
	targetQueue := indexedName(testQueue, index)
	t.Run("enumerate tasks", func(t *testing.T) {
		c := 0
		err = m.Tasks(ctx, targetQueue, func(name, status string) error {
			if expect := indexedName(task.Name, c); name != expect {
				t.Errorf("task name is %q, want %v", name, expect)
			}
			if c == 0 && status != StatusRunning {
				t.Errorf("task %v is %q, want %v", name, status, StatusRunning)
			} else if c != 0 && status != StatusWaiting {
				t.Errorf("task %v is %q, want %v", name, status, StatusWaiting)
			}
			c++
			return nil
		})
		if err != nil {
			t.Errorf("Task returns an error: %v", err)
		}
		if c != index+1 {
			t.Errorf("enumerated tasks are not enough: %v found, want %v", c, index+1)
		}
	})

	t.Run("handler returns an error", func(t *testing.T) {
		expected := fmt.Errorf("some error")
		err = m.Tasks(ctx, targetQueue, func(name, status string) error {
			return expected
		})
		if err != expected {
			t.Errorf("handler returns an error but Tasks doesn't: %v", err)
		}
	})

	t.Run("out-of-service queue manager", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()

		err = m.Tasks(ctx, targetQueue, func(name, status string) error {
			return nil
		})
		if err != ErrServiceFailure {
			t.Errorf("out-of-service queue manager returns no errors: %v", err)
		}
	})

	t.Run("canceled context", func(t *testing.T) {
		ctx2, cancel := context.WithCancel(ctx)
		cancel()

		err = m.Tasks(ctx2, targetQueue, func(name, status string) error {
			return nil
		})
		if err == nil {
			t.Error("context is canceled but no errors are returned")
		}
	})

}

func TestQueues(t *testing.T) {

	var err error
	ctx := context.Background()
	m := NewQueueManager()

	queueBase := "test-queue"
	task := script.Script{
		Name: "test-task",
	}
	err = insertDummyTasks(m, queueBase, task)
	if err != nil {
		t.Fatalf("insertDummyTasks returns an error: %v", err)
	}
	err = m.Stop(ctx, indexedName(queueBase, 5))
	if err != nil {
		t.Fatalf("Stop returnes an error: %v", err)
	}
	err = m.CreateWorkers(ctx, indexedName(queueBase, 9), 5, func(string) error {
		return nil
	})
	if err != nil {
		fmt.Println(m.Worker)
		t.Fatalf("CreateWorkers returns an error: %v", err)
	}

	t.Run("enumerate queues", func(t *testing.T) {
		c := 0
		err = m.Queues(ctx, func(name string, status cloud.QueueStatus) error {

			var base string
			var index int
			base, index, err = parseIndexedName(name)
			if err != nil {
				t.Fatalf("cannot parse the queue name %q: %v", name, err)
			}

			if base != queueBase || index > 9 {
				t.Errorf("not existing queue is returned: %v", name)
			}
			if index != 5 && status.Pending != 0 {
				t.Errorf("%v pending tasks, want %v", status.Pending, 0)
			} else if index == 5 && status.Pending != 5 {
				t.Errorf("%v pending tasks, want %v", status.Pending, 5)
			}
			if status.Running != 1 {
				t.Errorf("%v running tasks, want %v", status.Running, 1)
			}
			if index != 5 && status.Waiting != index {
				t.Errorf("%v waiting tasks, want %v", status.Waiting, index)
			}
			if index != 9 && status.Worker != 1 {
				t.Errorf("%v workers, want %v", status.Worker, 1)
			} else if index == 9 && status.Worker != 6 {
				t.Errorf("%v workers, want %v", status.Worker, 6)
			}
			c++
			return nil
		})
		if err != nil {
			t.Errorf("Queues returns an error: %v", err)
		}
		if c != 10 {
			t.Errorf("%v enumerated task, want %v", c, 10)
		}

	})

	t.Run("handler returns an error", func(t *testing.T) {
		expect := fmt.Errorf("some error")
		err = m.Queues(ctx, func(name string, status cloud.QueueStatus) error {
			return expect
		})
		if err != expect {
			t.Errorf("Queues returns %v, want %v", err, expect)
		}
	})

	t.Run("out-of-service queue manager", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()

		err = m.Queues(ctx, func(name string, status cloud.QueueStatus) error {
			return nil
		})
		if err != ErrServiceFailure {
			t.Errorf("Queues returns %v, want %v", err, ErrServiceFailure)
		}
	})

	t.Run("canceled context", func(t *testing.T) {
		ctx2, cancel := context.WithCancel(ctx)
		cancel()

		err = m.Queues(ctx2, func(name string, status cloud.QueueStatus) error {
			return nil
		})
		if err == nil {
			t.Error("context is canceled but no errors are returned")
		}
	})

}

func TestStop(t *testing.T) {

	var err error
	ctx := context.Background()
	m := NewQueueManager()

	queueBase := "test-queue"
	task := script.Script{
		Name: "test-task",
	}
	err = insertDummyTasks(m, queueBase, task)
	if err != nil {
		t.Fatalf("insertDumyTasks returns an error: %v", err)
	}

	t.Run("stop a queue", func(t *testing.T) {
		stopped := 5
		target := indexedName(queueBase, stopped)
		err = m.Stop(ctx, target)
		if err != nil {
			t.Fatalf("Stop returns an error: %v", err)
		}

		for name, tasks := range m.Status {
			for i, status := range tasks {
				if name == target {
					if i == 0 && status != StatusRunning {
						t.Errorf("task %v in %v is not %v: %v", i, name, StatusRunning, status)
					} else if i != 0 && status != StatusPending {
						t.Errorf("task %v in %v is not %v: %v", i, name, StatusPending, status)
					}
				} else {
					if i == 0 && status != StatusRunning {
						t.Errorf("task %v in %v is not %v: %v", i, name, StatusRunning, status)
					} else if i != 0 && status != StatusWaiting {
						t.Errorf("task %v in %v is not %v: %v", i, name, StatusWaiting, status)
					}
				}
			}
		}

	})

	t.Run("stop unexisting queue", func(t *testing.T) {
		err = m.Stop(ctx, "not existing queue")
		if err == nil {
			t.Error("try to stop unexisting queue but no errors are returned")
		}
	})

	t.Run("out-of-service queue manager", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()
		err = m.Stop(ctx, indexedName(queueBase, 0))
		if err != ErrServiceFailure {
			t.Errorf("out of service queue manager returns %v, want %v", err, ErrServiceFailure)
		}
	})

	t.Run("canceled context", func(t *testing.T) {
		ctx2, cancel := context.WithCancel(ctx)
		cancel()
		err = m.Stop(ctx2, indexedName(queueBase, 0))
		if err == nil {
			t.Error("context is canceled but no errors are returned")
		}
	})

}

func TestRestart(t *testing.T) {

	var err error
	ctx := context.Background()
	m := NewQueueManager()

	queueBase := "test-queue"
	task := script.Script{
		Name: "test-task",
	}
	err = insertDummyTasks(m, queueBase, task)
	if err != nil {
		t.Fatalf("insertDummyTasks returns an error: %v", err)
	}

	t.Run("restart a sopped queue", func(t *testing.T) {
		stopped := 4
		target := indexedName(queueBase, stopped)
		err = m.Stop(ctx, target)
		if err != nil {
			t.Fatalf("Stop returns an error: %v", err)
		}
		for i, status := range m.Status[target] {
			if i == 0 && status != StatusRunning {
				t.Errorf("task %v is not %v: %v", i, StatusRunning, status)
			} else if i != 0 && status != StatusPending {
				t.Errorf("task %v is not %v: %v", i, StatusPending, status)
			}
		}

		err = m.Restart(ctx, target)
		if err != nil {
			t.Fatalf("Restart returns an error: %v", err)
		}
		for i, status := range m.Status[target] {
			if i == 0 && status != StatusRunning {
				t.Errorf("task %v is not %v: %v", i, StatusRunning, status)
			} else if i != 0 && status != StatusWaiting {
				t.Errorf("task %v is not %v: %v", i, StatusWaiting, status)
			}
		}

	})

	t.Run("restart a not stopped queue", func(t *testing.T) {
		target := indexedName(queueBase, 6)
		err = m.Restart(ctx, target)
		if err == nil {
			t.Error("restart a running queue but no errors are returned")
		}
	})

	t.Run("restart a not existing queue", func(t *testing.T) {
		target := indexedName(queueBase, 99)
		err = m.Restart(ctx, target)
		if err == nil {
			t.Error("restart a not existing queue but no errors are returned")
		}
	})

	t.Run("out-of-service queue manager", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()

		target := indexedName(queueBase, 1)
		err = m.Restart(ctx, target)
		if err != ErrServiceFailure {
			t.Errorf("out-of-service queue manager returns %+v, want %v", err, ErrServiceFailure)
		}
	})

	t.Run("canceled context", func(t *testing.T) {
		ctx2, cancel := context.WithCancel(ctx)
		cancel()

		target := indexedName(queueBase, 1)
		err = m.Restart(ctx2, target)
		if err == nil {
			t.Error("context is canceled but no errors are returned")
		}
	})

}

func TestCreateWorkers(t *testing.T) {

	var err error
	ctx := context.Background()
	m := NewQueueManager()

	testQueue := "test-queue"
	task := script.Script{
		Name: "test-task",
	}
	err = m.Enqueue(ctx, testQueue, &task)
	if err != nil {
		t.Fatalf("Enqueue returns an error: %v", err)
	}
	if m.Worker[testQueue] != 1 {
		t.Fatalf("no worker for %v", testQueue)
	}

	t.Run("create workers", func(t *testing.T) {

		err = m.CreateWorkers(ctx, testQueue, 5, func(name string) error {
			if !strings.HasPrefix(name, testQueue) {
				t.Errorf("created worker %q doesn't have a name prefixed by %v", name, testQueue)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("CreateWorkers returns an error: %v", err)
		}
		if res := m.Worker[testQueue]; res != 6 {
			t.Errorf("%v workers for %v, want %v", res, testQueue, 6)
		}

	})

	t.Run("handler returns an error", func(t *testing.T) {
		expect := fmt.Errorf("some error")
		err = m.CreateWorkers(ctx, testQueue, 5, func(name string) error {
			return expect
		})
		if err != expect {
			t.Fatalf("CreateWorkers returns %v, want %v", err, expect)
		}
	})

	t.Run("create workers for not existing queue", func(t *testing.T) {
		err = m.CreateWorkers(ctx, testQueue+"a", 5, func(name string) error {
			return nil
		})
		if err != ErrQueueNotExist {
			t.Errorf("CreateWorkers returns %v, want %v", err, ErrQueueNotExist)
		}
	})

	t.Run("out-of-service queue manager", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()

		err = m.CreateWorkers(ctx, testQueue, 1, func(name string) error {
			return nil
		})
		if err != ErrServiceFailure {
			t.Errorf("CreateWorkers returns %v, want %v", err, ErrServiceFailure)
		}
	})

	t.Run("canceled context", func(t *testing.T) {
		ctx2, cancel := context.WithCancel(ctx)
		cancel()

		err = m.CreateWorkers(ctx2, testQueue, 1, func(name string) error {
			return nil
		})
		if err == nil {
			t.Error("context is canceled but no errors are returned")
		}
	})

}

func TestWorkers(t *testing.T) {

	var err error
	ctx := context.Background()
	m := NewQueueManager()

	testQueue := "test-queue"
	task := script.Script{
		Name: "test-task",
	}
	err = m.Enqueue(ctx, testQueue, &task)
	if err != nil {
		t.Fatalf("Enqueue returns an error: %v", err)
	}
	if m.Worker[testQueue] != 1 {
		t.Fatalf("no worker for %v", testQueue)
	}
	err = m.CreateWorkers(ctx, testQueue, 9, func(string) error {
		return nil
	})
	if err != nil {
		t.Fatalf("CreateWorkers returns an error: %v", err)
	}

	t.Run("enumerate workers", func(t *testing.T) {
		c := 0
		err = m.Workers(ctx, testQueue, func(name string) error {
			if !strings.HasPrefix(name, testQueue) {
				t.Errorf("found worker name doesn't have prefix %v: %v", testQueue, name)
			} else {
				c++
			}
			return nil
		})
		if err != nil {
			t.Errorf("Workers returns an error: %v", err)
		}
		if c != 10 {
			t.Errorf("%v workers are enumerated, want %v", c, 10)
		}
	})

	t.Run("enumerate workers for not existing queue", func(t *testing.T) {
		err = m.Workers(ctx, testQueue+"a", func(name string) error {
			return nil
		})
		if err != ErrQueueNotExist {
			t.Errorf("Workers returns %v, want %v", err, ErrQueueNotExist)
		}
	})

	t.Run("handler returns an error", func(t *testing.T) {
		expect := fmt.Errorf("some error")
		err = m.Workers(ctx, testQueue, func(name string) error {
			return expect
		})
		if err != expect {
			t.Errorf("Workers returns %v, want %v", err, expect)
		}
	})

	t.Run("out-of-service queue manager", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()

		err = m.Workers(ctx, testQueue, func(name string) error {
			return nil
		})
		if err != ErrServiceFailure {
			t.Errorf("Workers returns %v, want %v", err, ErrServiceFailure)
		}
	})

	t.Run("contex canceled", func(t *testing.T) {
		ctx2, cancel := context.WithCancel(ctx)
		cancel()

		err = m.Workers(ctx2, testQueue, func(name string) error {
			return nil
		})
		if err == nil {
			t.Error("context is canceled but no errors are returned")
		}
	})

}

func TestDeleteQueue(t *testing.T) {

	var err error
	ctx := context.Background()
	m := NewQueueManager()

	queueBase := "test-queue"
	task := script.Script{
		Name: "test-task",
	}
	err = insertDummyTasks(m, queueBase, task)
	if err != nil {
		t.Fatalf("insertDummyTasks returns an error: %v", err)
	}

	t.Run("delete a queue", func(t *testing.T) {
		target := indexedName(queueBase, 8)
		err = m.DeleteQueue(ctx, target)
		if err != nil {
			t.Fatalf("DeleteQueue returns an error: %v", err)
		}
		if _, exist := m.Status[target]; exist {
			t.Error("deleted queue still exists")
		}
		if _, exist := m.Script[target]; exist {
			t.Error("delete queue still exists")
		}
	})

	t.Run("delete a not existing queue", func(t *testing.T) {
		err = m.DeleteQueue(ctx, indexedName(queueBase, 20))
		if err != ErrQueueNotExist {
			t.Errorf("deleting a not existing queue but return %v, want %v", err, ErrQueueNotExist)
		}
	})

	t.Run("out-of-service queue manager", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()

		err = m.DeleteQueue(ctx, indexedName(queueBase, 5))
		if err != ErrServiceFailure {
			t.Errorf("out-of-service queue manager returns %v, want %v", err, ErrServiceFailure)
		}
	})

	t.Run("canceled context", func(t *testing.T) {
		ctx2, cancel := context.WithCancel(ctx)
		cancel()

		err = m.DeleteQueue(ctx2, indexedName(queueBase, 5))
		if err == nil {
			t.Error("canceled context but no errors are returned")
		}
	})

}

func TestDeleteTask(t *testing.T) {

	var err error
	ctx := context.Background()
	m := NewQueueManager()

	queueBase := "test-queue"
	task := script.Script{
		Name: "test-task",
	}
	err = insertDummyTasks(m, queueBase, task)
	if err != nil {
		t.Fatalf("insertDummyTasks returns an error: %v", err)
	}

	t.Run("delete a task", func(t *testing.T) {
		qIdx := 8
		tIdx := 3
		targetTask := indexedName(task.Name, tIdx)

		var exist bool
		for _, s := range m.Script[indexedName(queueBase, qIdx)] {
			if s.Name == targetTask {
				exist = true
			}
		}
		if !exist {
			t.Fatal("target task is not insered")
		}

		err = m.DeleteTask(ctx, indexedName(queueBase, qIdx), targetTask)
		if err != nil {
			t.Fatalf("DeleteTasks returns an error: %v", err)
		}

		for _, s := range m.Script[indexedName(queueBase, qIdx)] {
			if s.Name == targetTask {
				t.Error("deleted task still exists")
			}
		}

	})

	t.Run("delete the first task", func(t *testing.T) {
		err = m.DeleteTask(ctx, indexedName(queueBase, 5), indexedName(task.Name, 0))
		if err != nil {
			t.Fatalf("DeleteTasks returns an error: %v", err)
		}
		for _, s := range m.Script[indexedName(queueBase, 5)] {
			if s.Name == indexedName(task.Name, 0) {
				t.Error("deleted task still exists")
			}
		}
	})

	t.Run("delete the last task", func(t *testing.T) {
		err = m.DeleteTask(ctx, indexedName(queueBase, 2), indexedName(task.Name, 2))
		if err != nil {
			t.Fatalf("DeleteTasks returns an error: %v", err)
		}
		for _, s := range m.Script[indexedName(queueBase, 2)] {
			if s.Name == indexedName(task.Name, 2) {
				t.Error("deleted task still exists")
			}
		}
	})

	t.Run("delete a task in a not existing queue", func(t *testing.T) {
		err = m.DeleteTask(ctx, "random queue", "random task")
		if err == nil {
			t.Error("deleting a task in a not existing queue but no errors are returned")
		}
	})

	t.Run("delete a not existing task", func(t *testing.T) {
		err = m.DeleteTask(ctx, indexedName(queueBase, 1), indexedName(task.Name, 10))
		if err != ErrTaskNotExist {
			t.Errorf("deleting a not existing task returns %v, want %v", err, ErrTaskNotExist)
		}
	})

	t.Run("out-of-service queue manager", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()

		err = m.DeleteTask(ctx, indexedName(queueBase, 6), indexedName(task.Name, 2))
		if err != ErrServiceFailure {
			t.Errorf("DeleteTask returns %v, want %v", err, ErrServiceFailure)
		}
	})

	t.Run("canceled context", func(t *testing.T) {
		ctx2, cancel := context.WithCancel(ctx)
		cancel()

		err = m.DeleteTask(ctx2, indexedName(queueBase, 6), indexedName(task.Name, 2))
		if err == nil {
			t.Error("canceled context but no errors are returned")
		}
	})

}
