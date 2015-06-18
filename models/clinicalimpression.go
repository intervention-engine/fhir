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

import "time"

type ClinicalImpression struct {
	Id                     string                                                        `json:"-" bson:"_id"`
	Patient                *Reference                                                    `bson:"patient,omitempty" json:"patient,omitempty"`
	Assessor               *Reference                                                    `bson:"assessor,omitempty" json:"assessor,omitempty"`
	Status                 string                                                        `bson:"status,omitempty" json:"status,omitempty"`
	Date                   *FHIRDateTime                                                 `bson:"date,omitempty" json:"date,omitempty"`
	Description            string                                                        `bson:"description,omitempty" json:"description,omitempty"`
	Previous               *Reference                                                    `bson:"previous,omitempty" json:"previous,omitempty"`
	Problem                []Reference                                                   `bson:"problem,omitempty" json:"problem,omitempty"`
	TriggerCodeableConcept *CodeableConcept                                              `bson:"triggerCodeableConcept,omitempty" json:"triggerCodeableConcept,omitempty"`
	TriggerReference       *Reference                                                    `bson:"triggerReference,omitempty" json:"triggerReference,omitempty"`
	Investigations         []ClinicalImpressionClinicalImpressionInvestigationsComponent `bson:"investigations,omitempty" json:"investigations,omitempty"`
	Protocol               string                                                        `bson:"protocol,omitempty" json:"protocol,omitempty"`
	Summary                string                                                        `bson:"summary,omitempty" json:"summary,omitempty"`
	Finding                []ClinicalImpressionClinicalImpressionFindingComponent        `bson:"finding,omitempty" json:"finding,omitempty"`
	Resolved               []CodeableConcept                                             `bson:"resolved,omitempty" json:"resolved,omitempty"`
	RuledOut               []ClinicalImpressionClinicalImpressionRuledOutComponent       `bson:"ruledOut,omitempty" json:"ruledOut,omitempty"`
	Prognosis              string                                                        `bson:"prognosis,omitempty" json:"prognosis,omitempty"`
	Plan                   []Reference                                                   `bson:"plan,omitempty" json:"plan,omitempty"`
	Action                 []Reference                                                   `bson:"action,omitempty" json:"action,omitempty"`
}
type ClinicalImpressionClinicalImpressionInvestigationsComponent struct {
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Item []Reference      `bson:"item,omitempty" json:"item,omitempty"`
}
type ClinicalImpressionClinicalImpressionFindingComponent struct {
	Item  *CodeableConcept `bson:"item,omitempty" json:"item,omitempty"`
	Cause string           `bson:"cause,omitempty" json:"cause,omitempty"`
}
type ClinicalImpressionClinicalImpressionRuledOutComponent struct {
	Item   *CodeableConcept `bson:"item,omitempty" json:"item,omitempty"`
	Reason string           `bson:"reason,omitempty" json:"reason,omitempty"`
}

type ClinicalImpressionBundle struct {
	Type         string                          `json:"resourceType,omitempty"`
	Title        string                          `json:"title,omitempty"`
	Id           string                          `json:"id,omitempty"`
	Updated      time.Time                       `json:"updated,omitempty"`
	TotalResults int                             `json:"totalResults,omitempty"`
	Entry        []ClinicalImpressionBundleEntry `json:"entry,omitempty"`
	Category     ClinicalImpressionCategory      `json:"category,omitempty"`
}

type ClinicalImpressionBundleEntry struct {
	Title    string                     `json:"title,omitempty"`
	Id       string                     `json:"id,omitempty"`
	Content  ClinicalImpression         `json:"content,omitempty"`
	Category ClinicalImpressionCategory `json:"category,omitempty"`
}

type ClinicalImpressionCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
