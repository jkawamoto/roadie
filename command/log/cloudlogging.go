//
// command/log/cloudlogging.go
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

package log

import (
	"context"

	"cloud.google.com/go/logging/logadmin"
	"google.golang.org/api/iterator"
)

// CloudLoggingService implements LogEntryRequester interface.
// It requests logs to google cloud logging service.
type CloudLoggingService struct {
	// Context for this service.
	ctx context.Context
}

// NewCloudLoggingService creates a new CloudLoggingService with a given context.
func NewCloudLoggingService(ctx context.Context) (res *CloudLoggingService) {

	return &CloudLoggingService{
		ctx: ctx,
	}

}

// Entries get log entries matching with a given filter from given project logs.
// Found log entries will be passed a given handler one by one.
// If the handler returns non-nil value as an error, this function will end.
func (s *CloudLoggingService) Entries(project, filter string, handler EntryHandler) (err error) {

	client, err := logadmin.NewClient(s.ctx, project)
	if err != nil {
		return
	}
	defer client.Close()

	iter := client.Entries(s.ctx, logadmin.Filter(filter))
	for {
		select {
		case <-s.ctx.Done():
			return s.ctx.Err()

		default:
			e, err := iter.Next()
			if err == iterator.Done {
				return nil
			} else if err != nil {
				return err
			}
			if err := handler(e); err != nil {
				return err
			}
		}
	}
}
