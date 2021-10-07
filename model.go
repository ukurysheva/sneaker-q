package sneakerq

type Model struct {
	Id      int     `db:"id"`
	ShopId  int     `db:"shop_id" json:"-"`
	Title   string  `db:"title"`
	Size    string  `db:"size"`
	Price   float64 `db:"price"`
	Avail   string  `db:"avail"`
	ImgUrl  string  `db:"img_url"`
	PageUrl string  `db:"page_url"`
}

type SearchParams struct {
	Id        int `db:"id" json:"id"`
	ShopId    int `db:"shop_id" json:"-"`
	Sex       []string
	Size      []string `json:"size"`
	PriceFrom float64  `json:"price_from"`
	PriceTo   float64  `json:"price_to"`
	Avail     string
	Color     []string
}
