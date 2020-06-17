/*
   Copyright 2020 Authors of Arktos.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package server

import (
	"reflect"
	"testing"
)

func TestToCNIArgs(t *testing.T) {
	tcs := []struct{
		desc string
		input string
		expectingError bool
		expectedResult map[string]string
	}{
		{
			desc: "normal inputs",
			input: "FOO=BAR;ABC=123",
			expectingError: false,
			expectedResult: map[string]string{
				"ABC": "123",
				"FOO": "BAR",
			},
		},
		{
			desc: "empty input",
			input: "",
			expectingError: false,
			expectedResult: map[string]string{
			},
		},
		{
			desc: "invalid input with no key",
			input: "=value",
			expectingError: true,
		},
		{
			desc: "invalid input missing assigment char",
			input: "key",
			expectingError: true,
		},
		{
			desc: "valid input of empty value",
			input: "key=",
			expectingError: false,
			expectedResult: map[string]string{
				"key" : "",
			},
		},
		{
			desc: "invalid input with forbidden char",
			input: "key='k='",
			expectingError: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			kvs, err := toCNIArgs(tc.input)
			if tc.expectingError {
				if err == nil {
					t.Fatalf("expected error; got none")
				}
				t.Logf("got error: %v", err)
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if !reflect.DeepEqual(kvs, tc.expectedResult){
					t.Fatalf("expected %v; got %v", tc.expectedResult, kvs)
				}
			}
		})
	}
}