package models

import (
	"errors"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

// GetBSON translates the FHIR extension syntax to a syntax that is more suitable for storage and sorting in MongoDB.
//
// Extension {
//   Url: "http://example.org/fhir/extensions/foo",
//   ValueString: "bar",
// }
//
// becomes
//
// bson.M {
//   "@context": bson.M {
//     "foo": contextDefinition {
//       ID: "http://example.org/fhir/extensions/foo",
//       Type: "string",
//     },
//   },
//   "foo": "bar",
// }
func (e *Extension) GetBSON() (interface{}, error) {
	result := bson.M{"@context": bson.M{}}

	var i int
	if i = strings.LastIndex(e.Url, "/"); i < 0 || i == (len(e.Url)-1) {
		return nil, errors.New("Couldn't determine extension name for " + e.Url)
	}
	name := e.Url[i+1:]

	var fhirType string
	switch {
	case e.ValueBoolean != nil:
		fhirType = "boolean"
		result[name] = *e.ValueBoolean
	case e.ValueCodeableConcept != nil:
		fhirType = "CodeableConcept"
		result[name] = *e.ValueCodeableConcept
	case e.ValueDateTime != nil:
		fhirType = "dateTime"
		result[name] = *e.ValueDateTime
	case e.ValueInteger != nil:
		fhirType = "integer"
		result[name] = *e.ValueInteger
	case e.ValueRange != nil:
		fhirType = "Range"
		result[name] = *e.ValueRange
	default:
		fhirType = "string"
		result[name] = e.ValueString
	}

	result["@context"].(bson.M)[name] = contextDefinition{
		ID:   e.Url,
		Type: fhirType,
	}

	return result, nil
}

// SetBSON translates the stored extension syntax to the FHIR extension syntax.
//
// bson.M {
//   "@context": bson.M {
//     "foo": bson.M {
//       "@id": "http://example.org/fhir/extensions/foo",
//       "@type": "string",
//     },
//   },
//   "foo": "bar",
// }
//
// becomes
//
// Extension {
//   Url: "http://example.org/fhir/extensions/foo",
//   ValueString: "bar",
// }
func (e *Extension) SetBSON(raw bson.Raw) error {
	// Since we don't know the exact structure (property names), use a streaming approach with bson.RawD
	var rd bson.RawD
	err := raw.Unmarshal(&rd)
	if err != nil {
		return err
	}

	// Ensure there are only two sub-documents, then identify them
	if len(rd) != 2 {
		return errors.New("Couldn't properly unmarshal extension; unrecognized format in BSON")
	}
	var context map[string]contextDefinition
	var dataElement bson.RawDocElem
	for i := range rd {
		switch rd[i].Name {
		case "@context":
			rd[i].Value.Unmarshal(&context)
		default:
			dataElement = rd[i]
		}
	}
	if _, ok := context[dataElement.Name]; !ok {
		return errors.New("Couldn't properly unmarshal extension; key " + dataElement.Name + " not found in @context")
	}

	// Now set the URL and the right Value[x]
	e.Url = context[dataElement.Name].ID
	switch context[dataElement.Name].Type {
	case "boolean":
		e.ValueBoolean = new(bool)
		if err := dataElement.Value.Unmarshal(e.ValueBoolean); err != nil {
			return err
		}
	case "CodeableConcept":
		e.ValueCodeableConcept = new(CodeableConcept)
		if err := dataElement.Value.Unmarshal(e.ValueCodeableConcept); err != nil {
			return err
		}
	case "dateTime":
		e.ValueDateTime = new(FHIRDateTime)
		if err := dataElement.Value.Unmarshal(e.ValueDateTime); err != nil {
			return err
		}
	case "integer":
		e.ValueInteger = new(int32)
		if err := dataElement.Value.Unmarshal(e.ValueInteger); err != nil {
			return err
		}
	case "Range":
		e.ValueRange = new(Range)
		if err := dataElement.Value.Unmarshal(e.ValueRange); err != nil {
			return err
		}
	case "string":
		if err := dataElement.Value.Unmarshal(&e.ValueString); err != nil {
			return err
		}
	default:
		return errors.New("Couldn't determine extension value type from stored data")
	}

	return nil
}

type contextDefinition struct {
	ID   string `bson:"@id,omitempty"`
	Type string `bson:"@type,omitempty"`
}
