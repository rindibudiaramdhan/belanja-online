package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"belanja-online/internal/cart"
	"belanja-online/internal/db"
	"belanja-online/internal/items"
)

func main() {
	godotenv.Load()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	itemRepo := items.NewItemRepository(dbConn)
	itemService := items.NewItemService(itemRepo)
	itemHandler := items.NewItemHandler(itemService)

	cartRepo := cart.NewCartRepository(dbConn)
	cartService := cart.NewCartService(cartRepo)
	cartHandler := cart.NewCartHandler(cartService)

	// routes
	r.Get("/items", itemHandler.HandleGetItems)
	r.Get("/cart", cartHandler.HandleGetCart)
	r.Post("/cart", cartHandler.HandleAddToCart)
	r.Post("/cart/checkout", cartHandler.HandleCheckout)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
