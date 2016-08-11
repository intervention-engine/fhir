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

type Task struct {
	DomainResource `bson:",inline"`
	Identifier     *Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type           *CodeableConcept         `bson:"type,omitempty" json:"type,omitempty"`
	Description    string                   `bson:"description,omitempty" json:"description,omitempty"`
	PerformerType  []Coding                 `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Priority       string                   `bson:"priority,omitempty" json:"priority,omitempty"`
	Status         string                   `bson:"status,omitempty" json:"status,omitempty"`
	FailureReason  *CodeableConcept         `bson:"failureReason,omitempty" json:"failureReason,omitempty"`
	Subject        *Reference               `bson:"subject,omitempty" json:"subject,omitempty"`
	For            *Reference               `bson:"for,omitempty" json:"for,omitempty"`
	Definition     string                   `bson:"definition,omitempty" json:"definition,omitempty"`
	Created        *FHIRDateTime            `bson:"created,omitempty" json:"created,omitempty"`
	LastModified   *FHIRDateTime            `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	Creator        *Reference               `bson:"creator,omitempty" json:"creator,omitempty"`
	Owner          *Reference               `bson:"owner,omitempty" json:"owner,omitempty"`
	Parent         *Reference               `bson:"parent,omitempty" json:"parent,omitempty"`
	Input          []TaskParameterComponent `bson:"input,omitempty" json:"input,omitempty"`
	Output         []TaskOutputComponent    `bson:"output,omitempty" json:"output,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Task) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Task"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Task), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Task) GetBSON() (interface{}, error) {
	x.ResourceType = "Task"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "task" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type task Task

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Task) UnmarshalJSON(data []byte) (err error) {
	x2 := task{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Task(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Task) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Task"
	} else if x.ResourceType != "Task" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Task, instead received %s", x.ResourceType))
	}
	return nil
}

type TaskParameterComponent struct {
	BackboneElement      `bson:",inline"`
	Name                 string           `bson:"name,omitempty" json:"name,omitempty"`
	ValueAddress         *Address         `bson:"valueAddress,omitempty" json:"valueAddress,omitempty"`
	ValueAnnotation      *Annotation      `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueBase64Binary    string           `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCode            string           `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCoding          *Coding          `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint    *ContactPoint    `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty"`
	ValueDate            *FHIRDateTime    `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime        *FHIRDateTime    `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal         *float64         `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueHumanName       *HumanName       `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty"`
	ValueId              string           `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueIdentifier      *Identifier      `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty"`
	ValueInstant         *FHIRDateTime    `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueInteger         *int32           `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueMarkdown        string           `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty"`
	ValueMeta            *Meta            `bson:"valueMeta,omitempty" json:"valueMeta,omitempty"`
	ValueOid             string           `bson:"valueOid,omitempty" json:"valueOid,omitempty"`
	ValuePeriod          *Period          `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValuePositiveInt     *uint32          `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio           `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	ValueSampledData     *SampledData     `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueSignature       *Signature       `bson:"valueSignature,omitempty" json:"valueSignature,omitempty"`
	ValueString          string           `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueTime            *FHIRDateTime    `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueTiming          *Timing          `bson:"valueTiming,omitempty" json:"valueTiming,omitempty"`
	ValueUnsignedInt     *uint32          `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri             string           `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
}

