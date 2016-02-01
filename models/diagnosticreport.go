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

type DiagnosticReport struct {
	DomainResource    `bson:",inline"`
	Identifier        []Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string                           `bson:"status,omitempty" json:"status,omitempty"`
	Category          *CodeableConcept                 `bson:"category,omitempty" json:"category,omitempty"`
	Code              *CodeableConcept                 `bson:"code,omitempty" json:"code,omitempty"`
	Subject           *Reference                       `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter         *Reference                       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	EffectiveDateTime *FHIRDateTime                    `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *Period                          `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Issued            *FHIRDateTime                    `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer         *Reference                       `bson:"performer,omitempty" json:"performer,omitempty"`
	Request           []Reference                      `bson:"request,omitempty" json:"request,omitempty"`
	Specimen          []Reference                      `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Result            []Reference                      `bson:"result,omitempty" json:"result,omitempty"`
	ImagingStudy      []Reference                      `bson:"imagingStudy,omitempty" json:"imagingStudy,omitempty"`
	Image             []DiagnosticReportImageComponent `bson:"image,omitempty" json:"image,omitempty"`
	Conclusion        string                           `bson:"conclusion,omitempty" json:"conclusion,omitempty"`
	CodedDiagnosis    []CodeableConcept                `bson:"codedDiagnosis,omitempty" json:"codedDiagnosis,omitempty"`
	PresentedForm     []Attachment                     `bson:"presentedForm,omitempty" json:"presentedForm,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DiagnosticReport) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DiagnosticReport"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DiagnosticReport), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DiagnosticReport) GetBSON() (interface{}, error) {
	x.ResourceType = "DiagnosticReport"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "diagnosticReport" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type diagnosticReport DiagnosticReport

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DiagnosticReport) UnmarshalJSON(data []byte) (err error) {
	x2 := diagnosticReport{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DiagnosticReport(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DiagnosticReport) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DiagnosticReport"
	} else if x.ResourceType != "DiagnosticReport" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DiagnosticReport, instead received %s", x.ResourceType))
	}
	return nil
}

type DiagnosticReportImageComponent struct {
	Comment string     `bson:"comment,omitempty" json:"comment,omitempty"`
	Link    *Reference `bson:"link,omitempty" json:"link,omitempty"`
}

type DiagnosticReportPlus struct {
	DiagnosticReport             `bson:",inline"`
	DiagnosticReportPlusIncludes `bson:",inline"`
}

type DiagnosticReportPlusIncludes struct {
	IncludedImageResources                   *[]Media            `bson:"_includedImageResources,omitempty"`
	IncludedRequestReferralRequestResources  *[]ReferralRequest  `bson:"_includedRequestReferralRequestResources,omitempty"`
	IncludedRequestDiagnosticOrderResources  *[]DiagnosticOrder  `bson:"_includedRequestDiagnosticOrderResources,omitempty"`
	IncludedRequestProcedureRequestResources *[]ProcedureRequest `bson:"_includedRequestProcedureRequestResources,omitempty"`
	IncludedPerformerPractitionerResources   *[]Practitioner     `bson:"_includedPerformerPractitionerResources,omitempty"`
	IncludedPerformerOrganizationResources   *[]Organization     `bson:"_includedPerformerOrganizationResources,omitempty"`
	IncludedSubjectGroupResources            *[]Group            `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectDeviceResources           *[]Device           `bson:"_includedSubjectDeviceResources,omitempty"`
	IncludedSubjectPatientResources          *[]Patient          `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedSubjectLocationResources         *[]Location         `bson:"_includedSubjectLocationResources,omitempty"`
	IncludedEncounterResources               *[]Encounter        `bson:"_includedEncounterResources,omitempty"`
	IncludedResultResources                  *[]Observation      `bson:"_includedResultResources,omitempty"`
	IncludedPatientResources                 *[]Patient          `bson:"_includedPatientResources,omitempty"`
	IncludedSpecimenResources                *[]Specimen         `bson:"_includedSpecimenResources,omitempty"`
}

