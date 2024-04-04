package DTOs

import "time"

type AllBillsDTO struct {
	Name  string
	Value float64
	Date  time.Time
}
