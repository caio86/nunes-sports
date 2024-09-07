package ports

import (
	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type ProductRepository interface {
	FindByCode(code string) (*domain.Product, error)
	Save(product *domain.Product) (*domain.Product, error)
}
