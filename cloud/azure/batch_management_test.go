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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/arm/batch"
)

type mockBatchAcountServer struct {
	subscriptionID string
	resourceGroup  string
	location       string
	accounts       []batch.Account
	primaryKey     string
	secondaryKey   string
}

func (m *mockBatchAcountServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	paths := strings.Split(strings.TrimPrefix(req.URL.Path, "/"), "/")
	if len(paths) < 2 {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "invalid URL: %v", req.URL)
		return
	}
	if paths[1] != m.subscriptionID {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "subscription ID is %q, want %v", paths[1], m.subscriptionID)
		return
	}

	switch req.Method {
	case http.MethodGet:
		if len(paths) < 5 {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		if m.accounts == nil {
			m.accounts = []batch.Account{}
		}
		json.NewEncoder(res).Encode(&batch.AccountListResult{
			Value: &m.accounts,
		})
		return

	case http.MethodPost:
		if len(paths) < 9 {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		if paths[3] != m.resourceGroup {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(res, "resource group is %q, want %v", paths[3], m.resourceGroup)
			return
		}
		switch paths[8] {
		case "listKeys":
			json.NewEncoder(res).Encode(&batch.AccountKeys{
				AccountName: &paths[7],
				Primary:     &m.primaryKey,
				Secondary:   &m.secondaryKey,
			})
			return

		case "regenerateKeys":
			var params batch.AccountRegenerateKeyParameters
			json.NewDecoder(req.Body).Decode(&params)
			if params.KeyName == batch.Primary {
				m.primaryKey = fmt.Sprint(time.Now().Unix())
			} else {
				m.secondaryKey = fmt.Sprint(time.Now().Unix())
			}
			json.NewEncoder(res).Encode(&batch.AccountKeys{
				AccountName: &paths[7],
				Primary:     &m.primaryKey,
				Secondary:   &m.secondaryKey,
			})
			return

		}

	case http.MethodPut:
		if len(paths) < 8 {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		account := batch.Account{
			Name:     &paths[7],
			Location: &m.location,
		}
		m.accounts = append(m.accounts, account)
		json.NewEncoder(res).Encode(&account)
		return

	case http.MethodDelete:
		if len(paths) < 8 {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		selected := -1
		for i, a := range m.accounts {
			if *a.Name == paths[7] {
				selected = i
				break
			}
		}
		if selected == -1 {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		m.accounts = append(m.accounts[:selected], m.accounts[selected+1:]...)
		res.WriteHeader(http.StatusOK)
		return

	}
	res.WriteHeader(http.StatusNotImplemented)

}

func TestBatchManagementService(t *testing.T) {

	var err error
	testSubscriptionID := "00000000-0000-0000-0000-999999999999"
	testResourceGroup := "test-group"
	testAccount := "test-account"
	testLocation := "somewhere"
	testStorage := "test-storage"

	mock := mockBatchAcountServer{
		subscriptionID: testSubscriptionID,
		resourceGroup:  testResourceGroup,
		location:       testLocation,
	}
	server := httptest.NewServer(&mock)
	defer server.Close()

	ctx := context.Background()
	manager := batchAccountManager{
		client: batch.NewAccountClientWithBaseURI(server.URL, testSubscriptionID),
		Config: &Config{
			SubscriptionID:    testSubscriptionID,
			ResourceGroupName: testResourceGroup,
			BatchAccount:      testAccount,
			Location:          testLocation,
		},
		Logger: log.New(ioutil.Discard, "", log.LstdFlags),
	}

	accounts, err := manager.accounts(ctx)
	if err != nil {
		t.Fatalf("BatchAccounts returns an error: %v", err)
	}
	if len(accounts) != 0 {
		t.Errorf("%v accounts found, want %v", len(accounts), 0)
	}

	err = manager.create(ctx, testStorage)
	if err != nil {
		t.Fatalf("CreateBatchAccount returns an error: %v", err)
	}
	if len(mock.accounts) == 0 || *mock.accounts[0].Name != testAccount {
		t.Error("created account is not found.")
	}

	accounts, err = manager.accounts(ctx)
	if err != nil {
		t.Fatalf("BatchAccounts returns an error: %v", err)
	}
	if len(accounts) != 1 {
		t.Errorf("%v accounts found, want %v", len(accounts), 1)
	}
	if *accounts[0].Name != testAccount {
		t.Errorf("retrieve account name %v, want %v", *accounts[0].Name, testAccount)
	}

	key, err := manager.getKey(ctx)
	if err != nil {
		t.Fatalf("GetKey returns an error: %v", err)
	}
	key2, err := manager.getKey(ctx)
	if err != nil {
		t.Fatalf("GetKey returns an error: %v", err)
	}
	if string(key) != string(key2) {
		t.Errorf("generate keys are not matched, %v, %v", string(key), string(key2))
	}

	err = manager.delete(ctx)
	if err != nil {
		t.Fatalf("DeleteAccount returns an error: %v", err)
	}
	if len(mock.accounts) != 0 {
		t.Error("accounts are deleted but still there are some accounts")
	}

}
