package service

import (
	"testing"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type MockProductRepo struct {
	products map[string]*domain.Product
}

func NewMockProductRepo() *MockProductRepo {
	return &MockProductRepo{
		products: make(map[string]*domain.Product),
	}
}

func (m *MockProductRepo) FindByCode(code string) (*domain.Product, error) {
	product, ok := m.products[code]
	if !ok {
		return nil, ErrProductNotFound
	}

	return product, nil
}

func (m *MockProductRepo) Save(product *domain.Product) (*domain.Product, error) {
	m.products[product.Code] = product
	return product, nil
}

func TestGetProductByCode(t *testing.T) {
	tests := []struct {
		name        string
		code        string
		product     *domain.Product
		expectedErr error
	}{
		{
			name: "find product with code 1",
			code: "1",
			product: &domain.Product{
				Code:        "1",
				Name:        "Arroz",
				Description: "Comida",
				Price:       6,
			},
			expectedErr: nil,
		},
		{
			name:        "product does not exist",
			code:        "2",
			product:     &domain.Product{},
			expectedErr: ErrProductNotFound,
		},
		{
			name:        "invalid code",
			code:        "-1",
			product:     &domain.Product{},
			expectedErr: ErrProductCodeInvalid,
		},
		{
			name:        "invalid code with letters",
			code:        "abc",
			product:     &domain.Product{},
			expectedErr: ErrProductCodeInvalid,
		},
	}

	repo := NewMockProductRepo()
	svc := NewProductService(repo)

	for _, test := range tests {
		repo.products[test.product.Code] = test.product

		t.Run(test.name, func(t *testing.T) {
			got, err := svc.GetProductByCode(test.code)

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

func TestCreateProduct(t *testing.T) {
	repo := NewMockProductRepo()
	svc := NewProductService(repo)

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
			name: "create existing product",
			product: &domain.Product{
				Code:        "1",
				Name:        "Arroz-branco",
				Description: "Comida",
				Price:       2,
			},
			expectedErr: ErrProductAlreadyExists,
		},
		{
			name:        "empty product",
			product:     &domain.Product{},
			expectedErr: ErrProductCodeInvalid,
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