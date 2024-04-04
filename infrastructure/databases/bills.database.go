package databases

import (
	"errors"
	"skGab/Bills-management-service/domain/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BillsDatabase struct {
	Gorm *gorm.DB
}

// DATABASE CONNECTION
func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

// CREATE BILL
func (b *BillsDatabase) Create(billEntity *entities.BillEntity) error {
	// STORE ON DATABASE
	response := b.Gorm.Create(&billEntity)

	if response.Error != nil {
		return response.Error
	}

	if response.RowsAffected == 0 {
		return errors.New("algo aconteceu durante a criação da conta")
	}

	// RETURN NIL IF ANY ERRORS
	return nil
}
