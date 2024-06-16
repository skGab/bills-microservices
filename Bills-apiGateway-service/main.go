package main

import (
	"github.com/skGab/Bills-microservices/Bills-apigateway-service/app/services"
	"github.com/skGab/Bills-microservices/Bills-apigateway-service/infra/factory"
	"github.com/skGab/Bills-microservices/Bills-apigateway-service/infra/server"
)

func main() {
	//BUILD THE GATEWAY CONTROLLER
	gateway := factory.GateWayFactor(&services.Handlers{})

	// OPEN HTTPS SERVER
	server.RunHttps(":8080", gateway)

	//OPEN GRPC SERVER
	server.RunGrpc(":3030")
}
