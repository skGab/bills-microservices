package repositories

import "skGab/Bills-management-service/domain/entities"

type BillsRepository interface {
	Create(billEntity *entities.BillEntity) *error
}
