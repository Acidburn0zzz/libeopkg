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
	"encoding/xml"
	"os"
	"sort"
)

// A Group as seen through the eyes of XML
type Group struct {
	Name string // ID of this group, i.e. "multimedia"

	// Translated short name
	LocalName []LocalisedField

	Icon string // Display icon for this Group
}

// Groups is a simple helper wrapper for loading from components.xml files
type Groups struct {
	Groups []Group `xml:"Groups>Group"`
}

// GroupSort allows us to quickly sort our groups by name
type GroupSort []Group

func (g GroupSort) Len() int {
	return len(g)
}

func (g GroupSort) Less(a, b int) bool {
	return g[a].Name < g[b].Name
}

func (g GroupSort) Swap(a, b int) {
	g[a], g[b] = g[b], g[a]
}

// NewGroups will load the Groups data from the XML file
func NewGroups(xmlfile string) (*Groups, error) {
	fi, err := os.Open(xmlfile)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	grp := &Groups{}
	dec := xml.NewDecoder(fi)
	if err = dec.Decode(grp); err != nil {
		return nil, err
	}
	sort.Sort(GroupSort(grp.Groups))
	// Ensure there are no empty Lang= fields
	for i := range grp.Groups {
		group := &grp.Groups[i]
		FixMissingLocalLanguage(&group.LocalName)
	}
	return grp, nil
}
