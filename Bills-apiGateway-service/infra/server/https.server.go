package server

import (
	"log"
	"net/http"

	"github.com/skGab/Bills-microservices/Bills-apigateway-service/app/controllers"
)

func RunHttps(adress string) {
	http.Handle("/", controllers.GateWayPipe)

	log.Fatal(http.ListenAndServe(adress, nil))
}
