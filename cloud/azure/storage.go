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
	"github.com/jkawamoto/roadie/cloud"
	client "github.com/jkawamoto/roadie/cloud/azure/storage/client"
	"github.com/jkawamoto/roadie/cloud/azure/storage/client/storage_accounts"
	"github.com/jkawamoto/roadie/cloud/azure/storage/models"
)

const (
	// StorageAPIVersion defines API version of storage service.
	StorageAPIVersion = "2016-12-01"
)

// StorageService provides an interface for Azure's storage management service.
type StorageService struct {
	// Client of blob storage service
	blobClient storage.BlobStorageClient
	// Client of Azure resource manager
	armClient *client.StorageManagement
	// Configuration
	Config *AzureConfig
	// Logger
	Logger *log.Logger
	// Sleep time
	SleepTime time.Duration
}

// NewStorageService creates an interface of the storage service which has a
// given name and belongs to given subscription and location.
// If log.Logger logger is given, verbose mode is on and logging information will
// be written to the logger.
func NewStorageService(ctx context.Context, cfg *AzureConfig, logger *log.Logger) (s *StorageService, err error) {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}
	// Create a resource group if not exist.
	err = CreateResourceGroupIfNotExist(ctx, cfg, logger)
	if err != nil {
		return
	}

	// Create a storage account manager.
	manager := client.NewHTTPClient(strfmt.NewFormats())
	switch transport := manager.Transport.(type) {
	case *httptransport.Runtime:
		transport.DefaultAuthentication = httptransport.BearerToken(cfg.Token.AccessToken)
		manager.StorageAccounts.SetTransport(transport)
	}

	s = &StorageService{
		armClient: manager,
		Config:    cfg,
		Logger:    logger,
		SleepTime: DefaultSleepTime,
	}

	// Create an account if necessary.
	var exist bool
	exist, err = s.checkStorageAccount(ctx)
	if err != nil {
		return
	} else if !exist {
		err = s.createStorageAccount(ctx)
		if err != nil {
			return
		}
	}

	key, err := s.getStorageKey(ctx)
	if err != nil {
		return
	}
	cli, err := storage.NewBasicClient(s.Config.StorageAccount, key)
	if err != nil {
		return
	}
	s.blobClient = cli.GetBlobService()
	return

}

// TODO: The following methods should support both roadie based URL and Azure based URL.

// Upload a given stream to a given location.
func (s *StorageService) Upload(ctx context.Context, loc *url.URL, in io.Reader) (err error) {

	s.Logger.Println("Creating a blob at", loc)
	filename := strings.TrimPrefix(loc.Path, "/")
	return s.UploadWithMetadata(ctx, loc.Hostname(), filename, in, nil)

}

// UploadWithMetadata a given stream in a given container as a file named a given file name.
func (s *StorageService) UploadWithMetadata(ctx context.Context, container, filename string, in io.Reader, metadata map[string]string) (err error) {

	// Check the target container exists.
	containerRef := s.blobClient.GetContainerReference(container)
	_, err = containerRef.CreateIfNotExists(&storage.CreateContainerOptions{
		Access: storage.ContainerAccessTypeBlob,
	})
	if err != nil {
		return
	}

	s.Logger.Println("Creating append blob", filename)
	blob := containerRef.GetBlobReference(filename)
	blob.Metadata = storage.BlobMetadata(metadata)
	err = blob.CreateBlockBlob(nil)
	if err != nil {
		s.Logger.Println("Cannot create append blob", filename)
		return
	}
	err = blob.SetMetadata(nil)
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
		s.Logger.Println("Uploading data didn't finish")
	} else {
		s.Logger.Println("Finish uploading data")
	}
	return

}

