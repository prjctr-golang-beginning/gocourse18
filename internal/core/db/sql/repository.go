package sql

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"gocourse18/internal/core/db"
	"gocourse18/internal/core/helpers"
	"log"
	"strconv"
	"strings"
)

type Repository[E any] interface {
	Schema() db.Schema
	CreateOne(ctx context.Context, entity db.Entity) (db.PrimaryKey, error)
	FindOne(ctx context.Context, fields []string, pk db.PrimaryKey) (*E, error)
}

type Conn interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
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

	b := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(strconv.Quote(s.Schema().TableName())).
		Columns(helpers.EscapeColumns(columns)...).
		Values(vals...).
		Suffix(`RETURNING`).
		Suffix(strings.Join(entity.PrimaryKey().Fields(), `, `))

	query, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	res := new(E)
	rows, err := s.pool.QueryContext(ctx, query, args)
	if err != nil {
		return nil, err
	}
	e := new(E)
	for rows.Next() {
		err = rows.Scan(e)
		if err != nil {
			log.Fatal(err)
		}
	}
	rows.Close()

	return any(res).(db.PrimaryKeyable).PrimaryKey(), nil
}

func (s *repository[E]) FindOne(ctx context.Context, fields []string, pk db.PrimaryKey) (*E, error) {
	b := sq.
		Select(helpers.EscapeColumns(fields)...).
		From(strconv.Quote(s.Schema().TableName())).
		Where(pk.Fields(), pk.Values()).
		Limit(1)

	query, args, bErr := b.ToSql()
	if bErr != nil {
		return nil, bErr
	}

	rows, err := s.pool.QueryContext(ctx, query, args)
	if err != nil {
		return nil, err
	}
	e := new(E)
	for rows.Next() {
		err = rows.Scan(e)
		if err != nil {
			log.Fatal(err)
		}
	}
	rows.Close()

	return e, nil
}
