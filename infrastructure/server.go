package infrastructure

import "github.com/gin-gonic/gin"

type Server struct{}

// CONSTRUCTOR
func NewServer() *Server {
	return &Server{}
}

func (*Server) UpServer() {
	// GIN INSTANCE
	router := gin.Default()

	// RUN THE SERVER
	router.Run()
}
