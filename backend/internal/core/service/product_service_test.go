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

var products = []*domain.Product{
	{Code: "1", Name: "Arroz", Description: "Comida", Price: 6.00},
	{Code: "2", Name: "Carne", Description: "Comida", Price: 16.50},
	{Code: "3", Name: "Pippos", Description: "Comida", Price: 1.99},
	{Code: "4", Name: "Coca-cola", Description: "Bebida", Price: 6.99},
	{Code: "5", Name: "Guarana", Description: "Bebida", Price: 5.99},
}

func TestGetProducts(t *testing.T) {
	testCases := []struct {
		name           string
		page           int
		limit          int
		expectedResult []*domain.Product
	}{
		{"get all five", 0, 5, products[:5]},
		{"get first two", 0, 2, products[:2]},
		{"get second two", 1, 2, products[2:4]},
	}

	svc, _ := setupService(t, products)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := svc.GetProducts(tc.page, tc.limit)
			assertNoError(t, err)

			if !reflect.DeepEqual(got, tc.expectedResult) {
				t.Errorf("got %v, want %v", got, tc.expectedResult)
			}
		})
	}
}

func TestGetProductByCode(t *testing.T) {
	testCases := []struct {
		name        string
		code        string
		expectedErr error
	}{
		{"find product with id 1", "1", nil},
		{"product does not exist", "20", ErrProductNotFound},
	}

	svc, _ := setupService(t, products)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := svc.GetProductByCode(tc.code)

			if tc.expectedErr != nil {
				assertError(t, err, tc.expectedErr)
				assertNil(t, got)
			} else {
				assertNoError(t, err)
				assertNotNil(t, got)
			}
		})
	}
}

func TestCreateProduct(t *testing.T) {
	testCases := []struct {
		name        string
		product     *domain.Product
		expectedErr error
	}{
		{
			"create product", &domain.Product{
				Code:        "10",
				Name:        "Macarrao",
				Description: "Comida",
				Price:       1,
			}, nil,
		},
		{
			"create existing product", &domain.Product{
				Code:        "1",
				Name:        "Arroz-branco",
				Description: "Comida",
				Price:       2,
			}, ErrProductAlreadyExists,
		},
		{"empty product", &domain.Product{}, ErrProductIsEmpty},
		{
			"invalid price", &domain.Product{
				Code:        "20",
				Name:        "Carne",
				Description: "Comida",
				Price:       -2.5,
			}, ErrProductPriceInvalid,
		},
	}

	svc, _ := setupService(t, products)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := svc.CreateProduct(tc.product)

			if tc.expectedErr != nil {
				assertError(t, err, tc.expectedErr)
				assertNil(t, got)
			} else {
				assertNoError(t, err)
				assertNotNil(t, got)
			}
		})
	}
}

// Helper functions

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
		t.Error("expected non-nil, but got nil")
	}
}

func assertNil(t *testing.T, got *domain.Product) {
	t.Helper()
	if got != nil {
		t.Errorf("expected nil, but got %v", got)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("expected error %q, but got %q", want, got)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
