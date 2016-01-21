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

type MedicationAdministration struct {
	DomainResource            `bson:",inline"`
	Identifier                []Identifier                             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                    string                                   `bson:"status,omitempty" json:"status,omitempty"`
	Patient                   *Reference                               `bson:"patient,omitempty" json:"patient,omitempty"`
	Practitioner              *Reference                               `bson:"practitioner,omitempty" json:"practitioner,omitempty"`
	Encounter                 *Reference                               `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Prescription              *Reference                               `bson:"prescription,omitempty" json:"prescription,omitempty"`
	WasNotGiven               *bool                                    `bson:"wasNotGiven,omitempty" json:"wasNotGiven,omitempty"`
	ReasonNotGiven            []CodeableConcept                        `bson:"reasonNotGiven,omitempty" json:"reasonNotGiven,omitempty"`
	ReasonGiven               []CodeableConcept                        `bson:"reasonGiven,omitempty" json:"reasonGiven,omitempty"`
	EffectiveTimeDateTime     *FHIRDateTime                            `bson:"effectiveTimeDateTime,omitempty" json:"effectiveTimeDateTime,omitempty"`
	EffectiveTimePeriod       *Period                                  `bson:"effectiveTimePeriod,omitempty" json:"effectiveTimePeriod,omitempty"`
	MedicationCodeableConcept *CodeableConcept                         `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                               `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Device                    []Reference                              `bson:"device,omitempty" json:"device,omitempty"`
	Note                      string                                   `bson:"note,omitempty" json:"note,omitempty"`
	Dosage                    *MedicationAdministrationDosageComponent `bson:"dosage,omitempty" json:"dosage,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *MedicationAdministration) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		MedicationAdministration
	}{
		ResourceType:             "MedicationAdministration",
		MedicationAdministration: *resource,
	}
	return json.Marshal(x)
}

// The "medicationAdministration" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type medicationAdministration MedicationAdministration

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *MedicationAdministration) UnmarshalJSON(data []byte) (err error) {
	x2 := medicationAdministration{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = MedicationAdministration(x2)
	}
	return
}

type MedicationAdministrationDosageComponent struct {
	Text                string           `bson:"text,omitempty" json:"text,omitempty"`
	SiteCodeableConcept *CodeableConcept `bson:"siteCodeableConcept,omitempty" json:"siteCodeableConcept,omitempty"`
	SiteReference       *Reference       `bson:"siteReference,omitempty" json:"siteReference,omitempty"`
	Route               *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	Method              *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	Quantity            *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	RateRatio           *Ratio           `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateRange           *Range           `bson:"rateRange,omitempty" json:"rateRange,omitempty"`
}

type MedicationAdministrationPlus struct {
	MedicationAdministration             `bson:",inline"`
	MedicationAdministrationPlusIncludes `bson:",inline"`
}

type MedicationAdministrationPlusIncludes struct {
	IncludedPrescriptionResources              *[]MedicationOrder `bson:"_includedPrescriptionResources,omitempty"`
	IncludedPractitionerPractitionerResources  *[]Practitioner    `bson:"_includedPractitionerPractitionerResources,omitempty"`
	IncludedPractitionerPatientResources       *[]Patient         `bson:"_includedPractitionerPatientResources,omitempty"`
	IncludedPractitionerRelatedPersonResources *[]RelatedPerson   `bson:"_includedPractitionerRelatedPersonResources,omitempty"`
	IncludedPatientResources                   *[]Patient         `bson:"_includedPatientResources,omitempty"`
	IncludedMedicationResources                *[]Medication      `bson:"_includedMedicationResources,omitempty"`
	IncludedEncounterResources                 *[]Encounter       `bson:"_includedEncounterResources,omitempty"`
	IncludedDeviceResources                    *[]Device          `bson:"_includedDeviceResources,omitempty"`
}

func (m *MedicationAdministrationPlusIncludes) GetIncludedPrescriptionResource() (medicationOrder *MedicationOrder, err error) {
	if m.IncludedPrescriptionResources == nil {
		err = errors.New("Included medicationorders not requested")
	} else if len(*m.IncludedPrescriptionResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medicationOrder, but found %d", len(*m.IncludedPrescriptionResources))
	} else if len(*m.IncludedPrescriptionResources) == 1 {
		medicationOrder = &(*m.IncludedPrescriptionResources)[0]
	}
	return
}

func (m *MedicationAdministrationPlusIncludes) GetIncludedPractitionerPractitionerResource() (practitioner *Practitioner, err error) {
	if m.IncludedPractitionerPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedPractitionerPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedPractitionerPractitionerResources))
	} else if len(*m.IncludedPractitionerPractitionerResources) == 1 {
		practitioner = &(*m.IncludedPractitionerPractitionerResources)[0]
	}
	return
}

func (m *MedicationAdministrationPlusIncludes) GetIncludedPractitionerPatientResource() (patient *Patient, err error) {
	if m.IncludedPractitionerPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPractitionerPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPractitionerPatientResources))
	} else if len(*m.IncludedPractitionerPatientResources) == 1 {
		patient = &(*m.IncludedPractitionerPatientResources)[0]
	}
	return
}

func (m *MedicationAdministrationPlusIncludes) GetIncludedPractitionerRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if m.IncludedPractitionerRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*m.IncludedPractitionerRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*m.IncludedPractitionerRelatedPersonResources))
	} else if len(*m.IncludedPractitionerRelatedPersonResources) == 1 {
		relatedPerson = &(*m.IncludedPractitionerRelatedPersonResources)[0]
	}
	return
}

func (m *MedicationAdministrationPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if m.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResources))
	} else if len(*m.IncludedPatientResources) == 1 {
		patient = &(*m.IncludedPatientResources)[0]
	}
	return
}

func (m *MedicationAdministrationPlusIncludes) GetIncludedMedicationResource() (medication *Medication, err error) {
	if m.IncludedMedicationResources == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedMedicationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedMedicationResources))
	} else if len(*m.IncludedMedicationResources) == 1 {
		medication = &(*m.IncludedMedicationResources)[0]
	}
	return
}

func (m *MedicationAdministrationPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if m.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*m.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*m.IncludedEncounterResources))
	} else if len(*m.IncludedEncounterResources) == 1 {
		encounter = &(*m.IncludedEncounterResources)[0]
	}
	return
}

func (m *MedicationAdministrationPlusIncludes) GetIncludedDeviceResources() (devices []Device, err error) {
	if m.IncludedDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *m.IncludedDeviceResources
	}
	return
}

func (m *MedicationAdministrationPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedPrescriptionResources != nil {
		for _, r := range *m.IncludedPrescriptionResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedPractitionerPractitionerResources != nil {
		for _, r := range *m.IncludedPractitionerPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedPractitionerPatientResources != nil {
		for _, r := range *m.IncludedPractitionerPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedPractitionerRelatedPersonResources != nil {
		for _, r := range *m.IncludedPractitionerRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
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
	if m.IncludedEncounterResources != nil {
		for _, r := range *m.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedDeviceResources != nil {
		for _, r := range *m.IncludedDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
