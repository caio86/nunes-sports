package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
	"github.com/caio86/nunes-sports/backend/internal/core/service"
	"github.com/caio86/nunes-sports/backend/internal/mocks"
	"github.com/stretchr/testify/assert"
)

var products = []*domain.Product{
	{ID: "1", Name: "Arroz", Description: "Branco", Price: 6.49},
	{ID: "2", Name: "Carne", Description: "Cupim", Price: 49.99},
	{ID: "3", Name: "Feijão", Description: "Preto", Price: 6.49},
	{ID: "4", Name: "Café", Description: "Grãos", Price: 20.49},
}

func TestGet(t *testing.T) {
	svc := mocks.NewProductService()
	handler := New(svc)

	router := http.NewServeMux()
	router.HandleFunc("GET /product", handler.Get)

	svc.On("GetProducts", 1, 2).
		Return(products[0:2], int64(len(products)), nil)

	svc.On("GetProducts", 2, 2).
		Return(products[2:4], int64(len(products)), nil)

	svc.On("GetProducts", 0, 0).
		Return(products, int64(len(products)), nil)

	t.Run("get first page", func(t *testing.T) {
		req := newGetRequest(1, 2)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		var got []*domain.Product
		json.NewDecoder(res.Body).Decode(&got)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, products[0:2], got)
	})

	t.Run("get second page", func(t *testing.T) {
		req := newGetRequest(2, 2)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		var got []*domain.Product
		json.NewDecoder(res.Body).Decode(&got)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, products[2:4], got)
	})

	t.Run("get all", func(t *testing.T) {
		req := newGetRequest(0, 0)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		var got []*domain.Product
		json.NewDecoder(res.Body).Decode(&got)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, products, got)
	})

	t.Run("invalid page parameter", func(t *testing.T) {
		req := newGetRequest("a", 0)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		got := res.Body.String()

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Invalid page param\n", got)
	})

	t.Run("invalid limit parameter", func(t *testing.T) {
		req := newGetRequest(0, "a")
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		got := res.Body.String()

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Invalid limit param\n", got)
	})
}

func TestGetByID(t *testing.T) {
	svc := mocks.NewProductService()
	handler := New(svc)

	router := http.NewServeMux()
	router.HandleFunc("GET /product/{id}", handler.GetByID)

	svc.On("GetProductByID", "1").
		Return(products[0], nil)

	svc.On("GetProductByID", "5").
		Return(&domain.Product{}, service.ErrProductNotFound)

	svc.On("GetProductByID", "0").
		Return(&domain.Product{}, errors.New("Custom Error"))

	t.Run("get product with id 1", func(t *testing.T) {
		req := newGetByIDRequest(1)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		var got *domain.Product
		json.NewDecoder(res.Body).Decode(&got)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, products[0], got)
	})

	t.Run("return 404 for product not found", func(t *testing.T) {
		req := newGetByIDRequest(5)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		got := res.Body.String()

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "Product Not Found\n", got)
	})

	t.Run("custom error", func(t *testing.T) {
		req := newGetByIDRequest(0)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		got := res.Body.String()

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "Internal Server Error\n", got)
	})
}

func newGetRequest(page, limit interface{}) *http.Request {
	url := fmt.Sprintf("/product?page=%v&limit=%v", page, limit)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	return req
}

func newGetByIDRequest(id any) *http.Request {
	url := fmt.Sprintf("/product/%v", id)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	return req
}
