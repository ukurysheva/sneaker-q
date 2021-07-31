package shop

import (
	"log"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
	sneakerq "github.com/ukurysheva/sneaker-q"
)

var models []*sneakerq.Model
var NikeId int

func nikeParseModels(url string, shop_id int) []*sneakerq.Model {
	NikeId = shop_id
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
		ParseFunc: parseModels,
		Exporters: []export.Exporter{&export.JSON{FileName: "nike.json"}},
	}).Start()

	return models
}

func parseModels(g *geziyor.Geziyor, r *client.Response) {

	r.HTMLDoc.Find("figure").Each(func(i int, s *goquery.Selection) {
		modTitle := s.Find("div.product-card__title").Text()
		modPageUrlElem := s.Find("a.product-card__link-overlay")
		modPageUrl, exist := modPageUrlElem.Attr("href")
		if !exist {
			return
		}
		modPrice := s.Find("div.product-price").Text()
		modPrice = modifyPrice(modPrice)
		price, err := strconv.ParseFloat(modPrice, 64)
		if err != nil {
			log.Fatalf("err while parsing: %s", err.Error())
			return
		}
		model := &sneakerq.Model{Title: modTitle, Price: price, PageUrl: modPageUrl, ShopId: NikeId}
		models = append(models, model)
	})
}

func modifyPrice(pr string) string {
	var re = regexp.MustCompile(`\p{Z}|₽`)
	s := re.ReplaceAllString(pr, ``)
	return s
}
