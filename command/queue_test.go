//
// command/queue_test.go
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
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	yaml "gopkg.in/yaml.v2"

	"github.com/jkawamoto/roadie/cloud/mock"
	"github.com/jkawamoto/roadie/script"
)

// createDummyScript creates a script file and returns its file path for testing.
func createDummyScript(s *script.Script) (name string, err error) {

	tmp, err := ioutil.TempFile("", "")
	if err != nil {
		return
	}
	data, err := yaml.Marshal(s)
	if err != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return
	}
	tmp.Write(data)
	tmp.Close()

	return tmp.Name(), nil

}

func TestCmdQueueAdd(t *testing.T) {

	var err error

	// Prepare a test script file.
	name, err := createDummyScript(&script.Script{
		Source: "roadie://source/somefile.tar.gz",
		Run: []string{
			"cmd {{args}}",
		},
	})
	if err != nil {
		t.Fatalf("createDummyScript returns an error: %v", err)
	}
	defer os.Remove(name)

	// Prepare mock servers.
	p := mock.NewProvider()
	m := testMetadata(nil, p)

	testQueue := "test-queue"
	testTask := "test-task"
	opt := optQueueAdd{
		Metadata: m,
		SourceOpt: SourceOpt{
			Git: "git@github.com:jkawamoto/roadie.git",
		},
		QueueName:  testQueue,
		TaskName:   testTask,
		ScriptFile: name,
		ScriptArgs: []string{"args=abc"},
	}

	t.Run("task name is given", func(t *testing.T) {

		err = cmdQueueAdd(&opt)
		if err != nil {
			t.Fatalf("cmdQueueAdd returns an error: %v", err)
		}
		scripts := p.MockQueueManager.Script[testQueue]
		if len(scripts) == 0 {
			t.Fatal("script file is not found in the queue")
		}
		s := scripts[0]
		if expect := "https://github.com/jkawamoto/roadie.git"; s.Source != expect {
			t.Errorf("source section is %q, want %v", s.Source, expect)
		}
		if s.Name != testTask {
			t.Errorf("task name is %q, want %v", s.Name, testTask)
		}
		if expect := "roadie://result/" + s.Name; s.Result != expect {
			t.Errorf("result section is %q, want %v", s.Result, expect)
		}
		if len(s.Run) == 0 || s.Run[0] != "cmd abc" {
			t.Errorf("run section is %q, want %v", s.Run, "cmd abc")
		}

		if res := p.MockQueueManager.Worker[testQueue]; res != 1 {
			t.Errorf("%v workers running for %v, want %v", res, testQueue, 1)
		}

	})

	t.Run("task name is not given", func(t *testing.T) {
		opt.TaskName = ""

		err = cmdQueueAdd(&opt)
		if err != nil {
			t.Fatalf("cmdQueueAdd returns an error: %v", err)
		}
		scripts := p.MockQueueManager.Script[testQueue]
		if len(scripts) != 2 {
			t.Fatal("script file is not found in the queue")
		}
		s := scripts[1]
		if expect := "https://github.com/jkawamoto/roadie.git"; s.Source != expect {
			t.Errorf("source section is %q, want %v", s.Source, expect)
		}
		if expect := strings.ToLower(filepath.Base(name)); !strings.HasPrefix(s.Name, expect) {
			t.Errorf("task name doesn't have prefix %v: %q", expect, s.Name)
		}
		if expect := "roadie://result/" + s.Name; s.Result != expect {
			t.Errorf("result section is %q, want %v", s.Result, expect)
		}
		if len(s.Run) == 0 || s.Run[0] != "cmd abc" {
			t.Errorf("run section is %q, want %v", s.Run, "cmd abc")
		}

		if res := p.MockQueueManager.Worker[testQueue]; res != 1 {
			t.Errorf("%v workers running for %v, want %v", res, testQueue, 1)
		}

	})

}

