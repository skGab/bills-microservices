package usecases

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	DTOs "github.com/skGab/Bills-microservices/Bills-database-service/app/dtos"
	"github.com/skGab/Bills-microservices/Bills-database-service/domain/entities"
	"github.com/skGab/Bills-microservices/Bills-database-service/domain/repositories"
)

type BillsUsecases struct {
	Repository repositories.BillsRepository
}

// GET ALL BILLS
func (b *BillsUsecases) GetAllBills(clientID string) ([]DTOs.GetBillsDTO, error) {

	// CHECK FOR PROBLEMNS ON THE CLIENT ID
	if clientID == "nil" {
		return nil, errors.New("ID do cliente não encontrado no parametro da rota")
	}

	// CALL THE DATABASE TO DO THE OPERATION
	bills, err := b.Repository.GetAll(clientID)

	if err != nil {
		return nil, err
	}

	// MAP TO DTO
	var billsDTO []DTOs.GetBillsDTO

	for _, value := range bills {

		bill := DTOs.GetBillsDTO{
			ID:       value.ID,
			Name:     value.Name,
			Value:    value.Value,
			Due_date: value.Due_date.Format(time.DateOnly),
		}

		billsDTO = append(billsDTO, bill)
	}

	return billsDTO, nil
}

// CREATE BILL
func (b *BillsUsecases) CreateBill(formData *DTOs.CreateBillDTO) error {

	if formData == nil {
		return errors.New("dados não encontrados no corpo da requisição")
	}

	// CONVERT STRING DUE DATE TO TIME.TIME
	due_date, error := time.Parse(time.DateOnly, formData.Due_date)

	if error != nil {
		return error
	}

	// BUILD ENTITY
	billEntity := &entities.BillEntity{
		ID:       uuid.New().String(),
		Name:     formData.Name,
		Value:    formData.Value,
		Due_date: due_date,
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

// UPDATE BILL
func (b *BillsUsecases) UpdateBill(billID string, data interface{}) (bool, error) {

	if data == nil || billID == "" {
		return false, errors.New("ID ou dados não encontrados para atualização")
	}

	err := b.Repository.Update(billID, data)

	if err != nil {
		return false, err
	}

	return true, nil
}

// DELETE BILL
func (b *BillsUsecases) DeleteBill(billID string) (bool, error) {

	if billID == "" {
		return false, errors.New("ID da conta não encontrado no parametro")
	}

	err := b.Repository.Delete(billID)

	if err != nil {
		return false, err
	}

	return true, nil
}

// DELETE ALL BILLS
func (b *BillsUsecases) DeleteAllBills(billsIDs []string) (bool, error) {

	if len(billsIDs) == 0 {
		return false, errors.New("IDs não foram encontrados no corpo da requisição")
	}

	err := b.Repository.DeleteAll(billsIDs)

	if err != nil {
		return false, err
	}

	return true, nil
}
