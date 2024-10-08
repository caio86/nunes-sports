package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type MockProductService struct {
	store MockProductRepository
}

func (m *MockProductService) GetProducts() []*domain.Product {
	return m.store.FindAll()
}

func (m *MockProductService) GetProductByID(id int) (*domain.Product, error) {
	return m.store.FindByID(id)
}

func (m *MockProductService) CreateProduct(product *domain.Product) error {
	return m.store.Create(product)
}

func (m *MockProductService) UpdateProduct(product *domain.Product) error {
	return m.store.Update(product)
}

func (m *MockProductService) DeleteProduct(id int) error {
	return m.store.Delete(id)
}

type MockProductRepository struct {
	products map[int]*domain.Product
	lastID   int
}

func (s *MockProductRepository) FindAll() []*domain.Product {
	products := make([]*domain.Product, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}
	return products
}

func (s *MockProductRepository) FindByID(id int) (*domain.Product, error) {
	product, ok := s.products[id]
	if !ok {
		return nil, errors.New("could not find the requested product")
	}
	return product, nil
}

func (s *MockProductRepository) Create(product *domain.Product) error {
	s.lastID++
	product.ID = s.lastID

	s.products[product.ID] = product
	return nil
}

func (s *MockProductRepository) Update(product *domain.Product) error {
	s.products[product.ID] = product
	return nil
}

func (s *MockProductRepository) Delete(id int) error {
	delete(s.products, id)
	return nil
}

func TestGETProducts(t *testing.T) {
	store := MockProductRepository{
		map[int]*domain.Product{
			1: {ID: 1},
			2: {ID: 2},
		},
		2,
	}

	service := &MockProductService{store}
	handler := NewProductHandler(service)

	t.Run("get all products", func(t *testing.T) {
		request := newGetProductsRequest()
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		want := []*domain.Product{
			{ID: 1},
			{ID: 2},
		}

		got := getProductSliceFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertProducts(t, got, want)
	})

	t.Run("get product with id 1", func(t *testing.T) {
		request := newGetProductRequest(1)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		got := getProductFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertProductID(t, got.ID, 1)
	})

	t.Run("get product with id 2", func(t *testing.T) {
		request := newGetProductRequest(2)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		got := getProductFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertProductID(t, got.ID, 2)
	})

	t.Run("return 404 for missing product", func(t *testing.T) {
		request := newGetProductRequest(3)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestCreateProduct(t *testing.T) {
	store := MockProductRepository{
		map[int]*domain.Product{},
		0,
	}

	service := &MockProductService{store}
	handler := NewProductHandler(service)

	t.Run("create product", func(t *testing.T) {
		product := domain.Product{
			Name: "arroz",
		}

		json, _ := json.Marshal(product)

		request := newPostProductRequest(bytes.NewBuffer(json))
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusCreated)
		if store.products[1].Name != product.Name {
			t.Errorf("did not create product correctly, got %v, want %v", store.products[1], product)
		}
	})
}

func TestUpdateProduct(t *testing.T) {
	store := MockProductRepository{
		map[int]*domain.Product{
			1: {ID: 1, Name: "arroz"},
		},
		1,
	}

	service := &MockProductService{store}
	handler := NewProductHandler(service)

	t.Run("update product", func(t *testing.T) {
		updatedProduct := &domain.Product{
			Description: "arroz branco",
		}

		json, _ := json.Marshal(updatedProduct)

		request := newPutProductRequest(1, bytes.NewBuffer(json))
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		if store.products[1].Description != updatedProduct.Description {
			t.Errorf("did not update product correctly, got %v, want %v", store.products[1], updatedProduct)
		}
	})
}

func TestDeleteProduct(t *testing.T) {
	store := MockProductRepository{
		map[int]*domain.Product{
			1: {ID: 1},
		},
		1,
	}

	service := &MockProductService{store}
	handler := NewProductHandler(service)

	t.Run("delete product", func(t *testing.T) {
		request := newDeleteProductRequest(1)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNoContent)
		assertProducts(t, store.FindAll(), []*domain.Product{})
	})
}

func newGetProductsRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/products", nil)
	return req
}

func newGetProductRequest(id int) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/products/%d", id), nil)
	return req
}

func newPostProductRequest(body io.Reader) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/products", body)
	return req
}

func newPutProductRequest(id int, body io.Reader) *http.Request {
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/products/%d", id), body)
	return req
}

func newDeleteProductRequest(id int) *http.Request {
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/products/%d", id), nil)
	return req
}

func getProductSliceFromResponse(t *testing.T, body io.Reader) (products []*domain.Product) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&products)
	if err != nil {
		t.Fatalf("failed to decode response %q into a slice of products, '%v'", body, err)
	}

	return
}

func getProductFromResponse(t *testing.T, body io.Reader) (product *domain.Product) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&product)
	if err != nil {
		t.Fatalf("failed to decode response %q a product, '%v'", body, err)
	}

	return
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("status code mismatch, got %d, want %d", got, want)
	}
}

func assertProducts(t *testing.T, got, want []*domain.Product) {
	t.Helper()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertProductID(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
