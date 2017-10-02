// +build remote
//
// cloud/azure/batch_management_test.go
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

func TestCreateBatchAccount(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}
	cfg.BatchAccount = fmt.Sprintf("test%vbatch", time.Now().Unix())
	cfg.StorageAccount = fmt.Sprintf("test%vstorage", time.Now().Unix())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	s, err := NewBatchManagementService(ctx, cfg, logger)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = s.CreateBatchAccount(ctx)
	if err != nil {
		t.Error(err.Error())
	}

	accounts, err := s.BatchAccounts(ctx)
	if err != nil {
		t.Error(err.Error())
	} else if _, exists := accounts[cfg.BatchAccount]; !exists {
		t.Error("Created account does not exist")
	}

	err = s.DeleteAccount(ctx)
	if err != nil {
		t.Error(err.Error())
	}

	accounts, err = s.BatchAccounts(ctx)
	if err != nil {
		t.Error(err.Error())
	} else if _, exists := accounts[cfg.BatchAccount]; exists {
		t.Error("Deleted account still exists")
	}

	storage, err := NewStorageService(ctx, cfg, logger)
	if err != nil {
		t.Error(err.Error())
	}
	err = storage.deleteStorageAccount(ctx)
	if err != nil {
		t.Error(err.Error())
	}

}