func TestCmdQueueStatus(t *testing.T) {

	name, err := createDummyScript(&script.Script{
		Source: "roadie://source/somefile.tar.gz",
		Run: []string{
			"cmd",
		},
	})
	if err != nil {
		t.Fatalf("createDummyScript returns an error: %v", err)
	}
	defer os.Remove(name)

	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)
	queueName := "test-queue"
	taskName := "test-task"
	opt := optQueueAdd{
		Metadata:   m,
		SourceOpt:  SourceOpt{},
		ScriptFile: name,
	}

	for i := 0; i != 5; i++ {
		for j := 0; j != i+1; j++ {
			opt.QueueName = fmt.Sprintf("%v-%v", queueName, i)
			opt.TaskName = fmt.Sprintf("%v-%v", taskName, j)
			err = cmdQueueAdd(&opt)
			if err != nil {
				t.Fatalf("cmdQueueAdd returns an error: %v", err)
			}
		}
	}

	err = cmdQueueStatus(m)
	if err != nil {
		t.Fatalf("cmdQueueStatus returns an error: %v", err)
	}

	scanner := bufio.NewScanner(&output)
	var header bool
	for scanner.Scan() {
		if !header {
			if !strings.HasPrefix(scanner.Text(), QueueStatusHeaderName) {
				t.Errorf("output doesn't have a header: %v", scanner.Text())
			}
			header = true
		} else if !strings.HasPrefix(scanner.Text(), queueName) {
			t.Errorf("output is wrong: %v", scanner.Text())
		}
	}
	if !header {
		t.Error("output doesn't have a header")
	}

}

func TestCmdTaskStatus(t *testing.T) {

	name, err := createDummyScript(&script.Script{
		Source: "roadie://source/somefile.tar.gz",
	})
	if err != nil {
		t.Fatalf("createDummyScript returns an error: %v", err)
	}
	defer os.Remove(name)

	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)
	testQueue := "test-queue"
	taskName := "test-task"
	opt := optQueueAdd{
		Metadata:   m,
		SourceOpt:  SourceOpt{},
		QueueName:  testQueue,
		ScriptFile: name,
	}

	for i := 0; i != 5; i++ {
		opt.TaskName = fmt.Sprintf("%v-%v", taskName, i)
		err = cmdQueueAdd(&opt)
		if err != nil {
			t.Fatalf("cmdQueueAdd returns an error: %v", err)
		}
	}
	if size := len(p.MockQueueManager.Status[testQueue]); size != 5 {
		t.Fatalf("%v tasks in %v, want %v", size, testQueue, 5)
	}

	err = cmdTaskStatus(m, testQueue)
	if err != nil {
		t.Fatalf("cmdTaskStatus returns an error: %v", err)
	}

	scanner := bufio.NewScanner(&output)
	var header bool
	for scanner.Scan() {
		if !header {
			if !strings.HasPrefix(scanner.Text(), TaskStatusHeaderName) {
				t.Errorf("output doesn't have a header: %v", scanner.Text())
			}
			header = true
		} else {
			kv := strings.Split(scanner.Text(), "\t")
			if len(kv) == 1 {
				t.Fatalf("status is wrong: %v", kv)
			}
			if strings.HasSuffix(kv[0], "0") && kv[1] != mock.StatusRunning {
				t.Errorf("status of %v is %v, want %v", kv[0], kv[1], mock.StatusRunning)
			} else if !strings.HasSuffix(kv[0], "0") && kv[1] != mock.StatusWaiting {
				t.Errorf("status of %v is %v, want %v", kv[0], kv[1], mock.StatusWaiting)
			}
		}
	}
	if !header {
		t.Error("output doesn't have a header")
	}

}

