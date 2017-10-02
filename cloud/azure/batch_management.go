//
// cloud/azure/batch_management.go
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
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	client "github.com/jkawamoto/roadie/cloud/azure/batchmanagement/client"
	"github.com/jkawamoto/roadie/cloud/azure/batchmanagement/client/batch_account"
	"github.com/jkawamoto/roadie/cloud/azure/batchmanagement/models"
)

const (
	// BatchManagementAPIVersion defines API version of batch managmenet service.
	BatchManagementAPIVersion = "2017-01-01"
)

// BatchManagementService provides an interface for Azure's batch management
// service.
type BatchManagementService struct {
	client    *client.BatchManagement
	Config    *AzureConfig
	Logger    *log.Logger
	SleepTime time.Duration
}

// BatchAccountSet is a set of batch accounts.
type BatchAccountSet map[string]*models.BatchAccount

// NewBatchManagementService creates a new service for batch manager API.
func NewBatchManagementService(ctx context.Context, cfg *AzureConfig, logger *log.Logger) (service *BatchManagementService, err error) {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	// Create a resource group if not exist.
	err = CreateResourceGroupIfNotExist(ctx, cfg, logger)
	if err != nil {
		return
	}

	// Create a management client.
	mcli := client.NewHTTPClient(strfmt.NewFormats())
	switch transport := mcli.Transport.(type) {
	case *httptransport.Runtime:
		transport.Debug = apiAccessDebugMode
		mcli.BatchAccount.SetTransport(transport)
	}

	service = &BatchManagementService{
		client:    mcli,
		Config:    cfg,
		Logger:    logger,
		SleepTime: DefaultSleepTime,
	}
	return

}

// CreateBatchAccount creates a new batch account which has a name specified in
// the configuration given to construct this service.
func (s *BatchManagementService) CreateBatchAccount(ctx context.Context) (err error) {

	storage, err := NewStorageService(ctx, s.Config, s.Logger)
	if err != nil {
		return
	}
	storageInfo, err := storage.getStorageAccountInfo(ctx)
	if err != nil {
		return
	}

	s.Logger.Println("Creating batch account", s.Config.BatchAccount)
	created, creating, err := s.client.BatchAccount.BatchAccountCreate(
		batch_account.NewBatchAccountCreateParamsWithContext(ctx).
			WithAPIVersion(BatchManagementAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithAccountName(s.Config.BatchAccount).
			WithParameters(&models.BatchAccountCreateParameters{
				Location: &s.Config.Location,
				Properties: &models.BatchAccountBaseProperties{
					AutoStorage: &models.AutoStorageBaseProperties{
						StorageAccountID: &storageInfo.ID,
					},
					PoolAllocationMode: models.PoolAllocationModeBatchService,
				},
			}), httptransport.BearerToken(s.Config.Token.AccessToken))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case created != nil || creating != nil:
		var accounts BatchAccountSet
		for {
			s.Logger.Println("Waiting for creating a batch account")
			if accounts, err = s.BatchAccounts(ctx); err != nil {
				break
			} else if account, exists := accounts[s.Config.BatchAccount]; exists && account.Properties.ProvisioningState == ProvisioningSucceeded {
				s.Logger.Println("Created batch account", s.Config.BatchAccount)
				return
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-wait(s.SleepTime):
			}
		}

	default:
		err = fmt.Errorf("Unexpected case has occurred")

	}

	s.Logger.Println("Failed creating a batch account:", err.Error())
	return

}

// GetKey returns the primary key of the batch account. If any key does not
// exist, this function creates new keys.
func (s *BatchManagementService) GetKey(ctx context.Context) (key []byte, err error) {

	s.Logger.Println("Retrieving access keys")
	res, err := s.client.BatchAccount.BatchAccountGetKeys(
		batch_account.NewBatchAccountGetKeysParamsWithContext(ctx).
			WithAPIVersion(BatchManagementAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithAccountName(s.Config.BatchAccount), httptransport.BearerToken(s.Config.Token.AccessToken))
	if err == nil {
		s.Logger.Println("Retrieved access keys")
		return base64.StdEncoding.DecodeString(res.Payload.Primary)

	}

	s.Logger.Println("Generating a new key")
	res2, err2 := s.client.BatchAccount.BatchAccountRegenerateKey(
		batch_account.NewBatchAccountRegenerateKeyParamsWithContext(ctx).
			WithAPIVersion(BatchManagementAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithAccountName(s.Config.BatchAccount).
			WithParameters(&models.BatchAccountRegenerateKeyParameters{
				KeyName: toPtr("key"),
			}), httptransport.BearerToken(s.Config.Token.AccessToken))

	if err2 != nil {
		s.Logger.Println("Cannot generate a new key")
		return nil, err2
	}

	s.Logger.Println("Generated a new key")
	return base64.StdEncoding.DecodeString(res2.Payload.Primary)

}

// BatchAccounts retrieves a set of batch accounts defined in the subscription.
func (s *BatchManagementService) BatchAccounts(ctx context.Context) (set BatchAccountSet, err error) {

	s.Logger.Println("Retrieving batch accounts")
	res, err := s.client.BatchAccount.BatchAccountListByResourceGroup(
		batch_account.NewBatchAccountListByResourceGroupParamsWithContext(ctx).
			WithAPIVersion(BatchManagementAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName), httptransport.BearerToken(s.Config.Token.AccessToken))

	set = make(BatchAccountSet)
	for _, v := range res.Payload.Value {
		set[v.Name] = v
	}

	s.Logger.Println("Retrieved batch accounts")
	return

}

// DeleteAccount deletes the batch account of which name is given in the
// configuration.
func (s *BatchManagementService) DeleteAccount(ctx context.Context) (err error) {

	s.Logger.Println("Deleting batch acconut", s.Config.BatchAccount)
	deleted, deleting, err := s.client.BatchAccount.BatchAccountDelete(
		batch_account.NewBatchAccountDeleteParamsWithContext(ctx).
			WithAPIVersion(BatchManagementAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithAccountName(s.Config.BatchAccount), httptransport.BearerToken(s.Config.Token.AccessToken))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case deleted != nil || deleting != nil:
		var accounts BatchAccountSet
		for {
			s.Logger.Println("Waiting for deleting batch account", s.Config.BatchAccount)
			if accounts, err = s.BatchAccounts(ctx); err != nil {
				break
			} else if _, exists := accounts[s.Config.BatchAccount]; !exists {
				s.Logger.Println("Deleted batch account", s.Config.BatchAccount)
				return
			}

			select {
			case <-ctx.Done():
				err = ctx.Err()
				break
			case <-wait(s.SleepTime):
			}
		}

	default:
		err = fmt.Errorf("Unexpected case has occurred")

	}

	s.Logger.Println("Failed to delete a batch account:", err.Error())
	return

}
