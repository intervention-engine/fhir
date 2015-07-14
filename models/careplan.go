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

type CarePlan struct {
	Id          string                         `json:"-" bson:"_id"`
	Identifier  []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient     *Reference                     `bson:"patient,omitempty" json:"patient,omitempty"`
	Status      string                         `bson:"status,omitempty" json:"status,omitempty"`
	Period      *Period                        `bson:"period,omitempty" json:"period,omitempty"`
	Author      []Reference                    `bson:"author,omitempty" json:"author,omitempty"`
	Modified    *FHIRDateTime                  `bson:"modified,omitempty" json:"modified,omitempty"`
	Category    []CodeableConcept              `bson:"category,omitempty" json:"category,omitempty"`
	Concern     []Reference                    `bson:"concern,omitempty" json:"concern,omitempty"`
	Support     []Reference                    `bson:"support,omitempty" json:"support,omitempty"`
	Participant []CarePlanParticipantComponent `bson:"participant,omitempty" json:"participant,omitempty"`
	Goal        []Reference                    `bson:"goal,omitempty" json:"goal,omitempty"`
	Activity    []CarePlanActivityComponent    `bson:"activity,omitempty" json:"activity,omitempty"`
	Notes       string                         `bson:"notes,omitempty" json:"notes,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *CarePlan) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		CarePlan
	}{
		ResourceType: "CarePlan",
		CarePlan:     *resource,
	}
	return json.Marshal(x)
}

type CarePlanParticipantComponent struct {
	Role   *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Member *Reference       `bson:"member,omitempty" json:"member,omitempty"`
}

type CarePlanActivityComponent struct {
	ActionResulting []Reference                      `bson:"actionResulting,omitempty" json:"actionResulting,omitempty"`
	Notes           string                           `bson:"notes,omitempty" json:"notes,omitempty"`
	Reference       *Reference                       `bson:"reference,omitempty" json:"reference,omitempty"`
	Detail          *CarePlanActivityDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type CarePlanActivityDetailComponent struct {
	Category              string           `bson:"category,omitempty" json:"category,omitempty"`
	Code                  *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	ReasonCodeableConcept *CodeableConcept `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference       `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Goal                  []Reference      `bson:"goal,omitempty" json:"goal,omitempty"`
	Status                string           `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason          *CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Prohibited            *bool            `bson:"prohibited,omitempty" json:"prohibited,omitempty"`
	ScheduledTiming       *Timing          `bson:"scheduledTiming,omitempty" json:"scheduledTiming,omitempty"`
	ScheduledPeriod       *Period          `bson:"scheduledPeriod,omitempty" json:"scheduledPeriod,omitempty"`
	ScheduledString       string           `bson:"scheduledString,omitempty" json:"scheduledString,omitempty"`
	Location              *Reference       `bson:"location,omitempty" json:"location,omitempty"`
	Performer             []Reference      `bson:"performer,omitempty" json:"performer,omitempty"`
	Product               *Reference       `bson:"product,omitempty" json:"product,omitempty"`
	DailyAmount           *Quantity        `bson:"dailyAmount,omitempty" json:"dailyAmount,omitempty"`
	Quantity              *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Note                  string           `bson:"note,omitempty" json:"note,omitempty"`
}
