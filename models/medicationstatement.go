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

type MedicationStatement struct {
	DomainResource              `bson:",inline"`
	Identifier                  []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient                     *Reference                           `bson:"patient,omitempty" json:"patient,omitempty"`
	InformationSource           *Reference                           `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	DateAsserted                *FHIRDateTime                        `bson:"dateAsserted,omitempty" json:"dateAsserted,omitempty"`
	Status                      string                               `bson:"status,omitempty" json:"status,omitempty"`
	WasNotTaken                 *bool                                `bson:"wasNotTaken,omitempty" json:"wasNotTaken,omitempty"`
	ReasonNotTaken              []CodeableConcept                    `bson:"reasonNotTaken,omitempty" json:"reasonNotTaken,omitempty"`
	ReasonForUseCodeableConcept *CodeableConcept                     `bson:"reasonForUseCodeableConcept,omitempty" json:"reasonForUseCodeableConcept,omitempty"`
	ReasonForUseReference       *Reference                           `bson:"reasonForUseReference,omitempty" json:"reasonForUseReference,omitempty"`
	EffectiveDateTime           *FHIRDateTime                        `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod             *Period                              `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Note                        string                               `bson:"note,omitempty" json:"note,omitempty"`
	SupportingInformation       []Reference                          `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	MedicationCodeableConcept   *CodeableConcept                     `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference         *Reference                           `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Dosage                      []MedicationStatementDosageComponent `bson:"dosage,omitempty" json:"dosage,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *MedicationStatement) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "MedicationStatement"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to MedicationStatement), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *MedicationStatement) GetBSON() (interface{}, error) {
	x.ResourceType = "MedicationStatement"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "medicationStatement" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type medicationStatement MedicationStatement

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *MedicationStatement) UnmarshalJSON(data []byte) (err error) {
	x2 := medicationStatement{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = MedicationStatement(x2)
		return x.checkResourceType()
	}
	return
}

func (x *MedicationStatement) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "MedicationStatement"
	} else if x.ResourceType != "MedicationStatement" {
		return errors.New(fmt.Sprintf("Expected resourceType to be MedicationStatement, instead received %s", x.ResourceType))
	}
	return nil
}

type MedicationStatementDosageComponent struct {
	Text                    string           `bson:"text,omitempty" json:"text,omitempty"`
	Timing                  *Timing          `bson:"timing,omitempty" json:"timing,omitempty"`
	AsNeededBoolean         *bool            `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	SiteCodeableConcept     *CodeableConcept `bson:"siteCodeableConcept,omitempty" json:"siteCodeableConcept,omitempty"`
	SiteReference           *Reference       `bson:"siteReference,omitempty" json:"siteReference,omitempty"`
	Route                   *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	Method                  *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	QuantitySimpleQuantity  *Quantity        `bson:"quantitySimpleQuantity,omitempty" json:"quantitySimpleQuantity,omitempty"`
	QuantityRange           *Range           `bson:"quantityRange,omitempty" json:"quantityRange,omitempty"`
	RateRatio               *Ratio           `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateRange               *Range           `bson:"rateRange,omitempty" json:"rateRange,omitempty"`
	MaxDosePerPeriod        *Ratio           `bson:"maxDosePerPeriod,omitempty" json:"maxDosePerPeriod,omitempty"`
}

type MedicationStatementPlus struct {
	MedicationStatement             `bson:",inline"`
	MedicationStatementPlusIncludes `bson:",inline"`
}

type MedicationStatementPlusIncludes struct {
	IncludedPatientResources             *[]Patient       `bson:"_includedPatientResources,omitempty"`
	IncludedMedicationResources          *[]Medication    `bson:"_includedMedicationResources,omitempty"`
	IncludedSourcePractitionerResources  *[]Practitioner  `bson:"_includedSourcePractitionerResources,omitempty"`
	IncludedSourcePatientResources       *[]Patient       `bson:"_includedSourcePatientResources,omitempty"`
	IncludedSourceRelatedPersonResources *[]RelatedPerson `bson:"_includedSourceRelatedPersonResources,omitempty"`
}

func (m *MedicationStatementPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if m.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResources))
	} else if len(*m.IncludedPatientResources) == 1 {
		patient = &(*m.IncludedPatientResources)[0]
	}
	return
}

func (m *MedicationStatementPlusIncludes) GetIncludedMedicationResource() (medication *Medication, err error) {
	if m.IncludedMedicationResources == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedMedicationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedMedicationResources))
	} else if len(*m.IncludedMedicationResources) == 1 {
		medication = &(*m.IncludedMedicationResources)[0]
	}
	return
}

func (m *MedicationStatementPlusIncludes) GetIncludedSourcePractitionerResource() (practitioner *Practitioner, err error) {
	if m.IncludedSourcePractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedSourcePractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedSourcePractitionerResources))
	} else if len(*m.IncludedSourcePractitionerResources) == 1 {
		practitioner = &(*m.IncludedSourcePractitionerResources)[0]
	}
	return
}

func (m *MedicationStatementPlusIncludes) GetIncludedSourcePatientResource() (patient *Patient, err error) {
	if m.IncludedSourcePatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedSourcePatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedSourcePatientResources))
	} else if len(*m.IncludedSourcePatientResources) == 1 {
		patient = &(*m.IncludedSourcePatientResources)[0]
	}
	return
}

func (m *MedicationStatementPlusIncludes) GetIncludedSourceRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if m.IncludedSourceRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*m.IncludedSourceRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*m.IncludedSourceRelatedPersonResources))
	} else if len(*m.IncludedSourceRelatedPersonResources) == 1 {
		relatedPerson = &(*m.IncludedSourceRelatedPersonResources)[0]
	}
	return
}

func (m *MedicationStatementPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedPatientResources != nil {
		for _, r := range *m.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedMedicationResources != nil {
		for _, r := range *m.IncludedMedicationResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedSourcePractitionerResources != nil {
		for _, r := range *m.IncludedSourcePractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedSourcePatientResources != nil {
		for _, r := range *m.IncludedSourcePatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedSourceRelatedPersonResources != nil {
		for _, r := range *m.IncludedSourceRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
