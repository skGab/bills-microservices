package main

import (
	"skGab/Bills-management-service/infrastructure"
)

func main() {
	// SERVER INSTANCE
	server := infrastructure.NewServer()

	// RUN SERVER
	server.UpServer()
}
