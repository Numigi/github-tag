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
)


func getBranchSHA(client *GithubClient, owner, repo, branchName string) (string, error) {
	ctx := context.Background()
	branch, _, err := client.Repositories.GetBranch(ctx, owner, repo, branchName)
	if err != nil {
		return "", err
	}
	return *branch.Commit.SHA, nil
}


func getTagSHA(client *GithubClient, owner, repo, tagName string) (string, error) {
	ctx := context.Background()
	tag, _, err := client.Git.GetTag(ctx, owner, repo, tagName)
	if err != nil {
		return "", err
	}
	return *tag.SHA, nil
}
