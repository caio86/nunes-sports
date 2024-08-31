package ports

import "github.com/caio86/nunes-sports/backend/internal/core/domain"

type ProductRepository interface {
	FindAll() []*domain.Product
	FindByID(id int) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id int) error
}

type ProductService interface {
	GetProducts() []*domain.Product
	GetProductByID(id int) (*domain.Product, error)
	CreateProduct(product *domain.Product) error
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id int) error
}
