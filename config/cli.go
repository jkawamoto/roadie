//
// config/clid.go
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

package config

import (
	"context"

	"github.com/urfave/cli"
)

// FromCliContext returns a config object from a context of cli.
func FromCliContext(c *cli.Context) (conf *Config) {
	ctx, _ := c.App.Metadata["context"].(context.Context)
	conf, _ = FromContext(ctx)
	return conf
}
