// +build remote
//
// cloud/azure/compute_test.go
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

func TestCreateInstance(t *testing.T) {
	//t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	if cfg.OS.PublisherName == "" {
		str, _ := cfg.String()
		t.Fatal(str)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	s, err := NewComputeService(ctx, cfg, logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	name := fmt.Sprintf("test%v-instance", time.Now().Unix())
	task := script.Script{}
	err = s.CreateInstance(ctx, name, &task, 30)
	if err != nil {
		t.Fatal(err.Error())
	}

	instances, err := s.Instances(ctx)
	if err != nil {
		t.Error(err.Error())
	} else if _, exist := instances[name]; !exist {
		t.Error("Created instance is not found")
	}

	err = s.DeleteInstance(ctx, name)
	if err != nil {
		t.Fatal(err.Error())
	}

	instances, err = s.Instances(ctx)
	if err != nil {
		t.Error(err.Error())
	} else if _, exist := instances[name]; exist {
		t.Error("Deleted instance is found")
	}

}

func TestAvailableMachineTypes(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	ctx := context.Background()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	s, err := NewComputeService(ctx, cfg, logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := s.AvailableMachineTypes(ctx)
	if err != nil {
		t.Error(err.Error())
	}

	if len(res) == 0 {
		t.Error("There are no available machine types")
	}
	for _, v := range res {
		t.Logf("name: %v, description: %v", v.Name, v.Description)
	}

}

func TestImagePublishers(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	ctx := context.Background()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	s, err := NewComputeService(ctx, cfg, logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := s.ImagePublishers(ctx)
	if err != nil {
		t.Error(err.Error())
	}

	if len(res) == 0 {
		t.Error("There are no publishers")
	}
	for _, v := range res {
		t.Log(v.Name)
	}

}

func TestImageOffers(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	ctx := context.Background()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	s, err := NewComputeService(ctx, cfg, logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := s.ImageOffers(ctx, "Canonical")
	if err != nil {
		t.Error(err.Error())
	}

	if len(res) == 0 {
		t.Error("There are no images")
	}
	for _, v := range res {
		t.Log(v.Name)
	}

}

func TestImageSkus(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	ctx := context.Background()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	s, err := NewComputeService(ctx, cfg, logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := s.ImageSkus(ctx, "Canonical", "UbuntuServer")
	if err != nil {
		t.Error(err.Error())
	}

	if len(res) == 0 {
		t.Error("There are no images")
	}
	for _, v := range res {
		t.Log(v.Name)
	}

}

func TestImageVersions(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	ctx := context.Background()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	s, err := NewComputeService(ctx, cfg, logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := s.ImageVersions(ctx, "Canonical", "UbuntuServer", "16.10")
	if err != nil {
		t.Error(err.Error())
	}

	if len(res) == 0 {
		t.Error("There are no image versions")
	}
	for _, v := range res {
		t.Log(v.Name)
	}

}
