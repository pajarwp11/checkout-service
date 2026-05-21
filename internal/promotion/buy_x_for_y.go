package promotion

import (
	"checkout/internal/cart"
	"encoding/json"
	"math"
)

type BuyXForYPromotionService struct{}

func (s *BuyXForYPromotionService) Apply(promotion *Promotion, cart *cart.Cart) error {
	var cfg BuyXForYConfig

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

	groupCount := promotionItem.Quantity / cfg.ThresholdQty
	remainder := promotionItem.Quantity % cfg.ThresholdQty

	paidQty := (groupCount * cfg.PromotionQty) + remainder

	discountQty := promotionItem.Quantity - paidQty

	discount := float64(discountQty) * promotionItem.Price
	discount = math.Round(discount*100) / 100

	promotionItem.Discount += discount
	promotionItem.Total -= discount
	promotionItem.Total = math.Round(promotionItem.Total*100) / 100

	return nil
}
