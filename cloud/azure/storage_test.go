//
// cloud/azure/storage_test.go
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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strings"
	"testing"
	"time"

	arm_storage "github.com/Azure/azure-sdk-for-go/arm/storage"
	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/cloud/azure/mock"
	"github.com/jkawamoto/roadie/script"
)

func TestStorageService(t *testing.T) {

	var err error
	server := mock.NewStorageServer()
	defer server.Close()

	cli, err := server.GetClient()
	if err != nil {
		t.Fatalf("cannot get a client: %v", err)
	}

	t.Run("upload", func(t *testing.T) {
		testContainer := "upload"
		s := &StorageService{
			Client: cli.GetBlobService(),
			Logger: log.New(ioutil.Discard, "", log.LstdFlags),
		}

		var loc *url.URL
		ctx := context.Background()
		testFilename := "filename"
		loc, err = url.Parse(script.RoadieSchemePrefix + path.Join(testContainer, testFilename))
		if err != nil {
			t.Fatalf("cannot parse a URL: %v", err)
		}
		testData := "abcdefg"

		err = s.Upload(ctx, loc, strings.NewReader(testData))
		if err != nil {
			t.Fatalf("Upload returns an error: %v", err)
		}
		if data, exist := server.Items[testContainer][testFilename]; !exist || data.Body != testData {
			t.Errorf("uploaded file is %q, want %v", data.Body, testData)
		}
	})

	t.Run("download", func(t *testing.T) {
		testContainer := "download"
		s := &StorageService{
			Client: cli.GetBlobService(),
			Logger: log.New(ioutil.Discard, "", log.LstdFlags),
		}

		var loc *url.URL
		ctx := context.Background()
		testFilename := "filename"
		loc, err = url.Parse(script.RoadieSchemePrefix + path.Join(testContainer, testFilename))
		if err != nil {
			t.Fatalf("cannot parse a URL: %v", err)
		}
		testData := "abcdefg"
		err = s.Upload(ctx, loc, strings.NewReader(testData))
		if err != nil {
			t.Fatalf("Upload returns an error: %v", err)
		}

		var output bytes.Buffer
		err = s.Download(ctx, loc, &output)
		if err != nil {
			t.Fatalf("Download returns an error: %v", err)
		}
		if output.String() != testData {
			t.Errorf("downloaded file is %q, want %v", output.String(), testData)
		}
	})

	t.Run("get file info", func(t *testing.T) {
		testContainer := "fileinfo"
		s := &StorageService{
			Client: cli.GetBlobService(),
			Logger: log.New(ioutil.Discard, "", log.LstdFlags),
		}

		var loc *url.URL
		ctx := context.Background()
		testFilename := "filename"
		loc, err = url.Parse(script.RoadieSchemePrefix + path.Join(testContainer, testFilename))
		if err != nil {
			t.Fatalf("cannot parse a URL: %v", err)
		}
		testData := "abcdefg"
		err = s.Upload(ctx, loc, strings.NewReader(testData))
		if err != nil {
			t.Fatalf("Upload returns an error: %v", err)
		}

		var info *cloud.FileInfo
		info, err = s.GetFileInfo(ctx, loc)
		if err != nil {
			t.Fatalf("GetFileInfo returns an error: %v", err)
		}
		if info.Name != testFilename || info.URL.String() != loc.String() {
			t.Errorf("retrieved file info is incorrect: %v", info)
		}
	})

	t.Run("list", func(t *testing.T) {
		testContainer := "list"
		s := &StorageService{
			Client: cli.GetBlobService(),
			Logger: log.New(ioutil.Discard, "", log.LstdFlags),
		}

		var loc *url.URL
		ctx := context.Background()
		testFiles := []string{"file1", "file2", "file3"}
		for _, f := range testFiles {
			loc, err = url.Parse(script.RoadieSchemePrefix + path.Join(testContainer, f))
			if err != nil {
				t.Fatalf("cannot parse a URL: %v", err)
			}
			err = s.Upload(ctx, loc, strings.NewReader(f))
			if err != nil {
				t.Fatalf("Upload returns an error: %v", err)
			}
		}

		var query *url.URL
		query, err = url.Parse(script.RoadieSchemePrefix + testContainer)
		if err != nil {
			t.Fatalf("cannot parse a URL: %v", err)
		}
		err = s.List(ctx, query, func(info *cloud.FileInfo) error {
			if !strings.HasPrefix(info.Name, "file") {
				t.Errorf("retrieved file is incorrect: %v", info)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("List returns an error: %v", err)
		}

	})

	t.Run("delete", func(t *testing.T) {
		testContainer := "delete"
		s := &StorageService{
			Client: cli.GetBlobService(),
			Logger: log.New(ioutil.Discard, "", log.LstdFlags),
		}

		var loc *url.URL
		ctx := context.Background()
		testFilename := "filename"
		loc, err = url.Parse(script.RoadieSchemePrefix + path.Join(testContainer, testFilename))
		if err != nil {
			t.Fatalf("cannot parse a URL: %v", err)
		}
		testData := "abcdefg"
		err = s.Upload(ctx, loc, strings.NewReader(testData))
		if err != nil {
			t.Fatalf("Upload returns an error: %v", err)
		}

		err = s.Delete(ctx, loc)
		if err != nil {
			t.Fatalf("Delete returns an error: %v", err)
		}
		if _, exist := server.Items[testContainer][testFilename]; exist {
			t.Error("deleted file still exists")
		}
	})

	t.Run("upload with properties and metadata", func(t *testing.T) {
		testContainer := "upload-metadata"
		s := &StorageService{
			Client: cli.GetBlobService(),
			Logger: log.New(ioutil.Discard, "", log.LstdFlags),
		}

		ctx := context.Background()
		testFilename := "filename"
		testData := "abcdefg"
		err = s.UploadWithMetadata(ctx, testContainer, testFilename, strings.NewReader(testData), &storage.BlobProperties{
			ContentType: "text/yaml",
		}, map[string]string{
			"Version": "test",
		})
		if err != nil {
			t.Fatalf("Upload returns an error: %v", err)
		}
		if data, exist := server.Items[testContainer][testFilename]; !exist {
			t.Error("uploaded file isn't found")
		} else {
			if data.Body != testData {
				t.Errorf("body of the uploaded file is %q, want %v", data.Body, testData)
			}
			if data.Properties.ContentType != "text/yaml" {
				t.Errorf("content type property %q, want %v", data.Properties.ContentType, "text/yaml")
			}
			if len(data.Metadata) != 1 {
				t.Errorf("%v key-values are stored in the metadata, want %v", len(data.Metadata), 1)
			}
			for k, v := range data.Metadata {
				if k == "Version" && v[0] != "test" {
					t.Errorf("metadata %v is %q, want %v", k, v, "test")
				}
			}
		}
	})

	t.Run("get metadata", func(t *testing.T) {
		testContainer := "get-metadata"
		s := &StorageService{
			Client: cli.GetBlobService(),
			Logger: log.New(ioutil.Discard, "", log.LstdFlags),
		}

		ctx := context.Background()
		testFilename := "filename"
		testData := "abcdefg"
		err = s.UploadWithMetadata(ctx, testContainer, testFilename, strings.NewReader(testData), &storage.BlobProperties{
			ContentType: "text/yaml",
		}, map[string]string{
			"Version": "test",
		})
		if err != nil {
			t.Fatalf("UploadWithMetadate returns an error: %v", err)
		}

		metadata, err := s.GetMetadata(ctx, testContainer, testFilename)
		if err != nil {
			t.Fatalf("GetMetadata returns an error: %v", err)
		}
		if len(metadata) != 1 {
			t.Errorf("%v key-values are stored in the metadata, want %v", len(metadata), 1)
		}
		for k, v := range metadata {
			if k == "Version" && v != "test" {
				t.Errorf("metadata %v is %q, want %v", k, v, "test")
			}
		}
	})

}

// mockStorageAccountServer is a mock providing storage account managment service.
type mockStorageAccountServer struct {
	subscriptionID string
	accountName    string
	location       string
	accounts       []arm_storage.Account
	keys           []arm_storage.AccountKey
}

func (m *mockStorageAccountServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {

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
		if len(paths) < 8 {
			// List accounts.
			if m.accounts == nil {
				m.accounts = []arm_storage.Account{}
			}
			json.NewEncoder(res).Encode(&arm_storage.AccountListResult{
				Value: &m.accounts,
			})
			return
		}

		for _, a := range m.accounts {
			if *a.Name == paths[7] {
				json.NewEncoder(res).Encode(a)
				return
			}
		}

		res.WriteHeader(http.StatusNotFound)
		return

	case http.MethodPost:
		if len(paths) < 9 || paths[3] != m.accountName {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		if m.keys == nil {
			m.keys = []arm_storage.AccountKey{}
		}
		switch paths[8] {
		case "listKeys":
			json.NewEncoder(res).Encode(&arm_storage.AccountListKeysResult{
				Keys: &m.keys,
			})
			return

		case "regenerateKey":
			var key arm_storage.AccountKey
			json.NewDecoder(req.Body).Decode(&key)
			key.Value = toPtr(fmt.Sprint(time.Now().Unix()))

			m.keys = append(m.keys, key)
			json.NewEncoder(res).Encode(&arm_storage.AccountListKeysResult{
				Keys: &m.keys,
			})
			return

		}

	case http.MethodPut:
		if len(paths) < 7 || paths[3] != m.accountName {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(res, "resource group name is invald")
			return
		}
		var param arm_storage.AccountCreateParameters
		json.NewDecoder(req.Body).Decode(&param)
		if param.Kind != arm_storage.BlobStorage {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(res, "kind is %q, want %v", param.Kind, arm_storage.BlobStorage)
			return
		}
		if param.Sku == nil || param.Sku.Name != arm_storage.StandardRAGRS || param.Sku.Tier != arm_storage.Standard {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(res, "sku is wrong: %v", param.Sku)
			return
		}
		if param.Location == nil || *param.Location != m.location {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(res, "location is wrong: %v", param.Location)
			return
		}
		if param.AccountPropertiesCreateParameters == nil || param.AccountPropertiesCreateParameters.AccessTier != arm_storage.Hot {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(res, "property is wrong: %v", param.AccountPropertiesCreateParameters)
			return
		}
		m.accounts = append(m.accounts, arm_storage.Account{
			Name: &paths[7],
		})
		res.WriteHeader(http.StatusOK)
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
		return

	}
	res.WriteHeader(http.StatusNotImplemented)

}

func TestStorageAccountManager(t *testing.T) {

	var err error
	testSubscriptionID := "00000000-0000-0000-0000-999999999999"
	testAccount := "test-account"
	testLocation := "somewhere"

	mock := mockStorageAccountServer{
		subscriptionID: testSubscriptionID,
		accountName:    testAccount,
		location:       testLocation,
	}
	server := httptest.NewServer(&mock)
	defer server.Close()

	manager := storageAccountManager{
		client: arm_storage.NewAccountsClientWithBaseURI(server.URL, testSubscriptionID),
		Config: &Config{
			SubscriptionID: testSubscriptionID,
			AccountName:    testAccount,
			Location:       testLocation,
		},
		Logger: log.New(ioutil.Discard, "", log.LstdFlags),
	}

	err = manager.createIfNotExists(context.Background())
	if err != nil {
		t.Fatalf("createIfNotExists returns an error: %v", err)
	}
	if len(mock.accounts) == 0 || *mock.accounts[0].Name != testAccount {
		t.Error("created account is not found")
	}

	account, err := manager.getStorageAccountInfo()
	if err != nil {
		t.Fatalf("getStorageAccountInfo returns an error: %v", err)
	}
	if *account.Name != testAccount {
		t.Errorf("retrieved account name is %v, want %v", *account.Name, testAccount)
	}

	key, err := manager.getStorageKey(context.Background())
	if err != nil {
		t.Fatalf("getStorageKey returns an error: %v", err)
	}

	key2, err := manager.getStorageKey(context.Background())
	if err != nil {
		t.Fatalf("getStorageKey returns an error: %v", err)
	}
	if key != key2 {
		t.Errorf("getStorageKey regenerates keys even if some keys are already registered.")
	}

	err = manager.delete()
	if err != nil {
		t.Fatalf("delete returns an error: %v", err)
	}

}
