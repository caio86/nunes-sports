package service

import (
	"testing"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type MockProductRepo struct {
	products []*domain.Product
}

func NewMockProductRepo() *MockProductRepo {
	return &MockProductRepo{
		products: make([]*domain.Product, 0),
	}
}

func (m *MockProductRepo) Find() ([]*domain.Product, error) {
	result := make([]*domain.Product, 0)

	for _, product := range m.products {
		result = append(result, product)
	}

	return result, nil
}

func (m *MockProductRepo) FindByCode(code string) (*domain.Product, error) {
	for _, value := range m.products {
		if value.Code == code {
			return value, nil
		}
	}

	return nil, ErrProductNotFound
}

func (m *MockProductRepo) Save(product *domain.Product) (*domain.Product, error) {
	m.products = append(m.products, product)
	return product, nil
}

func TestGetProducts(t *testing.T) {
	products := []*domain.Product{
		{Code: "1", Name: "Arroz", Description: "Comida", Price: 6.50},
		{Code: "2", Name: "Carne", Description: "Comida", Price: 16.50},
		{Code: "3", Name: "Pippos", Description: "Comida", Price: 1.99},
		{Code: "4", Name: "Coca-cola", Description: "Bebida", Price: 6.99},
		{Code: "5", Name: "Guarana", Description: "Bebida", Price: 5.99},
	}

	repo := NewMockProductRepo()
	svc := NewProductService(repo)

	repo.products = products

	got, err := svc.GetProducts()

	assertNoError(t, err)

	if len(got) != len(products) {
		t.Errorf("got %d, want %d", len(got), len(products))
	}
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
		repo.products = append(repo.products, test.product)

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
