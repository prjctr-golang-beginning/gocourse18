package adapter

import (
	"context"
	"gocourse18/internal/core/db"
	"gocourse18/internal/domains/brand"
	"gocourse18/internal/domains/brand/model"
)

func NewBrandService(repo *brand.Repository) brand.Service {
	return &service{
		repository: repo,
		fields:     repo.Schema().Fields(),
	}
}

type service struct {
	repository *brand.Repository
	fields     []string
}

func (s *service) Create(ctx context.Context, entity *model.Brand) (db.PrimaryKey, error) {
	pk, err := s.repository.CreateOne(ctx, entity)
	if err != nil {
		return nil, err
	}

	return pk, nil
}

func (s *service) GetOne(ctx context.Context, pk db.PrimaryKey) (*model.Brand, error) {
	return s.repository.FindOne(ctx, s.fields, pk)
}
