package sneakerq

type Model struct {
	Id      int     `db:"id"`
	ShopId  int     `db:"shop_id"`
	Title   string  `db:"title"`
	Price   float64 `db:"price"`
	PageUrl string  `db:"page_url"`
}
