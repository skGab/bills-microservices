package factory

import (
	"github.com/skGab/bills-microservices/bills-gateway-service/app/controllers"
	"github.com/skGab/bills-microservices/bills-gateway-service/app/services"
)

func GateWayFactor(services *services.Handlers) *controllers.GateWayPipe {
	gateway := &controllers.GateWayPipe{Services: services}

	return gateway
}
