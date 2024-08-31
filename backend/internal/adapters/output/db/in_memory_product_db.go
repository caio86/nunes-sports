package db

import (
	"errors"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type InMemoryProductDB struct {
	store  []*domain.Product
	lastID int
}

func NewInMemoryProductDB() *InMemoryProductDB {
	return &InMemoryProductDB{
		store:  make([]*domain.Product, 0),
		lastID: 0,
	}
}

func (i *InMemoryProductDB) FindAllProducts() []*domain.Product {
	return i.store
}

func (i *InMemoryProductDB) FindProductByID(id int) (*domain.Product, error) {
	for _, product := range i.store {
		if product.ID == id {
			return product, nil
		}
	}

	return nil, errors.New("could not find the requested product")
}

func (i *InMemoryProductDB) CreateProduct(product *domain.Product) error {
	i.lastID++
	product.ID = i.lastID

	i.store = append(i.store, product)
	return nil
}

func (i *InMemoryProductDB) UpdateProduct(product *domain.Product) error {
	found, err := i.FindProductByID(product.ID)
	if err != nil {
		return err
	}

	if product.Name != "" {
		found.Name = product.Name
	}

	if product.Description != "" {
		found.Description = product.Description
	}

	if product.Price != 0 {
		found.Price = product.Price
	}

	return nil
}
