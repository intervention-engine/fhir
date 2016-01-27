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

type Observation struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                               `bson:"status,omitempty" json:"status,omitempty"`
	Category             *CodeableConcept                     `bson:"category,omitempty" json:"category,omitempty"`
	Code                 *CodeableConcept                     `bson:"code,omitempty" json:"code,omitempty"`
	Subject              *Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter            *Reference                           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	EffectiveDateTime    *FHIRDateTime                        `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod      *Period                              `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Issued               *FHIRDateTime                        `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer            []Reference                          `bson:"performer,omitempty" json:"performer,omitempty"`
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
	Comments             string                               `bson:"comments,omitempty" json:"comments,omitempty"`
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
	x := struct {
		ResourceType string `json:"resourceType"`
		Observation
	}{
		ResourceType: "Observation",
		Observation:  *resource,
	}
	return json.Marshal(x)
}

// The "observation" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type observation Observation

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Observation) UnmarshalJSON(data []byte) (err error) {
	x2 := observation{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Observation(x2)
	}
	return
}

type ObservationReferenceRangeComponent struct {
	Low     *Quantity        `bson:"low,omitempty" json:"low,omitempty"`
	High    *Quantity        `bson:"high,omitempty" json:"high,omitempty"`
	Meaning *CodeableConcept `bson:"meaning,omitempty" json:"meaning,omitempty"`
	Age     *Range           `bson:"age,omitempty" json:"age,omitempty"`
	Text    string           `bson:"text,omitempty" json:"text,omitempty"`
}

type ObservationRelatedComponent struct {
	Type   string     `bson:"type,omitempty" json:"type,omitempty"`
	Target *Reference `bson:"target,omitempty" json:"target,omitempty"`
}

type ObservationComponentComponent struct {
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
	ReferenceRange       []ObservationReferenceRangeComponent `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}

type ObservationPlus struct {
	Observation             `bson:",inline"`
	ObservationPlusIncludes `bson:",inline"`
}

type ObservationPlusIncludes struct {
	IncludedSubjectGroupResources                       *[]Group                 `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectDeviceResources                      *[]Device                `bson:"_includedSubjectDeviceResources,omitempty"`
	IncludedSubjectPatientResources                     *[]Patient               `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedSubjectLocationResources                    *[]Location              `bson:"_includedSubjectLocationResources,omitempty"`
	IncludedPatientResources                            *[]Patient               `bson:"_includedPatientResources,omitempty"`
	IncludedSpecimenResources                           *[]Specimen              `bson:"_includedSpecimenResources,omitempty"`
	IncludedPerformerPractitionerResources              *[]Practitioner          `bson:"_includedPerformerPractitionerResources,omitempty"`
	IncludedPerformerOrganizationResources              *[]Organization          `bson:"_includedPerformerOrganizationResources,omitempty"`
	IncludedPerformerPatientResources                   *[]Patient               `bson:"_includedPerformerPatientResources,omitempty"`
	IncludedPerformerRelatedPersonResources             *[]RelatedPerson         `bson:"_includedPerformerRelatedPersonResources,omitempty"`
	IncludedEncounterResources                          *[]Encounter             `bson:"_includedEncounterResources,omitempty"`
	IncludedRelatedtargetObservationResources           *[]Observation           `bson:"_includedRelatedtargetObservationResources,omitempty"`
	IncludedRelatedtargetQuestionnaireResponseResources *[]QuestionnaireResponse `bson:"_includedRelatedtargetQuestionnaireResponseResources,omitempty"`
	IncludedDeviceDeviceResources                       *[]Device                `bson:"_includedDeviceDeviceResources,omitempty"`
	IncludedDeviceDeviceMetricResources                 *[]DeviceMetric          `bson:"_includedDeviceDeviceMetricResources,omitempty"`
}

