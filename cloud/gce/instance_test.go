//
// cloud/gce/instance_test.go
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

package gce

import (
	"testing"
)

func TestNewComputeService(t *testing.T) {

	project := "sample-project"
	region := "us-central1-c"
	machine := "n1-standard-2"
	cfg := &GcpConfig{
		Project:     project,
		Zone:        region,
		MachineType: machine,
	}

	s := NewComputeService(cfg, nil)

	if s.Config.Project != project {
		t.Error("Project name doesn't match:", s.Config.Project)
	}
	if s.Config.Zone != region {
		t.Error("Zone name doesn't match:", s.Config.Zone)
	}
	if s.Config.MachineType != machine {
		t.Error("Machine type doesn't match:", s.Config.MachineType)
	}

	if s.Logger == nil {
		t.Error("Logger is nil")
	}

}
