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

type QuestionnaireAnswers struct {
	Id            string                              `json:"-" bson:"_id"`
	Identifier    *Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Questionnaire *Reference                          `bson:"questionnaire,omitempty" json:"questionnaire,omitempty"`
	Status        string                              `bson:"status,omitempty" json:"status,omitempty"`
	Subject       *Reference                          `bson:"subject,omitempty" json:"subject,omitempty"`
	Author        *Reference                          `bson:"author,omitempty" json:"author,omitempty"`
	Authored      *FHIRDateTime                       `bson:"authored,omitempty" json:"authored,omitempty"`
	Source        *Reference                          `bson:"source,omitempty" json:"source,omitempty"`
	Encounter     *Reference                          `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Group         *QuestionnaireAnswersGroupComponent `bson:"group,omitempty" json:"group,omitempty"`
}

type QuestionnaireAnswersGroupComponent struct {
	LinkId   string                                  `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Title    string                                  `bson:"title,omitempty" json:"title,omitempty"`
	Text     string                                  `bson:"text,omitempty" json:"text,omitempty"`
	Subject  *Reference                              `bson:"subject,omitempty" json:"subject,omitempty"`
	Group    []QuestionnaireAnswersGroupComponent    `bson:"group,omitempty" json:"group,omitempty"`
	Question []QuestionnaireAnswersQuestionComponent `bson:"question,omitempty" json:"question,omitempty"`
}

type QuestionnaireAnswersQuestionComponent struct {
	LinkId string                                        `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Text   string                                        `bson:"text,omitempty" json:"text,omitempty"`
	Answer []QuestionnaireAnswersQuestionAnswerComponent `bson:"answer,omitempty" json:"answer,omitempty"`
	Group  []QuestionnaireAnswersGroupComponent          `bson:"group,omitempty" json:"group,omitempty"`
}

type QuestionnaireAnswersQuestionAnswerComponent struct {
	ValueBoolean    *bool         `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueDecimal    *float64      `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueInteger    *int32        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate       *FHIRDateTime `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime   *FHIRDateTime `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueInstant    *FHIRDateTime `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueTime       *FHIRDateTime `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString     string        `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueUri        string        `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueAttachment *Attachment   `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCoding     *Coding       `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueQuantity   *Quantity     `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueReference  *Reference    `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
}

type QuestionnaireAnswersBundle struct {
	Type         string                            `json:"resourceType,omitempty"`
	Title        string                            `json:"title,omitempty"`
	Id           string                            `json:"id,omitempty"`
	Updated      time.Time                         `json:"updated,omitempty"`
	TotalResults int                               `json:"totalResults,omitempty"`
	Entry        []QuestionnaireAnswersBundleEntry `json:"entry,omitempty"`
	Category     QuestionnaireAnswersCategory      `json:"category,omitempty"`
}

type QuestionnaireAnswersBundleEntry struct {
	Title    string                       `json:"title,omitempty"`
	Id       string                       `json:"id,omitempty"`
	Content  QuestionnaireAnswers         `json:"content,omitempty"`
	Category QuestionnaireAnswersCategory `json:"category,omitempty"`
}

type QuestionnaireAnswersCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}

func (resource *QuestionnaireAnswers) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		QuestionnaireAnswers
	}{
		ResourceType:         "QuestionnaireAnswers",
		QuestionnaireAnswers: *resource,
	}
	return json.Marshal(x)
}
