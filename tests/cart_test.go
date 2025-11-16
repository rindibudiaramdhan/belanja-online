package tests

import (
	"belanja-online/internal/cart"
	"belanja-online/internal/cart/mocks"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestHandleAddToCart(t *testing.T) {
	// Arrange
	mockSvc := new(mocks.CartServiceI)
	mockSvc.On("Add", 1, 2).Return(nil)

	handler := cart.NewCartHandler(mockSvc)

	r := chi.NewRouter()
	r.Post("/cart", handler.HandleAddToCart)

	body := `{"item_id":1,"amount":2}`
	req := httptest.NewRequest("POST", "/cart", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, 200, w.Code)

	var resp map[string]string
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err)

	assert.Equal(t, "added to cart", resp["message"])

	// Verify mock was called
	mockSvc.AssertCalled(t, "Add", 1, 2)
}
