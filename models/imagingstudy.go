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

type ImagingStudy struct {
	DomainResource    `bson:",inline"`
	Started           *FHIRDateTime                 `bson:"started,omitempty" json:"started,omitempty"`
	Patient           *Reference                    `bson:"patient,omitempty" json:"patient,omitempty"`
	Uid               string                        `bson:"uid,omitempty" json:"uid,omitempty"`
	Accession         *Identifier                   `bson:"accession,omitempty" json:"accession,omitempty"`
	Identifier        []Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Order             []Reference                   `bson:"order,omitempty" json:"order,omitempty"`
	ModalityList      []Coding                      `bson:"modalityList,omitempty" json:"modalityList,omitempty"`
	Referrer          *Reference                    `bson:"referrer,omitempty" json:"referrer,omitempty"`
	Availability      string                        `bson:"availability,omitempty" json:"availability,omitempty"`
	Url               string                        `bson:"url,omitempty" json:"url,omitempty"`
	NumberOfSeries    *uint32                       `bson:"numberOfSeries,omitempty" json:"numberOfSeries,omitempty"`
	NumberOfInstances *uint32                       `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	Procedure         []Reference                   `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Interpreter       *Reference                    `bson:"interpreter,omitempty" json:"interpreter,omitempty"`
	Description       string                        `bson:"description,omitempty" json:"description,omitempty"`
	Series            []ImagingStudySeriesComponent `bson:"series,omitempty" json:"series,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ImagingStudy) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		ImagingStudy
	}{
		ResourceType: "ImagingStudy",
		ImagingStudy: *resource,
	}
	return json.Marshal(x)
}

// The "imagingStudy" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type imagingStudy ImagingStudy

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ImagingStudy) UnmarshalJSON(data []byte) (err error) {
	x2 := imagingStudy{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ImagingStudy(x2)
	}
	return
}

type ImagingStudySeriesComponent struct {
	Number            *uint32                               `bson:"number,omitempty" json:"number,omitempty"`
	Modality          *Coding                               `bson:"modality,omitempty" json:"modality,omitempty"`
	Uid               string                                `bson:"uid,omitempty" json:"uid,omitempty"`
	Description       string                                `bson:"description,omitempty" json:"description,omitempty"`
	NumberOfInstances *uint32                               `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	Availability      string                                `bson:"availability,omitempty" json:"availability,omitempty"`
	Url               string                                `bson:"url,omitempty" json:"url,omitempty"`
	BodySite          *Coding                               `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Laterality        *Coding                               `bson:"laterality,omitempty" json:"laterality,omitempty"`
	Started           *FHIRDateTime                         `bson:"started,omitempty" json:"started,omitempty"`
	Instance          []ImagingStudySeriesInstanceComponent `bson:"instance,omitempty" json:"instance,omitempty"`
}

type ImagingStudySeriesInstanceComponent struct {
	Number   *uint32      `bson:"number,omitempty" json:"number,omitempty"`
	Uid      string       `bson:"uid,omitempty" json:"uid,omitempty"`
	SopClass string       `bson:"sopClass,omitempty" json:"sopClass,omitempty"`
	Type     string       `bson:"type,omitempty" json:"type,omitempty"`
	Title    string       `bson:"title,omitempty" json:"title,omitempty"`
	Content  []Attachment `bson:"content,omitempty" json:"content,omitempty"`
}

type ImagingStudyPlus struct {
	ImagingStudy             `bson:",inline"`
	ImagingStudyPlusIncludes `bson:",inline"`
}

type ImagingStudyPlusIncludes struct {
	IncludedPatientResources *[]Patient         `bson:"_includedPatientResources,omitempty"`
	IncludedOrderResources   *[]DiagnosticOrder `bson:"_includedOrderResources,omitempty"`
}

func (i *ImagingStudyPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if i.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedPatientResources))
	} else if len(*i.IncludedPatientResources) == 1 {
		patient = &(*i.IncludedPatientResources)[0]
	}
	return
}

func (i *ImagingStudyPlusIncludes) GetIncludedOrderResources() (diagnosticOrders []DiagnosticOrder, err error) {
	if i.IncludedOrderResources == nil {
		err = errors.New("Included diagnosticOrders not requested")
	} else {
		diagnosticOrders = *i.IncludedOrderResources
	}
	return
}

func (i *ImagingStudyPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedPatientResources != nil {
		for _, r := range *i.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedOrderResources != nil {
		for _, r := range *i.IncludedOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