func TestCmdQueueLog(t *testing.T) {

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&stdout, p)
	m.Stderr = &stderr
	testQueue := "test-queue"
	now := time.Now()

	p.MockLogManager.QueueLogs = map[string][]mock.LogEntry{
		testQueue: {
			{Time: now.Add(-30 * time.Minute), Body: "task-1: aaaa"},
			{Time: now.Add(1 * time.Minute), Body: "task-2: bbbb"},
			{Time: now.Add(25 * time.Minute), Body: "task-1: aaaa", Stderr: true},
		},
	}

	t.Run("with time stamps", func(t *testing.T) {
		defer stdout.Reset()
		defer stderr.Reset()

		err := cmdQueueLog(m, testQueue, true)
		if err != nil {
			t.Fatalf("cmdQueueLog returns an error: %v", err)
		}

		var c int
		// Stdout
		scanner := bufio.NewScanner(&stdout)
		for c = 0; scanner.Scan(); c++ {
			sp := strings.Split(scanner.Text(), " ")
			if len(sp) < 3 {
				t.Fatalf("wrong log entry: %v", scanner.Text())
			}
			expect := p.MockLogManager.QueueLogs[testQueue][c]
			if expect.Time.Format(PrintTimeFormat) != strings.Join(sp[:2], " ") {
				t.Errorf("time %v, want %v", strings.Join(sp[:2], " "), expect.Time.Format(PrintTimeFormat))
			}
			if res := strings.Join(sp[2:], " "); res != expect.Body {
				t.Errorf("%v, want %v", res, expect.Body)
			}
		}
		if c != 2 {
			t.Errorf("%v log entries are outputted, want %v", c, 3)
		}

		// Stderr
		scanner = bufio.NewScanner(&stderr)
		for c = 0; scanner.Scan(); c++ {
			sp := strings.Split(scanner.Text(), " ")
			if len(sp) < 3 {
				t.Fatalf("wrong log entry: %v", scanner.Text())
			}
			expect := p.MockLogManager.QueueLogs[testQueue][c+2]
			if expect.Time.Format(PrintTimeFormat) != strings.Join(sp[:2], " ") {
				t.Errorf("time %v, want %v", strings.Join(sp[:2], " "), expect.Time.Format(PrintTimeFormat))
			}
			if res := strings.Join(sp[2:], " "); res != expect.Body {
				t.Errorf("%v, want %v", res, expect.Body)
			}
		}
		if c != 1 {
			t.Errorf("%v log entries are outputted to stderr, want %v", c, 1)
		}

	})

	t.Run("without time stamps", func(t *testing.T) {
		defer stdout.Reset()

		err := cmdQueueLog(m, testQueue, false)
		if err != nil {
			t.Fatalf("cmdQueueLog returns an error: %v", err)
		}

		var c int
		// Stdout
		scanner := bufio.NewScanner(&stdout)
		for c = 0; scanner.Scan(); c++ {
			expect := p.MockLogManager.QueueLogs[testQueue][c]
			if expect.Body != scanner.Text() {
				t.Errorf("%v, want %v", scanner.Text(), expect.Body)
			}
		}
		if c != 2 {
			t.Errorf("%v log entries are outputted, want %v", c, 2)
		}

		// Stderr
		scanner = bufio.NewScanner(&stderr)
		for c = 0; scanner.Scan(); c++ {
			expect := p.MockLogManager.QueueLogs[testQueue][c+2]
			if expect.Body != scanner.Text() {
				t.Errorf("%v, want %v", scanner.Text(), expect.Body)
			}
		}
		if c != 1 {
			t.Errorf("%v log entries are outputted to stderr, want %v", c, 1)
		}

	})

}

