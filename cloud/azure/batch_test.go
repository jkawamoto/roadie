// +build remote
//
// cloud/azure/batch_test.go
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

package azure

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jkawamoto/roadie/script"
)

func TestCreateJob(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.New(os.Stdout, "", log.Lshortfile)
	s, err := NewBatchService(ctx, cfg, logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	name := fmt.Sprintf("test%v-job", time.Now().Unix())
	err = s.CreateJob(ctx, name)
	if err != nil {
		t.Fatal(err.Error())
	}

}

func TestCreateInstanceByBatch(t *testing.T) {
	//t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.New(os.Stdout, "", log.Lshortfile)
	s, err := NewBatchService(ctx, cfg, logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	name := fmt.Sprintf("test%v-job", time.Now().Unix())
	err = s.CreateJob(ctx, name)
	if err != nil {
		t.Fatal(err.Error())
	}

	task := script.Script{
		Source: "https://github.com/itslab-kyushu/sss/releases/download/v0.3.2/sss_0.3.2_linux_amd64.tar.gz",
		Run: []string{
			`echo "abcdefg" > sample.dat`,
			"./sss_0.3.2_linux_amd64/sss local distribute sample.dat 5 2",
		},
		Upload: []string{
			"sample.dat.0.xz",
			"sample.dat.1.xz",
			"sample.dat.2.xz",
			"sample.dat.3.xz",
			"sample.dat.4.xz",
		},
		InstanceName: "sss",
	}
	err = s.CreateTask(ctx, name, &task)
	if err != nil {
		t.Error(err.Error())
	}

}

func TestBatchAvailableMachineTypes(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	s, err := NewBatchService(ctx, cfg, logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	types, err := s.AvailableMachineTypes(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(types) == 0 {
		t.Error("No machine types are found")
	}
	for _, v := range types {
		t.Log(v.Name, ":", v.Description)
	}

}
