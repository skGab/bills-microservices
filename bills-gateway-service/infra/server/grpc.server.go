package server

import (
	"log"
	"net"

	protoBuff "github.com/skGab/bills-microservices/bills-gateway-service/domain/proto"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	protoBuff.UnimplementedBillsDatabaseServiceServer
}

func RunGrpc(adress string) {
	// CREATE A LISTENER TO WATCH ON THE PORT 3030 AND WITH TPC NETWORK
	lis, err := net.Listen("tcp", adress)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// CREATE EMPTY ARRAY OF OPTIONS
	var opts []grpc.ServerOption

	// CREATE THE GRPC SERVER AND PASS THE CREATED ARRAY OPTIONS
	grpcServer := grpc.NewServer(opts...)

	// REGISTER THE SERVICE INSTANCE WITH THE GRPC SERVER
	protoBuff.RegisterBillsDatabaseServiceServer(grpcServer, &GrpcServer{})

	// START THE SERVER WITH THE DEFINED LISTENER
	// THE LISTENER HAS INFORMATION ABOUT THE PORT AND THE CONECTION'S TYPE
	grpcServer.Serve(lis)
}
