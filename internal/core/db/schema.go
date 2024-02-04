package db

import (
	"log"
	"reflect"
)

type Schema interface {
	TableName() string
	Fields() []string
}

type TableSchema struct {
	entity     any
	primaryKey []string

	escapeColumns bool

	_fields []string
	_table  string
}

func NewTableSchema(entity any) *TableSchema {
	res := &TableSchema{
		entity:        entity,
		escapeColumns: true,
	}

	res.cacheTableName()

	return res
}

func (s *TableSchema) Fields() []string {
	return []string{}
}

func (s *TableSchema) cacheTableName() {
	v := reflect.TypeOf(s.entity)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	f := v.Field(0)
	val, ok := f.Tag.Lookup("table_name")
	if !ok {
		log.Fatalf("Table name for entity %s not defined")
	}

	s._table = val
}

func (s *TableSchema) TableName() string {
	return s._table
}
