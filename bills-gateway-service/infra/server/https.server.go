package server

import (
	"log"
	"net/http"

	"github.com/skGab/bills-microservices/bills-gateway-service/app/controllers"
)

func RunHttps(adress string, gateway *controllers.GateWayPipe) {
	http.Handle("/gateway/:routingKey", gateway)

	log.Fatal(http.ListenAndServe(adress, nil))
}
