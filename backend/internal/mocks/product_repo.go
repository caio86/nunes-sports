package mocks

import (
	"github.com/caio86/nunes-sports/backend/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type productRepo struct {
	mock.Mock
}

func NewProductRepo() *productRepo {
	return &productRepo{}
}

func (m *productRepo) Find(offset, limit int) ([]*domain.Product, int64, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]*domain.Product), args.Get(1).(int64), args.Error(2)
}

func (m *productRepo) FindByID(id string) (*domain.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *productRepo) Save(product *domain.Product) (*domain.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *productRepo) Update(product *domain.Product) (*domain.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *productRepo) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
