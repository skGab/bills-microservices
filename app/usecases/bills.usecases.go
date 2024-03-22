package usecases

import (
	"fmt"
	DTOs "skGab/Bills-management-service/app/dtos"
	"skGab/Bills-management-service/domain/entities"
	repository "skGab/Bills-management-service/domain/repositories"

	"github.com/google/uuid"
)

type BillsUsecases struct {
	Repository repository.BillsRepository
}

func (b *BillsUsecases) CreateBill(formData *DTOs.CreateBillDTO) *error {
	// BUILD ENTITY
	billEntity := &entities.BillEntity{
		ID:    uuid.New(),
		Name:  formData.Name,
		Value: formData.Value,
		Date:  formData.Date,
	}

	// PASS ENTITY DATA TO REPOSITORY
	err := b.Repository.Create(billEntity)

	// return response
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
