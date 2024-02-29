package ordertype

type CreateOrderTypeModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateOrderType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListOrderType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
