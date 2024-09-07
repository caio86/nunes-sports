package service

import (
	"testing"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type MockProductRepo struct {
	FindAllfunc func() ([]*domain.Product, error)
}

func (m *MockProductRepo) FindAll() ([]*domain.Product, error)

func TestCreateProduct(t *testing.T) {
	svc := NewProductService(&MockProductRepo{})

	tests := []struct {
		name        string
		product     *domain.Product
		expectedErr error
	}{
		{
			name: "create product",
			product: &domain.Product{
				Code:        "1",
				Name:        "Arroz",
				Description: "Comida",
				Price:       1,
			},
			expectedErr: nil,
		},
		{
			name:        "empty product",
			product:     &domain.Product{},
			expectedErr: ErrProductNameRequired,
		},
		{
			name: "invalid price",
			product: &domain.Product{
				Code:        "2",
				Name:        "Carne",
				Description: "Comida",
				Price:       -2.5,
			},
			expectedErr: ErrProductPriceInvalid,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := svc.CreateProduct(test.product)

			if test.expectedErr != nil {
				assertError(t, err, test.expectedErr)
				assertProduct(t, got, nil)
			} else {
				assertNoError(t, err)
				assertProduct(t, got, test.product)
			}
		})
	}
}

// Helpers

func assertProduct(t *testing.T, got, want *domain.Product) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("got %q, when did not expect an error", err)
	}
}
