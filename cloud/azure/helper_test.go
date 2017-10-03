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
	"encoding/json"
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/jkawamoto/roadie/cloud/azure/auth"
)

func init() {
	apiAccessDebugMode = true
}

func GetTestConfig() (cfg *Config, err error) {

	logger := log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

	logger.Println("Loading a configuration file")
	cfg, err = NewConfigFromFile("./test_config.yml")
	if err != nil {
		return
	}

	logger.Println("Loading a token")
	token, err := auth.NewToken("token.json")
	if err != nil {
		return
	}
	if token.IsExpired() {
		logger.Println("Token was expired; refreshing it")
		token, err = auth.NewManualAuthorizer(cfg.TenantID, ClientID, &url.URL{}, "0000").RefreshToken(token)
		if err != nil {
			return
		}

		var fp *os.File
		fp, err = os.OpenFile("token.json", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		defer fp.Close()
		json.NewEncoder(fp).Encode(token)

	}
	cfg.Token = *token
	return

}

func TestParseRenamableURL(t *testing.T) {

	var (
		lhs string
		rhs string
	)
	lhs, rhs = parseRenamableURL("http://www.example.com/somedir/somefile:newfile")
	if lhs != "somefile" || rhs != "newfile" {
		t.Error("Parsed names are not correct:", lhs, rhs)
	}

	lhs, rhs = parseRenamableURL("http://www.example.com/somedir/somefile")
	if lhs != "somefile" || rhs != "somefile" {
		t.Error("Parsed names are not correct:", lhs, rhs)
	}

}
