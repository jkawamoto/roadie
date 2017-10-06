//
// cloud/azure/mock/storage.go
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

package mock

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/storage"
)

const (
	// StorageAPIVersion is the API version this mock server follows.
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

// Object represents an object stored in StorageServer.
type Object struct {
	Body       string
	Properties storage.BlobProperties
	Metadata   map[string][]string
}

// Container represents a container stored in StorageServer.
type Container map[string]Object

// StorageServer provides a mocked server.
// See https://docs.microsoft.com/en-us/rest/api/storageservices/
type StorageServer struct {
	// Items hold blob items in this mock server. This is a map of which keys
	// represent container names and the associated value represents a container.
	Items map[string]Container
	// server is a pointer for a httptest's server.
	server *httptest.Server
	// lock is a locker.
	lock sync.Mutex
}

// NewStorageServer creates a new mock storage server.
func NewStorageServer() (res *StorageServer) {

	res = &StorageServer{
		Items: make(map[string]Container),
	}
	res.server = httptest.NewServer(res)
	return

}

// Close closes the mock server.
func (s *StorageServer) Close() {
	s.server.Close()
}

// GetClient returns a storage client which is forareded to this mock server.
func (s *StorageServer) GetClient() (cli storage.Client, err error) {

	URL, err := url.Parse(s.server.URL)
	if err != nil {
		return
	}
	cli, err = storage.NewClient("name", base64.StdEncoding.EncodeToString([]byte("key")), URL.Host, StorageAPIVersion, false)
	if err != nil {
		return
	}
	httpClient := *cli.HTTPClient
	cli.HTTPClient = &httpClient
	cli.HTTPClient.Transport = NewForwardTransport(URL, nil)
	return

}

// ServeHTTP handles requests.
func (s *StorageServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	var container, filename string
	if sp := strings.SplitN(strings.TrimPrefix(req.URL.Path, "/"), "/", 2); len(sp) == 1 {
		container = sp[0]
	} else {
		container, filename = sp[0], sp[1]
	}

	s.lock.Lock()
	defer s.lock.Unlock()

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
				s.Items[container] = make(Container)
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
				s.Items[container][filename] = Object{
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
