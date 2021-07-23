package repository

import "github.com/jmoiron/sqlx"

type ParseWork struct {
	ImportModels
}

type ImportModels interface {
	GetSources(model)
	GetModels()
}

func NewParseWork(db *sqlx.DB) *ParseWork {
	return &ParseWork{
		ImportModels: NewImportModelsWork(db),
	}
}
