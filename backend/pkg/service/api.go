package service

import (
	"server/models"
	"server/pkg/repository"
)

type ApiService struct {
	repo repository.Api
}

func newApiService(repo repository.Api) *ApiService {
	return &ApiService{
		repo: repo,
	}
}

func (s *ApiService) NewWallet() (string, error) {
	return s.repo.NewWallet()
}

func (s *ApiService) GetWallet(id string) (models.Wallet, error) {
	wallet, err := s.repo.GetWallet(id)
	if err != nil {
		return models.Wallet{}, err
	}

	return wallet, nil
}
