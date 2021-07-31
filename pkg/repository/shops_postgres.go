package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	sneakerq "github.com/ukurysheva/sneaker-q"
)

type ShopsPostgres struct {
	db *sqlx.DB
}

func NewShopsPostgres(db *sqlx.DB) *ShopsPostgres {
	return &ShopsPostgres{db: db}
}

func (r *ShopsPostgres) GetShops() ([]sneakerq.Shop, error) {
	// shop := sneakerq.Shop{}
	var shopsList []sneakerq.Shop
	query := fmt.Sprintf("SELECT * FROM %s", shopsTable)
	err := r.db.Select(&shopsList, query)

	return shopsList, err
}
