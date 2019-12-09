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
	"strconv"
)


func GetNextTag(tagName string) (string, error) {
	major := getMajorVersion(tagName)

	minor, err := getMinorVersion(tagName)
	if err != nil {
		return "", err
	}

	incrementedMinor, err := incrementVersion(minor)
	if err != nil {
		return "", fmt.Errorf("Could not increment the minor version (%s) for the given tag (%v).", minor, tagName)
	}

	return fmt.Sprintf("%s.%s", major, incrementedMinor), nil
}


func getMajorVersion(tagName string) (string) {
	tagParts := strings.Split(tagName, ".")
	relevantTagParts := tagParts[:len(tagParts)-1]
	return strings.Join(relevantTagParts, ".")
}


func getMinorVersion(tagName string) (string, error) {
	versions := strings.Split(tagName, ".")

	if len(versions) < 2 {
		return "", fmt.Errorf("The given tag has no minor version: %v", tagName)
	}

	minor := versions[len(versions)-1];
	return minor, nil
}


func incrementVersion(version string) (string, error) {
	intVersion, err := strconv.Atoi(version)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(intVersion+1), nil
}
