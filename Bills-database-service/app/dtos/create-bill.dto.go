package DTOs

type CreateBillDTO struct {
	Name     string `json:"name" validate:"required"`
	Value    string `json:"value" validate:"required"`
	Due_date string `json:"due_date" validate:"required"`
}
