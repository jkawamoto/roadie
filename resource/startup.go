//
// resource/startup.go
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
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package resource

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/ttacon/chalk"
)

// StartupOpt defines variables used in startup template.
type StartupOpt struct {
	// Name of container.
	Name string
	// Body of script file.
	Script string
	// Options
	Options string
	// Container image.
	Image string
	// Number of retry.
	Retry int64
}

// Startup constructs a startup script by given options.
func Startup(opt *StartupOpt) (res string, err error) {

	startup, err := Asset("assets/startup.sh")
	if err != nil {
		fmt.Println(chalk.Red.Color("Startup script was not found."))
		return
	}

	buf := &bytes.Buffer{}
	temp, err := template.New("startup").Parse(string(startup))
	if err != nil {
		return
	}
	if err = temp.ExecuteTemplate(buf, "startup", opt); err != nil {
		return
	}

	res = buf.String()
	return

}
