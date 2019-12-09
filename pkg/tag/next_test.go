/*
Copyright © 2019 - today Numigi <contact@numigi.com>

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

package tag

import (
	"testing"
)


func Test_NextTag(t *testing.T) {
	tests := []struct {
		tagName    string
		expected   string
	}{
		{
			"abc.0",
			"abc.1",
		},
		{
			"10.0",
			"10.1",
		},
		{
			"10.0.1",
			"10.0.2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.tagName, func(t *testing.T) {
			got, err := GetNextTag(tt.tagName)

			if err != nil {
				t.Errorf(err.Error())
			}

			if tt.expected != got {
				t.Errorf("Expected %v, got %v", tt.expected, got)
			}
		})
	}
}


func Test_NextTag_WrongInputTagName(t *testing.T) {
	tests := []struct {
		testName    string
		tagName   string
	}{
		{
			"Empty TagName",
			"",
		},
		{
			"Non-digit minor version",
			"123.abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := GetNextTag(tt.tagName)

			if err == nil {
				t.Errorf("Expected NextTag to return an error.")
			}
		})
	}
}

