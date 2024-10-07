package service

import (
	"github.com/caio86/nunes-sports/backend/internal/core/domain"
	"github.com/caio86/nunes-sports/backend/internal/core/ports"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetProducts(page_index, page_size int) ([]*domain.Product, int64, error) {
	if page_index <= 0 || page_size < 0 {
		return nil, 0, domain.ErrProductInvalidPagination
	}

	products, total, err := s.repo.Find(page_index, page_size)
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (s *ProductService) CreateProduct(product *domain.Product) (*domain.Product, error) {
	if err := product.Validate(); err != nil {
		return nil, err
	}

	switch _, err := s.GetProductByID(product.ID); err {
	case domain.ErrProductNotFound:
		break
	case nil:
		return nil, domain.ErrProductAlreadyExists
	default:
		return nil, err
	}

	got, err := s.repo.Save(product)
	if err != nil {
		return nil, err
	}

	return got, nil
}

func (s *ProductService) GetProductByID(id string) (*domain.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, domain.ErrProductNotFound
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(id string) error {
	if _, err := s.GetProductByID(id); err != nil {
		return err
	}

	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
