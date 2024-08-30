package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type MockProductRepository struct {
	products map[int]*domain.Product
}

func (s *MockProductRepository) FindAllProducts() ([]*domain.Product, error) {
	products := make([]*domain.Product, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}
	return products, nil
}

func (s *MockProductRepository) FindProductByID(id int) (*domain.Product, error) {
	product, ok := s.products[id]
	if !ok {
		return nil, fmt.Errorf("product with id %d not found", id)
	}
	return product, nil
}

func TestGETProducts(t *testing.T) {
	store := MockProductRepository{
		map[int]*domain.Product{
			1: {ID: 1},
			2: {ID: 2},
		},
	}

	handler := &ProductHandler{&store}

	t.Run("get all products", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/products", nil)
		response := httptest.NewRecorder()

		handler.GetAllProducts(response, request)

		want := []*domain.Product{
			{ID: 1},
			{ID: 2},
		}

		got := getProductSliceFromResponse(t, response.Body)
		assertProducts(t, got, want)
	})

	t.Run("get product with id 1", func(t *testing.T) {
		request := newGetProductRequest(1)
		response := httptest.NewRecorder()

		handler.GetProduct(response, request)

		got := getProductFromResponse(t, response.Body)
		assertProductID(t, got.ID, 1)
	})

	t.Run("get product with id 2", func(t *testing.T) {
		request := newGetProductRequest(2)
		response := httptest.NewRecorder()

		handler.GetProduct(response, request)

		got := getProductFromResponse(t, response.Body)
		assertProductID(t, got.ID, 2)
	})
}

func newGetProductRequest(id int) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/products/%d", id), nil)
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
