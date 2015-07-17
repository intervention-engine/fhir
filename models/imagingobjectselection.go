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

import "encoding/json"

type ImagingObjectSelection struct {
	Id            string                                 `json:"id" bson:"_id"`
	Uid           string                                 `bson:"uid,omitempty" json:"uid,omitempty"`
	Patient       *Reference                             `bson:"patient,omitempty" json:"patient,omitempty"`
	Title         *CodeableConcept                       `bson:"title,omitempty" json:"title,omitempty"`
	Description   string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Author        *Reference                             `bson:"author,omitempty" json:"author,omitempty"`
	AuthoringTime *FHIRDateTime                          `bson:"authoringTime,omitempty" json:"authoringTime,omitempty"`
	Study         []ImagingObjectSelectionStudyComponent `bson:"study,omitempty" json:"study,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ImagingObjectSelection) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		ImagingObjectSelection
	}{
		ResourceType:           "ImagingObjectSelection",
		ImagingObjectSelection: *resource,
	}
	return json.Marshal(x)
}

type ImagingObjectSelectionStudyComponent struct {
	Uid    string                                  `bson:"uid,omitempty" json:"uid,omitempty"`
	Url    string                                  `bson:"url,omitempty" json:"url,omitempty"`
	Series []ImagingObjectSelectionSeriesComponent `bson:"series,omitempty" json:"series,omitempty"`
}

type ImagingObjectSelectionSeriesComponent struct {
	Uid      string                                    `bson:"uid,omitempty" json:"uid,omitempty"`
	Url      string                                    `bson:"url,omitempty" json:"url,omitempty"`
	Instance []ImagingObjectSelectionInstanceComponent `bson:"instance,omitempty" json:"instance,omitempty"`
}

type ImagingObjectSelectionInstanceComponent struct {
	SopClass string                                  `bson:"sopClass,omitempty" json:"sopClass,omitempty"`
	Uid      string                                  `bson:"uid,omitempty" json:"uid,omitempty"`
	Url      string                                  `bson:"url,omitempty" json:"url,omitempty"`
	Frames   []ImagingObjectSelectionFramesComponent `bson:"frames,omitempty" json:"frames,omitempty"`
}

type ImagingObjectSelectionFramesComponent struct {
	FrameNumbers []uint32 `bson:"frameNumbers,omitempty" json:"frameNumbers,omitempty"`
	Url          string   `bson:"url,omitempty" json:"url,omitempty"`
}
