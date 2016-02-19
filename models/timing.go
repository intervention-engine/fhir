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

type Timing struct {
	Event  []FHIRDateTime         `bson:"event,omitempty" json:"event,omitempty"`
	Repeat *TimingRepeatComponent `bson:"repeat,omitempty" json:"repeat,omitempty"`
	Code   *CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
}

type TimingRepeatComponent struct {
	BackboneElement `bson:",inline"`
	BoundsDuration  *Quantity `bson:"boundsDuration,omitempty" json:"boundsDuration,omitempty"`
	BoundsRange     *Range    `bson:"boundsRange,omitempty" json:"boundsRange,omitempty"`
	BoundsPeriod    *Period   `bson:"boundsPeriod,omitempty" json:"boundsPeriod,omitempty"`
	Count           *int32    `bson:"count,omitempty" json:"count,omitempty"`
	Duration        *float64  `bson:"duration,omitempty" json:"duration,omitempty"`
	DurationMax     *float64  `bson:"durationMax,omitempty" json:"durationMax,omitempty"`
	DurationUnits   string    `bson:"durationUnits,omitempty" json:"durationUnits,omitempty"`
	Frequency       *int32    `bson:"frequency,omitempty" json:"frequency,omitempty"`
	FrequencyMax    *int32    `bson:"frequencyMax,omitempty" json:"frequencyMax,omitempty"`
	Period          *float64  `bson:"period,omitempty" json:"period,omitempty"`
	PeriodMax       *float64  `bson:"periodMax,omitempty" json:"periodMax,omitempty"`
	PeriodUnits     string    `bson:"periodUnits,omitempty" json:"periodUnits,omitempty"`
	When            string    `bson:"when,omitempty" json:"when,omitempty"`
}
