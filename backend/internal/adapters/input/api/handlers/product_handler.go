package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/products/"))
	data := domain.Product{ID: id}

	renderJSON(w, http.StatusOK, data)
}

func renderJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
