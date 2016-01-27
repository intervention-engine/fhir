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

type EpisodeOfCare struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                                `bson:"status,omitempty" json:"status,omitempty"`
	StatusHistory        []EpisodeOfCareStatusHistoryComponent `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Type                 []CodeableConcept                     `bson:"type,omitempty" json:"type,omitempty"`
	Condition            []Reference                           `bson:"condition,omitempty" json:"condition,omitempty"`
	Patient              *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	ManagingOrganization *Reference                            `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Period               *Period                               `bson:"period,omitempty" json:"period,omitempty"`
	ReferralRequest      []Reference                           `bson:"referralRequest,omitempty" json:"referralRequest,omitempty"`
	CareManager          *Reference                            `bson:"careManager,omitempty" json:"careManager,omitempty"`
	CareTeam             []EpisodeOfCareCareTeamComponent      `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *EpisodeOfCare) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		EpisodeOfCare
	}{
		ResourceType:  "EpisodeOfCare",
		EpisodeOfCare: *resource,
	}
	return json.Marshal(x)
}

// The "episodeOfCare" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type episodeOfCare EpisodeOfCare

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *EpisodeOfCare) UnmarshalJSON(data []byte) (err error) {
	x2 := episodeOfCare{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = EpisodeOfCare(x2)
	}
	return
}

type EpisodeOfCareStatusHistoryComponent struct {
	Status string  `bson:"status,omitempty" json:"status,omitempty"`
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type EpisodeOfCareCareTeamComponent struct {
	Role   []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Period *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Member *Reference        `bson:"member,omitempty" json:"member,omitempty"`
}

type EpisodeOfCarePlus struct {
	EpisodeOfCare             `bson:",inline"`
	EpisodeOfCarePlusIncludes `bson:",inline"`
}

type EpisodeOfCarePlusIncludes struct {
	IncludedConditionResources              *[]Condition       `bson:"_includedConditionResources,omitempty"`
	IncludedIncomingreferralResources       *[]ReferralRequest `bson:"_includedIncomingreferralResources,omitempty"`
	IncludedPatientResources                *[]Patient         `bson:"_includedPatientResources,omitempty"`
	IncludedOrganizationResources           *[]Organization    `bson:"_includedOrganizationResources,omitempty"`
	IncludedTeammemberPractitionerResources *[]Practitioner    `bson:"_includedTeammemberPractitionerResources,omitempty"`
	IncludedTeammemberOrganizationResources *[]Organization    `bson:"_includedTeammemberOrganizationResources,omitempty"`
	IncludedCaremanagerResources            *[]Practitioner    `bson:"_includedCaremanagerResources,omitempty"`
}

func (e *EpisodeOfCarePlusIncludes) GetIncludedConditionResources() (conditions []Condition, err error) {
	if e.IncludedConditionResources == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *e.IncludedConditionResources
	}
	return
}

func (e *EpisodeOfCarePlusIncludes) GetIncludedIncomingreferralResources() (referralRequests []ReferralRequest, err error) {
	if e.IncludedIncomingreferralResources == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *e.IncludedIncomingreferralResources
	}
	return
}

func (e *EpisodeOfCarePlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if e.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResources))
	} else if len(*e.IncludedPatientResources) == 1 {
		patient = &(*e.IncludedPatientResources)[0]
	}
	return
}

func (e *EpisodeOfCarePlusIncludes) GetIncludedOrganizationResource() (organization *Organization, err error) {
	if e.IncludedOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResources))
	} else if len(*e.IncludedOrganizationResources) == 1 {
		organization = &(*e.IncludedOrganizationResources)[0]
	}
	return
}

func (e *EpisodeOfCarePlusIncludes) GetIncludedTeammemberPractitionerResource() (practitioner *Practitioner, err error) {
	if e.IncludedTeammemberPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedTeammemberPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedTeammemberPractitionerResources))
	} else if len(*e.IncludedTeammemberPractitionerResources) == 1 {
		practitioner = &(*e.IncludedTeammemberPractitionerResources)[0]
	}
	return
}

func (e *EpisodeOfCarePlusIncludes) GetIncludedTeammemberOrganizationResource() (organization *Organization, err error) {
	if e.IncludedTeammemberOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedTeammemberOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedTeammemberOrganizationResources))
	} else if len(*e.IncludedTeammemberOrganizationResources) == 1 {
		organization = &(*e.IncludedTeammemberOrganizationResources)[0]
	}
	return
}

func (e *EpisodeOfCarePlusIncludes) GetIncludedCaremanagerResource() (practitioner *Practitioner, err error) {
	if e.IncludedCaremanagerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedCaremanagerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedCaremanagerResources))
	} else if len(*e.IncludedCaremanagerResources) == 1 {
		practitioner = &(*e.IncludedCaremanagerResources)[0]
	}
	return
}

func (e *EpisodeOfCarePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedConditionResources != nil {
		for _, r := range *e.IncludedConditionResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedIncomingreferralResources != nil {
		for _, r := range *e.IncludedIncomingreferralResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPatientResources != nil {
		for _, r := range *e.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedOrganizationResources != nil {
		for _, r := range *e.IncludedOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedTeammemberPractitionerResources != nil {
		for _, r := range *e.IncludedTeammemberPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedTeammemberOrganizationResources != nil {
		for _, r := range *e.IncludedTeammemberOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedCaremanagerResources != nil {
		for _, r := range *e.IncludedCaremanagerResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
