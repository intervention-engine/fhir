package models

import "gopkg.in/mgo.v2/bson"

type Fact struct {
	Id                    string            `json:"-" bson:"_id"`
	TargetID              string            `json:"targetid" bson:"targetid"`
	StartDate             FHIRDateTime      `json:"startdate" bson:"startdate"`
	EndDate               FHIRDateTime      `json:"enddate" bson:"enddate"`
	BirthDate             FHIRDateTime      `json:"birthdate" bson:"birthdate"`
	Codes                 []CodeableConcept `json:"codes" bson:"codes"`
	ResultQuantity        Quantity          `json:"resultquantity" bson:"resultquantity"`
	ResultCodeableConcept CodeableConcept    `json:"resultcodeableconcept" bson:"resultcodeableconcept"`
	PatientID             string            `json:"patientid" bson:"patientid"`
}

func (p *Patient) ToFact() Fact {
	f := Fact{}
	f.BirthDate = p.BirthDate
	f.PatientID = p.Id
	f.TargetID = p.Id
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}

func (c *Condition) ToFact() Fact {
	f := Fact{}
	f.StartDate = c.OnsetDate
	f.EndDate = c.AbatementDate
	f.Codes = []CodeableConcept{c.Code}
	f.PatientID = c.Subject.ReferencedID
	f.TargetID = c.Id
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}

func (e *Encounter) ToFact() Fact {
	f := Fact{}
	f.StartDate = e.Period.Start
	f.EndDate = e.Period.End
	f.Codes = e.Type
	f.PatientID = e.Subject.ReferencedID
	f.TargetID = e.Id
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}

func (o *Observation) ToFact() Fact {
	f := Fact{}
	f.StartDate = o.AppliesPeriod.Start
	f.EndDate = o.AppliesPeriod.End
	f.ResultQuantity = o.ValueQuantity
	f.ResultCodeableConcept = o.ValueCodeableConcept
	f.Codes = []CodeableConcept{o.Name}
	f.PatientID = o.Subject.ReferencedID
	f.TargetID = o.Id
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}
