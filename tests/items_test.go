package tests

import (
	"belanja-online/internal/items"
	"belanja-online/internal/items/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestGetItems(t *testing.T) {
	// Arrange: mock repository
	mockRepo := new(mocks.ItemRepositoryI)
	mockRepo.
		On("Find", "Sa", 10, 0).
		Return([]items.Item{
			{ID: 1, Name: "Sabun", Stock: 20},
		}, nil)

	// service memakai mock repo
	itemService := items.NewItemService(mockRepo)
	itemHandler := items.NewItemHandler(itemService)

	r := chi.NewRouter()
	r.Get("/items", itemHandler.HandleGetItems)

	req := httptest.NewRequest("GET", "/items?name=Sa&page=1&limit=10", nil)
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Sabun")

	// verify mock dipanggil
	mockRepo.AssertExpectations(t)
}

func TestGetItems_ErrorFromRepo(t *testing.T) {
	mockRepo := new(mocks.ItemRepositoryI)
	mockRepo.
		On("Find", "Sa", 10, 0).
		Return([]items.Item{}, errors.New("db error"))

	itemService := items.NewItemService(mockRepo)
	itemHandler := items.NewItemHandler(itemService)

	r := chi.NewRouter()
	r.Get("/items", itemHandler.HandleGetItems)

	req := httptest.NewRequest("GET", "/items?name=Sa&page=1&limit=10", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "db error")

	mockRepo.AssertExpectations(t)
}
