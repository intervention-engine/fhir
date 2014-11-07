package models

import "gopkg.in/mgo.v2/bson"

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
