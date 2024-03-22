package infrastructure

import (
	"skGab/Bills-management-service/app/controllers"

	"github.com/gin-gonic/gin"
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

func (s *Server) UpServer() {
	// CORS
	s.Router.Use(corsMiddleware())

	// ROUTES HANDLE
	s.Router.POST("/bills/create", s.BillsControlller.CreateBill)

	// RUN THE SERVER
	s.Router.Run()
}

// CORS CONFIGURATION
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
