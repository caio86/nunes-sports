package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	data := domain.Product{ID: 2}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
