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

	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
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
		err = s.upload(ctx, testContainer, testFilename, strings.NewReader(testData), &storage.BlobProperties{
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
		err = s.upload(ctx, testContainer, testFilename, strings.NewReader(testData), &storage.BlobProperties{
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

func TestStorageManagerImplementation(t *testing.T) {

	var s cloud.StorageManager = &StorageService{}
	_ = s

}
