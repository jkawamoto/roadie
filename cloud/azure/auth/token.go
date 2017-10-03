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

	"github.com/Azure/go-autorest/autorest/adal"
)

// NewToken reads a file and returns a token in it.
func NewToken(filename string) (token *adal.Token, err error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	token = new(adal.Token)
	err = json.Unmarshal(data, token)
	return

}

// TokenError defines errors for a token.
type TokenError struct {
	ErrorSummary     string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorCodes       []int  `json:"error_codes"`
	Timestamp        string `json:"timestamp"`
	TraceID          string `json:"trace_id"`
	CorrelationID    string `json:"correlation_id"`
}

// Error returns a string representing this error.
func (e *TokenError) Error() string {
	return e.ErrorDescription
}
