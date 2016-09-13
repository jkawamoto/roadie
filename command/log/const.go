//
// command/log/const.go
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

const (
	// LogTimeFormat defines time format of Google Logging.
	LogTimeFormat = "2006-01-02T15:04:05Z"

	// EventSubtypeInsert means this event is creating an instance.
	EventSubtypeInsert = "compute.instances.insert"

	// EventSubtypeDelete means this event is deleting an instance.
	EventSubtypeDelete = "compute.instances.delete"
)
