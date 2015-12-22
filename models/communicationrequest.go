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

type CommunicationRequest struct {
	DomainResource    `bson:",inline"`
	Identifier        []Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Category          *CodeableConcept                       `bson:"category,omitempty" json:"category,omitempty"`
	Sender            *Reference                             `bson:"sender,omitempty" json:"sender,omitempty"`
	Recipient         []Reference                            `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Payload           []CommunicationRequestPayloadComponent `bson:"payload,omitempty" json:"payload,omitempty"`
	Medium            []CodeableConcept                      `bson:"medium,omitempty" json:"medium,omitempty"`
	Requester         *Reference                             `bson:"requester,omitempty" json:"requester,omitempty"`
	Status            string                                 `bson:"status,omitempty" json:"status,omitempty"`
	Encounter         *Reference                             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	ScheduledDateTime *FHIRDateTime                          `bson:"scheduledDateTime,omitempty" json:"scheduledDateTime,omitempty"`
	ScheduledPeriod   *Period                                `bson:"scheduledPeriod,omitempty" json:"scheduledPeriod,omitempty"`
	Reason            []CodeableConcept                      `bson:"reason,omitempty" json:"reason,omitempty"`
	RequestedOn       *FHIRDateTime                          `bson:"requestedOn,omitempty" json:"requestedOn,omitempty"`
	Subject           *Reference                             `bson:"subject,omitempty" json:"subject,omitempty"`
	Priority          *CodeableConcept                       `bson:"priority,omitempty" json:"priority,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *CommunicationRequest) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		CommunicationRequest
	}{
		ResourceType:         "CommunicationRequest",
		CommunicationRequest: *resource,
	}
	return json.Marshal(x)
}

// The "communicationRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type communicationRequest CommunicationRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *CommunicationRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := communicationRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = CommunicationRequest(x2)
	}
	return
}

type CommunicationRequestPayloadComponent struct {
	ContentString     string      `bson:"contentString,omitempty" json:"contentString,omitempty"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}
