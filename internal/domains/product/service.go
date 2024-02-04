package product

import (
	"context"
	"gocourse18/internal/core/db"
	"gocourse18/internal/domains/product/model"
)

type Service interface {
	Create(context.Context, *model.Product) (db.PrimaryKey, error)
	GetOne(context.Context, db.PrimaryKey) (*model.Product, error)
}
