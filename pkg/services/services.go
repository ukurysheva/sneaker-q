package services

import (
	sneakerq "github.com/ukurysheva/sneaker-q"
	"github.com/ukurysheva/sneaker-q/pkg/repository"
)

type Service struct {
	Model
}

type Model interface {
	GetShopModels(shop string) ([]sneakerq.Model, error)
	GetModelsByParams(searchParams sneakerq.SearchParams) ([]sneakerq.Model, error)
	GetModelById(id int) (sneakerq.Model, error)
	// Search(ModelParams sneakerq.Model) []sneakerq.Model // Get list of models, search by params
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Model: NewModelService(repo.Models),
	}
}
