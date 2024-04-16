package server

import (
	"log"
	"net"

	protoBuff "github.com/skGab/Bills-microservices/Bills-apigateway-service/grpc/database/proto"
	"google.golang.org/grpc"
)

type server struct {
	protoBuff.UnimplementedBillsDatabaseServiceServer
}

func Run() {
	// CREATE A LISTENER TO WATCH ON THE PORT 3030 AND WITH TPC NETWORK
	lis, err := net.Listen("tcp", "localhost:3030")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// CREATE EMPTY ARRAY OF OPTIONS
	var opts []grpc.ServerOption

	// CREATE THE GRPC SERVER AND PASS THE CREATED ARRAY OPTIONS
	grpcServer := grpc.NewServer(opts...)

	// REGISTER THE SERVICE INSTANCE WITH THE GRPC SERVER
	protoBuff.RegisterBillsDatabaseServiceServer(grpcServer, &server{})

	// START THE SERVER WITH THE DEFINED LISTENER
	// THE LISTENER HAS INFORMATION ABOUT THE PORT AND THE CONECTION'S TYPE
	grpcServer.Serve(lis)
}
