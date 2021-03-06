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

package util

import (
	"strconv"
	"strings"

	"github.com/kballard/go-shellquote"
)

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}
func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func CamelCase(s string) string {
	t := ""
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '_' {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t += string(c)
			continue
		}
		if isASCIILower(c) {
			c ^= ' ' // Make it upper case
		}
		t += string(c)
		// Convert upper case to lower case following an upper case character
		for i+1 < len(s) && isASCIIUpper(s[i+1]) {
			if i+2 < len(s) && isASCIILower(s[i+2]) {
				break
			}
			i++
			t += string(s[i] ^ ' ') // Make it lower case
		}
		// Accept lower case sequence that follows.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t += string(s[i])
		}
	}
	return t
}

func EscapeJson(jsoninput string) string {
	r := strings.NewReplacer(
		`{`, `\{`, `}`, `\}`)
	return r.Replace(jsoninput)
}

func CapitalizeMessage(msg string) string {
	if len(msg) == 0 {
		return msg
	}
	c := msg[0]
	// Return msg if already capitalized
	if !isASCIILower(c) {
		return msg
	}
	// Capitalize first character and append to rest of msg
	t := string(msg[1:])
	c ^= ' '
	t = string(c) + t
	return t
}

func UncapitalizeMessage(msg string) string {
	c := msg[0]
	// Return msg if already lower case
	if !isASCIIUpper(c) {
		return msg
	}
	t := string(msg[1:])
	c += 'a' - 'A'
	t = string(c) + t
	return t
}

func SplitCamelCase(name string) []string {
	out := []string{}
	if name == "" {
		return out
	}
	startIndex := 0
	for ii := 1; ii < len(name); ii++ {
		if isASCIIUpper(name[ii]) {
			out = append(out, name[startIndex:ii])
			startIndex = ii
		}
	}
	if startIndex < len(name) {
		out = append(out, name[startIndex:])
	}
	return out
}

// UnCamelCase converts camel case to lowercase separated by underscore
func UnCamelCase(name string) string {
	parts := SplitCamelCase(name)
	for ii := range parts {
		parts[ii] = strings.ToLower(parts[ii])
	}
	return strings.Join(parts, "_")
}

func QuoteArgs(cmd string) (string, error) {
	cmd = strings.TrimSpace(cmd)
	args, err := shellquote.Split(cmd)
	if err != nil {
		return "", err
	}
	for i := range args {
		args[i] = strconv.Quote(args[i])
	}
	return strings.Join(args, " "), nil
}
