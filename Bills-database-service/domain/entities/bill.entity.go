package entities

import (
	"time"
)

type BillEntity struct {
	ID       string
	Name     string
	Value    string
	Due_date time.Time
}
