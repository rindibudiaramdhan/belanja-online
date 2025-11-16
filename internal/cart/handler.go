package cart

import (
	"encoding/json"
	"net/http"
)

type CartHandler struct {
	service CartServiceI
}

func NewCartHandler(s CartServiceI) *CartHandler {
	return &CartHandler{service: s}
}

func (h *CartHandler) HandleAddToCart(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ItemID int `json:"item_id"`
		Amount int `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", 400)
		return
	}

	if err := h.service.Add(body.ItemID, body.Amount); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "added to cart"})
}

func (h *CartHandler) HandleGetCart(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.List()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Jika data == nil, jadikan slice kosong supaya JSON -> []
	if data == nil {
		data = []CartItem{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *CartHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	if err := h.service.Checkout(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "checkout success"})
}
