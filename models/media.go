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

type Media struct {
	DomainResource `bson:",inline"`
	Type           string           `bson:"type,omitempty" json:"type,omitempty"`
	Subtype        *CodeableConcept `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Identifier     []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject        *Reference       `bson:"subject,omitempty" json:"subject,omitempty"`
	Operator       *Reference       `bson:"operator,omitempty" json:"operator,omitempty"`
	View           *CodeableConcept `bson:"view,omitempty" json:"view,omitempty"`
	DeviceName     string           `bson:"deviceName,omitempty" json:"deviceName,omitempty"`
	Height         *uint32          `bson:"height,omitempty" json:"height,omitempty"`
	Width          *uint32          `bson:"width,omitempty" json:"width,omitempty"`
	Frames         *uint32          `bson:"frames,omitempty" json:"frames,omitempty"`
	Duration       *uint32          `bson:"duration,omitempty" json:"duration,omitempty"`
	Content        *Attachment      `bson:"content,omitempty" json:"content,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Media) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Media
	}{
		ResourceType: "Media",
		Media:        *resource,
	}
	return json.Marshal(x)
}

// The "media" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type media Media

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Media) UnmarshalJSON(data []byte) (err error) {
	x2 := media{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Media(x2)
	}
	return
}
