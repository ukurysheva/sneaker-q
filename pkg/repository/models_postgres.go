package repository

import "github.com/jmoiron/sqlx"

type ImportModelsWork struct {
	db *sqlx.DB
}

// Getting models from shoplists
func NewImportModelsWork(db *sqlx.DB) *ImportModelsWork {
	return &ImportModelsWork{db: db}
}
