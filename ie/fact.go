package ie

import (
	"gitlab.mitre.org/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

type Fact struct {
	Id                    string                   `json:"-" bson:"_id"`
	TargetID              string                   `json:"targetid" bson:"targetid"`
	StartDate             models.FHIRDateTime      `json:"startdate" bson:"startdate"`
	EndDate               models.FHIRDateTime      `json:"enddate" bson:"enddate"`
	BirthDate             models.FHIRDateTime      `json:"birthdate" bson:"birthdate"`
	Codes                 []models.CodeableConcept `json:"codes" bson:"codes"`
	ResultQuantity        models.Quantity          `json:"resultquantity" bson:"resultquantity"`
	ResultCodeableConcept models.CodeableConcept   `json:"resultcodeableconcept" bson:"resultcodeableconcept"`
	PatientID             string                   `json:"patientid" bson:"patientid"`
	Type                  string                   `json:"type" bson:"type"`
	Gender                string                   `json:"gender" bson:"gender"`
}

func FactFromPatient(p *models.Patient) Fact {
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

func FactFromCondition(c *models.Condition) Fact {
	f := Fact{}
	f.Type = "Condition"
	f.StartDate = c.OnsetDate
	f.EndDate = c.AbatementDate
	f.Codes = []models.CodeableConcept{c.Code}
	f.PatientID = c.Subject.ReferencedID
	f.TargetID = c.Id
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}

func FactFromEncounter(e *models.Encounter) Fact {
	f := Fact{}
	f.Type = "Encounter"
	f.StartDate = e.Period.Start
	f.EndDate = e.Period.End
	f.Codes = e.Type
	f.PatientID = e.Subject.ReferencedID
	f.TargetID = e.Id
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}

func FactFromObservation(o *models.Observation) Fact {
	f := Fact{}
	f.Type = "Observation"
	f.StartDate = o.AppliesPeriod.Start
	f.EndDate = o.AppliesPeriod.End
	f.ResultQuantity = o.ValueQuantity
	f.ResultCodeableConcept = o.ValueCodeableConcept
	f.Codes = []models.CodeableConcept{o.Name}
	f.PatientID = o.Subject.ReferencedID
	f.TargetID = o.Id
	i := bson.NewObjectId()
	f.Id = i.Hex()
	return f
}
