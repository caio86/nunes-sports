package ports

import "github.com/caio86/nunes-sports/backend/internal/core/domain"

type ProductRepository interface {
	FindAllProducts() []*domain.Product
	FindProductByID(id int) (*domain.Product, error)
	CreateProduct(product *domain.Product) error
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id int) error
}
