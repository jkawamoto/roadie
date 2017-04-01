//
// cloud/azure/error.go
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
	"fmt"
	"io"
	"io/ioutil"
	"reflect"

	"github.com/go-openapi/runtime"
)

// NewAPIError creates an error which is raised from an API call.
func NewAPIError(err error) error {
	switch e := err.(type) {
	case *runtime.APIError:
		response := reflect.ValueOf(e.Response)
		body := response.MethodByName("Body")
		res := body.Call(nil)[0]
		if res.CanInterface() {
			reader, ok := res.Interface().(io.ReadCloser)
			if ok {
				msg, err2 := ioutil.ReadAll(reader)
				if err2 == nil {
					fmt.Println(string(msg))
				}
			}
		}

		message := response.MethodByName("Message")
		res = message.Call(nil)[0]
		if res.CanInterface() {
			msg, ok := res.Interface().(string)
			if ok {
				return fmt.Errorf("API Error: %s", msg)
			}
		}

		return fmt.Errorf("API error: %v", response.FieldByName("resp"))

	default:
		return e
	}
}
