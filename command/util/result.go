//
// command/util/result.go
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

package util

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/jkawamoto/roadie/script"
	"github.com/ttacon/chalk"
)

// UpdateResultSection updates result section of the given script file.
func UpdateResultSection(s *script.Script, overwrite bool, warning io.Writer) {

	if s.Result == "" || overwrite {
		s.Result = script.RoadieSchemePrefix + filepath.Join(script.ResultPrefix, s.InstanceName)
	} else {
		fmt.Fprintf(
			warning,
			chalk.Red.Color("Since result section is given, all outputs will be stored in %s.\n"), s.Result)
		fmt.Fprintln(
			warning,
			chalk.Red.Color("Those buckets might not be retrieved from this program and manually downloading results is required."))
		fmt.Fprintln(
			warning,
			chalk.Red.Color("To manage outputs by this program, delete result section or set --overwrite-result-section flag."))
	}

}
