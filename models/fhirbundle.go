package models

import "time"

type FHIRBundle struct {
	Type         string
	Title        string
	Id           string
	Updated      time.Time
	TotalResults int
	Entries      []Patient
	Category     Category
}

type Category struct {
	Term   string
	Label  string
	Scheme string
}
