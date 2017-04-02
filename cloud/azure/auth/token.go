//
// cloud/azure/auth/token.go
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
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// Token defines token information.
type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	ExpiresOn    string `json:"expires_on"`
	Resource     string
	Scope        string
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
}

// NewToken reads a file and returns a token in it.
func NewToken(filename string) (token *Token, err error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	token = new(Token)
	err = json.Unmarshal(data, token)
	return

}

// Save stores this token to a file.
func (t *Token) Save(filename string, perm os.FileMode) (err error) {

	data, err := json.Marshal(t)
	if err != nil {
		return
	}
	return ioutil.WriteFile(filename, data, perm)

}

// Expired returns true if this token is expired.
func (t *Token) Expired() bool {

	value, err := strconv.ParseInt(t.ExpiresOn, 10, 64)
	if err != nil {
		return true
	}
	return value < time.Now().Unix()

}

// TokenError defines errors for a token.
type TokenError struct {
	ErrorSummary     string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorCodes       []int  `json:"error_codes"`
	Timestamp        string
	TraceID          string `json:"trace_id"`
	CorrelationID    string `json:"correlation_id"`
}

// Error returns a string representing this error.
func (e *TokenError) Error() string {
	return e.ErrorDescription
}
