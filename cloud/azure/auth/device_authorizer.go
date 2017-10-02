//
// cloud/azure/auth/device_authorizer.go
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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context/ctxhttp"
)

// DeviceCode defines a structure of a response to obtain a device code.
type DeviceCode struct {
	UserCode        string `json:"user_code"`
	DeviceCode      string `json:"device_code"`
	VerificationURL string `json:"verification_url"`
	ExpiresIn       string `json:"expires_in"`
	Interval        string `json:"interval"`
	Message         string `json:"message"`
}

// GetDeviceCode gets a device code.
func GetDeviceCode(ctx context.Context, clientID string) (code *DeviceCode, err error) {

	url := fmt.Sprintf(
		"https://login.microsoftonline.com/common/oauth2/devicecode?client_id=%v&resource=%v",
		clientID,
		url.QueryEscape("https://management.core.windows.net/"))
	//url.QueryEscape("00000002-0000-0000-c000-000000000000"))

	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")

	res, err := ctxhttp.Do(ctx, nil, req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	code = new(DeviceCode)
	err = json.NewDecoder(res.Body).Decode(&code)
	return

}

// AuthorizeDeviceCode runs authentication process by a device code.
func AuthorizeDeviceCode(ctx context.Context, clientID string, output io.Writer) (token *Token, err error) {

	code, err := GetDeviceCode(ctx, clientID)
	if err != nil {
		return
	}

	io.WriteString(output, code.Message)
	io.WriteString(output, "\n")

	expire, err := strconv.Atoi(code.ExpiresIn)
	if err != nil {
		return
	}
	interval, err := strconv.Atoi(code.Interval)
	if err != nil {
		return
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Duration(expire-30)*time.Second))
	defer cancel()

	var req *http.Request
	var res *http.Response
	for {

		body := fmt.Sprintf(
			"resource=%v&client_id=%v&grant_type=device_code&code=%v",
			url.QueryEscape("https://management.core.windows.net/"),
			//url.QueryEscape("00000002-0000-0000-c000-000000000000"),
			clientID,
			code.DeviceCode)
		req, err = http.NewRequest("Post", "https://login.microsoftonline.com/common/oauth2/token", strings.NewReader(body))
		if err != nil {
			return
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Accept", "application/json")

		res, err = ctxhttp.Do(ctx, nil, req)
		if err != nil {
			break
		}

		if res.StatusCode == 400 {
			var autherror TokenError
			err = json.NewDecoder(res.Body).Decode(&autherror)
			res.Body.Close()
			if err != nil {
				break
			}
			if strings.ToLower(autherror.ErrorSummary) != "authorization_pending" {
				break
			}

		} else {
			token = new(Token)
			err = json.NewDecoder(res.Body).Decode(token)
			res.Body.Close()
			if err != nil {
				return
			}
			break
		}

		select {
		case <-ctx.Done():
			err = ctx.Err()
			break
		case <-wait(time.Duration(interval) * time.Second):
		}

	}

	return

}

// Wait a given duration.
func wait(d time.Duration) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.Sleep(d)
		close(ch)
	}()
	return ch
}
