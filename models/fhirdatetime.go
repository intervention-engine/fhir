package models

import "time"

type Precision string

const (
	Date      = "date"
	Timestamp = "timestamp"
)

type FHIRDateTime struct {
	Time      time.Time
	Precision Precision
}

func (f *FHIRDateTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) <= 12 {
		f.Precision = Precision("date")
		f.Time, err = time.Parse("\"2006-01-02\"", string(data))
	} else {
		f.Precision = Precision("timestamp")
		f.Time = time.Time{}
		f.Time.UnmarshalJSON(data)
	}
	return err
}
