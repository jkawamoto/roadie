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
	"strings"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/script"
)

// LogManager defines a service interface for obtaining log entries.
type LogManager struct {
	storage *StorageService
	batch   *BatchService
	Config  *Config
	Logger  *log.Logger
}

// NewLogManager creates a new log manger for Azure.
func NewLogManager(ctx context.Context, cfg *Config, logger *log.Logger) (m *LogManager, err error) {

	storage, err := NewStorageService(ctx, cfg, logger)
	if err != nil {
		return
	}
	batch, err := NewBatchService(ctx, cfg, logger)
	if err != nil {
		return
	}

	m = &LogManager{
		storage: storage,
		batch:   batch,
		Config:  cfg,
		Logger:  logger,
	}
	return

}

// Get retrievs log entries.
func (m *LogManager) Get(ctx context.Context, instanceName string, from time.Time, handler cloud.LogHandler) (err error) {

	var urls []*url.URL
	var loc *url.URL
	for _, format := range []string{"%v-init.log", "%v.log"} {
		loc, err = url.Parse(script.RoadieSchemePrefix + path.Join(LogContainer, fmt.Sprintf(format, instanceName)))
		if err != nil {
			return
		}
		urls = append(urls, loc)
	}
	return m.get(ctx, urls, handler)

}

// get retrieves log files represented by a given URLs and sends each line to a given handler.
func (m *LogManager) get(ctx context.Context, urls []*url.URL, handler cloud.LogHandler) (err error) {

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
					fields := strings.SplitN(line, " ", 3)
					if len(fields) != 3 {
						continue
					}
					t, _ := time.Parse("2006/01/02 15:04:05", fmt.Sprintf("%v %v", fields[0], fields[1]))
					err = handler(t.UTC().In(time.Local), fields[2], false)
					if err != nil {
						reader.Close()
						writer.Close()
						ignore = true
					}
				}
			}
		}
	})

	wg.Go(func() (err error) {
		defer writer.Close()
		for _, loc := range urls {
			err = ignoreNotFoundError(m.storage.Download(ctx, loc, writer))
			if err != nil {
				break
			}
		}
		return
	})

	return wg.Wait()

}

// Delete instance log.
func (m *LogManager) Delete(ctx context.Context, instanceName string) error {

	var errs []error
	for _, format := range []string{"%v-init.log", "%v.log"} {
		loc, err := url.Parse(script.RoadieSchemePrefix + path.Join(LogContainer, fmt.Sprintf(format, instanceName)))
		if err != nil {
			errs = append(errs, err)
			continue
		}
		err = ignoreNotFoundError(m.storage.Delete(ctx, loc))
		if err != nil {
			errs = append(errs, err)
		}
	}
	err := ignoreNotFoundError(m.batch.DeleteJob(ctx, instanceName))
	if err != nil {
		//errs = append(errs, err)
		_ = err
	}

	if len(errs) != 0 {
		return errs[0]
	}
	return nil

}

// GetQueueLog retrieves log of a given queue.
func (m *LogManager) GetQueueLog(ctx context.Context, queue string, handler cloud.LogHandler) (err error) {

	queue = queueName(queue)
	tasks, err := m.batch.Tasks(ctx, queue)
	if err != nil {
		return
	}

	var urls []*url.URL
	var loc *url.URL
	loc, err = url.Parse(script.RoadieSchemePrefix + path.Join(LogContainer, fmt.Sprintf("%v-init.log", queue)))
	if err != nil {
		return
	}
	urls = append(urls, loc)

	for _, task := range tasks {
		loc, err = url.Parse(script.RoadieSchemePrefix + path.Join(LogContainer, fmt.Sprintf("%v.log", task.ID)))
		if err != nil {
			return
		}
		urls = append(urls, loc)
	}

	return m.get(ctx, urls, handler)

}

// GetTaskLog retrieves log of a given task.
func (m *LogManager) GetTaskLog(ctx context.Context, queue, task string, handler cloud.LogHandler) (err error) {

	// queue = queueName(queue)
	task = taskName(task)
	loc, err := url.Parse(script.RoadieSchemePrefix + path.Join(LogContainer, fmt.Sprintf("%v.log", task)))
	if err != nil {
		return
	}
	return m.get(ctx, []*url.URL{loc}, handler)

}

// ignoreNotFoundError is a wrapper function and ignores not found error.
func ignoreNotFoundError(err error) error {

	if err != nil {
		switch e := err.(type) {
		case storage.AzureStorageServiceError:
			if e.StatusCode == 404 {
				err = nil
			}
		}
	}
	return err

}
