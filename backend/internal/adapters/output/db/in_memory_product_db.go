package db

import (
	"fmt"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type InMemoryProductDB struct {
	store map[int]*domain.Product
}

func (i *InMemoryProductDB) FindAllProducts() []*domain.Product {
	products := make([]*domain.Product, 0, len(i.store))
	for _, product := range i.store {
		products = append(products, product)
	}
	return products
}

func (i *InMemoryProductDB) FindProductByID(id int) (*domain.Product, error) {
	product, ok := i.store[id]
	if !ok {
		return nil, fmt.Errorf("product with id %d not found", id)
	}
	return product, nil
}
