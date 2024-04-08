package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skGab/Bills-database-service/infrastructure/databases"
	"github.com/skGab/Bills-database-service/infrastructure/factor"
	"github.com/skGab/Bills-database-service/infrastructure/server"
)

func main() {
	// START DATABASE CONNECTION
	db := databases.DatabaseConnection()

	// INSTANCE OF BILLS CONTROLLER
	billsController := factor.BillsController(db)

	// BUILD SERVER
	server := server.NewServer(gin.Default(), billsController)

	// CONFIGURE AND RUN SERVER
	server.UpServer()
}
