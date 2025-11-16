package cart

import "belanja-online/internal/items"

type CartItem struct {
	Item   items.Item `json:"item"`
	Amount int        `json:"amount"`
}

type CartService struct {
	items *[]items.Item
	cart  []CartItem
}

func NewCartService(itemService *items.ItemService) *CartService {
	return &CartService{
		items: &itemService.Items,
		cart:  []CartItem{},
	}
}

func (s *CartService) Add(itemID, amount int) {
	for _, it := range *s.items {
		if it.ID == itemID {
			s.cart = append(s.cart, CartItem{
				Item:   it,
				Amount: amount,
			})
			break
		}
	}
}

func (s *CartService) List() []CartItem {
	return s.cart
}

func (s *CartService) Checkout() {
	s.cart = []CartItem{}
}
