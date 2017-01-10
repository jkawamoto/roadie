//
// command/data_test.go
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

package command

import (
	"context"
	"testing"
)

// TestCmdDataPut checks if wrong patterns are given, cmdDataPut returns error,
// and if empty pattern is given, it do nothing.
func TestCmdDataPut(t *testing.T) {

	// Test for wrong pattern.
	if err := cmdDataPut(context.Background(), "[b-a", ""); err == nil {
		t.Error("Give a wrong pattern but no errors occur.")
	} else {
		t.Logf("Wrong patter makes an error: %s", err.Error())
	}

	// Test for empty pattern.
	if err := cmdDataPut(context.Background(), "", ""); err != nil {
		t.Errorf(err.Error())
	}

}
