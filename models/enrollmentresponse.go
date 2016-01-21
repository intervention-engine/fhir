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

type EnrollmentResponse struct {
	DomainResource      `bson:",inline"`
	Identifier          []Identifier  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Request             *Reference    `bson:"request,omitempty" json:"request,omitempty"`
	Outcome             string        `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition         string        `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Ruleset             *Coding       `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset     *Coding       `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created             *FHIRDateTime `bson:"created,omitempty" json:"created,omitempty"`
	Organization        *Reference    `bson:"organization,omitempty" json:"organization,omitempty"`
	RequestProvider     *Reference    `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization *Reference    `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *EnrollmentResponse) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		EnrollmentResponse
	}{
		ResourceType:       "EnrollmentResponse",
		EnrollmentResponse: *resource,
	}
	return json.Marshal(x)
}

// The "enrollmentResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type enrollmentResponse EnrollmentResponse

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *EnrollmentResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := enrollmentResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = EnrollmentResponse(x2)
	}
	return
}

type EnrollmentResponsePlus struct {
	EnrollmentResponse             `bson:",inline"`
	EnrollmentResponsePlusIncludes `bson:",inline"`
}

type EnrollmentResponsePlusIncludes struct {
}

func (e *EnrollmentResponsePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}
