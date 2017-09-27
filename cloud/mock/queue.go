//
// cloud/mock/queue.go
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
	"strconv"
	"strings"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
)

var (
	// ErrQueueNotExist means a specified queue doesn't exist.
	ErrQueueNotExist = fmt.Errorf("specified queue does not exist")
	// ErrQueueIsNotStopped means a specified queue is not stopped.
	ErrQueueIsNotStopped = fmt.Errorf("specified queue is not stopped")
	// ErrTaskNotExist means a specified task doesn't exist.
	ErrTaskNotExist = fmt.Errorf("specifies task does not exist")
)

// indexedName creates a new name which has a base name and an index.
func indexedName(base string, index int) string {
	return fmt.Sprintf("%v-%v", base, index)
}

// parseIndexedName parses a name made by indexedName and returns the base
// name and the index.
func parseIndexedName(name string) (base string, index int, err error) {
	i := strings.LastIndex(name, "-")
	if i == -1 {
		err = fmt.Errorf("given name doesn't have any index: %v", name)
		return
	}
	base = name[:i]
	c, err := strconv.ParseInt(name[i+1:], 10, 0)
	if err != nil {
		return
	}
	index = int(c)
	return
}

// QueueManager is a mock service implementing cloud.QueueManager.
type QueueManager struct {
	Failure bool
	Status  map[string][]string
	Script  map[string][]*script.Script
	Worker  map[string]int
}

// NewQueueManager creates a new mock queue manager.
func NewQueueManager() *QueueManager {
	return &QueueManager{
		Status: make(map[string][]string),
		Script: make(map[string][]*script.Script),
		Worker: make(map[string]int),
	}
}

// Enqueue a new task to a given named queue.
func (m *QueueManager) Enqueue(ctx context.Context, queue string, task *script.Script) error {

	if m.Failure {
		return ErrServiceFailure
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if len(m.Status[queue]) == 0 {
		m.Status[queue] = append(m.Status[queue], StatusRunning)
	} else {
		m.Status[queue] = append(m.Status[queue], StatusWaiting)
	}
	m.Script[queue] = append(m.Script[queue], task)
	if m.Worker[queue] == 0 {
		m.Worker[queue] = 1
	}
	return nil

}

// Tasks retrieves tasks in a given names queue.
func (m *QueueManager) Tasks(ctx context.Context, queue string, handler cloud.QueueManagerTaskHandler) (err error) {

	if m.Failure {
		return ErrServiceFailure
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	for i, status := range m.Status[queue] {
		err = handler(m.Script[queue][i].Name, status)
		if err != nil {
			break
		}
	}
	return

}

// Queues retrieves existing queue names.
func (m *QueueManager) Queues(ctx context.Context, handler cloud.QueueStatusHandler) (err error) {

	if m.Failure {
		return ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	for name, worker := range m.Worker {

		s := cloud.QueueStatus{
			Waiting: 0,
			Pending: 0,
			Running: 0,
			Worker:  worker,
		}

		for _, status := range m.Status[name] {
			switch status {
			case StatusWaiting:
				s.Waiting++
			case StatusPending:
				s.Pending++
			case StatusRunning:
				s.Running++
			}
		}
		err = handler(name, s)
		if err != nil {
			return
		}
	}
	return

}

// Stop executing tasks in a given named queue.
func (m *QueueManager) Stop(ctx context.Context, queue string) error {

	if m.Failure {
		return ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	tasks, exist := m.Status[queue]
	if !exist {
		return ErrQueueNotExist
	}
	for i, task := range tasks {
		if task != StatusRunning {
			tasks[i] = StatusPending
		}
	}
	return nil

}

// Restart executing tasks in a given names queue.
func (m *QueueManager) Restart(ctx context.Context, queue string) error {

	if m.Failure {
		return ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	tasks, exist := m.Status[queue]
	if !exist {
		return ErrQueueNotExist
	}
	var updated bool
	for i, task := range tasks {
		if task == StatusPending {
			tasks[i] = StatusWaiting
			updated = true
		}
	}
	if !updated {
		return ErrQueueIsNotStopped
	}
	return nil

}

// CreateWorkers creates worker instances working for a given named queue.
func (m *QueueManager) CreateWorkers(ctx context.Context, queue string, n int, handler cloud.QueueManagerNameHandler) (err error) {

	if m.Failure {
		return ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if _, exist := m.Worker[queue]; !exist {
		return ErrQueueNotExist
	}
	for i := 0; i != n; i++ {
		err = handler(indexedName(queue, m.Worker[queue]+i))
		if err != nil {
			return
		}
	}
	m.Worker[queue] += n
	return

}

// Workers retrieves worker instance names for a given queue.
func (m *QueueManager) Workers(ctx context.Context, queue string, handler cloud.QueueManagerNameHandler) (err error) {

	if m.Failure {
		return ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if _, exist := m.Worker[queue]; !exist {
		return ErrQueueNotExist
	}
	for i := 0; i != m.Worker[queue]; i++ {
		err = handler(indexedName(queue, i))
		if err != nil {
			return
		}
	}

	return
}

// DeleteQueue deletes a given named queue.
func (m *QueueManager) DeleteQueue(ctx context.Context, queue string) error {

	if m.Failure {
		return ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, exist := m.Script[queue]
	if !exist {
		return ErrQueueNotExist
	}
	delete(m.Script, queue)
	_, exist = m.Status[queue]
	if !exist {
		return ErrQueueNotExist
	}
	delete(m.Status, queue)

	return nil
}

// DeleteTask deletes a given named task in a given named queue.
func (m *QueueManager) DeleteTask(ctx context.Context, queue, task string) error {

	if m.Failure {
		return ErrServiceFailure
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	tasks, exist := m.Script[queue]
	if !exist {
		return ErrQueueNotExist
	}
	var deleted bool
	for i, t := range tasks {
		if t.Name == task {
			statuses := m.Status[queue]
			m.Status[queue] = append(statuses[:i], statuses[i+1:]...)
			m.Script[queue] = append(tasks[:i], tasks[i+1:]...)
			deleted = true
			break
		}
	}
	if !deleted {
		return ErrTaskNotExist
	}
	return nil

}
