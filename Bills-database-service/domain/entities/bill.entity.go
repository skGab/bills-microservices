package entities

import (
	"time"
)

type BillEntity struct {
	ID    string
	Name  string
	Value float64
	Date  time.Time
}