type TaskOutputComponent struct {
	BackboneElement      `bson:",inline"`
	Name                 string           `bson:"name,omitempty" json:"name,omitempty"`
	ValueAddress         *Address         `bson:"valueAddress,omitempty" json:"valueAddress,omitempty"`
	ValueAnnotation      *Annotation      `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueBase64Binary    string           `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCode            string           `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCoding          *Coding          `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint    *ContactPoint    `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty"`
	ValueDate            *FHIRDateTime    `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime        *FHIRDateTime    `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal         *float64         `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueHumanName       *HumanName       `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty"`
	ValueId              string           `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueIdentifier      *Identifier      `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty"`
	ValueInstant         *FHIRDateTime    `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueInteger         *int32           `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueMarkdown        string           `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty"`
	ValueMeta            *Meta            `bson:"valueMeta,omitempty" json:"valueMeta,omitempty"`
	ValueOid             string           `bson:"valueOid,omitempty" json:"valueOid,omitempty"`
	ValuePeriod          *Period          `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValuePositiveInt     *uint32          `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio           `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	ValueSampledData     *SampledData     `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueSignature       *Signature       `bson:"valueSignature,omitempty" json:"valueSignature,omitempty"`
	ValueString          string           `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueTime            *FHIRDateTime    `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueTiming          *Timing          `bson:"valueTiming,omitempty" json:"valueTiming,omitempty"`
	ValueUnsignedInt     *uint32          `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri             string           `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
}

type TaskPlus struct {
	Task                     `bson:",inline"`
	TaskPlusRelatedResources `bson:",inline"`
}

type TaskPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByOwner                 *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByOwner,omitempty"`
	IncludedOrganizationResourcesReferencedByOwner                 *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOwner,omitempty"`
	IncludedDeviceResourcesReferencedByOwner                       *[]Device                `bson:"_includedDeviceResourcesReferencedByOwner,omitempty"`
	IncludedPatientResourcesReferencedByOwner                      *[]Patient               `bson:"_includedPatientResourcesReferencedByOwner,omitempty"`
	IncludedRelatedPersonResourcesReferencedByOwner                *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByOwner,omitempty"`
	IncludedTaskResourcesReferencedByParent                        *[]Task                  `bson:"_includedTaskResourcesReferencedByParent,omitempty"`
	IncludedPractitionerResourcesReferencedByCreator               *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByCreator,omitempty"`
	IncludedOrganizationResourcesReferencedByCreator               *[]Organization          `bson:"_includedOrganizationResourcesReferencedByCreator,omitempty"`
	IncludedDeviceResourcesReferencedByCreator                     *[]Device                `bson:"_includedDeviceResourcesReferencedByCreator,omitempty"`
	IncludedPatientResourcesReferencedByCreator                    *[]Patient               `bson:"_includedPatientResourcesReferencedByCreator,omitempty"`
	IncludedRelatedPersonResourcesReferencedByCreator              *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByCreator,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref     *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                   *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponsereference  *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponsereference,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequestreference   *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequestreference,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource     *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment        *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingParent                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingParent,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                     *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingItem                        *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                     *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                    *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated         *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject    *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestreference *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequestreference,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
}

