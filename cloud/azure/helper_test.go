//
// cloud/azure/helper_test.go
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
	"log"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/jkawamoto/roadie/cloud/azure/auth"
)

func GetTestConfig() (cfg *AzureConfig, err error) {

	logger := log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

	logger.Println("Loading a configuration file")
	cfg, err = NewAzureConfigFromFile("./test_config.yml")
	if err != nil {
		return
	}

	logger.Println("Loading a token")
	token, err := auth.NewToken("token.json")
	if err != nil {
		return
	}
	if token.Expired() {
		logger.Println("Token was expired; refreshing it")
		token, err = auth.NewManualAuthorizer(cfg.TenantID, cfg.ClientID, &url.URL{}, "0000").RefreshToken(token)
		if err != nil {
			return
		}
	}
	cfg.Token = *token
	return

}

func TestWait(t *testing.T) {

	select {
	case <-wait(1 * time.Minute):
		t.Fatal("Returned waiting 1min function first.")
	case <-wait(1 * time.Second):
		t.Log("Waiting 1 second returns first.")
	}

}
