package db

type Entity interface {
	Body() map[string]any
	PrimaryKey() PrimaryKey
}

type Payload struct {
	payload     map[string]any
	initialized bool
}

func (p *Payload) Add(field string, val any) {
	if !p.initialized {
		panic("Payload must be initialized with NewPayload(...)")
	}

	p.payload[field] = val
}

func (p *Payload) Body() map[string]any {
	return p.payload
}

func NewPayload() Payload {
	payload := Payload{
		payload:     make(map[string]any),
		initialized: true,
	}

	return payload
}
