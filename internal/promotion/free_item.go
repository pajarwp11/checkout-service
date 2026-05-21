package promotion

import (
	"checkout/internal/cart"
	"encoding/json"
	"math"
)

type FreeItemPromotionService struct{}

func (s *FreeItemPromotionService) Apply(promotion *Promotion, cart *cart.Cart) error {
	var cfg FreeItemConfig

	err := json.Unmarshal(
		promotion.Config,
		&cfg,
	)
	if err != nil {
		return err
	}

	mainItem, ok := cart.Items[promotion.SKU]
	if !ok {
		return nil
	}

	if mainItem.Quantity < cfg.ThresholdQty {
		return nil
	}

	freeQty := mainItem.Quantity * cfg.FreeItemQty

	freeItem, ok := cart.Items[cfg.FreeItemSKU]

	if !ok {
		return nil
	}

	discountQty := min(
		freeQty,
		freeItem.Quantity,
	)

	discount := float64(discountQty) * freeItem.Price
	discount = math.Round(discount*100) / 100

	freeItem.Discount += discount
	freeItem.Total -= discount
	freeItem.Total = math.Round(freeItem.Total*100) / 100

	return nil
}
