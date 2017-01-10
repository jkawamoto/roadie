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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jkawamoto/roadie/config"
	"github.com/urfave/cli"
)

func main() {

	// TODO: Construct a context at here and watch signals to cancel it when interrupted.

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = Author
	app.Email = Email
	app.Usage = "A easy way to run your programs on the cloud computing environment."

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound
	app.EnableBashCompletion = true
	app.Copyright = `roadie  Copyright (C) 2016  Junpei Kawamoto
This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions.

See https://jkawamoto.github.io/roadie/info/licenses/ for more
information.
`

	conf, err := config.NewConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Prepare to be canceled.
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGQUIT)
	go func() {
		<-sig
		cancel()
	}()

	// Set up metadata.
	app.Metadata = map[string]interface{}{
		"config":  conf,
		"context": ctx,
	}

	app.Run(os.Args)
}
