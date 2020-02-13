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
)



func CreateTag(client *GithubClient, owner, repo, branch, tagName, message string) (*github.Tag, error) {
	sha, err := getBranchSHA(client, owner, repo, branch)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	objectType := "commit"
	tag, _, err := client.Git.CreateTag(ctx, owner, repo, &github.Tag{
		Tag:     &tagName,
		SHA:     &sha,
		Message: &message,
		Object:  &github.GitObject{SHA: &sha, Type: &objectType},
	})
	if err != nil {
		return nil, err
	}

	_, _, err = client.Git.CreateRef(
		ctx,
		owner,
		repo,
		&github.Reference{
			Ref:    github.String(fmt.Sprintf("refs/tags/%v", *tag.Tag)),
			Object: &github.GitObject{SHA: tag.SHA},
		},
	)
	if err != nil {
		return nil, err
	}
	return tag, nil
}
