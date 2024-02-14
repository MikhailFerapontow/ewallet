package service

import "server/pkg/repository"

type Api interface {
	NewWallet() (string, error)
}

type Service struct {
	Api
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Api: newApiService(repo.Api),
	}
}
