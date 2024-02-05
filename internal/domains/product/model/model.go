package model

import (
	"github.com/google/uuid"
	"gocourse18/internal/core/db"
	"time"
)

func NewProduct() *Product {
	res := &Product{Payload: db.NewPayload()}

	return res
}

type Product struct {
	ID        uuid.UUID  `db:"id" table_name:"products"`
	BrandID   uuid.UUID  `db:"brand_id"`
	Status    *Status    `db:"status"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`

	db.Payload
}

func (p Product) PrimaryKey() db.PrimaryKey {
	return db.NewPrimaryKey(p.ID)
}
