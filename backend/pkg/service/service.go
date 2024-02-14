package service

import (
	"server/models"
	"server/pkg/repository"
)

type Api interface {
	NewWallet() (string, error)
	GetWallet(id string) (models.Wallet, error)
}

type Service struct {
	Api
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Api: newApiService(repo.Api),
	}
}
