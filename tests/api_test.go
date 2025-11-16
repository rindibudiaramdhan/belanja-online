package tests

import (
	"belanja-online/internal/cart"
	"belanja-online/internal/items"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()

	itemService := items.NewItemService()
	cartService := cart.NewCartService(itemService)

	itemHandler := items.NewItemHandler(itemService)
	cartHandler := cart.NewCartHandler(cartService)

	r.Get("/items", itemHandler.HandleGetItems)
	r.Get("/cart", cartHandler.HandleGetCart)
	r.Post("/cart", cartHandler.HandleAddToCart)
	r.Post("/cart/checkout", cartHandler.HandleCheckout)

	return r
}

// ------------------------
// TEST GET /items
// ------------------------

func TestGetItems(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest("GET", "/items?name=sa&page=1&limit=10", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.NotNil(t, resp["items"])
}

// ------------------------
// TEST POST /cart
// ------------------------

func TestAddToCart(t *testing.T) {
	r := setupRouter()

	body := `{"item_id": 1, "amount": 2}`
	req := httptest.NewRequest("POST", "/cart", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "added to cart", resp["message"])
}

// ------------------------
// TEST GET /cart
// ------------------------

func TestGetCart(t *testing.T) {
	r := setupRouter()

	// isi keranjang dulu
	body := `{"item_id": 2, "amount": 1}`
	reqAdd := httptest.NewRequest("POST", "/cart", strings.NewReader(body))
	reqAdd.Header.Set("Content-Type", "application/json")
	wAdd := httptest.NewRecorder()
	r.ServeHTTP(wAdd, reqAdd)

	// GET cart
	req := httptest.NewRequest("GET", "/cart", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.GreaterOrEqual(t, len(resp), 1)
	assert.Equal(t, 2, int(resp[0]["item"].(map[string]interface{})["id"].(float64)))
}

// ------------------------
// TEST POST /cart/checkout
// ------------------------

func TestCheckout(t *testing.T) {
	r := setupRouter()

	// Tambah data ke keranjang dulu
	body := `{"item_id": 3, "amount": 5}`
	reqAdd := httptest.NewRequest("POST", "/cart", strings.NewReader(body))
	reqAdd.Header.Set("Content-Type", "application/json")
	wAdd := httptest.NewRecorder()
	r.ServeHTTP(wAdd, reqAdd)

	// Checkout
	req := httptest.NewRequest("POST", "/cart/checkout", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "checkout success", resp["message"])

	// pastikan keranjang kosong
	reqGet := httptest.NewRequest("GET", "/cart", nil)
	wGet := httptest.NewRecorder()
	r.ServeHTTP(wGet, reqGet)

	var cartData []interface{}
	json.Unmarshal(wGet.Body.Bytes(), &cartData)
	assert.Equal(t, 0, len(cartData))
}
