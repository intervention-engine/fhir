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

type AllergyIntolerance struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Onset          *FHIRDateTime                         `bson:"onset,omitempty" json:"onset,omitempty"`
	RecordedDate   *FHIRDateTime                         `bson:"recordedDate,omitempty" json:"recordedDate,omitempty"`
	Recorder       *Reference                            `bson:"recorder,omitempty" json:"recorder,omitempty"`
	Patient        *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	Reporter       *Reference                            `bson:"reporter,omitempty" json:"reporter,omitempty"`
	Substance      *CodeableConcept                      `bson:"substance,omitempty" json:"substance,omitempty"`
	Status         string                                `bson:"status,omitempty" json:"status,omitempty"`
	Criticality    string                                `bson:"criticality,omitempty" json:"criticality,omitempty"`
	Type           string                                `bson:"type,omitempty" json:"type,omitempty"`
	Category       string                                `bson:"category,omitempty" json:"category,omitempty"`
	LastOccurence  *FHIRDateTime                         `bson:"lastOccurence,omitempty" json:"lastOccurence,omitempty"`
	Note           *Annotation                           `bson:"note,omitempty" json:"note,omitempty"`
	Reaction       []AllergyIntoleranceReactionComponent `bson:"reaction,omitempty" json:"reaction,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *AllergyIntolerance) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		AllergyIntolerance
	}{
		ResourceType:       "AllergyIntolerance",
		AllergyIntolerance: *resource,
	}
	return json.Marshal(x)
}

// The "allergyIntolerance" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type allergyIntolerance AllergyIntolerance

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *AllergyIntolerance) UnmarshalJSON(data []byte) (err error) {
	x2 := allergyIntolerance{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = AllergyIntolerance(x2)
	}
	return
}

type AllergyIntoleranceReactionComponent struct {
	Substance     *CodeableConcept  `bson:"substance,omitempty" json:"substance,omitempty"`
	Certainty     string            `bson:"certainty,omitempty" json:"certainty,omitempty"`
	Manifestation []CodeableConcept `bson:"manifestation,omitempty" json:"manifestation,omitempty"`
	Description   string            `bson:"description,omitempty" json:"description,omitempty"`
	Onset         *FHIRDateTime     `bson:"onset,omitempty" json:"onset,omitempty"`
	Severity      string            `bson:"severity,omitempty" json:"severity,omitempty"`
	ExposureRoute *CodeableConcept  `bson:"exposureRoute,omitempty" json:"exposureRoute,omitempty"`
	Note          *Annotation       `bson:"note,omitempty" json:"note,omitempty"`
}

type AllergyIntolerancePlus struct {
	AllergyIntolerance             `bson:",inline"`
	AllergyIntolerancePlusIncludes `bson:",inline"`
}

type AllergyIntolerancePlusIncludes struct {
	IncludedRecorderPractitionerResources  *[]Practitioner  `bson:"_includedRecorderPractitionerResources,omitempty"`
	IncludedRecorderPatientResources       *[]Patient       `bson:"_includedRecorderPatientResources,omitempty"`
	IncludedReporterPractitionerResources  *[]Practitioner  `bson:"_includedReporterPractitionerResources,omitempty"`
	IncludedReporterPatientResources       *[]Patient       `bson:"_includedReporterPatientResources,omitempty"`
	IncludedReporterRelatedPersonResources *[]RelatedPerson `bson:"_includedReporterRelatedPersonResources,omitempty"`
	IncludedPatientResources               *[]Patient       `bson:"_includedPatientResources,omitempty"`
}

func (a *AllergyIntolerancePlusIncludes) GetIncludedRecorderPractitionerResource() (practitioner *Practitioner, err error) {
	if a.IncludedRecorderPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedRecorderPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedRecorderPractitionerResources))
	} else if len(*a.IncludedRecorderPractitionerResources) == 1 {
		practitioner = &(*a.IncludedRecorderPractitionerResources)[0]
	}
	return
}

func (a *AllergyIntolerancePlusIncludes) GetIncludedRecorderPatientResource() (patient *Patient, err error) {
	if a.IncludedRecorderPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedRecorderPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedRecorderPatientResources))
	} else if len(*a.IncludedRecorderPatientResources) == 1 {
		patient = &(*a.IncludedRecorderPatientResources)[0]
	}
	return
}

func (a *AllergyIntolerancePlusIncludes) GetIncludedReporterPractitionerResource() (practitioner *Practitioner, err error) {
	if a.IncludedReporterPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedReporterPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedReporterPractitionerResources))
	} else if len(*a.IncludedReporterPractitionerResources) == 1 {
		practitioner = &(*a.IncludedReporterPractitionerResources)[0]
	}
	return
}

func (a *AllergyIntolerancePlusIncludes) GetIncludedReporterPatientResource() (patient *Patient, err error) {
	if a.IncludedReporterPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedReporterPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedReporterPatientResources))
	} else if len(*a.IncludedReporterPatientResources) == 1 {
		patient = &(*a.IncludedReporterPatientResources)[0]
	}
	return
}

func (a *AllergyIntolerancePlusIncludes) GetIncludedReporterRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedReporterRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedReporterRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedReporterRelatedPersonResources))
	} else if len(*a.IncludedReporterRelatedPersonResources) == 1 {
		relatedPerson = &(*a.IncludedReporterRelatedPersonResources)[0]
	}
	return
}

func (a *AllergyIntolerancePlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if a.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResources))
	} else if len(*a.IncludedPatientResources) == 1 {
		patient = &(*a.IncludedPatientResources)[0]
	}
	return
}

func (a *AllergyIntolerancePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedRecorderPractitionerResources != nil {
		for _, r := range *a.IncludedRecorderPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedRecorderPatientResources != nil {
		for _, r := range *a.IncludedRecorderPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedReporterPractitionerResources != nil {
		for _, r := range *a.IncludedReporterPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedReporterPatientResources != nil {
		for _, r := range *a.IncludedReporterPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedReporterRelatedPersonResources != nil {
		for _, r := range *a.IncludedReporterRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedPatientResources != nil {
		for _, r := range *a.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
