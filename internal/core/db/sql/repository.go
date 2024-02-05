package sql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"gocourse18/internal/core/db"
	"gocourse18/internal/core/helpers"
	"log"
)

type Repository[E any] interface {
	Schema() db.Schema
	CreateOne(ctx context.Context, entity db.Entity) (db.PrimaryKey, error)
	FindOne(ctx context.Context, fields []string, pk db.PrimaryKey) (*E, error)
}

type Conn interface {
	sqlx.QueryerContext
}

func NewRepository[E any](pool Conn, schema db.Schema) Repository[E] {
	return &repository[E]{
		pool:   pool,
		schema: schema,
	}
}

type WithExists struct {
	Exists bool `json:"exists"`
}

type repository[E any] struct {
	schema db.Schema
	pool   Conn
}

func (s *repository[E]) Schema() db.Schema {
	return s.schema
}

func (s *repository[E]) CreateOne(ctx context.Context, entity db.Entity) (db.PrimaryKey, error) {
	columns, vals := helpers.SplitMap(entity.Body())

	b := sq.StatementBuilder.
		Insert(s.Schema().TableName()).
		Columns(columns...).
		Values(vals...)

	query, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = s.pool.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return entity.PrimaryKey(), nil
}

func (s *repository[E]) FindOne(ctx context.Context, fields []string, pk db.PrimaryKey) (*E, error) {
	b := sq.
		Select(fields...).
		From(s.Schema().TableName()).
		Where(pk.OnlyEq()).
		Limit(1)

	query, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.pool.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	e := new(E)
	for rows.Next() {
		err = rows.StructScan(e)
		if err != nil {
			log.Fatal(err)
		}
	}
	rows.Close()

	return e, nil
}
