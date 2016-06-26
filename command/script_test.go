//
// command/script_test.go
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

package command

import "testing"

func TestLoadScript(t *testing.T) {

	_, err := loadScript("../test.yml", []string{"method=test"})
	if err != nil {
		t.Error(err.Error())
	}

}

func TestSetGitSource(t *testing.T) {

	s, err := loadScript("../test.yml", []string{"method=test"})
	if err != nil {
		t.Error(err.Error())
	}

	s.body.Source = ""
	s.setGitSource("https://github.com/jkawamoto/roadie-gcp.git")
	if s.body.Source != "https://github.com/jkawamoto/roadie-gcp.git" {
		t.Errorf("setGitSource doesn't work: %s", s.body.Source)
	}

}
