package domain

import "strconv"

type ProductErr string

const (
	ErrProductInvalidPagination   = ProductErr("invalid pagination")
	ErrProductIsEmpty             = ProductErr("empty product received")
	ErrProductIDRequired          = ProductErr("product ID is required")
	ErrProductIDInvalid           = ProductErr("product ID is invalid")
	ErrProductNameRequired        = ProductErr("product name is required")
	ErrProductDescriptionRequired = ProductErr("product description is required")
	ErrProductPriceInvalid        = ProductErr("product price must be greater than zero")
	ErrProductNotFound            = ProductErr("product not found")
	ErrProductAlreadyExists       = ProductErr("product already exists")
)

func (e ProductErr) Error() string {
	return string(e)
}

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
}

func (p *Product) Validate() error {
	if *p == (Product{}) {
		return ErrProductIsEmpty
	}

	if p.ID == "" {
		return ErrProductIDRequired
	}

	if _, err := strconv.Atoi(p.ID); err != nil {
		return ErrProductIDInvalid
	}

	if p.Name == "" {
		return ErrProductNameRequired
	}

	if p.Price <= 0 {
		return ErrProductPriceInvalid
	}

	return nil
}
