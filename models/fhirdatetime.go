package models

import (
	"encoding/json"
	"time"
)

type Precision string

const (
	Date      = "date"
	YearMonth = "year-month"
	Timestamp = "timestamp"
	Time      = "time"
)

type FHIRDateTime struct {
	Time      time.Time
	Precision Precision
}

func (f *FHIRDateTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) <= 12 {
		f.Precision = Precision("date")
		f.Time, err = time.ParseInLocation("\"2006-01-02\"", string(data), time.Local)
		if err != nil {
			f.Precision = Precision("year-month")
			f.Time, err = time.ParseInLocation("\"2006-01\"", string(data), time.Local)
		}
		if err != nil {
			// TODO: should move time into a separate type
			f.Precision = Precision("time")
			f.Time, err = time.ParseInLocation("\"15:04:05\"", string(data), time.Local)
		}
		if err != nil {
			f.Precision = ""
		}

	} else {
		f.Precision = Precision("timestamp")
		f.Time = time.Time{}
		f.Time.UnmarshalJSON(data)
	}
	return err
}

func (f FHIRDateTime) MarshalJSON() ([]byte, error) {
	if f.Precision == Timestamp {
		return json.Marshal(f.Time.Format(time.RFC3339))
	} else if f.Precision == YearMonth {
		return json.Marshal(f.Time.Format("2006-01"))
	} else if f.Precision == Time {
		return json.Marshal(f.Time.Format("15:04:05"))
	} else {
		return json.Marshal(f.Time.Format("2006-01-02"))
	}
}
