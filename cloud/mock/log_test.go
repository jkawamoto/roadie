//
// cloud/mock/log_test.go
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
	"io"
	"testing"
	"time"
)

func TestLogManagerGet(t *testing.T) {

	now := time.Now()

	m := NewLogManager()
	m.Logs = map[string][]LogEntry{
		"instance1": {
			{Time: now.Add(-30 * time.Minute), Body: "-30"},
			{Time: now.Add(1 * time.Minute), Body: "0"},
			{Time: now.Add(25 * time.Minute), Body: "25", Stderr: true},
		},
		"instance11": {
			{Time: now.Add(1 * time.Minute), Body: "0"},
			{Time: now.Add(4 * time.Minute), Body: "4", Stderr: true},
			{Time: now.Add(10 * time.Minute), Body: "10", Stderr: true},
		},
	}

	name := "instance1"
	t.Run("existing instance", func(t *testing.T) {

		var res []LogEntry
		err := m.Get(context.Background(), name, now, func(timestamp time.Time, line string, stderr bool) error {
			res = append(res, LogEntry{
				Time:   timestamp,
				Body:   line,
				Stderr: stderr,
			})
			return nil
		})
		if err != nil {
			t.Fatalf("Get of %q returns an error: %v", name, err)
		}

		expected := m.Logs[name][1:]
		if len(res) != len(expected) {
			t.Fatalf("%v log entries returned, want %v", len(res), len(expected))
		}
		for i, e := range expected {
			if res[i].Time != e.Time || res[i].Body != e.Body || res[i].Stderr != e.Stderr {
				t.Errorf("log entry %v is %v, want %v", i, res[i], e)
			}
		}

	})

	t.Run("keep alive", func(t *testing.T) {
		m.KeepAlive = true
		defer func() { m.KeepAlive = false }()

		err := m.Get(context.Background(), name, now, func(timestamp time.Time, line string, stderr bool) error {
			return nil
		})
		if err != io.EOF {
			t.Fatalf("Get of %q returns an error: %v", name, err)
		}

	})

	t.Run("break point", func(t *testing.T) {
		m.KeepAlive = true
		defer func() { m.KeepAlive = false }()

		breakPoint := now.Add(5 * time.Minute)
		m.Break = breakPoint

		var res []LogEntry
		// First call.
		err := m.Get(context.Background(), "instance11", now, func(timestamp time.Time, line string, stderr bool) error {
			res = append(res, LogEntry{
				Time:   timestamp,
				Body:   line,
				Stderr: stderr,
			})
			return nil
		})
		if err != nil {
			t.Fatalf("Get of %q returns an error: %v", name, err)
		}

		if len(res) != 2 {
			t.Fatalf("%v log entries returned, want %v", len(res), 2)
		}
		for _, e := range res {
			if e.Time.After(breakPoint) {
				t.Errorf("an entry issued after the break point is returned: %v", e)
			}
		}

		// Second call.
		res = []LogEntry{}
		err = m.Get(context.Background(), "instance11", breakPoint, func(timestamp time.Time, line string, stderr bool) error {
			res = append(res, LogEntry{
				Time:   timestamp,
				Body:   line,
				Stderr: stderr,
			})
			return nil
		})
		if err != io.EOF {
			t.Fatalf("Get of %q returns an error: %v", name, err)
		}

		if len(res) != 1 {
			t.Fatalf("%v log entries returned, want %v", len(res), 1)
		}
		for _, e := range res {
			if e.Time.Before(breakPoint) {
				t.Errorf("an entry issued before the break point is returned: %v", e)
			}
		}

	})

	t.Run("not existing instance", func(t *testing.T) {

		var res []LogEntry
		err := m.Get(context.Background(), "dummy_instance", now, func(timestamp time.Time, line string, stderr bool) error {
			res = append(res, LogEntry{
				Time:   timestamp,
				Body:   line,
				Stderr: stderr,
			})
			return nil
		})
		if err != nil {
			t.Fatalf("Get of %q returns an error: %v", name, err)
		}

		if len(res) != 0 {
			t.Errorf("%v log entries are returned, want %v", len(res), 0)
		}

	})

	t.Run("handler returns an error", func(t *testing.T) {

		expected := fmt.Errorf("some error")
		err := m.Get(context.Background(), name, now, func(timestamp time.Time, line string, stderr bool) error {
			return expected
		})
		if err != expected {
			t.Errorf("Get returns %+v, want %v", err, expected)
		}

	})

	t.Run("out-of-service", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()

		err := m.Get(context.Background(), name, now, func(timestamp time.Time, line string, stderr bool) error {
			return nil
		})
		if err != ErrServiceFailure {
			t.Errorf("Get returns %+v, want %v", err, ErrServiceFailure)
		}

	})

	t.Run("canceled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := m.Get(ctx, name, now, func(timestamp time.Time, line string, stderr bool) error {
			return nil
		})
		if err == nil {
			t.Error("context is canceled but no errors are returned")
		}

	})

}