func TestCmdTaskLog(t *testing.T) {

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&stdout, p)
	m.Stderr = &stderr
	testQueue := "test-queue"
	now := time.Now()

	p.MockLogManager.QueueLogs = map[string][]mock.LogEntry{
		testQueue: {
			{Time: now.Add(-30 * time.Minute), Body: "task-1: aaaa"},
			{Time: now.Add(1 * time.Minute), Body: "task-2: bbbb"},
			{Time: now.Add(25 * time.Minute), Body: "task-1: aaaa", Stderr: true},
		},
	}

	t.Run("with time stamps", func(t *testing.T) {
		defer stdout.Reset()
		defer stderr.Reset()

		err := cmdTaskLog(m, testQueue, "task-1", true)
		if err != nil {
			t.Fatalf("cmdTaskLog returns an error: %v", err)
		}

		var c int
		// Stdout
		scanner := bufio.NewScanner(&stdout)
		for c = 0; scanner.Scan(); c++ {
			sp := strings.Split(scanner.Text(), " ")
			if len(sp) < 2 {
				t.Fatalf("wrong log entry: %v", scanner.Text())
			}
			expect := p.MockLogManager.QueueLogs[testQueue][c]
			if expect.Time.Format(PrintTimeFormat) != strings.Join(sp[:2], " ") {
				t.Errorf("time %v, want %v", strings.Join(sp[:2], " "), expect.Time.Format(PrintTimeFormat))
			}
			if sp[2] != "aaaa" {
				t.Errorf("%v, want %v", sp[2], "aaaa")
			}
		}
		if c != 1 {
			t.Errorf("%v log entries are outputted, want %v", c, 1)
		}

		// Stderr
		scanner = bufio.NewScanner(&stderr)
		for c = 0; scanner.Scan(); c++ {
			sp := strings.Split(scanner.Text(), " ")
			if len(sp) < 2 {
				t.Fatalf("wrong log entry: %v", scanner.Text())
			}
			expect := p.MockLogManager.QueueLogs[testQueue][c+2]
			if expect.Time.Format(PrintTimeFormat) != strings.Join(sp[:2], " ") {
				t.Errorf("time %v, want %v", strings.Join(sp[:2], " "), expect.Time.Format(PrintTimeFormat))
			}
			if sp[2] != "aaaa" {
				t.Errorf("%v, want %v", sp[2], "aaaa")
			}
		}
		if c != 1 {
			t.Errorf("%v log entries are outputted to stderr, want %v", c, 1)
		}

	})

	t.Run("without time stamps", func(t *testing.T) {
		defer stdout.Reset()
		defer stderr.Reset()

		err := cmdTaskLog(m, testQueue, "task-1", false)
		if err != nil {
			t.Fatalf("cmdTaskLog returns an error: %v", err)
		}

		var c int
		// Stdout
		scanner := bufio.NewScanner(&stdout)
		for c = 0; scanner.Scan(); c++ {
			if scanner.Text() != "aaaa" {
				t.Errorf("%v, want %v", scanner.Text(), "aaaa")
			}
		}
		if c != 1 {
			t.Errorf("%v log entries are outputted, want %v", c, 1)
		}

		// Stderr
		scanner = bufio.NewScanner(&stderr)
		for c = 0; scanner.Scan(); c++ {
			if scanner.Text() != "aaaa" {
				t.Errorf("%v, want %v", scanner.Text(), "aaaa")
			}
		}
		if c != 1 {
			t.Errorf("%v log entries are outputted to stderr, want %v", c, 1)
		}

	})

}

func TestCmdQueueInstanceList(t *testing.T) {

	name, err := createDummyScript(&script.Script{
		Source: "roadie://source/somefile.tar.gz",
		Run: []string{
			"cmd",
		},
	})
	if err != nil {
		t.Fatalf("createDummyScript returns an error: %v", err)
	}
	defer os.Remove(name)

	var output bytes.Buffer
	p := mock.NewProvider()
	m := testMetadata(&output, p)
	testQueue := "test-queue"
	taskName := "test-task"
	opt := optQueueAdd{
		Metadata:   m,
		SourceOpt:  SourceOpt{},
		QueueName:  testQueue,
		TaskName:   taskName,
		ScriptFile: name,
	}
	err = cmdQueueAdd(&opt)
	if err != nil {
		t.Fatalf("cmdQueueAdd returns an error: %v", err)
	}
	if res := p.MockQueueManager.Worker[opt.QueueName]; res != 1 {
		t.Fatalf("%v running instances, want %v", res, 1)
	}

	err = cmdQueueInstanceAdd(m, testQueue, 3)
	if err != nil {
		t.Fatalf("cmdQueueInstanceAdd returns an error: %v", err)
	}

	output.Reset()
	err = cmdQueueInstanceList(m, testQueue)
	if err != nil {
		t.Fatalf("cmdQueueInstanceList returns an error: %v", err)
	}

	var c int
	scanner := bufio.NewScanner(&output)
	for c = 0; scanner.Scan(); c++ {
		if expect := fmt.Sprintf("%v-%v", testQueue, c); scanner.Text() != expect {
			t.Errorf("instance name %q, want %v", scanner.Text(), expect)
		}
	}
	if c != 4 {
		t.Errorf("%v instances listed up, want %v", c, 4)
	}

}

