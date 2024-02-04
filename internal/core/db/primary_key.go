package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type PKType string

const (
	PkTypeID      PKType = "ID"
	PkTypeComplex PKType = "Complex"
	KeyID         string = "id"
)

type PrimaryKeyable interface {
	PrimaryKey() PrimaryKey
}

type PrimaryKey interface {
	IsComplex() bool
	Type() PKType
	UUID() (uuid.UUID, bool)
	Fields() []string
	Values() []any
}

type PrimaryKeySrc interface {
	int | int64 | string | uuid.UUID | map[string]any
}

func NewPrimaryKey[S PrimaryKeySrc](key S) PrimaryKey {
	switch val := any(key).(type) {
	case int, int64, string, uuid.UUID:
		return &primaryKey{
			tp:      PkTypeID,
			key:     map[string]any{KeyID: val},
			_fields: []string{KeyID},
			_values: []any{&val},
		}
	case map[string]any:
		pk := &primaryKey{
			tp:      PkTypeComplex,
			key:     val,
			_fields: make([]string, 0, len(val)),
			_values: make([]any, 0, len(val)),
		}

		for field, value := range val {
			pk._fields = append(pk._fields, field)
			pk._values = append(pk._values, &value)
		}

		return pk
	}

	return nil
}

type primaryKey struct {
	tp      PKType
	key     sq.Eq
	_fields []string
	_values []any
}

func (s *primaryKey) Type() PKType {
	return s.tp
}

func (s *primaryKey) Fields() []string {
	return s._fields
}

func (s *primaryKey) Values() []any {
	return s._values
}

func (s *primaryKey) IsComplex() bool {
	return s.Type() == PkTypeComplex
}

func (s *primaryKey) IsID() bool {
	return s.Type() == PkTypeID
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
