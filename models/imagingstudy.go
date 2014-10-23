// Copyright (c) 2011-2014, HL7, Inc & The MITRE Corporation
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

type ImagingStudy struct {
	Id                  string                        `json:"-" bson:"_id"`
	DateTime            FHIRDateTime                  `bson:"dateTime"`
	Subject             Reference                     `bson:"subject"`
	Uid                 string                        `bson:"uid"`
	AccessionNo         Identifier                    `bson:"accessionNo"`
	Identifier          []Identifier                  `bson:"identifier"`
	Order               []Reference                   `bson:"order"`
	Modality            []string                      `bson:"modality"`
	Referrer            Reference                     `bson:"referrer"`
	Availability        string                        `bson:"availability"`
	Url                 string                        `bson:"url"`
	NumberOfSeries      float64                       `bson:"numberOfSeries"`
	NumberOfInstances   float64                       `bson:"numberOfInstances"`
	ClinicalInformation string                        `bson:"clinicalInformation"`
	Procedure           []Coding                      `bson:"procedure"`
	Interpreter         Reference                     `bson:"interpreter"`
	Description         string                        `bson:"description"`
	Series              []ImagingStudySeriesComponent `bson:"series"`
}

// This is an ugly hack to deal with embedded structures in the spec instance
type ImagingStudySeriesInstanceComponent struct {
	Number     float64   `bson:"number"`
	Uid        string    `bson:"uid"`
	Sopclass   string    `bson:"sopclass"`
	Type       string    `bson:"type"`
	Title      string    `bson:"title"`
	Url        string    `bson:"url"`
	Attachment Reference `bson:"attachment"`
}

// This is an ugly hack to deal with embedded structures in the spec series
type ImagingStudySeriesComponent struct {
	Number            float64                               `bson:"number"`
	Modality          string                                `bson:"modality"`
	Uid               string                                `bson:"uid"`
	Description       string                                `bson:"description"`
	NumberOfInstances float64                               `bson:"numberOfInstances"`
	Availability      string                                `bson:"availability"`
	Url               string                                `bson:"url"`
	BodySite          Coding                                `bson:"bodySite"`
	DateTime          FHIRDateTime                          `bson:"dateTime"`
	Instance          []ImagingStudySeriesInstanceComponent `bson:"instance"`
}
