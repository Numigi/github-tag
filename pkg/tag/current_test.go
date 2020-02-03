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
	"fmt"
	"strings"
	"testing"
)


func Test_CurrentTag(t *testing.T) {
	type args struct {
		owner  string
		repo   string
		filter string
	}
	tests := []struct {
		name    string
		filter  string
		prefix  string
		wantErr bool // do we except the test to raise an error
	}{
		{
			"test_if_filter_given__current_tag_starts_with_filter",
			"",
			"zzz",
			false,
		},
		{
			"test_if_filter_given__current_tag_starts_with_filter",
			testFilter,
			testFilter,
			false,
		},
		{
			"When we ask for a filter that matches no tag, then an error is raised.",
			"NothingStartsWithThisFilter",
			"",
			true,
		},
	}
	for _, tt := range tests {
		client, err := getTestingGithubClient()
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCurrentTag(client, testOwner, testRepo, tt.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("CurrentTag() errored with %v: %v", tt.prefix, err)
				return
			}

			if !strings.HasPrefix(got, tt.prefix) {
				t.Errorf("CurrentTag() = %v, while it should start with %v", got, tt.prefix)
			}
		})
	}
}


func Test_CurrentTagTarballURL(t *testing.T) {
	client, err := getTestingGithubClient()
	if(err != nil){
		t.Errorf(err.Error())
		return
	}

	currentTag, err := GetCurrentTag(client, testOwner, testRepo, testFilter)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	got, err := GetCurrentTagTarballURL(client, testOwner, testRepo, testFilter)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	want := fmt.Sprintf("https://api.github.com/repos/Numigi/%v/tarball/%s", testRepo, currentTag)

	if got != want {
		t.Errorf("LatestTarballUrl() = %v, want %v", got, want)
	}
}
