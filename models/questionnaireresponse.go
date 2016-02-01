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
	"errors"
	"fmt"
)

type QuestionnaireResponse struct {
	DomainResource `bson:",inline"`
	Identifier     *Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Questionnaire  *Reference                           `bson:"questionnaire,omitempty" json:"questionnaire,omitempty"`
	Status         string                               `bson:"status,omitempty" json:"status,omitempty"`
	Subject        *Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Author         *Reference                           `bson:"author,omitempty" json:"author,omitempty"`
	Authored       *FHIRDateTime                        `bson:"authored,omitempty" json:"authored,omitempty"`
	Source         *Reference                           `bson:"source,omitempty" json:"source,omitempty"`
	Encounter      *Reference                           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Group          *QuestionnaireResponseGroupComponent `bson:"group,omitempty" json:"group,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "QuestionnaireResponse"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to QuestionnaireResponse), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *QuestionnaireResponse) GetBSON() (interface{}, error) {
	x.ResourceType = "QuestionnaireResponse"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "questionnaireResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type questionnaireResponse QuestionnaireResponse

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *QuestionnaireResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := questionnaireResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = QuestionnaireResponse(x2)
		return x.checkResourceType()
	}
	return
}

func (x *QuestionnaireResponse) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "QuestionnaireResponse"
	} else if x.ResourceType != "QuestionnaireResponse" {
		return errors.New(fmt.Sprintf("Expected resourceType to be QuestionnaireResponse, instead received %s", x.ResourceType))
	}
	return nil
}

type QuestionnaireResponseGroupComponent struct {
	LinkId   string                                   `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Title    string                                   `bson:"title,omitempty" json:"title,omitempty"`
	Text     string                                   `bson:"text,omitempty" json:"text,omitempty"`
	Subject  *Reference                               `bson:"subject,omitempty" json:"subject,omitempty"`
	Group    []QuestionnaireResponseGroupComponent    `bson:"group,omitempty" json:"group,omitempty"`
	Question []QuestionnaireResponseQuestionComponent `bson:"question,omitempty" json:"question,omitempty"`
}

type QuestionnaireResponseQuestionComponent struct {
	LinkId string                                         `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Text   string                                         `bson:"text,omitempty" json:"text,omitempty"`
	Answer []QuestionnaireResponseQuestionAnswerComponent `bson:"answer,omitempty" json:"answer,omitempty"`
}

type QuestionnaireResponseQuestionAnswerComponent struct {
	ValueBoolean    *bool                                 `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueDecimal    *float64                              `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueInteger    *int32                                `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate       *FHIRDateTime                         `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime   *FHIRDateTime                         `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueInstant    *FHIRDateTime                         `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueTime       *FHIRDateTime                         `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString     string                                `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueUri        string                                `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueAttachment *Attachment                           `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCoding     *Coding                               `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueQuantity   *Quantity                             `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueReference  *Reference                            `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Group           []QuestionnaireResponseGroupComponent `bson:"group,omitempty" json:"group,omitempty"`
}

type QuestionnaireResponsePlus struct {
	QuestionnaireResponse             `bson:",inline"`
	QuestionnaireResponsePlusIncludes `bson:",inline"`
}

type QuestionnaireResponsePlusIncludes struct {
	IncludedQuestionnaireResources       *[]Questionnaire `bson:"_includedQuestionnaireResources,omitempty"`
	IncludedAuthorPractitionerResources  *[]Practitioner  `bson:"_includedAuthorPractitionerResources,omitempty"`
	IncludedAuthorDeviceResources        *[]Device        `bson:"_includedAuthorDeviceResources,omitempty"`
	IncludedAuthorPatientResources       *[]Patient       `bson:"_includedAuthorPatientResources,omitempty"`
	IncludedAuthorRelatedPersonResources *[]RelatedPerson `bson:"_includedAuthorRelatedPersonResources,omitempty"`
	IncludedPatientResources             *[]Patient       `bson:"_includedPatientResources,omitempty"`
	IncludedEncounterResources           *[]Encounter     `bson:"_includedEncounterResources,omitempty"`
	IncludedSourcePractitionerResources  *[]Practitioner  `bson:"_includedSourcePractitionerResources,omitempty"`
	IncludedSourcePatientResources       *[]Patient       `bson:"_includedSourcePatientResources,omitempty"`
	IncludedSourceRelatedPersonResources *[]RelatedPerson `bson:"_includedSourceRelatedPersonResources,omitempty"`
}

