//
// cloud/azure/storage.go
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
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/storage"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/jkawamoto/azure/auth"
	"github.com/jkawamoto/roadie/cloud"
	client "github.com/jkawamoto/roadie/cloud/azure/storage/client"
	"github.com/jkawamoto/roadie/cloud/azure/storage/client/storage_accounts"
	"github.com/jkawamoto/roadie/cloud/azure/storage/models"
)

const (
	// StorageAPIVersion defines API version of storage service.
	StorageAPIVersion = "2016-12-01"

	// StorageServiceResourceGroupName defines the default resource group name
	// used in Roadie.
	StorageServiceResourceGroupName = "roadie"

	// StorageServiceContainerName defines the default container name used in
	// Roadie.
	StorageServiceContainerName = "roadie"
)

// StorageService provides an interface for Azure's storage management service.
type StorageService struct {
	// Client of blob storage service
	blobClient storage.BlobStorageClient
	// Client of Azure resource manager
	armClient *client.StorageManagement
	// SubscriptID
	SubscriptionID string
	// Location
	Location string
	// Resource group name
	ResourceGroupName string
	// Logger
	Logger *log.Logger
	// Sleep time
	SleepTime time.Duration
	// Container name
	ContainerName string
}

// NewStorageService creates an interface of the storage service which has a
// given name and belongs to given subscription and location.
// If io.Writer out is given, verbose mode is on and logging information will
// be written to the writer.
func NewStorageService(ctx context.Context, token *auth.Token, subscriptionID, location, account string, out io.Writer) (s *StorageService, err error) {

	if out == nil {
		out = ioutil.Discard
	}
	logger := log.New(out, "", log.LstdFlags)

	// Create a resource group if not exist.
	err = CreateResourceGroupIfNotExist(ctx, token, subscriptionID, location, StorageServiceResourceGroupName, logger)
	if err != nil {
		return
	}

	// Create a storage account manager.
	manager := client.NewHTTPClient(strfmt.NewFormats())
	switch transport := manager.Transport.(type) {
	case *httptransport.Runtime:
		transport.DefaultAuthentication = httptransport.BearerToken(token.AccessToken)
		manager.StorageAccounts.SetTransport(transport)
	}

	s = &StorageService{
		armClient:         manager,
		SubscriptionID:    subscriptionID,
		Location:          location,
		ResourceGroupName: StorageServiceResourceGroupName,
		Logger:            logger,
		SleepTime:         30 * time.Second,
		ContainerName:     StorageServiceContainerName,
	}

	// Create an account if necessary.
	var exist bool
	exist, err = s.checkStorageAccount(ctx, account)
	if err != nil {
		return
	} else if !exist {
		err = s.createStorageAccount(ctx, account)
		if err != nil {
			return
		}
	}

	key, err := s.getStorageKey(ctx, account)
	if err != nil {
		return
	}
	cli, err := storage.NewBasicClient(account, key)
	if err != nil {
		return
	}
	s.blobClient = cli.GetBlobService()

	err = s.prepareContainer()
	return

}

// PrepareContainer creates the associated container if not exists.
func (s *StorageService) prepareContainer() (err error) {

	_, err = s.blobClient.GetContainerReference(s.ContainerName).CreateIfNotExists(
		&storage.CreateContainerOptions{
			Access: storage.ContainerAccessTypeBlob,
		})
	return

}

// Upload a given stream to a given location.
func (s *StorageService) Upload(ctx context.Context, loc *url.URL, in io.Reader) (err error) {

	filename := strings.TrimPrefix(loc.Path, "/")
	s.Logger.Println("Creating append blob", filename)

	blob := s.blobClient.GetContainerReference(s.ContainerName).GetBlobReference(filename)
	err = blob.CreateBlockBlob(nil)
	if err != nil {
		s.Logger.Println("Cannot create append blob", filename)
		return
	}
	s.Logger.Println("Created append blob", filename)

	s.Logger.Println("Uploading data")
	reader := bufio.NewReader(in)
	buf := make([]byte, 1024*1024*4)
	var size int
	for {
		size, err = reader.Read(buf)
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			break
		} else if size > 0 {
			err = blob.AppendBlock(buf[0:size], nil)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		s.Logger.Println("Uploading data didn't finishe")
	} else {
		s.Logger.Println("Finish uploading data")
	}
	return

}

