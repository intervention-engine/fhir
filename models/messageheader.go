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

type MessageHeader struct {
	Id          string                         `json:"-" bson:"_id"`
	Identifier  string                         `bson:"identifier"`
	Timestamp   FHIRDateTime                   `bson:"timestamp"`
	Event       Coding                         `bson:"event"`
	Response    MessageHeaderResponseComponent `bson:"response"`
	Source      MessageSourceComponent         `bson:"source"`
	Destination []MessageDestinationComponent  `bson:"destination"`
	Enterer     Reference                      `bson:"enterer"`
	Author      Reference                      `bson:"author"`
	Receiver    Reference                      `bson:"receiver"`
	Responsible Reference                      `bson:"responsible"`
	Reason      CodeableConcept                `bson:"reason"`
	Data        []Reference                    `bson:"data"`
}

// This is an ugly hack to deal with embedded structures in the spec response
type MessageHeaderResponseComponent struct {
	Identifier string    `bson:"identifier"`
	Code       string    `bson:"code"`
	Details    Reference `bson:"details"`
}

// This is an ugly hack to deal with embedded structures in the spec source
type MessageSourceComponent struct {
	Name     string       `bson:"name"`
	Software string       `bson:"software"`
	Version  string       `bson:"version"`
	Contact  ContactPoint `bson:"contact"`
	Endpoint string       `bson:"endpoint"`
}

// This is an ugly hack to deal with embedded structures in the spec destination
type MessageDestinationComponent struct {
	Name     string    `bson:"name"`
	Target   Reference `bson:"target"`
	Endpoint string    `bson:"endpoint"`
}
