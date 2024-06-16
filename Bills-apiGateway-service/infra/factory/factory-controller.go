package factory

import (
	"github.com/skGab/Bills-microservices/Bills-apigateway-service/app/controllers"
	"github.com/skGab/Bills-microservices/Bills-apigateway-service/app/services"
)

func GateWayFactor(services *services.Handlers) *controllers.GateWayPipe {
	gateway := &controllers.GateWayPipe{Services: services}

	return gateway
}
