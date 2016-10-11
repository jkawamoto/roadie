//
// command/resource/woker.go
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

package resource

import (
	"bytes"
	"html/template"
)

// WorkerStartupOpt defines variables for startup script of queue worker instance.
type WorkerStartupOpt struct {
	// Project ID.
	ProjectID string
	// Queue name.
	Name string
	// Instance name.
	InstanceName string
	// Version of Roadie queue manager. The format is x.y.z.
	Version string
}

// WorkerStartup returns a startup script for worker instances.
func WorkerStartup(opt *WorkerStartupOpt) (res string, err error) {

	startup, err := Asset("assets/worker.sh")
	if err != nil {
		return
	}

	temp, err := template.New("startup").Parse(string(startup))
	if err != nil {
		return
	}
	buf := &bytes.Buffer{}

	err = temp.ExecuteTemplate(buf, "startup", opt)
	if err != nil {
		return
	}

	res = buf.String()
	return

}
