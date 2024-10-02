package product

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/caio86/nunes-sports/backend/internal/adapters/input/api/dto"
	"github.com/caio86/nunes-sports/backend/internal/core/domain"
	"github.com/caio86/nunes-sports/backend/internal/core/ports"
	"github.com/caio86/nunes-sports/backend/internal/core/service"
)

type Handler struct {
	svc ports.ProductService
}

func New(service ports.ProductService) *Handler {
	return &Handler{
		svc: service,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	var pageNumber int
	var pageSize int
	var err error

	if res := queryParams.Get("page"); res == "" {
		pageNumber = 1
	} else {
		pageNumber, err = strconv.Atoi(res)
		if err != nil {
			http.Error(w, "Invalid page param", http.StatusBadRequest)
			return
		}
	}

	if res := queryParams.Get("limit"); res == "" {
		pageSize = 0
	} else {
		pageSize, err = strconv.Atoi(res)
		if err != nil {
			http.Error(w, "Invalid limit param", http.StatusBadRequest)
			return
		}
	}

	data, total, err := h.svc.GetProducts(pageNumber, pageSize)
	if err != nil {
		log.Printf("Failed to get products: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := dto.GetProductsResponse{
		Products: make([]dto.ProductResponse, len(data)),
		Page:     pageNumber,
		Limit:    pageSize,
		Total:    total,
	}
	for i, v := range data {
		res.Products[i].ID = v.ID
		res.Products[i].Name = v.Name
		res.Products[i].Description = v.Description
		res.Products[i].Price = v.Price
	}

	if err := renderJSON(w, http.StatusOK, res); err != nil {
		return
	}
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	data, err := h.svc.GetProductByID(id)
	switch err {
	case service.ErrProductNotFound:
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	case nil:
		break
	default:
		log.Printf("Failed to get product: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := dto.GetProductByIDResponse{
		ProductResponse: dto.ProductResponse{
			ID:          data.ID,
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
		},
	}

	if err := renderJSON(w, http.StatusOK, res); err != nil {
		return
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body dto.CreateProductRequest
	if err := decodeJSON(w, r, &body); err != nil {
		log.Println(err)
		return
	}

	domain := &domain.Product{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
	}

	product, err := h.svc.CreateProduct(domain)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := dto.CreateProductResponse{
		ProductResponse: dto.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
	}

	renderJSON(w, http.StatusCreated, res)
}

func renderJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("Failed to send response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}
	return nil
}

func decodeJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dst); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return err
	}
	return nil
}
