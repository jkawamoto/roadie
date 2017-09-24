//
// cloud/mock/log.go
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
	"context"
	"fmt"
	"io"
	"time"

	"github.com/jkawamoto/roadie/cloud"
)

var (
	zeroTime = time.Time{}
	farTime  = time.Now().Add(3600 * time.Hour)
)

// LogManager is a mock log manager.
type LogManager struct {
	// Failure is set true, all methods will return ErrServiceFailure.
	Failure bool
	// Logs is a map to maintain log entries. The key of the map is instance names
	// and associated values are log entries for the instance.
	Logs map[string][]LogEntry
	// KeepAlive is true then Get returns nil otherwise io.EOF.
	KeepAlive bool
	// Break defines a break point when Get returns entries before that point.
	// Onece, the break point is used, it will be removed and KeepAlive will be
	// false.
	Break time.Time
}

// LogEntry defines a log entry which has a time and a body.
type LogEntry struct {
	// Time is the time the log entry posted.
	Time time.Time
	// Body is the message body of this log entry.
	Body string
	// Stderr to output this entry to stderr instead of stdout.
	Stderr bool
}

// NewLogManager creates a new mock log manager.
func NewLogManager() *LogManager {
	return &LogManager{}
}

// Get instance log.
func (m *LogManager) Get(ctx context.Context, instanceName string, after time.Time, handler cloud.LogHandler) (err error) {

	if m.Failure {
		return ErrServiceFailure
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	for _, e := range m.Logs[instanceName] {
		if e.Time.After(after) && (m.Break == zeroTime || e.Time.Before(m.Break)) {
			err = handler(e.Time, e.Body, e.Stderr)
			if err != nil {
				return
			}
		}
	}

	if m.Break != zeroTime {
		m.Break = zeroTime
		return nil
	}
	if err == nil && m.KeepAlive {
		return io.EOF
	}
	return

}

// Delete instance log.
func (m *LogManager) Delete(ctx context.Context, instanceName string) error {

	if m.Failure {
		return ErrServiceFailure
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if _, exist := m.Logs[instanceName]; !exist {
		return fmt.Errorf("logs for %q do not exist", instanceName)
	}
	delete(m.Logs, instanceName)
	return nil

}

func (m *LogManager) GetQueueLog(ctx context.Context, queue string, handler cloud.LogHandler) error {
	return nil
}

func (m *LogManager) GetTaskLog(ctx context.Context, queue, task string, handler cloud.LogHandler) error {
	return nil
}
