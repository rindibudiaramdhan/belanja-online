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

func TestAddToCart(t *testing.T) {
	itemService := items.NewItemService()
	cartService := cart.NewCartService(itemService)
	cartHandler := cart.NewCartHandler(cartService)

	r := chi.NewRouter()
	r.Post("/cart", cartHandler.HandleAddToCart)

	body := []byte(`{"item_id":1,"amount":2}`)
	req := httptest.NewRequest("POST", "/cart", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	if len(cartService.List()) != 1 {
		t.Errorf("Expected 1 item in cart, got %d", len(cartService.List()))
	}
}
