package models

import "gopkg.in/mgo.v2/bson"

func (p *Patient) ToFact() Fact {
	f := Fact{}
	f.Type = "Patient"
	f.BirthDate = p.BirthDate
	f.PatientID = p.Id
	f.TargetID = p.Id
	f.Gender = p.Gender.Coding[0].Code
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}
