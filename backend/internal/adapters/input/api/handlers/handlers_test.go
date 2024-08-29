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

func TestGETProducts(t *testing.T) {
	t.Run("get product with id 1", func(t *testing.T) {
		request := newGetProductRequest(1)
		response := httptest.NewRecorder()

		GetProduct(response, request)

		want := domain.Product{ID: 1}

		assertJSONResponseBody(t, response.Body.String(), want)
	})

	t.Run("get product with id 2", func(t *testing.T) {
		request := newGetProductRequest(2)
		response := httptest.NewRecorder()

		GetProduct(response, request)

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
