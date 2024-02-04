package adapter

import (
	"context"
	"gocourse18/internal/core/db"
	"gocourse18/internal/domains/product"
	"gocourse18/internal/domains/product/model"
)

func NewProductService(repo *product.Repository) product.Service {
	return &service{
		repository: repo,
		fields:     repo.Schema().Fields(),
	}
}

type service struct {
	repository *product.Repository
	fields     []string
}

func (s *service) Create(ctx context.Context, entity *model.Product) (db.PrimaryKey, error) {
	pk, err := s.repository.CreateOne(ctx, entity)
	if err != nil {
		return nil, err
	}

	return pk, nil
}

func (s *service) GetOne(ctx context.Context, pk db.PrimaryKey) (*model.Product, error) {
	return s.repository.FindOne(ctx, s.fields, pk)
}
