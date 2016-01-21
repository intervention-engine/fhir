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

type Immunization struct {
	DomainResource      `bson:",inline"`
	Identifier          []Identifier                               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              string                                     `bson:"status,omitempty" json:"status,omitempty"`
	Date                *FHIRDateTime                              `bson:"date,omitempty" json:"date,omitempty"`
	VaccineCode         *CodeableConcept                           `bson:"vaccineCode,omitempty" json:"vaccineCode,omitempty"`
	Patient             *Reference                                 `bson:"patient,omitempty" json:"patient,omitempty"`
	WasNotGiven         *bool                                      `bson:"wasNotGiven,omitempty" json:"wasNotGiven,omitempty"`
	Reported            *bool                                      `bson:"reported,omitempty" json:"reported,omitempty"`
	Performer           *Reference                                 `bson:"performer,omitempty" json:"performer,omitempty"`
	Requester           *Reference                                 `bson:"requester,omitempty" json:"requester,omitempty"`
	Encounter           *Reference                                 `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Manufacturer        *Reference                                 `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Location            *Reference                                 `bson:"location,omitempty" json:"location,omitempty"`
	LotNumber           string                                     `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	ExpirationDate      *FHIRDateTime                              `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	Site                *CodeableConcept                           `bson:"site,omitempty" json:"site,omitempty"`
	Route               *CodeableConcept                           `bson:"route,omitempty" json:"route,omitempty"`
	DoseQuantity        *Quantity                                  `bson:"doseQuantity,omitempty" json:"doseQuantity,omitempty"`
	Note                []Annotation                               `bson:"note,omitempty" json:"note,omitempty"`
	Explanation         *ImmunizationExplanationComponent          `bson:"explanation,omitempty" json:"explanation,omitempty"`
	Reaction            []ImmunizationReactionComponent            `bson:"reaction,omitempty" json:"reaction,omitempty"`
	VaccinationProtocol []ImmunizationVaccinationProtocolComponent `bson:"vaccinationProtocol,omitempty" json:"vaccinationProtocol,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Immunization) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Immunization
	}{
		ResourceType: "Immunization",
		Immunization: *resource,
	}
	return json.Marshal(x)
}

// The "immunization" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type immunization Immunization

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Immunization) UnmarshalJSON(data []byte) (err error) {
	x2 := immunization{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Immunization(x2)
	}
	return
}

type ImmunizationExplanationComponent struct {
	Reason         []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	ReasonNotGiven []CodeableConcept `bson:"reasonNotGiven,omitempty" json:"reasonNotGiven,omitempty"`
}

type ImmunizationReactionComponent struct {
	Date     *FHIRDateTime `bson:"date,omitempty" json:"date,omitempty"`
	Detail   *Reference    `bson:"detail,omitempty" json:"detail,omitempty"`
	Reported *bool         `bson:"reported,omitempty" json:"reported,omitempty"`
}

type ImmunizationVaccinationProtocolComponent struct {
	DoseSequence     *uint32           `bson:"doseSequence,omitempty" json:"doseSequence,omitempty"`
	Description      string            `bson:"description,omitempty" json:"description,omitempty"`
	Authority        *Reference        `bson:"authority,omitempty" json:"authority,omitempty"`
	Series           string            `bson:"series,omitempty" json:"series,omitempty"`
	SeriesDoses      *uint32           `bson:"seriesDoses,omitempty" json:"seriesDoses,omitempty"`
	TargetDisease    []CodeableConcept `bson:"targetDisease,omitempty" json:"targetDisease,omitempty"`
	DoseStatus       *CodeableConcept  `bson:"doseStatus,omitempty" json:"doseStatus,omitempty"`
	DoseStatusReason *CodeableConcept  `bson:"doseStatusReason,omitempty" json:"doseStatusReason,omitempty"`
}

type ImmunizationPlus struct {
	Immunization             `bson:",inline"`
	ImmunizationPlusIncludes `bson:",inline"`
}

type ImmunizationPlusIncludes struct {
	IncludedRequesterResources    *[]Practitioner `bson:"_includedRequesterResources,omitempty"`
	IncludedPerformerResources    *[]Practitioner `bson:"_includedPerformerResources,omitempty"`
	IncludedReactionResources     *[]Observation  `bson:"_includedReactionResources,omitempty"`
	IncludedManufacturerResources *[]Organization `bson:"_includedManufacturerResources,omitempty"`
	IncludedPatientResources      *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedLocationResources     *[]Location     `bson:"_includedLocationResources,omitempty"`
}

func (i *ImmunizationPlusIncludes) GetIncludedRequesterResource() (practitioner *Practitioner, err error) {
	if i.IncludedRequesterResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*i.IncludedRequesterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*i.IncludedRequesterResources))
	} else if len(*i.IncludedRequesterResources) == 1 {
		practitioner = &(*i.IncludedRequesterResources)[0]
	}
	return
}

func (i *ImmunizationPlusIncludes) GetIncludedPerformerResource() (practitioner *Practitioner, err error) {
	if i.IncludedPerformerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*i.IncludedPerformerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*i.IncludedPerformerResources))
	} else if len(*i.IncludedPerformerResources) == 1 {
		practitioner = &(*i.IncludedPerformerResources)[0]
	}
	return
}

func (i *ImmunizationPlusIncludes) GetIncludedReactionResource() (observation *Observation, err error) {
	if i.IncludedReactionResources == nil {
		err = errors.New("Included observations not requested")
	} else if len(*i.IncludedReactionResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 observation, but found %d", len(*i.IncludedReactionResources))
	} else if len(*i.IncludedReactionResources) == 1 {
		observation = &(*i.IncludedReactionResources)[0]
	}
	return
}

func (i *ImmunizationPlusIncludes) GetIncludedManufacturerResource() (organization *Organization, err error) {
	if i.IncludedManufacturerResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*i.IncludedManufacturerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*i.IncludedManufacturerResources))
	} else if len(*i.IncludedManufacturerResources) == 1 {
		organization = &(*i.IncludedManufacturerResources)[0]
	}
	return
}

func (i *ImmunizationPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if i.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedPatientResources))
	} else if len(*i.IncludedPatientResources) == 1 {
		patient = &(*i.IncludedPatientResources)[0]
	}
	return
}

func (i *ImmunizationPlusIncludes) GetIncludedLocationResource() (location *Location, err error) {
	if i.IncludedLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*i.IncludedLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*i.IncludedLocationResources))
	} else if len(*i.IncludedLocationResources) == 1 {
		location = &(*i.IncludedLocationResources)[0]
	}
	return
}

func (i *ImmunizationPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedRequesterResources != nil {
		for _, r := range *i.IncludedRequesterResources {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedPerformerResources != nil {
		for _, r := range *i.IncludedPerformerResources {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedReactionResources != nil {
		for _, r := range *i.IncludedReactionResources {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedManufacturerResources != nil {
		for _, r := range *i.IncludedManufacturerResources {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedPatientResources != nil {
		for _, r := range *i.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedLocationResources != nil {
		for _, r := range *i.IncludedLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