func TestCmdQueueInstanceAdd(t *testing.T) {

	name, err := createDummyScript(&script.Script{
		Source: "roadie://source/somefile.tar.gz",
		Run: []string{
			"cmd",
		},
	})
	if err != nil {
		t.Fatalf("createDummyScript returns an error: %v", err)
	}
	defer os.Remove(name)

	p := mock.NewProvider()
	m := testMetadata(nil, p)
	testQueue := "test-queue"
	taskName := "test-task"
	opt := optQueueAdd{
		Metadata:   m,
		SourceOpt:  SourceOpt{},
		QueueName:  testQueue,
		TaskName:   taskName,
		ScriptFile: name,
	}
	err = cmdQueueAdd(&opt)
	if err != nil {
		t.Fatalf("cmdQueueAdd returns an error: %v", err)
	}
	if res := p.MockQueueManager.Worker[testQueue]; res != 1 {
		t.Fatalf("%v running instances, want %v", res, 1)
	}

	err = cmdQueueInstanceAdd(m, testQueue, 5)
	if err != nil {
		t.Fatalf("cmdQueueInstanceAdd returns an error: %v", err)
	}
	if res := p.MockQueueManager.Worker[testQueue]; res != 6 {
		t.Errorf("%v running instances, want %v", res, 6)
	}

}

func TestCmdQueueStop(t *testing.T) {

	name, err := createDummyScript(&script.Script{
		Source: "roadie://source/somefile.tar.gz",
		Run: []string{
			"cmd",
		},
	})
	if err != nil {
		t.Fatalf("createDummyScript returns an error: %v", err)
	}
	defer os.Remove(name)

	p := mock.NewProvider()
	m := testMetadata(nil, p)
	testQueue := "test-queue"
	taskName := "test-task"
	opt := optQueueAdd{
		Metadata:   m,
		SourceOpt:  SourceOpt{},
		QueueName:  testQueue,
		ScriptFile: name,
	}

	for i := 0; i != 5; i++ {
		opt.TaskName = fmt.Sprintf("%v-%v", taskName, i)
		err = cmdQueueAdd(&opt)
		if err != nil {
			t.Fatalf("cmdQueueAdd returns an error: %v", err)
		}
	}
	if size := len(p.MockQueueManager.Status[testQueue]); size != 5 {
		t.Fatalf("%v tasks in %v, want %v", size, testQueue, 5)
	}

	err = cmdQueueStop(m, testQueue)
	if err != nil {
		t.Fatalf("cmdQueueStop returns an error: %v", err)
	}
	for i, status := range p.MockQueueManager.Status[testQueue] {
		if i == 0 && status != mock.StatusRunning {
			t.Errorf("task %v is %v, want %v", i, status, mock.StatusRunning)
		} else if i != 0 && status != mock.StatusPending {
			t.Errorf("task %v is %v, want %v", i, status, mock.StatusPending)
		}
	}

}

