package checkout

import (
	"checkout/internal/product"
	"checkout/internal/promotion"
	"encoding/json"
	"testing"
)

type ProductRepositoryMock struct{}

func (r *ProductRepositoryMock) GetBySKUs(skus []string) ([]product.Product, error) {
	return []product.Product{
		{
			SKU:   "120P90",
			Name:  "Google Home",
			Price: 49.99,
			Stock: 10,
		},
	}, nil
}

type PromotionRepositoryMock struct{}

func (r *PromotionRepositoryMock) GetBySKUs(skus []string) ([]promotion.Promotion, error) {

	cfg, _ := json.Marshal(
		promotion.BuyXForYConfig{
			ThresholdQty: 3,
			PromotionQty: 2,
		},
	)

	return []promotion.Promotion{
		{
			SKU:    "120P90",
			Type:   promotion.TypeBuyXForY,
			Config: cfg,
		},
	}, nil
}

func TestDoCheckout_Buy3Pay2GoogleHome(t *testing.T) {
	productRepo := &ProductRepositoryMock{}

	promotionRepo := &PromotionRepositoryMock{}

	promotionEngine := promotion.NewEngine()

	service := NewService(
		productRepo,
		promotionRepo,
		*promotionEngine,
	)

	request := &CheckoutRequest{
		Items: []CheckoutItem{
			{
				SKU:      "120P90",
				Quantity: 3,
			},
		},
	}

	response, err := service.DoCheckout(request)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if response.Total != 99.98 {
		t.Fatalf(
			"expected total 99.98, got %v",
			response.Total,
		)
	}

	if response.Discount != 49.99 {
		t.Fatalf(
			"expected discount 49.99, got %v",
			response.Discount,
		)
	}
}
