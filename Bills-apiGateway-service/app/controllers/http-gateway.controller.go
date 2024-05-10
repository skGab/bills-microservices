package controllers

import (
	"net/http"

	"github.com/skGab/Bills-microservices/Bills-apigateway-service/app/services"

	"github.com/gin-gonic/gin"
)

type GateWayPipe struct {
	services *services.Handlers
}

type url struct {
	name  string  `url:"name" binding:"required"`
	total float32 `url:"total" binding:"required"`
	age   int     `url:"age" binding:"required"`
}

func (gp *GateWayPipe) Run(gin *gin.Context) {

	// VERIFY IF HAS DATA ON THE PARAM
	// GET THE ACTION
	var urlData *url

	if err := gin.ShouldBindUri(urlData); err != nil {
		gin.JSON(http.StatusBadRequest, err)
	}

	var body *interface{}

	if err := gin.ShouldBindJSON(body); err != nil {
		gin.JSON(http.StatusBadRequest, err)
	}

	switch action {
	case "database":

		result, err := gp.services.DatabaseHandle()

		if err != nil {
			gin.JSON(http.StatusInternalServerError, err)
		}

		gin.JSON(http.StatusOK, result)
	default:

		gin.JSON(http.StatusBadRequest, struct{ msg string }{msg: "Algo deu errado durante a requisição"})
	}

}
