//
// cloud/azure/auth/manual_authorizer.go
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

package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// authorizeEndpoint defines the URL of the authorize endpoint.
	// It has following place holders in that order:
	// - tenant ID where the application is registered,
	// - application's client ID,
	// - redirect URI,
	// - state information.
	authorizeEndpoint = "https://login.microsoftonline.com/%v/oauth2/authorize?client_id=%v&response_type=code&redirect_uri=%v&response_mode=query&state=%v&resource=%v"

	// tokenEndpoint defines a URL to be used for obtaining authorization token.
	// Ut has the following place holder:
	// - tenant ID where the application is registered.
	tokenEndpoint = "https://login.microsoftonline.com/%v/oauth2/token"
)

// ManualAuthorizer is an authorizer which gets an authorization code manually.
type ManualAuthorizer struct {
	AuthorizeURL string
	tenantID     string
	clientID     string
	redirect     *url.URL
	state        string
}

// NewManualAuthorizer creates a new manual authorizer.
func NewManualAuthorizer(tenantID, clientID string, redirect *url.URL, state string) (a *ManualAuthorizer) {

	return &ManualAuthorizer{
		AuthorizeURL: fmt.Sprintf(
			authorizeEndpoint, tenantID, clientID, url.QueryEscape(redirect.String()),
			state, url.QueryEscape("https://management.core.windows.net/")),
		tenantID: tenantID,
		clientID: clientID,
		redirect: redirect,
		state:    state,
	}

}

// GetAuthorizeURL returns a URL where the user should open.
func (a *ManualAuthorizer) GetAuthorizeURL() string {
	return a.AuthorizeURL
}

// RequestToken requests an authorization token.
func (a *ManualAuthorizer) RequestToken(authorizationCode string) (token *Token, err error) {

	request := make(url.Values)
	request.Add("grant_type", "authorization_code")
	request.Add("client_id", a.clientID)
	request.Add("code", authorizationCode)
	request.Add("redirect_uri", a.redirect.String())
	request.Add("resource", "https://management.core.windows.net/")

	return requestToken(a.tenantID, request)

}

// RefreshToken refreshes the authorization token.
func (a *ManualAuthorizer) RefreshToken(token *Token) (newToken *Token, err error) {

	request := make(url.Values)
	request.Add("grant_type", "refresh_token")
	request.Add("client_id", a.clientID)
	request.Add("refresh_token", token.RefreshToken)
	request.Add("resource", "https://management.core.windows.net/")
	return requestToken(a.tenantID, request)

}

// requestToken requests a token.
func requestToken(tenantID string, request url.Values) (token *Token, err error) {

	res, err := http.PostForm(fmt.Sprintf(tokenEndpoint, tenantID), request)
	if err != nil {
		return
	}

	var body []byte
	if res.StatusCode != 200 {
		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var info TokenError
		if err = json.Unmarshal(body, &info); err != nil {
			return nil, err
		}
		return nil, &info

	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	token = &Token{}
	err = json.Unmarshal(body, token)
	return

}
