//
// command/log_test.go
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
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jkawamoto/roadie/cloud/mock"
)

func TestCmdLog(t *testing.T) {

	var err error
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&stdout, p)
	m.Stderr = &stderr
	now := time.Now()

	p.MockLogManager.Logs = map[string][]mock.LogEntry{
		"instance1": {
			{Time: now.Add(-30 * time.Minute), Body: "-30"},
			{Time: now, Body: "0"},
			{Time: now.Add(25 * time.Minute), Body: "25", Stderr: true},
		},
		"instance11": {
			{Time: now.Add(-30 * time.Minute), Body: "-30"},
			{Time: now, Body: "0"},
			{Time: now.Add(25 * time.Minute), Body: "25"},
		},
	}

	t.Run("without time stamps", func(t *testing.T) {
		defer stdout.Reset()
		defer stderr.Reset()

		opt := optLog{
			Metadata:     m,
			InstanceName: "instance1",
			Timestamp:    false,
		}
		err = cmdLog(&opt)
		if err != nil {
			t.Fatalf("cmdLog returns an error: %v", err)
		}

		var msgs []string
		// Stdout
		msgs = strings.Split(strings.TrimRight(stdout.String(), "\n"), "\n")
		if len(msgs) != 2 {
			t.Errorf("%v log entries written to stdout, want %v", len(msgs), 2)
		}
		for i, e := range msgs {
			if expect := p.MockLogManager.Logs[opt.InstanceName][i].Body; expect != e {
				t.Errorf("log entry of %v is %q, want %q", i, e, expect)
			}
		}

		// Stderr
		msgs = strings.Split(strings.TrimRight(stderr.String(), "\n"), "\n")
		if len(msgs) != 1 {
			t.Errorf("%v log entries written to stderr, want %v", len(msgs), 1)
		}
		for i, e := range msgs {
			if expect := p.MockLogManager.Logs[opt.InstanceName][i+2].Body; expect != e {
				t.Errorf("log entry of %v is %q, want %q", i, e, expect)
			}
		}

	})

	t.Run("with time stamps", func(t *testing.T) {
		defer stdout.Reset()
		defer stderr.Reset()

		opt := optLog{
			Metadata:     m,
			InstanceName: "instance11",
			Timestamp:    true,
		}
		err = cmdLog(&opt)
		if err != nil {
			t.Fatalf("cmdLog returns an error: %v", err)
		}

		var msgs []string
		// Stdout
		msgs = strings.Split(strings.TrimRight(stdout.String(), "\n"), "\n")
		if len(msgs) != 3 {
			t.Errorf("%v log entries written to stdout, want %v", len(msgs), 3)
		}
		for i, e := range msgs {
			target := p.MockLogManager.Logs[opt.InstanceName][i]
			expect := fmt.Sprintf("%v %v", target.Time.Format(PrintTimeFormat), target.Body)
			if expect != e {
				t.Errorf("log entry of %v is %q, want %q", i, e, expect)
			}
		}

	})

	t.Run("set from option", func(t *testing.T) {
		defer stdout.Reset()
		defer stderr.Reset()

		opt := optLog{
			Metadata:     m,
			InstanceName: "instance11",
			After:        now,
		}
		err = cmdLog(&opt)
		if err != nil {
			t.Fatalf("cmdLog returns an error: %v", err)
		}

		var msgs []string
		// Stdout
		msgs = strings.Split(strings.TrimRight(stdout.String(), "\n"), "\n")
		if len(msgs) != 1 {
			t.Fatalf("%v log entries written to stdout, want %v", len(msgs), 1)
		}
		for i, e := range msgs {
			if expect := p.MockLogManager.Logs[opt.InstanceName][i+2].Body; expect != e {
				t.Errorf("log entry of %v is %q, want %q", i, e, expect)
			}
		}

	})

	t.Run("set follow option", func(t *testing.T) {
		p.MockLogManager.KeepAlive = true
		p.MockLogManager.Break = now
		defer func() {
			stdout.Reset()
			stderr.Reset()
			p.MockLogManager.KeepAlive = false
			p.MockLogManager.Break = time.Time{}
		}()

		var msgs []string
		opt := optLog{
			Metadata:     m,
			InstanceName: "instance11",
			Follow:       true,
			SleepTime:    100 * time.Millisecond,
		}

		err = cmdLog(&opt)
		if err != nil {
			t.Fatalf("cmdLog returns an error: %v", err)
		}
		msgs = strings.Split(strings.TrimRight(stdout.String(), "\n"), "\n")
		if len(msgs) != 3 {
			t.Fatalf("%v log entries written to stdout, want %v", len(msgs), 3)
		}
		for i, e := range msgs {
			if expect := p.MockLogManager.Logs[opt.InstanceName][i].Body; expect != e {
				t.Errorf("log entry of %v is %q, want %q", i, e, expect)
			}
		}

	})

	t.Run("out-of-service", func(t *testing.T) {
		p.MockLogManager.Failure = true
		defer func() {
			stdout.Reset()
			stderr.Reset()
			p.MockLogManager.Failure = false
		}()

		opt := optLog{
			Metadata:     m,
			InstanceName: "instance1",
		}
		err = cmdLog(&opt)
		if err != mock.ErrServiceFailure {
			t.Error("out-of-service log manager doesn't return any errors")
		}

	})

	t.Run("cancel", func(t *testing.T) {
		defer func() {
			stdout.Reset()
			stderr.Reset()
		}()
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		m.Context = ctx
		opt := optLog{
			Metadata:     m,
			InstanceName: "instance1",
		}
		err = cmdLog(&opt)
		if err == nil {
			t.Error("canceled but no errors are returned")
		}

	})

}
