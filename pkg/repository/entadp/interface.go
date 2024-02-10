package entadp

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"context"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context, c *dto.User) error
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, id int, c *dto.User) error
	GetById(ctx context.Context, id int) (*dto.User, error)
	ListUser(ctx context.Context) ([]*ent.User, error)
}

type TablesRepository interface {
	CreateTable(ctx context.Context, c *dto.Tables) error
	DeleteTable(ctx context.Context, id int) error
	UpdateTable(ctx context.Context, id int, c *dto.Tables) error
	ListTable(ctx context.Context) ([]*ent.Tables, error)
}

type TablesTypeRepostiory interface {
	CreateTableType(ctx context.Context, c *dto.TablesType) error
	DeleteTableType(ctx context.Context, id int) error
	UpdateTableType(ctx context.Context, id int, c *dto.TablesType) error
}

type ReservationRepository interface {
	CreateReservation(ctx context.Context, c *dto.Reservation) error
	DeleteReservation(ctx context.Context, id int) error
	UpdateReservation(ctx context.Context, id int, c *dto.Reservation) error
	GetByNameReservation(ctx context.Context, name string) (*dto.Reservation, error)
	CheckAvailability(ctx context.Context, startTime, endTime time.Time, tableID int) (bool, error)
}

type PriceRepository interface {
	CreatePrice(ctx context.Context, c *dto.Price) error
	DeletePrice(ctx context.Context, id int) error
	UpdatePrice(ctx context.Context, id int, c *dto.Price) error
	ListPrice(ctx context.Context) ([]*ent.Price, error)
}
