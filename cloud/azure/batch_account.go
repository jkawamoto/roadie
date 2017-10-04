//
// cloud/azure/batch_account.go
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
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/arm/batch"
	"github.com/Azure/go-autorest/autorest"
)

// batchAccountManager provides an interface for Azure's batch management
// service.
type batchAccountManager struct {
	Config    *Config
	Logger    *log.Logger
	SleepTime time.Duration
	client    batch.AccountClient
}

// newBatchAccountManager creates a new batch account manager.
func newBatchAccountManager(ctx context.Context, cfg *Config, logger *log.Logger) (manager *batchAccountManager, err error) {

	cli := batch.NewAccountClient(cfg.SubscriptionID)
	cli.Authorizer = autorest.NewBearerAuthorizer(&cfg.Token)

	manager = &batchAccountManager{
		client:    cli,
		Config:    cfg,
		Logger:    logger,
		SleepTime: DefaultSleepTime,
	}
	return

}

// create creates a new batch account which has a name specified in
// the configuration given to construct this service.
func (s *batchAccountManager) create(ctx context.Context, storageID string) (err error) {

	s.Logger.Printf("Creating batch account %q", s.Config.AccountName)
	resCh, errCh := s.client.Create(s.Config.AccountName, s.Config.AccountName, batch.AccountCreateParameters{
		Location: &s.Config.Location,
		AccountCreateProperties: &batch.AccountCreateProperties{
			AutoStorage: &batch.AutoStorageBaseProperties{
				StorageAccountID: &storageID,
			},
			PoolAllocationMode: batch.BatchService,
		},
	}, ctx.Done())

	select {
	case <-resCh:
	case err = <-errCh:
	case <-ctx.Done():
		err = ctx.Err()
	}

	if err != nil {
		s.Logger.Printf("Failed creating a batch account: %v", err)
	} else {
		s.Logger.Printf("Created batch account %q", s.Config.AccountName)
	}
	return

}

// getKey returns the primary key of the batch account. If any key does not
// exist, this function creates new keys.
func (s *batchAccountManager) getKey(ctx context.Context) (key []byte, err error) {

	s.Logger.Println("Retrieving access keys")

	keys, err := s.client.GetKeys(s.Config.AccountName, s.Config.AccountName)
	if err == nil && keys.Primary != nil {
		s.Logger.Println("Retrieved access keys")
		return base64.StdEncoding.DecodeString(*keys.Primary)
	}

	s.Logger.Println("Generating a new key")
	keys, err = s.client.RegenerateKey(s.Config.AccountName, s.Config.AccountName, batch.AccountRegenerateKeyParameters{
		KeyName: batch.Primary,
	})
	if err != nil {
		return
	}

	s.Logger.Println("Generated a new key")
	return base64.StdEncoding.DecodeString(*keys.Primary)

}

// accounts retrieves a set of batch accounts defined in the subscription.
func (s *batchAccountManager) accounts(ctx context.Context) (res []batch.Account, err error) {

	s.Logger.Println("Retrieving batch accounts")
	list, err := s.client.List()
	if err != nil {
		return
	}
	res = *list.Value
	s.Logger.Println("Retrieved batch accounts")
	return

}

// delete deletes the batch account of which name is given in the
// configuration.
func (s *batchAccountManager) delete(ctx context.Context) (err error) {

	s.Logger.Printf("Deleting batch acconut %q", s.Config.AccountName)
	resCh, errCh := s.client.Delete(s.Config.AccountName, s.Config.AccountName, ctx.Done())
	select {
	case <-resCh:
	case err = <-errCh:
	case <-ctx.Done():
		err = ctx.Err()
	}

	if err != nil {
		s.Logger.Printf("Failed to delete batch account %q: %v", s.Config.AccountName, err)
	} else {
		s.Logger.Printf("Deleted batch account %q", s.Config.AccountName)
	}
	return

}
