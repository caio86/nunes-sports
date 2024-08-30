package main

import (
	"log"
	"net/http"

	"github.com/caio86/nunes-sports/backend/internal/adapters/input/api/handlers"
	"github.com/caio86/nunes-sports/backend/internal/adapters/output/db"
)

func main() {
	store := db.NewInMemoryProductDB()
	handler := handlers.NewProductHandler(store)

	log.Fatal(http.ListenAndServe(":5000", handler))
}
