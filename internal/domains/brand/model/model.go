package model

import (
	"github.com/google/uuid"
	"gocourse18/internal/core/db"
	"time"
)

func NewBrand() *Brand {
	res := &Brand{}
	res.Payload = db.NewPayload(res)

	return res
}

type Brand struct {
	ID        uuid.UUID  `json:"id" table_name:"brands"`
	Name      string     `json:"name"`
	Code      string     `json:"code"`
	Alias     string     `json:"alias"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	db.Payload
}

func (b Brand) PrimaryKey() db.PrimaryKey {
	return db.NewPrimaryKey(b.ID)
}
