package cart

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
	Total    float64
	SubTotal float64
	Discount float64
}

type Cart struct {
	Items map[string]*CartItem
}
