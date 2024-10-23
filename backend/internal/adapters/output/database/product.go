package database

import (
	"context"

	"github.com/caio86/nunes-sports/backend/internal/adapters/output/database/repository"
	"github.com/caio86/nunes-sports/backend/internal/core/domain"
	"github.com/caio86/nunes-sports/backend/internal/core/ports"
	"github.com/jackc/pgx/v5"
)

type productRepository struct {
	conn *pgx.Conn
}

func NewProductRepository(conn *pgx.Conn) ports.ProductRepository {
	return &productRepository{
		conn: conn,
	}
}

func (r *productRepository) Find(offset, limit int) ([]*domain.Product, int64, error) {
	var result []repository.Product
	var err error

	ctx := context.Background()
	repo := repository.New(r.conn)

	total, err := repo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	if limit == 0 {
		result, err = repo.FindAll(ctx)
	} else {
		offset := (offset - 1) * limit
		result, err = repo.
			Find(ctx, repository.FindParams{
				Limit:  int32(limit),
				Offset: int32(offset),
			})
	}
	if err != nil {
		return nil, 0, err
	}

	products := make([]*domain.Product, len(result))

	for i, v := range result {
		product := &domain.Product{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
		}
		products[i] = product
	}

	return products, total, nil
}

func (r *productRepository) FindByID(id string) (*domain.Product, error) {
	ctx := context.Background()
	repo := repository.New(r.conn)
	result, err := repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return (*domain.Product)(&result), nil
}

func (r *productRepository) Save(product *domain.Product) (*domain.Product, error) {
	ctx := context.Background()
	repo := repository.New(r.conn)
	result, err := repo.Create(ctx, repository.CreateParams{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	})
	if err != nil {
		return nil, err
	}

	return (*domain.Product)(&result), nil
}

func (r *productRepository) Update(product *domain.Product) (*domain.Product, error) {
	ctx := context.Background()
	repo := repository.New(r.conn)
	result, err := repo.Update(ctx, repository.UpdateParams{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	})
	if err != nil {
		return nil, err
	}

	return (*domain.Product)(&result), nil
}

func (r *productRepository) Delete(id string) error {
	ctx := context.Background()
	repo := repository.New(r.conn)
	err := repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
