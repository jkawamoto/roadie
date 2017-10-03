//
// cloud/azure/subscriptions.go
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

// This source file is associated with Azure's Subscriptions API of which
// Swagger's clients are stored in `subscriptions` directory.

package azure

import (
	"context"

	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/go-openapi/strfmt"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/cloud/azure/subscriptions/client"
	"github.com/jkawamoto/roadie/cloud/azure/subscriptions/client/subscriptions"
)

const (
	// SubscriptionsAPIVersion defines API version of sbscriptions service.
	SubscriptionsAPIVersion = "2016-06-01"
)

// Locations gets list of locations in a given subscription.
func Locations(ctx context.Context, token *adal.Token, subscriptionID string) (regions []cloud.Region, err error) {

	cli := client.NewHTTPClient(strfmt.NewFormats())
	res, err := cli.Subscriptions.SubscriptionsListLocations(
		subscriptions.NewSubscriptionsListLocationsParamsWithContext(ctx).
			WithAPIVersion(SubscriptionsAPIVersion).
			WithSubscriptionID(subscriptionID), httptransport.BearerToken(token.AccessToken))
	if err != nil {
		return nil, NewAPIError(err)
	}

	regions = make([]cloud.Region, len(res.Payload.Value))
	for i, v := range res.Payload.Value {
		regions[i] = cloud.Region{
			Name:   v.Name,
			Status: v.DisplayName,
		}
	}

	return

}
