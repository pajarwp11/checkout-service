package promotion

import (
	"checkout/internal/cart"
	"encoding/json"
	"math"
)

type DiscountPromotionService struct{}

func (e *DiscountPromotionService) Apply(promotion *Promotion, cart *cart.Cart) error {
	var cfg DiscountConfig

	err := json.Unmarshal(
		promotion.Config,
		&cfg,
	)
	if err != nil {
		return err
	}

	promotionItem, ok := cart.Items[promotion.SKU]
	if !ok {
		return nil
	}

	if promotionItem.Quantity < cfg.ThresholdQty {
		return nil
	}

	itemTotalPrice := float64(promotionItem.Quantity) * promotionItem.Price
	discount := itemTotalPrice * cfg.Discount
	discount = math.Round(discount*100) / 100

	promotionItem.Discount += discount
	promotionItem.Total -= discount
	promotionItem.Total = math.Round(promotionItem.Total*100) / 100

	return nil
}
