package sneakerq

type Shop struct {
	Id        int
	Title     string
	Link      string
	ClassName string `db:"class_name"`
}

type ShopList struct {
	List []*Shop
}

type MenuPage struct {
	Title string
	Link  string
}
