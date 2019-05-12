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
)

// A Distribution as seen through the eyes of XML
type Distribution struct {
	SourceName string // Name of source to match source repos

	// Translated description
	Description []LocalisedField

	Version    string // Published version number for compatibility
	Type       string // Type of repository (should always be main, really. Just descriptive)
	BinaryName string // Name of the binary repository

	Obsoletes []string `xml:"Obsoletes>Package"` // Package names that are no longer supported

	obsmap map[string]bool
}

// NewDistribution will load the Distribution data from the XML file
func NewDistribution(xmlfile string) (dist *Distribution, err error) {
	fi, err := os.Open(xmlfile)
	if err != nil {
		return
	}
	defer fi.Close()
	dist = &Distribution{
		obsmap: make(map[string]bool),
	}
	dec := xml.NewDecoder(fi)
	if err = dec.Decode(dist); err != nil {
		return
	}
	for _, p := range dist.Obsoletes {
		dist.obsmap[p] = true
	}
	return
}

// IsObsolete will allow quickly determination of whether the package name
// was marked obsolete and should be hidden from the index
func (d *Distribution) IsObsolete(id string) bool {
	return d.obsmap[id]
}
