package ports

import (
	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type ProductRepository interface {
	Find(offset, limit uint) ([]*domain.Product, error)
	FindByCode(code string) (*domain.Product, error)
	Save(product *domain.Product) (*domain.Product, error)
	Update(product *domain.Product) (*domain.Product, error)
	Delete(code string) error
}

type ProductService interface {
	GetProducts(page_index, page_size uint) ([]*domain.Product, error)
	GetProductByCode(code string) (*domain.Product, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
}
