package shop

import (
	"log"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	sneakerq "github.com/ukurysheva/sneaker-q"
)

var NikeId int

func NikeParseMenu(shopInfo sneakerq.Shop) []*MenuItem {
	var menu []*MenuItem

	// Parse main Menu - get women and men collections
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{shopInfo.Link},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("li.pre-desktop-menu-item").Each(func(i int, s *goquery.Selection) {
				menuLink := s.Find("a.pre-desktop-menu-link")
				menuTitle := menuLink.Text()
				if menuTitle == "Женщины" || menuTitle == "Мужчины" {
					menuPageUrl, exist := menuLink.Attr("href")
					if !exist {
						return
					}
					menuItem := &MenuItem{Title: menuTitle, Link: menuPageUrl}
					menu = append(menu, menuItem)
				}

			})
		},
	}).Start()

	var shoesUrl []*MenuItem
	for _, v := range menu {
		geziyor.NewGeziyor(&geziyor.Options{
			StartURLs: []string{v.Link},
			ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
				submenu := r.HTMLDoc.Find("div.sticky-inner-wrapper")

				nav := submenu.Find("nav")
				subList := nav.Find("ul")
				subList.Find("li").Each(func(i int, s *goquery.Selection) {
					menuLink := s.Find("a")
					menuTitle := menuLink.Text()
					menuPageType, exist := menuLink.Attr("aria-label")
					if exist && menuPageType == "Обувь" {
						menuPageUrl, existHref := menuLink.Attr("href")
						if existHref {
							menuItem := &MenuItem{Title: menuTitle, Link: menuPageUrl}
							shoesUrl = append(shoesUrl, menuItem)
						}

					}
				})
			},
		}).Start()
	}

	return shoesUrl
}

// TODO: instead of append - send to channel
// https://stackoverflow.com/questions/18499352/golang-concurrency-how-to-append-to-the-same-slice-from-different-goroutines
func NikeParseModels(shopInfo sneakerq.Shop, menus []*MenuItem) []*sneakerq.Model {
	var models []*sneakerq.Model

	for _, v := range menus {
		geziyor.NewGeziyor(&geziyor.Options{
			StartURLs: []string{v.Link},
			ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
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
					model := &sneakerq.Model{Title: modTitle, Price: price, PageUrl: modPageUrl, ShopId: shopInfo.Id}
					models = append(models, model)
				})
			},
		}).Start()
	}

	return models
}

// func getModels()
func modifyPrice(pr string) string {
	var re = regexp.MustCompile(`\p{Z}|₽`)
	s := re.ReplaceAllString(pr, ``)
	return s
}
