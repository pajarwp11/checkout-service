package promotion

import "checkout/internal/cart"

type Engine struct {
	processors map[PromotionType]PromotionService
}

func NewEngine() *Engine {
	return &Engine{
		processors: map[PromotionType]PromotionService{
			TypeFreeItem: &FreeItemPromotionService{},
			TypeBuyXForY: &BuyXForYPromotionService{},
			TypeDiscount: &DiscountPromotionService{},
		},
	}
}

func (e *Engine) Apply(promotion *Promotion, cart *cart.Cart) error {
	processor, ok := e.processors[promotion.Type]
	if !ok {
		return nil
	}

	return processor.Apply(
		promotion,
		cart,
	)
}
