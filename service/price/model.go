package price

type CreatePriceModel struct {
	Price     float64 `json:"price"`
	PriceName string  `json:"price_name"`
}

type UpdatePriceModel struct {
	Price     float64 `json:"price"`
	PriceName string  `json:"price_name"`
}

type PriceModel struct {
	Price     float64 `json:"price"`
	PriceName string  `json:"price_name"`
}
