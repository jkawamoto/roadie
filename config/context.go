//
// config/context.go
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
	"fmt"
)

// Key type to attach and obtaine Config from contexts.
type key int

// ConfigKey is the key used store config into a context.
const configKey key = 0

// ErrNoConfig defines an error which be raised when a given context doesn't have
// any Config.
var ErrNoConfig = fmt.Errorf("Given context doesn't have a config")

// NewContext returns a context a given Config is attached to.
func NewContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, configKey, cfg)
}

// FromContext returns a Config from a given context.
func FromContext(ctx context.Context) (cfg *Config, err error) {
	cfg, ok := ctx.Value(configKey).(*Config)
	if !ok {
		err = ErrNoConfig
	}
	return
}
