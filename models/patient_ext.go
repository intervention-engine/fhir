package models

import "gopkg.in/mgo.v2/bson"

func (p *Patient) ToFact() Fact {
	f := Fact{}
	f.BirthDate = p.BirthDate
	f.PatientID = p.Id
	f.TargetID = p.Id
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}
