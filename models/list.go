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

type List struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Title          string               `bson:"title,omitempty" json:"title,omitempty"`
	Code           *CodeableConcept     `bson:"code,omitempty" json:"code,omitempty"`
	Subject        *Reference           `bson:"subject,omitempty" json:"subject,omitempty"`
	Source         *Reference           `bson:"source,omitempty" json:"source,omitempty"`
	Encounter      *Reference           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Status         string               `bson:"status,omitempty" json:"status,omitempty"`
	Date           *FHIRDateTime        `bson:"date,omitempty" json:"date,omitempty"`
	OrderedBy      *CodeableConcept     `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Mode           string               `bson:"mode,omitempty" json:"mode,omitempty"`
	Note           string               `bson:"note,omitempty" json:"note,omitempty"`
	Entry          []ListEntryComponent `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason    *CodeableConcept     `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *List) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		List
	}{
		ResourceType: "List",
		List:         *resource,
	}
	return json.Marshal(x)
}

// The "list" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type list List

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *List) UnmarshalJSON(data []byte) (err error) {
	x2 := list{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = List(x2)
	}
	return
}

type ListEntryComponent struct {
	Flag    *CodeableConcept `bson:"flag,omitempty" json:"flag,omitempty"`
	Deleted *bool            `bson:"deleted,omitempty" json:"deleted,omitempty"`
	Date    *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Item    *Reference       `bson:"item,omitempty" json:"item,omitempty"`
}
