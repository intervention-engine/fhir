// Copyright (c) 2011-2015, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import "encoding/json"

type Provenance struct {
	Id        string                      `json:"-" bson:"_id"`
	Target    []Reference                 `bson:"target,omitempty" json:"target,omitempty"`
	Period    *Period                     `bson:"period,omitempty" json:"period,omitempty"`
	Recorded  *FHIRDateTime               `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Reason    *CodeableConcept            `bson:"reason,omitempty" json:"reason,omitempty"`
	Location  *Reference                  `bson:"location,omitempty" json:"location,omitempty"`
	Policy    []string                    `bson:"policy,omitempty" json:"policy,omitempty"`
	Agent     []ProvenanceAgentComponent  `bson:"agent,omitempty" json:"agent,omitempty"`
	Entity    []ProvenanceEntityComponent `bson:"entity,omitempty" json:"entity,omitempty"`
	Signature []Signature                 `bson:"signature,omitempty" json:"signature,omitempty"`
}

type ProvenanceAgentComponent struct {
	Role               *Coding    `bson:"role,omitempty" json:"role,omitempty"`
	Type               *Coding    `bson:"type,omitempty" json:"type,omitempty"`
	ReferenceUri       string     `bson:"referenceUri,omitempty" json:"referenceUri,omitempty"`
	ReferenceReference *Reference `bson:"referenceReference,omitempty" json:"referenceReference,omitempty"`
	Display            string     `bson:"display,omitempty" json:"display,omitempty"`
}

type ProvenanceEntityComponent struct {
	Role      string                    `bson:"role,omitempty" json:"role,omitempty"`
	Type      *Coding                   `bson:"type,omitempty" json:"type,omitempty"`
	Reference string                    `bson:"reference,omitempty" json:"reference,omitempty"`
	Display   string                    `bson:"display,omitempty" json:"display,omitempty"`
	Agent     *ProvenanceAgentComponent `bson:"agent,omitempty" json:"agent,omitempty"`
}

type ProvenanceBundle struct {
	Id    string                  `json:"id,omitempty"`
	Type  string                  `json:"resourceType,omitempty"`
	Base  string                  `json:"base,omitempty"`
	Total int                     `json:"total,omitempty"`
	Link  []BundleLinkComponent   `json:"link,omitempty"`
	Entry []ProvenanceBundleEntry `json:"entry,omitempty"`
}

type ProvenanceBundleEntry struct {
	Id       string                `json:"id,omitempty"`
	Base     string                `json:"base,omitempty"`
	Link     []BundleLinkComponent `json:"link,omitempty"`
	Resource Provenance            `json:"resource,omitempty"`
}

func (resource *Provenance) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Provenance
	}{
		ResourceType: "Provenance",
		Provenance:   *resource,
	}
	return json.Marshal(x)
}
