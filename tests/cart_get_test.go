package tests

import (
	"belanja-online/internal/cart"
	"belanja-online/internal/items"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestGetCart(t *testing.T) {
	itemService := items.NewItemService()
	cartService := cart.NewCartService(itemService)
	cartService.Add(1, 2)

	cartHandler := cart.NewCartHandler(cartService)

	r := chi.NewRouter()
	r.Get("/cart", cartHandler.HandleGetCart)

	req := httptest.NewRequest("GET", "/cart", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	if !contains(w.Body.String(), `"item":`) {
		t.Errorf("Expected cart content in response")
	}
}
