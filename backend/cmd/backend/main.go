package main

import (
	"log"
	"net/http"

	"github.com/caio86/nunes-sports/backend/internal/adapters/input/api/handlers"
	"github.com/caio86/nunes-sports/backend/internal/adapters/output/db"
	"github.com/caio86/nunes-sports/backend/internal/core/service"
)

func main() {
	store := db.NewInMemoryProductDB()
	service := service.NewProductServiceImpl(store)
	handler := handlers.NewProductHandler(service)

	log.Fatal(http.ListenAndServe(":5000", handler))
}
