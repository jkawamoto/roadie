//
// command/log/requester.go
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

import logging "google.golang.org/api/logging/v2beta1"

// EntryRequester is an interface used in GetLogEntries.
// This interface requests supplying Do method which process a request of
// obtaining log entries.
type EntryRequester interface {
	Do(*logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error)
}

// EntryRequesterFunc will be used to implement EntryRequester interface
// on functions.
type EntryRequesterFunc func(*logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error)

// Do implements EntryRequester interface.
func (f EntryRequesterFunc) Do(req *logging.ListLogEntriesRequest) (*logging.ListLogEntriesResponse, error) {
	return f(req)
}
