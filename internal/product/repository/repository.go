package product

import "checkout/internal/product"

type Repository interface {
	GetBySKUs(skus []string) ([]product.Product, error)
}
