package service

import (
	"reflect"
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

func (m *MockProductRepo) Find(offset, limit int) ([]*domain.Product, error) {
	result := make([]*domain.Product, 0)

	start := offset * limit
	end := start + limit

	result = m.products[start:end]

	return result, nil
}

func (m *MockProductRepo) FindByID(code uint) (*domain.Product, error) {
	for _, value := range m.products {
		if value.ID == code {
			return value, nil
		}
	}

	return nil, ErrProductNotFound
}

func (m *MockProductRepo) Save(product *domain.Product) (*domain.Product, error) {
	m.products = append(m.products, product)
	return product, nil
}

var products = []*domain.Product{
	{ID: 1, Name: "Arroz", Description: "Comida", Price: 6.00},
	{ID: 2, Name: "Carne", Description: "Comida", Price: 16.50},
	{ID: 3, Name: "Pippos", Description: "Comida", Price: 1.99},
	{ID: 4, Name: "Coca-cola", Description: "Bebida", Price: 6.99},
	{ID: 5, Name: "Guarana", Description: "Bebida", Price: 5.99},
}

func TestGetProducts(t *testing.T) {
	tests := []struct {
		name           string
		page           int
		limit          int
		expectedResult []*domain.Product
	}{
		{name: "get all five", page: 0, limit: 5, expectedResult: products[:5]},
		{name: "get first two", page: 0, limit: 2, expectedResult: products[:2]},
		{name: "get second two", page: 1, limit: 2, expectedResult: products[2:4]},
	}

	svc, _ := setupService(t, products)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := svc.GetProducts(test.page, test.limit)

			assertNoError(t, err)

			if !reflect.DeepEqual(got, test.expectedResult) {
				t.Errorf("got %v, want %v", got, test.expectedResult)
			}
		})
	}
}

func TestGetProductByID(t *testing.T) {
	tests := []struct {
		name        string
		code        uint
		expectedErr error
	}{
		{
			name:        "find product with code 1",
			code:        1,
			expectedErr: nil,
		},
		{
			name:        "product does not exist",
			code:        20,
			expectedErr: ErrProductNotFound,
		},
	}

	svc, _ := setupService(t, products)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := svc.GetProductByID(test.code)

			if test.expectedErr != nil {
				assertError(t, err, test.expectedErr)
				assertNil(t, got)
			} else {
				assertNoError(t, err)
				assertNotNil(t, got)
			}
		})
	}
}

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		name        string
		product     *domain.Product
		expectedErr error
	}{
		{
			name: "create product",
			product: &domain.Product{
				ID:          10,
				Name:        "Macarrao",
				Description: "Comida",
				Price:       1,
			},
			expectedErr: nil,
		},
		{
			name: "create existing product",
			product: &domain.Product{
				ID:          1,
				Name:        "Arroz-branco",
				Description: "Comida",
				Price:       2,
			},
			expectedErr: ErrProductAlreadyExists,
		},
		{
			name:        "empty product",
			product:     &domain.Product{},
			expectedErr: ErrProductIsEmpty,
		},
		{
			name: "invalid price",
			product: &domain.Product{
				ID:          20,
				Name:        "Carne",
				Description: "Comida",
				Price:       -2.5,
			},
			expectedErr: ErrProductPriceInvalid,
		},
	}

	svc, _ := setupService(t, products)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := svc.CreateProduct(test.product)

			if test.expectedErr != nil {
				assertError(t, err, test.expectedErr)
				assertNil(t, got)
			} else {
				assertNoError(t, err)
				assertNotNil(t, got)
			}
		})
	}
}

// Helpers

func setupService(t *testing.T, products []*domain.Product) (*ProductService, *MockProductRepo) {
	t.Helper()

	repo := NewMockProductRepo()
	svc := NewProductService(repo)

	repo.products = products

	return svc, repo
}

func assertNotNil(t *testing.T, got *domain.Product) {
	t.Helper()

	if got == nil {
		t.Error("got nil, when did not expect nil")
	}
}

func assertNil(t *testing.T, got *domain.Product) {
	t.Helper()

	if got != nil {
		t.Errorf("got %v, want nil", got)
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
