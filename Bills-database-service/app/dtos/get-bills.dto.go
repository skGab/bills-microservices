package DTOs

import "time"

type GetBillsDTO struct {
	Name  string
	Value float64
	Date  time.Time
}
