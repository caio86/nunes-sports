package service

import (
	"github.com/caio86/nunes-sports/backend/internal/core/domain"
	"github.com/caio86/nunes-sports/backend/internal/core/ports"
)

type ProductServiceImpl struct {
	repo ports.ProductRepository
}

func NewProductServiceImpl(repo ports.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		repo: repo,
	}
}

func (p *ProductServiceImpl) GetProducts() []*domain.Product {
	return p.repo.FindAll()
}

func (p *ProductServiceImpl) GetProductByID(id int) (*domain.Product, error) {
	return p.repo.FindByID(id)
}

func (p *ProductServiceImpl) CreateProduct(product *domain.Product) error {
	return p.repo.Create(product)
}

func (p *ProductServiceImpl) UpdateProduct(product *domain.Product) error {
	return p.repo.Update(product)
}

func (p *ProductServiceImpl) DeleteProduct(id int) error {
	return p.repo.Delete(id)
}
