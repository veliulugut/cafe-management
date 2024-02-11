package entadp

import "cafe-management/ent"

var _ RepositoryInterface = (*Repository)(nil)

type Repository struct {
	user        UserRepository
	tables      TablesRepository
	price       PriceRepository
	tablestype  TablesTypeRepostiory
	reservation ReservationRepository
}

func NewRepository(dbClient *ent.Client) *Repository {
	return &Repository{
		user:       NewUserRepository(dbClient),
		tables:     NewTablesRepository(dbClient),
		price:      NewPriceRepository(dbClient),
		tablestype: NewTableTypeRepository(dbClient),
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
