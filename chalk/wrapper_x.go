// +build !windows
//
// chalk/wrapper_x.go
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

package chalk

import (
	baseChalk "github.com/ttacon/chalk"
)

var (
	//Colors
	Black      = baseChalk.Black
	Red        = baseChalk.Red
	Green      = baseChalk.Green
	Yellow     = baseChalk.Yellow
	Blue       = baseChalk.Blue
	Magenta    = baseChalk.Magenta
	Cyan       = baseChalk.Cyan
	White      = baseChalk.White
	ResetColor = baseChalk.ResetColor

	// Text Styles
	Bold          = baseChalk.Bold
	Dim           = baseChalk.Dim
	Italic        = baseChalk.Italic
	Underline     = baseChalk.Underline
	Inverse       = baseChalk.Inverse
	Hidden        = baseChalk.Hidden
	Strikethrough = baseChalk.Strikethrough

	Reset = baseChalk.Reset
)