func TestLogMamagerDelete(t *testing.T) {

	now := time.Now()
	m := NewLogManager()
	m.Logs = map[string][]LogEntry{
		"instance1": {
			{Time: now.Add(-30 * time.Minute), Body: "-30"},
			{Time: now, Body: "0"},
			{Time: now.Add(25 * time.Minute), Body: "25", Stderr: true},
		},
		"instance11": {
			{Time: now.Add(-30 * time.Minute), Body: "-30"},
			{Time: now, Body: "0"},
			{Time: now.Add(25 * time.Minute), Body: "25", Stderr: true},
		},
		"instance21": {
			{Time: now.Add(-30 * time.Minute), Body: "-30"},
			{Time: now, Body: "0"},
			{Time: now.Add(25 * time.Minute), Body: "25", Stderr: true},
		},
	}

	name := "instance1"
	t.Run("delete logs for an existing instance", func(t *testing.T) {

		err := m.Delete(context.Background(), "instance1")
		if err != nil {
			t.Fatalf("Delete returns an error: %v", err)
		}
		if _, exist := m.Logs[name]; exist {
			t.Error("deleted log entries still exist")
		}

	})

	t.Run("delete logs for not existing instance", func(t *testing.T) {

		err := m.Delete(context.Background(), "dummy_instance")
		if err == nil {
			t.Error("deleting logs for not existing instance but no errors are returned")
		}

	})

	t.Run("out-of-service", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()

		err := m.Delete(context.Background(), "instance11")
		if err != ErrServiceFailure {
			t.Error("out-of-service manager doesn't return ErrServiceFailure")
		}

	})

	t.Run("canceled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := m.Delete(ctx, "instance21")
		if err == nil {
			t.Error("context is canceled but no errors are returned")
		}

	})

}

func TestGetQueueLog(t *testing.T) {

	m := NewLogManager()
	now := time.Now()
	queue := "test-queue"
	m.QueueLogs = map[string][]LogEntry{
		queue: {
			{Time: now.Add(-30 * time.Minute), Body: "task-1: aaaa"},
			{Time: now.Add(1 * time.Minute), Body: "task-2: bbbb"},
			{Time: now.Add(25 * time.Minute), Body: "task-1: aaaa", Stderr: true},
		},
	}

	t.Run("retrieve logs", func(t *testing.T) {
		var c int
		err := m.GetQueueLog(context.Background(), queue, func(d time.Time, list string, stderr bool) error {
			if expect := m.QueueLogs[queue][c]; expect.Time != d || expect.Body != list || expect.Stderr != stderr {
				t.Errorf("(%v, %v, %v), want %v", d, list, stderr, expect)
			}
			c++
			return nil
		})
		if err != nil {
			t.Fatalf("GetQueueLog returns an error: %v", err)
		}
	})

	t.Run("handler returns an error", func(t *testing.T) {
		expect := fmt.Errorf("some error")
		err := m.GetQueueLog(context.Background(), queue, func(t time.Time, list string, stderr bool) error {
			return expect
		})
		if err != expect {
			t.Fatalf("GetQueueLog returns %v, want %v", err, expect)
		}
	})

	t.Run("out-of-service", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()

		err := m.GetQueueLog(context.Background(), "queue", func(t time.Time, list string, stderr bool) error {
			return nil
		})
		if err != ErrServiceFailure {
			t.Error("out-of-service manager doesn't return ErrServiceFailure")
		}
	})

	t.Run("canceled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := m.GetQueueLog(ctx, "queue", func(t time.Time, list string, stderr bool) error {
			return nil
		})
		if err == nil {
			t.Error("context is canceled but no errors are returned")
		}
	})

}

func TestGetTaskLog(t *testing.T) {

	m := NewLogManager()
	now := time.Now()
	queue := "test-queue"
	m.QueueLogs = map[string][]LogEntry{
		queue: {
			{Time: now.Add(-30 * time.Minute), Body: "task-1: aaaa"},
			{Time: now.Add(1 * time.Minute), Body: "task-2: bbbb"},
			{Time: now.Add(25 * time.Minute), Body: "task-1: aaaa", Stderr: true},
		},
	}

	t.Run("retrieve logs", func(t *testing.T) {

		var c int
		err := m.GetTaskLog(context.Background(), queue, "task-1", func(d time.Time, list string, stderr bool) error {
			if list != "aaaa" {
				t.Errorf("wrong entry: %v", list)
			}
			c++
			return nil
		})
		if err != nil {
			t.Fatalf("GetQueueLog returns an error: %v", err)
		}
		if c != 2 {
			t.Errorf("%v log entries are found, want %v", c, 2)
		}

	})

	t.Run("handler returns an error", func(t *testing.T) {
		expect := fmt.Errorf("some error")
		err := m.GetTaskLog(context.Background(), queue, "task-1", func(t time.Time, list string, stderr bool) error {
			return expect
		})
		if err != expect {
			t.Fatalf("GetTaskLog returns %v, want %v", err, expect)
		}
	})

	t.Run("out-of-service", func(t *testing.T) {
		m.Failure = true
		defer func() { m.Failure = false }()

		err := m.GetTaskLog(context.Background(), "queue", "task", func(t time.Time, list string, stderr bool) error {
			return nil
		})
		if err != ErrServiceFailure {
			t.Error("out-of-service manager doesn't return ErrServiceFailure")
		}
	})

	t.Run("canceled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := m.GetTaskLog(ctx, "queue", "task", func(t time.Time, list string, stderr bool) error {
			return nil
		})
		if err == nil {
			t.Error("context is canceled but no errors are returned")
		}
	})

}
