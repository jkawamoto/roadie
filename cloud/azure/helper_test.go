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
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/jkawamoto/azure/auth"
)

type TestConfig struct {
	SubscriptionID string `json:"subscription_id"`
	ClientID       string `json:"client_id"`
	Token          *auth.Token
}

func GetTestConfig() (cfg *TestConfig, err error) {

	cfg = new(TestConfig)
	data, err := ioutil.ReadFile("./test_config.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return
	}

	redirect, err := url.Parse("http://localhost:53612/")
	state := "12346"
	if err != nil {
		return
	}

	token, err := auth.NewToken("token.json")
	if err != nil {

		var authorizer *auth.Authorizer
		authorizer, err = auth.NewAuthorizer("common", cfg.ClientID, redirect, state)
		if err != nil {
			return
		}
		defer authorizer.Close()

		fmt.Println(authorizer.GetAuthorizeURL())

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		token, err = authorizer.WaitResponse(ctx)
		if err != nil {
			return
		}

		if err = token.Save("token.json", 0644); err != nil {
			return
		}

	} else if token.Expired() {

		authorizer := auth.NewManualAuthorizer("common", cfg.ClientID, redirect, state)
		token, err = authorizer.RefreshToken(token)
		if err != nil {
			return
		}

		if err = token.Save("token.json", 0644); err != nil {
			return
		}

	}

	cfg.Token = token
	return

}
