// +build !windows
//
// command/decorator_x.go
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

package command

import "github.com/ttacon/chalk"

// ColoredDecorator is a decorator which provides decorating strings.
var ColoredDecorator = Decorator{
	Black:   chalk.Black.Color,
	Red:     chalk.Red.Color,
	Green:   chalk.Green.Color,
	Yellow:  chalk.Yellow.Color,
	Blue:    chalk.Blue.Color,
	Magenta: chalk.Magenta.Color,
	Cyan:    chalk.Cyan.Color,
	White:   chalk.White.Color,
	Bold:    chalk.Bold.TextStyle,
}

// MonoDecorator is a decorator which does not decorate anything.
var MonoDecorator = Decorator{
	Black:   NoDetorate,
	Red:     NoDetorate,
	Green:   NoDetorate,
	Yellow:  NoDetorate,
	Blue:    NoDetorate,
	Magenta: NoDetorate,
	Cyan:    NoDetorate,
	White:   NoDetorate,
	Bold:    NoDetorate,
}
