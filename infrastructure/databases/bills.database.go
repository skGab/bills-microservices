package databases

import (
	"skGab/Bills-management-service/domain/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BillsDatabase struct {
	Gorm *gorm.DB
}

func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func (*BillsDatabase) Create(billEntity *entities.BillEntity) *error {
	return nil
}
