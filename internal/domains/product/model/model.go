package model

import (
	"github.com/google/uuid"
	"gocourse18/internal/core/db"
	"time"
)

func NewProduct() *Product {
	res := &Product{}
	res.Payload = db.NewPayload(res)

	return res
}

type Product struct {
	ID        uuid.UUID  `json:"id" table_name:"products"`
	BrandID   uuid.UUID  `json:"brand_id"`
	Status    Status     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	InsertedAt time.Time `json:"inserted_at"`

	db.Payload
}

func (p Product) PrimaryKey() db.PrimaryKey {
	return db.NewPrimaryKey(p.ID)
}
