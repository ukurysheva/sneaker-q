package parser

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/ukurysheva/sneaker-q/internal/parser/shop"
	"github.com/ukurysheva/sneaker-q/pkg/repository"
)

type ParserTask struct {
	repo *repository.Repository
}

func NewParserTask(repo *repository.Repository) *ParserTask {
	return &ParserTask{repo: repo}
}

func (pt *ParserTask) ParseTask() {
	shops, err := pt.repo.Shops.GetShops()

	if err != nil {
		logrus.Fatalf("error getting shops db: %s", err.Error())
		// add log - fatal cron while getting shops
		return
	}
	fmt.Println(shops)

	// looping through shops classes and adding models
	for _, shopInfo := range shops {
		menus := shop.ParseMenu(shopInfo)

		// Instead of getting like that - make a goroutine and get models from a channel
		models := shop.ParseModels(shopInfo, menus)
		pt.repo.Models.AddModelsList(models)
		if err != nil {
			log.Fatalf("err while adding models for %s: %s", shopInfo.Title, err.Error())
			//  add log
			continue
		}
	}
}
