package main

import (
	"log"
	"net/http"

	"github.com/caio86/nunes-sports/backend/internal/adapters/input/api/handlers"
	"github.com/caio86/nunes-sports/backend/internal/adapters/output/db"
	"github.com/caio86/nunes-sports/backend/internal/core/service"
)

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func main() {
	store := db.NewInMemoryProductDB()
	service := service.NewProductServiceImpl(store)
	handler := handlers.NewProductHandler(service)

	log.Fatal(http.ListenAndServe(":5000", CORS(handler.ServeHTTP)))
}
