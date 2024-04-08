package databases

import (
	"errors"
	"fmt"

	"github.com/skGab/Bills-microservices/Bills-database-service/domain/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BillsDatabase struct {
	DB *gorm.DB
}

// CREATE THE CONNECTION WITH THE SELECTED DATABASE
func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

// GET ALL BILLS
func (bills *BillsDatabase) GetAll(clientID int) ([]entities.BillEntity, error) {
	var billsEntity []entities.BillEntity

	response := bills.DB.Where("id = ?", clientID).Find(&billsEntity)

	if response.Error != nil {
		fmt.Println(response.Error)
		return nil, response.Error
	}

	if response.RowsAffected == 0 {
		return nil, errors.New("algo aconteceu durante a busca de contas")
	}

	return billsEntity, nil
}

// CREATE BILL
func (bills *BillsDatabase) Create(billEntity *entities.BillEntity) error {
	response := bills.DB.Create(billEntity)

	if response.Error != nil {
		fmt.Println(response.Error)

		return response.Error
	}

	if response.RowsAffected == 0 {
		return errors.New("algo aconteceu durante a criação da conta")
	}

	return nil
}

// UPDATE BILL
func (bill *BillsDatabase) Update(billID int, data interface{}) error {
	response := bill.DB.Model(&entities.BillEntity{}).Where("id = ?", billID).Updates(data)

	if response.Error != nil {
		return response.Error
	}

	if response.RowsAffected == 0 {
		return errors.New("algo inesperado aconteceu durante a atualização da conta")
	}

	return nil
}

// DELETE BILL
func (bill *BillsDatabase) Delete(billID int) error {

	response := bill.DB.Delete(&entities.BillEntity{}, billID)

	if response.Error != nil {
		return response.Error
	}

	if response.RowsAffected == 0 {
		return errors.New("algo inesperado aconteceu ao deletar a tarefa")
	}

	return nil
}

// DELETE ALL BILLS
func (bill *BillsDatabase) DeleteAll(billsIDs []int) error {
	response := bill.DB.Delete(&entities.BillEntity{}, billsIDs)

	if response.Error != nil {
		return response.Error
	}

	if response.RowsAffected == 0 {
		return errors.New("algo inesperado aconteceu ao deletar a tarefa")
	}

	// RETURN NIL IF ANY ERRORS
	return nil
}
