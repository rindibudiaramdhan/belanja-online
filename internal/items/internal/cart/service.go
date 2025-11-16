package cart

import "belanja-online/internal/items"

type CartItem struct {
	ItemID int `json:"item_id"`
	Qty    int `json:"qty"`
}

type CartService struct {
	cart        []CartItem
	itemService *items.ItemService
}

func NewCartService(itemService *items.ItemService) *CartService {
	return &CartService{
		cart:        []CartItem{},
		itemService: itemService,
	}
}

func (s *CartService) AddToCart(itemID, qty int) {
	s.cart = append(s.cart, CartItem{ItemID: itemID, Qty: qty})
}

func (s *CartService) GetCart() []CartItem {
	return s.cart
}

func (s *CartService) Checkout() {
	s.cart = []CartItem{} // kosongkan keranjang
}
