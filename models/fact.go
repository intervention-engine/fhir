package models

type Fact struct {
	Id                    string            `json:"-" bson:"_id"`
	TargetID              string            `json:"targetid" bson:"targetid"`
	StartDate             FHIRDateTime      `json:"startdate" bson:"startdate"`
	EndDate               FHIRDateTime      `json:"enddate" bson:"enddate"`
	BirthDate             FHIRDateTime      `json:"birthdate" bson:"birthdate"`
	Codes                 []CodeableConcept `json:"codes" bson:"codes"`
	ResultQuantity        Quantity          `json:"resultquantity" bson:"resultquantity"`
	ResultCodeableConcept CodeableConcept   `json:"resultcodeableconcept" bson:"resultcodeableconcept"`
	PatientID             string            `json:"patientid" bson:"patientid"`
	Type									string						`json:"type" bson:"type"`
}
