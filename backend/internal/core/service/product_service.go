package service

import (
	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type ProductServiceErr string

const (
	ErrProductNameRequired        = ProductServiceErr("product name is required")
	ErrProductDescriptionRequired = ProductServiceErr("product description is required")
	ErrProductPriceInvalid        = ProductServiceErr("product price must be greater than zero")
	ErrProductNotFound            = ProductServiceErr("product not found")
	ErrProductAlreadyExists       = ProductServiceErr("product already exists")
)

func (e ProductServiceErr) Error() string {
	return string(e)
}

func CreateProduct(product *domain.Product) (*domain.Product, error) {
	if err := validateProduct(product); err != nil {
		return nil, err
	}

	return product, nil
}

func validateProduct(product *domain.Product) error {
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
