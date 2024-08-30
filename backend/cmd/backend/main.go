package main

import (
	"log"
	"net/http"

	"github.com/caio86/nunes-sports/backend/internal/adapters/input/api/handlers"
	"github.com/caio86/nunes-sports/backend/internal/adapters/output/db"
)

func main() {
	mux := http.NewServeMux()
	store := &db.InMemoryProductDB{}
	handler := handlers.NewProductHandler(store)

	mux.HandleFunc("/api/v1/products", handler.GetAllProducts)
	mux.HandleFunc("/api/v1/products/", handler.GetProduct)

	log.Fatal(http.ListenAndServe(":5000", mux))
}
