package service

import (
	"reflect"
	"testing"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
	"github.com/caio86/nunes-sports/backend/internal/mocks"
)

var products = []*domain.Product{
	{ID: "1", Name: "Arroz", Description: "Comida", Price: 6.00},
	{ID: "2", Name: "Carne", Description: "Comida", Price: 16.50},
	{ID: "3", Name: "Pippos", Description: "Comida", Price: 1.99},
	{ID: "4", Name: "Coca-cola", Description: "Bebida", Price: 6.99},
	{ID: "5", Name: "Guarana", Description: "Bebida", Price: 5.99},
}

func TestGetProducts(t *testing.T) {
	testCases := []struct {
		name           string
		page_index     int
		page_size      int
		expectedResult []*domain.Product
	}{
		{"get all five", 1, 5, products[:5]},
		{"get first two", 1, 2, products[:2]},
		{"get second two", 2, 2, products[2:4]},
	}

	repo := mocks.NewProductRepo()
	svc := NewProductService(repo)

	repo.On("Find", 1, 5).
		Return(products, int64(5), nil)

	repo.On("Find", 1, 2).
		Return(products[:2], int64(5), nil)

	repo.On("Find", 2, 2).
		Return(products[2:4], int64(5), nil)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, _, err := svc.GetProducts(tc.page_index, tc.page_size)
			assertNoError(t, err)

			if len(got) != tc.page_size {
				t.Errorf("want size %d got size %d", tc.page_size, len(got))
			}

			if !reflect.DeepEqual(got, tc.expectedResult) {
				t.Errorf("got %v, want %v", got, tc.expectedResult)
			}
		})
	}
}

func TestGetProductByID(t *testing.T) {
	testCases := []struct {
		name        string
		id          string
		expectedErr error
	}{
		{"find product with id 1", "1", nil},
		{"product does not exist", "20", ErrProductNotFound},
	}

	repo := mocks.NewProductRepo()
	svc := NewProductService(repo)

	repo.On("FindByID", "1").
		Return(products[0], nil)
	repo.On("FindByID", "20").
		Return(products[0], ErrProductNotFound)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := svc.GetProductByID(tc.id)

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
				ID:          "10",
				Name:        "Macarrao",
				Description: "Comida",
				Price:       1,
			}, nil,
		},
		{
			"create existing product", &domain.Product{
				ID:          "1",
				Name:        "Arroz-branco",
				Description: "Comida",
				Price:       2,
			}, ErrProductAlreadyExists,
		},
		{"empty product", &domain.Product{}, ErrProductIsEmpty},
		{
			"invalid price", &domain.Product{
				ID:          "20",
				Name:        "Carne",
				Description: "Comida",
				Price:       -2.5,
			}, ErrProductPriceInvalid,
		},
	}

	repo := mocks.NewProductRepo()
	svc := NewProductService(repo)

	repo.On("FindByID", "1").
		Return(products[0], nil)
	repo.On("FindByID", "10").
		Return(&domain.Product{}, ErrProductNotFound)
	repo.On("Save", testCases[0].product).
		Return(testCases[0].product, nil)

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
