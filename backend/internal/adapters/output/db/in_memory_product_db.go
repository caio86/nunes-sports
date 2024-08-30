package db

import (
	"errors"

	"github.com/caio86/nunes-sports/backend/internal/core/domain"
)

type InMemoryProductDB struct {
	store  map[int]*domain.Product
	lastID int
}

func NewInMemoryProductDB() *InMemoryProductDB {
	return &InMemoryProductDB{
		store:  make(map[int]*domain.Product),
		lastID: 0,
	}
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
		return nil, errors.New("could not find the requested product")
	}
	return product, nil
}

func (i *InMemoryProductDB) CreateProduct(product *domain.Product) error {
	i.lastID++
	product.ID = i.lastID

	_, err := i.FindProductByID(product.ID)
	if err == nil {
		return errors.New("product already exists")
	}

	i.store[product.ID] = product
	return nil
}
