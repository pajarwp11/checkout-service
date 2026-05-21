package checkout

type CheckoutItem struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem
}

type CheckoutResponse struct {
	Total    float64 `json:"total"`
	SubTotal float64 `json:"sub_total"`
	Discount float64 `json:"discount"`
}
