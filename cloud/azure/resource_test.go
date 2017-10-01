// +build remote
//
// cloud/azure/resource_test.go
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
)

func TestResourceService(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	ctx := context.Background()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	s := NewResourceService(cfg, logger)

	name := fmt.Sprintf("resource%v", time.Now().Unix())
	err = s.CreateResourceGroup(ctx, name)
	if err != nil {
		t.Fatal(err.Error())
	}

	groups, err := s.ResourceGroups(ctx)
	if err != nil {
		t.Error(err.Error())
	}
	if _, exist := groups[name]; !exist {
		t.Error("Created resource group isn't found")
	}

	err = s.DeleteResourceGroup(ctx, name)
	if err != nil {
		t.Error(err.Error())
	}

	groups, err = s.ResourceGroups(ctx)
	if err != nil {
		t.Error(err.Error())
	}
	if _, exist := groups[name]; exist {
		t.Error("Deleted resource group is found")
	}

}
