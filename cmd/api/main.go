package main

import (
	"log"
	"net/http"

	"checkout/internal/checkout"
	checkoutHandler "checkout/internal/checkout/handler"
	db "checkout/internal/repository/db"

	productRepository "checkout/internal/repository/db/product"
	promotionRepository "checkout/internal/repository/db/promotion"

	"checkout/internal/promotion"
)

func main() {
	mysqlDB, err := db.NewMySQLConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer mysqlDB.Close()

	productRepo := productRepository.NewProductRepository(
		mysqlDB,
	)

	promotionRepo := promotionRepository.NewPromotionRepository(
		mysqlDB,
	)

	promotionEngine := promotion.NewEngine()

	checkoutService := checkout.NewService(
		productRepo,
		promotionRepo,
		*promotionEngine,
	)

	checkoutHandler := checkoutHandler.NewHandler(
		checkoutService,
	)

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/checkout",
		checkoutHandler.Checkout,
	)

	log.Println("server running on :8080")

	err = http.ListenAndServe(
		":8080",
		mux,
	)
	if err != nil {
		log.Fatal(err)
	}
}
