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

type DocumentReference struct {
	DomainResource   `bson:",inline"`
	MasterIdentifier *Identifier                           `bson:"masterIdentifier,omitempty" json:"masterIdentifier,omitempty"`
	Identifier       []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject          *Reference                            `bson:"subject,omitempty" json:"subject,omitempty"`
	Type             *CodeableConcept                      `bson:"type,omitempty" json:"type,omitempty"`
	Class            *CodeableConcept                      `bson:"class,omitempty" json:"class,omitempty"`
	Author           []Reference                           `bson:"author,omitempty" json:"author,omitempty"`
	Custodian        *Reference                            `bson:"custodian,omitempty" json:"custodian,omitempty"`
	Authenticator    *Reference                            `bson:"authenticator,omitempty" json:"authenticator,omitempty"`
	Created          *FHIRDateTime                         `bson:"created,omitempty" json:"created,omitempty"`
	Indexed          *FHIRDateTime                         `bson:"indexed,omitempty" json:"indexed,omitempty"`
	Status           string                                `bson:"status,omitempty" json:"status,omitempty"`
	DocStatus        *CodeableConcept                      `bson:"docStatus,omitempty" json:"docStatus,omitempty"`
	RelatesTo        []DocumentReferenceRelatesToComponent `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Description      string                                `bson:"description,omitempty" json:"description,omitempty"`
	SecurityLabel    []CodeableConcept                     `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Content          []DocumentReferenceContentComponent   `bson:"content,omitempty" json:"content,omitempty"`
	Context          *DocumentReferenceContextComponent    `bson:"context,omitempty" json:"context,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DocumentReference) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DocumentReference"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DocumentReference), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DocumentReference) GetBSON() (interface{}, error) {
	x.ResourceType = "DocumentReference"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "documentReference" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type documentReference DocumentReference

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DocumentReference) UnmarshalJSON(data []byte) (err error) {
	x2 := documentReference{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DocumentReference(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DocumentReference) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DocumentReference"
	} else if x.ResourceType != "DocumentReference" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DocumentReference, instead received %s", x.ResourceType))
	}
	return nil
}

type DocumentReferenceRelatesToComponent struct {
	BackboneElement `bson:",inline"`
	Code            string     `bson:"code,omitempty" json:"code,omitempty"`
	Target          *Reference `bson:"target,omitempty" json:"target,omitempty"`
}

type DocumentReferenceContentComponent struct {
	BackboneElement `bson:",inline"`
	Attachment      *Attachment `bson:"attachment,omitempty" json:"attachment,omitempty"`
	Format          []Coding    `bson:"format,omitempty" json:"format,omitempty"`
}

type DocumentReferenceContextComponent struct {
	BackboneElement   `bson:",inline"`
	Encounter         *Reference                                 `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Event             []CodeableConcept                          `bson:"event,omitempty" json:"event,omitempty"`
	Period            *Period                                    `bson:"period,omitempty" json:"period,omitempty"`
	FacilityType      *CodeableConcept                           `bson:"facilityType,omitempty" json:"facilityType,omitempty"`
	PracticeSetting   *CodeableConcept                           `bson:"practiceSetting,omitempty" json:"practiceSetting,omitempty"`
	SourcePatientInfo *Reference                                 `bson:"sourcePatientInfo,omitempty" json:"sourcePatientInfo,omitempty"`
	Related           []DocumentReferenceContextRelatedComponent `bson:"related,omitempty" json:"related,omitempty"`
}

type DocumentReferenceContextRelatedComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ref             *Reference  `bson:"ref,omitempty" json:"ref,omitempty"`
}

type DocumentReferencePlus struct {
	DocumentReference                     `bson:",inline"`
	DocumentReferencePlusRelatedResources `bson:",inline"`
}

type DocumentReferencePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedBySubject            *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySubject,omitempty"`
	IncludedGroupResourcesReferencedBySubject                   *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                  *[]Device                `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByAuthenticator      *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthenticator,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthenticator      *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAuthenticator,omitempty"`
	IncludedOrganizationResourcesReferencedByCustodian          *[]Organization          `bson:"_includedOrganizationResourcesReferencedByCustodian,omitempty"`
	IncludedPractitionerResourcesReferencedByAuthor             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthor             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAuthor,omitempty"`
	IncludedDeviceResourcesReferencedByAuthor                   *[]Device                `bson:"_includedDeviceResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByAuthor                  *[]Patient               `bson:"_includedPatientResourcesReferencedByAuthor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAuthor            *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByAuthor,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter             *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedDocumentReferenceResourcesReferencedByRelatesto     *[]DocumentReference     `bson:"_includedDocumentReferenceResourcesReferencedByRelatesto,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                  *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedConsentResourcesReferencingSource                *[]Consent               `bson:"_revIncludedConsentResourcesReferencingSource,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatesto   *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatesto,omitempty"`
	RevIncludedContractResourcesReferencingTtopic               *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject              *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse        *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource  *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon         *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                    *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces    *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition  *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces     *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition   *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity             *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedPractitionerResourceReferencedBySubject() (practitioner *Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedBySubject == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPractitionerResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPractitionerResourcesReferencedBySubject))
	} else if len(*d.IncludedPractitionerResourcesReferencedBySubject) == 1 {
		practitioner = &(*d.IncludedPractitionerResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if d.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedGroupResourcesReferencedBySubject))
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*d.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*d.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedBySubject))
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedByPatient))
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthenticator() (practitioner *Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByAuthenticator == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPractitionerResourcesReferencedByAuthenticator) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPractitionerResourcesReferencedByAuthenticator))
	} else if len(*d.IncludedPractitionerResourcesReferencedByAuthenticator) == 1 {
		practitioner = &(*d.IncludedPractitionerResourcesReferencedByAuthenticator)[0]
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedOrganizationResourceReferencedByAuthenticator() (organization *Organization, err error) {
	if d.IncludedOrganizationResourcesReferencedByAuthenticator == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*d.IncludedOrganizationResourcesReferencedByAuthenticator) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*d.IncludedOrganizationResourcesReferencedByAuthenticator))
	} else if len(*d.IncludedOrganizationResourcesReferencedByAuthenticator) == 1 {
		organization = &(*d.IncludedOrganizationResourcesReferencedByAuthenticator)[0]
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedOrganizationResourceReferencedByCustodian() (organization *Organization, err error) {
	if d.IncludedOrganizationResourcesReferencedByCustodian == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*d.IncludedOrganizationResourcesReferencedByCustodian) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*d.IncludedOrganizationResourcesReferencedByCustodian))
	} else if len(*d.IncludedOrganizationResourcesReferencedByCustodian) == 1 {
		organization = &(*d.IncludedOrganizationResourcesReferencedByCustodian)[0]
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedPractitionerResourcesReferencedByAuthor() (practitioners []Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *d.IncludedPractitionerResourcesReferencedByAuthor
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedOrganizationResourcesReferencedByAuthor() (organizations []Organization, err error) {
	if d.IncludedOrganizationResourcesReferencedByAuthor == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *d.IncludedOrganizationResourcesReferencedByAuthor
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedDeviceResourcesReferencedByAuthor() (devices []Device, err error) {
	if d.IncludedDeviceResourcesReferencedByAuthor == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *d.IncludedDeviceResourcesReferencedByAuthor
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedPatientResourcesReferencedByAuthor() (patients []Patient, err error) {
	if d.IncludedPatientResourcesReferencedByAuthor == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *d.IncludedPatientResourcesReferencedByAuthor
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByAuthor() (relatedPeople []RelatedPerson, err error) {
	if d.IncludedRelatedPersonResourcesReferencedByAuthor == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *d.IncludedRelatedPersonResourcesReferencedByAuthor
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if d.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*d.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*d.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*d.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*d.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedDocumentReferenceResourceReferencedByRelatesto() (documentReference *DocumentReference, err error) {
	if d.IncludedDocumentReferenceResourcesReferencedByRelatesto == nil {
		err = errors.New("Included documentreferences not requested")
	} else if len(*d.IncludedDocumentReferenceResourcesReferencedByRelatesto) > 1 {
		err = fmt.Errorf("Expected 0 or 1 documentReference, but found %d", len(*d.IncludedDocumentReferenceResourcesReferencedByRelatesto))
	} else if len(*d.IncludedDocumentReferenceResourcesReferencedByRelatesto) == 1 {
		documentReference = &(*d.IncludedDocumentReferenceResourcesReferencedByRelatesto)[0]
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedConsentResourcesReferencingSource() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingSource == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingSource
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatesto() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatesto == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelatesto
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if d.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *d.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingItem
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if d.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *d.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if d.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *d.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if d.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *d.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if d.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *d.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if d.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *d.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if d.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *d.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if d.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *d.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedPractitionerResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedBySubject {
			rsc := (*d.IncludedPractitionerResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedGroupResourcesReferencedBySubject {
			rsc := (*d.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*d.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedPatientResourcesReferencedBySubject {
			rsc := (*d.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByAuthenticator != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByAuthenticator {
			rsc := (*d.IncludedPractitionerResourcesReferencedByAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByAuthenticator != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByAuthenticator {
			rsc := (*d.IncludedOrganizationResourcesReferencedByAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByCustodian != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByCustodian {
			rsc := (*d.IncludedOrganizationResourcesReferencedByCustodian)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*d.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*d.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*d.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*d.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for idx := range *d.IncludedRelatedPersonResourcesReferencedByAuthor {
			rsc := (*d.IncludedRelatedPersonResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *d.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*d.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDocumentReferenceResourcesReferencedByRelatesto != nil {
		for idx := range *d.IncludedDocumentReferenceResourcesReferencedByRelatesto {
			rsc := (*d.IncludedDocumentReferenceResourcesReferencedByRelatesto)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DocumentReferencePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingData {
			rsc := (*d.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingSource {
			rsc := (*d.RevIncludedConsentResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatesto != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatesto {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelatesto)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*d.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingSubject {
			rsc := (*d.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingTopic {
			rsc := (*d.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *d.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*d.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*d.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*d.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*d.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *d.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*d.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *d.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*d.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*d.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*d.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DocumentReferencePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedPractitionerResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedBySubject {
			rsc := (*d.IncludedPractitionerResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedGroupResourcesReferencedBySubject {
			rsc := (*d.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*d.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedPatientResourcesReferencedBySubject {
			rsc := (*d.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByAuthenticator != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByAuthenticator {
			rsc := (*d.IncludedPractitionerResourcesReferencedByAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByAuthenticator != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByAuthenticator {
			rsc := (*d.IncludedOrganizationResourcesReferencedByAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByCustodian != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByCustodian {
			rsc := (*d.IncludedOrganizationResourcesReferencedByCustodian)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*d.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*d.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*d.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*d.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for idx := range *d.IncludedRelatedPersonResourcesReferencedByAuthor {
			rsc := (*d.IncludedRelatedPersonResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *d.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*d.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDocumentReferenceResourcesReferencedByRelatesto != nil {
		for idx := range *d.IncludedDocumentReferenceResourcesReferencedByRelatesto {
			rsc := (*d.IncludedDocumentReferenceResourcesReferencedByRelatesto)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingData {
			rsc := (*d.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingSource {
			rsc := (*d.RevIncludedConsentResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatesto != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatesto {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelatesto)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*d.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingSubject {
			rsc := (*d.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingTopic {
			rsc := (*d.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *d.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*d.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*d.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*d.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*d.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *d.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*d.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *d.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*d.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*d.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*d.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
