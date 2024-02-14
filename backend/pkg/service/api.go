package service

import "server/pkg/repository"

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
