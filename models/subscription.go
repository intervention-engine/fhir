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

type Subscription struct {
	DomainResource `bson:",inline"`
	Criteria       string                        `bson:"criteria,omitempty" json:"criteria,omitempty"`
	Contact        []ContactPoint                `bson:"contact,omitempty" json:"contact,omitempty"`
	Reason         string                        `bson:"reason,omitempty" json:"reason,omitempty"`
	Status         string                        `bson:"status,omitempty" json:"status,omitempty"`
	Error          string                        `bson:"error,omitempty" json:"error,omitempty"`
	Channel        *SubscriptionChannelComponent `bson:"channel,omitempty" json:"channel,omitempty"`
	End            *FHIRDateTime                 `bson:"end,omitempty" json:"end,omitempty"`
	Tag            []Coding                      `bson:"tag,omitempty" json:"tag,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Subscription) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Subscription
	}{
		ResourceType: "Subscription",
		Subscription: *resource,
	}
	return json.Marshal(x)
}

// The "subscription" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type subscription Subscription

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Subscription) UnmarshalJSON(data []byte) (err error) {
	x2 := subscription{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Subscription(x2)
	}
	return
}

type SubscriptionChannelComponent struct {
	Type     string `bson:"type,omitempty" json:"type,omitempty"`
	Endpoint string `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	Payload  string `bson:"payload,omitempty" json:"payload,omitempty"`
	Header   string `bson:"header,omitempty" json:"header,omitempty"`
}

type SubscriptionPlus struct {
	Subscription             `bson:",inline"`
	SubscriptionPlusIncludes `bson:",inline"`
}

type SubscriptionPlusIncludes struct {
}

func (s *SubscriptionPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}
