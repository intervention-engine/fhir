// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
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

type Observation struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn              []Reference                          `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Status               string                               `bson:"status,omitempty" json:"status,omitempty"`
	Category             []CodeableConcept                    `bson:"category,omitempty" json:"category,omitempty"`
	Code                 *CodeableConcept                     `bson:"code,omitempty" json:"code,omitempty"`
	Subject              *Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Context              *Reference                           `bson:"context,omitempty" json:"context,omitempty"`
	EffectiveDateTime    *FHIRDateTime                        `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod      *Period                              `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Issued               *FHIRDateTime                        `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer            []Reference                          `bson:"performer,omitempty" json:"performer,omitempty"`
	ValueQuantity        *Quantity                            `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept                     `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          string                               `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean         *bool                                `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueRange           *Range                               `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                               `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                         `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueAttachment      *Attachment                          `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueTime            *FHIRDateTime                        `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime        *FHIRDateTime                        `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                              `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept                     `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation       *CodeableConcept                     `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	Comment              string                               `bson:"comment,omitempty" json:"comment,omitempty"`
	BodySite             *CodeableConcept                     `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Method               *CodeableConcept                     `bson:"method,omitempty" json:"method,omitempty"`
	Specimen             *Reference                           `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Device               *Reference                           `bson:"device,omitempty" json:"device,omitempty"`
	ReferenceRange       []ObservationReferenceRangeComponent `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
	Related              []ObservationRelatedComponent        `bson:"related,omitempty" json:"related,omitempty"`
	Component            []ObservationComponentComponent      `bson:"component,omitempty" json:"component,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Observation) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Observation"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Observation), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Observation) GetBSON() (interface{}, error) {
	x.ResourceType = "Observation"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "observation" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type observation Observation

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Observation) UnmarshalJSON(data []byte) (err error) {
	x2 := observation{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = Observation(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Observation) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Observation"
	} else if x.ResourceType != "Observation" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Observation, instead received %s", x.ResourceType))
	}
	return nil
}

type ObservationReferenceRangeComponent struct {
	BackboneElement `bson:",inline"`
	Low             *Quantity         `bson:"low,omitempty" json:"low,omitempty"`
	High            *Quantity         `bson:"high,omitempty" json:"high,omitempty"`
	Type            *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	AppliesTo       []CodeableConcept `bson:"appliesTo,omitempty" json:"appliesTo,omitempty"`
	Age             *Range            `bson:"age,omitempty" json:"age,omitempty"`
	Text            string            `bson:"text,omitempty" json:"text,omitempty"`
}

type ObservationRelatedComponent struct {
	BackboneElement `bson:",inline"`
	Type            string     `bson:"type,omitempty" json:"type,omitempty"`
	Target          *Reference `bson:"target,omitempty" json:"target,omitempty"`
}

type ObservationComponentComponent struct {
	BackboneElement      `bson:",inline"`
	Code                 *CodeableConcept                     `bson:"code,omitempty" json:"code,omitempty"`
	ValueQuantity        *Quantity                            `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept                     `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          string                               `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueRange           *Range                               `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                               `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                         `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueAttachment      *Attachment                          `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueTime            *FHIRDateTime                        `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime        *FHIRDateTime                        `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                              `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept                     `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation       *CodeableConcept                     `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	ReferenceRange       []ObservationReferenceRangeComponent `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}

type ObservationPlus struct {
	Observation                     `bson:",inline"`
	ObservationPlusRelatedResources `bson:",inline"`
}

