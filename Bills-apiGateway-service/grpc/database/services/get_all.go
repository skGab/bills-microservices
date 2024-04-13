package services

import (
	"context"

	pb "github.com/skGab/Bills-microservices/Bills-apigateway-service/grpc/database"
	"google.golang.org/protobuf/proto"
)

type BillsDatabaseService struct {
	savedFeatures []*pb.Bills // read-only after initialized
}

func (b *BillsDatabaseService) GetAll(ctx context.Context, client *pb.Client) (*pb.Bills, error) {
	for _, bill := range b.savedFeatures {
		if proto.Equal(bill, client) {
			return bill, nil
		}
	}
	// No feature was found, return an unnamed feature
	return &pb.Bills{}, nil
}