func TestCmdQueueRestart(t *testing.T) {

	name, err := createDummyScript(&script.Script{
		Source: "roadie://source/somefile.tar.gz",
		Run: []string{
			"cmd",
		},
	})
	if err != nil {
		t.Fatalf("createDummyScript returns an error: %v", err)
	}
	defer os.Remove(name)

	p := mock.NewProvider()
	m := testMetadata(nil, p)
	testQueue := "test-queue"
	taskName := "test-task"
	opt := optQueueAdd{
		Metadata:   m,
		SourceOpt:  SourceOpt{},
		QueueName:  testQueue,
		ScriptFile: name,
	}

	for i := 0; i != 5; i++ {
		opt.TaskName = fmt.Sprintf("%v-%v", taskName, i)
		err = cmdQueueAdd(&opt)
		if err != nil {
			t.Fatalf("cmdQueueAdd returns an error: %v", err)
		}
	}
	err = cmdQueueStop(m, testQueue)
	if err != nil {
		t.Fatalf("cmdQueueStop returns an error: %v", err)
	}

	err = cmdQueueRestart(m, testQueue)
	if err != nil {
		t.Fatalf("cmdQueueRestart returns an error: %v", err)
	}
	for i, status := range p.MockQueueManager.Status[testQueue] {
		if i == 0 && status != mock.StatusRunning {
			t.Errorf("task %v is %v, want %v", i, status, mock.StatusRunning)
		} else if i != 0 && status != mock.StatusWaiting {
			t.Errorf("task %v is %v, want %v", i, status, mock.StatusWaiting)
		}
	}

}

func TestCmdQueueDelete(t *testing.T) {

	name, err := createDummyScript(&script.Script{
		Source: "roadie://source/somefile.tar.gz",
		Run: []string{
			"cmd",
		},
	})
	if err != nil {
		t.Fatalf("createDummyScript returns an error: %v", err)
	}
	defer os.Remove(name)

	p := mock.NewProvider()
	m := testMetadata(nil, p)
	testQueue := "test-queue"
	taskName := "test-task"
	opt := optQueueAdd{
		Metadata:   m,
		SourceOpt:  SourceOpt{},
		QueueName:  testQueue,
		ScriptFile: name,
	}

	for i := 0; i != 5; i++ {
		opt.TaskName = fmt.Sprintf("%v-%v", taskName, i)
		err = cmdQueueAdd(&opt)
		if err != nil {
			t.Fatalf("cmdQueueAdd returns an error: %v", err)
		}
	}

	err = cmdQueueDelete(m, testQueue)
	if err != nil {
		t.Fatalf("cmdQueueDelete returns an error: %v", err)
	}
	_, exist := p.MockQueueManager.Status[testQueue]
	if exist {
		t.Error("deleted queue still exists")
	}

}

func TestCmdTaskDelete(t *testing.T) {

	name, err := createDummyScript(&script.Script{
		Source: "roadie://source/somefile.tar.gz",
		Run: []string{
			"cmd",
		},
	})
	if err != nil {
		t.Fatalf("createDummyScript returns an error: %v", err)
	}
	defer os.Remove(name)

	p := mock.NewProvider()
	m := testMetadata(nil, p)
	testQueue := "test-queue"
	taskName := "test-task"
	opt := optQueueAdd{
		Metadata:   m,
		SourceOpt:  SourceOpt{},
		QueueName:  testQueue,
		ScriptFile: name,
	}

	for i := 0; i != 5; i++ {
		opt.TaskName = fmt.Sprintf("%v-%v", taskName, i)
		err = cmdQueueAdd(&opt)
		if err != nil {
			t.Fatalf("cmdQueueAdd returns an error: %v", err)
		}
	}

	target := fmt.Sprintf("%v-%v", taskName, 3)
	err = cmdTaskDelete(m, testQueue, target)
	if err != nil {
		t.Fatalf("cmdTaskDelete returns an error: %v", err)
	}
	for _, s := range p.MockQueueManager.Script[testQueue] {
		if s.Name == target {
			t.Error("deleted task still exists")
		}
	}

}
