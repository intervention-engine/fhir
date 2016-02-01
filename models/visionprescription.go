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

type VisionPrescription struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	DateWritten           *FHIRDateTime                         `bson:"dateWritten,omitempty" json:"dateWritten,omitempty"`
	Patient               *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	Prescriber            *Reference                            `bson:"prescriber,omitempty" json:"prescriber,omitempty"`
	Encounter             *Reference                            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	ReasonCodeableConcept *CodeableConcept                      `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                            `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Dispense              []VisionPrescriptionDispenseComponent `bson:"dispense,omitempty" json:"dispense,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *VisionPrescription) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "VisionPrescription"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to VisionPrescription), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *VisionPrescription) GetBSON() (interface{}, error) {
	x.ResourceType = "VisionPrescription"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "visionPrescription" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type visionPrescription VisionPrescription

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *VisionPrescription) UnmarshalJSON(data []byte) (err error) {
	x2 := visionPrescription{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = VisionPrescription(x2)
		return x.checkResourceType()
	}
	return
}

func (x *VisionPrescription) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "VisionPrescription"
	} else if x.ResourceType != "VisionPrescription" {
		return errors.New(fmt.Sprintf("Expected resourceType to be VisionPrescription, instead received %s", x.ResourceType))
	}
	return nil
}

type VisionPrescriptionDispenseComponent struct {
	Product   *Coding   `bson:"product,omitempty" json:"product,omitempty"`
	Eye       string    `bson:"eye,omitempty" json:"eye,omitempty"`
	Sphere    *float64  `bson:"sphere,omitempty" json:"sphere,omitempty"`
	Cylinder  *float64  `bson:"cylinder,omitempty" json:"cylinder,omitempty"`
	Axis      *int32    `bson:"axis,omitempty" json:"axis,omitempty"`
	Prism     *float64  `bson:"prism,omitempty" json:"prism,omitempty"`
	Base      string    `bson:"base,omitempty" json:"base,omitempty"`
	Add       *float64  `bson:"add,omitempty" json:"add,omitempty"`
	Power     *float64  `bson:"power,omitempty" json:"power,omitempty"`
	BackCurve *float64  `bson:"backCurve,omitempty" json:"backCurve,omitempty"`
	Diameter  *float64  `bson:"diameter,omitempty" json:"diameter,omitempty"`
	Duration  *Quantity `bson:"duration,omitempty" json:"duration,omitempty"`
	Color     string    `bson:"color,omitempty" json:"color,omitempty"`
	Brand     string    `bson:"brand,omitempty" json:"brand,omitempty"`
	Notes     string    `bson:"notes,omitempty" json:"notes,omitempty"`
}

type VisionPrescriptionPlus struct {
	VisionPrescription             `bson:",inline"`
	VisionPrescriptionPlusIncludes `bson:",inline"`
}

type VisionPrescriptionPlusIncludes struct {
	IncludedPrescriberResources *[]Practitioner `bson:"_includedPrescriberResources,omitempty"`
	IncludedPatientResources    *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedEncounterResources  *[]Encounter    `bson:"_includedEncounterResources,omitempty"`
}

func (v *VisionPrescriptionPlusIncludes) GetIncludedPrescriberResource() (practitioner *Practitioner, err error) {
	if v.IncludedPrescriberResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*v.IncludedPrescriberResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*v.IncludedPrescriberResources))
	} else if len(*v.IncludedPrescriberResources) == 1 {
		practitioner = &(*v.IncludedPrescriberResources)[0]
	}
	return
}

func (v *VisionPrescriptionPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if v.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*v.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*v.IncludedPatientResources))
	} else if len(*v.IncludedPatientResources) == 1 {
		patient = &(*v.IncludedPatientResources)[0]
	}
	return
}

func (v *VisionPrescriptionPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if v.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*v.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*v.IncludedEncounterResources))
	} else if len(*v.IncludedEncounterResources) == 1 {
		encounter = &(*v.IncludedEncounterResources)[0]
	}
	return
}

func (v *VisionPrescriptionPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.IncludedPrescriberResources != nil {
		for _, r := range *v.IncludedPrescriberResources {
			resourceMap[r.Id] = &r
		}
	}
	if v.IncludedPatientResources != nil {
		for _, r := range *v.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if v.IncludedEncounterResources != nil {
		for _, r := range *v.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
