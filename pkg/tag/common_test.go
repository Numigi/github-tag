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
	"os"
)

const (
    testOwner  = "numigi"
    testRepo   = "numigi_web_services_ci"
    testCommit = "c11c44833499f57eb57677cf86c14a86c143001f"
    testBranch = "master"
    testFilter = "fake"
)


func getTestingGithubClient() (*GithubClient, error) {
	accessTokenVarEnv := "GITHUB_ACCESS_TOKEN"
	accessToken := os.Getenv(accessTokenVarEnv)
	if accessToken == "" {
		return nil, fmt.Errorf("$%v is not defined", accessTokenVarEnv)
	}
	return GetGithubClientFromToken(accessToken)
}
