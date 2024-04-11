package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/skGab/Bills-microservices/Bills-database-service/infrastructure/databases"
	"github.com/skGab/Bills-microservices/Bills-database-service/infrastructure/factor"
	"github.com/skGab/Bills-microservices/Bills-database-service/infrastructure/server"
)

func main() {
	// DISABLE THE LOG BUILD ON TERMINAL
	gin.SetMode(gin.ReleaseMode)

	// START DATABASE CONNECTION
	db := databases.DatabaseConnection()

	// MIGRATE ENTITYS TO DB TABLES
	// db.AutoMigrate(&entities.BillEntity{})

	// CREATE THE INSTANCE OF VALIDATOR CLASS
	validate := validator.New(validator.WithRequiredStructEnabled())

	// INSTANCE OF BILLS CONTROLLER
	billsController := factor.BillsController(db, validate)

	// BUILD SERVER
	server := server.New(gin.Default(), billsController)

	// CONFIGURE AND RUN SERVER
	server.UpServer()
}
