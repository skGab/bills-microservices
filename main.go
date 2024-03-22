package main

import (
	"skGab/Bills-management-service/infrastructure"
	"skGab/Bills-management-service/infrastructure/databases"
	"skGab/Bills-management-service/infrastructure/factor"

	"github.com/gin-gonic/gin"
)

func main() {
	// START DATABASE CONNECTION
	db := databases.DatabaseConnection()

	// INSTANCE OF BILLS CONTROLLER
	billsController := factor.BillsController(db)

	// BUILD SERVER
	server := infrastructure.NewServer(gin.Default(), billsController)

	// CONFIGURE AND RUN SERVER
	server.UpServer()
}
