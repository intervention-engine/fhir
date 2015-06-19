package models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MedicationLookup func(id string) (Medication, error)

func MedicationFinder(database *mgo.Database, idString string) (Medication, error) {
	var id bson.ObjectId
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return Medication{}, errors.New("Invalid id")
	}

	c := database.C("medications")
	result := Medication{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return Medication{}, err
	}
	return result, nil
}

func BindMedicationLookup(database *mgo.Database) MedicationLookup {
	return func(id string) (Medication, error) {
		return MedicationFinder(database, id)
	}
}
