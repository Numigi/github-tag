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
	"github.com/google/go-github/v28/github"
)


func Test_CreateTag(t *testing.T) {
	expectedNewTag, err := _getNextTag(testOwner, testRepo, testFilter)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	_, err = _createNextTag(testOwner, testRepo, testBranch, testFilter, "")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	newTag, err := _getCurrentTag(testOwner, testRepo, testFilter)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if newTag != expectedNewTag {
		t.Errorf("The new current tag (%v) does not match the expected value (%v).", newTag, expectedNewTag)
	}
}


func Test_CreateTag__CommitMessage(t *testing.T) {
	message := "Some commit detail message."

	tag, err := _createNextTag(testOwner, testRepo, testBranch, testFilter, message)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if *tag.Message != message {
		t.Errorf("The message (%s) for the created tag does not match the expected value (%s).", *tag.Message, message)
	}
}


func _createNextTag(owner, repo, branch, filter, message string) (*github.Tag, error) {
	nextTag, err := _getNextTag(owner, repo, filter)
	if err != nil {
		return nil, err;
	}

	client, err := getGithubClient()
	if err != nil {
		return nil, err;
	}

	return CreateTag(client, owner, repo, branch, nextTag, message)
}



func _getNextTag(owner, repo, filter string) (string, error) {
	currentTag, err := _getCurrentTag(owner, repo, filter)
	if err != nil {
		return "", err;
	}

	return GetNextTag(currentTag)
}


func _getCurrentTag(owner, repo, filter string) (string, error) {
	client, err := getGithubClient()
	if err != nil {
		return "", err;
	}

	return GetCurrentTag(client, owner, repo, filter)
}
