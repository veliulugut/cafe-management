package entadp

import "cafe-management/ent"

var _ RepositoryInterface = (*Repository)(nil)

type Repository struct {
	user        UserRepository
	tables      TablesRepository
	price       PriceRepository
	tablestype  TablesTypeRepostiory
	reservation ReservationRepository
	product     ProductRepository
	order       OrderRepository
	ordertype   OrdersTypeRepository
	menu        MenuRepository
}

func NewRepository(dbClient *ent.Client) *Repository {
	return &Repository{
		user:        NewUserRepository(dbClient),
		tables:      NewTablesRepository(dbClient),
		price:       NewPriceRepository(dbClient),
		tablestype:  NewTableTypeRepository(dbClient),
		product:     NewProductRepository(dbClient),
		reservation: NewReservation(dbClient),
		order:       NewOrderRepository(dbClient),
		ordertype:   NewOrdersTypeRepository(dbClient),
		menu:        NewMenuRepository(dbClient),
	}
}

func (r *Repository) Price() PriceRepository {
	return r.price
}

func (r *Repository) Reservation() ReservationRepository {
	return r.reservation
}

func (r *Repository) Table() TablesRepository {
	return r.tables
}

func (r *Repository) TableTypes() TablesTypeRepostiory {
	return r.tablestype
}

func (r *Repository) User() UserRepository {
	return r.user
}

func (r *Repository) Product() ProductRepository {
	return r.product
}

func (r *Repository) Order() OrderRepository {
	return r.order
}

func (r *Repository) OrderType() OrdersTypeRepository {
	return r.ordertype
}

func (r *Repository) Menu() MenuRepository {
	return r.menu
}
