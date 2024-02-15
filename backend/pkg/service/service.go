package service

import (
	"server/models"
	"server/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

type Api interface {
	NewWallet() (models.Wallet, error)
	GetWallet(id string) (models.Wallet, error)
	SendMoney(fromId string, input models.SendMoneyInput) error
	GetHistory(id string) ([]models.Transaction, error)
}

type Service struct {
	Api
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Api: newApiService(repo.Api),
	}
}
