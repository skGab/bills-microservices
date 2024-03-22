package DTOs

import "time"

type CreateBillDTO struct {
	Name  string    `json:"name"`
	Value float64   `json:"value"`
	Date  time.Time `json:"date"`
}
