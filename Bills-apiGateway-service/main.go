package main

import "github.com/skGab/Bills-microservices/Bills-apigateway-service/infra/server"

func main() {
	// OPEN HTTPS SERVER
	server.RunHttps(":8080")

	//OPEN GRPC SERVER
	server.RunGrpc(":3030")
}
