package repository

import (
	"fmt"
	"server/models"

	"github.com/jmoiron/sqlx"
)

type ApiPostgres struct {
	db *sqlx.DB
}

func newApiPostgres(db *sqlx.DB) *ApiPostgres {
	return &ApiPostgres{
		db: db,
	}
}

func (r *ApiPostgres) NewWallet() (string, error) {
	var id string

	query := fmt.Sprintf("INSERT INTO %s (balance) VALUES ('100') RETURNING id", walletTable)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (r *ApiPostgres) GetWallet(id string) (models.Wallet, error) {
	var wallet models.Wallet

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", walletTable)
	err := r.db.Get(&wallet, query, id)
	return wallet, err
}