// Download a file associated from a given location and write it to a given
// writer.
func (s *StorageService) Download(ctx context.Context, loc *url.URL, out io.Writer) (err error) {

	filename := strings.TrimPrefix(loc.Path, "/")
	s.Logger.Println("Downloading blob", filename)

	reader, err := s.blobClient.GetContainerReference(s.ContainerName).GetBlobReference(filename).Get(nil)
	if err != nil {
		return
	}
	defer reader.Close()

	_, err = io.Copy(out, reader)
	s.Logger.Println("Downloaded blob", filename)
	return

}

// GetFileInfo gets information of file in a given location.
func (s *StorageService) GetFileInfo(ctx context.Context, loc *url.URL) (info *cloud.FileInfo, err error) {

	filename := strings.TrimPrefix(loc.Path, "/")
	s.Logger.Println("Retrieving information of file", filename)

	blob := s.blobClient.GetContainerReference(s.ContainerName).GetBlobReference(filename)
	err = blob.GetProperties(
		&storage.GetBlobPropertiesOptions{})
	if err != nil {
		return
	}

	info = &cloud.FileInfo{
		Name:        filepath.Base(filename),
		URL:         loc,
		TimeCreated: time.Time(blob.Properties.LastModified),
		Size:        blob.Properties.ContentLength,
	}

	s.Logger.Println("Retrieved information of file", filename)
	return

}

// List up files matching a given prefix.
// It takes a handler; information of found files are sent to it.
func (s *StorageService) List(ctx context.Context, loc *url.URL, handler cloud.FileInfoHandler) (err error) {

	prefix := strings.TrimPrefix(loc.Path, "/")
	s.Logger.Println("Retrieving blobs matching ", prefix)
	res, err := s.blobClient.GetContainerReference(s.ContainerName).ListBlobs(storage.ListBlobsParameters{
		Prefix: prefix,
	})
	if err != nil {
		return
	}

	for _, v := range res.Blobs {

		u := *loc
		u.Path = path.Join(u.Path, filepath.Base(v.Name))
		err = handler(&cloud.FileInfo{
			Name:        filepath.Base(v.Name),
			URL:         &u,
			TimeCreated: time.Time(v.Properties.LastModified),
			Size:        v.Properties.ContentLength,
		})
		if err != nil {
			return
		}

	}
	return

}

// Delete a file in a given location.
func (s *StorageService) Delete(ctx context.Context, loc *url.URL) (err error) {

	filename := strings.TrimPrefix(loc.Path, "/")
	s.Logger.Println("Deleting blob", filename)
	err = s.blobClient.GetContainerReference(s.ContainerName).GetBlobReference(filename).Delete(nil)
	if err != nil {
		s.Logger.Println("Cannot delete blob", filename, ":", err.Error())
	} else {
		s.Logger.Println("Deleted blob", filename)
	}
	return

}

// checkStorageAccount checks existence of a given account in a given subscription.
func (s *StorageService) checkStorageAccount(ctx context.Context, account string) (bool, error) {

	s.Logger.Println("Checking existence of storage account", account)
	list, err := s.armClient.StorageAccounts.StorageAccountsList(
		storage_accounts.NewStorageAccountsListParamsWithContext(ctx).
			WithAPIVersion(StorageAPIVersion).
			WithSubscriptionID(s.SubscriptionID))

	if err == nil {
		for _, v := range list.Payload.Value {
			if v.Name == account {
				s.Logger.Println("Found storage account", account)
				return true, nil
			}
		}
	}

	s.Logger.Println("Cannot find storage account", account)
	return false, err

}

