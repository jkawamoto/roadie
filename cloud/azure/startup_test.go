//
// cloud/azure/startup_test.go
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

package azure

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/jkawamoto/roadie/script"
)

func TestStartup(t *testing.T) {

	cfg := NewConfig()
	task := &script.Script{
		Name: "test-instance",
	}

	res, err := StartupScript(cfg, task)
	if err != nil {
		t.Fatalf("StartupScript returns an error: %v", err)
	}
	data, err := base64.StdEncoding.DecodeString(res)
	if err != nil {
		t.Fatalf("DecodeString returns an error: %v", err)
	}

	startup := string(data)
	if strings.Contains(startup, "&lt;") {
		t.Errorf("URLs are encoded: %v", startup)
	}
	if !strings.Contains(startup, DefaultOSPublisherName) {
		t.Errorf("doesn't have any configuration: %v", startup)
	}
	if !strings.Contains(startup, "name: test-instance") {
		t.Errorf("doesn't have a script: %v", startup)
	}

}
