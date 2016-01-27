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

type MedicationDispense struct {
	DomainResource            `bson:",inline"`
	Identifier                *Identifier                                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                    string                                         `bson:"status,omitempty" json:"status,omitempty"`
	Patient                   *Reference                                     `bson:"patient,omitempty" json:"patient,omitempty"`
	Dispenser                 *Reference                                     `bson:"dispenser,omitempty" json:"dispenser,omitempty"`
	AuthorizingPrescription   []Reference                                    `bson:"authorizingPrescription,omitempty" json:"authorizingPrescription,omitempty"`
	Type                      *CodeableConcept                               `bson:"type,omitempty" json:"type,omitempty"`
	Quantity                  *Quantity                                      `bson:"quantity,omitempty" json:"quantity,omitempty"`
	DaysSupply                *Quantity                                      `bson:"daysSupply,omitempty" json:"daysSupply,omitempty"`
	MedicationCodeableConcept *CodeableConcept                               `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                                     `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	WhenPrepared              *FHIRDateTime                                  `bson:"whenPrepared,omitempty" json:"whenPrepared,omitempty"`
	WhenHandedOver            *FHIRDateTime                                  `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty"`
	Destination               *Reference                                     `bson:"destination,omitempty" json:"destination,omitempty"`
	Receiver                  []Reference                                    `bson:"receiver,omitempty" json:"receiver,omitempty"`
	Note                      string                                         `bson:"note,omitempty" json:"note,omitempty"`
	DosageInstruction         []MedicationDispenseDosageInstructionComponent `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	Substitution              *MedicationDispenseSubstitutionComponent       `bson:"substitution,omitempty" json:"substitution,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *MedicationDispense) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		MedicationDispense
	}{
		ResourceType:       "MedicationDispense",
		MedicationDispense: *resource,
	}
	return json.Marshal(x)
}

// The "medicationDispense" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type medicationDispense MedicationDispense

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *MedicationDispense) UnmarshalJSON(data []byte) (err error) {
	x2 := medicationDispense{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = MedicationDispense(x2)
	}
	return
}

type MedicationDispenseDosageInstructionComponent struct {
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

type MedicationDispenseSubstitutionComponent struct {
	Type             *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Reason           []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	ResponsibleParty []Reference       `bson:"responsibleParty,omitempty" json:"responsibleParty,omitempty"`
}

type MedicationDispensePlus struct {
	MedicationDispense             `bson:",inline"`
	MedicationDispensePlusIncludes `bson:",inline"`
}

type MedicationDispensePlusIncludes struct {
	IncludedReceiverPractitionerResources *[]Practitioner    `bson:"_includedReceiverPractitionerResources,omitempty"`
	IncludedReceiverPatientResources      *[]Patient         `bson:"_includedReceiverPatientResources,omitempty"`
	IncludedDestinationResources          *[]Location        `bson:"_includedDestinationResources,omitempty"`
	IncludedMedicationResources           *[]Medication      `bson:"_includedMedicationResources,omitempty"`
	IncludedResponsiblepartyResources     *[]Practitioner    `bson:"_includedResponsiblepartyResources,omitempty"`
	IncludedDispenserResources            *[]Practitioner    `bson:"_includedDispenserResources,omitempty"`
	IncludedPrescriptionResources         *[]MedicationOrder `bson:"_includedPrescriptionResources,omitempty"`
	IncludedPatientResources              *[]Patient         `bson:"_includedPatientResources,omitempty"`
}

func (m *MedicationDispensePlusIncludes) GetIncludedReceiverPractitionerResources() (practitioners []Practitioner, err error) {
	if m.IncludedReceiverPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *m.IncludedReceiverPractitionerResources
	}
	return
}

func (m *MedicationDispensePlusIncludes) GetIncludedReceiverPatientResources() (patients []Patient, err error) {
	if m.IncludedReceiverPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *m.IncludedReceiverPatientResources
	}
	return
}

func (m *MedicationDispensePlusIncludes) GetIncludedDestinationResource() (location *Location, err error) {
	if m.IncludedDestinationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*m.IncludedDestinationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*m.IncludedDestinationResources))
	} else if len(*m.IncludedDestinationResources) == 1 {
		location = &(*m.IncludedDestinationResources)[0]
	}
	return
}

func (m *MedicationDispensePlusIncludes) GetIncludedMedicationResource() (medication *Medication, err error) {
	if m.IncludedMedicationResources == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedMedicationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedMedicationResources))
	} else if len(*m.IncludedMedicationResources) == 1 {
		medication = &(*m.IncludedMedicationResources)[0]
	}
	return
}

func (m *MedicationDispensePlusIncludes) GetIncludedResponsiblepartyResources() (practitioners []Practitioner, err error) {
	if m.IncludedResponsiblepartyResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *m.IncludedResponsiblepartyResources
	}
	return
}

func (m *MedicationDispensePlusIncludes) GetIncludedDispenserResource() (practitioner *Practitioner, err error) {
	if m.IncludedDispenserResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedDispenserResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedDispenserResources))
	} else if len(*m.IncludedDispenserResources) == 1 {
		practitioner = &(*m.IncludedDispenserResources)[0]
	}
	return
}

func (m *MedicationDispensePlusIncludes) GetIncludedPrescriptionResources() (medicationOrders []MedicationOrder, err error) {
	if m.IncludedPrescriptionResources == nil {
		err = errors.New("Included medicationOrders not requested")
	} else {
		medicationOrders = *m.IncludedPrescriptionResources
	}
	return
}

func (m *MedicationDispensePlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if m.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResources))
	} else if len(*m.IncludedPatientResources) == 1 {
		patient = &(*m.IncludedPatientResources)[0]
	}
	return
}

func (m *MedicationDispensePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedReceiverPractitionerResources != nil {
		for _, r := range *m.IncludedReceiverPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedReceiverPatientResources != nil {
		for _, r := range *m.IncludedReceiverPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedDestinationResources != nil {
		for _, r := range *m.IncludedDestinationResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedMedicationResources != nil {
		for _, r := range *m.IncludedMedicationResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedResponsiblepartyResources != nil {
		for _, r := range *m.IncludedResponsiblepartyResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedDispenserResources != nil {
		for _, r := range *m.IncludedDispenserResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedPrescriptionResources != nil {
		for _, r := range *m.IncludedPrescriptionResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedPatientResources != nil {
		for _, r := range *m.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
