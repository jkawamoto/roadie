// +build windows
//
// chalkd/wrapper_windows.go
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

type DummyColor struct {
}

func (d DummyColor) Color(val string) string {
	return val
}

var (
	//Colors
	Black      = DummyColor{}
	Red        = DummyColor{}
	Green      = DummyColor{}
	Yellow     = DummyColor{}
	Blue       = DummyColor{}
	Magenta    = DummyColor{}
	Cyan       = DummyColor{}
	White      = DummyColor{}
	ResetColor = DummyColor{}

	// Text Styles
	Bold          = baseChalk.TextStyle{}
	Dim           = baseChalk.TextStyle{}
	Italic        = baseChalk.TextStyle{}
	Underline     = baseChalk.TextStyle{}
	Inverse       = baseChalk.TextStyle{}
	Hidden        = baseChalk.TextStyle{}
	Strikethrough = baseChalk.TextStyle{}

	Reset = baseChalk.Reset
)
