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
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"path"
	"strings"
	"time"

	arm_storage "github.com/Azure/azure-sdk-for-go/arm/storage"
	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/jkawamoto/roadie/cloud"
)

const (
	// DefaultAccessPolicyExpiryTime defines a default expiry time.
	DefaultAccessPolicyExpiryTime = 30 * 24 * time.Hour
)

// storageAccountManager provides functions to manage a storage account.
type storageAccountManager struct {
	// Configuration
	Config *Config
	// Logger
	Logger *log.Logger
	// client for Azure's account resource manager.
	client arm_storage.AccountsClient
}

// newStorageAccountManager creates a new account manager.
func newStorageAccountManager(cfg *Config, logger *log.Logger) *storageAccountManager {

	cli := arm_storage.NewAccountsClient(cfg.SubscriptionID)
	cli.Authorizer = autorest.NewBearerAuthorizer(&cfg.Token)
	return &storageAccountManager{
		Config: cfg,
		Logger: logger,
		client: cli,
	}

}

// createIfNotExists checks the associated storage account exists. If not, creates
// it.
func (s *storageAccountManager) createIfNotExists(ctx context.Context) (err error) {

	s.Logger.Printf("Checking storage account %q exists", s.Config.StorageAccount)
	accounts, err := s.client.List()
	if err != nil {
		return
	}

	var exist bool
	for _, a := range *accounts.Value {
		if *a.Name == s.Config.StorageAccount {
			exist = true
			break
		}
	}
	if !exist {
		s.Logger.Printf("Storage account %q doesn't exist and creating it", s.Config.StorageAccount)
		resCh, errCh := s.client.Create(s.Config.ResourceGroupName, s.Config.StorageAccount, arm_storage.AccountCreateParameters{
			Sku: &arm_storage.Sku{
				Name: arm_storage.StandardRAGRS,
				Tier: arm_storage.Standard,
			},
			Kind:     arm_storage.BlobStorage,
			Location: &s.Config.Location,
			AccountPropertiesCreateParameters: &arm_storage.AccountPropertiesCreateParameters{
				AccessTier: arm_storage.Hot,
			},
		}, ctx.Done())

		select {
		case _ = <-resCh:
		case err = <-errCh:
		case <-ctx.Done():
			err = ctx.Err()
		}

	}
	return

}

// getStorageAccountInfo retrieves information of the associated storage account.
func (s *storageAccountManager) getStorageAccountInfo() (arm_storage.Account, error) {

	return s.client.GetProperties(s.Config.ResourceGroupName, s.Config.StorageAccount)

}

// getStorageKey returns a storage access key. If the associated storage account
// doesn't have any access keys, this function will generate it.
func (s *storageAccountManager) getStorageKey(ctx context.Context) (key string, err error) {

	s.Logger.Printf("Obtaining an access key for storage %q", s.Config.StorageAccount)
	res, err := s.client.ListKeys(s.Config.ResourceGroupName, s.Config.StorageAccount)
	if err != nil {
		return
	}
	for len(*res.Keys) == 0 {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
		}

		s.Logger.Printf("Access keys for %q don't exist and generating it", s.Config.StorageAccount)
		res, err = s.client.RegenerateKey(s.Config.ResourceGroupName, s.Config.StorageAccount, arm_storage.AccountRegenerateKeyParameters{
			KeyName: toPtr("key"),
		})
		if err != nil {
			return
		}
	}
	key = *(*res.Keys)[0].Value
	return

}

// delete deletes the associated storage account.
func (s *storageAccountManager) delete() (err error) {

	s.Logger.Printf("Deleting storage account %q", s.Config.StorageAccount)
	_, err = s.client.Delete(s.Config.ResourceGroupName, s.Config.StorageAccount)
	return

}

// StorageService provides an interface for Azure's storage management service.
type StorageService struct {
	// Client of blob storage service
	blobClient storage.BlobStorageClient
	// Configuration
	Config *Config
	// Logger
	Logger *log.Logger
	// AccessPolicyExpiryTime defines the expiry time for accessing containers.
	AccessPolicyExpiryTime time.Duration
}

