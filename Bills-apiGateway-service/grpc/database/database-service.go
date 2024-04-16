package bills_database

import (
	"context"

	protoBuff "github.com/skGab/Bills-microservices/Bills-apigateway-service/grpc/database/proto"
)

type BillsDatabaseService struct {
}

func (b *BillsDatabaseService) GetAll(ctx context.Context, client *protoBuff.Client) (*protoBuff.Bills, error) {
	return &protoBuff.Bills{}, nil
}
