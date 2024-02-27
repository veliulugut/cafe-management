package dto

import "time"

type Menu struct {
	MenuID       int       `json:"menu_id"`
	Name         string    `json:"name"`
	Category     string    `json:"category"`
	Price        float64   `json:"price"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	MenuImageUrl string    `json:"menu_image_url"`
	UpdatedAt    time.Time `json:"updated_at"`
}