// Download a file associated from a given location and write it to a given
// writer.
func (s *StorageService) Download(ctx context.Context, loc *url.URL, out io.Writer) (err error) {

	s.Logger.Println("Downloading a blob from", loc)
	filename := strings.TrimPrefix(loc.Path, "/")
	reader, err := s.blobClient.GetContainerReference(loc.Hostname()).GetBlobReference(filename).Get(nil)
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

	s.Logger.Println("Retrieving information of file in", loc)
	filename := strings.TrimPrefix(loc.Path, "/")
	blob := s.blobClient.GetContainerReference(loc.Hostname()).GetBlobReference(filename)
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

// GetMetadata retrives metadata of a given named file.
func (s *StorageService) GetMetadata(ctx context.Context, container, filename string) (metadata map[string]string, err error) {

	s.Logger.Println("Retrieving metadata of file", filename)
	blob := s.blobClient.GetContainerReference(container).GetBlobReference(filename)
	err = blob.GetMetadata(nil)
	if err != nil {
		s.Logger.Println("Get metadata of file", filename)
	}
	metadata = blob.Metadata
	return

}

// List up files matching a given prefix.
// It takes a handler; information of found files are sent to it.
func (s *StorageService) List(ctx context.Context, loc *url.URL, handler cloud.FileInfoHandler) (err error) {

	s.Logger.Println("Retrieving blobs matching", loc)
	prefix := strings.TrimPrefix(loc.Path, "/")
	res, err := s.blobClient.GetContainerReference(loc.Hostname()).ListBlobs(storage.ListBlobsParameters{
		Prefix: prefix,
	})
	if err != nil {
		switch e := err.(type) {
		case storage.AzureStorageServiceError:
			if e.StatusCode == 404 {
				s.Logger.Println("Finished retrieving blobs")
				err = nil
			}
		}
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

	s.Logger.Println("Deleting a blob in", loc)
	filename := strings.TrimPrefix(loc.Path, "/")
	err = s.blobClient.GetContainerReference(loc.Hostname()).GetBlobReference(filename).Delete(nil)
	if err != nil {
		s.Logger.Println("Cannot delete blob", filename, ":", err.Error())
	} else {
		s.Logger.Println("Deleted blob", filename)
	}
	return

}

// getFileURL returns an Azure storage URL of a file identified by the given
// container name and file name.
// Note that, this URL shouldn't use out of this package. In other packages,
// URL must starts with `roadie://`.
func (s *StorageService) getFileURL(container, filename string) string {
	return s.blobClient.GetContainerReference(container).GetBlobReference(filename).GetURL()
}

// checkStorageAccount checks existence of a given account in a given subscription.
func (s *StorageService) checkStorageAccount(ctx context.Context) (bool, error) {

	s.Logger.Println("Checking existence of storage account", s.Config.StorageAccount)
	list, err := s.armClient.StorageAccounts.StorageAccountsList(
		storage_accounts.NewStorageAccountsListParamsWithContext(ctx).
			WithAPIVersion(StorageAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID))

	if err == nil {
		for _, v := range list.Payload.Value {
			if v.Name == s.Config.StorageAccount {
				s.Logger.Println("Found storage account", s.Config.StorageAccount)
				return true, nil
			}
		}
	}

	s.Logger.Println("Cannot find storage account", s.Config.StorageAccount)
	return false, err

}

// createStorageAccount creates a given account in a given subscription.
func (s *StorageService) createStorageAccount(ctx context.Context) (err error) {

	s.Logger.Println("Creating storage account", s.Config.StorageAccount)
	created, creating, err := s.armClient.StorageAccounts.StorageAccountsCreate(
		storage_accounts.NewStorageAccountsCreateParamsWithContext(ctx).
			WithAPIVersion(StorageAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithAccountName(s.Config.StorageAccount).
			WithParameters(&models.StorageAccountCreateParameters{
				Kind:     toPtr(models.StorageAccountKindBlobStorage),
				Location: &s.Config.Location,
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

	case created != nil || creating != nil:
		var info *models.StorageAccount
		for {
			s.Logger.Println("Waiting for creating storage account", s.Config.StorageAccount)
			if info, err = s.getStorageAccountInfo(ctx); err != nil {
				break
			} else if info.Properties.ProvisioningState == ProvisioningSucceeded {
				s.Logger.Println("Created storage account", s.Config.StorageAccount)
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

	s.Logger.Println("Cannot create storage account", s.Config.StorageAccount, ":", err.Error())
	return

}

func (s *StorageService) getStorageAccountInfo(ctx context.Context) (info *models.StorageAccount, err error) {

	s.Logger.Println("Retrieving information of storage account", s.Config.StorageAccount)
	res, err := s.armClient.StorageAccounts.StorageAccountsGetProperties(
		storage_accounts.NewStorageAccountsGetPropertiesParamsWithContext(ctx).
			WithAPIVersion(StorageAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithAccountName(s.Config.StorageAccount))
	if err != nil {
		return
	}

	s.Logger.Println("Retrieved information of storage account", s.Config.StorageAccount)
	info = res.Payload
	return

}

// deleteStorageAccount deletes a given named storage account.
func (s *StorageService) deleteStorageAccount(ctx context.Context) (err error) {

	s.Logger.Println("Deleting storage account", s.Config.StorageAccount)
	deleted, deleting, err := s.armClient.StorageAccounts.StorageAccountsDelete(
		storage_accounts.NewStorageAccountsDeleteParamsWithContext(ctx).
			WithAPIVersion(StorageAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithAccountName(s.Config.StorageAccount))

	switch {
	case err != nil:
		err = NewAPIError(err)

	case deleted != nil || deleting != nil:
		var exist bool
		for {
			s.Logger.Println("Waiting for deleting storage account", s.Config.StorageAccount)
			if exist, err = s.checkStorageAccount(ctx); err != nil {
				break
			} else if !exist {
				s.Logger.Println("Deleted storage account", s.Config.StorageAccount)
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

	s.Logger.Println("Cannot delete storage account", s.Config.StorageAccount, ":", err.Error())
	return

}

// getStorageKey returns a storage access key of a given account.
func (s *StorageService) getStorageKey(ctx context.Context) (key string, err error) {

	s.Logger.Println("Retrieving access keys for account", s.Config.StorageAccount)
	keyList, err := s.armClient.StorageAccounts.StorageAccountsListKeys(
		storage_accounts.NewStorageAccountsListKeysParamsWithContext(ctx).
			WithAPIVersion(StorageAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithAccountName(s.Config.StorageAccount))

	if err != nil || len(keyList.Payload.Keys) == 0 {
		keys, err := s.generateStorageKeys(ctx)
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
func (s *StorageService) generateStorageKeys(ctx context.Context) (keys []*models.StorageAccountKey, err error) {

	s.Logger.Println("Generating access keys for account", s.Config.StorageAccount)
	res, err := s.armClient.StorageAccounts.StorageAccountsRegenerateKey(
		storage_accounts.NewStorageAccountsRegenerateKeyParamsWithContext(ctx).
			WithAPIVersion(StorageAPIVersion).
			WithSubscriptionID(s.Config.SubscriptionID).
			WithResourceGroupName(s.Config.ResourceGroupName).
			WithAccountName(s.Config.StorageAccount))

	if err != nil {
		err = NewAPIError(err)
		s.Logger.Println("Cannot generate access keys")
		return
	}

	s.Logger.Println("Generated access keys")
	keys = res.Payload.Keys
	return

}
