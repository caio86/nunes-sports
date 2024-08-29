package main

import (
	"log"
	"net/http"

	"github.com/caio86/nunes-sports/backend/internal/adapters/input/api/handlers"
)

func main() {
	mux := http.NewServeMux()
	handler := handlers.GetProduct

	mux.HandleFunc("GET /api/v1/products", handler)

	log.Fatal(http.ListenAndServe(":5000", mux))
}
