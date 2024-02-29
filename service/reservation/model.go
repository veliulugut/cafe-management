package reservation

import "time"

type CreateReservationModel struct {
	FullName    string    `json:"full_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	TableID     int       `json:"table_id"`
	PhoneNumber string    `json:"phone_number"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Status      string    `json:"status"`
}

type UpdateReservationModel struct {
	FullName    string    `json:"full_name"`
	UpdatedAt   time.Time `json:"updated_at"`
	TableID     int       `json:"table_id"`
	PhoneNumber string    `json:"phone_number"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Status      string    `json:"status"`
}

type ReservationModel struct {
	FullName    string    `json:"full_name"`
	TableID     int       `json:"table_id"`
	PhoneNumber string    `json:"phone_number"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
