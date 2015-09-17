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

type Procedure struct {
	Id                    string                          `json:"id" bson:"_id"`
	Identifier            []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject               *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Status                string                          `bson:"status,omitempty" json:"status,omitempty"`
	Category              *CodeableConcept                `bson:"category,omitempty" json:"category,omitempty"`
	Code                  *CodeableConcept                `bson:"code,omitempty" json:"code,omitempty"`
	NotPerformed          *bool                           `bson:"notPerformed,omitempty" json:"notPerformed,omitempty"`
	ReasonNotPerformed    []CodeableConcept               `bson:"reasonNotPerformed,omitempty" json:"reasonNotPerformed,omitempty"`
	BodySite              []CodeableConcept               `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	ReasonCodeableConcept *CodeableConcept                `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                      `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Performer             []ProcedurePerformerComponent   `bson:"performer,omitempty" json:"performer,omitempty"`
	PerformedDateTime     *FHIRDateTime                   `bson:"performedDateTime,omitempty" json:"performedDateTime,omitempty"`
	PerformedPeriod       *Period                         `bson:"performedPeriod,omitempty" json:"performedPeriod,omitempty"`
	Encounter             *Reference                      `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Location              *Reference                      `bson:"location,omitempty" json:"location,omitempty"`
	Outcome               *CodeableConcept                `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Report                []Reference                     `bson:"report,omitempty" json:"report,omitempty"`
	Complication          []CodeableConcept               `bson:"complication,omitempty" json:"complication,omitempty"`
	FollowUp              []CodeableConcept               `bson:"followUp,omitempty" json:"followUp,omitempty"`
	Request               *Reference                      `bson:"request,omitempty" json:"request,omitempty"`
	Notes                 []Annotation                    `bson:"notes,omitempty" json:"notes,omitempty"`
	FocalDevice           []ProcedureFocalDeviceComponent `bson:"focalDevice,omitempty" json:"focalDevice,omitempty"`
	Used                  []Reference                     `bson:"used,omitempty" json:"used,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Procedure) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Procedure
	}{
		ResourceType: "Procedure",
		Procedure:    *resource,
	}
	return json.Marshal(x)
}

type ProcedurePerformerComponent struct {
	Actor *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
	Role  *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type ProcedureFocalDeviceComponent struct {
	Action      *CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Manipulated *Reference       `bson:"manipulated,omitempty" json:"manipulated,omitempty"`
}
