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
	CreateProduct(product *domain.Product) error
}

type AddProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductHandler struct {
	store ProductRepository
	http.Handler
}

func NewProductHandler(store ProductRepository) *ProductHandler {
	p := new(ProductHandler)

	p.store = store

	router := http.NewServeMux()

	router.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			p.GetAllProducts(w, r)
		case http.MethodPost:
			p.CreateProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/api/v1/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			p.GetProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	p.Handler = router

	return p
}

func (p *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	data := p.store.FindAllProducts()

	renderJSON(w, http.StatusOK, data)
}

func (p *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/products/"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := p.store.FindProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	renderJSON(w, http.StatusOK, data)
}

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req AddProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	err := p.store.CreateProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func renderJSON(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
