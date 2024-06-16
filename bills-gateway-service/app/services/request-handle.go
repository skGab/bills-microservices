package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Payload struct {
	Id         *string
	RoutingKey string
	Data       interface{}
}

func (*Handlers) HandleRequest(gin *gin.Context) *Payload {
	// GETTING THE ROUTING KEY FROM HEADER
	routingKey := gin.Request.Header.Get("x-routing-key")

	// GETTING ROUTE ID
	id := gin.Param("id")

	if routingKey == "" {
		gin.JSON(http.StatusBadRequest, "Routing Key não encontrada no Header da requisição")
	}

	var body *Body

	if err := gin.ShouldBindJSON(&body); err != nil {
		gin.JSON(http.StatusBadRequest, err)
	}

	payload := &Payload{Id: &id, RoutingKey: routingKey, Data: body}

	return payload
}
