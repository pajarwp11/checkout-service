package checkout

import (
	"checkout/internal/product"
	"checkout/internal/promotion"
	"encoding/json"
	"testing"
)

type ProductRepositoryMock struct {
	items map[string]product.Product
}

func (r *ProductRepositoryMock) GetBySKUs(skus []string) ([]product.Product, error) {
	result := make([]product.Product, 0)

	for _, sku := range skus {
		if p, ok := r.items[sku]; ok {
			result = append(result, p)
		}
	}

	return result, nil
}

type PromotionRepositoryMock struct {
	promos []promotion.Promotion
}

func (r *PromotionRepositoryMock) GetBySKUs(skus []string) ([]promotion.Promotion, error) {
	return r.promos, nil
}

func mustJSON(v any) json.RawMessage {
	b, _ := json.Marshal(v)
	return b
}

func TestDoCheckout_Buy3Pay2GoogleHome(t *testing.T) {
	productRepo := &ProductRepositoryMock{
		items: map[string]product.Product{
			"120P90": {
				SKU:   "120P90",
				Name:  "Google Home",
				Price: 49.99,
				Stock: 10,
			},
		},
	}

	promotionRepo := &PromotionRepositoryMock{
		promos: []promotion.Promotion{
			{
				SKU:  "120P90",
				Type: promotion.TypeBuyXForY,
				Config: mustJSON(promotion.BuyXForYConfig{
					ThresholdQty: 3,
					PromotionQty: 2,
				}),
			},
		},
	}

	service := NewService(productRepo, promotionRepo, *promotion.NewEngine())

	req := &CheckoutRequest{
		Items: []CheckoutItem{
			{SKU: "120P90", Quantity: 3},
		},
	}

	res, err := service.DoCheckout(req)
	if err != nil {
		t.Fatal(err)
	}

	if res.SubTotal != 149.97 {
		t.Fatalf("expected subtotal 149.97 got %v", res.SubTotal)
	}

	if res.Discount != 49.99 {
		t.Fatalf("expected discount 49.99 got %v", res.Discount)
	}

	if res.Total != 99.98 {
		t.Fatalf("expected total 99.98 got %v", res.Total)
	}
}

func TestDoCheckout_MacBookPro_FreeRaspberryPi(t *testing.T) {
	productRepo := &ProductRepositoryMock{
		items: map[string]product.Product{
			"43N23P": {
				SKU:   "43N23P",
				Name:  "MacBook Pro",
				Price: 5399.99,
				Stock: 5,
			},
			"234234": {
				SKU:   "234234",
				Name:  "Raspberry Pi B",
				Price: 30.00,
				Stock: 2,
			},
		},
	}

	promotionRepo := &PromotionRepositoryMock{
		promos: []promotion.Promotion{
			{
				SKU:  "43N23P",
				Type: promotion.TypeFreeItem,
				Config: mustJSON(promotion.FreeItemConfig{
					ThresholdQty: 1,
					FreeItemSKU:  "234234",
					FreeItemQty:  1,
				}),
			},
		},
	}

	service := NewService(productRepo, promotionRepo, *promotion.NewEngine())

	req := &CheckoutRequest{
		Items: []CheckoutItem{
			{SKU: "43N23P", Quantity: 1},
			{SKU: "234234", Quantity: 1},
		},
	}

	res, err := service.DoCheckout(req)
	if err != nil {
		t.Fatal(err)
	}

	if res.SubTotal != 5429.99 {
		t.Fatalf("expected subtotal 5429.99 got %v", res.SubTotal)
	}

	if res.Discount != 30 {
		t.Fatalf("expected discount 30 got %v", res.Discount)
	}

	if res.Total != 5399.99 {
		t.Fatalf("expected total 5399.99 got %v", res.Total)
	}
}

func TestDoCheckout_AlexaDiscount(t *testing.T) {
	productRepo := &ProductRepositoryMock{
		items: map[string]product.Product{
			"A304SD": {
				SKU:   "A304SD",
				Name:  "Alexa Speaker",
				Price: 109.50,
				Stock: 10,
			},
		},
	}

	promotionRepo := &PromotionRepositoryMock{
		promos: []promotion.Promotion{
			{
				SKU:  "A304SD",
				Type: promotion.TypeDiscount,
				Config: mustJSON(promotion.DiscountConfig{
					ThresholdQty: 3,
					Discount:     0.10,
				}),
			},
		},
	}

	service := NewService(productRepo, promotionRepo, *promotion.NewEngine())

	req := &CheckoutRequest{
		Items: []CheckoutItem{
			{SKU: "A304SD", Quantity: 3},
		},
	}

	res, err := service.DoCheckout(req)
	if err != nil {
		t.Fatal(err)
	}

	if res.SubTotal != 328.50 {
		t.Fatalf("expected subtotal 328.50 got %v", res.SubTotal)
	}

	if res.Discount != 32.85 {
		t.Fatalf("expected discount %v got %v", 32.85, res.Discount)
	}

	if res.Total != 295.65 {
		t.Fatalf("expected total %v got %v", 295.65, res.Total)
	}
}
