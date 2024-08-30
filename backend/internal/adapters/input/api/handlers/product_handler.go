package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type ProductRepository interface {
	FindAllProducts() []*domain.Product
	FindProductByID(id int) (*domain.Product, error)
}

type ProductHandler struct {
	store ProductRepository
}

func NewProductHandler(store ProductRepository) *ProductHandler {
	return &ProductHandler{
		store: store,
	}
}

func (p *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	data := p.store.FindAllProducts()

	renderJSON(w, http.StatusOK, data)
}

func (p *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/products/"))
	data, err := p.store.FindProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(w, http.StatusOK, data)
}

func renderJSON(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
