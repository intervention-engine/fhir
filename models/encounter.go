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

type Encounter struct {
	DomainResource   `bson:",inline"`
	Identifier       []Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status           string                             `bson:"status,omitempty" json:"status,omitempty"`
	StatusHistory    []EncounterStatusHistoryComponent  `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Class            string                             `bson:"class,omitempty" json:"class,omitempty"`
	Type             []CodeableConcept                  `bson:"type,omitempty" json:"type,omitempty"`
	Priority         *CodeableConcept                   `bson:"priority,omitempty" json:"priority,omitempty"`
	Patient          *Reference                         `bson:"patient,omitempty" json:"patient,omitempty"`
	EpisodeOfCare    []Reference                        `bson:"episodeOfCare,omitempty" json:"episodeOfCare,omitempty"`
	IncomingReferral []Reference                        `bson:"incomingReferral,omitempty" json:"incomingReferral,omitempty"`
	Participant      []EncounterParticipantComponent    `bson:"participant,omitempty" json:"participant,omitempty"`
	Appointment      *Reference                         `bson:"appointment,omitempty" json:"appointment,omitempty"`
	Period           *Period                            `bson:"period,omitempty" json:"period,omitempty"`
	Length           *Quantity                          `bson:"length,omitempty" json:"length,omitempty"`
	Reason           []CodeableConcept                  `bson:"reason,omitempty" json:"reason,omitempty"`
	Indication       []Reference                        `bson:"indication,omitempty" json:"indication,omitempty"`
	Hospitalization  *EncounterHospitalizationComponent `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Location         []EncounterLocationComponent       `bson:"location,omitempty" json:"location,omitempty"`
	ServiceProvider  *Reference                         `bson:"serviceProvider,omitempty" json:"serviceProvider,omitempty"`
	PartOf           *Reference                         `bson:"partOf,omitempty" json:"partOf,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Encounter) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Encounter"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Encounter), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Encounter) GetBSON() (interface{}, error) {
	x.ResourceType = "Encounter"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "encounter" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type encounter Encounter

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Encounter) UnmarshalJSON(data []byte) (err error) {
	x2 := encounter{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Encounter(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Encounter) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Encounter"
	} else if x.ResourceType != "Encounter" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Encounter, instead received %s", x.ResourceType))
	}
	return nil
}

type EncounterStatusHistoryComponent struct {
	Status string  `bson:"status,omitempty" json:"status,omitempty"`
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type EncounterParticipantComponent struct {
	Type       []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Period     *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Individual *Reference        `bson:"individual,omitempty" json:"individual,omitempty"`
}

type EncounterHospitalizationComponent struct {
	PreAdmissionIdentifier *Identifier       `bson:"preAdmissionIdentifier,omitempty" json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference        `bson:"origin,omitempty" json:"origin,omitempty"`
	AdmitSource            *CodeableConcept  `bson:"admitSource,omitempty" json:"admitSource,omitempty"`
	AdmittingDiagnosis     []Reference       `bson:"admittingDiagnosis,omitempty" json:"admittingDiagnosis,omitempty"`
	ReAdmission            *CodeableConcept  `bson:"reAdmission,omitempty" json:"reAdmission,omitempty"`
	DietPreference         []CodeableConcept `bson:"dietPreference,omitempty" json:"dietPreference,omitempty"`
	SpecialCourtesy        []CodeableConcept `bson:"specialCourtesy,omitempty" json:"specialCourtesy,omitempty"`
	SpecialArrangement     []CodeableConcept `bson:"specialArrangement,omitempty" json:"specialArrangement,omitempty"`
	Destination            *Reference        `bson:"destination,omitempty" json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept  `bson:"dischargeDisposition,omitempty" json:"dischargeDisposition,omitempty"`
	DischargeDiagnosis     []Reference       `bson:"dischargeDiagnosis,omitempty" json:"dischargeDiagnosis,omitempty"`
}

type EncounterLocationComponent struct {
	Location *Reference `bson:"location,omitempty" json:"location,omitempty"`
	Status   string     `bson:"status,omitempty" json:"status,omitempty"`
	Period   *Period    `bson:"period,omitempty" json:"period,omitempty"`
}

type EncounterPlus struct {
	Encounter             `bson:",inline"`
	EncounterPlusIncludes `bson:",inline"`
}

type EncounterPlusIncludes struct {
	IncludedEpisodeofcareResources            *[]EpisodeOfCare   `bson:"_includedEpisodeofcareResources,omitempty"`
	IncludedIncomingreferralResources         *[]ReferralRequest `bson:"_includedIncomingreferralResources,omitempty"`
	IncludedPractitionerResources             *[]Practitioner    `bson:"_includedPractitionerResources,omitempty"`
	IncludedAppointmentResources              *[]Appointment     `bson:"_includedAppointmentResources,omitempty"`
	IncludedPartofResources                   *[]Encounter       `bson:"_includedPartofResources,omitempty"`
	IncludedProcedureResources                *[]Procedure       `bson:"_includedProcedureResources,omitempty"`
	IncludedParticipantPractitionerResources  *[]Practitioner    `bson:"_includedParticipantPractitionerResources,omitempty"`
	IncludedParticipantRelatedPersonResources *[]RelatedPerson   `bson:"_includedParticipantRelatedPersonResources,omitempty"`
	IncludedConditionResources                *[]Condition       `bson:"_includedConditionResources,omitempty"`
	IncludedPatientResources                  *[]Patient         `bson:"_includedPatientResources,omitempty"`
	IncludedLocationResources                 *[]Location        `bson:"_includedLocationResources,omitempty"`
	IncludedIndicationConditionResources      *[]Condition       `bson:"_includedIndicationConditionResources,omitempty"`
	IncludedIndicationProcedureResources      *[]Procedure       `bson:"_includedIndicationProcedureResources,omitempty"`
}

func (e *EncounterPlusIncludes) GetIncludedEpisodeofcareResources() (episodeOfCares []EpisodeOfCare, err error) {
	if e.IncludedEpisodeofcareResources == nil {
		err = errors.New("Included episodeOfCares not requested")
	} else {
		episodeOfCares = *e.IncludedEpisodeofcareResources
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedIncomingreferralResources() (referralRequests []ReferralRequest, err error) {
	if e.IncludedIncomingreferralResources == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *e.IncludedIncomingreferralResources
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedPractitionerResource() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResources))
	} else if len(*e.IncludedPractitionerResources) == 1 {
		practitioner = &(*e.IncludedPractitionerResources)[0]
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedAppointmentResource() (appointment *Appointment, err error) {
	if e.IncludedAppointmentResources == nil {
		err = errors.New("Included appointments not requested")
	} else if len(*e.IncludedAppointmentResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 appointment, but found %d", len(*e.IncludedAppointmentResources))
	} else if len(*e.IncludedAppointmentResources) == 1 {
		appointment = &(*e.IncludedAppointmentResources)[0]
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedPartofResource() (encounter *Encounter, err error) {
	if e.IncludedPartofResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*e.IncludedPartofResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*e.IncludedPartofResources))
	} else if len(*e.IncludedPartofResources) == 1 {
		encounter = &(*e.IncludedPartofResources)[0]
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedProcedureResources() (procedures []Procedure, err error) {
	if e.IncludedProcedureResources == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *e.IncludedProcedureResources
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedParticipantPractitionerResource() (practitioner *Practitioner, err error) {
	if e.IncludedParticipantPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedParticipantPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedParticipantPractitionerResources))
	} else if len(*e.IncludedParticipantPractitionerResources) == 1 {
		practitioner = &(*e.IncludedParticipantPractitionerResources)[0]
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedParticipantRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if e.IncludedParticipantRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*e.IncludedParticipantRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*e.IncludedParticipantRelatedPersonResources))
	} else if len(*e.IncludedParticipantRelatedPersonResources) == 1 {
		relatedPerson = &(*e.IncludedParticipantRelatedPersonResources)[0]
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedConditionResources() (conditions []Condition, err error) {
	if e.IncludedConditionResources == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *e.IncludedConditionResources
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if e.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResources))
	} else if len(*e.IncludedPatientResources) == 1 {
		patient = &(*e.IncludedPatientResources)[0]
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedLocationResource() (location *Location, err error) {
	if e.IncludedLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*e.IncludedLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*e.IncludedLocationResources))
	} else if len(*e.IncludedLocationResources) == 1 {
		location = &(*e.IncludedLocationResources)[0]
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedIndicationConditionResources() (conditions []Condition, err error) {
	if e.IncludedIndicationConditionResources == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *e.IncludedIndicationConditionResources
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedIndicationProcedureResources() (procedures []Procedure, err error) {
	if e.IncludedIndicationProcedureResources == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *e.IncludedIndicationProcedureResources
	}
	return
}

func (e *EncounterPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedEpisodeofcareResources != nil {
		for _, r := range *e.IncludedEpisodeofcareResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedIncomingreferralResources != nil {
		for _, r := range *e.IncludedIncomingreferralResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPractitionerResources != nil {
		for _, r := range *e.IncludedPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedAppointmentResources != nil {
		for _, r := range *e.IncludedAppointmentResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPartofResources != nil {
		for _, r := range *e.IncludedPartofResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedProcedureResources != nil {
		for _, r := range *e.IncludedProcedureResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedParticipantPractitionerResources != nil {
		for _, r := range *e.IncludedParticipantPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedParticipantRelatedPersonResources != nil {
		for _, r := range *e.IncludedParticipantRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedConditionResources != nil {
		for _, r := range *e.IncludedConditionResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPatientResources != nil {
		for _, r := range *e.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedLocationResources != nil {
		for _, r := range *e.IncludedLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedIndicationConditionResources != nil {
		for _, r := range *e.IncludedIndicationConditionResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedIndicationProcedureResources != nil {
		for _, r := range *e.IncludedIndicationProcedureResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
