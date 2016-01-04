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
	DomainResource `bson:",inline"`
	Target         []Reference                 `bson:"target,omitempty" json:"target,omitempty"`
	Period         *Period                     `bson:"period,omitempty" json:"period,omitempty"`
	Recorded       *FHIRDateTime               `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Reason         []CodeableConcept           `bson:"reason,omitempty" json:"reason,omitempty"`
	Activity       *CodeableConcept            `bson:"activity,omitempty" json:"activity,omitempty"`
	Location       *Reference                  `bson:"location,omitempty" json:"location,omitempty"`
	Policy         []string                    `bson:"policy,omitempty" json:"policy,omitempty"`
	Agent          []ProvenanceAgentComponent  `bson:"agent,omitempty" json:"agent,omitempty"`
	Entity         []ProvenanceEntityComponent `bson:"entity,omitempty" json:"entity,omitempty"`
	Signature      []Signature                 `bson:"signature,omitempty" json:"signature,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
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

// The "provenance" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type provenance Provenance

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Provenance) UnmarshalJSON(data []byte) (err error) {
	x2 := provenance{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Provenance(x2)
	}
	return
}

type ProvenanceAgentComponent struct {
	Role         *Coding                                `bson:"role,omitempty" json:"role,omitempty"`
	Actor        *Reference                             `bson:"actor,omitempty" json:"actor,omitempty"`
	UserId       *Identifier                            `bson:"userId,omitempty" json:"userId,omitempty"`
	RelatedAgent []ProvenanceAgentRelatedAgentComponent `bson:"relatedAgent,omitempty" json:"relatedAgent,omitempty"`
}

type ProvenanceAgentRelatedAgentComponent struct {
	Type   *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Target string           `bson:"target,omitempty" json:"target,omitempty"`
}

type ProvenanceEntityComponent struct {
	Role      string                    `bson:"role,omitempty" json:"role,omitempty"`
	Type      *Coding                   `bson:"type,omitempty" json:"type,omitempty"`
	Reference string                    `bson:"reference,omitempty" json:"reference,omitempty"`
	Display   string                    `bson:"display,omitempty" json:"display,omitempty"`
	Agent     *ProvenanceAgentComponent `bson:"agent,omitempty" json:"agent,omitempty"`
}
