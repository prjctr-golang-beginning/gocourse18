package brand

import (
	"gocourse18/internal/core/db"
	"gocourse18/internal/core/db/sql"
	"gocourse18/internal/domains/brand/model"
	"sync"
)

var (
	schema *db.TableSchema
	once   sync.Once
)

func Schema() *db.TableSchema {
	once.Do(func() {
		schema = db.NewTableSchema(&model.Brand{})
	})

	return schema
}

func NewBrandRepository(pool sql.Conn) *Repository {
	return &Repository{
		Repository: sql.NewRepository[model.Brand](pool, Schema()),
	}
}

type Repository struct {
	sql.Repository[model.Brand]
}
