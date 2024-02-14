package repository

import "github.com/jmoiron/sqlx"

type Api interface {
	NewWallet() (string, error)
}

type Repository struct {
	Api
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Api: newApiPostgres(db),
	}
}
