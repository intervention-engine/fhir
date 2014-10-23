// Copyright (c) 2011-2014, HL7, Inc & The MITRE Corporation
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

type CarePlan struct {
	Id          string                         `json:"-" bson:"_id"`
	Identifier  []Identifier                   `bson:"identifier"`
	Patient     Reference                      `bson:"patient"`
	Status      string                         `bson:"status"`
	Period      Period                         `bson:"period"`
	Modified    FHIRDateTime                   `bson:"modified"`
	Concern     []Reference                    `bson:"concern"`
	Participant []CarePlanParticipantComponent `bson:"participant"`
	Goal        []CarePlanGoalComponent        `bson:"goal"`
	Activity    []CarePlanActivityComponent    `bson:"activity"`
	Notes       string                         `bson:"notes"`
}

// This is an ugly hack to deal with embedded structures in the spec participant
type CarePlanParticipantComponent struct {
	Role   CodeableConcept `bson:"role"`
	Member Reference       `bson:"member"`
}

// This is an ugly hack to deal with embedded structures in the spec goal
type CarePlanGoalComponent struct {
	Description string      `bson:"description"`
	Status      string      `bson:"status"`
	Notes       string      `bson:"notes"`
	Concern     []Reference `bson:"concern"`
}

// This is an ugly hack to deal with embedded structures in the spec simple
type CarePlanActivitySimpleComponent struct {
	Category        string          `bson:"category"`
	Code            CodeableConcept `bson:"code"`
	ScheduledTiming Timing          `bson:"scheduledTiming"`
	ScheduledPeriod Period          `bson:"scheduledPeriod"`
	ScheduledString string          `bson:"scheduledString"`
	Location        Reference       `bson:"location"`
	Performer       []Reference     `bson:"performer"`
	Product         Reference       `bson:"product"`
	DailyAmount     Quantity        `bson:"dailyAmount"`
	Quantity        Quantity        `bson:"quantity"`
	Details         string          `bson:"details"`
}

// This is an ugly hack to deal with embedded structures in the spec activity
type CarePlanActivityComponent struct {
	Goal            []Reference                     `bson:"goal"`
	Status          string                          `bson:"status"`
	Prohibited      bool                            `bson:"prohibited"`
	ActionResulting []Reference                     `bson:"actionResulting"`
	Notes           string                          `bson:"notes"`
	Detail          Reference                       `bson:"detail"`
	Simple          CarePlanActivitySimpleComponent `bson:"simple"`
}
