//
// cloud/azure/auth/authorizer.go
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
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
)

import "github.com/Azure/go-autorest/autorest/adal"

// Authorizer provides methods to get an authorization code and refresh it.
type Authorizer struct {
	authorizer *ManualAuthorizer
	receiver   receiver
	listener   net.Listener
}

// authorizationCode represents an authorization code.
type authorizationCode struct {
	AdminConsent string
	Code         string
	SessionState string
	State        string
}

// failedResponse represents an error response.
type failedResponse struct {
	ErrorCode        string
	ErrorDescription string
	State            string
}

// Error returns a string explaining this error.
func (e *failedResponse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode, e.ErrorDescription)
}

// NewAuthorizer creates a new authorizer.
func NewAuthorizer(tenantID, clientID string, redirect *url.URL, state string) (a *Authorizer, err error) {

	listener, err := net.Listen("tcp", redirect.Host)
	if err != nil {
		return
	}

	a = &Authorizer{
		authorizer: NewManualAuthorizer(tenantID, clientID, redirect, state),
		listener:   listener,
		receiver: receiver{
			ResultChannel: make(chan *authorizationCode, 1),
			ErrorChannel:  make(chan *failedResponse, 1),
		},
	}

	go http.Serve(a.listener, &a.receiver)
	return

}

// GetAuthorizeURL returns a URL which users should access.
func (a *Authorizer) GetAuthorizeURL() string {
	return a.authorizer.GetAuthorizeURL()
}

// WaitResponse waits the user acceseses an authorization URL and grants
// access.
func (a *Authorizer) WaitResponse(ctx context.Context) (*adal.Token, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	case code := <-a.receiver.ResultChannel:
		if code.State != a.authorizer.state {
			return nil, fmt.Errorf("State information is not matched: %s", code.State)
		}
		log.Println("Received an authorization code, then requesting a token...")
		return a.authorizer.RequestToken(code.Code)

	case err := <-a.receiver.ErrorChannel:
		return nil, err
	}

}

// RefreshToken refreshes a token.
func (a *Authorizer) RefreshToken(token *adal.Token) (newToken *adal.Token, err error) {
	return a.authorizer.RefreshToken(token)
}

// Close authorization process. It must be called.
func (a *Authorizer) Close() error {
	return a.listener.Close()
}

// receiver implements http.Handler and provides a service which receives
// an authorization code.
type receiver struct {
	ResultChannel chan *authorizationCode
	ErrorChannel  chan *failedResponse
}

// ServeHTTP receives a http.Request and writes a response to a given response
// writer.
func (r *receiver) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	queries := request.URL.Query()
	if errCode := queries.Get("error"); errCode != "" {

		fmt.Fprintln(response, "Done")
		r.ErrorChannel <- &failedResponse{
			ErrorCode:        errCode,
			ErrorDescription: queries.Get("error_description"),
			State:            queries.Get("state"),
		}

	} else if code := queries.Get("code"); code == "" {

		fmt.Fprintln(response, "Invarid")
		r.ErrorChannel <- &failedResponse{}

	} else {

		fmt.Fprintln(response, "Approved")
		r.ResultChannel <- &authorizationCode{
			AdminConsent: queries.Get("admin_consent"),
			Code:         queries.Get("code"),
			SessionState: queries.Get("session_state"),
			State:        queries.Get("state"),
		}
	}

}
