package tests

import (
	"belanja-online/internal/cart"
	"belanja-online/internal/items"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestCheckout(t *testing.T) {
	itemService := items.NewItemService()
	cartService := cart.NewCartService(itemService)
	cartService.Add(1, 3)

	cartHandler := cart.NewCartHandler(cartService)

	r := chi.NewRouter()
	r.Post("/cart/checkout", cartHandler.HandleCheckout)

	req := httptest.NewRequest("POST", "/cart/checkout", bytes.NewBuffer([]byte("{}")))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d", w.Code)
	}

	if len(cartService.List()) != 0 {
		t.Errorf("Cart should be empty after checkout")
	}
}