func (q *QuestionnaireResponsePlusIncludes) GetIncludedQuestionnaireResource() (questionnaire *Questionnaire, err error) {
	if q.IncludedQuestionnaireResources == nil {
		err = errors.New("Included questionnaires not requested")
	} else if len(*q.IncludedQuestionnaireResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 questionnaire, but found %d", len(*q.IncludedQuestionnaireResources))
	} else if len(*q.IncludedQuestionnaireResources) == 1 {
		questionnaire = &(*q.IncludedQuestionnaireResources)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusIncludes) GetIncludedAuthorPractitionerResource() (practitioner *Practitioner, err error) {
	if q.IncludedAuthorPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*q.IncludedAuthorPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*q.IncludedAuthorPractitionerResources))
	} else if len(*q.IncludedAuthorPractitionerResources) == 1 {
		practitioner = &(*q.IncludedAuthorPractitionerResources)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusIncludes) GetIncludedAuthorDeviceResource() (device *Device, err error) {
	if q.IncludedAuthorDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*q.IncludedAuthorDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*q.IncludedAuthorDeviceResources))
	} else if len(*q.IncludedAuthorDeviceResources) == 1 {
		device = &(*q.IncludedAuthorDeviceResources)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusIncludes) GetIncludedAuthorPatientResource() (patient *Patient, err error) {
	if q.IncludedAuthorPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*q.IncludedAuthorPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*q.IncludedAuthorPatientResources))
	} else if len(*q.IncludedAuthorPatientResources) == 1 {
		patient = &(*q.IncludedAuthorPatientResources)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusIncludes) GetIncludedAuthorRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if q.IncludedAuthorRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*q.IncludedAuthorRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*q.IncludedAuthorRelatedPersonResources))
	} else if len(*q.IncludedAuthorRelatedPersonResources) == 1 {
		relatedPerson = &(*q.IncludedAuthorRelatedPersonResources)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if q.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*q.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*q.IncludedPatientResources))
	} else if len(*q.IncludedPatientResources) == 1 {
		patient = &(*q.IncludedPatientResources)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if q.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*q.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*q.IncludedEncounterResources))
	} else if len(*q.IncludedEncounterResources) == 1 {
		encounter = &(*q.IncludedEncounterResources)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusIncludes) GetIncludedSourcePractitionerResource() (practitioner *Practitioner, err error) {
	if q.IncludedSourcePractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*q.IncludedSourcePractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*q.IncludedSourcePractitionerResources))
	} else if len(*q.IncludedSourcePractitionerResources) == 1 {
		practitioner = &(*q.IncludedSourcePractitionerResources)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusIncludes) GetIncludedSourcePatientResource() (patient *Patient, err error) {
	if q.IncludedSourcePatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*q.IncludedSourcePatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*q.IncludedSourcePatientResources))
	} else if len(*q.IncludedSourcePatientResources) == 1 {
		patient = &(*q.IncludedSourcePatientResources)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusIncludes) GetIncludedSourceRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if q.IncludedSourceRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*q.IncludedSourceRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*q.IncludedSourceRelatedPersonResources))
	} else if len(*q.IncludedSourceRelatedPersonResources) == 1 {
		relatedPerson = &(*q.IncludedSourceRelatedPersonResources)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.IncludedQuestionnaireResources != nil {
		for _, r := range *q.IncludedQuestionnaireResources {
			resourceMap[r.Id] = &r
		}
	}
	if q.IncludedAuthorPractitionerResources != nil {
		for _, r := range *q.IncludedAuthorPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if q.IncludedAuthorDeviceResources != nil {
		for _, r := range *q.IncludedAuthorDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if q.IncludedAuthorPatientResources != nil {
		for _, r := range *q.IncludedAuthorPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if q.IncludedAuthorRelatedPersonResources != nil {
		for _, r := range *q.IncludedAuthorRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if q.IncludedPatientResources != nil {
		for _, r := range *q.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if q.IncludedEncounterResources != nil {
		for _, r := range *q.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	if q.IncludedSourcePractitionerResources != nil {
		for _, r := range *q.IncludedSourcePractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if q.IncludedSourcePatientResources != nil {
		for _, r := range *q.IncludedSourcePatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if q.IncludedSourceRelatedPersonResources != nil {
		for _, r := range *q.IncludedSourceRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
