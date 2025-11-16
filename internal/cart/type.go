package cart

import "belanja-online/internal/items"

type CartItem struct {
	ID     int        `json:"id"`
	Item   items.Item `json:"item"`
	Amount int        `json:"amount"`
}