type ObservationPlusRelatedResources struct {
	IncludedGroupResourcesReferencedBySubject                            *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                           *[]Device                     `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                          *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                         *[]Location                   `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedReferralRequestResourcesReferencedByBasedon                  *[]ReferralRequest            `bson:"_includedReferralRequestResourcesReferencedByBasedon,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                         *[]CarePlan                   `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedMedicationRequestResourcesReferencedByBasedon                *[]MedicationRequest          `bson:"_includedMedicationRequestResourcesReferencedByBasedon,omitempty"`
	IncludedNutritionOrderResourcesReferencedByBasedon                   *[]NutritionOrder             `bson:"_includedNutritionOrderResourcesReferencedByBasedon,omitempty"`
	IncludedProcedureRequestResourcesReferencedByBasedon                 *[]ProcedureRequest           `bson:"_includedProcedureRequestResourcesReferencedByBasedon,omitempty"`
	IncludedDeviceRequestResourcesReferencedByBasedon                    *[]DeviceRequest              `bson:"_includedDeviceRequestResourcesReferencedByBasedon,omitempty"`
	IncludedImmunizationRecommendationResourcesReferencedByBasedon       *[]ImmunizationRecommendation `bson:"_includedImmunizationRecommendationResourcesReferencedByBasedon,omitempty"`
	IncludedPatientResourcesReferencedByPatient                          *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedSpecimenResourcesReferencedBySpecimen                        *[]Specimen                   `bson:"_includedSpecimenResourcesReferencedBySpecimen,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByContext                    *[]EpisodeOfCare              `bson:"_includedEpisodeOfCareResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByContext                        *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByContext,omitempty"`
	IncludedPractitionerResourcesReferencedByPerformer                   *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer                   *[]Organization               `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                        *[]Patient                    `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer                  *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                      *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedObservationResourcesReferencedByRelatedtarget                *[]Observation                `bson:"_includedObservationResourcesReferencedByRelatedtarget,omitempty"`
	IncludedSequenceResourcesReferencedByRelatedtarget                   *[]Sequence                   `bson:"_includedSequenceResourcesReferencedByRelatedtarget,omitempty"`
	IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget      *[]QuestionnaireResponse      `bson:"_includedQuestionnaireResponseResourcesReferencedByRelatedtarget,omitempty"`
	IncludedDeviceResourcesReferencedByDevice                            *[]Device                     `bson:"_includedDeviceResourcesReferencedByDevice,omitempty"`
	IncludedDeviceMetricResourcesReferencedByDevice                      *[]DeviceMetric               `bson:"_includedDeviceMetricResourcesReferencedByDevice,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref            *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref            *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                      *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                      *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                      *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                    *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                    *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                     *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1                 *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2                 *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref           *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingSubject                       *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                     *[]Contract                   `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                  *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                 *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource           *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedChargeItemResourcesReferencingService                     *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingService,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingSuccessor            *[]ServiceDefinition          `bson:"_revIncludedServiceDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDerivedfrom          *[]ServiceDefinition          `bson:"_revIncludedServiceDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingPredecessor          *[]ServiceDefinition          `bson:"_revIncludedServiceDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingComposedof           *[]ServiceDefinition          `bson:"_revIncludedServiceDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDependson            *[]ServiceDefinition          `bson:"_revIncludedServiceDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                   *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                  *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor           *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom         *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor         *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof          *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1      *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2      *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedRequestGroupResourcesReferencingDefinition                *[]RequestGroup               `bson:"_revIncludedRequestGroupResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                  *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest             *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                    *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref                   *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                      *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                           *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                           *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedProcedureResourcesReferencingPartof                       *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingPartof,omitempty"`
	RevIncludedListResourcesReferencingItem                              *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedImmunizationResourcesReferencingReaction                  *[]Immunization               `bson:"_revIncludedImmunizationResourcesReferencingReaction,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces              *[]ProcedureRequest           `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon               *[]ProcedureRequest           `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedObservationResourcesReferencingRelatedtarget              *[]Observation                `bson:"_revIncludedObservationResourcesReferencingRelatedtarget,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                    *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                    *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                     *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingPartof             *[]MedicationStatement        `bson:"_revIncludedMedicationStatementResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon           *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                          *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingResult                *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingResult,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                      *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail               *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                    *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                      *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated               *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingParent           *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingParent,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject          *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest                *[]ProcessResponse            `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingFindingref          *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingFindingref,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingInvestigation       *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingInvestigation,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom             *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor             *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof              *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1          *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2          *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (o *ObservationPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if o.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*o.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*o.IncludedGroupResourcesReferencedBySubject))
	} else if len(*o.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*o.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if o.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*o.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*o.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if o.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*o.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*o.IncludedPatientResourcesReferencedBySubject))
	} else if len(*o.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*o.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if o.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*o.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*o.IncludedLocationResourcesReferencedBySubject))
	} else if len(*o.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*o.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedReferralRequestResourcesReferencedByBasedon() (referralRequests []ReferralRequest, err error) {
	if o.IncludedReferralRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *o.IncludedReferralRequestResourcesReferencedByBasedon
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByBasedon() (carePlans []CarePlan, err error) {
	if o.IncludedCarePlanResourcesReferencedByBasedon == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *o.IncludedCarePlanResourcesReferencedByBasedon
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedMedicationRequestResourcesReferencedByBasedon() (medicationRequests []MedicationRequest, err error) {
	if o.IncludedMedicationRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included medicationRequests not requested")
	} else {
		medicationRequests = *o.IncludedMedicationRequestResourcesReferencedByBasedon
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedNutritionOrderResourcesReferencedByBasedon() (nutritionOrders []NutritionOrder, err error) {
	if o.IncludedNutritionOrderResourcesReferencedByBasedon == nil {
		err = errors.New("Included nutritionOrders not requested")
	} else {
		nutritionOrders = *o.IncludedNutritionOrderResourcesReferencedByBasedon
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedProcedureRequestResourcesReferencedByBasedon() (procedureRequests []ProcedureRequest, err error) {
	if o.IncludedProcedureRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included procedureRequests not requested")
	} else {
		procedureRequests = *o.IncludedProcedureRequestResourcesReferencedByBasedon
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedDeviceRequestResourcesReferencedByBasedon() (deviceRequests []DeviceRequest, err error) {
	if o.IncludedDeviceRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included deviceRequests not requested")
	} else {
		deviceRequests = *o.IncludedDeviceRequestResourcesReferencedByBasedon
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedImmunizationRecommendationResourcesReferencedByBasedon() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if o.IncludedImmunizationRecommendationResourcesReferencedByBasedon == nil {
		err = errors.New("Included immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *o.IncludedImmunizationRecommendationResourcesReferencedByBasedon
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if o.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*o.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*o.IncludedPatientResourcesReferencedByPatient))
	} else if len(*o.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*o.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedSpecimenResourceReferencedBySpecimen() (specimen *Specimen, err error) {
	if o.IncludedSpecimenResourcesReferencedBySpecimen == nil {
		err = errors.New("Included specimen not requested")
	} else if len(*o.IncludedSpecimenResourcesReferencedBySpecimen) > 1 {
		err = fmt.Errorf("Expected 0 or 1 specimen, but found %d", len(*o.IncludedSpecimenResourcesReferencedBySpecimen))
	} else if len(*o.IncludedSpecimenResourcesReferencedBySpecimen) == 1 {
		specimen = &(*o.IncludedSpecimenResourcesReferencedBySpecimen)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedEpisodeOfCareResourceReferencedByContext() (episodeOfCare *EpisodeOfCare, err error) {
	if o.IncludedEpisodeOfCareResourcesReferencedByContext == nil {
		err = errors.New("Included episodeofcares not requested")
	} else if len(*o.IncludedEpisodeOfCareResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 episodeOfCare, but found %d", len(*o.IncludedEpisodeOfCareResourcesReferencedByContext))
	} else if len(*o.IncludedEpisodeOfCareResourcesReferencedByContext) == 1 {
		episodeOfCare = &(*o.IncludedEpisodeOfCareResourcesReferencedByContext)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedEncounterResourceReferencedByContext() (encounter *Encounter, err error) {
	if o.IncludedEncounterResourcesReferencedByContext == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*o.IncludedEncounterResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*o.IncludedEncounterResourcesReferencedByContext))
	} else if len(*o.IncludedEncounterResourcesReferencedByContext) == 1 {
		encounter = &(*o.IncludedEncounterResourcesReferencedByContext)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByPerformer() (practitioners []Practitioner, err error) {
	if o.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *o.IncludedPractitionerResourcesReferencedByPerformer
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByPerformer() (organizations []Organization, err error) {
	if o.IncludedOrganizationResourcesReferencedByPerformer == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *o.IncludedOrganizationResourcesReferencedByPerformer
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedPatientResourcesReferencedByPerformer() (patients []Patient, err error) {
	if o.IncludedPatientResourcesReferencedByPerformer == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *o.IncludedPatientResourcesReferencedByPerformer
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByPerformer() (relatedPeople []RelatedPerson, err error) {
	if o.IncludedRelatedPersonResourcesReferencedByPerformer == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *o.IncludedRelatedPersonResourcesReferencedByPerformer
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if o.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*o.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*o.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*o.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*o.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedObservationResourceReferencedByRelatedtarget() (observation *Observation, err error) {
	if o.IncludedObservationResourcesReferencedByRelatedtarget == nil {
		err = errors.New("Included observations not requested")
	} else if len(*o.IncludedObservationResourcesReferencedByRelatedtarget) > 1 {
		err = fmt.Errorf("Expected 0 or 1 observation, but found %d", len(*o.IncludedObservationResourcesReferencedByRelatedtarget))
	} else if len(*o.IncludedObservationResourcesReferencedByRelatedtarget) == 1 {
		observation = &(*o.IncludedObservationResourcesReferencedByRelatedtarget)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedSequenceResourceReferencedByRelatedtarget() (sequence *Sequence, err error) {
	if o.IncludedSequenceResourcesReferencedByRelatedtarget == nil {
		err = errors.New("Included sequences not requested")
	} else if len(*o.IncludedSequenceResourcesReferencedByRelatedtarget) > 1 {
		err = fmt.Errorf("Expected 0 or 1 sequence, but found %d", len(*o.IncludedSequenceResourcesReferencedByRelatedtarget))
	} else if len(*o.IncludedSequenceResourcesReferencedByRelatedtarget) == 1 {
		sequence = &(*o.IncludedSequenceResourcesReferencedByRelatedtarget)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedQuestionnaireResponseResourceReferencedByRelatedtarget() (questionnaireResponse *QuestionnaireResponse, err error) {
	if o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget == nil {
		err = errors.New("Included questionnaireresponses not requested")
	} else if len(*o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget) > 1 {
		err = fmt.Errorf("Expected 0 or 1 questionnaireResponse, but found %d", len(*o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget))
	} else if len(*o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget) == 1 {
		questionnaireResponse = &(*o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedDeviceResourceReferencedByDevice() (device *Device, err error) {
	if o.IncludedDeviceResourcesReferencedByDevice == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedDeviceResourcesReferencedByDevice) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedDeviceResourcesReferencedByDevice))
	} else if len(*o.IncludedDeviceResourcesReferencedByDevice) == 1 {
		device = &(*o.IncludedDeviceResourcesReferencedByDevice)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedDeviceMetricResourceReferencedByDevice() (deviceMetric *DeviceMetric, err error) {
	if o.IncludedDeviceMetricResourcesReferencedByDevice == nil {
		err = errors.New("Included devicemetrics not requested")
	} else if len(*o.IncludedDeviceMetricResourcesReferencedByDevice) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceMetric, but found %d", len(*o.IncludedDeviceMetricResourcesReferencedByDevice))
	} else if len(*o.IncludedDeviceMetricResourcesReferencedByDevice) == 1 {
		deviceMetric = &(*o.IncludedDeviceMetricResourcesReferencedByDevice)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if o.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *o.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingService() (chargeItems []ChargeItem, err error) {
	if o.RevIncludedChargeItemResourcesReferencingService == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *o.RevIncludedChargeItemResourcesReferencingService
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if o.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *o.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if o.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *o.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if o.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *o.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if o.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *o.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if o.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *o.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if o.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *o.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if o.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *o.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if o.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *o.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPartof() (procedures []Procedure, err error) {
	if o.RevIncludedProcedureResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *o.RevIncludedProcedureResourcesReferencingPartof
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if o.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *o.RevIncludedListResourcesReferencingItem
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingReaction() (immunizations []Immunization, err error) {
	if o.RevIncludedImmunizationResourcesReferencingReaction == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *o.RevIncludedImmunizationResourcesReferencingReaction
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if o.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *o.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if o.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *o.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingRelatedtarget() (observations []Observation, err error) {
	if o.RevIncludedObservationResourcesReferencingRelatedtarget == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *o.RevIncludedObservationResourcesReferencingRelatedtarget
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingPartof() (medicationStatements []MedicationStatement, err error) {
	if o.RevIncludedMedicationStatementResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *o.RevIncludedMedicationStatementResourcesReferencingPartof
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if o.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *o.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if o.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *o.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingResult() (diagnosticReports []DiagnosticReport, err error) {
	if o.RevIncludedDiagnosticReportResourcesReferencingResult == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *o.RevIncludedDiagnosticReportResourcesReferencingResult
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if o.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *o.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *o.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingParent() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingParent == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingParent
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if o.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *o.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingFindingref() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingFindingref == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingFindingref
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingInvestigation() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingInvestigation == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingInvestigation
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedGroupResourcesReferencedBySubject {
			rsc := (*o.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*o.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedPatientResourcesReferencedBySubject {
			rsc := (*o.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedLocationResourcesReferencedBySubject {
			rsc := (*o.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedReferralRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedReferralRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedReferralRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*o.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedMedicationRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedMedicationRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedMedicationRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedNutritionOrderResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedNutritionOrderResourcesReferencedByBasedon {
			rsc := (*o.IncludedNutritionOrderResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedProcedureRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedProcedureRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedProcedureRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedDeviceRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedDeviceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedImmunizationRecommendationResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedImmunizationRecommendationResourcesReferencedByBasedon {
			rsc := (*o.IncludedImmunizationRecommendationResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *o.IncludedPatientResourcesReferencedByPatient {
			rsc := (*o.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedSpecimenResourcesReferencedBySpecimen != nil {
		for idx := range *o.IncludedSpecimenResourcesReferencedBySpecimen {
			rsc := (*o.IncludedSpecimenResourcesReferencedBySpecimen)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *o.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*o.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *o.IncludedEncounterResourcesReferencedByContext {
			rsc := (*o.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*o.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*o.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*o.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*o.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *o.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*o.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedObservationResourcesReferencedByRelatedtarget != nil {
		for idx := range *o.IncludedObservationResourcesReferencedByRelatedtarget {
			rsc := (*o.IncludedObservationResourcesReferencedByRelatedtarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedSequenceResourcesReferencedByRelatedtarget != nil {
		for idx := range *o.IncludedSequenceResourcesReferencedByRelatedtarget {
			rsc := (*o.IncludedSequenceResourcesReferencedByRelatedtarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget != nil {
		for idx := range *o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget {
			rsc := (*o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedByDevice != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedByDevice {
			rsc := (*o.IncludedDeviceResourcesReferencedByDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceMetricResourcesReferencedByDevice != nil {
		for idx := range *o.IncludedDeviceMetricResourcesReferencedByDevice {
			rsc := (*o.IncludedDeviceMetricResourcesReferencedByDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *ObservationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*o.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*o.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*o.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSubject {
			rsc := (*o.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*o.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *o.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*o.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedChargeItemResourcesReferencingService != nil {
		for idx := range *o.RevIncludedChargeItemResourcesReferencingService {
			rsc := (*o.RevIncludedChargeItemResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*o.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *o.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*o.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*o.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*o.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*o.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*o.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedProcedureResourcesReferencingPartof {
			rsc := (*o.RevIncludedProcedureResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingReaction != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingReaction {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingReaction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *o.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*o.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingRelatedtarget != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingRelatedtarget {
			rsc := (*o.RevIncludedObservationResourcesReferencingRelatedtarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*o.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*o.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationStatementResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedMedicationStatementResourcesReferencingPartof {
			rsc := (*o.RevIncludedMedicationStatementResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingResult != nil {
		for idx := range *o.RevIncludedDiagnosticReportResourcesReferencingResult {
			rsc := (*o.RevIncludedDiagnosticReportResourcesReferencingResult)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *o.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*o.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*o.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*o.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*o.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingParent != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingParent {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingFindingref != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingFindingref {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingFindingref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *ObservationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedGroupResourcesReferencedBySubject {
			rsc := (*o.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*o.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedPatientResourcesReferencedBySubject {
			rsc := (*o.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedLocationResourcesReferencedBySubject {
			rsc := (*o.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedReferralRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedReferralRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedReferralRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*o.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedMedicationRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedMedicationRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedMedicationRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedNutritionOrderResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedNutritionOrderResourcesReferencedByBasedon {
			rsc := (*o.IncludedNutritionOrderResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedProcedureRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedProcedureRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedProcedureRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedDeviceRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedDeviceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedImmunizationRecommendationResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedImmunizationRecommendationResourcesReferencedByBasedon {
			rsc := (*o.IncludedImmunizationRecommendationResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *o.IncludedPatientResourcesReferencedByPatient {
			rsc := (*o.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedSpecimenResourcesReferencedBySpecimen != nil {
		for idx := range *o.IncludedSpecimenResourcesReferencedBySpecimen {
			rsc := (*o.IncludedSpecimenResourcesReferencedBySpecimen)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *o.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*o.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *o.IncludedEncounterResourcesReferencedByContext {
			rsc := (*o.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*o.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*o.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*o.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*o.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *o.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*o.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedObservationResourcesReferencedByRelatedtarget != nil {
		for idx := range *o.IncludedObservationResourcesReferencedByRelatedtarget {
			rsc := (*o.IncludedObservationResourcesReferencedByRelatedtarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedSequenceResourcesReferencedByRelatedtarget != nil {
		for idx := range *o.IncludedSequenceResourcesReferencedByRelatedtarget {
			rsc := (*o.IncludedSequenceResourcesReferencedByRelatedtarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget != nil {
		for idx := range *o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget {
			rsc := (*o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedByDevice != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedByDevice {
			rsc := (*o.IncludedDeviceResourcesReferencedByDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceMetricResourcesReferencedByDevice != nil {
		for idx := range *o.IncludedDeviceMetricResourcesReferencedByDevice {
			rsc := (*o.IncludedDeviceMetricResourcesReferencedByDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*o.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*o.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*o.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSubject {
			rsc := (*o.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*o.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *o.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*o.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedChargeItemResourcesReferencingService != nil {
		for idx := range *o.RevIncludedChargeItemResourcesReferencingService {
			rsc := (*o.RevIncludedChargeItemResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*o.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *o.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*o.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*o.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*o.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*o.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*o.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedProcedureResourcesReferencingPartof {
			rsc := (*o.RevIncludedProcedureResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingReaction != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingReaction {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingReaction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *o.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*o.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingRelatedtarget != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingRelatedtarget {
			rsc := (*o.RevIncludedObservationResourcesReferencingRelatedtarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*o.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*o.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationStatementResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedMedicationStatementResourcesReferencingPartof {
			rsc := (*o.RevIncludedMedicationStatementResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingResult != nil {
		for idx := range *o.RevIncludedDiagnosticReportResourcesReferencingResult {
			rsc := (*o.RevIncludedDiagnosticReportResourcesReferencingResult)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *o.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*o.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*o.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*o.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*o.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingParent != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingParent {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingFindingref != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingFindingref {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingFindingref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
