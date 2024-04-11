package repositories

import "github.com/skGab/Bills-microservices/Bills-database-service/domain/entities"

type BillsRepository interface {
	GetAll(clientID string) ([]entities.BillEntity, error)
	Create(billEntity *entities.BillEntity) error
	Update(billID string, data interface{}) error
	Delete(billID string) error
	DeleteAll(billsIDs []string) error
}
