package db

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"reflect"
)

type Entity interface {
	Body() map[string]any
	PrimaryKey() PrimaryKey
}

type BodyEmptyer interface {
	IsEmpty() bool
	Body() map[string]any
}

type Payload struct {
	payload map[string]any
	initial []byte
	dirty   []string

	initialized bool
}

func (p *Payload) Add(field string, val any) {
	if !p.initialized {
		panic("Payload must be initialized with NewPayload(...)")
	}

	p.payload[field] = val
	p.dirty = append(p.dirty, field)
}

func (p *Payload) Delete(field string) {
	if !p.initialized {
		panic("Payload must be initialized with NewPayload(...)")
	}

	delete(p.payload, field)
}

func (p *Payload) IsInitialized() bool {
	return p.initialized
}

func (p *Payload) IsEmpty() bool {
	return len(p.payload) == 0
}

func (p *Payload) GetInitialEntity() []byte {
	return p.initial
}

func (p *Payload) GetDirty() []string {
	return p.dirty
}

func (p *Payload) Body() map[string]any {
	return p.payload
}

func NewPayload(entity any) Payload {
	payload := Payload{
		payload:     make(map[string]any),
		initial:     nil,
		dirty:       make([]string, 0),
		initialized: true,
	}

	if entity != nil && !reflect.ValueOf(entity).IsNil() {
		var err error
		payload.initial, err = json.Marshal(entity)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return payload
}

type IdHolder interface {
	Id() uuid.UUID
}
