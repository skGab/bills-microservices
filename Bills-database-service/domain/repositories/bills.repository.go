package repositories

import "github.com/skGab/Bills-microservices/Bills-database-service/domain/entities"

type BillsRepository interface {
	GetAll(clientID int) ([]entities.BillEntity, error)
	Create(billEntity *entities.BillEntity) error
	Update(billID int, data interface{}) error
	Delete(billID int) error
	DeleteAll(billsIDs []int) error
}
