//
// cloud/azure/provider.go
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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/jkawamoto/roadie/cloud"
	"github.com/jkawamoto/roadie/cloud/azure/auth"
)

// Provider provides information of the service provider for Azure.
type Provider struct {
	Config *Config
	Logger *log.Logger
}

// NewProvider creates a new provider for Azure service.
func NewProvider(ctx context.Context, cfg *Config, logger *log.Logger, forceAuth bool) (provider *Provider, err error) {

	if cfg.Token.AccessToken == "" || forceAuth {

		logger.Println("Access token is not given.")
		var token *adal.Token
		token, err = auth.AuthorizeDeviceCode(ctx, ClientID, os.Stdout)
		if err != nil {
			return
		}
		cfg.Token = *token

	} else if cfg.Token.IsExpired() {

		logger.Println("Access token is expired; refreshing now.")
		authorizer := auth.NewManualAuthorizer(cfg.TenantID, ClientID, nil, fmt.Sprintf("%v", time.Now().Unix()))
		var token *adal.Token
		token, err = authorizer.RefreshToken(&cfg.Token)
		if err != nil {
			logger.Printf("Cannot refresh a token: %v\nObtaining a new token...", err)
			token, err = auth.AuthorizeDeviceCode(ctx, ClientID, os.Stdout)
			if err != nil {
				return
			}
		}
		cfg.Token = *token

	}

	provider = &Provider{
		Config: cfg,
		Logger: logger,
	}
	return

}

// InstanceManager returns an instance manager interface.
func (p *Provider) InstanceManager(ctx context.Context) (cloud.InstanceManager, error) {
	return NewInstanceManager(ctx, p.Config, p.Logger)
}

// QueueManager returns a queue manager interface.
func (p *Provider) QueueManager(ctx context.Context) (cloud.QueueManager, error) {
	return NewQueueManager(ctx, p.Config, p.Logger)
}

// StorageManager returns a storage manager interface.
func (p *Provider) StorageManager(ctx context.Context) (cloud.StorageManager, error) {
	return NewStorageService(ctx, p.Config, p.Logger)
}

// LogManager returns a log manager interface.
func (p *Provider) LogManager(ctx context.Context) (cloud.LogManager, error) {
	return NewLogManager(ctx, p.Config, p.Logger)
}

// ResourceManager returns a resource manager interface.
func (p *Provider) ResourceManager(ctx context.Context) (cloud.ResourceManager, error) {
	return NewResourceManager(p.Config, p.Logger), nil
}
