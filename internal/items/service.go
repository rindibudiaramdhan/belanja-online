package items

import "strings"

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type ItemService struct {
	Items []Item
}

func NewItemService() *ItemService {
	return &ItemService{
		Items: []Item{
			{ID: 1, Name: "Sabun", Stock: 20},
			{ID: 2, Name: "Shampoo", Stock: 15},
			{ID: 3, Name: "Sikat Gigi", Stock: 30},
		},
	}
}

func (s *ItemService) Search(name string, page, limit int) []Item {
	// simple search
	var result []Item
	for _, item := range s.Items {
		if name == "" || containsIgnoreCase(item.Name, name) {
			result = append(result, item)
		}
	}

	// pagination
	start := (page - 1) * limit
	end := start + limit
	if start > len(result) {
		return []Item{}
	}
	if end > len(result) {
		end = len(result)
	}

	return result[start:end]
}

func containsIgnoreCase(a, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}
