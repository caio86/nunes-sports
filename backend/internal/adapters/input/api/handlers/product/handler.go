package product

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/caio86/nunes-sports/backend/internal/core/ports"
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
	var pageIndex int
	var pageSize int
	var err error

	if res := queryParams.Get("page"); res == "" {
		pageIndex = 0
	} else {
		pageIndex, err = strconv.Atoi(res)
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

	data, _, err := h.svc.GetProducts(pageIndex, pageSize)
	if err != nil {
		log.Printf("Failed to get products %v", err)
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
		log.Printf("Failed to send response %v", err)
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
