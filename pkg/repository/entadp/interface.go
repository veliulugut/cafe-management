//go:generate mockgen -destination=repository_mocks.go -package=entadp cafe-management/pkg/repository/entadp RepositoryInterface
package entadp

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"context"
	"time"
)

type RepositoryInterface interface {
	User() UserRepository
	Table() TablesRepository
	TableTypes() TablesTypeRepostiory
	Reservation() ReservationRepository
	Price() PriceRepository
	Product() ProductRepository
	Order() OrderRepository
	OrderType() OrdersTypeRepository
	Menu() MenuRepository
	QRCode() QRCodeRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, c *dto.User) error
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, id int, c *dto.User) error
	GetById(ctx context.Context, id int) (*dto.User, error)
	ListUser(ctx context.Context) ([]*ent.User, error)
	GetByUserName(ctx context.Context, name string) (*dto.User, error)
	GetByUserEmail(ctx context.Context, email string) (*dto.User, error)
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
	GetReservationHistory(ctx context.Context, username string) ([]*ent.Reservation, error)
	CancelReservation(ctx context.Context, id int) error
	ListReservationByTable(ctx context.Context, tableID int) (*dto.Reservation, error)
}

type PriceRepository interface {
	CreatePrice(ctx context.Context, c *dto.Price) error
	DeletePrice(ctx context.Context, id int) error
	UpdatePrice(ctx context.Context, id int, c *dto.Price) error
	ListPrice(ctx context.Context) ([]*ent.Price, error)
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, c *dto.Product) error
	DeleteProduct(ctx context.Context, id int) error
	UpdateProduct(ctx context.Context, id int, c *dto.Product) error
	ListProduct(ctx context.Context) ([]*ent.Product, error)
	GetById(ctx context.Context, id int) (*dto.Product, error)
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, c *dto.Order) error
	DeleteOrder(ctx context.Context, id int) error
	UpdateOrder(ctx context.Context, id int, c *dto.Order) error
	ListOrder(ctx context.Context) ([]*ent.Order, error)
	GetById(ctx context.Context, id int) (*dto.Order, error)
}

type OrdersTypeRepository interface {
	CreateOrderType(ctx context.Context, c *dto.OrderType) error
	DeleteOrderType(ctx context.Context, id int) error
	UpdateOrderType(ctx context.Context, id int, c *dto.OrderType) error
	ListOrderType(ctx context.Context) ([]*ent.OrderType, error)
}

type MenuRepository interface {
	CreateMenu(ctx context.Context, c *dto.Menu) error
	DeleteMenu(ctx context.Context, id int) error
	UpdateMenu(ctx context.Context, id int, c *dto.Menu) error
	ListMenu(ctx context.Context) ([]*ent.Menu, error)
	GetById(ctx context.Context, id int) (*dto.Menu, error)
}

type QRCodeRepository interface {
	CreateQRCode(ctx context.Context, url string, image []byte) error
	UpdateQRCode(ctx context.Context, id int, c *dto.QRCode) error
	DeleteQRCode(ctx context.Context, id int) error
	ListQRCode(ctx context.Context) ([]*ent.QrCode, error)
}
