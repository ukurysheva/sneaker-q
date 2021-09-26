package service

import "github.com/ukurysheva/sneaker-q/pkg/repository"

type ModelService struct {
	repo repository.Models
}

func NewModelService(repo repository.Models) *ModelService {
	return &ModelService{
		repo: repo,
	}
}
