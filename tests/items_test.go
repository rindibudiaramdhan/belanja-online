package tests

import (
	"belanja-online/internal/items"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestGetItems(t *testing.T) {
	// Arrange
	itemService := items.NewItemService()
	itemHandler := items.NewItemHandler(itemService)

	r := chi.NewRouter()
	r.Get("/items", itemHandler.HandleGetItems)

	req := httptest.NewRequest("GET", "/items?name=Sa&page=1&limit=10", nil)
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	body := w.Body.String()
	if !contains(body, "Sabun") {
		t.Errorf("Expected to find Sabun in result, got %s", body)
	}
}

func contains(s, substr string) bool {
	return (len(s) >= len(substr)) && (string(s[:len(substr)]) == substr || contains(s[1:], substr))
}
