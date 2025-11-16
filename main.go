package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"belanja-online/internal/cart"
	"belanja-online/internal/items"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	itemService := items.NewItemService()
	cartService := cart.NewCartService(itemService)

	itemHandler := items.NewItemHandler(itemService)
	cartHandler := cart.NewCartHandler(cartService)

	r.Get("/items", itemHandler.HandleGetItems)
	r.Get("/cart", cartHandler.HandleGetCart)
	r.Post("/cart", cartHandler.HandleAddToCart)
	r.Post("/cart/checkout", cartHandler.HandleCheckout)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