// createStorageAccount creates a given account in a given subscription.
func (s *StorageService) createStorageAccount(ctx context.Context, account string) (err error) {

	s.Logger.Println("Creating storage account", account)
	created, creating, err := s.armClient.StorageAccounts.StorageAccountsCreate(
		storage_accounts.NewStorageAccountsCreateParamsWithContext(ctx).
			WithAPIVersion(StorageAPIVersion).
			WithSubscriptionID(s.SubscriptionID).
			WithResourceGroupName(s.ResourceGroupName).
			WithAccountName(account).
			WithParameters(&models.StorageAccountCreateParameters{
				Kind:     toPtr(models.StorageAccountKindBlobStorage),
				Location: &s.Location,
				Sku: &models.Sku{
					Name: toPtr(models.SkuNameStandardRAGRS),
					Tier: models.SkuTierStandard,
				},
				Properties: &models.StorageAccountPropertiesCreateParameters{
					AccessTier: models.StorageAccountPropertiesCreateParametersAccessTierHot,
				},
			}))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case created != nil:
		s.Logger.Println("Created storage account", account)
		return

	case creating != nil:
		var exist bool
		for {
			s.Logger.Println("Waiting for creating storage account", account)
			if exist, err = s.checkStorageAccount(ctx, account); err != nil {
				break
			} else if exist {
				s.Logger.Println("Created storage account", account)
				return
			}
			time.Sleep(s.SleepTime)
		}

	default:
		err = fmt.Errorf("Unexpected case has occurred")

	}

	s.Logger.Println("Cannot create storage account", account, ":", err.Error())
	return

}

// deleteStorageAccount deletes a given named storage account.
func (s *StorageService) deleteStorageAccount(ctx context.Context, account string) (err error) {

	s.Logger.Println("Deleting storage account", account)
	deleted, deleting, err := s.armClient.StorageAccounts.StorageAccountsDelete(
		storage_accounts.NewStorageAccountsDeleteParamsWithContext(ctx).
			WithAPIVersion(StorageAPIVersion).
			WithSubscriptionID(s.SubscriptionID).
			WithResourceGroupName(s.ResourceGroupName).
			WithAccountName(account))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case deleted != nil:
		s.Logger.Println("Deleted storage account", account)

	case deleting != nil:
		var exist bool
		for {
			s.Logger.Println("Waiting for deleting storage account", account)
			if exist, err = s.checkStorageAccount(ctx, account); err != nil {
				break
			} else if !exist {
				s.Logger.Println("Deleted storage account", account)
				return
			}
			time.Sleep(s.SleepTime)
		}

	default:
		err = fmt.Errorf("Unexpected case has occurred")

	}

	s.Logger.Println("Cannot delete storage account", account, ":", err.Error())
	return

}

// getStorageKey returns a storage access key of a given account.
func (s *StorageService) getStorageKey(ctx context.Context, account string) (key string, err error) {

	s.Logger.Println("Retrieving access keys for account", account)
	keyList, err := s.armClient.StorageAccounts.StorageAccountsListKeys(
		storage_accounts.NewStorageAccountsListKeysParamsWithContext(ctx).
			WithAPIVersion(StorageAPIVersion).
			WithSubscriptionID(s.SubscriptionID).
			WithResourceGroupName(s.ResourceGroupName).
			WithAccountName(account))

	if err != nil || len(keyList.Payload.Keys) == 0 {
		keys, err := s.generateStorageKeys(ctx, account)
		if err != nil {
			s.Logger.Println("Cannot retrieve access keys")
			return "", err
		}
		key = keys[0].Value

	} else {
		key = keyList.Payload.Keys[0].Value

	}

	s.Logger.Println("Retrieved access keys")
	return

}

// generateStorageKeys generates access keys for a given account.
func (s *StorageService) generateStorageKeys(ctx context.Context, account string) (keys []*models.StorageAccountKey, err error) {

	s.Logger.Println("Generating access keys for account", account)
	res, err := s.armClient.StorageAccounts.StorageAccountsRegenerateKey(
		storage_accounts.NewStorageAccountsRegenerateKeyParamsWithContext(ctx).
			WithAPIVersion(StorageAPIVersion).
			WithSubscriptionID(s.SubscriptionID).
			WithResourceGroupName(s.ResourceGroupName).
			WithAccountName(account))

	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot generate access keys")
		return
	}

	s.Logger.Println("Generated access keys")
	keys = res.Payload.Keys
	return

}
