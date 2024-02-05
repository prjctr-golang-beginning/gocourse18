package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type PKType string

const (
	KeyID string = "id"
)

type PrimaryKey interface {
	UUID() (uuid.UUID, bool)
	Fields() []string
	OnlyEq() sq.Eq
}

type PrimaryKeySrc interface {
	int | int64 | string | uuid.UUID | map[string]any
}

func NewPrimaryKey[S PrimaryKeySrc](key S) PrimaryKey {
	switch val := any(key).(type) {
	case int, int64, string, uuid.UUID:
		return &primaryKey{
			key:     map[string]any{KeyID: val},
			_fields: []string{KeyID},
		}
	}

	return nil
}

type primaryKey struct {
	key     sq.Eq
	_fields []string
}

func (s *primaryKey) Fields() []string {
	return s._fields
}

func (s *primaryKey) UUID() (uuid.UUID, bool) {
	if id, ok := s.key[KeyID]; ok {
		if res, ok := id.(uuid.UUID); ok {
			return res, true
		}
	}

	return [16]byte{}, false
}

func (s *primaryKey) OnlyEq() sq.Eq {
	return s.key
}
