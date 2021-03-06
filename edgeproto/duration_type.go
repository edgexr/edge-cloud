// Copyright 2022 MobiledgeX, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package edgeproto

import (
	"encoding/json"
	fmt "fmt"
	reflect "reflect"
	"time"
)

// Duration Type allows protobufs to store/manage duration
// as a raw int64 value, but accept string values via json/yaml
// marshalling for friendly user input.

type Duration int64

func (e *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	err := unmarshal(&str)
	if err == nil {
		dur, err := time.ParseDuration(str)
		if err != nil {
			return err
		}
		*e = Duration(dur)
		return nil
	}
	var val int64
	err = unmarshal(&val)
	if err == nil {
		*e = Duration(val)
		return nil
	}
	return fmt.Errorf("Invalid duration type")
}

func (e Duration) MarshalYAML() (interface{}, error) {
	dur := time.Duration(e)
	return dur.String(), nil
}

func (e *Duration) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err == nil {
		dur, err := time.ParseDuration(str)
		if err != nil {
			return &json.UnmarshalTypeError{
				Value: "string " + str,
				Type:  reflect.TypeOf(Duration(0)),
			}
		}
		*e = Duration(dur)
		return nil
	}
	var val int64
	err = json.Unmarshal(b, &val)
	if err == nil {
		*e = Duration(val)
		return nil
	}
	return &json.UnmarshalTypeError{
		Value: "value " + str,
		Type:  reflect.TypeOf(Duration(0)),
	}
}

func (e Duration) MarshalJSON() ([]byte, error) {
	dur := time.Duration(e)
	return json.Marshal(dur.String())
}

func (e Duration) TimeDuration() time.Duration {
	return time.Duration(e)
}

// DecodeHook for use with mapstructure package.
func DecodeHook(from, to reflect.Type, data interface{}) (interface{}, error) {
	if from.Kind() == reflect.String {
		switch to {
		case reflect.TypeOf(Duration(0)):
			dur, err := time.ParseDuration(data.(string))
			if err != nil {
				return data, NewDurationParseError(data.(string), err)
			}
			return Duration(dur), nil
		case reflect.TypeOf(time.Duration(0)):
			dur, err := time.ParseDuration(data.(string))
			if err != nil {
				return data, NewDurationParseError(data.(string), err)
			}
			return dur, nil
		case reflect.TypeOf(time.Time{}):
			return time.Parse(time.RFC3339, data.(string))
		case reflect.TypeOf(Udec64{}):
			return ParseUdec64(data.(string))
		}
	}

	// decode enums
	return EnumDecodeHook(from, to, data)
}

// DurationParseError wraps a time.Duration parse error so that
// it can be recognized as such from errors returned from map decode.
type DurationParseError struct {
	Value string
	Err   error
}

func (e *DurationParseError) Error() string {
	return e.Err.Error()
}

func NewDurationParseError(val string, err error) *DurationParseError {
	dpe := DurationParseError{
		Value: val,
		Err:   err,
	}
	return &dpe
}
