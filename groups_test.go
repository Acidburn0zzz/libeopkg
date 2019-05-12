//
// Copyright © 2017-2019 Solus Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package libeopkg

import (
	"testing"
)

const (
	notAGroup     = "testdata/not.xml"
	groupTestFile = "testdata/groups.xml"
)

func TestGroups(t *testing.T) {
	grp, err := NewGroups(groupTestFile)
	if err != nil {
		t.Fatalf("Failed to load good file: %s", err)
	}
	if grp == nil {
		t.Fatalf("Failed to get group")
	}

	var want *Group
	for i := range grp.Groups {
		c := &grp.Groups[i]
		if c.Name == "multimedia" {
			want = c
			break
		}
	}

	if want == nil {
		t.Fatal("Cannot find desired group multimedia")
	}

	if len(want.LocalName) != 23 {
		t.Fatalf("Invalid number of LocalNames: %d", len(want.LocalName))
	}

	if want.LocalName[0].Lang != "en" {
		t.Fatalf("First element should be en by default: %s", want.LocalName[0].Lang)
	}
	if want.LocalName[22].Lang != "zh_CN" {
		t.Fatalf("Wrong lang on last element: %s", want.LocalName[22].Lang)
	}
	if want.Icon != "multimedia-volume-control" {
		t.Fatalf("Wrong icon: %s", want.Icon)
	}
}

func TestGroupsMissing(t *testing.T) {
	_, err := NewGroups(notAFile)
	if err == nil {
		t.Fatalf("Should have failed to load missing file: %s", notAFile)
	}
}

func TestGroupsInvalid(t *testing.T) {
	_, err := NewGroups(notAGroup)
	if err == nil {
		t.Fatalf("Should have failed to load invalid file: %s", notAGroup)
	}
}
