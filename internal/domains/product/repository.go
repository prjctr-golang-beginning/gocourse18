package product

import (
	"gocourse18/internal/core/db"
	"gocourse18/internal/core/db/sql"
	"gocourse18/internal/domains/product/model"
	"sync"
)

var (
	schema *db.TableSchema
	once   sync.Once
)

func Schema() *db.TableSchema {
	once.Do(func() {
		schema = db.NewTableSchema(model.Product{})
	})

	return schema
}

func NewProductRepository(pool sql.Conn) *Repository {
	return &Repository{
		Repository: sql.NewRepository[model.Product](pool, Schema()),
	}
}

type Repository struct {
	sql.Repository[model.Product]
}
