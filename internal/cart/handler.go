package cart

import (
	"encoding/json"
	"net/http"
)

type CartHandler struct {
	service *CartService
}

func NewCartHandler(s *CartService) *CartHandler {
	return &CartHandler{service: s}
}

func (h *CartHandler) HandleAddToCart(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ItemID int `json:"item_id"`
		Amount int `json:"amount"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	h.service.Add(body.ItemID, body.Amount)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "added to cart",
	})
}

func (h *CartHandler) HandleGetCart(w http.ResponseWriter, r *http.Request) {
	data := h.service.List()
	json.NewEncoder(w).Encode(data)
}

func (h *CartHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	h.service.Checkout()
	json.NewEncoder(w).Encode(map[string]string{
		"message": "checkout success",
	})
}
