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

type ReferralRequest struct {
	DomainResource        `bson:",inline"`
	Status                string            `bson:"status,omitempty" json:"status,omitempty"`
	Identifier            []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Date                  *FHIRDateTime     `bson:"date,omitempty" json:"date,omitempty"`
	Type                  *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Specialty             *CodeableConcept  `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Priority              *CodeableConcept  `bson:"priority,omitempty" json:"priority,omitempty"`
	Patient               *Reference        `bson:"patient,omitempty" json:"patient,omitempty"`
	Requester             *Reference        `bson:"requester,omitempty" json:"requester,omitempty"`
	Recipient             []Reference       `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Encounter             *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateSent              *FHIRDateTime     `bson:"dateSent,omitempty" json:"dateSent,omitempty"`
	Reason                *CodeableConcept  `bson:"reason,omitempty" json:"reason,omitempty"`
	Description           string            `bson:"description,omitempty" json:"description,omitempty"`
	ServiceRequested      []CodeableConcept `bson:"serviceRequested,omitempty" json:"serviceRequested,omitempty"`
	SupportingInformation []Reference       `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	FulfillmentTime       *Period           `bson:"fulfillmentTime,omitempty" json:"fulfillmentTime,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ReferralRequest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ReferralRequest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ReferralRequest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ReferralRequest) GetBSON() (interface{}, error) {
	x.ResourceType = "ReferralRequest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "referralRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type referralRequest ReferralRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ReferralRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := referralRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ReferralRequest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ReferralRequest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ReferralRequest"
	} else if x.ResourceType != "ReferralRequest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ReferralRequest, instead received %s", x.ResourceType))
	}
	return nil
}

type ReferralRequestPlus struct {
	ReferralRequest             `bson:",inline"`
	ReferralRequestPlusIncludes `bson:",inline"`
}

type ReferralRequestPlusIncludes struct {
	IncludedRequesterPractitionerResources *[]Practitioner `bson:"_includedRequesterPractitionerResources,omitempty"`
	IncludedRequesterOrganizationResources *[]Organization `bson:"_includedRequesterOrganizationResources,omitempty"`
	IncludedRequesterPatientResources      *[]Patient      `bson:"_includedRequesterPatientResources,omitempty"`
	IncludedPatientResources               *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedRecipientPractitionerResources *[]Practitioner `bson:"_includedRecipientPractitionerResources,omitempty"`
	IncludedRecipientOrganizationResources *[]Organization `bson:"_includedRecipientOrganizationResources,omitempty"`
}

func (r *ReferralRequestPlusIncludes) GetIncludedRequesterPractitionerResource() (practitioner *Practitioner, err error) {
	if r.IncludedRequesterPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*r.IncludedRequesterPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*r.IncludedRequesterPractitionerResources))
	} else if len(*r.IncludedRequesterPractitionerResources) == 1 {
		practitioner = &(*r.IncludedRequesterPractitionerResources)[0]
	}
	return
}

func (r *ReferralRequestPlusIncludes) GetIncludedRequesterOrganizationResource() (organization *Organization, err error) {
	if r.IncludedRequesterOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*r.IncludedRequesterOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*r.IncludedRequesterOrganizationResources))
	} else if len(*r.IncludedRequesterOrganizationResources) == 1 {
		organization = &(*r.IncludedRequesterOrganizationResources)[0]
	}
	return
}

func (r *ReferralRequestPlusIncludes) GetIncludedRequesterPatientResource() (patient *Patient, err error) {
	if r.IncludedRequesterPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedRequesterPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedRequesterPatientResources))
	} else if len(*r.IncludedRequesterPatientResources) == 1 {
		patient = &(*r.IncludedRequesterPatientResources)[0]
	}
	return
}

func (r *ReferralRequestPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if r.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResources))
	} else if len(*r.IncludedPatientResources) == 1 {
		patient = &(*r.IncludedPatientResources)[0]
	}
	return
}

func (r *ReferralRequestPlusIncludes) GetIncludedRecipientPractitionerResources() (practitioners []Practitioner, err error) {
	if r.IncludedRecipientPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *r.IncludedRecipientPractitionerResources
	}
	return
}

func (r *ReferralRequestPlusIncludes) GetIncludedRecipientOrganizationResources() (organizations []Organization, err error) {
	if r.IncludedRecipientOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *r.IncludedRecipientOrganizationResources
	}
	return
}

func (r *ReferralRequestPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedRequesterPractitionerResources != nil {
		for _, r := range *r.IncludedRequesterPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedRequesterOrganizationResources != nil {
		for _, r := range *r.IncludedRequesterOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedRequesterPatientResources != nil {
		for _, r := range *r.IncludedRequesterPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedPatientResources != nil {
		for _, r := range *r.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedRecipientPractitionerResources != nil {
		for _, r := range *r.IncludedRecipientPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedRecipientOrganizationResources != nil {
		for _, r := range *r.IncludedRecipientOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
