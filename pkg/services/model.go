package services

import (
	sneakerq "github.com/ukurysheva/sneaker-q"
	"github.com/ukurysheva/sneaker-q/pkg/repository"
)

type ModelService struct {
	repo repository.Models
}

func NewModelService(repo repository.Models) *ModelService {
	return &ModelService{
		repo: repo,
	}
}

func (model *ModelService) GetModelById(id int) (sneakerq.Model, error) {
	return model.repo.GetModelById(id)
}
func (model *ModelService) GetShopModels(shopname string) ([]sneakerq.Model, error) {
	return model.repo.GetShopModels(shopname)
}

func (model *ModelService) GetModelsByParams(searchParams sneakerq.SearchParams) ([]sneakerq.Model, error) {
	return model.repo.GetModelsByParams(searchParams)
}
