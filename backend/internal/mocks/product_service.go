package mocks

import (
	"github.com/caio86/nunes-sports/backend/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type productService struct {
	mock.Mock
}

func NewProductService() *productService {
	return &productService{}
}

func (m *productService) GetProducts(page_index, page_size int) ([]*domain.Product, int64, error) {
	args := m.Called(page_index, page_size)
	return args.Get(0).([]*domain.Product), args.Get(1).(int64), args.Error(2)
}

func (m *productService) GetProductByID(id string) (*domain.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *productService) CreateProduct(product *domain.Product) (*domain.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *productService) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *productService) DeleteProduct(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
