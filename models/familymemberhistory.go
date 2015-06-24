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

import (
	"encoding/json"
	"time"
)

type FamilyMemberHistory struct {
	Id              string                                  `json:"-" bson:"_id"`
	Identifier      []Identifier                            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient         *Reference                              `bson:"patient,omitempty" json:"patient,omitempty"`
	Date            *FHIRDateTime                           `bson:"date,omitempty" json:"date,omitempty"`
	Name            string                                  `bson:"name,omitempty" json:"name,omitempty"`
	Relationship    *CodeableConcept                        `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Gender          *CodeableConcept                        `bson:"gender,omitempty" json:"gender,omitempty"`
	BornPeriod      *Period                                 `bson:"bornPeriod,omitempty" json:"bornPeriod,omitempty"`
	BornDate        *FHIRDateTime                           `bson:"bornDate,omitempty" json:"bornDate,omitempty"`
	BornString      string                                  `bson:"bornString,omitempty" json:"bornString,omitempty"`
	AgeAge          *Quantity                               `bson:"ageAge,omitempty" json:"ageAge,omitempty"`
	AgeRange        *Range                                  `bson:"ageRange,omitempty" json:"ageRange,omitempty"`
	AgeString       string                                  `bson:"ageString,omitempty" json:"ageString,omitempty"`
	DeceasedBoolean *bool                                   `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedAge     *Quantity                               `bson:"deceasedAge,omitempty" json:"deceasedAge,omitempty"`
	DeceasedRange   *Range                                  `bson:"deceasedRange,omitempty" json:"deceasedRange,omitempty"`
	DeceasedDate    *FHIRDateTime                           `bson:"deceasedDate,omitempty" json:"deceasedDate,omitempty"`
	DeceasedString  string                                  `bson:"deceasedString,omitempty" json:"deceasedString,omitempty"`
	Note            string                                  `bson:"note,omitempty" json:"note,omitempty"`
	Condition       []FamilyMemberHistoryConditionComponent `bson:"condition,omitempty" json:"condition,omitempty"`
}

type FamilyMemberHistoryConditionComponent struct {
	Type        *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Outcome     *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OnsetAge    *Quantity        `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetRange  *Range           `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetString string           `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	Note        string           `bson:"note,omitempty" json:"note,omitempty"`
}

type FamilyMemberHistoryBundle struct {
	Type         string                           `json:"resourceType,omitempty"`
	Title        string                           `json:"title,omitempty"`
	Id           string                           `json:"id,omitempty"`
	Updated      time.Time                        `json:"updated,omitempty"`
	TotalResults int                              `json:"totalResults,omitempty"`
	Entry        []FamilyMemberHistoryBundleEntry `json:"entry,omitempty"`
	Category     FamilyMemberHistoryCategory      `json:"category,omitempty"`
}

type FamilyMemberHistoryBundleEntry struct {
	Title    string                      `json:"title,omitempty"`
	Id       string                      `json:"id,omitempty"`
	Content  FamilyMemberHistory         `json:"content,omitempty"`
	Category FamilyMemberHistoryCategory `json:"category,omitempty"`
}

type FamilyMemberHistoryCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}

func (resource *FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		FamilyMemberHistory
	}{
		ResourceType:        "FamilyMemberHistory",
		FamilyMemberHistory: *resource,
	}
	return json.Marshal(x)
}
