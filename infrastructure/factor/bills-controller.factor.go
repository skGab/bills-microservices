package factor

import (
	"skGab/Bills-management-service/app/controllers"
	"skGab/Bills-management-service/app/usecases"
	"skGab/Bills-management-service/infrastructure/databases"

	"gorm.io/gorm"
)

func BillsController(db *gorm.DB) *controllers.BillsController {
	// DATABASE INSTANCE
	billsDatabase := &databases.BillsDatabase{}

	// USECASES INSTANCE
	billsUsecases := &usecases.BillsUsecases{
		Repository: billsDatabase,
	}

	// CONTROLLER INSTANCE
	billsController := &controllers.BillsController{
		Usecases: billsUsecases,
	}

	return billsController
}
