package menu

import "time"

type CreateMenuModel struct {
	MenuID       int       `json:"menu_id"`
	Name         string    `json:"name"`
	Category     string    `json:"category"`
	Price        float64   `json:"price"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"example:"2006-01-02T15:04:05Z"`
	MenuImageUrl string    `json:"menu_image_url"`
	UpdatedAt    time.Time `json:"created_at"example:"2006-01-02T15:04:05Z"`
}

type UpdateMenuModel struct {
	MenuID       int       `json:"menu_id"`
	Name         string    `json:"name"`
	Category     string    `json:"category"`
	Price        float64   `json:"price"`
	Description  string    `json:"description"`
	MenuImageUrl string    `json:"menu_image_url"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type MenuModel struct {
	MenuID       int       `json:"menu_id"`
	Name         string    `json:"name"`
	Category     string    `json:"category"`
	Price        float64   `json:"price"`
	Description  string    `json:"description"`
	MenuImageUrl string    `json:"menu_image_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
