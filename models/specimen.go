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

type Specimen struct {
	DomainResource      `bson:",inline"`
	Identifier          []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              string                       `bson:"status,omitempty" json:"status,omitempty"`
	Type                *CodeableConcept             `bson:"type,omitempty" json:"type,omitempty"`
	Parent              []Reference                  `bson:"parent,omitempty" json:"parent,omitempty"`
	Subject             *Reference                   `bson:"subject,omitempty" json:"subject,omitempty"`
	AccessionIdentifier *Identifier                  `bson:"accessionIdentifier,omitempty" json:"accessionIdentifier,omitempty"`
	ReceivedTime        *FHIRDateTime                `bson:"receivedTime,omitempty" json:"receivedTime,omitempty"`
	Collection          *SpecimenCollectionComponent `bson:"collection,omitempty" json:"collection,omitempty"`
	Treatment           []SpecimenTreatmentComponent `bson:"treatment,omitempty" json:"treatment,omitempty"`
	Container           []SpecimenContainerComponent `bson:"container,omitempty" json:"container,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Specimen) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Specimen
	}{
		ResourceType: "Specimen",
		Specimen:     *resource,
	}
	return json.Marshal(x)
}

// The "specimen" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type specimen Specimen

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Specimen) UnmarshalJSON(data []byte) (err error) {
	x2 := specimen{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Specimen(x2)
	}
	return
}

type SpecimenCollectionComponent struct {
	Collector         *Reference       `bson:"collector,omitempty" json:"collector,omitempty"`
	Comment           []string         `bson:"comment,omitempty" json:"comment,omitempty"`
	CollectedDateTime *FHIRDateTime    `bson:"collectedDateTime,omitempty" json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period          `bson:"collectedPeriod,omitempty" json:"collectedPeriod,omitempty"`
	Quantity          *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Method            *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	BodySite          *CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
}

type SpecimenTreatmentComponent struct {
	Description string           `bson:"description,omitempty" json:"description,omitempty"`
	Procedure   *CodeableConcept `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Additive    []Reference      `bson:"additive,omitempty" json:"additive,omitempty"`
}

type SpecimenContainerComponent struct {
	Identifier              []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Description             string           `bson:"description,omitempty" json:"description,omitempty"`
	Type                    *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Capacity                *Quantity        `bson:"capacity,omitempty" json:"capacity,omitempty"`
	SpecimenQuantity        *Quantity        `bson:"specimenQuantity,omitempty" json:"specimenQuantity,omitempty"`
	AdditiveCodeableConcept *CodeableConcept `bson:"additiveCodeableConcept,omitempty" json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *Reference       `bson:"additiveReference,omitempty" json:"additiveReference,omitempty"`
}

type SpecimenPlus struct {
	Specimen             `bson:",inline"`
	SpecimenPlusIncludes `bson:",inline"`
}

type SpecimenPlusIncludes struct {
	IncludedParentResources           *[]Specimen     `bson:"_includedParentResources,omitempty"`
	IncludedSubjectGroupResources     *[]Group        `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectDeviceResources    *[]Device       `bson:"_includedSubjectDeviceResources,omitempty"`
	IncludedSubjectPatientResources   *[]Patient      `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedSubjectSubstanceResources *[]Substance    `bson:"_includedSubjectSubstanceResources,omitempty"`
	IncludedPatientResources          *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedCollectorResources        *[]Practitioner `bson:"_includedCollectorResources,omitempty"`
}

func (s *SpecimenPlusIncludes) GetIncludedParentResources() (specimen []Specimen, err error) {
	if s.IncludedParentResources == nil {
		err = errors.New("Included specimen not requested")
	} else {
		specimen = *s.IncludedParentResources
	}
	return
}

func (s *SpecimenPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if s.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*s.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*s.IncludedSubjectGroupResources))
	} else if len(*s.IncludedSubjectGroupResources) == 1 {
		group = &(*s.IncludedSubjectGroupResources)[0]
	}
	return
}

func (s *SpecimenPlusIncludes) GetIncludedSubjectDeviceResource() (device *Device, err error) {
	if s.IncludedSubjectDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*s.IncludedSubjectDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*s.IncludedSubjectDeviceResources))
	} else if len(*s.IncludedSubjectDeviceResources) == 1 {
		device = &(*s.IncludedSubjectDeviceResources)[0]
	}
	return
}

func (s *SpecimenPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if s.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedSubjectPatientResources))
	} else if len(*s.IncludedSubjectPatientResources) == 1 {
		patient = &(*s.IncludedSubjectPatientResources)[0]
	}
	return
}

func (s *SpecimenPlusIncludes) GetIncludedSubjectSubstanceResource() (substance *Substance, err error) {
	if s.IncludedSubjectSubstanceResources == nil {
		err = errors.New("Included substances not requested")
	} else if len(*s.IncludedSubjectSubstanceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*s.IncludedSubjectSubstanceResources))
	} else if len(*s.IncludedSubjectSubstanceResources) == 1 {
		substance = &(*s.IncludedSubjectSubstanceResources)[0]
	}
	return
}

func (s *SpecimenPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if s.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedPatientResources))
	} else if len(*s.IncludedPatientResources) == 1 {
		patient = &(*s.IncludedPatientResources)[0]
	}
	return
}

func (s *SpecimenPlusIncludes) GetIncludedCollectorResource() (practitioner *Practitioner, err error) {
	if s.IncludedCollectorResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*s.IncludedCollectorResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*s.IncludedCollectorResources))
	} else if len(*s.IncludedCollectorResources) == 1 {
		practitioner = &(*s.IncludedCollectorResources)[0]
	}
	return
}

func (s *SpecimenPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedParentResources != nil {
		for _, r := range *s.IncludedParentResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedSubjectGroupResources != nil {
		for _, r := range *s.IncludedSubjectGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedSubjectDeviceResources != nil {
		for _, r := range *s.IncludedSubjectDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedSubjectPatientResources != nil {
		for _, r := range *s.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedSubjectSubstanceResources != nil {
		for _, r := range *s.IncludedSubjectSubstanceResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedPatientResources != nil {
		for _, r := range *s.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedCollectorResources != nil {
		for _, r := range *s.IncludedCollectorResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
