package repository

import (
	"server/models"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository.go

type Api interface {
	NewWallet() (models.Wallet, error)
	GetWallet(id string) (models.Wallet, error)
	SendMoney(fromId string, input models.SendMoneyInput) error
	GetHistory(id string) ([]models.Transaction, error)
}

type Repository struct {
	Api
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Api: newApiPostgres(db),
	}
}
