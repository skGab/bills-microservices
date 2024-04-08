package entities

import (
	"time"

	"github.com/google/uuid"
)

type BillEntity struct {
	ID    uuid.UUID
	Name  string
	Value float64
	Date  time.Time
}
