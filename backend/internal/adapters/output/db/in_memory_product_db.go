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

func (i *InMemoryProductDB) FindAll() []*domain.Product {
	return i.store
}

func (i *InMemoryProductDB) FindByID(id int) (*domain.Product, error) {
	for _, product := range i.store {
		if product.ID == id {
			return product, nil
		}
	}

	return nil, errors.New("could not find the requested product")
}

func (i *InMemoryProductDB) Create(product *domain.Product) error {
	i.lastID++
	product.ID = i.lastID

	i.store = append(i.store, product)
	return nil
}

func (i *InMemoryProductDB) Update(product *domain.Product) error {
	found, err := i.FindByID(product.ID)
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

func (i *InMemoryProductDB) Delete(id int) error {
	for index, product := range i.store {
		if product.ID == id {
			i.store = append(i.store[:index], i.store[index+1:]...)
			return nil
		}
	}

	return errors.New("could not find the requested product")
}
