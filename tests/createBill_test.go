package databases

import (
	"skGab/Bills-management-service/domain/entities"
	"skGab/Bills-management-service/infrastructure/databases"
	"testing"
	"time"
)

func CreateTest(test *testing.T) {
	// GET DATA
	billEntity := &entities.BillEntity{
		Name:  "CEMIG",
		Value: 3.000,
		Date:  time.Now(),
	}

	// STORE ENTITY ON DATABASE
	response := databases.BillsDatabase
}

// SHOULD TEST  WITH CORRECT DATA
// SHOULD TEST IF DONT PASS A FORMDATA
// SHOULLD TEST IF PASS DIFERENT FIELDS
// SHOULLD TEST IF PASS DIFERENT TYPES
