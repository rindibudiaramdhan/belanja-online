package items

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ItemHandler struct {
	service *ItemService
}

func NewItemHandler(s *ItemService) *ItemHandler {
	return &ItemHandler{service: s}
}

func (h *ItemHandler) HandleGetItems(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 {
		limit = 10
	}

	data := h.service.Search(name, page, limit)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": data,
		"meta": map[string]int{
			"page":  page,
			"limit": limit,
		},
	})
}
