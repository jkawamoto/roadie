//
// command/util/gcs_test.go
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

package util

import (
	"os"
	"testing"

	"golang.org/x/net/context"
)

func TestNewStorage(t *testing.T) {

	id := os.Getenv("PROJECT_ID")
	if id == "" {
		t.Log("Skip this test because no project id is given.")
		return
	}

	_, err := NewStorage(context.Background(), id, id)
	if err != nil {
		t.Error(err.Error())
	}

}

func TestUpload(t *testing.T) {

	// s, err := NewStorage("jkawamoto-ppls", "jkawamoto-ppls")
	// if err != nil {
	// 	t.Error(err.Error())
	// }
	//
	// location, err := url.Parse("gs://jkawamoto-ppls/.roadie/gcs_test.go")
	// if err != nil {
	// 	t.Error(err.Error())
	// }

	// if err := s.Upload("./gcs_test.go", location); err != nil {
	// 	t.Error(err.Error())
	// }

}
