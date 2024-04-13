package server

import (
	"fmt"
	"log"
	"net"

	protoBuff "github.com/skGab/Bills-microservices/Bills-apigateway-service/grpc/database"
	"github.com/skGab/Bills-microservices/Bills-apigateway-service/grpc/database/services"
	"google.golang.org/grpc"
)

func newServer() *services.BillsDatabaseService {
	return &services.BillsDatabaseService{
		savedFeatures: ""
	}
}

func Run() {
	port := ":3030"
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	protoBuff.RegisterBillsDatabaseServer(grpcServer, newServer())

	// START THE SERVER WITH THE DEFINED LISTENER
	// THE LISTENER HAS INFORMATION ABOUT THE PORT AND THE CONECTION'S TYPE
	grpcServer.Serve(lis)
}
