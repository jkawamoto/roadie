//
// cloud/azure/log_manager.go
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
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"path"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
)

// LogManager defines a service interface for obtaining log entries.
type LogManager struct {
	service *StorageService
	Config  *AzureConfig
	Logger  *log.Logger
}

// NewLogManager creates a new log manger for Azure.
func NewLogManager(ctx context.Context, cfg *AzureConfig, logger *log.Logger) (m *LogManager, err error) {

	service, err := NewStorageService(ctx, cfg, logger)
	if err != nil {
		return
	}

	m = &LogManager{
		service: service,
		Config:  cfg,
		Logger:  logger,
	}
	return

}

// Get retrievs log entries.
func (m *LogManager) Get(ctx context.Context, instanceName string, from time.Time, handler cloud.LogHandler) (err error) {

	loc, err := url.Parse(script.RoadieSchemePrefix + path.Join(LogContainer, fmt.Sprintf("%v.log", instanceName)))
	if err != nil {
		return
	}

	ch := make(chan string)
	wg, ctx := errgroup.WithContext(ctx)
	reader, writer := io.Pipe()

	wg.Go(func() error {
		defer reader.Close()
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
		return nil
	})

	wg.Go(func() (err error) {
		var ignore bool
		for {
			select {
			case <-ctx.Done():
				reader.Close()
				writer.Close()
				ignore = true
				err = ctx.Err()
			case line, ok := <-ch:
				if !ok {
					return
				}
				if !ignore {
					err = handler(time.Time{}, line, false)
					if err != nil {
						reader.Close()
						writer.Close()
						ignore = true
					}
				}
			}
		}
	})

	wg.Go(func() error {
		defer writer.Close()
		return m.service.Download(ctx, loc, writer)
	})

	return wg.Wait()

}

// Delete instance log.
func (m *LogManager) Delete(ctx context.Context, instanceName string) error {
	return fmt.Errorf("not implemented")
}

// GetQueueLog retrievs log entries from a given queue.
func (m *LogManager) GetQueueLog(ctx context.Context, queue string, handler cloud.LogHandler) error {
	return fmt.Errorf("not implemented")
}

// GetTaskLog retrieves log entries from a task in a queue.
func (m *LogManager) GetTaskLog(ctx context.Context, queue, task string, handler cloud.LogHandler) error {
	return fmt.Errorf("not implemented")
}
