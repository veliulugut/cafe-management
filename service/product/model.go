package product

type CreateProductModel struct {
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	ProductType string  `json:"product_type"`
}

type UpdateProductModel struct {
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	ProductType string  `json:"product_type"`
}

type ProductModel struct {
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	ProductType string  `json:"product_type"`
}
