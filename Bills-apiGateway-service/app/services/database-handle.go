package services

type Handlers struct{}

// GetAll(clientID string) ([]entities.BillEntity, error)
// 	Create(billEntity *entities.BillEntity) error
// 	Update(billID string, data interface{}) error
// 	Delete(billID string) error
// 	DeleteAll(billsIDs []string) error

type Param struct {
	id     *string
	action *string
}
type Body struct {
	id   *int
	ids  *[]int
	data *interface{}
}

func (hs *Handlers) DatabaseHandle(Body, Param) (string, error) {

	// SE ID NA ROTA + ACTION = GET
	// SE DADOS NO REQUEST + ACTION = POST
	// SE ID + DADOS NO REQUEST + ACTION = UPDATE
	// SE ID NO REQUEST + ACTION = DELETE
	// SE IDS NO REQUEST + ACTION = DELETE ALL

	return "Request redirecionado com sucesso", nil
}
