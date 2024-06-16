package services

type Handlers struct{}

// GetAll(clientID string) ([]entities.BillEntity, error)
// 	Create(billEntity *entities.BillEntity) error
// 	Update(billID string, data interface{}) error
// 	Delete(billID string) error
// 	DeleteAll(billsIDs []string) error

type Params struct {
	id         *string `param:"id"`
	routingKey string  `param:"action"`
}
type Body struct {
	id   *int         `body:"id"`
	ids  *[]int       `body:"ids"`
	data *interface{} `body:"data"`
}

func (hs *Handlers) DatabaseHandle(body Body, params Params) (string, error) {

	// SE ID NA ROTA + ACTION = GET
	// SE DADOS NO REQUEST + ACTION = POST
	// SE ID + DADOS NO REQUEST + ACTION = UPDATE
	// SE ID NO REQUEST + ACTION = DELETE
	// SE IDS NO REQUEST + ACTION = DELETE ALL

	return "Request redirecionado com sucesso", nil
}
