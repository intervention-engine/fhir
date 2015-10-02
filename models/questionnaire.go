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

type Questionnaire struct {
	Id          string                       `json:"id" bson:"_id"`
	Identifier  []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version     string                       `bson:"version,omitempty" json:"version,omitempty"`
	Status      string                       `bson:"status,omitempty" json:"status,omitempty"`
	Date        *FHIRDateTime                `bson:"date,omitempty" json:"date,omitempty"`
	Publisher   string                       `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Telecom     []ContactPoint               `bson:"telecom,omitempty" json:"telecom,omitempty"`
	SubjectType []string                     `bson:"subjectType,omitempty" json:"subjectType,omitempty"`
	Group       *QuestionnaireGroupComponent `bson:"group,omitempty" json:"group,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Questionnaire) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Questionnaire
	}{
		ResourceType:  "Questionnaire",
		Questionnaire: *resource,
	}
	return json.Marshal(x)
}

type QuestionnaireGroupComponent struct {
	LinkId   string                           `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Title    string                           `bson:"title,omitempty" json:"title,omitempty"`
	Concept  []Coding                         `bson:"concept,omitempty" json:"concept,omitempty"`
	Text     string                           `bson:"text,omitempty" json:"text,omitempty"`
	Required *bool                            `bson:"required,omitempty" json:"required,omitempty"`
	Repeats  *bool                            `bson:"repeats,omitempty" json:"repeats,omitempty"`
	Group    []QuestionnaireGroupComponent    `bson:"group,omitempty" json:"group,omitempty"`
	Question []QuestionnaireQuestionComponent `bson:"question,omitempty" json:"question,omitempty"`
}

type QuestionnaireQuestionComponent struct {
	LinkId   string                        `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Concept  []Coding                      `bson:"concept,omitempty" json:"concept,omitempty"`
	Text     string                        `bson:"text,omitempty" json:"text,omitempty"`
	Type     string                        `bson:"type,omitempty" json:"type,omitempty"`
	Required *bool                         `bson:"required,omitempty" json:"required,omitempty"`
	Repeats  *bool                         `bson:"repeats,omitempty" json:"repeats,omitempty"`
	Options  *Reference                    `bson:"options,omitempty" json:"options,omitempty"`
	Option   []Coding                      `bson:"option,omitempty" json:"option,omitempty"`
	Group    []QuestionnaireGroupComponent `bson:"group,omitempty" json:"group,omitempty"`
}
