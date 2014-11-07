package models

import "gopkg.in/mgo.v2/bson"

func (o *Observation) ToFact() Fact {
	f := Fact{}
  f.Type = "Observation"
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
