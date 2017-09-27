//
// cloud/mock/constant.go
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

import "fmt"

const (
	// StatusRunning means an instance is still running.
	StatusRunning = "running"
	// StatusTerminated means an instance has been terminated.
	StatusTerminated = "terminated"
	// StatusPending means a task is pending now.
	StatusPending = "pending"
	// StatusWaiting means a task is waiting to be run.
	StatusWaiting = "waiting"
)

var (
	// ErrServiceFailure is an error used in tests.
	ErrServiceFailure = fmt.Errorf("this service is out of order")
)