func (t *TaskPlusRelatedResources) GetIncludedPractitionerResourceReferencedByOwner() (practitioner *Practitioner, err error) {
	if t.IncludedPractitionerResourcesReferencedByOwner == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*t.IncludedPractitionerResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*t.IncludedPractitionerResourcesReferencedByOwner))
	} else if len(*t.IncludedPractitionerResourcesReferencedByOwner) == 1 {
		practitioner = &(*t.IncludedPractitionerResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOwner() (organization *Organization, err error) {
	if t.IncludedOrganizationResourcesReferencedByOwner == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*t.IncludedOrganizationResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*t.IncludedOrganizationResourcesReferencedByOwner))
	} else if len(*t.IncludedOrganizationResourcesReferencedByOwner) == 1 {
		organization = &(*t.IncludedOrganizationResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedDeviceResourceReferencedByOwner() (device *Device, err error) {
	if t.IncludedDeviceResourcesReferencedByOwner == nil {
		err = errors.New("Included devices not requested")
	} else if len(*t.IncludedDeviceResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*t.IncludedDeviceResourcesReferencedByOwner))
	} else if len(*t.IncludedDeviceResourcesReferencedByOwner) == 1 {
		device = &(*t.IncludedDeviceResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPatientResourceReferencedByOwner() (patient *Patient, err error) {
	if t.IncludedPatientResourcesReferencedByOwner == nil {
		err = errors.New("Included patients not requested")
	} else if len(*t.IncludedPatientResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*t.IncludedPatientResourcesReferencedByOwner))
	} else if len(*t.IncludedPatientResourcesReferencedByOwner) == 1 {
		patient = &(*t.IncludedPatientResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByOwner() (relatedPerson *RelatedPerson, err error) {
	if t.IncludedRelatedPersonResourcesReferencedByOwner == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*t.IncludedRelatedPersonResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*t.IncludedRelatedPersonResourcesReferencedByOwner))
	} else if len(*t.IncludedRelatedPersonResourcesReferencedByOwner) == 1 {
		relatedPerson = &(*t.IncludedRelatedPersonResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedTaskResourceReferencedByParent() (task *Task, err error) {
	if t.IncludedTaskResourcesReferencedByParent == nil {
		err = errors.New("Included tasks not requested")
	} else if len(*t.IncludedTaskResourcesReferencedByParent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 task, but found %d", len(*t.IncludedTaskResourcesReferencedByParent))
	} else if len(*t.IncludedTaskResourcesReferencedByParent) == 1 {
		task = &(*t.IncludedTaskResourcesReferencedByParent)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPractitionerResourceReferencedByCreator() (practitioner *Practitioner, err error) {
	if t.IncludedPractitionerResourcesReferencedByCreator == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*t.IncludedPractitionerResourcesReferencedByCreator) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*t.IncludedPractitionerResourcesReferencedByCreator))
	} else if len(*t.IncludedPractitionerResourcesReferencedByCreator) == 1 {
		practitioner = &(*t.IncludedPractitionerResourcesReferencedByCreator)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedOrganizationResourceReferencedByCreator() (organization *Organization, err error) {
	if t.IncludedOrganizationResourcesReferencedByCreator == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*t.IncludedOrganizationResourcesReferencedByCreator) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*t.IncludedOrganizationResourcesReferencedByCreator))
	} else if len(*t.IncludedOrganizationResourcesReferencedByCreator) == 1 {
		organization = &(*t.IncludedOrganizationResourcesReferencedByCreator)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedDeviceResourceReferencedByCreator() (device *Device, err error) {
	if t.IncludedDeviceResourcesReferencedByCreator == nil {
		err = errors.New("Included devices not requested")
	} else if len(*t.IncludedDeviceResourcesReferencedByCreator) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*t.IncludedDeviceResourcesReferencedByCreator))
	} else if len(*t.IncludedDeviceResourcesReferencedByCreator) == 1 {
		device = &(*t.IncludedDeviceResourcesReferencedByCreator)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPatientResourceReferencedByCreator() (patient *Patient, err error) {
	if t.IncludedPatientResourcesReferencedByCreator == nil {
		err = errors.New("Included patients not requested")
	} else if len(*t.IncludedPatientResourcesReferencedByCreator) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*t.IncludedPatientResourcesReferencedByCreator))
	} else if len(*t.IncludedPatientResourcesReferencedByCreator) == 1 {
		patient = &(*t.IncludedPatientResourcesReferencedByCreator)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByCreator() (relatedPerson *RelatedPerson, err error) {
	if t.IncludedRelatedPersonResourcesReferencedByCreator == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*t.IncludedRelatedPersonResourcesReferencedByCreator) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*t.IncludedRelatedPersonResourcesReferencedByCreator))
	} else if len(*t.IncludedRelatedPersonResourcesReferencedByCreator) == 1 {
		relatedPerson = &(*t.IncludedRelatedPersonResourcesReferencedByCreator)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if t.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *t.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if t.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *t.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if t.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *t.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if t.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *t.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if t.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *t.RevIncludedPaymentNoticeResourcesReferencingResponsereference
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if t.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *t.RevIncludedPaymentNoticeResourcesReferencingRequestreference
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if t.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *t.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if t.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *t.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if t.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *t.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if t.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *t.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedTaskResourcesReferencingParent() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingParent == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingParent
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if t.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *t.RevIncludedListResourcesReferencingItem
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if t.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *t.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if t.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *t.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if t.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *t.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *t.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if t.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *t.RevIncludedProcessResponseResourcesReferencingRequestreference
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if t.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *t.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.IncludedPractitionerResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPractitionerResourcesReferencedByOwner {
			rsc := (*t.IncludedPractitionerResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByOwner {
			rsc := (*t.IncludedOrganizationResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedDeviceResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedDeviceResourcesReferencedByOwner {
			rsc := (*t.IncludedDeviceResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByOwner {
			rsc := (*t.IncludedPatientResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedRelatedPersonResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedRelatedPersonResourcesReferencedByOwner {
			rsc := (*t.IncludedRelatedPersonResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedTaskResourcesReferencedByParent != nil {
		for idx := range *t.IncludedTaskResourcesReferencedByParent {
			rsc := (*t.IncludedTaskResourcesReferencedByParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPractitionerResourcesReferencedByCreator != nil {
		for idx := range *t.IncludedPractitionerResourcesReferencedByCreator {
			rsc := (*t.IncludedPractitionerResourcesReferencedByCreator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByCreator != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByCreator {
			rsc := (*t.IncludedOrganizationResourcesReferencedByCreator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedDeviceResourcesReferencedByCreator != nil {
		for idx := range *t.IncludedDeviceResourcesReferencedByCreator {
			rsc := (*t.IncludedDeviceResourcesReferencedByCreator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByCreator != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByCreator {
			rsc := (*t.IncludedPatientResourcesReferencedByCreator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedRelatedPersonResourcesReferencedByCreator != nil {
		for idx := range *t.IncludedRelatedPersonResourcesReferencedByCreator {
			rsc := (*t.IncludedRelatedPersonResourcesReferencedByCreator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (t *TaskPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*t.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingSubject {
			rsc := (*t.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingTopic {
			rsc := (*t.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *t.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*t.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *t.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*t.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingParent != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingParent {
			rsc := (*t.RevIncludedTaskResourcesReferencingParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*t.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *t.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*t.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*t.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*t.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *t.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*t.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*t.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *t.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*t.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *t.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*t.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (t *TaskPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.IncludedPractitionerResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPractitionerResourcesReferencedByOwner {
			rsc := (*t.IncludedPractitionerResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByOwner {
			rsc := (*t.IncludedOrganizationResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedDeviceResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedDeviceResourcesReferencedByOwner {
			rsc := (*t.IncludedDeviceResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByOwner {
			rsc := (*t.IncludedPatientResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedRelatedPersonResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedRelatedPersonResourcesReferencedByOwner {
			rsc := (*t.IncludedRelatedPersonResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedTaskResourcesReferencedByParent != nil {
		for idx := range *t.IncludedTaskResourcesReferencedByParent {
			rsc := (*t.IncludedTaskResourcesReferencedByParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPractitionerResourcesReferencedByCreator != nil {
		for idx := range *t.IncludedPractitionerResourcesReferencedByCreator {
			rsc := (*t.IncludedPractitionerResourcesReferencedByCreator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByCreator != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByCreator {
			rsc := (*t.IncludedOrganizationResourcesReferencedByCreator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedDeviceResourcesReferencedByCreator != nil {
		for idx := range *t.IncludedDeviceResourcesReferencedByCreator {
			rsc := (*t.IncludedDeviceResourcesReferencedByCreator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByCreator != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByCreator {
			rsc := (*t.IncludedPatientResourcesReferencedByCreator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedRelatedPersonResourcesReferencedByCreator != nil {
		for idx := range *t.IncludedRelatedPersonResourcesReferencedByCreator {
			rsc := (*t.IncludedRelatedPersonResourcesReferencedByCreator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*t.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingSubject {
			rsc := (*t.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingTopic {
			rsc := (*t.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *t.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*t.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *t.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*t.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingParent != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingParent {
			rsc := (*t.RevIncludedTaskResourcesReferencingParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*t.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *t.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*t.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*t.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*t.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *t.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*t.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*t.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *t.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*t.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *t.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*t.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
