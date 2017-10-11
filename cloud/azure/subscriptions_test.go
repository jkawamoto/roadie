// +build remote
//
// cloud/azure/subscriptions_test.go
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
	"context"
	"testing"
)

func TestLocations(t *testing.T) {
	t.SkipNow()

	var err error
	cfg, err := GetTestConfig()
	if err != nil {
		t.Skip("Test configuration is not supplied, skip tests.")
	}

	ctx := context.Background()
	regions, err := Locations(ctx, &cfg.Token, cfg.SubscriptionID)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(regions) == 0 {
		t.Error("There are no available locations")
	}
	for _, v := range regions {
		t.Log(v.Name)
	}

}
