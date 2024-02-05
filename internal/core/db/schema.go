package db

import (
	"context"
	"fmt"
	"gocourse18/internal/core/db/nosql/cache"
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
	var tags []string

	// Отримуємо тип переданої змінної
	t := reflect.TypeOf(s.entity)

	// Перевіряємо, чи val є покажчиком, і отримуємо тип, на який він вказує
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Переконуємося, що ми працюємо зі структурою
	if t.Kind() != reflect.Struct {
		fmt.Println("Provided value is not a struct!")
		return tags
	}

	// Проходимо по всіх полях структури
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Отримуємо значення JSON тега
		tag := field.Tag.Get("db")
		if tag != "" && tag != "-" {
			tags = append(tags, tag)
		}
	}

	return tags
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

	_ = cache.Store(context.Background(), fmt.Sprintf(`%T_tablename`, s.entity), val, 0)
}

func (s *TableSchema) TableName() string {
	name, _ := cache.Load(context.Background(), fmt.Sprintf(`%T_tablename`, s.entity))
	return name
}
