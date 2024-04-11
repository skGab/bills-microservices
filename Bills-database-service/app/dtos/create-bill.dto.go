package DTOs

type CreateBillDTO struct {
	Name     string `json:"name" binding:"required"`
	Value    string `json:"value" binding:"required"`
	Due_date string `json:"due_date" binding:"required"`
}
