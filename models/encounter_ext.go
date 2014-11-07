package models

import "gopkg.in/mgo.v2/bson"

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
