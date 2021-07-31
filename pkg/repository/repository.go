package repository

import (
	"github.com/jmoiron/sqlx"
	sneakerq "github.com/ukurysheva/sneaker-q"
)

type Repository struct {
	Shops
	Models
}

type Shops interface {
	GetShops() ([]sneakerq.Shop, error)
}

type Models interface {
	AddModelsList([]*sneakerq.Model) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Shops:  NewShopsPostgres(db),
		Models: NewModelsPostgres(db),
	}
}
