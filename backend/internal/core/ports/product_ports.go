package ports

import (
	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type ProductRepository interface {
	Find() ([]*domain.Product, error)
	FindByCode(code string) (*domain.Product, error)
	Save(product *domain.Product) (*domain.Product, error)
}

type ProductService interface {
	GetProducts() ([]*domain.Product, error)
	GetProductByCode(code string) (*domain.Product, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
}
