package brand

import (
	"context"
	"gocourse18/internal/core/db"
	"gocourse18/internal/domains/brand/model"
)

type Service interface {
	Create(context.Context, *model.Brand) (db.PrimaryKey, error)
	GetOne(context.Context, db.PrimaryKey) (*model.Brand, error)
}
