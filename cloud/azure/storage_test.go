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
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
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
	"github.com/jkawamoto/roadie/script"
)

const (
	StorageAPIVersion = "2017-06-01"
)

// ForwardTransport is a http.RoundTripper which forwards requests to a given
// target URL's host.
type ForwardTransport struct {
	target *url.URL
	parent http.RoundTripper
}

// NewForwardTransport creates a round tripper which forwards requests to a given
// target's host. It also sends request a given parent round tripper.
// If the parent is not given, http.DefaultTransport will be used.
func NewForwardTransport(target *url.URL, parent http.RoundTripper) http.RoundTripper {
	if parent == nil {
		parent = http.DefaultTransport
	}
	return &ForwardTransport{
		target: target,
		parent: parent,
	}
}

// RoundTrip updates request URL's hosts to the target URL's ones, and then
// sends the request to the parent round tripper.
func (r *ForwardTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Host = r.target.Host
	return r.parent.RoundTrip(req)
}

// mockObject represents an object stored in mockStorageServer.
type mockObject struct {
	Body       string
	Properties storage.BlobProperties
	Metadata   map[string][]string
}

// mockContainer represents a container stored in mockStorageServer.
type mockContainer map[string]mockObject

// See https://docs.microsoft.com/en-us/rest/api/storageservices/
type mockStorageServer struct {
	// Items hold blob items in this mock server. This is a map of which keys
	// represent container names and the associated value represents a container.
	Items map[string]mockContainer
	// server is a pointer for a httptest's server.
	server *httptest.Server
}

// newMockStorageServer creates a new mock storage server.
func newMockStorageServer() (res *mockStorageServer) {

	res = &mockStorageServer{
		Items: make(map[string]mockContainer),
	}
	res.server = httptest.NewServer(res)
	return

}

func (s *mockStorageServer) Close() {
	s.server.Close()
}

func (s *mockStorageServer) GetClient() (cli storage.Client, err error) {

	URL, err := url.Parse(s.server.URL)
	if err != nil {
		return
	}
	cli, err = storage.NewClient("name", base64.StdEncoding.EncodeToString([]byte("key")), URL.Host, StorageAPIVersion, false)
	if err != nil {
		return
	}
	cli.HTTPClient.Transport = NewForwardTransport(URL, nil)
	return

}

