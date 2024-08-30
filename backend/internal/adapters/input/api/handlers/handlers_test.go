package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type MockProductRepositoy struct {
	products map[int]*domain.Product
}

func (s *MockProductRepositoy) FindByID(id int) (*domain.Product, error) {
	product, ok := s.products[id]
	if !ok {
		return nil, fmt.Errorf("product with id %d not found", id)
	}
	return product, nil
}

func TestGETProducts(t *testing.T) {
	store := MockProductRepositoy{
		map[int]*domain.Product{
			1: {ID: 1},
			2: {ID: 2},
		},
	}

	handler := &ProductHandler{&store}

	t.Run("get product with id 1", func(t *testing.T) {
		request := newGetProductRequest(1)
		response := httptest.NewRecorder()

		handler.GetProduct(response, request)

		want := domain.Product{ID: 1}

		assertJSONResponseBody(t, response.Body.String(), want)
	})

	t.Run("get product with id 2", func(t *testing.T) {
		request := newGetProductRequest(2)
		response := httptest.NewRecorder()

		handler.GetProduct(response, request)

		want := domain.Product{ID: 2}

		assertJSONResponseBody(t, response.Body.String(), want)
	})
}

func newGetProductRequest(id int) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/products/%d", id), nil)
	return req
}

func assertJSONResponseBody(t *testing.T, got string, want any) {
	t.Helper()

	wantJSON := new(strings.Builder)
	if err := json.NewEncoder(wantJSON).Encode(want); err != nil {
		t.Errorf("failed to encode %s", err)
	}

	if got != wantJSON.String() {
		t.Errorf("got %s, want %s", got, wantJSON.String())
	}
}
