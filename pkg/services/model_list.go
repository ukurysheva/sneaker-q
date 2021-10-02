package services

import (
	sneakerq "github.com/ukurysheva/sneaker-q"
	"github.com/ukurysheva/sneaker-q/pkg/repository"
)

type ModelListService struct {
	repo repository.Models
}

func NewModelListService(repo repository.Models) *ModelListService {
	return &ModelListService{
		repo: repo,
	}
}

func (mlist *ModelListService) GetShopModels(shopname string) ([]sneakerq.Model, error) {
	return mlist.repo.GetShopModels(shopname)
}

func (mlist *ModelListService) GetModelsByParams(searchParams sneakerq.SearchParams) ([]sneakerq.Model, error) {
	return mlist.repo.GetModelsByParams(searchParams)
}