// NewStorageService creates an interface of the storage service which has a
// given name and belongs to given subscription and location.
// If log.Logger logger is given, verbose mode is on and logging information will
// be written to the logger.
func NewStorageService(ctx context.Context, cfg *Config, logger *log.Logger) (s *StorageService, err error) {

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}
	// Create a resource group if not exist.
	err = CreateResourceGroupIfNotExist(ctx, cfg, logger)
	if err != nil {
		return
	}

	// Create a storage account manager and prepare the account.
	manager := newStorageAccountManager(cfg, logger)
	err = manager.createIfNotExists(ctx)
	if err != nil {
		return
	}
	key, err := manager.getStorageKey(ctx)
	if err != nil {
		return
	}

	// Create a blob storage client.
	cli, err := storage.NewBasicClient(cfg.StorageAccount, key)
	if err != nil {
		return
	}

	s = &StorageService{
		blobClient:             cli.GetBlobService(),
		Config:                 cfg,
		Logger:                 logger,
		AccessPolicyExpiryTime: DefaultAccessPolicyExpiryTime,
	}
	return

}

// TODO: The following methods should support both roadie based URL and Azure based URL.

// Upload a given stream to a given location.
func (s *StorageService) Upload(ctx context.Context, loc *url.URL, in io.Reader) (err error) {

	s.Logger.Println("Creating a blob at", loc)
	filename := strings.TrimPrefix(loc.Path, "/")
	return s.upload(ctx, loc.Hostname(), filename, in, nil, nil)

}

// upload a given stream in a given container as a file named a given file name.
func (s *StorageService) upload(
	ctx context.Context, container, filename string, in io.Reader, props *storage.BlobProperties, metadata storage.BlobMetadata) (err error) {

	// Check the target container exists.
	containerRef := s.blobClient.GetContainerReference(container)
	created, err := containerRef.CreateIfNotExists(&storage.CreateContainerOptions{
		Access: storage.ContainerAccessTypeBlob,
	})
	if err != nil {
		return
	}
	if created {
		err = containerRef.SetPermissions(storage.ContainerPermissions{
			AccessType: storage.ContainerAccessTypeBlob,
			AccessPolicies: []storage.ContainerAccessPolicy{
				storage.ContainerAccessPolicy{
					ID:         "full-access",
					StartTime:  time.Now(),
					ExpiryTime: time.Now().Add(s.AccessPolicyExpiryTime),
					CanRead:    true,
					CanWrite:   true,
					CanDelete:  true,
				},
			},
		}, nil)
		if err != nil {
			return
		}
	}

	s.Logger.Printf("Checking the uploading file %v exists\n", filename)
	exists, err := containerRef.GetBlobReference(filename).DeleteIfExists(nil)
	if err != nil {
		s.Logger.Println("Cannot check the existence of the uploading file")
		return
	} else if exists {
		s.Logger.Println("Old file has been deleted")
	}

	s.Logger.Println("Creating blob", filename)
	blob := containerRef.GetBlobReference(filename)
	err = blob.PutAppendBlob(nil)
	if err != nil {
		s.Logger.Println("Cannot create the blob", filename)
		return
	}
	if props != nil {
		blob.Properties = *props
		err = blob.SetProperties(nil)
		if err != nil {
			s.Logger.Println("Cannot set properties to blob", filename)
			return
		}
	}
	if metadata != nil {
		blob.Metadata = metadata
		err = blob.SetMetadata(nil)
		if err != nil {
			s.Logger.Println("Cannot set metadata to blob", filename)
			return
		}
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
		Name:        path.Base(filename),
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
		u.Path = path.Join(u.Path, path.Base(v.Name))
		err = handler(&cloud.FileInfo{
			Name:        path.Base(v.Name),
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

// deleteAccount deletes the associated storage account.
func (s *StorageService) deleteAccount(ctx context.Context) (err error) {

	manager := newStorageAccountManager(s.Config, s.Logger)
	return manager.delete()

}
