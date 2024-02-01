package dto

import "time"

type Reservation struct {
	FullName    string    `json:"full_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	TableID     int       `json:"table_id"`
	PhoneNumber string    `json:"phone_number"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
