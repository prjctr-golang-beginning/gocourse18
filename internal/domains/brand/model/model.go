package model

import (
	"github.com/google/uuid"
	"gocourse18/internal/core/db"
	"time"
)

func NewBrand() *Brand {
	res := &Brand{Payload: db.NewPayload()}

	return res
}

type Brand struct {
	ID        uuid.UUID  `db:"id" table_name:"brands"`
	Name      string     `db:"name"`
	Code      *string    `db:"code"`
	Alias     *string    `db:"alias"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`

	db.Payload
}

func (b Brand) PrimaryKey() db.PrimaryKey {
	return db.NewPrimaryKey(b.ID)
}
