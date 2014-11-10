package models

import "gopkg.in/mgo.v2/bson"

func (q *Query) ToPipeline() []bson.M {
	pipeline := []bson.M{{"$group": bson.M{"_id": "$targetid", "gender": bson.M{"$max": "$gender"}, "birthdate": bson.M{"$max": "$birthdate"}, "entries": bson.M{"$addToSet": bson.M{"startdate": "$startdate", "enddate": "$enddate", "codes": "$codes", "type": "$type"}}}}}
	for _, extension := range q.Parameter {
		switch extension.Url {
		case "http://interventionengine.org/patientgender":
			pipeline = append(pipeline, bson.M{"$match": bson.M{"gender": extension.ValueString}})
		}
	}

	pipeline = append(pipeline, bson.M{"$group": bson.M{"_id": nil, "total": bson.M{"$sum": 1}}})
	return pipeline
}
