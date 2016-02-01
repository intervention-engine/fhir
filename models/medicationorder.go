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

type MedicationOrder struct {
	DomainResource            `bson:",inline"`
	Identifier                []Identifier                                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	DateWritten               *FHIRDateTime                               `bson:"dateWritten,omitempty" json:"dateWritten,omitempty"`
	Status                    string                                      `bson:"status,omitempty" json:"status,omitempty"`
	DateEnded                 *FHIRDateTime                               `bson:"dateEnded,omitempty" json:"dateEnded,omitempty"`
	ReasonEnded               *CodeableConcept                            `bson:"reasonEnded,omitempty" json:"reasonEnded,omitempty"`
	Patient                   *Reference                                  `bson:"patient,omitempty" json:"patient,omitempty"`
	Prescriber                *Reference                                  `bson:"prescriber,omitempty" json:"prescriber,omitempty"`
	Encounter                 *Reference                                  `bson:"encounter,omitempty" json:"encounter,omitempty"`
	ReasonCodeableConcept     *CodeableConcept                            `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference           *Reference                                  `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                      string                                      `bson:"note,omitempty" json:"note,omitempty"`
	MedicationCodeableConcept *CodeableConcept                            `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                                  `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	DosageInstruction         []MedicationOrderDosageInstructionComponent `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	DispenseRequest           *MedicationOrderDispenseRequestComponent    `bson:"dispenseRequest,omitempty" json:"dispenseRequest,omitempty"`
	Substitution              *MedicationOrderSubstitutionComponent       `bson:"substitution,omitempty" json:"substitution,omitempty"`
	PriorPrescription         *Reference                                  `bson:"priorPrescription,omitempty" json:"priorPrescription,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *MedicationOrder) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "MedicationOrder"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to MedicationOrder), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *MedicationOrder) GetBSON() (interface{}, error) {
	x.ResourceType = "MedicationOrder"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "medicationOrder" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type medicationOrder MedicationOrder

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *MedicationOrder) UnmarshalJSON(data []byte) (err error) {
	x2 := medicationOrder{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = MedicationOrder(x2)
		return x.checkResourceType()
	}
	return
}

func (x *MedicationOrder) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "MedicationOrder"
	} else if x.ResourceType != "MedicationOrder" {
		return errors.New(fmt.Sprintf("Expected resourceType to be MedicationOrder, instead received %s", x.ResourceType))
	}
	return nil
}

type MedicationOrderDosageInstructionComponent struct {
	Text                    string           `bson:"text,omitempty" json:"text,omitempty"`
	AdditionalInstructions  *CodeableConcept `bson:"additionalInstructions,omitempty" json:"additionalInstructions,omitempty"`
	Timing                  *Timing          `bson:"timing,omitempty" json:"timing,omitempty"`
	AsNeededBoolean         *bool            `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	SiteCodeableConcept     *CodeableConcept `bson:"siteCodeableConcept,omitempty" json:"siteCodeableConcept,omitempty"`
	SiteReference           *Reference       `bson:"siteReference,omitempty" json:"siteReference,omitempty"`
	Route                   *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	Method                  *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	DoseRange               *Range           `bson:"doseRange,omitempty" json:"doseRange,omitempty"`
	DoseSimpleQuantity      *Quantity        `bson:"doseSimpleQuantity,omitempty" json:"doseSimpleQuantity,omitempty"`
	RateRatio               *Ratio           `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateRange               *Range           `bson:"rateRange,omitempty" json:"rateRange,omitempty"`
	MaxDosePerPeriod        *Ratio           `bson:"maxDosePerPeriod,omitempty" json:"maxDosePerPeriod,omitempty"`
}

type MedicationOrderDispenseRequestComponent struct {
	MedicationCodeableConcept *CodeableConcept `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference       `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	ValidityPeriod            *Period          `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
	NumberOfRepeatsAllowed    *uint32          `bson:"numberOfRepeatsAllowed,omitempty" json:"numberOfRepeatsAllowed,omitempty"`
	Quantity                  *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ExpectedSupplyDuration    *Quantity        `bson:"expectedSupplyDuration,omitempty" json:"expectedSupplyDuration,omitempty"`
}

type MedicationOrderSubstitutionComponent struct {
	Type   *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Reason *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}

type MedicationOrderPlus struct {
	MedicationOrder             `bson:",inline"`
	MedicationOrderPlusIncludes `bson:",inline"`
}

type MedicationOrderPlusIncludes struct {
	IncludedPrescriberResources *[]Practitioner `bson:"_includedPrescriberResources,omitempty"`
	IncludedPatientResources    *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedMedicationResources *[]Medication   `bson:"_includedMedicationResources,omitempty"`
	IncludedEncounterResources  *[]Encounter    `bson:"_includedEncounterResources,omitempty"`
}

func (m *MedicationOrderPlusIncludes) GetIncludedPrescriberResource() (practitioner *Practitioner, err error) {
	if m.IncludedPrescriberResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedPrescriberResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedPrescriberResources))
	} else if len(*m.IncludedPrescriberResources) == 1 {
		practitioner = &(*m.IncludedPrescriberResources)[0]
	}
	return
}

func (m *MedicationOrderPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if m.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResources))
	} else if len(*m.IncludedPatientResources) == 1 {
		patient = &(*m.IncludedPatientResources)[0]
	}
	return
}

func (m *MedicationOrderPlusIncludes) GetIncludedMedicationResource() (medication *Medication, err error) {
	if m.IncludedMedicationResources == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedMedicationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedMedicationResources))
	} else if len(*m.IncludedMedicationResources) == 1 {
		medication = &(*m.IncludedMedicationResources)[0]
	}
	return
}

func (m *MedicationOrderPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if m.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*m.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*m.IncludedEncounterResources))
	} else if len(*m.IncludedEncounterResources) == 1 {
		encounter = &(*m.IncludedEncounterResources)[0]
	}
	return
}

func (m *MedicationOrderPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedPrescriberResources != nil {
		for _, r := range *m.IncludedPrescriberResources {
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
	return resourceMap
}
