package promotion

import "checkout/internal/promotion"

type Repository interface {
	GetBySKUs(skus []string) ([]promotion.Promotion, error)
}
