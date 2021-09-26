package service

import (
	sneakerq "github.com/ukurysheva/sneaker-q"
	"github.com/ukurysheva/sneaker-q/pkg/repository"
)

type Service struct {
	ModelItem
	ModelList
}

type ModelItem interface {
}

type ModelList interface {
	GetShopModels(shop string) ([]sneakerq.Model, error)
	// Search(ModelParams sneakerq.Model) []sneakerq.Model // Get list of models, search by params
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		ModelItem: NewModelService(repo.Models),
		ModelList: NewModelListService(repo.Models),
	}
}
