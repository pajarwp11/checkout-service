package checkout

import (
	"checkout/internal/cart"
	"checkout/internal/product"
	"checkout/internal/promotion"
	"math"
)

func (s *Service) DoCheckout(request *CheckoutRequest) (*CheckoutResponse, error) {
	skus := make([]string, 0, len(request.Items))
	for _, item := range request.Items {
		skus = append(skus, item.SKU)
	}

	products, err := s.productRepo.GetBySKUs(skus)
	if err != nil {
		return nil, err
	}

	productMap := make(map[string]product.Product)
	for _, p := range products {
		productMap[p.SKU] = p
	}

	promotions, err := s.promotionRepo.GetBySKUs(skus)
	if err != nil {
		return nil, err
	}

	promotionMap := make(map[string][]promotion.Promotion)
	for _, promo := range promotions {
		promotionMap[promo.SKU] = append(promotionMap[promo.SKU], promo)
	}

	carts := cart.Cart{
		Items: make(map[string]*cart.CartItem),
	}
	for _, item := range request.Items {
		productData, ok := productMap[item.SKU]
		if !ok {
			return nil, ErrProductNotFound
		}

		if productData.Stock < item.Quantity {
			return nil, ErrInsufficientStock
		}

		total := productData.Price * float64(item.Quantity)
		total = math.Round(total*100) / 100
		cartItem := cart.CartItem{
			Name:     productData.Name,
			Price:    productData.Price,
			Quantity: item.Quantity,
			Total:    total,
			SubTotal: total,
		}
		carts.Items[productData.SKU] = &cartItem
	}

	for sku, promotions := range promotionMap {
		_, ok := carts.Items[sku]
		if !ok {
			continue
		}

		for _, promo := range promotions {
			err := s.promotionEngine.Apply(
				&promo,
				&carts,
			)
			if err != nil {
				return nil, err
			}
		}
	}

	var response CheckoutResponse
	for _, c := range carts.Items {
		response.Total += c.Total
		response.SubTotal += c.SubTotal
		response.Discount += c.Discount
	}

	return &response, nil
}
