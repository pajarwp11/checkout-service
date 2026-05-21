package checkout

import (
	"checkout/internal/checkout"
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	service *checkout.Service
}

func NewHandler(service *checkout.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request checkout.CheckoutRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			map[string]string{
				"message": "invalid request body",
			},
		)

		return
	}

	response, err := h.service.DoCheckout(&request)
	if err != nil {
		switch {
		case errors.Is(err, checkout.ErrProductNotFound):
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(
				map[string]string{
					"message": err.Error(),
				},
			)
			return

		case errors.Is(err, checkout.ErrInsufficientStock):
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(
				map[string]string{
					"message": err.Error(),
				},
			)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(
			map[string]string{
				"message": "internal server error",
			},
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)
}
