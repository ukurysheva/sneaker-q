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
	for _, shopObj := range shops {
		shoplist, err := shop.Call(shopObj.ClassName, shopObj.Link, shopObj.Id)
		// for _, model := range shoplist {
		fmt.Println(shoplist)
		pt.repo.Models.AddModelsList(shoplist)
		// }
		if err != nil {
			log.Fatalf("err while adding models for %s: %s", shopObj.Title, err.Error())
			//  add log
			continue
		}
	}
}
