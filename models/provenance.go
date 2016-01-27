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

type Provenance struct {
	DomainResource `bson:",inline"`
	Target         []Reference                 `bson:"target,omitempty" json:"target,omitempty"`
	Period         *Period                     `bson:"period,omitempty" json:"period,omitempty"`
	Recorded       *FHIRDateTime               `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Reason         []CodeableConcept           `bson:"reason,omitempty" json:"reason,omitempty"`
	Activity       *CodeableConcept            `bson:"activity,omitempty" json:"activity,omitempty"`
	Location       *Reference                  `bson:"location,omitempty" json:"location,omitempty"`
	Policy         []string                    `bson:"policy,omitempty" json:"policy,omitempty"`
	Agent          []ProvenanceAgentComponent  `bson:"agent,omitempty" json:"agent,omitempty"`
	Entity         []ProvenanceEntityComponent `bson:"entity,omitempty" json:"entity,omitempty"`
	Signature      []Signature                 `bson:"signature,omitempty" json:"signature,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Provenance) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Provenance
	}{
		ResourceType: "Provenance",
		Provenance:   *resource,
	}
	return json.Marshal(x)
}

// The "provenance" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type provenance Provenance

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Provenance) UnmarshalJSON(data []byte) (err error) {
	x2 := provenance{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Provenance(x2)
	}
	return
}

type ProvenanceAgentComponent struct {
	Role         *Coding                                `bson:"role,omitempty" json:"role,omitempty"`
	Actor        *Reference                             `bson:"actor,omitempty" json:"actor,omitempty"`
	UserId       *Identifier                            `bson:"userId,omitempty" json:"userId,omitempty"`
	RelatedAgent []ProvenanceAgentRelatedAgentComponent `bson:"relatedAgent,omitempty" json:"relatedAgent,omitempty"`
}

type ProvenanceAgentRelatedAgentComponent struct {
	Type   *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Target string           `bson:"target,omitempty" json:"target,omitempty"`
}

type ProvenanceEntityComponent struct {
	Role      string                    `bson:"role,omitempty" json:"role,omitempty"`
	Type      *Coding                   `bson:"type,omitempty" json:"type,omitempty"`
	Reference string                    `bson:"reference,omitempty" json:"reference,omitempty"`
	Display   string                    `bson:"display,omitempty" json:"display,omitempty"`
	Agent     *ProvenanceAgentComponent `bson:"agent,omitempty" json:"agent,omitempty"`
}

type ProvenancePlus struct {
	Provenance             `bson:",inline"`
	ProvenancePlusIncludes `bson:",inline"`
}

type ProvenancePlusIncludes struct {
	IncludedAgentPractitionerResources  *[]Practitioner  `bson:"_includedAgentPractitionerResources,omitempty"`
	IncludedAgentOrganizationResources  *[]Organization  `bson:"_includedAgentOrganizationResources,omitempty"`
	IncludedAgentDeviceResources        *[]Device        `bson:"_includedAgentDeviceResources,omitempty"`
	IncludedAgentPatientResources       *[]Patient       `bson:"_includedAgentPatientResources,omitempty"`
	IncludedAgentRelatedPersonResources *[]RelatedPerson `bson:"_includedAgentRelatedPersonResources,omitempty"`
	IncludedPatientResources            *[]Patient       `bson:"_includedPatientResources,omitempty"`
	IncludedLocationResources           *[]Location      `bson:"_includedLocationResources,omitempty"`
}

func (p *ProvenancePlusIncludes) GetIncludedAgentPractitionerResource() (practitioner *Practitioner, err error) {
	if p.IncludedAgentPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedAgentPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedAgentPractitionerResources))
	} else if len(*p.IncludedAgentPractitionerResources) == 1 {
		practitioner = &(*p.IncludedAgentPractitionerResources)[0]
	}
	return
}

func (p *ProvenancePlusIncludes) GetIncludedAgentOrganizationResource() (organization *Organization, err error) {
	if p.IncludedAgentOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedAgentOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedAgentOrganizationResources))
	} else if len(*p.IncludedAgentOrganizationResources) == 1 {
		organization = &(*p.IncludedAgentOrganizationResources)[0]
	}
	return
}

func (p *ProvenancePlusIncludes) GetIncludedAgentDeviceResource() (device *Device, err error) {
	if p.IncludedAgentDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*p.IncludedAgentDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*p.IncludedAgentDeviceResources))
	} else if len(*p.IncludedAgentDeviceResources) == 1 {
		device = &(*p.IncludedAgentDeviceResources)[0]
	}
	return
}

func (p *ProvenancePlusIncludes) GetIncludedAgentPatientResource() (patient *Patient, err error) {
	if p.IncludedAgentPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedAgentPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedAgentPatientResources))
	} else if len(*p.IncludedAgentPatientResources) == 1 {
		patient = &(*p.IncludedAgentPatientResources)[0]
	}
	return
}

func (p *ProvenancePlusIncludes) GetIncludedAgentRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if p.IncludedAgentRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*p.IncludedAgentRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*p.IncludedAgentRelatedPersonResources))
	} else if len(*p.IncludedAgentRelatedPersonResources) == 1 {
		relatedPerson = &(*p.IncludedAgentRelatedPersonResources)[0]
	}
	return
}

func (p *ProvenancePlusIncludes) GetIncludedPatientResources() (patients []Patient, err error) {
	if p.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *p.IncludedPatientResources
	}
	return
}

func (p *ProvenancePlusIncludes) GetIncludedLocationResource() (location *Location, err error) {
	if p.IncludedLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*p.IncludedLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*p.IncludedLocationResources))
	} else if len(*p.IncludedLocationResources) == 1 {
		location = &(*p.IncludedLocationResources)[0]
	}
	return
}

func (p *ProvenancePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedAgentPractitionerResources != nil {
		for _, r := range *p.IncludedAgentPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedAgentOrganizationResources != nil {
		for _, r := range *p.IncludedAgentOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedAgentDeviceResources != nil {
		for _, r := range *p.IncludedAgentDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedAgentPatientResources != nil {
		for _, r := range *p.IncludedAgentPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedAgentRelatedPersonResources != nil {
		for _, r := range *p.IncludedAgentRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedPatientResources != nil {
		for _, r := range *p.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedLocationResources != nil {
		for _, r := range *p.IncludedLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
