//
// command/queue.go
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

package command

import "github.com/urfave/cli"

// CmdQueueList lists up existing queue information.
// Each information should have queue name, the number of items in the queue,
// the number of instances working to the queue.
func CmdQueueList(c *cli.Context) error {
	return nil
}

func CmdQueueStatus(c *cli.Context) error {
	return nil
}
func CmdQueueInstanceShow(c *cli.Context) error {
	return nil
}
func CmdQueueInstanceAdd(c *cli.Context) error {
	return nil
}
func CmdQueueStop(c *cli.Context) error {
	return nil
}
func CmdQueueRestart(c *cli.Context) error {
	return nil
}
