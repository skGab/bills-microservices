package bills_database

import (
	"context"
	"errors"

	protoBuff "github.com/skGab/Bills-microservices/Bills-apigateway-service/grpc/database/proto"
)

type BillsDatabaseService struct {
}

func (b *BillsDatabaseService) GetAll(ctx context.Context, client *protoBuff.Client) (*protoBuff.Bills, error) {
	// VALIDATE THE REQUEST
	if client == nil {
		return nil, errors.New("nenhum ID encontrado no corpo da requisição")
	}

	// SEARCH FOR THE CLIENT ON THE DATABASE

	// CHECK FOR ERROS AFTER THE SEARCH
	// RETURN THE CLIENT
	return &protoBuff.Bills{}, nil
}
