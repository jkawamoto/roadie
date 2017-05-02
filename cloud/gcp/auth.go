//
// cloud/gcp/auth.go
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

package gcp

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const (
	// authorizeEndpoint defines a URL to obtain an authorization code.
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	// tokenEndpoint defines a URL to obtain a token and refresh a token.
	tokenEndpoint = "https://www.googleapis.com/oauth2/v4/token"
)

var (
	// CodeVerifierChars defines a set of characters used to generate a code verifier.
	CodeVerifierChars []byte
)

func init() {
	for b := byte('a'); b <= byte('z'); b++ {
		CodeVerifierChars = append(CodeVerifierChars, b)
	}
	for b := byte('A'); b <= byte('Z'); b++ {
		CodeVerifierChars = append(CodeVerifierChars, b)
	}
	for b := byte('0'); b <= byte('9'); b++ {
		CodeVerifierChars = append(CodeVerifierChars, b)
	}
	CodeVerifierChars = append(CodeVerifierChars, byte('-'), byte('.'), byte('_'), byte('~'))
}

// authorizationCode defines returned values while getting an authorization code.
type authorizationCode struct {
	Code  string
	State string
}

// codeReciever is a local HTTP server to recieve an authorization code.
type codeReciever struct {
	Result chan *authorizationCode
	Error  chan string
}

// newCodeReciever create a new codeReciever.
func newCodeReciever() *codeReciever {
	return &codeReciever{
		Result: make(chan *authorizationCode, 1),
		Error:  make(chan string, 1),
	}
}

func (r *codeReciever) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	queries := req.URL.Query()
	if errCode := queries.Get("error"); errCode != "" {
		fmt.Fprintln(res, "Done")
		r.Error <- errCode

	} else if code := queries.Get("code"); code == "" {
		fmt.Fprintln(res, "Invarid")
		r.Error <- ""

	} else {
		fmt.Fprintln(res, "Approved")
		r.Result <- &authorizationCode{
			Code:  queries.Get("code"),
			State: queries.Get("state"),
		}

	}

}

// errorMessage defines a structure of an error message of getting an access token.
type errorMessage struct {
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

// NewAuthorizationConfig creates a new configuration for authorization.
func NewAuthorizationConfig(port int) *oauth2.Config {

	redirect := fmt.Sprintf("http://127.0.0.1:%v", port)
	return &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		RedirectURL: redirect,
		Scopes:      []string{gcpScope},
	}
}

// RequestToken requests a new authorization token.
func RequestToken(ctx context.Context, output io.Writer) (token *oauth2.Token, err error) {

	codeVerifier, err := GenerateCodeVerifier()
	if err != nil {
		return
	}
	codeChallenge := string(codeVerifier)

	port := 18029
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", port))
	if err != nil {
		return
	}
	defer listener.Close()

	cfg := NewAuthorizationConfig(port)
	state := fmt.Sprintf("%v", time.Now().Unix())
	endpoint := cfg.AuthCodeURL(state, oauth2.SetAuthURLParam("code_challenge_method", "plain"), oauth2.SetAuthURLParam("code_challenge", codeChallenge))
	// TODO: Print message to open the following link.
	fmt.Fprintln(output, endpoint)

	receiver := newCodeReciever()
	go http.Serve(listener, receiver)

	var code *authorizationCode
	select {
	case code = <-receiver.Result:
		fmt.Println(code)
	case errMsg := <-receiver.Error:
		err = fmt.Errorf("Failed authorization: %v", errMsg)
		return
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	cfg.Endpoint.TokenURL = fmt.Sprintf("%v?code_verifier=%v", cfg.Endpoint.TokenURL, string(codeVerifier))
	return cfg.Exchange(ctx, code.Code)

}

// GenerateCodeVerifier generates a plain code verifier.
func GenerateCodeVerifier() (codeVerifier []byte, err error) {

	max := big.NewInt(int64(len(CodeVerifierChars)))
	var n *big.Int
	for i := 0; i != 128; i++ {
		n, err = rand.Int(rand.Reader, max)
		if err != nil {
			return
		}
		codeVerifier = append(codeVerifier, CodeVerifierChars[n.Int64()])
	}
	return

}
