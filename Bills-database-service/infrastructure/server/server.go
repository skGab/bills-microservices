package server

import (
	"github.com/gin-gonic/gin"
	"github.com/skGab/Bills-database-service/app/controllers"
)

type Server struct {
	Router           *gin.Engine
	BillsControlller *controllers.BillsController
}

// CONSTRUCTOR
func NewServer(gin *gin.Engine, billsController *controllers.BillsController) *Server {
	return &Server{
		Router:           gin,
		BillsControlller: billsController,
	}
}

func (server *Server) UpServer() {
	// CORS CONFIGURATION
	server.Router.Use(CorsMiddleware())

	// CREATE THE API PATHS
	server.Router.GET("/bills/getAll/clientID", server.BillsControlller.GetBills)
	server.Router.POST("/bills/create", server.BillsControlller.CreateBill)
	server.Router.PUT("/bills/update/billID", server.BillsControlller.UpdateBill)
	server.Router.DELETE("/bills/delete/billID", server.BillsControlller.DeleteBill)
	server.Router.DELETE("/bills/deleteAll/billsIDs", server.BillsControlller.DeleteAllBills)

	// RUN THE SERVER
	server.Router.Run()
}
