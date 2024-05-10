package controllers

import (
	"net/http"

	"github.com/skGab/Bills-microservices/Bills-apigateway-service/app/handlers"

	"github.com/gin-gonic/gin"
)

type GateWayPipe struct {
	handlers *handlers.Handlers
}

// RIGHT HERE THE LOGIC TO ROUTE THE INCOMING REQUESTS
// Here i would get the request and check his type, them i would redirect to the right grpc method.
func (gp *GateWayPipe) Run(gin *gin.Context) {

	action := gin.Param("action")

	switch action {
	case "database":

		result, err := gp.handlers.DatabaseHandle()

		if err {
			gin.JSON(http.StatusInternalServerError, err.error())
		}

		gin.JSON(http.StatusOK, result)
	default:
		gin.JSON(http.StatusBadRequest, gin.H{"Algo deu errado durante o request"})
	}

}