func (s *mockStorageServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	var container, filename string
	if sp := strings.SplitN(strings.TrimPrefix(req.URL.Path, "/"), "/", 2); len(sp) == 1 {
		container = sp[0]
	} else {
		container, filename = sp[0], sp[1]
	}

	switch req.Method {
	case http.MethodHead:
		_, exist := s.Items[container]
		if !exist {
			res.WriteHeader(http.StatusNotFound)
			return
		}

		if filename == "" {
			res.WriteHeader(http.StatusNotImplemented)

		} else {

			data, exist := s.Items[container][filename]
			if !exist {
				res.WriteHeader(http.StatusNotFound)
				return
			}
			res.Header().Add("Last-Modified", time.Now().Format(time.RFC1123))
			res.Header().Add("Content-Length", fmt.Sprint(len(data.Body)))

		}

	case http.MethodGet:
		_, exist := s.Items[container]
		if !exist {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		switch {
		case filename == "" && req.URL.Query().Get("restype") == "container" && req.URL.Query().Get("comp") == "list":
			// List blobs.
			var blobs []storage.Blob
			for name := range s.Items[container] {
				blobs = append(blobs, storage.Blob{
					Name:     name,
					Snapshot: time.Now(),
				})
			}
			data, err := xml.Marshal(&storage.BlobListResponse{
				Blobs: blobs,
			})
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				io.WriteString(res, err.Error())
				return
			}
			data = []byte(strings.Replace(strings.Replace(
				string(data),
				"<CopyCompletionTime></CopyCompletionTime>",
				fmt.Sprintf("<CopyCompletionTime>%v</CopyCompletionTime>", time.Now().Format(time.RFC1123)),
				-1),
				"<Last-Modified></Last-Modified>",
				fmt.Sprintf("<Last-Modified>%v</Last-Modified>", time.Now().Format(time.RFC1123)),
				-1))
			res.Write(data)

		case filename != "" && req.URL.Query().Get("comp") == "metadata":
			// Get blob metadata.
			data, exist := s.Items[container][filename]
			if !exist {
				res.WriteHeader(http.StatusNotFound)
				return
			}
			for key, value := range data.Metadata {
				res.Header().Add("X-Ms-Meta-"+key, strings.Join(value, ","))
			}

		case filename != "":
			// Get blob.
			data, exist := s.Items[container][filename]
			if !exist {
				res.WriteHeader(http.StatusNotFound)
				return
			}
			io.WriteString(res, data.Body)

		default:
			res.WriteHeader(http.StatusNotImplemented)

		}

	case http.MethodPut:
		switch {
		case filename == "" && req.URL.Query().Get("restype") == "container":
			// Operations for container.

			switch req.URL.Query().Get("comp") {
			case "acl":
				// Set Container ACL
				var props storage.SignedIdentifiers
				err := xml.NewDecoder(req.Body).Decode(&props)
				if err != nil {
					res.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(res, err)
					return
				}
				if len(props.SignedIdentifiers) == 0 {
					res.WriteHeader(http.StatusBadRequest)
					fmt.Fprint(res, "no signed identifiers are given")
					return
				}
				acl := props.SignedIdentifiers[0]
				if acl.ID != "full-access" {
					res.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(res, "ID is %q, want %q", acl.ID, "full-access")
					return
				}

			default:
				// Create container.
				if _, exist := s.Items[container]; exist {
					res.WriteHeader(http.StatusConflict)
					return
				}
				s.Items[container] = make(mockContainer)
				res.WriteHeader(http.StatusCreated)
			}

		case filename != "" && req.URL.Query().Get("comp") == "appendblock":
			// Append block.
			if blobType := req.Header.Get("X-Ms-Blob-Type"); blobType != "AppendBlob" {
				res.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(res, "Blob type is %q, want %v", blobType, "AppendBlob")
				return
			}
			if data, err := ioutil.ReadAll(req.Body); err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				io.WriteString(res, err.Error())
			} else {
				obj := s.Items[container][filename]
				obj.Body += string(data)
				s.Items[container][filename] = obj
				res.WriteHeader(http.StatusCreated)
			}

		case filename != "" && req.URL.Query().Get("comp") == "properties":
			// Set blob properties.
			obj := s.Items[container][filename]
			obj.Properties.ContentType = req.Header.Get("X-Ms-Blob-Content-Type")
			s.Items[container][filename] = obj
			res.WriteHeader(http.StatusOK)

		case filename != "" && req.URL.Query().Get("comp") == "metadata":
			// Set blob metadata.
			// The format of metadata is "X-Ms-Meta-name:value" in the header.
			obj := s.Items[container][filename]
			for key, value := range req.Header {
				if strings.HasPrefix(key, "X-Ms-Meta-") {
					obj.Metadata[strings.TrimPrefix(key, "X-Ms-Meta-")] = value
				}
			}
			s.Items[container][filename] = obj
			res.WriteHeader(http.StatusOK)

		case filename != "":
			// Put blob.
			if _, exist := s.Items[container][filename]; exist {
				res.WriteHeader(http.StatusBadRequest)
			} else {
				s.Items[container][filename] = mockObject{
					Metadata: make(map[string][]string),
				}
				res.WriteHeader(http.StatusCreated)
			}

		default:
			// Unknown request.
			fmt.Println(req.URL)
			res.WriteHeader(http.StatusBadRequest)

		}

	case http.MethodDelete:
		_, exist := s.Items[container]
		if !exist {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		if filename == "" {
			res.WriteHeader(http.StatusNotImplemented)
			return
		}
		_, exist = s.Items[container][filename]
		if !exist {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		delete(s.Items[container], filename)
		res.WriteHeader(http.StatusAccepted)

	default:
		res.WriteHeader(http.StatusNotImplemented)

	}

}

func TestStorageService(t *testing.T) {

	var err error
	server := newMockStorageServer()
	defer server.Close()

	cli, err := server.GetClient()
	if err != nil {
		t.Fatalf("cannot get a client: %v", err)
	}

	t.Run("upload", func(t *testing.T) {
		testContainer := "upload"
		s := &StorageService{
			blobClient: cli.GetBlobService(),
			Logger:     log.New(ioutil.Discard, "", log.LstdFlags),
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
			blobClient: cli.GetBlobService(),
			Logger:     log.New(ioutil.Discard, "", log.LstdFlags),
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
			blobClient: cli.GetBlobService(),
			Logger:     log.New(ioutil.Discard, "", log.LstdFlags),
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
			blobClient: cli.GetBlobService(),
			Logger:     log.New(ioutil.Discard, "", log.LstdFlags),
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
			blobClient: cli.GetBlobService(),
			Logger:     log.New(ioutil.Discard, "", log.LstdFlags),
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
			blobClient: cli.GetBlobService(),
			Logger:     log.New(ioutil.Discard, "", log.LstdFlags),
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
			blobClient: cli.GetBlobService(),
			Logger:     log.New(ioutil.Discard, "", log.LstdFlags),
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
