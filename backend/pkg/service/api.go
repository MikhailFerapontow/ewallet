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

func (s *ApiService) SendMoney(fromId string, input models.SendMoneyInput) error {
	return s.repo.SendMoney(fromId, input)
}

func (s *ApiService) GetHistory(id string) ([]models.Transaction, error) {
	transactions, err := s.repo.GetHistory(id)
	if err != nil {
		return []models.Transaction{}, err
	}

	return transactions, nil
}
