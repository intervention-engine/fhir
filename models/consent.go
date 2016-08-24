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

type Consent struct {
	DomainResource   `bson:",inline"`
	Identifier       *Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status           string                   `bson:"status,omitempty" json:"status,omitempty"`
	Category         []CodeableConcept        `bson:"category,omitempty" json:"category,omitempty"`
	DateTime         *FHIRDateTime            `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Period           *Period                  `bson:"period,omitempty" json:"period,omitempty"`
	Patient          *Reference               `bson:"patient,omitempty" json:"patient,omitempty"`
	Consentor        []Reference              `bson:"consentor,omitempty" json:"consentor,omitempty"`
	Organization     *Reference               `bson:"organization,omitempty" json:"organization,omitempty"`
	SourceAttachment *Attachment              `bson:"sourceAttachment,omitempty" json:"sourceAttachment,omitempty"`
	SourceIdentifier *Identifier              `bson:"sourceIdentifier,omitempty" json:"sourceIdentifier,omitempty"`
	SourceReference  *Reference               `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	Policy           string                   `bson:"policy,omitempty" json:"policy,omitempty"`
	Recipient        []Reference              `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Purpose          []Coding                 `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Except           []ConsentExceptComponent `bson:"except,omitempty" json:"except,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Consent) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Consent"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Consent), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Consent) GetBSON() (interface{}, error) {
	x.ResourceType = "Consent"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "consent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type consent Consent

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Consent) UnmarshalJSON(data []byte) (err error) {
	x2 := consent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Consent(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Consent) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Consent"
	} else if x.ResourceType != "Consent" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Consent, instead received %s", x.ResourceType))
	}
	return nil
}

type ConsentExceptComponent struct {
	BackboneElement `bson:",inline"`
	Type            string                        `bson:"type,omitempty" json:"type,omitempty"`
	Period          *Period                       `bson:"period,omitempty" json:"period,omitempty"`
	Actor           []ConsentExceptActorComponent `bson:"actor,omitempty" json:"actor,omitempty"`
	Action          []CodeableConcept             `bson:"action,omitempty" json:"action,omitempty"`
	SecurityLabel   []Coding                      `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Purpose         []Coding                      `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Class           []Coding                      `bson:"class,omitempty" json:"class,omitempty"`
	Code            []Coding                      `bson:"code,omitempty" json:"code,omitempty"`
	Data            []ConsentExceptDataComponent  `bson:"data,omitempty" json:"data,omitempty"`
}

type ConsentExceptActorComponent struct {
	BackboneElement `bson:",inline"`
	Role            *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Reference       *Reference       `bson:"reference,omitempty" json:"reference,omitempty"`
}

type ConsentExceptDataComponent struct {
	BackboneElement `bson:",inline"`
	Meaning         string     `bson:"meaning,omitempty" json:"meaning,omitempty"`
	Reference       *Reference `bson:"reference,omitempty" json:"reference,omitempty"`
}

type ConsentPlus struct {
	Consent                     `bson:",inline"`
	ConsentPlusRelatedResources `bson:",inline"`
}

type ConsentPlusRelatedResources struct {
	IncludedConsentResourcesReferencedBySource                     *[]Consent               `bson:"_includedConsentResourcesReferencedBySource,omitempty"`
	IncludedContractResourcesReferencedBySource                    *[]Contract              `bson:"_includedContractResourcesReferencedBySource,omitempty"`
	IncludedQuestionnaireResponseResourcesReferencedBySource       *[]QuestionnaireResponse `bson:"_includedQuestionnaireResponseResourcesReferencedBySource,omitempty"`
	IncludedDocumentReferenceResourcesReferencedBySource           *[]DocumentReference     `bson:"_includedDocumentReferenceResourcesReferencedBySource,omitempty"`
	IncludedPractitionerResourcesReferencedByActor                 *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByActor,omitempty"`
	IncludedGroupResourcesReferencedByActor                        *[]Group                 `bson:"_includedGroupResourcesReferencedByActor,omitempty"`
	IncludedOrganizationResourcesReferencedByActor                 *[]Organization          `bson:"_includedOrganizationResourcesReferencedByActor,omitempty"`
	IncludedCareTeamResourcesReferencedByActor                     *[]CareTeam              `bson:"_includedCareTeamResourcesReferencedByActor,omitempty"`
	IncludedDeviceResourcesReferencedByActor                       *[]Device                `bson:"_includedDeviceResourcesReferencedByActor,omitempty"`
	IncludedPatientResourcesReferencedByActor                      *[]Patient               `bson:"_includedPatientResourcesReferencedByActor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByActor                *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByActor,omitempty"`
	IncludedPatientResourcesReferencedByPatient                    *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization          *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedPractitionerResourcesReferencedByRecipient             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByRecipient,omitempty"`
	IncludedGroupResourcesReferencedByRecipient                    *[]Group                 `bson:"_includedGroupResourcesReferencedByRecipient,omitempty"`
	IncludedOrganizationResourcesReferencedByRecipient             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByRecipient,omitempty"`
	IncludedCareTeamResourcesReferencedByRecipient                 *[]CareTeam              `bson:"_includedCareTeamResourcesReferencedByRecipient,omitempty"`
	IncludedDeviceResourcesReferencedByRecipient                   *[]Device                `bson:"_includedDeviceResourcesReferencedByRecipient,omitempty"`
	IncludedPatientResourcesReferencedByRecipient                  *[]Patient               `bson:"_includedPatientResourcesReferencedByRecipient,omitempty"`
	IncludedRelatedPersonResourcesReferencedByRecipient            *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByRecipient,omitempty"`
	IncludedPractitionerResourcesReferencedByConsentor             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByConsentor,omitempty"`
	IncludedOrganizationResourcesReferencedByConsentor             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByConsentor,omitempty"`
	IncludedPatientResourcesReferencedByConsentor                  *[]Patient               `bson:"_includedPatientResourcesReferencedByConsentor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByConsentor            *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByConsentor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                     *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedConsentResourcesReferencingSource                   *[]Consent               `bson:"_revIncludedConsentResourcesReferencingSource,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref     *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                   *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequestreference   *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequestreference,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponsereference  *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponsereference,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource     *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon            *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                       *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedListResourcesReferencingItem                        *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces       *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon        *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces        *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon         *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                    *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated         *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject    *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestreference *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequestreference,omitempty"`
}

func (c *ConsentPlusRelatedResources) GetIncludedConsentResourceReferencedBySource() (consent *Consent, err error) {
	if c.IncludedConsentResourcesReferencedBySource == nil {
		err = errors.New("Included consents not requested")
	} else if len(*c.IncludedConsentResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 consent, but found %d", len(*c.IncludedConsentResourcesReferencedBySource))
	} else if len(*c.IncludedConsentResourcesReferencedBySource) == 1 {
		consent = &(*c.IncludedConsentResourcesReferencedBySource)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedContractResourceReferencedBySource() (contract *Contract, err error) {
	if c.IncludedContractResourcesReferencedBySource == nil {
		err = errors.New("Included contracts not requested")
	} else if len(*c.IncludedContractResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 contract, but found %d", len(*c.IncludedContractResourcesReferencedBySource))
	} else if len(*c.IncludedContractResourcesReferencedBySource) == 1 {
		contract = &(*c.IncludedContractResourcesReferencedBySource)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedQuestionnaireResponseResourceReferencedBySource() (questionnaireResponse *QuestionnaireResponse, err error) {
	if c.IncludedQuestionnaireResponseResourcesReferencedBySource == nil {
		err = errors.New("Included questionnaireresponses not requested")
	} else if len(*c.IncludedQuestionnaireResponseResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 questionnaireResponse, but found %d", len(*c.IncludedQuestionnaireResponseResourcesReferencedBySource))
	} else if len(*c.IncludedQuestionnaireResponseResourcesReferencedBySource) == 1 {
		questionnaireResponse = &(*c.IncludedQuestionnaireResponseResourcesReferencedBySource)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedDocumentReferenceResourceReferencedBySource() (documentReference *DocumentReference, err error) {
	if c.IncludedDocumentReferenceResourcesReferencedBySource == nil {
		err = errors.New("Included documentreferences not requested")
	} else if len(*c.IncludedDocumentReferenceResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 documentReference, but found %d", len(*c.IncludedDocumentReferenceResourcesReferencedBySource))
	} else if len(*c.IncludedDocumentReferenceResourcesReferencedBySource) == 1 {
		documentReference = &(*c.IncludedDocumentReferenceResourcesReferencedBySource)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedPractitionerResourceReferencedByActor() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByActor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByActor))
	} else if len(*c.IncludedPractitionerResourcesReferencedByActor) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByActor)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedGroupResourceReferencedByActor() (group *Group, err error) {
	if c.IncludedGroupResourcesReferencedByActor == nil {
		err = errors.New("Included groups not requested")
	} else if len(*c.IncludedGroupResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*c.IncludedGroupResourcesReferencedByActor))
	} else if len(*c.IncludedGroupResourcesReferencedByActor) == 1 {
		group = &(*c.IncludedGroupResourcesReferencedByActor)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedOrganizationResourceReferencedByActor() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByActor == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByActor))
	} else if len(*c.IncludedOrganizationResourcesReferencedByActor) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByActor)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedCareTeamResourceReferencedByActor() (careTeam *CareTeam, err error) {
	if c.IncludedCareTeamResourcesReferencedByActor == nil {
		err = errors.New("Included careteams not requested")
	} else if len(*c.IncludedCareTeamResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 careTeam, but found %d", len(*c.IncludedCareTeamResourcesReferencedByActor))
	} else if len(*c.IncludedCareTeamResourcesReferencedByActor) == 1 {
		careTeam = &(*c.IncludedCareTeamResourcesReferencedByActor)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedDeviceResourceReferencedByActor() (device *Device, err error) {
	if c.IncludedDeviceResourcesReferencedByActor == nil {
		err = errors.New("Included devices not requested")
	} else if len(*c.IncludedDeviceResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*c.IncludedDeviceResourcesReferencedByActor))
	} else if len(*c.IncludedDeviceResourcesReferencedByActor) == 1 {
		device = &(*c.IncludedDeviceResourcesReferencedByActor)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedPatientResourceReferencedByActor() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByActor == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByActor))
	} else if len(*c.IncludedPatientResourcesReferencedByActor) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByActor)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByActor() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByActor == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedByActor))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByActor) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedByActor)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPatient))
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*c.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByRecipient() (practitioners []Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByRecipient == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *c.IncludedPractitionerResourcesReferencedByRecipient
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedGroupResourcesReferencedByRecipient() (groups []Group, err error) {
	if c.IncludedGroupResourcesReferencedByRecipient == nil {
		err = errors.New("Included groups not requested")
	} else {
		groups = *c.IncludedGroupResourcesReferencedByRecipient
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByRecipient() (organizations []Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByRecipient == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *c.IncludedOrganizationResourcesReferencedByRecipient
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedCareTeamResourcesReferencedByRecipient() (careTeams []CareTeam, err error) {
	if c.IncludedCareTeamResourcesReferencedByRecipient == nil {
		err = errors.New("Included careTeams not requested")
	} else {
		careTeams = *c.IncludedCareTeamResourcesReferencedByRecipient
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedDeviceResourcesReferencedByRecipient() (devices []Device, err error) {
	if c.IncludedDeviceResourcesReferencedByRecipient == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *c.IncludedDeviceResourcesReferencedByRecipient
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedPatientResourcesReferencedByRecipient() (patients []Patient, err error) {
	if c.IncludedPatientResourcesReferencedByRecipient == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResourcesReferencedByRecipient
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByRecipient() (relatedPeople []RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByRecipient == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *c.IncludedRelatedPersonResourcesReferencedByRecipient
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByConsentor() (practitioners []Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByConsentor == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *c.IncludedPractitionerResourcesReferencedByConsentor
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByConsentor() (organizations []Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByConsentor == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *c.IncludedOrganizationResourcesReferencedByConsentor
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedPatientResourcesReferencedByConsentor() (patients []Patient, err error) {
	if c.IncludedPatientResourcesReferencedByConsentor == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResourcesReferencedByConsentor
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByConsentor() (relatedPeople []RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByConsentor == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *c.IncludedRelatedPersonResourcesReferencedByConsentor
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedConsentResourcesReferencingSource() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingSource == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingSource
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if c.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *c.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *ConsentPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequestreference
	}
	return
}

func (c *ConsentPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedConsentResourcesReferencedBySource != nil {
		for idx := range *c.IncludedConsentResourcesReferencedBySource {
			rsc := (*c.IncludedConsentResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedContractResourcesReferencedBySource != nil {
		for idx := range *c.IncludedContractResourcesReferencedBySource {
			rsc := (*c.IncludedContractResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedQuestionnaireResponseResourcesReferencedBySource != nil {
		for idx := range *c.IncludedQuestionnaireResponseResourcesReferencedBySource {
			rsc := (*c.IncludedQuestionnaireResponseResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDocumentReferenceResourcesReferencedBySource != nil {
		for idx := range *c.IncludedDocumentReferenceResourcesReferencedBySource {
			rsc := (*c.IncludedDocumentReferenceResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByActor != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByActor {
			rsc := (*c.IncludedPractitionerResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedByActor != nil {
		for idx := range *c.IncludedGroupResourcesReferencedByActor {
			rsc := (*c.IncludedGroupResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByActor != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByActor {
			rsc := (*c.IncludedOrganizationResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCareTeamResourcesReferencedByActor != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByActor {
			rsc := (*c.IncludedCareTeamResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByActor != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByActor {
			rsc := (*c.IncludedDeviceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByActor != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByActor {
			rsc := (*c.IncludedPatientResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByActor != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByActor {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*c.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByRecipient {
			rsc := (*c.IncludedPractitionerResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedGroupResourcesReferencedByRecipient {
			rsc := (*c.IncludedGroupResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByRecipient {
			rsc := (*c.IncludedOrganizationResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCareTeamResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByRecipient {
			rsc := (*c.IncludedCareTeamResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByRecipient {
			rsc := (*c.IncludedDeviceResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByRecipient {
			rsc := (*c.IncludedPatientResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByRecipient {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByConsentor != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByConsentor {
			rsc := (*c.IncludedPractitionerResourcesReferencedByConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByConsentor != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByConsentor {
			rsc := (*c.IncludedOrganizationResourcesReferencedByConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByConsentor != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByConsentor {
			rsc := (*c.IncludedPatientResourcesReferencedByConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByConsentor != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByConsentor {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ConsentPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingData {
			rsc := (*c.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingSource != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingSource {
			rsc := (*c.RevIncludedConsentResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ConsentPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedConsentResourcesReferencedBySource != nil {
		for idx := range *c.IncludedConsentResourcesReferencedBySource {
			rsc := (*c.IncludedConsentResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedContractResourcesReferencedBySource != nil {
		for idx := range *c.IncludedContractResourcesReferencedBySource {
			rsc := (*c.IncludedContractResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedQuestionnaireResponseResourcesReferencedBySource != nil {
		for idx := range *c.IncludedQuestionnaireResponseResourcesReferencedBySource {
			rsc := (*c.IncludedQuestionnaireResponseResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDocumentReferenceResourcesReferencedBySource != nil {
		for idx := range *c.IncludedDocumentReferenceResourcesReferencedBySource {
			rsc := (*c.IncludedDocumentReferenceResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByActor != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByActor {
			rsc := (*c.IncludedPractitionerResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedByActor != nil {
		for idx := range *c.IncludedGroupResourcesReferencedByActor {
			rsc := (*c.IncludedGroupResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByActor != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByActor {
			rsc := (*c.IncludedOrganizationResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCareTeamResourcesReferencedByActor != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByActor {
			rsc := (*c.IncludedCareTeamResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByActor != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByActor {
			rsc := (*c.IncludedDeviceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByActor != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByActor {
			rsc := (*c.IncludedPatientResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByActor != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByActor {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*c.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByRecipient {
			rsc := (*c.IncludedPractitionerResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedGroupResourcesReferencedByRecipient {
			rsc := (*c.IncludedGroupResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByRecipient {
			rsc := (*c.IncludedOrganizationResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCareTeamResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByRecipient {
			rsc := (*c.IncludedCareTeamResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByRecipient {
			rsc := (*c.IncludedDeviceResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByRecipient {
			rsc := (*c.IncludedPatientResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByRecipient != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByRecipient {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByConsentor != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByConsentor {
			rsc := (*c.IncludedPractitionerResourcesReferencedByConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByConsentor != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByConsentor {
			rsc := (*c.IncludedOrganizationResourcesReferencedByConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByConsentor != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByConsentor {
			rsc := (*c.IncludedPatientResourcesReferencedByConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByConsentor != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByConsentor {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingData {
			rsc := (*c.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingSource != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingSource {
			rsc := (*c.RevIncludedConsentResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
