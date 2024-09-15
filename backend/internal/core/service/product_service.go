package service

import (
	"github.com/caio86/nunes-sports/backend/internal/core/domain"
	"github.com/caio86/nunes-sports/backend/internal/core/ports"
)

type ProductServiceErr string

const (
	ErrProductIsEmpty             = ProductServiceErr("empty product received")
	ErrProductNameRequired        = ProductServiceErr("product name is required")
	ErrProductDescriptionRequired = ProductServiceErr("product description is required")
	ErrProductPriceInvalid        = ProductServiceErr("product price must be greater than zero")
	ErrProductNotFound            = ProductServiceErr("product not found")
	ErrProductAlreadyExists       = ProductServiceErr("product already exists")
)

func (e ProductServiceErr) Error() string {
	return string(e)
}

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetProducts(page, limit int) ([]*domain.Product, error) {
	products, err := s.repo.Find(page, limit)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) CreateProduct(product *domain.Product) (*domain.Product, error) {
	if err := validateProduct(product); err != nil {
		return nil, err
	}

	switch _, err := s.GetProductByID(product.ID); err {
	case ErrProductNotFound:
		break
	case nil:
		return nil, ErrProductAlreadyExists
	default:
		return nil, err
	}

	got, err := s.repo.Save(product)
	if err != nil {
		return nil, err
	}

	return got, nil
}

func (s *ProductService) GetProductByID(id uint) (*domain.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, ErrProductNotFound
	}

	return product, nil
}

func validateProduct(product *domain.Product) error {
	if *product == (domain.Product{}) {
		return ErrProductIsEmpty
	}

	if product.Name == "" {
		return ErrProductNameRequired
	}

	if product.Description == "" {
		return ErrProductDescriptionRequired
	}

	if product.Price <= 0 {
		return ErrProductPriceInvalid
	}

	return nil
}
