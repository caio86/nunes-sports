package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/caio86/nunes-sports/backend/internal/adapters/input/api/dto"
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

	nextPage := pageNumber + 1
	previousPage := pageNumber - 1
	if previousPage <= 0 {
		previousPage = 1
	}

	links := dto.PaginationLinks{
		Next:     fmt.Sprintf("%s?page=%d&limit=%d", r.URL.Path, nextPage, pageSize),
		Previous: fmt.Sprintf("%s?page=%d&limit=%d", r.URL.Path, previousPage, pageSize),
	}

	res := dto.GetProductsResponse{
		Data:  make([]dto.ProductResponse, len(data)),
		Page:  pageNumber,
		Limit: pageSize,
		Total: total,
		Links: links,
	}
	for i, v := range data {
		res.Data[i].ID = v.ID
		res.Data[i].Name = v.Name
		res.Data[i].Description = v.Description
		res.Data[i].Price = v.Price
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

	if err := renderJSON(w, http.StatusOK, data); err != nil {
		return
	}
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
