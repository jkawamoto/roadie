//
// cloud/azure/resource.go
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

// This source file is associated with Azure's Resource Management API of which
// Swagger's clients are stored in `resource` directory.

package azure

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/jkawamoto/azure/auth"
	client "github.com/jkawamoto/roadie/cloud/azure/resource/client"
	"github.com/jkawamoto/roadie/cloud/azure/resource/client/resource_groups"
	"github.com/jkawamoto/roadie/cloud/azure/resource/models"
)

const (
	// ResourceAPIVersion defines API version of resource service.
	ResourceAPIVersion = "2016-09-01"
)

// ResourceService provides an interface for Azure's resource management service.
type ResourceService struct {
	client         *client.ResourceManagementClient
	token          *auth.Token
	SubscriptionID string
	Logger         *log.Logger
	SleepTime      time.Duration
}

// ResourceGroupSet is a map of which key represents a name of a resource group
// and the associated value represents a resource group object.
type ResourceGroupSet map[string]*models.ResourceGroup

// NewResourceService creates a resource service associated with a given
// subscription.
func NewResourceService(token *auth.Token, subscriptionID string, logger *log.Logger) *ResourceService {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags|log.Lshortfile)
	}

	return &ResourceService{
		client:         client.NewHTTPClient(strfmt.NewFormats()),
		token:          token,
		SubscriptionID: subscriptionID,
		Logger:         logger,
		SleepTime:      30 * time.Second,
	}

}

// CreateResourceGroup creates a resource group which has a given name in a
// given location. The created resource group will belong to the subscription
// specified whtn this resource service was created.
func (s *ResourceService) CreateResourceGroup(ctx context.Context, location, name string) (err error) {

	s.Logger.Println("Creating resource group", name)
	created, creating, err := s.client.ResourceGroups.ResourceGroupsCreateOrUpdate(
		resource_groups.NewResourceGroupsCreateOrUpdateParamsWithContext(ctx).
			WithAPIVersion(ResourceAPIVersion).
			WithSubscriptionID(s.SubscriptionID).
			WithResourceGroupName(name).
			WithParameters(&models.ResourceGroup{
				Location: &location,
			}), httptransport.BearerToken(s.token.AccessToken))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case created != nil:
		s.Logger.Println("Created resource group", name)
		return

	case creating != nil:
		var groups ResourceGroupSet
		for {
			s.Logger.Println("Waiting for creating resource group", name)
			if groups, err = s.ResourceGroups(ctx); err != nil {
				break
			} else if _, exist := groups[name]; exist {
				s.Logger.Println("Created resource group", name)
				return
			}
			time.Sleep(s.SleepTime)
		}

	default:
		err = fmt.Errorf("Unexpected case has occurred")

	}

	s.Logger.Println("Cannot create resource group", name, ":", err.Error())
	return

}

// CheckExistence checkes a given named resource group exists.
func (s *ResourceService) CheckExistence(ctx context.Context, name string) bool {

	_, err := s.client.ResourceGroups.ResourceGroupsCheckExistence(
		resource_groups.NewResourceGroupsCheckExistenceParamsWithContext(ctx).
			WithAPIVersion(ResourceAPIVersion).
			WithSubscriptionID(s.SubscriptionID).
			WithResourceGroupName(name), httptransport.BearerToken(s.token.AccessToken))
	return (err == nil)

}

// ResourceGroups retrieves a set of resource groups belonging to the
// subscrption specified when this resource service was created.
func (s *ResourceService) ResourceGroups(ctx context.Context) (groups ResourceGroupSet, err error) {

	s.Logger.Println("Retrieving resource groups")
	res, err := s.client.ResourceGroups.ResourceGroupsList(
		resource_groups.NewResourceGroupsListParamsWithContext(ctx).
			WithAPIVersion(ResourceAPIVersion).
			WithSubscriptionID(s.SubscriptionID), httptransport.BearerToken(s.token.AccessToken))
	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot retrieve resource groups:", err.Error())
		return
	}

	groups = make(ResourceGroupSet)
	for _, v := range res.Payload.Value {
		groups[v.Name] = v
	}

	s.Logger.Println("Retrieved resource groups")
	return

}

// DeleteResourceGroup deletes a given named resource group.
func (s *ResourceService) DeleteResourceGroup(ctx context.Context, name string) (err error) {

	s.Logger.Println("Deleting resource group", name)
	deleted, deleting, err := s.client.ResourceGroups.ResourceGroupsDelete(
		resource_groups.NewResourceGroupsDeleteParamsWithContext(ctx).
			WithAPIVersion(ResourceAPIVersion).
			WithSubscriptionID(s.SubscriptionID).
			WithResourceGroupName(name), httptransport.BearerToken(s.token.AccessToken))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case deleted != nil:
		s.Logger.Println("Deleted resource group", name)
		return

	case deleting != nil:
		var groups ResourceGroupSet
		for {
			s.Logger.Println("Waiting for deleting resource group", name)
			if groups, err = s.ResourceGroups(ctx); err != nil {
				break
			} else if _, exist := groups[name]; !exist {
				s.Logger.Println("Deleted resource group", name)
				return
			}
			time.Sleep(s.SleepTime)
		}

	default:
		err = fmt.Errorf("Unexpected case has occurred")

	}

	s.Logger.Println("Cannot delete resource group", name, ":", err.Error())
	return

}

// CreateResourceGroupIfNotExist checks a given named resource group exists in
// a given subscription and location. If not exists, this function creates a
// new resource group.
func CreateResourceGroupIfNotExist(ctx context.Context, token *auth.Token, subscriptionID, location, name string, logger *log.Logger) (err error) {

	service := NewResourceService(token, subscriptionID, logger)
	if service.CheckExistence(ctx, name) {
		return
	}

	return service.CreateResourceGroup(ctx, location, name)

}