func (d *DiagnosticReportPlusIncludes) GetIncludedImageResource() (media *Media, err error) {
	if d.IncludedImageResources == nil {
		err = errors.New("Included media not requested")
	} else if len(*d.IncludedImageResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 media, but found %d", len(*d.IncludedImageResources))
	} else if len(*d.IncludedImageResources) == 1 {
		media = &(*d.IncludedImageResources)[0]
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedRequestReferralRequestResources() (referralRequests []ReferralRequest, err error) {
	if d.IncludedRequestReferralRequestResources == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *d.IncludedRequestReferralRequestResources
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedRequestDiagnosticOrderResources() (diagnosticOrders []DiagnosticOrder, err error) {
	if d.IncludedRequestDiagnosticOrderResources == nil {
		err = errors.New("Included diagnosticOrders not requested")
	} else {
		diagnosticOrders = *d.IncludedRequestDiagnosticOrderResources
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedRequestProcedureRequestResources() (procedureRequests []ProcedureRequest, err error) {
	if d.IncludedRequestProcedureRequestResources == nil {
		err = errors.New("Included procedureRequests not requested")
	} else {
		procedureRequests = *d.IncludedRequestProcedureRequestResources
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedPerformerPractitionerResource() (practitioner *Practitioner, err error) {
	if d.IncludedPerformerPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPerformerPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPerformerPractitionerResources))
	} else if len(*d.IncludedPerformerPractitionerResources) == 1 {
		practitioner = &(*d.IncludedPerformerPractitionerResources)[0]
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedPerformerOrganizationResource() (organization *Organization, err error) {
	if d.IncludedPerformerOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*d.IncludedPerformerOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*d.IncludedPerformerOrganizationResources))
	} else if len(*d.IncludedPerformerOrganizationResources) == 1 {
		organization = &(*d.IncludedPerformerOrganizationResources)[0]
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if d.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedSubjectGroupResources))
	} else if len(*d.IncludedSubjectGroupResources) == 1 {
		group = &(*d.IncludedSubjectGroupResources)[0]
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedSubjectDeviceResource() (device *Device, err error) {
	if d.IncludedSubjectDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedSubjectDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedSubjectDeviceResources))
	} else if len(*d.IncludedSubjectDeviceResources) == 1 {
		device = &(*d.IncludedSubjectDeviceResources)[0]
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if d.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedSubjectPatientResources))
	} else if len(*d.IncludedSubjectPatientResources) == 1 {
		patient = &(*d.IncludedSubjectPatientResources)[0]
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedSubjectLocationResource() (location *Location, err error) {
	if d.IncludedSubjectLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*d.IncludedSubjectLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*d.IncludedSubjectLocationResources))
	} else if len(*d.IncludedSubjectLocationResources) == 1 {
		location = &(*d.IncludedSubjectLocationResources)[0]
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if d.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*d.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*d.IncludedEncounterResources))
	} else if len(*d.IncludedEncounterResources) == 1 {
		encounter = &(*d.IncludedEncounterResources)[0]
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedResultResources() (observations []Observation, err error) {
	if d.IncludedResultResources == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *d.IncludedResultResources
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if d.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResources))
	} else if len(*d.IncludedPatientResources) == 1 {
		patient = &(*d.IncludedPatientResources)[0]
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedSpecimenResources() (specimen []Specimen, err error) {
	if d.IncludedSpecimenResources == nil {
		err = errors.New("Included specimen not requested")
	} else {
		specimen = *d.IncludedSpecimenResources
	}
	return
}

func (d *DiagnosticReportPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedImageResources != nil {
		for _, r := range *d.IncludedImageResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedRequestReferralRequestResources != nil {
		for _, r := range *d.IncludedRequestReferralRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedRequestDiagnosticOrderResources != nil {
		for _, r := range *d.IncludedRequestDiagnosticOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedRequestProcedureRequestResources != nil {
		for _, r := range *d.IncludedRequestProcedureRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPerformerPractitionerResources != nil {
		for _, r := range *d.IncludedPerformerPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPerformerOrganizationResources != nil {
		for _, r := range *d.IncludedPerformerOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSubjectGroupResources != nil {
		for _, r := range *d.IncludedSubjectGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSubjectDeviceResources != nil {
		for _, r := range *d.IncludedSubjectDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSubjectPatientResources != nil {
		for _, r := range *d.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSubjectLocationResources != nil {
		for _, r := range *d.IncludedSubjectLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedEncounterResources != nil {
		for _, r := range *d.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedResultResources != nil {
		for _, r := range *d.IncludedResultResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResources != nil {
		for _, r := range *d.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSpecimenResources != nil {
		for _, r := range *d.IncludedSpecimenResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
