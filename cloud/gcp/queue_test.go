// +build remote
//
// cloud/gcp/queue_test.go
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

package gcp

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jkawamoto/roadie/script"
)

func TestNewQueueService(t *testing.T) {

	project := "sample-project"
	region := "us-central1-c"
	machine := "n1-standard-2"
	cfg := &Config{
		Project:     project,
		Zone:        region,
		MachineType: machine,
	}

	ctx := context.Background()
	s, err := NewQueueService(ctx, cfg, nil)
	if err != nil {
		t.Error(err.Error())
	}

	if s.Config.Project != project {
		t.Error("Project name doesn't match:", s.Config.Project)
	}
	if s.Config.Zone != region {
		t.Error("Zone name doesn't match:", s.Config.Zone)
	}
	if s.Config.MachineType != machine {
		t.Error("Machine type doesn't match:", s.Config.MachineType)
	}

	if s.Logger == nil {
		t.Error("Logger is nil")
	}

}

func TestEnqueuAndFetch(t *testing.T) {

	cfg := GetConfig()
	if cfg == nil {
		t.Skip("Config file to access GCP is not given")
	}

	ctx := context.Background()
	service, err := NewQueueService(ctx, cfg, log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile))
	if err != nil {
		t.Fatal(err.Error())
	}
	queue := fmt.Sprintf("test-queue-%v", time.Now().Unix())
	taskName := "test-task"

	err = service.Enqueue(ctx, queue, &script.Script{
		InstanceName: taskName,
		Run:          []string{"echo test"},
		Result:       fmt.Sprintf("gs://%v/%v/result/test", cfg.Bucket, StoragePrefix),
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	task, err := service.Fetch(ctx, queue)
	if err != nil {
		t.Error(err.Error())
	} else if task == nil {
		t.Error("Cannot fatch the added task")
	} else if task.Name != taskName {
		t.Error("Name of the fetched task is not correct:", task.Name)
	}

	service.DeleteQueue(ctx, queue)

}
