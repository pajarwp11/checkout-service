package promotion

import (
	"checkout/internal/cart"
	"encoding/json"
)

type PromotionService interface {
	Apply(promotion *Promotion, cart *cart.Cart) error
}

type PromotionType string

const (
	TypeFreeItem PromotionType = "FREE_ITEM"
	TypeBuyXForY PromotionType = "BUY_X_FOR_Y"
	TypeDiscount PromotionType = "DISCOUNT"
)

type Promotion struct {
	ID     int64           `json:"id"`
	SKU    string          `json:"sku"`
	Type   PromotionType   `json:"type"`
	Config json.RawMessage `json:"config"`
}

type FreeItemConfig struct {
	ThresholdQty int    `json:"threshold_qty"`
	FreeItemSKU  string `json:"free_item_sku"`
	FreeItemQty  int    `json:"free_item_qty"`
}

type BuyXForYConfig struct {
	ThresholdQty int `json:"threshold_qty"`
	PromotionQty int `json:"promotion_qty"`
}

type DiscountConfig struct {
	ThresholdQty int     `json:"threshold_qty"`
	Discount     float64 `json:"discount"`
}
