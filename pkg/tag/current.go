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
	"context"
	"fmt"
	"github.com/google/go-github/v28/github"
	"strings"
)


func GetCurrentTag(client *GithubClient, owner, repo, filter string) (string, error) {
	tag, err := getCurrentTagObject(client, owner, repo, filter)
	if err != nil {
		return "", err
	}
	return *tag.Name, nil
}


func GetCurrentTagTarballURL(client *GithubClient, owner, repo, filter string) (string, error) {
	tag, err := getCurrentTagObject(client, owner, repo, filter)
	if err != nil {
		return "", err
	}
	return *tag.TarballURL, nil
}


func getCurrentTagObject(client *GithubClient, owner, repo, filter string) (*github.RepositoryTag, error) {
	page := 0
	for {
		tags, err := listTags(client, owner, repo, page)

		if err != nil {
			return nil, err
		}

		if len(tags) == 0 {
			return nil, fmt.Errorf("No tag found for %s/%s with the filter %s", owner, repo, filter)
		}

		latestTag := getFirstMatchingTag(&tags, filter)

		if latestTag != nil {
			return latestTag, nil
		}

		page++
	}
}


const numberOfTagPerPage = 100


func listTags(client *GithubClient, owner string, repo string, page int) ([]*github.RepositoryTag, error) {
	listOptions := github.ListOptions {
		Page:    page,
		PerPage: numberOfTagPerPage,
	}
	ctx := context.Background()
	tags, _, err := client.Repositories.ListTags(ctx, owner, repo, &listOptions)
	return tags, err
}


func getFirstMatchingTag(tags *[]*github.RepositoryTag, filter string) *github.RepositoryTag {
	for _, tag := range *tags {
		if strings.HasPrefix(*tag.Name, filter) {
			return tag
		}
	}
	return nil
}
