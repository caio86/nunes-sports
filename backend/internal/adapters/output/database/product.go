package database

import (
	"github.com/caio86/nunes-sports/backend/internal/core/domain"
	"github.com/caio86/nunes-sports/backend/internal/core/ports"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ports.ProductRepository {
	db.AutoMigrate(&domain.Product{})
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Find(page, limit int) ([]*domain.Product, int64, error) {
	var products []*domain.Product
	var total int64
	var result *gorm.DB

	r.db.Model(&domain.Product{}).Count(&total)

	if limit == 0 {
		result = r.db.Find(&products)
	} else {
		offset := (page - 1) * limit
		result = r.db.
			Limit(limit).
			Offset(offset).
			Find(&products)
	}

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return products, total, nil
}

func (r *productRepository) FindByID(id string) (*domain.Product, error) {
	var product *domain.Product

	result := r.db.First(&product, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *productRepository) Save(product *domain.Product) (*domain.Product, error) {
	result := r.db.Create(product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *productRepository) Update(product *domain.Product) (*domain.Product, error) {
	result := r.db.Save(product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *productRepository) Delete(id string) error {
	result := r.db.Delete(&domain.Product{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
