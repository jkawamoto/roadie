//
// command/log/conv.go
//
// Copyright (c) 2016 Junpei Kawamoto
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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//

package log

import (
	"reflect"
	"strings"

	"github.com/golang/protobuf/ptypes/struct"
)

// TagKey defines a structure tag name for ConvertStructPB.
const TagKey = "structpb"

// ConvertStructPB converts a structpb.Struct object to a concrete object.
func ConvertStructPB(src *structpb.Struct, dest interface{}) error {

	r := reflect.Indirect(reflect.ValueOf(dest))
	for i := 0; i < r.NumField(); i++ {

		target := r.Field(i)
		targetType := r.Type().Field(i)

		name := targetType.Tag.Get(TagKey)
		if name == "" {
			name = strings.ToLower(targetType.Name)
		}

		if v, ok := src.GetFields()[name]; ok {
			switch t := v.GetKind().(type) {
			case *structpb.Value_BoolValue:
				target.SetBool(t.BoolValue)
			case *structpb.Value_ListValue:
				target.Set(reflect.ValueOf(t.ListValue))
			case *structpb.Value_NullValue:
				target.Set(reflect.ValueOf(t.NullValue))
			case *structpb.Value_NumberValue:
				target.Set(reflect.ValueOf(t.NumberValue))
			case *structpb.Value_StringValue:
				target.SetString(t.StringValue)
			case *structpb.Value_StructValue:
				ConvertStructPB(t.StructValue, target.Addr().Interface())
			}

		}

	}

	return nil

}
