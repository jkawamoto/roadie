//
// main.go
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

package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jkawamoto/roadie/command"
	"github.com/urfave/cli"
)

func main() {

	// Prepare to be canceled.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGQUIT)
	go func() {
		<-sig
		cancel()
	}()

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = Author
	app.Email = Email
	app.Usage = "A easy way to run your programs on the cloud computing environment."

	app.Metadata = map[string]interface{}{
		"context": ctx,
	}
	app.Flags = GlobalFlags
	app.Before = command.PrepareCommand
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound
	app.EnableBashCompletion = true
	app.Copyright = `roadie  Copyright (C) 2016-2017 Junpei Kawamoto <junpei.kawamoto@acm.org>

   This program comes with ABSOLUTELY NO WARRANTY.
   This is free software, and you are welcome to redistribute it
   under certain conditions.

   See https://jkawamoto.github.io/roadie/info/licenses/ for more
   information.`

	app.Run(os.Args)

}