func (o *ObservationPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if o.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*o.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*o.IncludedSubjectGroupResources))
	} else if len(*o.IncludedSubjectGroupResources) == 1 {
		group = &(*o.IncludedSubjectGroupResources)[0]
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedSubjectDeviceResource() (device *Device, err error) {
	if o.IncludedSubjectDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedSubjectDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedSubjectDeviceResources))
	} else if len(*o.IncludedSubjectDeviceResources) == 1 {
		device = &(*o.IncludedSubjectDeviceResources)[0]
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if o.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*o.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*o.IncludedSubjectPatientResources))
	} else if len(*o.IncludedSubjectPatientResources) == 1 {
		patient = &(*o.IncludedSubjectPatientResources)[0]
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedSubjectLocationResource() (location *Location, err error) {
	if o.IncludedSubjectLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*o.IncludedSubjectLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*o.IncludedSubjectLocationResources))
	} else if len(*o.IncludedSubjectLocationResources) == 1 {
		location = &(*o.IncludedSubjectLocationResources)[0]
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if o.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*o.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*o.IncludedPatientResources))
	} else if len(*o.IncludedPatientResources) == 1 {
		patient = &(*o.IncludedPatientResources)[0]
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedSpecimenResource() (specimen *Specimen, err error) {
	if o.IncludedSpecimenResources == nil {
		err = errors.New("Included specimen not requested")
	} else if len(*o.IncludedSpecimenResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 specimen, but found %d", len(*o.IncludedSpecimenResources))
	} else if len(*o.IncludedSpecimenResources) == 1 {
		specimen = &(*o.IncludedSpecimenResources)[0]
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedPerformerPractitionerResources() (practitioners []Practitioner, err error) {
	if o.IncludedPerformerPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *o.IncludedPerformerPractitionerResources
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedPerformerOrganizationResources() (organizations []Organization, err error) {
	if o.IncludedPerformerOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *o.IncludedPerformerOrganizationResources
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedPerformerPatientResources() (patients []Patient, err error) {
	if o.IncludedPerformerPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *o.IncludedPerformerPatientResources
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedPerformerRelatedPersonResources() (relatedPeople []RelatedPerson, err error) {
	if o.IncludedPerformerRelatedPersonResources == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *o.IncludedPerformerRelatedPersonResources
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if o.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*o.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*o.IncludedEncounterResources))
	} else if len(*o.IncludedEncounterResources) == 1 {
		encounter = &(*o.IncludedEncounterResources)[0]
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedRelatedtargetObservationResource() (observation *Observation, err error) {
	if o.IncludedRelatedtargetObservationResources == nil {
		err = errors.New("Included observations not requested")
	} else if len(*o.IncludedRelatedtargetObservationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 observation, but found %d", len(*o.IncludedRelatedtargetObservationResources))
	} else if len(*o.IncludedRelatedtargetObservationResources) == 1 {
		observation = &(*o.IncludedRelatedtargetObservationResources)[0]
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedRelatedtargetQuestionnaireResponseResource() (questionnaireResponse *QuestionnaireResponse, err error) {
	if o.IncludedRelatedtargetQuestionnaireResponseResources == nil {
		err = errors.New("Included questionnaireresponses not requested")
	} else if len(*o.IncludedRelatedtargetQuestionnaireResponseResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 questionnaireResponse, but found %d", len(*o.IncludedRelatedtargetQuestionnaireResponseResources))
	} else if len(*o.IncludedRelatedtargetQuestionnaireResponseResources) == 1 {
		questionnaireResponse = &(*o.IncludedRelatedtargetQuestionnaireResponseResources)[0]
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedDeviceDeviceResource() (device *Device, err error) {
	if o.IncludedDeviceDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedDeviceDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedDeviceDeviceResources))
	} else if len(*o.IncludedDeviceDeviceResources) == 1 {
		device = &(*o.IncludedDeviceDeviceResources)[0]
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedDeviceDeviceMetricResource() (deviceMetric *DeviceMetric, err error) {
	if o.IncludedDeviceDeviceMetricResources == nil {
		err = errors.New("Included devicemetrics not requested")
	} else if len(*o.IncludedDeviceDeviceMetricResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceMetric, but found %d", len(*o.IncludedDeviceDeviceMetricResources))
	} else if len(*o.IncludedDeviceDeviceMetricResources) == 1 {
		deviceMetric = &(*o.IncludedDeviceDeviceMetricResources)[0]
	}
	return
}

func (o *ObservationPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedSubjectGroupResources != nil {
		for _, r := range *o.IncludedSubjectGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedSubjectDeviceResources != nil {
		for _, r := range *o.IncludedSubjectDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedSubjectPatientResources != nil {
		for _, r := range *o.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedSubjectLocationResources != nil {
		for _, r := range *o.IncludedSubjectLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPatientResources != nil {
		for _, r := range *o.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedSpecimenResources != nil {
		for _, r := range *o.IncludedSpecimenResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPerformerPractitionerResources != nil {
		for _, r := range *o.IncludedPerformerPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPerformerOrganizationResources != nil {
		for _, r := range *o.IncludedPerformerOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPerformerPatientResources != nil {
		for _, r := range *o.IncludedPerformerPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPerformerRelatedPersonResources != nil {
		for _, r := range *o.IncludedPerformerRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedEncounterResources != nil {
		for _, r := range *o.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedRelatedtargetObservationResources != nil {
		for _, r := range *o.IncludedRelatedtargetObservationResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedRelatedtargetQuestionnaireResponseResources != nil {
		for _, r := range *o.IncludedRelatedtargetQuestionnaireResponseResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedDeviceDeviceResources != nil {
		for _, r := range *o.IncludedDeviceDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedDeviceDeviceMetricResources != nil {
		for _, r := range *o.IncludedDeviceDeviceMetricResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
