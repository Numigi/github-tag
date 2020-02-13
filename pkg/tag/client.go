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
	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

// Alias to avoid propagation of dependency of go-github.
// Modules that use github-tag will not need to explicitly depend on go-github.
type GithubClient = github.Client


func GetGithubClientFromToken(token string) (*GithubClient, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc), nil
}
