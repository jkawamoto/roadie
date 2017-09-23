//
// command/prepare_test.go
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

package command

import (
	"context"
	"io"
	"io/ioutil"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jkawamoto/roadie/cloud/mock"
	"github.com/jkawamoto/roadie/config"
	colorable "github.com/mattn/go-colorable"
)

// testMetadata creates metadata for testings.
func testMetadata(output io.Writer) (m *Metadata) {
	if output == nil {
		output = ioutil.Discard
	}
	m = &Metadata{
		Config:   &config.Config{},
		Context:  context.Background(),
		provider: mock.NewProvider(),
		Stdout:   colorable.NewNonColorable(output),
		Spinner:  spinner.New(spinner.CharSets[14], 100*time.Millisecond),
	}
	m.Spinner.Writer = ioutil.Discard
	return
}
