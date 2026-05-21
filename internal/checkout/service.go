package checkout

import (
	"errors"

	productRepository "checkout/internal/product/repository"
	"checkout/internal/promotion"
	promotionRepository "checkout/internal/promotion/repository"
)

var (
	ErrProductNotFound   = errors.New("product not found")
	ErrInsufficientStock = errors.New("insufficient stock")
)

type Service struct {
	productRepo     productRepository.Repository
	promotionRepo   promotionRepository.Repository
	promotionEngine promotion.Engine
}

func NewService(
	productRepo productRepository.Repository,
	promotionRepo promotionRepository.Repository,
	promotionEngine promotion.Engine,
) *Service {
	return &Service{
		productRepo:     productRepo,
		promotionRepo:   promotionRepo,
		promotionEngine: promotionEngine,
	}
}
