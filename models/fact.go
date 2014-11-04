package models

import "gopkg.in/mgo.v2/bson"

type Fact struct {
	Id                    string            `json:"-" bson:"_id"`
	StartDate             FHIRDateTime      `json:"startdate" bson:"startdate"`
	EndDate               FHIRDateTime      `json:"enddate" bson:"enddate"`
	BirthDate             FHIRDateTime      `json:"birthdate" bson:"birthdate"`
	Codes                 []CodeableConcept `json:"codes" bson:"codes"`
	ResultQuantity        Quantity          `json:"resultquantity" bson:"resultquantity"`
	ResultCodeableConcept CodableConcept    `json:"resultcodeableconcept" bson:"resultcodeableconcept"`
	PatientID             string            `json:"patientid" bson:"patientid"`
}

func (p *Patient) ToFact() (f Fact) {
	f := Fact{}
	f.BirthDate = p.BirthDate
	f.PatientID = p.Id
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}

func (c *Condition) ToFact() (f Fact) {
	f := Fact{}
	f.StartDate = c.OnsetDate
	f.EndDate = c.AbatementDate
	f.Codes = []CodeableConcept{c.Code}
	f.PatientID = c.Subject.ReferencedID
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}

func (e *Encounter) ToFact() (f Fact) {
	f := Fact{}
	f.StartDate = e.Period.Start
	f.EndDate = e.Period.End
	f.Codes = e.Type
	f.PatientID = e.Subject.ReferencedID
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}

func (o *Observation) ToFact() (f Fact) {
	f := Fact{}
	f.StartDate = o.AppliesPeriod.Start
	f.EndDate = o.AppliesPeriod.End
	f.ResultQuantity = o.ValueQuantity
	f.ResultCodeableConcept = o.ValueCodeableConcept
	f.Codes = []CodeableConcept{o.Name}
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}
