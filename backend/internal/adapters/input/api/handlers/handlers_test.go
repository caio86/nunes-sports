package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

func TestGETProducts(t *testing.T) {
	t.Run("get product with id 1", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/products/1", nil)
		response := httptest.NewRecorder()

		GetProduct(response, request)

		got := response.Body.String()
		want := new(strings.Builder)
		json.NewEncoder(want).Encode(domain.Product{ID: 1})

		if got != want.String() {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("get product with id 2", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/products/2", nil)
		response := httptest.NewRecorder()

		GetProduct(response, request)

		got := response.Body.String()
		want := new(strings.Builder)
		json.NewEncoder(want).Encode(domain.Product{ID: 2})

		if got != want.String() {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
