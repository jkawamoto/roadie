//
// command/log/cloudlogging.go
//
// Copyright (c) 2016 Junpei Kawamoto
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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package log

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	logging "google.golang.org/api/logging/v2beta1"
)

// CloudLoggingService implements LogEntryRequester interface.
// It requests logs to google cloud logging service.
type CloudLoggingService struct {
	service *logging.Service
}

// NewCloudLoggingService creates a new CloudLoggingService with a given context.
func NewCloudLoggingService(ctx context.Context) (res *CloudLoggingService, err error) {

	client, err := google.DefaultClient(ctx, logging.CloudPlatformReadOnlyScope)
	if err != nil {
		return
	}

	service, err := logging.New(client)
	if err != nil {
		return
	}

	return &CloudLoggingService{service: service}, nil

}

// Do requests a given request with the specified context.
func (s *CloudLoggingService) Do(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {

	return s.service.Entries.List(req).Do()

}
