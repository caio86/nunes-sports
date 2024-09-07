package service

import (
	"strconv"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
	"github.com/caio86/nunes-sports/backend/internal/core/ports"
)

type ProductServiceErr string

const (
	ErrProductCodeInvalid         = ProductServiceErr("product code must be an int")
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

func (s *ProductService) CreateProduct(product *domain.Product) (*domain.Product, error) {
	if err := validateProduct(product); err != nil {
		return nil, err
	}

	switch _, err := s.GetProductByCode(product.Code); err {
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

func (s *ProductService) GetProductByCode(code string) (*domain.Product, error) {
	product, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, ErrProductNotFound
	}

	return product, nil
}

func validateProductCode(code string) error {
	if _, err := strconv.ParseInt(code, 10, 64); err != nil {
		return ErrProductCodeInvalid
	}

	return nil
}

func validateProduct(product *domain.Product) error {
	if err := validateProductCode(product.Code); err != nil {
		return err
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
